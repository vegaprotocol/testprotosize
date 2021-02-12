// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: testsize.proto

package proto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *ThingWithSize) Validate() error {
	if !(this.Size_ > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Size_", fmt.Errorf(`value '%v' must be greater than '0'`, this.Size_))
	}
	return nil
}
