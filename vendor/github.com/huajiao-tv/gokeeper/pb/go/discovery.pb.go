// Code generated by protoc-gen-go. DO NOT EDIT.
// discovery.proto is a deprecated file.

package pb

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/lnnujxxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//服务发现列表事件
type DiscoveryEventType int32

const (
	// type
	DiscoveryEventType__DISCOVERY_EVENT_TYPE_ DiscoveryEventType = 0
	//服务列表未更新
	DiscoveryEventType_DISCOVERY_EVENT_NONE DiscoveryEventType = 1
	//服务列表更新
	DiscoveryEventType_DISCOVERY_EVENT_UPDATE DiscoveryEventType = 2
)

var DiscoveryEventType_name = map[int32]string{
	0: "_DISCOVERY_EVENT_TYPE_",
	1: "DISCOVERY_EVENT_NONE",
	2: "DISCOVERY_EVENT_UPDATE",
}

var DiscoveryEventType_value = map[string]int32{
	"_DISCOVERY_EVENT_TYPE_": 0,
	"DISCOVERY_EVENT_NONE":   1,
	"DISCOVERY_EVENT_UPDATE": 2,
}

func (x DiscoveryEventType) String() string {
	return proto.EnumName(DiscoveryEventType_name, int32(x))
}

func (DiscoveryEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{0}
}

