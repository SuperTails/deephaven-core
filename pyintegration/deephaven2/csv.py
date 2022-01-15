#
#   Copyright (c) 2016-2021 Deephaven Data Labs and Patent Pending
#
""" The deephaven.csv module supports reading an external CSV file into a Deephaven table and writing a
Deephaven table out as a CSV file.
"""
from enum import Enum
from typing import Dict, Any, List

import jpy

from deephaven2 import DHError
from deephaven2.dtypes import DType
from deephaven2.table import Table

_JCsvSpecs = jpy.get_type("io.deephaven.csv.CsvSpecs")
_JInferenceSpecs = jpy.get_type("io.deephaven.csv.InferenceSpecs")
_JTableHeader = jpy.get_type("io.deephaven.qst.table.TableHeader")
_JCsvTools = jpy.get_type("io.deephaven.csv.CsvTools")


class Inference(Enum):
    """ An Enum of predefined inference specs.

    Inference specifications contains the configuration and logic for inferring an acceptable parser from string values.
    """

    STRINGS = _JInferenceSpecs.strings()
    """ Configured parsers: strings only.
    """

    MINIMAL = _JInferenceSpecs.minimal()
    """ Configured parsers: BOOL, LONG, DOUBLE, INSTANT, STRING.
    """

    STANDARD = _JInferenceSpecs.standard()
    """ Configured parsers: BOOL, INT, LONG, DOUBLE, DATETIME, CHAR, STRING.
    """

    STANDARD_TIMES = _JInferenceSpecs.standardTimes()
    """ Configured parsers: BOOL, DATETIME, CHAR, STRING, SECONDS.
    """

def _build_header(header: Dict[str, DType] = None):
    if not header:
        return None

    table_header_builder = _JTableHeader.builder()
    for k, v in header.items():
        table_header_builder.putHeaders(k, v.qst_type)

    return table_header_builder.build()


def read(path: str,
         header: Dict[str, DType] = None,
         inference: Any = Inference.STANDARD,
         headless: bool = False,
         delimiter: str = ",",
         quote: str = "\"",
         ignore_surrounding_spaces: bool = True,
         trim: bool = False) -> Table:
    """ Read the CSV data specified by the path parameter as a table.

    Args:
        path (str): a file path or a URL string
        header (Dict[str, DType]): a dict to define the table columns with key being the name, value being the data type
        inference (csv.Inference): an Enum value specifying the rules for data type inference, default is STANDARD_TIMES
        headless (bool): indicates if the CSV data is headless, default is False
        delimiter (str): the delimiter used by the CSV, default is the comma
        quote (str): the quote character for the CSV, default is double quote
        ignore_surrounding_spaces (bool): indicates whether surrounding white space should be ignored for unquoted text
            fields, default is True
        trim (bool) : indicates whether to trim white space inside a quoted string, default is False

    Returns:
        a table

    Raises:
        DHError
    """
    try:
        csv_specs_builder = _JCsvSpecs.builder()

        # build the head spec
        table_header = _build_header(header)
        if table_header:
            csv_specs_builder.header(table_header)

        csv_specs = (csv_specs_builder.inference(inference.value)
                     .hasHeaderRow(not headless)
                     .delimiter(ord(delimiter))
                     .quote(ord(quote))
                     .ignoreSurroundingSpaces(ignore_surrounding_spaces)
                     .trim(trim)
                     .build())

        j_table = _JCsvTools.readCsv(path, csv_specs)

        return Table(j_table=j_table)
    except Exception as e:
        raise DHError(e, "read csv failed") from e


def write(table: Table, path: str, cols: List[str] = []) -> None:
    """ Write a table to a standard CSV file.

    Args:
        table (Table): the source table
        path (str): the path of the CSV file
        cols (List[str]): the names of the columns to be written out

    Raises:
        DHError
    """
    try:
        _JCsvTools.writeCsv(table.j_table, False, path, *cols)
    except Exception as e:
        raise DHError("write csv failed.") from e