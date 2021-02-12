# Makefile

.PHONY: default
default: proto

.PHONY: proto
proto:
	cd proto && \
	protoc \
		-I. \
		-Ivendor \
		-Ivendor/github.com/protocolbuffers/protobuf/src \
		--go_out=plugins=grpc,paths=source_relative:. \
		--govalidators_out=paths=source_relative:. \
		testsize.proto

.PHONY: run
run:
	go run ./cmd/testsize/main.go

# Expect to see:
# proto/testsize.validator.pb.go:20:11: this.Size_ undefined (type *ThingWithSize has no field or method Size_)
# proto/testsize.validator.pb.go:21:123: this.Size_ undefined (type *ThingWithSize has no field or method Size_)
