// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: calculatePB.proto

package CalcAvg

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

type CalcAvgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Num float32 `protobuf:"fixed32,1,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *CalcAvgRequest) Reset() {
	*x = CalcAvgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatePB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcAvgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcAvgRequest) ProtoMessage() {}

func (x *CalcAvgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculatePB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcAvgRequest.ProtoReflect.Descriptor instead.
func (*CalcAvgRequest) Descriptor() ([]byte, []int) {
	return file_calculatePB_proto_rawDescGZIP(), []int{0}
}

func (x *CalcAvgRequest) GetNum() float32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type CalcAvgResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result float32 `protobuf:"fixed32,1,opt,name=Result,proto3" json:"Result,omitempty"`
}

func (x *CalcAvgResponse) Reset() {
	*x = CalcAvgResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculatePB_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcAvgResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcAvgResponse) ProtoMessage() {}

func (x *CalcAvgResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculatePB_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcAvgResponse.ProtoReflect.Descriptor instead.
func (*CalcAvgResponse) Descriptor() ([]byte, []int) {
	return file_calculatePB_proto_rawDescGZIP(), []int{1}
}

func (x *CalcAvgResponse) GetResult() float32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculatePB_proto protoreflect.FileDescriptor

var file_calculatePB_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x42, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x61, 0x6c, 0x63, 0x41, 0x76, 0x67, 0x22, 0x22, 0x0a, 0x0e,
	0x63, 0x61, 0x6c, 0x63, 0x41, 0x76, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6e, 0x75, 0x6d,
	0x22, 0x29, 0x0a, 0x0f, 0x63, 0x61, 0x6c, 0x63, 0x41, 0x76, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x54, 0x0a, 0x10, 0x43,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x40, 0x0a, 0x07, 0x43, 0x61, 0x6c, 0x63, 0x41, 0x76, 0x67, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x41, 0x76, 0x67, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x41, 0x76, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x41, 0x76, 0x67, 0x2e, 0x63, 0x61,
	0x6c, 0x63, 0x41, 0x76, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28,
	0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x43, 0x61, 0x6c, 0x63, 0x41, 0x76, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calculatePB_proto_rawDescOnce sync.Once
	file_calculatePB_proto_rawDescData = file_calculatePB_proto_rawDesc
)

func file_calculatePB_proto_rawDescGZIP() []byte {
	file_calculatePB_proto_rawDescOnce.Do(func() {
		file_calculatePB_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculatePB_proto_rawDescData)
	})
	return file_calculatePB_proto_rawDescData
}

var file_calculatePB_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_calculatePB_proto_goTypes = []interface{}{
	(*CalcAvgRequest)(nil),  // 0: calcAvg.calcAvgRequest
	(*CalcAvgResponse)(nil), // 1: calcAvg.calcAvgResponse
}
var file_calculatePB_proto_depIdxs = []int32{
	0, // 0: calcAvg.CalculateService.CalcAvg:input_type -> calcAvg.calcAvgRequest
	1, // 1: calcAvg.CalculateService.CalcAvg:output_type -> calcAvg.calcAvgResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calculatePB_proto_init() }
func file_calculatePB_proto_init() {
	if File_calculatePB_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculatePB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcAvgRequest); i {
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
		file_calculatePB_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcAvgResponse); i {
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
			RawDescriptor: file_calculatePB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculatePB_proto_goTypes,
		DependencyIndexes: file_calculatePB_proto_depIdxs,
		MessageInfos:      file_calculatePB_proto_msgTypes,
	}.Build()
	File_calculatePB_proto = out.File
	file_calculatePB_proto_rawDesc = nil
	file_calculatePB_proto_goTypes = nil
	file_calculatePB_proto_depIdxs = nil
}