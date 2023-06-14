// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.6.1
// source: calculator/calculator.proto

package calculator

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operand1 int32 `protobuf:"varint,1,opt,name=Operand1,proto3" json:"Operand1,omitempty"`
	Operand2 int32 `protobuf:"varint,2,opt,name=Operand2,proto3" json:"Operand2,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetOperand1() int32 {
	if x != nil {
		return x.Operand1
	}
	return 0
}

func (x *AddRequest) GetOperand2() int32 {
	if x != nil {
		return x.Operand2
	}
	return 0
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type SubtractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operand1 int32 `protobuf:"varint,1,opt,name=Operand1,proto3" json:"Operand1,omitempty"`
	Operand2 int32 `protobuf:"varint,2,opt,name=Operand2,proto3" json:"Operand2,omitempty"`
}

func (x *SubtractRequest) Reset() {
	*x = SubtractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubtractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubtractRequest) ProtoMessage() {}

func (x *SubtractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubtractRequest.ProtoReflect.Descriptor instead.
func (*SubtractRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *SubtractRequest) GetOperand1() int32 {
	if x != nil {
		return x.Operand1
	}
	return 0
}

func (x *SubtractRequest) GetOperand2() int32 {
	if x != nil {
		return x.Operand2
	}
	return 0
}

type SubtractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *SubtractResponse) Reset() {
	*x = SubtractResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubtractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubtractResponse) ProtoMessage() {}

func (x *SubtractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubtractResponse.ProtoReflect.Descriptor instead.
func (*SubtractResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *SubtractResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type MultplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operand1 int32 `protobuf:"varint,1,opt,name=Operand1,proto3" json:"Operand1,omitempty"`
	Operand2 int32 `protobuf:"varint,2,opt,name=Operand2,proto3" json:"Operand2,omitempty"`
}

func (x *MultplyRequest) Reset() {
	*x = MultplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultplyRequest) ProtoMessage() {}

func (x *MultplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultplyRequest.ProtoReflect.Descriptor instead.
func (*MultplyRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculator_proto_rawDescGZIP(), []int{4}
}

func (x *MultplyRequest) GetOperand1() int32 {
	if x != nil {
		return x.Operand1
	}
	return 0
}

func (x *MultplyRequest) GetOperand2() int32 {
	if x != nil {
		return x.Operand2
	}
	return 0
}

type MultiplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *MultiplyResponse) Reset() {
	*x = MultiplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiplyResponse) ProtoMessage() {}

func (x *MultiplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiplyResponse.ProtoReflect.Descriptor instead.
func (*MultiplyResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculator_proto_rawDescGZIP(), []int{5}
}

func (x *MultiplyResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type DivisionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operand1 int32 `protobuf:"varint,1,opt,name=Operand1,proto3" json:"Operand1,omitempty"`
	Operand2 int32 `protobuf:"varint,2,opt,name=Operand2,proto3" json:"Operand2,omitempty"`
}

func (x *DivisionRequest) Reset() {
	*x = DivisionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivisionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivisionRequest) ProtoMessage() {}

func (x *DivisionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivisionRequest.ProtoReflect.Descriptor instead.
func (*DivisionRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculator_proto_rawDescGZIP(), []int{6}
}

func (x *DivisionRequest) GetOperand1() int32 {
	if x != nil {
		return x.Operand1
	}
	return 0
}

func (x *DivisionRequest) GetOperand2() int32 {
	if x != nil {
		return x.Operand2
	}
	return 0
}

type DivisionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *DivisionResponse) Reset() {
	*x = DivisionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivisionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivisionResponse) ProtoMessage() {}

func (x *DivisionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivisionResponse.ProtoReflect.Descriptor instead.
func (*DivisionResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculator_proto_rawDescGZIP(), []int{7}
}

func (x *DivisionResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_calculator_proto protoreflect.FileDescriptor

var file_calculator_calculator_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x44, 0x0a, 0x0a, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x6e, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x6e, 0x64, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x32, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x32, 0x22,
	0x25, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x49, 0x0a, 0x0f, 0x53, 0x75, 0x62, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x6e, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x6e, 0x64, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64,
	0x32, 0x22, 0x2a, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x48, 0x0a,
	0x0e, 0x4d, 0x75, 0x6c, 0x74, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x32, 0x22, 0x2a, 0x0a, 0x10, 0x4d, 0x75, 0x6c, 0x74, 0x69,
	0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x49, 0x0a, 0x0f, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e,
	0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e,
	0x64, 0x31, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x32, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x32, 0x22, 0x2a,
	0x0a, 0x10, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x9e, 0x02, 0x0a, 0x0a, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x38, 0x0a, 0x03, 0x41, 0x64, 0x64,
	0x12, 0x16, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12,
	0x1b, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x75, 0x62,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x75, 0x62, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x08,
	0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x12, 0x1b,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x69, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2f,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_calculator_calculator_proto_rawDescOnce sync.Once
	file_calculator_calculator_proto_rawDescData = file_calculator_calculator_proto_rawDesc
)

func file_calculator_calculator_proto_rawDescGZIP() []byte {
	file_calculator_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_calculator_proto_rawDescData)
	})
	return file_calculator_calculator_proto_rawDescData
}

var file_calculator_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_calculator_calculator_proto_goTypes = []interface{}{
	(*AddRequest)(nil),       // 0: calculator.AddRequest
	(*AddResponse)(nil),      // 1: calculator.AddResponse
	(*SubtractRequest)(nil),  // 2: calculator.SubtractRequest
	(*SubtractResponse)(nil), // 3: calculator.SubtractResponse
	(*MultplyRequest)(nil),   // 4: calculator.MultplyRequest
	(*MultiplyResponse)(nil), // 5: calculator.MultiplyResponse
	(*DivisionRequest)(nil),  // 6: calculator.DivisionRequest
	(*DivisionResponse)(nil), // 7: calculator.DivisionResponse
}
var file_calculator_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.Calculator.Add:input_type -> calculator.AddRequest
	2, // 1: calculator.Calculator.Subtract:input_type -> calculator.SubtractRequest
	4, // 2: calculator.Calculator.Multiply:input_type -> calculator.MultplyRequest
	6, // 3: calculator.Calculator.Divide:input_type -> calculator.DivisionRequest
	1, // 4: calculator.Calculator.Add:output_type -> calculator.AddResponse
	3, // 5: calculator.Calculator.Subtract:output_type -> calculator.SubtractResponse
	5, // 6: calculator.Calculator.Multiply:output_type -> calculator.MultiplyResponse
	7, // 7: calculator.Calculator.Divide:output_type -> calculator.DivisionResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_calculator_proto_init() }
func file_calculator_calculator_proto_init() {
	if File_calculator_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubtractRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubtractResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultplyRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiplyResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivisionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivisionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_calculator_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_calculator_proto_msgTypes,
	}.Build()
	File_calculator_calculator_proto = out.File
	file_calculator_calculator_proto_rawDesc = nil
	file_calculator_calculator_proto_goTypes = nil
	file_calculator_calculator_proto_depIdxs = nil
}
