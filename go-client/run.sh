/opt/protoc/bin/protoc --plugin=protoc-gen-go=/opt/protoc-gen-go --plugin=protoc-gen-go-grpc=/opt/protoc-gen-go-grpc \
	--go_out=/generated/go --go-grpc_out=/generated/go \
	--go_opt=module=github.com/deephaven/deephaven-core/go-client --go-grpc_opt=module=github.com/deephaven/deephaven-core/go-client \
	-I../includes \
	/includes/deephaven/proto/ticket.proto \
	/includes/deephaven/proto/session.proto \
	/includes/deephaven/proto/table.proto \
	/includes/deephaven/proto/application.proto \
	/includes/deephaven/proto/inputtable.proto \
	/includes/deephaven/proto/partitionedtable.proto

#	../deephaven/proto/console.proto \
#	../deephaven/proto/object.proto \