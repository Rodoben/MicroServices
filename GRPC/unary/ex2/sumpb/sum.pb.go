// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: sum.proto

package sumpb

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

type Numbers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num1 float32 `protobuf:"fixed32,1,opt,name=num1,proto3" json:"num1,omitempty"`
	Num2 float32 `protobuf:"fixed32,2,opt,name=num2,proto3" json:"num2,omitempty"`
}

func (x *Numbers) Reset() {
	*x = Numbers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sum_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Numbers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Numbers) ProtoMessage() {}

func (x *Numbers) ProtoReflect() protoreflect.Message {
	mi := &file_sum_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Numbers.ProtoReflect.Descriptor instead.
func (*Numbers) Descriptor() ([]byte, []int) {
	return file_sum_proto_rawDescGZIP(), []int{0}
}

func (x *Numbers) GetNum1() float32 {
	if x != nil {
		return x.Num1
	}
	return 0
}

func (x *Numbers) GetNum2() float32 {
	if x != nil {
		return x.Num2
	}
	return 0
}

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Numbers *Numbers `protobuf:"bytes,1,opt,name=numbers,proto3" json:"numbers,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sum_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sum_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_sum_proto_rawDescGZIP(), []int{1}
}

func (x *SumRequest) GetNumbers() *Numbers {
	if x != nil {
		return x.Numbers
	}
	return nil
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sum_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sum_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_sum_proto_rawDescGZIP(), []int{2}
}

func (x *SumResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_sum_proto protoreflect.FileDescriptor

var file_sum_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x73, 0x75, 0x6d,
	0x22, 0x31, 0x0a, 0x07, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x75, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6e, 0x75, 0x6d, 0x31, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6e,
	0x75, 0x6d, 0x32, 0x22, 0x34, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x26, 0x0a, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x75, 0x6d, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x52, 0x07, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x75, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0x38, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a,
	0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x0f, 0x2e, 0x73, 0x75, 0x6d, 0x2e, 0x53, 0x75, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x73, 0x75, 0x6d, 0x2e, 0x53, 0x75, 0x6d,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x73,
	0x75, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sum_proto_rawDescOnce sync.Once
	file_sum_proto_rawDescData = file_sum_proto_rawDesc
)

func file_sum_proto_rawDescGZIP() []byte {
	file_sum_proto_rawDescOnce.Do(func() {
		file_sum_proto_rawDescData = protoimpl.X.CompressGZIP(file_sum_proto_rawDescData)
	})
	return file_sum_proto_rawDescData
}

var file_sum_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sum_proto_goTypes = []interface{}{
	(*Numbers)(nil),     // 0: sum.Numbers
	(*SumRequest)(nil),  // 1: sum.SumRequest
	(*SumResponse)(nil), // 2: sum.SumResponse
}
var file_sum_proto_depIdxs = []int32{
	0, // 0: sum.SumRequest.numbers:type_name -> sum.Numbers
	1, // 1: sum.SumService.Sum:input_type -> sum.SumRequest
	2, // 2: sum.SumService.Sum:output_type -> sum.SumResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sum_proto_init() }
func file_sum_proto_init() {
	if File_sum_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sum_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Numbers); i {
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
		file_sum_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_sum_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
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
			RawDescriptor: file_sum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sum_proto_goTypes,
		DependencyIndexes: file_sum_proto_depIdxs,
		MessageInfos:      file_sum_proto_msgTypes,
	}.Build()
	File_sum_proto = out.File
	file_sum_proto_rawDesc = nil
	file_sum_proto_goTypes = nil
	file_sum_proto_depIdxs = nil
}