//
//MIT License
//Copyright(c) 2020 Futurewei Cloud
//
//Permission is hereby granted,
//free of charge, to any person obtaining a copy of this software and associated documentation files(the "Software"), to deal in the Software without restriction,
//including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and / or sell copies of the Software, and to permit persons
//to whom the Software is furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
//WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: securitygroup.proto

package schema

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

type SecurityGroupConfiguration_Direction int32

const (
	SecurityGroupConfiguration_EGRESS  SecurityGroupConfiguration_Direction = 0
	SecurityGroupConfiguration_INGRESS SecurityGroupConfiguration_Direction = 1
)

// Enum value maps for SecurityGroupConfiguration_Direction.
var (
	SecurityGroupConfiguration_Direction_name = map[int32]string{
		0: "EGRESS",
		1: "INGRESS",
	}
	SecurityGroupConfiguration_Direction_value = map[string]int32{
		"EGRESS":  0,
		"INGRESS": 1,
	}
)

func (x SecurityGroupConfiguration_Direction) Enum() *SecurityGroupConfiguration_Direction {
	p := new(SecurityGroupConfiguration_Direction)
	*p = x
	return p
}

func (x SecurityGroupConfiguration_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecurityGroupConfiguration_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_securitygroup_proto_enumTypes[0].Descriptor()
}

func (SecurityGroupConfiguration_Direction) Type() protoreflect.EnumType {
	return &file_securitygroup_proto_enumTypes[0]
}

func (x SecurityGroupConfiguration_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SecurityGroupConfiguration_Direction.Descriptor instead.
func (SecurityGroupConfiguration_Direction) EnumDescriptor() ([]byte, []int) {
	return file_securitygroup_proto_rawDescGZIP(), []int{0, 0}
}

type SecurityGroupConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RevisionNumber     uint32                                          `protobuf:"varint,1,opt,name=revision_number,json=revisionNumber,proto3" json:"revision_number,omitempty"`
	RequestId          string                                          `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Id                 string                                          `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	UpdateType         UpdateType                                      `protobuf:"varint,4,opt,name=update_type,json=updateType,proto3,enum=alcor.schema.UpdateType" json:"update_type,omitempty"` // DELTA (default) or FULL
	VpcId              string                                          `protobuf:"bytes,5,opt,name=vpc_id,json=vpcId,proto3" json:"vpc_id,omitempty"`
	Name               string                                          `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	SecurityGroupRules []*SecurityGroupConfiguration_SecurityGroupRule `protobuf:"bytes,7,rep,name=security_group_rules,json=securityGroupRules,proto3" json:"security_group_rules,omitempty"`
}

func (x *SecurityGroupConfiguration) Reset() {
	*x = SecurityGroupConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_securitygroup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityGroupConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityGroupConfiguration) ProtoMessage() {}

func (x *SecurityGroupConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_securitygroup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityGroupConfiguration.ProtoReflect.Descriptor instead.
func (*SecurityGroupConfiguration) Descriptor() ([]byte, []int) {
	return file_securitygroup_proto_rawDescGZIP(), []int{0}
}

func (x *SecurityGroupConfiguration) GetRevisionNumber() uint32 {
	if x != nil {
		return x.RevisionNumber
	}
	return 0
}

func (x *SecurityGroupConfiguration) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *SecurityGroupConfiguration) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SecurityGroupConfiguration) GetUpdateType() UpdateType {
	if x != nil {
		return x.UpdateType
	}
	return UpdateType_DELTA
}

func (x *SecurityGroupConfiguration) GetVpcId() string {
	if x != nil {
		return x.VpcId
	}
	return ""
}

func (x *SecurityGroupConfiguration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecurityGroupConfiguration) GetSecurityGroupRules() []*SecurityGroupConfiguration_SecurityGroupRule {
	if x != nil {
		return x.SecurityGroupRules
	}
	return nil
}

type SecurityGroupState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationType OperationType               `protobuf:"varint,1,opt,name=operation_type,json=operationType,proto3,enum=alcor.schema.OperationType" json:"operation_type,omitempty"`
	Configuration *SecurityGroupConfiguration `protobuf:"bytes,2,opt,name=configuration,proto3" json:"configuration,omitempty"`
}

func (x *SecurityGroupState) Reset() {
	*x = SecurityGroupState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_securitygroup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityGroupState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityGroupState) ProtoMessage() {}

func (x *SecurityGroupState) ProtoReflect() protoreflect.Message {
	mi := &file_securitygroup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityGroupState.ProtoReflect.Descriptor instead.
func (*SecurityGroupState) Descriptor() ([]byte, []int) {
	return file_securitygroup_proto_rawDescGZIP(), []int{1}
}

func (x *SecurityGroupState) GetOperationType() OperationType {
	if x != nil {
		return x.OperationType
	}
	return OperationType_INFO
}

func (x *SecurityGroupState) GetConfiguration() *SecurityGroupConfiguration {
	if x != nil {
		return x.Configuration
	}
	return nil
}

type SecurityGroupConfiguration_SecurityGroupRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperationType   OperationType                        `protobuf:"varint,1,opt,name=operation_type,json=operationType,proto3,enum=alcor.schema.OperationType" json:"operation_type,omitempty"`
	SecurityGroupId string                               `protobuf:"bytes,2,opt,name=security_group_id,json=securityGroupId,proto3" json:"security_group_id,omitempty"`
	Id              string                               `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Direction       SecurityGroupConfiguration_Direction `protobuf:"varint,4,opt,name=direction,proto3,enum=alcor.schema.SecurityGroupConfiguration_Direction" json:"direction,omitempty"`
	Ethertype       EtherType                            `protobuf:"varint,5,opt,name=ethertype,proto3,enum=alcor.schema.EtherType" json:"ethertype,omitempty"`
	Protocol        Protocol                             `protobuf:"varint,6,opt,name=protocol,proto3,enum=alcor.schema.Protocol" json:"protocol,omitempty"`
	PortRangeMin    uint32                               `protobuf:"varint,7,opt,name=port_range_min,json=portRangeMin,proto3" json:"port_range_min,omitempty"`
	PortRangeMax    uint32                               `protobuf:"varint,8,opt,name=port_range_max,json=portRangeMax,proto3" json:"port_range_max,omitempty"`
	RemoteIpPrefix  string                               `protobuf:"bytes,9,opt,name=remote_ip_prefix,json=remoteIpPrefix,proto3" json:"remote_ip_prefix,omitempty"`
	RemoteGroupId   string                               `protobuf:"bytes,10,opt,name=remote_group_id,json=remoteGroupId,proto3" json:"remote_group_id,omitempty"`
}

