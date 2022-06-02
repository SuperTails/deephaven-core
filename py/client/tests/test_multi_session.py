

import unittest
from time import sleep

import pyarrow as pa
from pyarrow import csv

from pydeephaven import DHError
from pydeephaven import Session
from pydeephaven.session import SYNC_ONCE, SYNC_REPEATED
from tests.testbase import BaseTestCase


class MultiSessionTestCase(BaseTestCase):
    def test_persistent_tables(self):
        with Session() as session1:
            session1 = Session()
            session1.run_script('t = None')
            t = session1.empty_table(10)
            session1.bind_table('t', t)

        with Session(sync_tables=SYNC_ONCE) as session2:
            self.assertIn('t', session2.tables)

    def test_shared_tables(self):
        session1 = Session(sync_tables=SYNC_REPEATED)
        session1.run_script('t = None')

        session2 = Session()
        t = session2.empty_table(10)
        session2.bind_table('t', t)

        sleep(1)

        self.assertIn('t', session1.tables)
