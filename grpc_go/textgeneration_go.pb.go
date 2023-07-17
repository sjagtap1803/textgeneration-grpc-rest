//Toy gRPC project

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: textgeneration_go.proto

package grpc_gateway

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

// request message containing the input prompt
type TextGenPrompt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prompt  string `protobuf:"bytes,1,opt,name=prompt,proto3" json:"prompt,omitempty"`
	Cleanup bool   `protobuf:"varint,2,opt,name=cleanup,proto3" json:"cleanup,omitempty"`
}

func (x *TextGenPrompt) Reset() {
	*x = TextGenPrompt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_textgeneration_go_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextGenPrompt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextGenPrompt) ProtoMessage() {}

func (x *TextGenPrompt) ProtoReflect() protoreflect.Message {
	mi := &file_textgeneration_go_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextGenPrompt.ProtoReflect.Descriptor instead.
func (*TextGenPrompt) Descriptor() ([]byte, []int) {
	return file_textgeneration_go_proto_rawDescGZIP(), []int{0}
}

func (x *TextGenPrompt) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *TextGenPrompt) GetCleanup() bool {
	if x != nil {
		return x.Cleanup
	}
	return false
}

// response message containing output generated by the model
type TextGenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TextGenResponse) Reset() {
	*x = TextGenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_textgeneration_go_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextGenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextGenResponse) ProtoMessage() {}

