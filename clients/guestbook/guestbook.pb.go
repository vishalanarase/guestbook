// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: guestbook.proto

package guestbook

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Token         string                 `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddMessageRequest) Reset() {
	*x = AddMessageRequest{}
	mi := &file_guestbook_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMessageRequest) ProtoMessage() {}

func (x *AddMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_guestbook_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMessageRequest.ProtoReflect.Descriptor instead.
func (*AddMessageRequest) Descriptor() ([]byte, []int) {
	return file_guestbook_proto_rawDescGZIP(), []int{0}
}

func (x *AddMessageRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AddMessageRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type AddMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddMessageResponse) Reset() {
	*x = AddMessageResponse{}
	mi := &file_guestbook_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddMessageResponse) ProtoMessage() {}

func (x *AddMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_guestbook_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddMessageResponse.ProtoReflect.Descriptor instead.
func (*AddMessageResponse) Descriptor() ([]byte, []int) {
	return file_guestbook_proto_rawDescGZIP(), []int{1}
}

func (x *AddMessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetMessagesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Token         string                 `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMessagesRequest) Reset() {
	*x = GetMessagesRequest{}
	mi := &file_guestbook_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesRequest) ProtoMessage() {}

func (x *GetMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_guestbook_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetMessagesRequest) Descriptor() ([]byte, []int) {
	return file_guestbook_proto_rawDescGZIP(), []int{2}
}

func (x *GetMessagesRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Messages      []string               `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMessageResponse) Reset() {
	*x = GetMessageResponse{}
	mi := &file_guestbook_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageResponse) ProtoMessage() {}

func (x *GetMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_guestbook_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageResponse.ProtoReflect.Descriptor instead.
func (*GetMessageResponse) Descriptor() ([]byte, []int) {
	return file_guestbook_proto_rawDescGZIP(), []int{3}
}

func (x *GetMessageResponse) GetMessages() []string {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_guestbook_proto protoreflect.FileDescriptor

var file_guestbook_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x67, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x67, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x5f, 0x0a, 0x11,
	0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x2e, 0x0a,
	0x12, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x2a, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x30, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0xad, 0x01, 0x0a, 0x10,
	0x47, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4b, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c,
	0x2e, 0x67, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67,
	0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x41, 0x64, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x75,
	0x65, 0x73, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x75, 0x65,
	0x73, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x41, 0x5a, 0x3f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x73, 0x68, 0x61, 0x6c,
	0x61, 0x6e, 0x61, 0x72, 0x61, 0x73, 0x65, 0x2f, 0x67, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x6f,
	0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x75, 0x65, 0x73, 0x74,
	0x62, 0x6f, 0x6f, 0x6b, 0x3b, 0x67, 0x75, 0x65, 0x73, 0x74, 0x62, 0x6f, 0x6f, 0x6b, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_guestbook_proto_rawDescOnce sync.Once
	file_guestbook_proto_rawDescData []byte
)

func file_guestbook_proto_rawDescGZIP() []byte {
	file_guestbook_proto_rawDescOnce.Do(func() {
		file_guestbook_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_guestbook_proto_rawDesc), len(file_guestbook_proto_rawDesc)))
	})
	return file_guestbook_proto_rawDescData
}

var file_guestbook_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_guestbook_proto_goTypes = []any{
	(*AddMessageRequest)(nil),  // 0: guestbook.AddMessageRequest
	(*AddMessageResponse)(nil), // 1: guestbook.AddMessageResponse
	(*GetMessagesRequest)(nil), // 2: guestbook.GetMessagesRequest
	(*GetMessageResponse)(nil), // 3: guestbook.GetMessageResponse
}
var file_guestbook_proto_depIdxs = []int32{
	0, // 0: guestbook.GuestbookService.AddMessage:input_type -> guestbook.AddMessageRequest
	2, // 1: guestbook.GuestbookService.GetMessage:input_type -> guestbook.GetMessagesRequest
	1, // 2: guestbook.GuestbookService.AddMessage:output_type -> guestbook.AddMessageResponse
	3, // 3: guestbook.GuestbookService.GetMessage:output_type -> guestbook.GetMessageResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_guestbook_proto_init() }
func file_guestbook_proto_init() {
	if File_guestbook_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_guestbook_proto_rawDesc), len(file_guestbook_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_guestbook_proto_goTypes,
		DependencyIndexes: file_guestbook_proto_depIdxs,
		MessageInfos:      file_guestbook_proto_msgTypes,
	}.Build()
	File_guestbook_proto = out.File
	file_guestbook_proto_goTypes = nil
	file_guestbook_proto_depIdxs = nil
}