func (x *SecurityGroupConfiguration_SecurityGroupRule) Reset() {
	*x = SecurityGroupConfiguration_SecurityGroupRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_securitygroup_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityGroupConfiguration_SecurityGroupRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityGroupConfiguration_SecurityGroupRule) ProtoMessage() {}

func (x *SecurityGroupConfiguration_SecurityGroupRule) ProtoReflect() protoreflect.Message {
	mi := &file_securitygroup_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityGroupConfiguration_SecurityGroupRule.ProtoReflect.Descriptor instead.
func (*SecurityGroupConfiguration_SecurityGroupRule) Descriptor() ([]byte, []int) {
	return file_securitygroup_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SecurityGroupConfiguration_SecurityGroupRule) GetOperationType() OperationType {
	if x != nil {
		return x.OperationType
	}
	return OperationType_INFO
}

func (x *SecurityGroupConfiguration_SecurityGroupRule) GetSecurityGroupId() string {
	if x != nil {
		return x.SecurityGroupId
	}
	return ""
}

func (x *SecurityGroupConfiguration_SecurityGroupRule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SecurityGroupConfiguration_SecurityGroupRule) GetDirection() SecurityGroupConfiguration_Direction {
	if x != nil {
		return x.Direction
	}
	return SecurityGroupConfiguration_EGRESS
}

func (x *SecurityGroupConfiguration_SecurityGroupRule) GetEthertype() EtherType {
	if x != nil {
		return x.Ethertype
	}
	return EtherType_IPV4
}

func (x *SecurityGroupConfiguration_SecurityGroupRule) GetProtocol() Protocol {
	if x != nil {
		return x.Protocol
	}
	return Protocol_TCP
}

func (x *SecurityGroupConfiguration_SecurityGroupRule) GetPortRangeMin() uint32 {
	if x != nil {
		return x.PortRangeMin
	}
	return 0
}

func (x *SecurityGroupConfiguration_SecurityGroupRule) GetPortRangeMax() uint32 {
	if x != nil {
		return x.PortRangeMax
	}
	return 0
}

func (x *SecurityGroupConfiguration_SecurityGroupRule) GetRemoteIpPrefix() string {
	if x != nil {
		return x.RemoteIpPrefix
	}
	return ""
}

func (x *SecurityGroupConfiguration_SecurityGroupRule) GetRemoteGroupId() string {
	if x != nil {
		return x.RemoteGroupId
	}
	return ""
}

var File_securitygroup_proto protoreflect.FileDescriptor

var file_securitygroup_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x6c, 0x63, 0x6f, 0x72, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdf, 0x06, 0x0a, 0x1a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x72, 0x65, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x61, 0x6c, 0x63, 0x6f, 0x72, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x76, 0x70, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x70, 0x63, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x6c,
	0x0a, 0x14, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x61,
	0x6c, 0x63, 0x6f, 0x72, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x12, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x1a, 0xee, 0x03, 0x0a,
	0x11, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x75,
	0x6c, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x61, 0x6c, 0x63,
	0x6f, 0x72, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x50, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e, 0x61, 0x6c, 0x63, 0x6f, 0x72, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x09, 0x65, 0x74, 0x68, 0x65, 0x72, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x61, 0x6c, 0x63, 0x6f, 0x72, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x45, 0x74, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x09, 0x65, 0x74, 0x68, 0x65, 0x72, 0x74, 0x79, 0x70, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e,
	0x61, 0x6c, 0x63, 0x6f, 0x72, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12,
	0x24, 0x0a, 0x0e, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x69,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x4d, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x72, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x6d, 0x61, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x61, 0x78, 0x12, 0x28, 0x0a, 0x10, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x70, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x49, 0x70, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x26, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x24, 0x0a,
	0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x47,
	0x52, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x47, 0x52, 0x45, 0x53,
	0x53, 0x10, 0x01, 0x22, 0xa8, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x61, 0x6c, 0x63, 0x6f, 0x72, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0d, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4e,
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x6c, 0x63, 0x6f, 0x72, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x54,
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x75, 0x74, 0x75, 0x72, 0x65, 0x77, 0x65, 0x69, 0x2e,
	0x61, 0x6c, 0x63, 0x6f, 0x72, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x42, 0x0d, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x5a, 0x27, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75, 0x74, 0x75, 0x72, 0x65, 0x77, 0x65,
	0x69, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x6c, 0x63, 0x6f, 0x72, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_securitygroup_proto_rawDescOnce sync.Once
	file_securitygroup_proto_rawDescData = file_securitygroup_proto_rawDesc
)

