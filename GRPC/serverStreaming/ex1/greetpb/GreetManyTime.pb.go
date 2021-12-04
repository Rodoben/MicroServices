// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: GreetManyTime.proto

package greetpb

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

type Greeting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *Greeting) Reset() {
	*x = Greeting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GreetManyTime_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Greeting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Greeting) ProtoMessage() {}

func (x *Greeting) ProtoReflect() protoreflect.Message {
	mi := &file_GreetManyTime_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Greeting.ProtoReflect.Descriptor instead.
func (*Greeting) Descriptor() ([]byte, []int) {
	return file_GreetManyTime_proto_rawDescGZIP(), []int{0}
}

func (x *Greeting) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Greeting) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type GreetManyTimesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greeting *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
}

func (x *GreetManyTimesRequest) Reset() {
	*x = GreetManyTimesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GreetManyTime_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetManyTimesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetManyTimesRequest) ProtoMessage() {}

func (x *GreetManyTimesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_GreetManyTime_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetManyTimesRequest.ProtoReflect.Descriptor instead.
func (*GreetManyTimesRequest) Descriptor() ([]byte, []int) {
	return file_GreetManyTime_proto_rawDescGZIP(), []int{1}
}

func (x *GreetManyTimesRequest) GetGreeting() *Greeting {
	if x != nil {
		return x.Greeting
	}
	return nil
}

type GreetManyTimesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GreetManyTimesResponse) Reset() {
	*x = GreetManyTimesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GreetManyTime_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GreetManyTimesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetManyTimesResponse) ProtoMessage() {}

func (x *GreetManyTimesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_GreetManyTime_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetManyTimesResponse.ProtoReflect.Descriptor instead.
func (*GreetManyTimesResponse) Descriptor() ([]byte, []int) {
	return file_GreetManyTime_proto_rawDescGZIP(), []int{2}
}

func (x *GreetManyTimesResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_GreetManyTime_proto protoreflect.FileDescriptor

var file_GreetManyTime_proto_rawDesc = []byte{
	0x0a, 0x13, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x22, 0x46, 0x0a, 0x08,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x22, 0x44, 0x0a, 0x15, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e,
	0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a,
	0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x52, 0x08, 0x67, 0x72, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x30, 0x0a, 0x16, 0x47, 0x72,
	0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x61, 0x0a, 0x0c,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0e,
	0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1c,
	0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x2e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GreetManyTime_proto_rawDescOnce sync.Once
	file_GreetManyTime_proto_rawDescData = file_GreetManyTime_proto_rawDesc
)

func file_GreetManyTime_proto_rawDescGZIP() []byte {
	file_GreetManyTime_proto_rawDescOnce.Do(func() {
		file_GreetManyTime_proto_rawDescData = protoimpl.X.CompressGZIP(file_GreetManyTime_proto_rawDescData)
	})
	return file_GreetManyTime_proto_rawDescData
}

var file_GreetManyTime_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_GreetManyTime_proto_goTypes = []interface{}{
	(*Greeting)(nil),               // 0: greet.Greeting
	(*GreetManyTimesRequest)(nil),  // 1: greet.GreetManyTimesRequest
	(*GreetManyTimesResponse)(nil), // 2: greet.GreetManyTimesResponse
}
var file_GreetManyTime_proto_depIdxs = []int32{
	0, // 0: greet.GreetManyTimesRequest.greeting:type_name -> greet.Greeting
	1, // 1: greet.GreetService.GreetManyTimes:input_type -> greet.GreetManyTimesRequest
	2, // 2: greet.GreetService.GreetManyTimes:output_type -> greet.GreetManyTimesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GreetManyTime_proto_init() }
func file_GreetManyTime_proto_init() {
	if File_GreetManyTime_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GreetManyTime_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Greeting); i {
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
		file_GreetManyTime_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetManyTimesRequest); i {
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
		file_GreetManyTime_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GreetManyTimesResponse); i {
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
			RawDescriptor: file_GreetManyTime_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_GreetManyTime_proto_goTypes,
		DependencyIndexes: file_GreetManyTime_proto_depIdxs,
		MessageInfos:      file_GreetManyTime_proto_msgTypes,
	}.Build()
	File_GreetManyTime_proto = out.File
	file_GreetManyTime_proto_rawDesc = nil
	file_GreetManyTime_proto_goTypes = nil
	file_GreetManyTime_proto_depIdxs = nil
}
