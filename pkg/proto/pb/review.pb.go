// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.21.12
// source: pkg/proto/review.proto

package pb

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

type Review struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductID     uint64                 `protobuf:"varint,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
	Mark          uint32                 `protobuf:"varint,3,opt,name=Mark,proto3" json:"Mark,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Review) Reset() {
	*x = Review{}
	mi := &file_pkg_proto_review_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_review_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_pkg_proto_review_proto_rawDescGZIP(), []int{0}
}

func (x *Review) GetProductID() uint64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *Review) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Review) GetMark() uint32 {
	if x != nil {
		return x.Mark
	}
	return 0
}

type AddReviewReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Review        *Review                `protobuf:"bytes,1,opt,name=Review,proto3" json:"Review,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddReviewReq) Reset() {
	*x = AddReviewReq{}
	mi := &file_pkg_proto_review_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddReviewReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReviewReq) ProtoMessage() {}

func (x *AddReviewReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_review_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReviewReq.ProtoReflect.Descriptor instead.
func (*AddReviewReq) Descriptor() ([]byte, []int) {
	return file_pkg_proto_review_proto_rawDescGZIP(), []int{1}
}

func (x *AddReviewReq) GetReview() *Review {
	if x != nil {
		return x.Review
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Error         string                 `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
	Status        int32                  `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_pkg_proto_review_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_review_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_pkg_proto_review_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Response) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type GetReviewsReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ProductID     uint64                 `protobuf:"varint,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetReviewsReq) Reset() {
	*x = GetReviewsReq{}
	mi := &file_pkg_proto_review_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReviewsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReviewsReq) ProtoMessage() {}

func (x *GetReviewsReq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_review_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReviewsReq.ProtoReflect.Descriptor instead.
func (*GetReviewsReq) Descriptor() ([]byte, []int) {
	return file_pkg_proto_review_proto_rawDescGZIP(), []int{3}
}

func (x *GetReviewsReq) GetProductID() uint64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

type GetReviewsResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Reviews       []*Review              `protobuf:"bytes,1,rep,name=Reviews,proto3" json:"Reviews,omitempty"`
	Response      *Response              `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetReviewsResp) Reset() {
	*x = GetReviewsResp{}
	mi := &file_pkg_proto_review_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReviewsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReviewsResp) ProtoMessage() {}

func (x *GetReviewsResp) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_review_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReviewsResp.ProtoReflect.Descriptor instead.
func (*GetReviewsResp) Descriptor() ([]byte, []int) {
	return file_pkg_proto_review_proto_rawDescGZIP(), []int{4}
}

func (x *GetReviewsResp) GetReviews() []*Review {
	if x != nil {
		return x.Reviews
	}
	return nil
}

func (x *GetReviewsResp) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_pkg_proto_review_proto protoreflect.FileDescriptor

var file_pkg_proto_review_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x06, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4d, 0x61,
	0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x4d, 0x61, 0x72, 0x6b, 0x22, 0x2f,
	0x0a, 0x0c, 0x41, 0x64, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x12, 0x1f,
	0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22,
	0x38, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x2d, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x22, 0x5a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x07, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x52, 0x07, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x25, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x32, 0x58, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x0d, 0x2e, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x09, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x10,
	0x5a, 0x0e, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_review_proto_rawDescOnce sync.Once
	file_pkg_proto_review_proto_rawDescData = file_pkg_proto_review_proto_rawDesc
)

func file_pkg_proto_review_proto_rawDescGZIP() []byte {
	file_pkg_proto_review_proto_rawDescOnce.Do(func() {
		file_pkg_proto_review_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_review_proto_rawDescData)
	})
	return file_pkg_proto_review_proto_rawDescData
}

var file_pkg_proto_review_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_proto_review_proto_goTypes = []any{
	(*Review)(nil),         // 0: review
	(*AddReviewReq)(nil),   // 1: AddReviewReq
	(*Response)(nil),       // 2: Response
	(*GetReviewsReq)(nil),  // 3: GetReviewsReq
	(*GetReviewsResp)(nil), // 4: GetReviewsResp
}
var file_pkg_proto_review_proto_depIdxs = []int32{
	0, // 0: AddReviewReq.Review:type_name -> review
	0, // 1: GetReviewsResp.Reviews:type_name -> review
	2, // 2: GetReviewsResp.response:type_name -> Response
	1, // 3: ReviewService.Add:input_type -> AddReviewReq
	3, // 4: ReviewService.Get:input_type -> GetReviewsReq
	2, // 5: ReviewService.Add:output_type -> Response
	4, // 6: ReviewService.Get:output_type -> GetReviewsResp
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_proto_review_proto_init() }
func file_pkg_proto_review_proto_init() {
	if File_pkg_proto_review_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_proto_review_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_review_proto_goTypes,
		DependencyIndexes: file_pkg_proto_review_proto_depIdxs,
		MessageInfos:      file_pkg_proto_review_proto_msgTypes,
	}.Build()
	File_pkg_proto_review_proto = out.File
	file_pkg_proto_review_proto_rawDesc = nil
	file_pkg_proto_review_proto_goTypes = nil
	file_pkg_proto_review_proto_depIdxs = nil
}
