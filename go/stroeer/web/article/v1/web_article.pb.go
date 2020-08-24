// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.4
// source: stroeer/web/article/v1/web_article.proto

package article

import (
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/stroeer/tapir/go/stroeer/core/v1"
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

type RelatedWebArticleRole int32

const (
	RelatedWebArticleRole_RELATED_WEB_ARTICLE_ROLE_UNSPECIFIED RelatedWebArticleRole = 0
	RelatedWebArticleRole_RELATED_WEB_ARTICLE_ROLE_EDITORIAL   RelatedWebArticleRole = 1
)

// Enum value maps for RelatedWebArticleRole.
var (
	RelatedWebArticleRole_name = map[int32]string{
		0: "RELATED_WEB_ARTICLE_ROLE_UNSPECIFIED",
		1: "RELATED_WEB_ARTICLE_ROLE_EDITORIAL",
	}
	RelatedWebArticleRole_value = map[string]int32{
		"RELATED_WEB_ARTICLE_ROLE_UNSPECIFIED": 0,
		"RELATED_WEB_ARTICLE_ROLE_EDITORIAL":   1,
	}
)

func (x RelatedWebArticleRole) Enum() *RelatedWebArticleRole {
	p := new(RelatedWebArticleRole)
	*p = x
	return p
}

func (x RelatedWebArticleRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RelatedWebArticleRole) Descriptor() protoreflect.EnumDescriptor {
	return file_stroeer_web_article_v1_web_article_proto_enumTypes[0].Descriptor()
}

func (RelatedWebArticleRole) Type() protoreflect.EnumType {
	return &file_stroeer_web_article_v1_web_article_proto_enumTypes[0]
}

func (x RelatedWebArticleRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RelatedWebArticleRole.Descriptor instead.
func (RelatedWebArticleRole) EnumDescriptor() ([]byte, []int) {
	return file_stroeer_web_article_v1_web_article_proto_rawDescGZIP(), []int{0}
}

type WebArticle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	WebUrl       string            `protobuf:"bytes,2,opt,name=web_url,json=webUrl,proto3" json:"web_url,omitempty"`
	CanonicalUrl string            `protobuf:"bytes,3,opt,name=canonical_url,json=canonicalUrl,proto3" json:"canonical_url,omitempty"`
	Type         v1.ContentType    `protobuf:"varint,4,opt,name=type,proto3,enum=stroeer.core.v1.ContentType" json:"type,omitempty"`
	SectionPath  string            `protobuf:"bytes,5,opt,name=section_path,json=sectionPath,proto3" json:"section_path,omitempty"`
	Headline     string            `protobuf:"bytes,6,opt,name=headline,proto3" json:"headline,omitempty"`
	Fields       map[string]string `protobuf:"bytes,7,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Elements     []*v1.Element     `protobuf:"bytes,8,rep,name=elements,proto3" json:"elements,omitempty"`
	Authors      []*v1.Author      `protobuf:"bytes,9,rep,name=authors,proto3" json:"authors,omitempty"`
	BodyNodes    []*v1.BodyNode    `protobuf:"bytes,10,rep,name=body_nodes,json=bodyNodes,proto3" json:"body_nodes,omitempty"`
}

func (x *WebArticle) Reset() {
	*x = WebArticle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_web_article_v1_web_article_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebArticle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebArticle) ProtoMessage() {}

func (x *WebArticle) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_web_article_v1_web_article_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebArticle.ProtoReflect.Descriptor instead.
func (*WebArticle) Descriptor() ([]byte, []int) {
	return file_stroeer_web_article_v1_web_article_proto_rawDescGZIP(), []int{0}
}

func (x *WebArticle) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WebArticle) GetWebUrl() string {
	if x != nil {
		return x.WebUrl
	}
	return ""
}

func (x *WebArticle) GetCanonicalUrl() string {
	if x != nil {
		return x.CanonicalUrl
	}
	return ""
}

func (x *WebArticle) GetType() v1.ContentType {
	if x != nil {
		return x.Type
	}
	return v1.ContentType_CONTENT_TYPE_UNSPECIFIED
}

func (x *WebArticle) GetSectionPath() string {
	if x != nil {
		return x.SectionPath
	}
	return ""
}

func (x *WebArticle) GetHeadline() string {
	if x != nil {
		return x.Headline
	}
	return ""
}

func (x *WebArticle) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *WebArticle) GetElements() []*v1.Element {
	if x != nil {
		return x.Elements
	}
	return nil
}

func (x *WebArticle) GetAuthors() []*v1.Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

func (x *WebArticle) GetBodyNodes() []*v1.BodyNode {
	if x != nil {
		return x.BodyNodes
	}
	return nil
}

type RelatedWebArticle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Article *WebArticle           `protobuf:"bytes,1,opt,name=article,proto3" json:"article,omitempty"`
	Role    RelatedWebArticleRole `protobuf:"varint,2,opt,name=role,proto3,enum=stroeer.web.article.v1.RelatedWebArticleRole" json:"role,omitempty"`
}

func (x *RelatedWebArticle) Reset() {
	*x = RelatedWebArticle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_web_article_v1_web_article_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelatedWebArticle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelatedWebArticle) ProtoMessage() {}

func (x *RelatedWebArticle) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_web_article_v1_web_article_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelatedWebArticle.ProtoReflect.Descriptor instead.
func (*RelatedWebArticle) Descriptor() ([]byte, []int) {
	return file_stroeer_web_article_v1_web_article_proto_rawDescGZIP(), []int{1}
}

func (x *RelatedWebArticle) GetArticle() *WebArticle {
	if x != nil {
		return x.Article
	}
	return nil
}

func (x *RelatedWebArticle) GetRole() RelatedWebArticleRole {
	if x != nil {
		return x.Role
	}
	return RelatedWebArticleRole_RELATED_WEB_ARTICLE_ROLE_UNSPECIFIED
}

var File_stroeer_web_article_v1_web_article_proto protoreflect.FileDescriptor

var file_stroeer_web_article_v1_web_article_proto_rawDesc = []byte{
	0x0a, 0x28, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x65, 0x62, 0x5f, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x73, 0x74, 0x72, 0x6f,
	0x65, 0x65, 0x72, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x1a, 0x1d, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf1, 0x03, 0x0a, 0x0a, 0x57, 0x65, 0x62, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x77, 0x65, 0x62, 0x55, 0x72, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x6e,
	0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x12, 0x30,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x73,
	0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x46, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x41, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x34, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x6f,
	0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x31, 0x0a,
	0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73,
	0x12, 0x38, 0x0a, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x09, 0x62, 0x6f, 0x64, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x94, 0x01, 0x0a, 0x11, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x57, 0x65, 0x62, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x73,
	0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x65, 0x62, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x52, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65,
	0x72, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x57, 0x65, 0x62, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x2a, 0x69, 0x0a, 0x15,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x57, 0x65, 0x62, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x28, 0x0a, 0x24, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x45, 0x44,
	0x5f, 0x57, 0x45, 0x42, 0x5f, 0x41, 0x52, 0x54, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x52, 0x4f, 0x4c,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x26, 0x0a, 0x22, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x45, 0x44, 0x5f, 0x57, 0x45, 0x42, 0x5f, 0x41,
	0x52, 0x54, 0x49, 0x43, 0x4c, 0x45, 0x5f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x45, 0x44, 0x49, 0x54,
	0x4f, 0x52, 0x49, 0x41, 0x4c, 0x10, 0x01, 0x42, 0x59, 0x0a, 0x19, 0x64, 0x65, 0x2e, 0x73, 0x74,
	0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x70, 0x69, 0x72,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x77, 0x65, 0x62, 0x2f,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stroeer_web_article_v1_web_article_proto_rawDescOnce sync.Once
	file_stroeer_web_article_v1_web_article_proto_rawDescData = file_stroeer_web_article_v1_web_article_proto_rawDesc
)

func file_stroeer_web_article_v1_web_article_proto_rawDescGZIP() []byte {
	file_stroeer_web_article_v1_web_article_proto_rawDescOnce.Do(func() {
		file_stroeer_web_article_v1_web_article_proto_rawDescData = protoimpl.X.CompressGZIP(file_stroeer_web_article_v1_web_article_proto_rawDescData)
	})
	return file_stroeer_web_article_v1_web_article_proto_rawDescData
}

var file_stroeer_web_article_v1_web_article_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_stroeer_web_article_v1_web_article_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_stroeer_web_article_v1_web_article_proto_goTypes = []interface{}{
	(RelatedWebArticleRole)(0), // 0: stroeer.web.article.v1.RelatedWebArticleRole
	(*WebArticle)(nil),         // 1: stroeer.web.article.v1.WebArticle
	(*RelatedWebArticle)(nil),  // 2: stroeer.web.article.v1.RelatedWebArticle
	nil,                        // 3: stroeer.web.article.v1.WebArticle.FieldsEntry
	(v1.ContentType)(0),        // 4: stroeer.core.v1.ContentType
	(*v1.Element)(nil),         // 5: stroeer.core.v1.Element
	(*v1.Author)(nil),          // 6: stroeer.core.v1.Author
	(*v1.BodyNode)(nil),        // 7: stroeer.core.v1.BodyNode
}
var file_stroeer_web_article_v1_web_article_proto_depIdxs = []int32{
	4, // 0: stroeer.web.article.v1.WebArticle.type:type_name -> stroeer.core.v1.ContentType
	3, // 1: stroeer.web.article.v1.WebArticle.fields:type_name -> stroeer.web.article.v1.WebArticle.FieldsEntry
	5, // 2: stroeer.web.article.v1.WebArticle.elements:type_name -> stroeer.core.v1.Element
	6, // 3: stroeer.web.article.v1.WebArticle.authors:type_name -> stroeer.core.v1.Author
	7, // 4: stroeer.web.article.v1.WebArticle.body_nodes:type_name -> stroeer.core.v1.BodyNode
	1, // 5: stroeer.web.article.v1.RelatedWebArticle.article:type_name -> stroeer.web.article.v1.WebArticle
	0, // 6: stroeer.web.article.v1.RelatedWebArticle.role:type_name -> stroeer.web.article.v1.RelatedWebArticleRole
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_stroeer_web_article_v1_web_article_proto_init() }
func file_stroeer_web_article_v1_web_article_proto_init() {
	if File_stroeer_web_article_v1_web_article_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stroeer_web_article_v1_web_article_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebArticle); i {
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
		file_stroeer_web_article_v1_web_article_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelatedWebArticle); i {
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
			RawDescriptor: file_stroeer_web_article_v1_web_article_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stroeer_web_article_v1_web_article_proto_goTypes,
		DependencyIndexes: file_stroeer_web_article_v1_web_article_proto_depIdxs,
		EnumInfos:         file_stroeer_web_article_v1_web_article_proto_enumTypes,
		MessageInfos:      file_stroeer_web_article_v1_web_article_proto_msgTypes,
	}.Build()
	File_stroeer_web_article_v1_web_article_proto = out.File
	file_stroeer_web_article_v1_web_article_proto_rawDesc = nil
	file_stroeer_web_article_v1_web_article_proto_goTypes = nil
	file_stroeer_web_article_v1_web_article_proto_depIdxs = nil
}
