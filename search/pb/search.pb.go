// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.1
// source: search.proto

package pb

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

type Post struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Post) Reset() {
	*x = Post{}
	mi := &file_search_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Post) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type PostBlogRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostBlogRequest) Reset() {
	*x = PostBlogRequest{}
	mi := &file_search_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostBlogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostBlogRequest) ProtoMessage() {}

func (x *PostBlogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostBlogRequest.ProtoReflect.Descriptor instead.
func (*PostBlogRequest) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{1}
}

func (x *PostBlogRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PostBlogRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type PostBlogResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Post          *Post                  `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostBlogResponse) Reset() {
	*x = PostBlogResponse{}
	mi := &file_search_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostBlogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostBlogResponse) ProtoMessage() {}

func (x *PostBlogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostBlogResponse.ProtoReflect.Descriptor instead.
func (*PostBlogResponse) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{2}
}

func (x *PostBlogResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type GetPostRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	mi := &file_search_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPostResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Post          *Post                  `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPostResponse) Reset() {
	*x = GetPostResponse{}
	mi := &file_search_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostResponse) ProtoMessage() {}

func (x *GetPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostResponse.ProtoReflect.Descriptor instead.
func (*GetPostResponse) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{4}
}

func (x *GetPostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type GetPostsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Skip          uint64                 `protobuf:"varint,1,opt,name=skip,proto3" json:"skip,omitempty"`
	Take          uint64                 `protobuf:"varint,2,opt,name=take,proto3" json:"take,omitempty"`
	Ids           []string               `protobuf:"bytes,3,rep,name=ids,proto3" json:"ids,omitempty"`
	Query         string                 `protobuf:"bytes,4,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPostsRequest) Reset() {
	*x = GetPostsRequest{}
	mi := &file_search_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsRequest) ProtoMessage() {}

func (x *GetPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsRequest.ProtoReflect.Descriptor instead.
func (*GetPostsRequest) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{5}
}

func (x *GetPostsRequest) GetSkip() uint64 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *GetPostsRequest) GetTake() uint64 {
	if x != nil {
		return x.Take
	}
	return 0
}

func (x *GetPostsRequest) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GetPostsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type GetPostsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Posts         []*Post                `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPostsResponse) Reset() {
	*x = GetPostsResponse{}
	mi := &file_search_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsResponse) ProtoMessage() {}

func (x *GetPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsResponse.ProtoReflect.Descriptor instead.
func (*GetPostsResponse) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{6}
}

func (x *GetPostsResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

var File_search_proto protoreflect.FileDescriptor

const file_search_proto_rawDesc = "" +
	"\n" +
	"\fsearch.proto\x12\x02pb\"L\n" +
	"\x04Post\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\"G\n" +
	"\x0fPostBlogRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\"0\n" +
	"\x10PostBlogResponse\x12\x1c\n" +
	"\x04post\x18\x01 \x01(\v2\b.pb.PostR\x04post\" \n" +
	"\x0eGetPostRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"/\n" +
	"\x0fGetPostResponse\x12\x1c\n" +
	"\x04post\x18\x01 \x01(\v2\b.pb.PostR\x04post\"a\n" +
	"\x0fGetPostsRequest\x12\x12\n" +
	"\x04skip\x18\x01 \x01(\x04R\x04skip\x12\x12\n" +
	"\x04take\x18\x02 \x01(\x04R\x04take\x12\x10\n" +
	"\x03ids\x18\x03 \x03(\tR\x03ids\x12\x14\n" +
	"\x05query\x18\x04 \x01(\tR\x05query\"2\n" +
	"\x10GetPostsResponse\x12\x1e\n" +
	"\x05posts\x18\x01 \x03(\v2\b.pb.PostR\x05posts2\xb1\x01\n" +
	"\rSearchService\x125\n" +
	"\bPostBlog\x12\x13.pb.PostBlogRequest\x1a\x14.pb.PostBlogResponse\x122\n" +
	"\aGetPost\x12\x12.pb.GetPostRequest\x1a\x13.pb.GetPostResponse\x125\n" +
	"\bGetPosts\x12\x13.pb.GetPostsRequest\x1a\x14.pb.GetPostsResponseB\x04Z\x02./b\x06proto3"

var (
	file_search_proto_rawDescOnce sync.Once
	file_search_proto_rawDescData []byte
)

func file_search_proto_rawDescGZIP() []byte {
	file_search_proto_rawDescOnce.Do(func() {
		file_search_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_search_proto_rawDesc), len(file_search_proto_rawDesc)))
	})
	return file_search_proto_rawDescData
}

var file_search_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_search_proto_goTypes = []any{
	(*Post)(nil),             // 0: pb.Post
	(*PostBlogRequest)(nil),  // 1: pb.PostBlogRequest
	(*PostBlogResponse)(nil), // 2: pb.PostBlogResponse
	(*GetPostRequest)(nil),   // 3: pb.GetPostRequest
	(*GetPostResponse)(nil),  // 4: pb.GetPostResponse
	(*GetPostsRequest)(nil),  // 5: pb.GetPostsRequest
	(*GetPostsResponse)(nil), // 6: pb.GetPostsResponse
}
var file_search_proto_depIdxs = []int32{
	0, // 0: pb.PostBlogResponse.post:type_name -> pb.Post
	0, // 1: pb.GetPostResponse.post:type_name -> pb.Post
	0, // 2: pb.GetPostsResponse.posts:type_name -> pb.Post
	1, // 3: pb.SearchService.PostBlog:input_type -> pb.PostBlogRequest
	3, // 4: pb.SearchService.GetPost:input_type -> pb.GetPostRequest
	5, // 5: pb.SearchService.GetPosts:input_type -> pb.GetPostsRequest
	2, // 6: pb.SearchService.PostBlog:output_type -> pb.PostBlogResponse
	4, // 7: pb.SearchService.GetPost:output_type -> pb.GetPostResponse
	6, // 8: pb.SearchService.GetPosts:output_type -> pb.GetPostsResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_search_proto_init() }
func file_search_proto_init() {
	if File_search_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_search_proto_rawDesc), len(file_search_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_search_proto_goTypes,
		DependencyIndexes: file_search_proto_depIdxs,
		MessageInfos:      file_search_proto_msgTypes,
	}.Build()
	File_search_proto = out.File
	file_search_proto_goTypes = nil
	file_search_proto_depIdxs = nil
}