func file_securitygroup_proto_rawDescGZIP() []byte {
	file_securitygroup_proto_rawDescOnce.Do(func() {
		file_securitygroup_proto_rawDescData = protoimpl.X.CompressGZIP(file_securitygroup_proto_rawDescData)
	})
	return file_securitygroup_proto_rawDescData
}

var file_securitygroup_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_securitygroup_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_securitygroup_proto_goTypes = []interface{}{
	(SecurityGroupConfiguration_Direction)(0),            // 0: alcor.schema.SecurityGroupConfiguration.Direction
	(*SecurityGroupConfiguration)(nil),                   // 1: alcor.schema.SecurityGroupConfiguration
	(*SecurityGroupState)(nil),                           // 2: alcor.schema.SecurityGroupState
	(*SecurityGroupConfiguration_SecurityGroupRule)(nil), // 3: alcor.schema.SecurityGroupConfiguration.SecurityGroupRule
	(UpdateType)(0),                                      // 4: alcor.schema.UpdateType
	(OperationType)(0),                                   // 5: alcor.schema.OperationType
	(EtherType)(0),                                       // 6: alcor.schema.EtherType
	(Protocol)(0),                                        // 7: alcor.schema.Protocol
}
var file_securitygroup_proto_depIdxs = []int32{
	4, // 0: alcor.schema.SecurityGroupConfiguration.update_type:type_name -> alcor.schema.UpdateType
	3, // 1: alcor.schema.SecurityGroupConfiguration.security_group_rules:type_name -> alcor.schema.SecurityGroupConfiguration.SecurityGroupRule
	5, // 2: alcor.schema.SecurityGroupState.operation_type:type_name -> alcor.schema.OperationType
	1, // 3: alcor.schema.SecurityGroupState.configuration:type_name -> alcor.schema.SecurityGroupConfiguration
	5, // 4: alcor.schema.SecurityGroupConfiguration.SecurityGroupRule.operation_type:type_name -> alcor.schema.OperationType
	0, // 5: alcor.schema.SecurityGroupConfiguration.SecurityGroupRule.direction:type_name -> alcor.schema.SecurityGroupConfiguration.Direction
	6, // 6: alcor.schema.SecurityGroupConfiguration.SecurityGroupRule.ethertype:type_name -> alcor.schema.EtherType
	7, // 7: alcor.schema.SecurityGroupConfiguration.SecurityGroupRule.protocol:type_name -> alcor.schema.Protocol
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_securitygroup_proto_init() }
func file_securitygroup_proto_init() {
	if File_securitygroup_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_securitygroup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityGroupConfiguration); i {
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
		file_securitygroup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityGroupState); i {
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
		file_securitygroup_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityGroupConfiguration_SecurityGroupRule); i {
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
			RawDescriptor: file_securitygroup_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_securitygroup_proto_goTypes,
		DependencyIndexes: file_securitygroup_proto_depIdxs,
		EnumInfos:         file_securitygroup_proto_enumTypes,
		MessageInfos:      file_securitygroup_proto_msgTypes,
	}.Build()
	File_securitygroup_proto = out.File
	file_securitygroup_proto_rawDesc = nil
	file_securitygroup_proto_goTypes = nil
	file_securitygroup_proto_depIdxs = nil
}