cmake_minimum_required(VERSION 3.16)
project(read_table_with_arrow_flight)

set(CMAKE_CXX_STANDARD 17)


add_subdirectory(../../cpp-client/deephaven deephaven_dir)
#find_package(deephaven REQUIRED)

find_package(Arrow REQUIRED)
find_package(ArrowFlight REQUIRED HINTS ${Arrow_DIR})
find_package(Protobuf REQUIRED)
find_package(gRPC REQUIRED)
find_package(Threads REQUIRED)

add_executable(read_table_with_arrow_flight main.cc)

target_link_libraries(read_table_with_arrow_flight deephaven::client)