func (x *TextGenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_textgeneration_go_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextGenResponse.ProtoReflect.Descriptor instead.
func (*TextGenResponse) Descriptor() ([]byte, []int) {
	return file_textgeneration_go_proto_rawDescGZIP(), []int{1}
}

func (x *TextGenResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// request message for uploading a file
type FileUpload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to UploadOneof:
	//
	//	*FileUpload_Filename
	//	*FileUpload_Filedata
	UploadOneof isFileUpload_UploadOneof `protobuf_oneof:"upload_oneof"`
}

func (x *FileUpload) Reset() {
	*x = FileUpload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_textgeneration_go_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUpload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUpload) ProtoMessage() {}

func (x *FileUpload) ProtoReflect() protoreflect.Message {
	mi := &file_textgeneration_go_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUpload.ProtoReflect.Descriptor instead.
func (*FileUpload) Descriptor() ([]byte, []int) {
	return file_textgeneration_go_proto_rawDescGZIP(), []int{2}
}

func (m *FileUpload) GetUploadOneof() isFileUpload_UploadOneof {
	if m != nil {
		return m.UploadOneof
	}
	return nil
}

func (x *FileUpload) GetFilename() string {
	if x, ok := x.GetUploadOneof().(*FileUpload_Filename); ok {
		return x.Filename
	}
	return ""
}

func (x *FileUpload) GetFiledata() []byte {
	if x, ok := x.GetUploadOneof().(*FileUpload_Filedata); ok {
		return x.Filedata
	}
	return nil
}

type isFileUpload_UploadOneof interface {
	isFileUpload_UploadOneof()
}

type FileUpload_Filename struct {
	Filename string `protobuf:"bytes,1,opt,name=filename,proto3,oneof"`
}

type FileUpload_Filedata struct {
	Filedata []byte `protobuf:"bytes,2,opt,name=filedata,proto3,oneof"`
}

func (*FileUpload_Filename) isFileUpload_UploadOneof() {}

func (*FileUpload_Filedata) isFileUpload_UploadOneof() {}

var File_textgeneration_go_proto protoreflect.FileDescriptor

var file_textgeneration_go_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x65, 0x78, 0x74, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x65, 0x78, 0x74, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x41, 0x0a, 0x0d, 0x54, 0x65, 0x78,
	0x74, 0x47, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x6c, 0x65, 0x61, 0x6e, 0x75, 0x70, 0x22, 0x25, 0x0a, 0x0f,
	0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x22, 0x58, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x1c, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x48, 0x00, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0e, 0x0a,
	0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x32, 0xaf, 0x04,
	0x0a, 0x0e, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x52, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54, 0x65, 0x78, 0x74, 0x47,
	0x65, 0x6e, 0x12, 0x1d, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x5f, 0x0a, 0x19, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x1d, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x1a, 0x1f, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x61, 0x0a, 0x1b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x50, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x5c, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x1d, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x1a,
	0x1f, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x51, 0x0a, 0x10, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x2e, 0x74, 0x65, 0x78,
	0x74, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x1f, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0f, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1a, 0x2e, 0x74,
	0x65, 0x78, 0x74, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x1f, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x47, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_textgeneration_go_proto_rawDescOnce sync.Once
	file_textgeneration_go_proto_rawDescData = file_textgeneration_go_proto_rawDesc
)

func file_textgeneration_go_proto_rawDescGZIP() []byte {
	file_textgeneration_go_proto_rawDescOnce.Do(func() {
		file_textgeneration_go_proto_rawDescData = protoimpl.X.CompressGZIP(file_textgeneration_go_proto_rawDescData)
	})
	return file_textgeneration_go_proto_rawDescData
}

var file_textgeneration_go_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_textgeneration_go_proto_goTypes = []interface{}{
	(*TextGenPrompt)(nil),   // 0: textgeneration.TextGenPrompt
	(*TextGenResponse)(nil), // 1: textgeneration.TextGenResponse
	(*FileUpload)(nil),      // 2: textgeneration.FileUpload
}
var file_textgeneration_go_proto_depIdxs = []int32{
	0, // 0: textgeneration.TextGeneration.ProcessTextGen:input_type -> textgeneration.TextGenPrompt
	0, // 1: textgeneration.TextGeneration.ProcessTextGenMultiStream:input_type -> textgeneration.TextGenPrompt
	0, // 2: textgeneration.TextGeneration.ProcessTextGenStreamRequest:input_type -> textgeneration.TextGenPrompt
	0, // 3: textgeneration.TextGeneration.ProcessTextGenStream:input_type -> textgeneration.TextGenPrompt
	2, // 4: textgeneration.TextGeneration.SingleFileUpload:input_type -> textgeneration.FileUpload
	2, // 5: textgeneration.TextGeneration.MultiFileUpload:input_type -> textgeneration.FileUpload
	1, // 6: textgeneration.TextGeneration.ProcessTextGen:output_type -> textgeneration.TextGenResponse
	1, // 7: textgeneration.TextGeneration.ProcessTextGenMultiStream:output_type -> textgeneration.TextGenResponse
	1, // 8: textgeneration.TextGeneration.ProcessTextGenStreamRequest:output_type -> textgeneration.TextGenResponse
	1, // 9: textgeneration.TextGeneration.ProcessTextGenStream:output_type -> textgeneration.TextGenResponse
	1, // 10: textgeneration.TextGeneration.SingleFileUpload:output_type -> textgeneration.TextGenResponse
	1, // 11: textgeneration.TextGeneration.MultiFileUpload:output_type -> textgeneration.TextGenResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_textgeneration_go_proto_init() }
func file_textgeneration_go_proto_init() {
	if File_textgeneration_go_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_textgeneration_go_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextGenPrompt); i {
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
		file_textgeneration_go_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextGenResponse); i {
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
		file_textgeneration_go_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUpload); i {
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
	file_textgeneration_go_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*FileUpload_Filename)(nil),
		(*FileUpload_Filedata)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_textgeneration_go_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_textgeneration_go_proto_goTypes,
		DependencyIndexes: file_textgeneration_go_proto_depIdxs,
		MessageInfos:      file_textgeneration_go_proto_msgTypes,
	}.Build()
	File_textgeneration_go_proto = out.File
	file_textgeneration_go_proto_rawDesc = nil
	file_textgeneration_go_proto_goTypes = nil
	file_textgeneration_go_proto_depIdxs = nil
}
