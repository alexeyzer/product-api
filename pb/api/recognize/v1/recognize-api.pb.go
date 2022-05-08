// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: api/recognize/v1/recognize-api.proto

package recognize_api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RecognizePhotoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *RecognizePhotoRequest) Reset() {
	*x = RecognizePhotoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_recognize_v1_recognize_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecognizePhotoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecognizePhotoRequest) ProtoMessage() {}

func (x *RecognizePhotoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_recognize_v1_recognize_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecognizePhotoRequest.ProtoReflect.Descriptor instead.
func (*RecognizePhotoRequest) Descriptor() ([]byte, []int) {
	return file_api_recognize_v1_recognize_api_proto_rawDescGZIP(), []int{0}
}

func (x *RecognizePhotoRequest) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

type RecognizePhotoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Category string `protobuf:"bytes,1,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *RecognizePhotoResponse) Reset() {
	*x = RecognizePhotoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_recognize_v1_recognize_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecognizePhotoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecognizePhotoResponse) ProtoMessage() {}

func (x *RecognizePhotoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_recognize_v1_recognize_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecognizePhotoResponse.ProtoReflect.Descriptor instead.
func (*RecognizePhotoResponse) Descriptor() ([]byte, []int) {
	return file_api_recognize_v1_recognize_api_proto_rawDescGZIP(), []int{1}
}

func (x *RecognizePhotoResponse) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

var File_api_recognize_v1_recognize_api_proto protoreflect.FileDescriptor

var file_api_recognize_v1_recognize_api_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x2d, 0x61, 0x70, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x2d, 0x0a, 0x15, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x50, 0x68, 0x6f,
	0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22,
	0x34, 0x0a, 0x16, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x50, 0x68, 0x6f, 0x74,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x32, 0x8e, 0x01, 0x0a, 0x13, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e,
	0x69, 0x7a, 0x65, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a,
	0x0e, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x12,
	0x24, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x50, 0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a,
	0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x50,
	0x68, 0x6f, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x67, 0x6e,
	0x69, 0x7a, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x65, 0x78, 0x65, 0x79, 0x7a, 0x65, 0x72, 0x2f, 0x72,
	0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x7a, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_recognize_v1_recognize_api_proto_rawDescOnce sync.Once
	file_api_recognize_v1_recognize_api_proto_rawDescData = file_api_recognize_v1_recognize_api_proto_rawDesc
)

func file_api_recognize_v1_recognize_api_proto_rawDescGZIP() []byte {
	file_api_recognize_v1_recognize_api_proto_rawDescOnce.Do(func() {
		file_api_recognize_v1_recognize_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_recognize_v1_recognize_api_proto_rawDescData)
	})
	return file_api_recognize_v1_recognize_api_proto_rawDescData
}

var file_api_recognize_v1_recognize_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_recognize_v1_recognize_api_proto_goTypes = []interface{}{
	(*RecognizePhotoRequest)(nil),  // 0: recognize.api.recognizePhotoRequest
	(*RecognizePhotoResponse)(nil), // 1: recognize.api.recognizePhotoResponse
}
var file_api_recognize_v1_recognize_api_proto_depIdxs = []int32{
	0, // 0: recognize.api.RecognizeApiService.recognizePhoto:input_type -> recognize.api.recognizePhotoRequest
	1, // 1: recognize.api.RecognizeApiService.recognizePhoto:output_type -> recognize.api.recognizePhotoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_recognize_v1_recognize_api_proto_init() }
func file_api_recognize_v1_recognize_api_proto_init() {
	if File_api_recognize_v1_recognize_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_recognize_v1_recognize_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecognizePhotoRequest); i {
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
		file_api_recognize_v1_recognize_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecognizePhotoResponse); i {
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
			RawDescriptor: file_api_recognize_v1_recognize_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_recognize_v1_recognize_api_proto_goTypes,
		DependencyIndexes: file_api_recognize_v1_recognize_api_proto_depIdxs,
		MessageInfos:      file_api_recognize_v1_recognize_api_proto_msgTypes,
	}.Build()
	File_api_recognize_v1_recognize_api_proto = out.File
	file_api_recognize_v1_recognize_api_proto_rawDesc = nil
	file_api_recognize_v1_recognize_api_proto_goTypes = nil
	file_api_recognize_v1_recognize_api_proto_depIdxs = nil
}