//实例信息,可参考Instance结构体
type Instance struct {
	//实例id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	//服务名称
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	//机房，bjcc/bjyt/bjdt
	Zone string `protobuf:"bytes,3,opt,name=zone,proto3" json:"zone,omitempty"`
	//环境，prod(线上)/pre(预发布)/test(测试)/dev(开发)
	Env string `protobuf:"bytes,4,opt,name=env,proto3" json:"env,omitempty"`
	//机器地址
	Hostname string `protobuf:"bytes,5,opt,name=hostname,proto3" json:"hostname,omitempty"`
	//地址 scheme -> address  例如： http -> 127.0.0.1:8080、grpc -> 127.0.0.1:8081
	Addrs map[string]string `protobuf:"bytes,6,rep,name=addrs,proto3" json:"addrs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//元组数据，例如weight权重信息
	Metadata map[string]string `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//注册时间
	RegTime int64 `protobuf:"varint,8,opt,name=reg_time,json=regTime,proto3" json:"reg_time,omitempty"`
	//更新时间
	UpdateTime           int64    `protobuf:"varint,9,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}
func (*Instance) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{0}
}

func (m *Instance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instance.Unmarshal(m, b)
}
func (m *Instance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instance.Marshal(b, m, deterministic)
}
func (m *Instance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instance.Merge(m, src)
}
func (m *Instance) XXX_Size() int {
	return xxx_messageInfo_Instance.Size(m)
}
func (m *Instance) XXX_DiscardUnknown() {
	xxx_messageInfo_Instance.DiscardUnknown(m)
}

var xxx_messageInfo_Instance proto.InternalMessageInfo

func (m *Instance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Instance) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Instance) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *Instance) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

func (m *Instance) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *Instance) GetAddrs() map[string]string {
	if m != nil {
		return m.Addrs
	}
	return nil
}

func (m *Instance) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Instance) GetRegTime() int64 {
	if m != nil {
		return m.RegTime
	}
	return 0
}

func (m *Instance) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

//实例群组
type ZoneGroup struct {
	//实例群组
	ZoneGroup            map[string]*Instance `protobuf:"bytes,1,rep,name=zone_group,json=zoneGroup,proto3" json:"zone_group,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ZoneGroup) Reset()         { *m = ZoneGroup{} }
func (m *ZoneGroup) String() string { return proto.CompactTextString(m) }
func (*ZoneGroup) ProtoMessage()    {}
func (*ZoneGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{1}
}

func (m *ZoneGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZoneGroup.Unmarshal(m, b)
}
func (m *ZoneGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZoneGroup.Marshal(b, m, deterministic)
}
func (m *ZoneGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZoneGroup.Merge(m, src)
}
func (m *ZoneGroup) XXX_Size() int {
	return xxx_messageInfo_ZoneGroup.Size(m)
}
func (m *ZoneGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_ZoneGroup.DiscardUnknown(m)
}

var xxx_messageInfo_ZoneGroup proto.InternalMessageInfo

func (m *ZoneGroup) GetZoneGroup() map[string]*Instance {
	if m != nil {
		return m.ZoneGroup
	}
	return nil
}

//服务信息，可参考Service结构体
type Service struct {
	//服务名称，project.service[.sub_service],例如live.session
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	//实例信息，根据zone（机房）进行分组
	Instances map[string]*ZoneGroup `protobuf:"bytes,2,rep,name=instances,proto3" json:"instances,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//元组数据，用于扩展
	Metadata map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//service更新时间
	UpdateTime int64 `protobuf:"varint,4,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	//service版本
	Version              int64    `protobuf:"varint,5,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Service) Reset()         { *m = Service{} }
func (m *Service) String() string { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()    {}
func (*Service) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{2}
}

func (m *Service) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Service.Unmarshal(m, b)
}
func (m *Service) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Service.Marshal(b, m, deterministic)
}
func (m *Service) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Service.Merge(m, src)
}
func (m *Service) XXX_Size() int {
	return xxx_messageInfo_Service.Size(m)
}
func (m *Service) XXX_DiscardUnknown() {
	xxx_messageInfo_Service.DiscardUnknown(m)
}

var xxx_messageInfo_Service proto.InternalMessageInfo

func (m *Service) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *Service) GetInstances() map[string]*ZoneGroup {
	if m != nil {
		return m.Instances
	}
	return nil
}

func (m *Service) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Service) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func (m *Service) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

//服务注册请求体
type RegisterReq struct {
	//实例信息
	Instance *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	//实例续约时间，单位秒
	LeaseSecond          int64    `protobuf:"varint,2,opt,name=lease_second,json=leaseSecond,proto3" json:"lease_second,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{3}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *RegisterReq) GetLeaseSecond() int64 {
	if m != nil {
		return m.LeaseSecond
	}
	return 0
}

//服务周期续约请求体
type KeepAliveReq struct {
	//实例信息
	Instance *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	//实例续约时间，单位秒
	LeaseSecond          int64    `protobuf:"varint,2,opt,name=lease_second,json=leaseSecond,proto3" json:"lease_second,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeepAliveReq) Reset()         { *m = KeepAliveReq{} }
func (m *KeepAliveReq) String() string { return proto.CompactTextString(m) }
func (*KeepAliveReq) ProtoMessage()    {}
func (*KeepAliveReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{4}
}

func (m *KeepAliveReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeepAliveReq.Unmarshal(m, b)
}
func (m *KeepAliveReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeepAliveReq.Marshal(b, m, deterministic)
}
func (m *KeepAliveReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeepAliveReq.Merge(m, src)
}
func (m *KeepAliveReq) XXX_Size() int {
	return xxx_messageInfo_KeepAliveReq.Size(m)
}
func (m *KeepAliveReq) XXX_DiscardUnknown() {
	xxx_messageInfo_KeepAliveReq.DiscardUnknown(m)
}

var xxx_messageInfo_KeepAliveReq proto.InternalMessageInfo

func (m *KeepAliveReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *KeepAliveReq) GetLeaseSecond() int64 {
	if m != nil {
		return m.LeaseSecond
	}
	return 0
}

//服务解除注册请求体
type DeregisterReq struct {
	//实例信息
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DeregisterReq) Reset()         { *m = DeregisterReq{} }
func (m *DeregisterReq) String() string { return proto.CompactTextString(m) }
func (*DeregisterReq) ProtoMessage()    {}
func (*DeregisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{5}
}

func (m *DeregisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeregisterReq.Unmarshal(m, b)
}
func (m *DeregisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeregisterReq.Marshal(b, m, deterministic)
}
func (m *DeregisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeregisterReq.Merge(m, src)
}
func (m *DeregisterReq) XXX_Size() int {
	return xxx_messageInfo_DeregisterReq.Size(m)
}
func (m *DeregisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeregisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeregisterReq proto.InternalMessageInfo

func (m *DeregisterReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

//批量获取服务节点的请求体
type PollsReq struct {
	//服务信息 service_name -> version
	PollServices map[string]int64 `protobuf:"bytes,1,rep,name=poll_services,json=pollServices,proto3" json:"poll_services,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	//环境
	Env string `protobuf:"bytes,2,opt,name=env,proto3" json:"env,omitempty"`
	//订阅者
	Subscriber           string   `protobuf:"bytes,3,opt,name=subscriber,proto3" json:"subscriber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PollsReq) Reset()         { *m = PollsReq{} }
func (m *PollsReq) String() string { return proto.CompactTextString(m) }
func (*PollsReq) ProtoMessage()    {}
func (*PollsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{6}
}

func (m *PollsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollsReq.Unmarshal(m, b)
}
func (m *PollsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollsReq.Marshal(b, m, deterministic)
}
func (m *PollsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollsReq.Merge(m, src)
}
func (m *PollsReq) XXX_Size() int {
	return xxx_messageInfo_PollsReq.Size(m)
}
func (m *PollsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PollsReq.DiscardUnknown(m)
}

var xxx_messageInfo_PollsReq proto.InternalMessageInfo

func (m *PollsReq) GetPollServices() map[string]int64 {
	if m != nil {
		return m.PollServices
	}
	return nil
}

func (m *PollsReq) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

func (m *PollsReq) GetSubscriber() string {
	if m != nil {
		return m.Subscriber
	}
	return ""
}

//批量获取服务节点的返回体，方便后续扩展字段，单独声明一个结构体
type PollsResp struct {
	//服务发现事件类型
	EventType DiscoveryEventType `protobuf:"varint,1,opt,name=event_type,json=eventType,proto3,enum=keeper.DiscoveryEventType" json:"event_type,omitempty"`
	//服务信息，service_name -> service
	Services             map[string]*Service `protobuf:"bytes,2,rep,name=services,proto3" json:"services,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PollsResp) Reset()         { *m = PollsResp{} }
func (m *PollsResp) String() string { return proto.CompactTextString(m) }
func (*PollsResp) ProtoMessage()    {}
func (*PollsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e7ff60feb39c8d0, []int{7}
}

func (m *PollsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PollsResp.Unmarshal(m, b)
}
func (m *PollsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PollsResp.Marshal(b, m, deterministic)
}
func (m *PollsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PollsResp.Merge(m, src)
}
func (m *PollsResp) XXX_Size() int {
	return xxx_messageInfo_PollsResp.Size(m)
}
func (m *PollsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_PollsResp.DiscardUnknown(m)
}

var xxx_messageInfo_PollsResp proto.InternalMessageInfo

func (m *PollsResp) GetEventType() DiscoveryEventType {
	if m != nil {
		return m.EventType
	}
	return DiscoveryEventType__DISCOVERY_EVENT_TYPE_
}

func (m *PollsResp) GetServices() map[string]*Service {
	if m != nil {
		return m.Services
	}
	return nil
}

func init() {
	proto.RegisterEnum("keeper.DiscoveryEventType", DiscoveryEventType_name, DiscoveryEventType_value)
	proto.RegisterType((*Instance)(nil), "keeper.Instance")
	proto.RegisterMapType((map[string]string)(nil), "keeper.Instance.AddrsEntry")
	proto.RegisterMapType((map[string]string)(nil), "keeper.Instance.MetadataEntry")
	proto.RegisterType((*ZoneGroup)(nil), "keeper.ZoneGroup")
	proto.RegisterMapType((map[string]*Instance)(nil), "keeper.ZoneGroup.ZoneGroupEntry")
	proto.RegisterType((*Service)(nil), "keeper.Service")
	proto.RegisterMapType((map[string]*ZoneGroup)(nil), "keeper.Service.InstancesEntry")
	proto.RegisterMapType((map[string]string)(nil), "keeper.Service.MetadataEntry")
	proto.RegisterType((*RegisterReq)(nil), "keeper.RegisterReq")
	proto.RegisterType((*KeepAliveReq)(nil), "keeper.KeepAliveReq")
	proto.RegisterType((*DeregisterReq)(nil), "keeper.DeregisterReq")
	proto.RegisterType((*PollsReq)(nil), "keeper.PollsReq")
	proto.RegisterMapType((map[string]int64)(nil), "keeper.PollsReq.PollServicesEntry")
	proto.RegisterType((*PollsResp)(nil), "keeper.PollsResp")
	proto.RegisterMapType((map[string]*Service)(nil), "keeper.PollsResp.ServicesEntry")
}

func init() { proto.RegisterFile("discovery.proto", fileDescriptor_1e7ff60feb39c8d0) }

var fileDescriptor_1e7ff60feb39c8d0 = []byte{
	// 863 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x51, 0x8f, 0xdb, 0x44,
	0x10, 0xbe, 0xb5, 0x2f, 0x77, 0xf6, 0xe4, 0x2e, 0xcd, 0x2d, 0x47, 0x31, 0x2e, 0xb4, 0x47, 0x24,
	0xe0, 0x84, 0x90, 0x0f, 0x82, 0x90, 0xda, 0x1e, 0x55, 0x75, 0xd7, 0xb3, 0xa0, 0x02, 0x72, 0x91,
	0x13, 0x2a, 0xb5, 0x0f, 0x58, 0x4e, 0x3c, 0x0d, 0x56, 0x1d, 0xaf, 0xf1, 0x3a, 0x96, 0xd2, 0x9f,
	0x82, 0x78, 0x81, 0x5f, 0x81, 0x78, 0xe2, 0x2f, 0xc0, 0xbf, 0xe0, 0x85, 0x47, 0x9e, 0x91, 0xd7,
	0x5e, 0xc7, 0x71, 0x8e, 0x56, 0x20, 0x78, 0x9b, 0x9d, 0x99, 0x6f, 0x77, 0xe6, 0xfb, 0x66, 0x77,
	0xe1, 0x9a, 0x1f, 0xf0, 0x29, 0xcb, 0x30, 0x59, 0x5a, 0x71, 0xc2, 0x52, 0x46, 0x77, 0x9e, 0x21,
	0xc6, 0x98, 0x98, 0x6f, 0xcc, 0x18, 0x9b, 0x85, 0x78, 0xe2, 0xc5, 0xc1, 0x89, 0x17, 0x45, 0x2c,
	0xf5, 0xd2, 0x80, 0x45, 0xbc, 0xc8, 0x32, 0x6f, 0x94, 0x51, 0xb1, 0x9a, 0x2c, 0x9e, 0x9e, 0xe0,
	0x3c, 0x4e, 0xcb, 0x2d, 0xcc, 0xd7, 0x32, 0x2f, 0x0c, 0x7c, 0x2f, 0xc5, 0x13, 0x69, 0x14, 0x81,
	0xde, 0x0f, 0x2a, 0x68, 0x0f, 0x23, 0x9e, 0x7a, 0xd1, 0x14, 0x69, 0x07, 0x94, 0xc0, 0x37, 0xc8,
	0x11, 0x39, 0xd6, 0x1d, 0x25, 0xf0, 0xe9, 0x5b, 0xb0, 0xc7, 0x31, 0xc9, 0x82, 0x29, 0xba, 0x91,
	0x37, 0x47, 0x43, 0x11, 0x91, 0x76, 0xe9, 0x1b, 0x78, 0x73, 0xa4, 0x14, 0xb6, 0x9f, 0xb3, 0x08,
	0x0d, 0x55, 0x84, 0x84, 0x4d, 0xbb, 0xa0, 0x62, 0x94, 0x19, 0xdb, 0xc2, 0x95, 0x9b, 0xd4, 0x04,
	0xed, 0x1b, 0xc6, 0x53, 0xb1, 0x49, 0x4b, 0xb8, 0xab, 0x35, 0xfd, 0x10, 0x5a, 0x9e, 0xef, 0x27,
	0xdc, 0xd8, 0x39, 0x52, 0x8f, 0xdb, 0xfd, 0x1b, 0x56, 0xd1, 0xad, 0x25, 0xab, 0xb2, 0xce, 0xf2,
	0xa8, 0x1d, 0xa5, 0xc9, 0xd2, 0x29, 0x32, 0xe9, 0x5d, 0xd0, 0xe6, 0x98, 0x7a, 0xbe, 0x97, 0x7a,
	0xc6, 0xae, 0x40, 0xdd, 0xdc, 0x40, 0x7d, 0x59, 0x26, 0x14, 0xc0, 0x2a, 0x9f, 0xbe, 0x0e, 0x5a,
	0x82, 0x33, 0x37, 0x0d, 0xe6, 0x68, 0x68, 0x47, 0xe4, 0x58, 0x75, 0x76, 0x13, 0x9c, 0x8d, 0x83,
	0x39, 0xd2, 0x5b, 0xd0, 0x5e, 0xc4, 0x39, 0x37, 0x45, 0x54, 0x17, 0x51, 0x28, 0x5c, 0x79, 0x82,
	0x79, 0x1b, 0x60, 0x55, 0x4c, 0xde, 0xe6, 0x33, 0x5c, 0x96, 0x74, 0xe5, 0x26, 0x3d, 0x84, 0x56,
	0xe6, 0x85, 0x0b, 0x49, 0x54, 0xb1, 0xb8, 0xab, 0xdc, 0x26, 0xe6, 0x29, 0xec, 0xaf, 0x15, 0xf4,
	0x4f, 0xc0, 0xbd, 0xef, 0x09, 0xe8, 0x4f, 0x58, 0x84, 0x9f, 0x26, 0x6c, 0x11, 0xd3, 0xfb, 0x00,
	0x39, 0xcb, 0xee, 0x2c, 0x5f, 0x19, 0x44, 0xb4, 0x7f, 0x24, 0xdb, 0xaf, 0xd2, 0x56, 0x56, 0x41,
	0x80, 0xfe, 0x5c, 0xae, 0xcd, 0x01, 0x74, 0xd6, 0x83, 0x57, 0x14, 0xf3, 0x4e, 0xbd, 0x98, 0x76,
	0xbf, 0xdb, 0xa4, 0xb7, 0x5e, 0xde, 0x1f, 0x0a, 0xec, 0x8e, 0x8a, 0x91, 0xd8, 0x98, 0x18, 0xb2,
	0x39, 0x31, 0x9f, 0x80, 0x1e, 0x94, 0xbb, 0x70, 0x43, 0x59, 0x57, 0xaf, 0xdc, 0xa6, 0x3a, 0xa6,
	0x94, 0x7d, 0x05, 0xa0, 0x77, 0x6a, 0xd2, 0xab, 0x02, 0xfc, 0x66, 0x13, 0xfc, 0x77, 0xca, 0x37,
	0xe4, 0xdd, 0x6e, 0xca, 0x4b, 0x0d, 0xd8, 0xcd, 0x30, 0xe1, 0x01, 0x8b, 0xc4, 0x90, 0xaa, 0x8e,
	0x5c, 0x9a, 0x97, 0xd0, 0x59, 0x2f, 0xe9, 0x0a, 0xca, 0xde, 0x5d, 0xa7, 0xec, 0x60, 0x43, 0x92,
	0xff, 0x6c, 0x1e, 0xbe, 0x86, 0xb6, 0x83, 0xb3, 0x80, 0xa7, 0x98, 0x38, 0xf8, 0x2d, 0x7d, 0x1f,
	0x34, 0xc9, 0x8f, 0xc0, 0x5f, 0x25, 0x57, 0x95, 0x91, 0x2b, 0x14, 0xa2, 0xc7, 0xd1, 0xe5, 0x38,
	0x65, 0x91, 0x2f, 0x76, 0x57, 0x9d, 0xb6, 0xf0, 0x8d, 0x84, 0xab, 0xe7, 0xc2, 0xde, 0xe7, 0x88,
	0xf1, 0x59, 0x18, 0x64, 0xf8, 0xbf, 0x1c, 0x70, 0x0f, 0xf6, 0x2f, 0x30, 0xf9, 0xb7, 0x2d, 0xf4,
	0x7e, 0x23, 0xa0, 0x0d, 0x59, 0x18, 0xf2, 0x1c, 0x3a, 0x82, 0xfd, 0x98, 0x85, 0xa1, 0x5b, 0x8e,
	0x18, 0x2f, 0x6f, 0x44, 0x4f, 0xe2, 0x65, 0xa2, 0x30, 0xca, 0x11, 0x29, 0x34, 0x3c, 0x87, 0x9f,
	0x7f, 0xff, 0x45, 0x6d, 0x7d, 0x47, 0x14, 0x8d, 0x38, 0x7b, 0x71, 0x2d, 0x2c, 0x5f, 0x30, 0x65,
	0xf5, 0x82, 0xdd, 0x04, 0xe0, 0x8b, 0x09, 0x9f, 0x26, 0xc1, 0x04, 0x93, 0xf2, 0xb5, 0xab, 0x79,
	0xcc, 0xfb, 0x70, 0xb0, 0x71, 0xc0, 0xcb, 0x44, 0x55, 0xeb, 0xa2, 0xfe, 0x4a, 0x40, 0x2f, 0x6b,
	0xe5, 0x31, 0xbd, 0x03, 0x80, 0x19, 0x46, 0xa9, 0x9b, 0x2e, 0xe3, 0x82, 0x92, 0x4e, 0xdf, 0x94,
	0x2d, 0x5d, 0xc8, 0xff, 0xc1, 0xce, 0x53, 0xc6, 0xcb, 0x18, 0x1d, 0x1d, 0xa5, 0x49, 0x4f, 0x41,
	0xab, 0xb8, 0x28, 0xae, 0xd7, 0xad, 0x06, 0x17, 0x3c, 0xb6, 0xd6, 0xea, 0x74, 0x2a, 0x80, 0xf9,
	0x05, 0xec, 0xbf, 0xac, 0x85, 0xb7, 0xd7, 0xe7, 0xfc, 0x5a, 0xe3, 0xfa, 0xd5, 0x7a, 0x7a, 0xef,
	0x29, 0xd0, 0xcd, 0x5a, 0xa9, 0x09, 0xd7, 0xdd, 0x8b, 0x87, 0xa3, 0x07, 0x97, 0x8f, 0x6c, 0xe7,
	0xb1, 0x6b, 0x3f, 0xb2, 0x07, 0x63, 0x77, 0xfc, 0x78, 0x68, 0xbb, 0xdd, 0x2d, 0x6a, 0xc0, 0x61,
	0x33, 0x34, 0xb8, 0x1c, 0xd8, 0x5d, 0x92, 0xa3, 0x9a, 0x91, 0xaf, 0x86, 0x17, 0x67, 0x63, 0xbb,
	0xab, 0xf4, 0xff, 0x24, 0xa0, 0x57, 0x07, 0xe5, 0x4f, 0x84, 0xbc, 0x1e, 0xf4, 0x15, 0x59, 0x5d,
	0xed, 0xc2, 0x98, 0xd7, 0xad, 0xe2, 0xab, 0xb4, 0xe4, 0x57, 0x69, 0xd9, 0xf9, 0x57, 0xd9, 0xdb,
	0xa2, 0xa7, 0xa0, 0x57, 0x93, 0x4f, 0x0f, 0x25, 0xb6, 0x7e, 0x19, 0x5e, 0x00, 0xbe, 0x07, 0xb0,
	0x9a, 0x6a, 0xfa, 0x6a, 0xa5, 0x56, 0x7d, 0xd2, 0x5f, 0x00, 0xef, 0x43, 0x4b, 0xe8, 0x43, 0xbb,
	0xcd, 0xd1, 0x35, 0x0f, 0x36, 0x04, 0xec, 0x6d, 0x1d, 0x93, 0x0f, 0xc8, 0xf9, 0xc7, 0x40, 0xa7,
	0x6c, 0x6e, 0x95, 0xf2, 0x95, 0x59, 0xe7, 0x9d, 0x8a, 0x8b, 0x61, 0x7e, 0xc8, 0x67, 0x64, 0x48,
	0x9e, 0x28, 0xf1, 0xe4, 0x27, 0x42, 0x7e, 0x54, 0xd4, 0x07, 0xc3, 0xf3, 0xc9, 0x8e, 0x38, 0xfc,
	0xa3, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x95, 0xf4, 0x20, 0x04, 0x6a, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DiscoveryClient is the client API for Discovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscoveryClient interface {
	//服务注册接口,服务就绪后调用
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*empty.Empty, error)
	//服务心跳接口
	KeepAlive(ctx context.Context, in *KeepAliveReq, opts ...grpc.CallOption) (*empty.Empty, error)
	//服务解除注册接口，服务下线时调用
	Deregister(ctx context.Context, in *DeregisterReq, opts ...grpc.CallOption) (*empty.Empty, error)
	//批量获取服务的节点信息,如果没有相关数据更新，则阻塞一段时间
	Polls(ctx context.Context, opts ...grpc.CallOption) (Discovery_PollsClient, error)
}

type discoveryClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryClient(cc *grpc.ClientConn) DiscoveryClient {
	return &discoveryClient{cc}
}

func (c *discoveryClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/keeper.Discovery/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) KeepAlive(ctx context.Context, in *KeepAliveReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/keeper.Discovery/KeepAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) Deregister(ctx context.Context, in *DeregisterReq, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/keeper.Discovery/Deregister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryClient) Polls(ctx context.Context, opts ...grpc.CallOption) (Discovery_PollsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Discovery_serviceDesc.Streams[0], "/keeper.Discovery/Polls", opts...)
	if err != nil {
		return nil, err
	}
	x := &discoveryPollsClient{stream}
	return x, nil
}

type Discovery_PollsClient interface {
	Send(*PollsReq) error
	Recv() (*PollsResp, error)
	grpc.ClientStream
}

type discoveryPollsClient struct {
	grpc.ClientStream
}

func (x *discoveryPollsClient) Send(m *PollsReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *discoveryPollsClient) Recv() (*PollsResp, error) {
	m := new(PollsResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DiscoveryServer is the server API for Discovery service.
type DiscoveryServer interface {
	//服务注册接口,服务就绪后调用
	Register(context.Context, *RegisterReq) (*empty.Empty, error)
	//服务心跳接口
	KeepAlive(context.Context, *KeepAliveReq) (*empty.Empty, error)
	//服务解除注册接口，服务下线时调用
	Deregister(context.Context, *DeregisterReq) (*empty.Empty, error)
	//批量获取服务的节点信息,如果没有相关数据更新，则阻塞一段时间
	Polls(Discovery_PollsServer) error
}

// UnimplementedDiscoveryServer can be embedded to have forward compatible implementations.
type UnimplementedDiscoveryServer struct {
}

func (*UnimplementedDiscoveryServer) Register(ctx context.Context, req *RegisterReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedDiscoveryServer) KeepAlive(ctx context.Context, req *KeepAliveReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeepAlive not implemented")
}
func (*UnimplementedDiscoveryServer) Deregister(ctx context.Context, req *DeregisterReq) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deregister not implemented")
}
func (*UnimplementedDiscoveryServer) Polls(srv Discovery_PollsServer) error {
	return status.Errorf(codes.Unimplemented, "method Polls not implemented")
}

func RegisterDiscoveryServer(s *grpc.Server, srv DiscoveryServer) {
	s.RegisterService(&_Discovery_serviceDesc, srv)
}

func _Discovery_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.Discovery/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeepAliveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.Discovery/KeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).KeepAlive(ctx, req.(*KeepAliveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/keeper.Discovery/Deregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServer).Deregister(ctx, req.(*DeregisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Discovery_Polls_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DiscoveryServer).Polls(&discoveryPollsServer{stream})
}

type Discovery_PollsServer interface {
	Send(*PollsResp) error
	Recv() (*PollsReq, error)
	grpc.ServerStream
}

type discoveryPollsServer struct {
	grpc.ServerStream
}

func (x *discoveryPollsServer) Send(m *PollsResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *discoveryPollsServer) Recv() (*PollsReq, error) {
	m := new(PollsReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Discovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "keeper.Discovery",
	HandlerType: (*DiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Discovery_Register_Handler,
		},
		{
			MethodName: "KeepAlive",
			Handler:    _Discovery_KeepAlive_Handler,
		},
		{
			MethodName: "Deregister",
			Handler:    _Discovery_Deregister_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Polls",
			Handler:       _Discovery_Polls_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "discovery.proto",
}
