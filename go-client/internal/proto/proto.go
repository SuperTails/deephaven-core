package proto

//go:generate -command protoc-gen-go protoc --go_out=../.. --go_opt=module=github.com/deephaven/deephaven-core/go-client --experimental_allow_proto3_optional -I../../../proto/proto-backplane-grpc/src/main/proto

//go:generate protoc-gen-go deephaven/proto/application.proto
//go:generate protoc-gen-go deephaven/proto/console.proto
//go:generate protoc-gen-go deephaven/proto/inputtable.proto
//go:generate protoc-gen-go deephaven/proto/object.proto
//go:generate protoc-gen-go deephaven/proto/session.proto
//go:generate protoc-gen-go deephaven/proto/table.proto
//go:generate protoc-gen-go deephaven/proto/ticket.proto

//go:generate -command protoc-gen-go-grpc protoc --go-grpc_out=../.. --go-grpc_opt=module=github.com/deephaven/deephaven-core/go-client --experimental_allow_proto3_optional -I../../../proto/proto-backplane-grpc/src/main/proto

//go:generate protoc-gen-go-grpc deephaven/proto/application.proto
//go:generate protoc-gen-go-grpc deephaven/proto/console.proto
//go:generate protoc-gen-go-grpc deephaven/proto/inputtable.proto
//go:generate protoc-gen-go-grpc deephaven/proto/object.proto
//go:generate protoc-gen-go-grpc deephaven/proto/session.proto
//go:generate protoc-gen-go-grpc deephaven/proto/table.proto
//go:generate protoc-gen-go-grpc deephaven/proto/ticket.proto
