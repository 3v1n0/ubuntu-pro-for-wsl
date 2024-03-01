// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: agentapi.proto

package agentapi

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentapi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_agentapi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_agentapi_proto_rawDescGZIP(), []int{0}
}

type ProAttachInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ProAttachInfo) Reset() {
	*x = ProAttachInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentapi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProAttachInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProAttachInfo) ProtoMessage() {}

func (x *ProAttachInfo) ProtoReflect() protoreflect.Message {
	mi := &file_agentapi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProAttachInfo.ProtoReflect.Descriptor instead.
func (*ProAttachInfo) Descriptor() ([]byte, []int) {
	return file_agentapi_proto_rawDescGZIP(), []int{1}
}

func (x *ProAttachInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LandscapeConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config string `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *LandscapeConfig) Reset() {
	*x = LandscapeConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentapi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LandscapeConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LandscapeConfig) ProtoMessage() {}

func (x *LandscapeConfig) ProtoReflect() protoreflect.Message {
	mi := &file_agentapi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LandscapeConfig.ProtoReflect.Descriptor instead.
func (*LandscapeConfig) Descriptor() ([]byte, []int) {
	return file_agentapi_proto_rawDescGZIP(), []int{2}
}

func (x *LandscapeConfig) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

type SubscriptionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"` // The ID of the Ubuntu Pro for WSL product on the Microsoft Store.
	// Types that are assignable to SubscriptionType:
	//
	//	*SubscriptionInfo_None
	//	*SubscriptionInfo_User
	//	*SubscriptionInfo_Organization
	//	*SubscriptionInfo_MicrosoftStore
	SubscriptionType isSubscriptionInfo_SubscriptionType `protobuf_oneof:"subscriptionType"`
}

func (x *SubscriptionInfo) Reset() {
	*x = SubscriptionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentapi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionInfo) ProtoMessage() {}

func (x *SubscriptionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_agentapi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionInfo.ProtoReflect.Descriptor instead.
func (*SubscriptionInfo) Descriptor() ([]byte, []int) {
	return file_agentapi_proto_rawDescGZIP(), []int{3}
}

func (x *SubscriptionInfo) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (m *SubscriptionInfo) GetSubscriptionType() isSubscriptionInfo_SubscriptionType {
	if m != nil {
		return m.SubscriptionType
	}
	return nil
}

func (x *SubscriptionInfo) GetNone() *Empty {
	if x, ok := x.GetSubscriptionType().(*SubscriptionInfo_None); ok {
		return x.None
	}
	return nil
}

func (x *SubscriptionInfo) GetUser() *Empty {
	if x, ok := x.GetSubscriptionType().(*SubscriptionInfo_User); ok {
		return x.User
	}
	return nil
}

func (x *SubscriptionInfo) GetOrganization() *Empty {
	if x, ok := x.GetSubscriptionType().(*SubscriptionInfo_Organization); ok {
		return x.Organization
	}
	return nil
}

func (x *SubscriptionInfo) GetMicrosoftStore() *Empty {
	if x, ok := x.GetSubscriptionType().(*SubscriptionInfo_MicrosoftStore); ok {
		return x.MicrosoftStore
	}
	return nil
}

type isSubscriptionInfo_SubscriptionType interface {
	isSubscriptionInfo_SubscriptionType()
}

type SubscriptionInfo_None struct {
	None *Empty `protobuf:"bytes,2,opt,name=none,proto3,oneof"` // There is no active subscription.
}

type SubscriptionInfo_User struct {
	User *Empty `protobuf:"bytes,3,opt,name=user,proto3,oneof"` // The subscription is managed by the user with a pro token from the GUI or the registry.
}

type SubscriptionInfo_Organization struct {
	Organization *Empty `protobuf:"bytes,4,opt,name=organization,proto3,oneof"` // The subscription is managed by the sysadmin with a pro token from the registry.
}

type SubscriptionInfo_MicrosoftStore struct {
	MicrosoftStore *Empty `protobuf:"bytes,5,opt,name=microsoftStore,proto3,oneof"` // The subscription is managed via the Microsoft store.
}

func (*SubscriptionInfo_None) isSubscriptionInfo_SubscriptionType() {}

func (*SubscriptionInfo_User) isSubscriptionInfo_SubscriptionType() {}

func (*SubscriptionInfo_Organization) isSubscriptionInfo_SubscriptionType() {}

func (*SubscriptionInfo_MicrosoftStore) isSubscriptionInfo_SubscriptionType() {}

type DistroInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WslName     string `protobuf:"bytes,1,opt,name=wsl_name,json=wslName,proto3" json:"wsl_name,omitempty"`
	Id          string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	VersionId   string `protobuf:"bytes,3,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	PrettyName  string `protobuf:"bytes,4,opt,name=pretty_name,json=prettyName,proto3" json:"pretty_name,omitempty"`
	ProAttached bool   `protobuf:"varint,5,opt,name=pro_attached,json=proAttached,proto3" json:"pro_attached,omitempty"`
	Hostname    string `protobuf:"bytes,6,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *DistroInfo) Reset() {
	*x = DistroInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentapi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DistroInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DistroInfo) ProtoMessage() {}

func (x *DistroInfo) ProtoReflect() protoreflect.Message {
	mi := &file_agentapi_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DistroInfo.ProtoReflect.Descriptor instead.
func (*DistroInfo) Descriptor() ([]byte, []int) {
	return file_agentapi_proto_rawDescGZIP(), []int{4}
}

func (x *DistroInfo) GetWslName() string {
	if x != nil {
		return x.WslName
	}
	return ""
}

func (x *DistroInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DistroInfo) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

func (x *DistroInfo) GetPrettyName() string {
	if x != nil {
		return x.PrettyName
	}
	return ""
}

func (x *DistroInfo) GetProAttached() bool {
	if x != nil {
		return x.ProAttached
	}
	return false
}

func (x *DistroInfo) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port uint32 `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agentapi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_agentapi_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_agentapi_proto_rawDescGZIP(), []int{5}
}

func (x *Port) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_agentapi_proto protoreflect.FileDescriptor

var file_agentapi_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x25, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x29, 0x0a, 0x0f, 0x4c, 0x61,
	0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x84, 0x02, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x04, 0x6e, 0x6f, 0x6e, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x61, 0x70,
	0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x12,
	0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52,
	0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a,
	0x0e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x6f, 0x66, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x69,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x48, 0x00, 0x52, 0x0e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73,
	0x6f, 0x66, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x42, 0x12, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0xb6, 0x01, 0x0a,
	0x0a, 0x44, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x77,
	0x73, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77,
	0x73, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x74, 0x74, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x74,
	0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x5f, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1a, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x32, 0xc5, 0x02, 0x0a, 0x02, 0x55, 0x49, 0x12, 0x46, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x50, 0x72, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x72, 0x6f, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x1a, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00,
	0x12, 0x44, 0x0a, 0x14, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x4c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61,
	0x70, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x61, 0x70, 0x69, 0x2e, 0x4c, 0x61, 0x6e, 0x64, 0x73, 0x63, 0x61, 0x70, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x0f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0f,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x12, 0x44, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0f, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0e, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x12, 0x0f, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x32, 0x46, 0x0a, 0x0b, 0x57, 0x53, 0x4c,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x14, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x61, 0x70, 0x69,
	0x2e, 0x44, 0x69, 0x73, 0x74, 0x72, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x0e, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x2f, 0x75, 0x62, 0x75, 0x6e, 0x74, 0x75,
	0x2d, 0x70, 0x72, 0x6f, 0x2d, 0x66, 0x6f, 0x72, 0x2d, 0x77, 0x73, 0x6c, 0x2f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agentapi_proto_rawDescOnce sync.Once
	file_agentapi_proto_rawDescData = file_agentapi_proto_rawDesc
)

func file_agentapi_proto_rawDescGZIP() []byte {
	file_agentapi_proto_rawDescOnce.Do(func() {
		file_agentapi_proto_rawDescData = protoimpl.X.CompressGZIP(file_agentapi_proto_rawDescData)
	})
	return file_agentapi_proto_rawDescData
}

var file_agentapi_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_agentapi_proto_goTypes = []interface{}{
	(*Empty)(nil),            // 0: agentapi.Empty
	(*ProAttachInfo)(nil),    // 1: agentapi.ProAttachInfo
	(*LandscapeConfig)(nil),  // 2: agentapi.LandscapeConfig
	(*SubscriptionInfo)(nil), // 3: agentapi.SubscriptionInfo
	(*DistroInfo)(nil),       // 4: agentapi.DistroInfo
	(*Port)(nil),             // 5: agentapi.Port
}
var file_agentapi_proto_depIdxs = []int32{
	0,  // 0: agentapi.SubscriptionInfo.none:type_name -> agentapi.Empty
	0,  // 1: agentapi.SubscriptionInfo.user:type_name -> agentapi.Empty
	0,  // 2: agentapi.SubscriptionInfo.organization:type_name -> agentapi.Empty
	0,  // 3: agentapi.SubscriptionInfo.microsoftStore:type_name -> agentapi.Empty
	1,  // 4: agentapi.UI.ApplyProToken:input_type -> agentapi.ProAttachInfo
	2,  // 5: agentapi.UI.ApplyLandscapeConfig:input_type -> agentapi.LandscapeConfig
	0,  // 6: agentapi.UI.Ping:input_type -> agentapi.Empty
	0,  // 7: agentapi.UI.GetSubscriptionInfo:input_type -> agentapi.Empty
	0,  // 8: agentapi.UI.NotifyPurchase:input_type -> agentapi.Empty
	4,  // 9: agentapi.WSLInstance.Connected:input_type -> agentapi.DistroInfo
	3,  // 10: agentapi.UI.ApplyProToken:output_type -> agentapi.SubscriptionInfo
	0,  // 11: agentapi.UI.ApplyLandscapeConfig:output_type -> agentapi.Empty
	0,  // 12: agentapi.UI.Ping:output_type -> agentapi.Empty
	3,  // 13: agentapi.UI.GetSubscriptionInfo:output_type -> agentapi.SubscriptionInfo
	3,  // 14: agentapi.UI.NotifyPurchase:output_type -> agentapi.SubscriptionInfo
	5,  // 15: agentapi.WSLInstance.Connected:output_type -> agentapi.Port
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_agentapi_proto_init() }
func file_agentapi_proto_init() {
	if File_agentapi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agentapi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_agentapi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProAttachInfo); i {
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
		file_agentapi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LandscapeConfig); i {
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
		file_agentapi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionInfo); i {
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
		file_agentapi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DistroInfo); i {
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
		file_agentapi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
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
	file_agentapi_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*SubscriptionInfo_None)(nil),
		(*SubscriptionInfo_User)(nil),
		(*SubscriptionInfo_Organization)(nil),
		(*SubscriptionInfo_MicrosoftStore)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_agentapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_agentapi_proto_goTypes,
		DependencyIndexes: file_agentapi_proto_depIdxs,
		MessageInfos:      file_agentapi_proto_msgTypes,
	}.Build()
	File_agentapi_proto = out.File
	file_agentapi_proto_rawDesc = nil
	file_agentapi_proto_goTypes = nil
	file_agentapi_proto_depIdxs = nil
}
