// Code generated by protoc-gen-go. DO NOT EDIT.
// source: board.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type BoardDetailsReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Fqbn                 string    `protobuf:"bytes,2,opt,name=fqbn,proto3" json:"fqbn,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BoardDetailsReq) Reset()         { *m = BoardDetailsReq{} }
func (m *BoardDetailsReq) String() string { return proto.CompactTextString(m) }
func (*BoardDetailsReq) ProtoMessage()    {}
func (*BoardDetailsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_937f74b042f92c0f, []int{0}
}

func (m *BoardDetailsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardDetailsReq.Unmarshal(m, b)
}
func (m *BoardDetailsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardDetailsReq.Marshal(b, m, deterministic)
}
func (m *BoardDetailsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardDetailsReq.Merge(m, src)
}
func (m *BoardDetailsReq) XXX_Size() int {
	return xxx_messageInfo_BoardDetailsReq.Size(m)
}
func (m *BoardDetailsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardDetailsReq.DiscardUnknown(m)
}

var xxx_messageInfo_BoardDetailsReq proto.InternalMessageInfo

func (m *BoardDetailsReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

func (m *BoardDetailsReq) GetFqbn() string {
	if m != nil {
		return m.Fqbn
	}
	return ""
}

type BoardDetailsResp struct {
	Name                 string          `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ConfigOptions        []*ConfigOption `protobuf:"bytes,3,rep,name=config_options,json=configOptions,proto3" json:"config_options,omitempty"`
	RequiredTools        []*RequiredTool `protobuf:"bytes,4,rep,name=required_tools,json=requiredTools,proto3" json:"required_tools,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *BoardDetailsResp) Reset()         { *m = BoardDetailsResp{} }
func (m *BoardDetailsResp) String() string { return proto.CompactTextString(m) }
func (*BoardDetailsResp) ProtoMessage()    {}
func (*BoardDetailsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_937f74b042f92c0f, []int{1}
}

func (m *BoardDetailsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardDetailsResp.Unmarshal(m, b)
}
func (m *BoardDetailsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardDetailsResp.Marshal(b, m, deterministic)
}
func (m *BoardDetailsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardDetailsResp.Merge(m, src)
}
func (m *BoardDetailsResp) XXX_Size() int {
	return xxx_messageInfo_BoardDetailsResp.Size(m)
}
func (m *BoardDetailsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardDetailsResp.DiscardUnknown(m)
}

var xxx_messageInfo_BoardDetailsResp proto.InternalMessageInfo

func (m *BoardDetailsResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BoardDetailsResp) GetConfigOptions() []*ConfigOption {
	if m != nil {
		return m.ConfigOptions
	}
	return nil
}

func (m *BoardDetailsResp) GetRequiredTools() []*RequiredTool {
	if m != nil {
		return m.RequiredTools
	}
	return nil
}

type ConfigOption struct {
	Option               string         `protobuf:"bytes,1,opt,name=option,proto3" json:"option,omitempty"`
	OptionLabel          string         `protobuf:"bytes,2,opt,name=option_label,json=optionLabel,proto3" json:"option_label,omitempty"`
	Values               []*ConfigValue `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ConfigOption) Reset()         { *m = ConfigOption{} }
func (m *ConfigOption) String() string { return proto.CompactTextString(m) }
func (*ConfigOption) ProtoMessage()    {}
func (*ConfigOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_937f74b042f92c0f, []int{2}
}

func (m *ConfigOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigOption.Unmarshal(m, b)
}
func (m *ConfigOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigOption.Marshal(b, m, deterministic)
}
func (m *ConfigOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigOption.Merge(m, src)
}
func (m *ConfigOption) XXX_Size() int {
	return xxx_messageInfo_ConfigOption.Size(m)
}
func (m *ConfigOption) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigOption.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigOption proto.InternalMessageInfo

func (m *ConfigOption) GetOption() string {
	if m != nil {
		return m.Option
	}
	return ""
}

func (m *ConfigOption) GetOptionLabel() string {
	if m != nil {
		return m.OptionLabel
	}
	return ""
}

func (m *ConfigOption) GetValues() []*ConfigValue {
	if m != nil {
		return m.Values
	}
	return nil
}

type ConfigValue struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	ValueLabel           string   `protobuf:"bytes,2,opt,name=value_label,json=valueLabel,proto3" json:"value_label,omitempty"`
	Selected             bool     `protobuf:"varint,3,opt,name=selected,proto3" json:"selected,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigValue) Reset()         { *m = ConfigValue{} }
func (m *ConfigValue) String() string { return proto.CompactTextString(m) }
func (*ConfigValue) ProtoMessage()    {}
func (*ConfigValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_937f74b042f92c0f, []int{3}
}

func (m *ConfigValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigValue.Unmarshal(m, b)
}
func (m *ConfigValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigValue.Marshal(b, m, deterministic)
}
func (m *ConfigValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigValue.Merge(m, src)
}
func (m *ConfigValue) XXX_Size() int {
	return xxx_messageInfo_ConfigValue.Size(m)
}
func (m *ConfigValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigValue.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigValue proto.InternalMessageInfo

func (m *ConfigValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ConfigValue) GetValueLabel() string {
	if m != nil {
		return m.ValueLabel
	}
	return ""
}

func (m *ConfigValue) GetSelected() bool {
	if m != nil {
		return m.Selected
	}
	return false
}

type RequiredTool struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Packager             string   `protobuf:"bytes,3,opt,name=packager,proto3" json:"packager,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequiredTool) Reset()         { *m = RequiredTool{} }
func (m *RequiredTool) String() string { return proto.CompactTextString(m) }
func (*RequiredTool) ProtoMessage()    {}
func (*RequiredTool) Descriptor() ([]byte, []int) {
	return fileDescriptor_937f74b042f92c0f, []int{4}
}

func (m *RequiredTool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequiredTool.Unmarshal(m, b)
}
func (m *RequiredTool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequiredTool.Marshal(b, m, deterministic)
}
func (m *RequiredTool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequiredTool.Merge(m, src)
}
func (m *RequiredTool) XXX_Size() int {
	return xxx_messageInfo_RequiredTool.Size(m)
}
func (m *RequiredTool) XXX_DiscardUnknown() {
	xxx_messageInfo_RequiredTool.DiscardUnknown(m)
}

var xxx_messageInfo_RequiredTool proto.InternalMessageInfo

func (m *RequiredTool) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RequiredTool) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RequiredTool) GetPackager() string {
	if m != nil {
		return m.Packager
	}
	return ""
}

type BoardListReq struct {
	Instance             *Instance `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BoardListReq) Reset()         { *m = BoardListReq{} }
func (m *BoardListReq) String() string { return proto.CompactTextString(m) }
func (*BoardListReq) ProtoMessage()    {}
func (*BoardListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_937f74b042f92c0f, []int{5}
}

func (m *BoardListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardListReq.Unmarshal(m, b)
}
func (m *BoardListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardListReq.Marshal(b, m, deterministic)
}
func (m *BoardListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardListReq.Merge(m, src)
}
func (m *BoardListReq) XXX_Size() int {
	return xxx_messageInfo_BoardListReq.Size(m)
}
func (m *BoardListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardListReq.DiscardUnknown(m)
}

var xxx_messageInfo_BoardListReq proto.InternalMessageInfo

func (m *BoardListReq) GetInstance() *Instance {
	if m != nil {
		return m.Instance
	}
	return nil
}

type BoardListResp struct {
	Serial               []*AttachedSerialBoard  `protobuf:"bytes,1,rep,name=serial,proto3" json:"serial,omitempty"`
	Network              []*AttachedNetworkBoard `protobuf:"bytes,2,rep,name=network,proto3" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *BoardListResp) Reset()         { *m = BoardListResp{} }
func (m *BoardListResp) String() string { return proto.CompactTextString(m) }
func (*BoardListResp) ProtoMessage()    {}
func (*BoardListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_937f74b042f92c0f, []int{6}
}

func (m *BoardListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoardListResp.Unmarshal(m, b)
}
func (m *BoardListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoardListResp.Marshal(b, m, deterministic)
}
func (m *BoardListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoardListResp.Merge(m, src)
}
func (m *BoardListResp) XXX_Size() int {
	return xxx_messageInfo_BoardListResp.Size(m)
}
func (m *BoardListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BoardListResp.DiscardUnknown(m)
}

var xxx_messageInfo_BoardListResp proto.InternalMessageInfo

func (m *BoardListResp) GetSerial() []*AttachedSerialBoard {
	if m != nil {
		return m.Serial
	}
	return nil
}

func (m *BoardListResp) GetNetwork() []*AttachedNetworkBoard {
	if m != nil {
		return m.Network
	}
	return nil
}

type AttachedNetworkBoard struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Fqbn                 string   `protobuf:"bytes,2,opt,name=fqbn,proto3" json:"fqbn,omitempty"`
	Info                 string   `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Port                 uint64   `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttachedNetworkBoard) Reset()         { *m = AttachedNetworkBoard{} }
func (m *AttachedNetworkBoard) String() string { return proto.CompactTextString(m) }
func (*AttachedNetworkBoard) ProtoMessage()    {}
func (*AttachedNetworkBoard) Descriptor() ([]byte, []int) {
	return fileDescriptor_937f74b042f92c0f, []int{7}
}

func (m *AttachedNetworkBoard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachedNetworkBoard.Unmarshal(m, b)
}
func (m *AttachedNetworkBoard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachedNetworkBoard.Marshal(b, m, deterministic)
}
func (m *AttachedNetworkBoard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachedNetworkBoard.Merge(m, src)
}
func (m *AttachedNetworkBoard) XXX_Size() int {
	return xxx_messageInfo_AttachedNetworkBoard.Size(m)
}
func (m *AttachedNetworkBoard) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachedNetworkBoard.DiscardUnknown(m)
}

var xxx_messageInfo_AttachedNetworkBoard proto.InternalMessageInfo

func (m *AttachedNetworkBoard) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AttachedNetworkBoard) GetFqbn() string {
	if m != nil {
		return m.Fqbn
	}
	return ""
}

func (m *AttachedNetworkBoard) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func (m *AttachedNetworkBoard) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AttachedNetworkBoard) GetPort() uint64 {
	if m != nil {
		return m.Port
	}
	return 0
}

type AttachedSerialBoard struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Fqbn                 string   `protobuf:"bytes,2,opt,name=fqbn,proto3" json:"fqbn,omitempty"`
	Port                 string   `protobuf:"bytes,3,opt,name=port,proto3" json:"port,omitempty"`
	SerialNumber         string   `protobuf:"bytes,4,opt,name=serialNumber,proto3" json:"serialNumber,omitempty"`
	ProductID            string   `protobuf:"bytes,5,opt,name=productID,proto3" json:"productID,omitempty"`
	VendorID             string   `protobuf:"bytes,6,opt,name=vendorID,proto3" json:"vendorID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttachedSerialBoard) Reset()         { *m = AttachedSerialBoard{} }
func (m *AttachedSerialBoard) String() string { return proto.CompactTextString(m) }
func (*AttachedSerialBoard) ProtoMessage()    {}
func (*AttachedSerialBoard) Descriptor() ([]byte, []int) {
	return fileDescriptor_937f74b042f92c0f, []int{8}
}

func (m *AttachedSerialBoard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttachedSerialBoard.Unmarshal(m, b)
}
func (m *AttachedSerialBoard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttachedSerialBoard.Marshal(b, m, deterministic)
}
func (m *AttachedSerialBoard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttachedSerialBoard.Merge(m, src)
}
func (m *AttachedSerialBoard) XXX_Size() int {
	return xxx_messageInfo_AttachedSerialBoard.Size(m)
}
func (m *AttachedSerialBoard) XXX_DiscardUnknown() {
	xxx_messageInfo_AttachedSerialBoard.DiscardUnknown(m)
}

var xxx_messageInfo_AttachedSerialBoard proto.InternalMessageInfo

func (m *AttachedSerialBoard) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AttachedSerialBoard) GetFqbn() string {
	if m != nil {
		return m.Fqbn
	}
	return ""
}

func (m *AttachedSerialBoard) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *AttachedSerialBoard) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *AttachedSerialBoard) GetProductID() string {
	if m != nil {
		return m.ProductID
	}
	return ""
}

func (m *AttachedSerialBoard) GetVendorID() string {
	if m != nil {
		return m.VendorID
	}
	return ""
}

func init() {
	proto.RegisterType((*BoardDetailsReq)(nil), "arduino.BoardDetailsReq")
	proto.RegisterType((*BoardDetailsResp)(nil), "arduino.BoardDetailsResp")
	proto.RegisterType((*ConfigOption)(nil), "arduino.ConfigOption")
	proto.RegisterType((*ConfigValue)(nil), "arduino.ConfigValue")
	proto.RegisterType((*RequiredTool)(nil), "arduino.RequiredTool")
	proto.RegisterType((*BoardListReq)(nil), "arduino.BoardListReq")
	proto.RegisterType((*BoardListResp)(nil), "arduino.BoardListResp")
	proto.RegisterType((*AttachedNetworkBoard)(nil), "arduino.AttachedNetworkBoard")
	proto.RegisterType((*AttachedSerialBoard)(nil), "arduino.AttachedSerialBoard")
}

func init() { proto.RegisterFile("board.proto", fileDescriptor_937f74b042f92c0f) }

var fileDescriptor_937f74b042f92c0f = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0x96, 0xfb, 0xc8, 0x63, 0xec, 0xf0, 0x58, 0x02, 0xb2, 0xaa, 0x22, 0x82, 0xc5, 0x21, 0x07,
	0x9a, 0x4a, 0x05, 0x89, 0x0b, 0x1c, 0x28, 0xb9, 0x44, 0xaa, 0x8a, 0xb4, 0x54, 0x08, 0x71, 0x09,
	0x6b, 0x7b, 0x93, 0xae, 0xea, 0xec, 0x3a, 0xbb, 0xeb, 0xf4, 0x86, 0xf8, 0x2d, 0xfc, 0x01, 0xfe,
	0x22, 0xda, 0x87, 0x8d, 0x53, 0x72, 0xe9, 0xc9, 0xf3, 0xcd, 0x7c, 0xdf, 0xec, 0xcc, 0xb7, 0x2b,
	0x43, 0x98, 0x0a, 0x22, 0xf3, 0x49, 0x29, 0x85, 0x16, 0xa8, 0x4b, 0x64, 0x5e, 0x31, 0x2e, 0x8e,
	0xa2, 0x4c, 0xac, 0x56, 0x82, 0xbb, 0x74, 0x72, 0x05, 0x0f, 0xcf, 0x0d, 0x6b, 0x4a, 0x35, 0x61,
	0x85, 0xc2, 0x74, 0x8d, 0x4e, 0xa0, 0xc7, 0xb8, 0xd2, 0x84, 0x67, 0x34, 0x0e, 0x46, 0xc1, 0x38,
	0x3c, 0x7b, 0x3c, 0xf1, 0xe2, 0xc9, 0xcc, 0x17, 0x70, 0x43, 0x41, 0x08, 0x0e, 0x16, 0xeb, 0x94,
	0xc7, 0x7b, 0xa3, 0x60, 0xdc, 0xc7, 0x36, 0x4e, 0x7e, 0x07, 0xf0, 0x68, 0xbb, 0xad, 0x2a, 0x0d,
	0x91, 0x93, 0x15, 0xad, 0x89, 0x26, 0x46, 0xef, 0xe1, 0x41, 0x26, 0xf8, 0x82, 0x2d, 0xe7, 0xa2,
	0xd4, 0x4c, 0x70, 0x15, 0xef, 0x8f, 0xf6, 0xc7, 0xe1, 0xd9, 0xd3, 0xe6, 0xc4, 0x4f, 0xb6, 0xfc,
	0xd9, 0x56, 0xf1, 0x20, 0x6b, 0x21, 0x65, 0xd4, 0x92, 0xae, 0x2b, 0x26, 0x69, 0x3e, 0xd7, 0x42,
	0x14, 0x2a, 0x3e, 0xb8, 0xa3, 0xc6, 0xbe, 0x7c, 0x25, 0x44, 0x81, 0x07, 0xb2, 0x85, 0x54, 0x72,
	0x0b, 0x51, 0xbb, 0x39, 0x7a, 0x06, 0x1d, 0x37, 0x84, 0xdd, 0xba, 0x8f, 0x3d, 0x42, 0x2f, 0x21,
	0x72, 0xd1, 0xbc, 0x20, 0x29, 0x2d, 0xfc, 0xfc, 0xa1, 0xcb, 0x5d, 0x98, 0x14, 0x7a, 0x0d, 0x9d,
	0x0d, 0x29, 0x2a, 0x5a, 0x8f, 0x3f, 0xbc, 0x33, 0xfe, 0x57, 0x53, 0xc4, 0x9e, 0x93, 0xfc, 0x80,
	0xb0, 0x95, 0x46, 0x43, 0x38, 0xb4, 0x05, 0x7f, 0xac, 0x03, 0xe8, 0x05, 0x84, 0x36, 0xd8, 0x3a,
	0x14, 0x6c, 0xca, 0x9d, 0x79, 0x04, 0x3d, 0x45, 0x0b, 0x9a, 0x69, 0x9a, 0xc7, 0xfb, 0xa3, 0x60,
	0xdc, 0xc3, 0x0d, 0x4e, 0xbe, 0x41, 0xd4, 0xde, 0xbc, 0xb1, 0x3e, 0x68, 0x59, 0x1f, 0x43, 0x77,
	0x43, 0xa5, 0x32, 0xfb, 0xba, 0xe6, 0x35, 0x34, 0x9d, 0x4b, 0x92, 0xdd, 0x90, 0x25, 0x95, 0xb6,
	0x73, 0x1f, 0x37, 0x38, 0xf9, 0x00, 0x91, 0xbd, 0xd8, 0x0b, 0xa6, 0xf4, 0xfd, 0x1f, 0x4b, 0xf2,
	0x13, 0x06, 0x2d, 0xb9, 0x2a, 0xd1, 0x5b, 0xe8, 0x28, 0x2a, 0x19, 0x29, 0xe2, 0xc0, 0x3a, 0x77,
	0xdc, 0xa8, 0x3f, 0x6a, 0x4d, 0xb2, 0x6b, 0x9a, 0x7f, 0xb1, 0x65, 0xab, 0xc2, 0x9e, 0x8b, 0xde,
	0x41, 0x97, 0x53, 0x7d, 0x2b, 0xe4, 0x4d, 0xbc, 0x67, 0x65, 0xcf, 0xff, 0x93, 0x5d, 0xba, 0xba,
	0xd3, 0xd5, 0xec, 0xe4, 0x57, 0x00, 0xc3, 0x5d, 0x8c, 0x9d, 0x0e, 0xed, 0x78, 0xd9, 0x26, 0xc7,
	0xf8, 0x42, 0x78, 0x5f, 0x6c, 0x6c, 0x9c, 0x24, 0x79, 0x2e, 0xa9, 0x32, 0xef, 0xcf, 0x3a, 0xe9,
	0xa1, 0x61, 0x97, 0x42, 0xea, 0xf8, 0x70, 0x14, 0x8c, 0x0f, 0xb0, 0x8d, 0x93, 0x3f, 0x01, 0x3c,
	0xd9, 0xb1, 0xdb, 0x7d, 0x26, 0xb0, 0x3d, 0xfd, 0x04, 0x26, 0x46, 0x09, 0x44, 0xce, 0x99, 0xcb,
	0x6a, 0x95, 0x52, 0xe9, 0xc7, 0xd8, 0xca, 0xa1, 0x63, 0xe8, 0x97, 0x52, 0xe4, 0x55, 0xa6, 0x67,
	0x53, 0x3b, 0x50, 0x1f, 0xff, 0x4b, 0x98, 0x3b, 0xdf, 0x50, 0x9e, 0x0b, 0x39, 0x9b, 0xc6, 0x1d,
	0x77, 0xe7, 0x35, 0x3e, 0x7f, 0xf5, 0x3d, 0x59, 0x32, 0x7d, 0x5d, 0xa5, 0x93, 0x4c, 0xac, 0x4e,
	0xbd, 0xd1, 0xf5, 0xf7, 0x24, 0x2b, 0xd8, 0xa9, 0x2c, 0xb3, 0xb4, 0x63, 0x7f, 0x28, 0x6f, 0xfe,
	0x06, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x9e, 0x7a, 0xa4, 0x76, 0x04, 0x00, 0x00,
}
