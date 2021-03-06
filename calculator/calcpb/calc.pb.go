// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: calculator/calcpb/calc.proto

package calcpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CalcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 int32 `protobuf:"varint,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 int32 `protobuf:"varint,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *CalcRequest) Reset() {
	*x = CalcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcRequest) ProtoMessage() {}

func (x *CalcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcRequest.ProtoReflect.Descriptor instead.
func (*CalcRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{0}
}

func (x *CalcRequest) GetNum1() int32 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *CalcRequest) GetNum2() int32 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type CalcResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Answer int32 `protobuf:"varint,1,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *CalcResponse) Reset() {
	*x = CalcResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcResponse) ProtoMessage() {}

func (x *CalcResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcResponse.ProtoReflect.Descriptor instead.
func (*CalcResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{1}
}

func (x *CalcResponse) GetAnswer() int32 {
	if x != nil {
		return x.Answer
	}
	return 0
}

type PrimeDecompRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Integer int32 `protobuf:"varint,1,opt,name=integer,proto3" json:"integer,omitempty"`
}

func (x *PrimeDecompRequest) Reset() {
	*x = PrimeDecompRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeDecompRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeDecompRequest) ProtoMessage() {}

func (x *PrimeDecompRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeDecompRequest.ProtoReflect.Descriptor instead.
func (*PrimeDecompRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{2}
}

func (x *PrimeDecompRequest) GetInteger() int32 {
	if x != nil {
		return x.Integer
	}
	return 0
}

type PrimeDecompResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prime int64 `protobuf:"varint,1,opt,name=prime,proto3" json:"prime,omitempty"`
}

func (x *PrimeDecompResponse) Reset() {
	*x = PrimeDecompResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrimeDecompResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrimeDecompResponse) ProtoMessage() {}

func (x *PrimeDecompResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrimeDecompResponse.ProtoReflect.Descriptor instead.
func (*PrimeDecompResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{3}
}

func (x *PrimeDecompResponse) GetPrime() int64 {
	if x != nil {
		return x.Prime
	}
	return 0
}

type AverageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *AverageRequest) Reset() {
	*x = AverageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageRequest) ProtoMessage() {}

func (x *AverageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageRequest.ProtoReflect.Descriptor instead.
func (*AverageRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{4}
}

func (x *AverageRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type AverageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Average float32 `protobuf:"fixed32,1,opt,name=average,proto3" json:"average,omitempty"`
}

func (x *AverageResponse) Reset() {
	*x = AverageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AverageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AverageResponse) ProtoMessage() {}

func (x *AverageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AverageResponse.ProtoReflect.Descriptor instead.
func (*AverageResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{5}
}

func (x *AverageResponse) GetAverage() float32 {
	if x != nil {
		return x.Average
	}
	return 0
}

type SquareRootRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *SquareRootRequest) Reset() {
	*x = SquareRootRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquareRootRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareRootRequest) ProtoMessage() {}

func (x *SquareRootRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareRootRequest.ProtoReflect.Descriptor instead.
func (*SquareRootRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{6}
}

func (x *SquareRootRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type SquareRootResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumberRoot float64 `protobuf:"fixed64,1,opt,name=number_root,json=numberRoot,proto3" json:"number_root,omitempty"`
}

func (x *SquareRootResponse) Reset() {
	*x = SquareRootResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calcpb_calc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquareRootResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquareRootResponse) ProtoMessage() {}

func (x *SquareRootResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calcpb_calc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquareRootResponse.ProtoReflect.Descriptor instead.
func (*SquareRootResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calcpb_calc_proto_rawDescGZIP(), []int{7}
}

func (x *SquareRootResponse) GetNumberRoot() float64 {
	if x != nil {
		return x.NumberRoot
	}
	return 0
}

var File_calculator_calcpb_calc_proto protoreflect.FileDescriptor

var file_calculator_calcpb_calc_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x35, 0x0a, 0x0b, 0x43, 0x61,
	0x6c, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d,
	0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x75, 0x6d,
	0x32, 0x22, 0x26, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0x2e, 0x0a, 0x12, 0x50, 0x72, 0x69,
	0x6d, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x22, 0x2b, 0x0a, 0x13, 0x50, 0x72, 0x69,
	0x6d, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x6d, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x2b, 0x0a, 0x0f, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x22, 0x2b, 0x0a,
	0x11, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x35, 0x0a, 0x12, 0x53, 0x71,
	0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x6f, 0x6f,
	0x74, 0x32, 0xc2, 0x02, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3a, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x61, 0x6c, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43,
	0x61, 0x6c, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x59, 0x0a,
	0x12, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x50, 0x72, 0x69, 0x6d, 0x65, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x4d, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x41, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x4d, 0x0a, 0x0a, 0x53, 0x71, 0x75, 0x61, 0x72,
	0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x1d, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_calculator_calcpb_calc_proto_rawDescOnce sync.Once
	file_calculator_calcpb_calc_proto_rawDescData = file_calculator_calcpb_calc_proto_rawDesc
)

func file_calculator_calcpb_calc_proto_rawDescGZIP() []byte {
	file_calculator_calcpb_calc_proto_rawDescOnce.Do(func() {
		file_calculator_calcpb_calc_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_calcpb_calc_proto_rawDescData)
	})
	return file_calculator_calcpb_calc_proto_rawDescData
}

var file_calculator_calcpb_calc_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_calculator_calcpb_calc_proto_goTypes = []interface{}{
	(*CalcRequest)(nil),         // 0: calculator.CalcRequest
	(*CalcResponse)(nil),        // 1: calculator.CalcResponse
	(*PrimeDecompRequest)(nil),  // 2: calculator.PrimeDecompRequest
	(*PrimeDecompResponse)(nil), // 3: calculator.PrimeDecompResponse
	(*AverageRequest)(nil),      // 4: calculator.AverageRequest
	(*AverageResponse)(nil),     // 5: calculator.AverageResponse
	(*SquareRootRequest)(nil),   // 6: calculator.SquareRootRequest
	(*SquareRootResponse)(nil),  // 7: calculator.SquareRootResponse
}
var file_calculator_calcpb_calc_proto_depIdxs = []int32{
	0, // 0: calculator.CalcService.Sum:input_type -> calculator.CalcRequest
	2, // 1: calculator.CalcService.PrimeDecomposition:input_type -> calculator.PrimeDecompRequest
	4, // 2: calculator.CalcService.ComputeAverage:input_type -> calculator.AverageRequest
	6, // 3: calculator.CalcService.SquareRoot:input_type -> calculator.SquareRootRequest
	1, // 4: calculator.CalcService.Sum:output_type -> calculator.CalcResponse
	3, // 5: calculator.CalcService.PrimeDecomposition:output_type -> calculator.PrimeDecompResponse
	5, // 6: calculator.CalcService.ComputeAverage:output_type -> calculator.AverageResponse
	7, // 7: calculator.CalcService.SquareRoot:output_type -> calculator.SquareRootResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculator_calcpb_calc_proto_init() }
func file_calculator_calcpb_calc_proto_init() {
	if File_calculator_calcpb_calc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_calcpb_calc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcRequest); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcResponse); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeDecompRequest); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrimeDecompResponse); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageRequest); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AverageResponse); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SquareRootRequest); i {
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
		file_calculator_calcpb_calc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SquareRootResponse); i {
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
			RawDescriptor: file_calculator_calcpb_calc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_calcpb_calc_proto_goTypes,
		DependencyIndexes: file_calculator_calcpb_calc_proto_depIdxs,
		MessageInfos:      file_calculator_calcpb_calc_proto_msgTypes,
	}.Build()
	File_calculator_calcpb_calc_proto = out.File
	file_calculator_calcpb_calc_proto_rawDesc = nil
	file_calculator_calcpb_calc_proto_goTypes = nil
	file_calculator_calcpb_calc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalcServiceClient is the client API for CalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcServiceClient interface {
	// Unary
	Sum(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
	// Server streaming
	PrimeDecomposition(ctx context.Context, in *PrimeDecompRequest, opts ...grpc.CallOption) (CalcService_PrimeDecompositionClient, error)
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalcService_ComputeAverageClient, error)
	SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error)
}

type calcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcServiceClient(cc grpc.ClientConnInterface) CalcServiceClient {
	return &calcServiceClient{cc}
}

func (c *calcServiceClient) Sum(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalcService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) PrimeDecomposition(ctx context.Context, in *PrimeDecompRequest, opts ...grpc.CallOption) (CalcService_PrimeDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalcService_serviceDesc.Streams[0], "/calculator.CalcService/PrimeDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServicePrimeDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalcService_PrimeDecompositionClient interface {
	Recv() (*PrimeDecompResponse, error)
	grpc.ClientStream
}

type calcServicePrimeDecompositionClient struct {
	grpc.ClientStream
}

func (x *calcServicePrimeDecompositionClient) Recv() (*PrimeDecompResponse, error) {
	m := new(PrimeDecompResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalcService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalcService_serviceDesc.Streams[1], "/calculator.CalcService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServiceComputeAverageClient{stream}
	return x, nil
}

type CalcService_ComputeAverageClient interface {
	Send(*AverageRequest) error
	CloseAndRecv() (*AverageResponse, error)
	grpc.ClientStream
}

type calcServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calcServiceComputeAverageClient) Send(m *AverageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcServiceComputeAverageClient) CloseAndRecv() (*AverageResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AverageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServiceClient) SquareRoot(ctx context.Context, in *SquareRootRequest, opts ...grpc.CallOption) (*SquareRootResponse, error) {
	out := new(SquareRootResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalcService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServiceServer is the server API for CalcService service.
type CalcServiceServer interface {
	// Unary
	Sum(context.Context, *CalcRequest) (*CalcResponse, error)
	// Server streaming
	PrimeDecomposition(*PrimeDecompRequest, CalcService_PrimeDecompositionServer) error
	ComputeAverage(CalcService_ComputeAverageServer) error
	SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error)
}

// UnimplementedCalcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalcServiceServer struct {
}

func (*UnimplementedCalcServiceServer) Sum(context.Context, *CalcRequest) (*CalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalcServiceServer) PrimeDecomposition(*PrimeDecompRequest, CalcService_PrimeDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeDecomposition not implemented")
}
func (*UnimplementedCalcServiceServer) ComputeAverage(CalcService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}
func (*UnimplementedCalcServiceServer) SquareRoot(context.Context, *SquareRootRequest) (*SquareRootResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}

func RegisterCalcServiceServer(s *grpc.Server, srv CalcServiceServer) {
	s.RegisterService(&_CalcService_serviceDesc, srv)
}

func _CalcService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalcService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Sum(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_PrimeDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeDecompRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalcServiceServer).PrimeDecomposition(m, &calcServicePrimeDecompositionServer{stream})
}

type CalcService_PrimeDecompositionServer interface {
	Send(*PrimeDecompResponse) error
	grpc.ServerStream
}

type calcServicePrimeDecompositionServer struct {
	grpc.ServerStream
}

func (x *calcServicePrimeDecompositionServer) Send(m *PrimeDecompResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalcService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServiceServer).ComputeAverage(&calcServiceComputeAverageServer{stream})
}

type CalcService_ComputeAverageServer interface {
	SendAndClose(*AverageResponse) error
	Recv() (*AverageRequest, error)
	grpc.ServerStream
}

type calcServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calcServiceComputeAverageServer) SendAndClose(m *AverageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcServiceComputeAverageServer) Recv() (*AverageRequest, error) {
	m := new(AverageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalcService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SquareRootRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalcService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).SquareRoot(ctx, req.(*SquareRootRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalcService",
	HandlerType: (*CalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalcService_Sum_Handler,
		},
		{
			MethodName: "SquareRoot",
			Handler:    _CalcService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeDecomposition",
			Handler:       _CalcService_PrimeDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalcService_ComputeAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calcpb/calc.proto",
}
