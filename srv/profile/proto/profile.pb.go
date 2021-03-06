// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/profile.proto

package profile

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	HotelIds             []string `protobuf:"bytes,1,rep,name=hotelIds,proto3" json:"hotelIds,omitempty"`
	Locale               string   `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_16247a1ae8103e3c, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetHotelIds() []string {
	if m != nil {
		return m.HotelIds
	}
	return nil
}

func (m *Request) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

type Result struct {
	Hotels               []*Hotel `protobuf:"bytes,1,rep,name=hotels,proto3" json:"hotels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_16247a1ae8103e3c, []int{1}
}
func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (dst *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(dst, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetHotels() []*Hotel {
	if m != nil {
		return m.Hotels
	}
	return nil
}

type Hotel struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,3,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Address              *Address `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	Images               []*Image `protobuf:"bytes,6,rep,name=images,proto3" json:"images,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hotel) Reset()         { *m = Hotel{} }
func (m *Hotel) String() string { return proto.CompactTextString(m) }
func (*Hotel) ProtoMessage()    {}
func (*Hotel) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_16247a1ae8103e3c, []int{2}
}
func (m *Hotel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hotel.Unmarshal(m, b)
}
func (m *Hotel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hotel.Marshal(b, m, deterministic)
}
func (dst *Hotel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hotel.Merge(dst, src)
}
func (m *Hotel) XXX_Size() int {
	return xxx_messageInfo_Hotel.Size(m)
}
func (m *Hotel) XXX_DiscardUnknown() {
	xxx_messageInfo_Hotel.DiscardUnknown(m)
}

var xxx_messageInfo_Hotel proto.InternalMessageInfo

func (m *Hotel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Hotel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Hotel) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *Hotel) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Hotel) GetAddress() *Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Hotel) GetImages() []*Image {
	if m != nil {
		return m.Images
	}
	return nil
}

type Address struct {
	StreetNumber         string   `protobuf:"bytes,1,opt,name=streetNumber,proto3" json:"streetNumber,omitempty"`
	StreetName           string   `protobuf:"bytes,2,opt,name=streetName,proto3" json:"streetName,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	Country              string   `protobuf:"bytes,5,opt,name=country,proto3" json:"country,omitempty"`
	PostalCode           string   `protobuf:"bytes,6,opt,name=postalCode,proto3" json:"postalCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_16247a1ae8103e3c, []int{3}
}
func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (dst *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(dst, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetStreetNumber() string {
	if m != nil {
		return m.StreetNumber
	}
	return ""
}

func (m *Address) GetStreetName() string {
	if m != nil {
		return m.StreetName
	}
	return ""
}

func (m *Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Address) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Address) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

type Image struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Default              bool     `protobuf:"varint,2,opt,name=default,proto3" json:"default,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_profile_16247a1ae8103e3c, []int{4}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (dst *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(dst, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Image) GetDefault() bool {
	if m != nil {
		return m.Default
	}
	return false
}

func init() {
	proto.RegisterType((*Request)(nil), "profile.Request")
	proto.RegisterType((*Result)(nil), "profile.Result")
	proto.RegisterType((*Hotel)(nil), "profile.Hotel")
	proto.RegisterType((*Address)(nil), "profile.Address")
	proto.RegisterType((*Image)(nil), "profile.Image")
}

func init() { proto.RegisterFile("proto/profile.proto", fileDescriptor_profile_16247a1ae8103e3c) }

var fileDescriptor_profile_16247a1ae8103e3c = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x41, 0xae, 0xda, 0x30,
	0x10, 0x86, 0x15, 0x20, 0x09, 0x4c, 0x2a, 0x8a, 0xa6, 0x55, 0x65, 0xb1, 0xa8, 0xa2, 0x2c, 0x2a,
	0xd4, 0x05, 0x45, 0xb0, 0xac, 0xba, 0xa8, 0xba, 0x68, 0xd9, 0x54, 0x95, 0x6f, 0x10, 0xe2, 0xa1,
	0x44, 0x32, 0x71, 0x6a, 0x3b, 0x0b, 0x8e, 0xd5, 0x33, 0xbc, 0x8b, 0x3d, 0xd9, 0x71, 0x20, 0xef,
	0xed, 0xe6, 0xff, 0x66, 0xec, 0x99, 0x7f, 0x34, 0xf0, 0xae, 0xd5, 0xca, 0xaa, 0x2f, 0xad, 0x56,
	0xe7, 0x5a, 0xd2, 0xd6, 0x2b, 0x4c, 0x83, 0x2c, 0xbe, 0x41, 0xca, 0xe9, 0x5f, 0x47, 0xc6, 0xe2,
	0x1a, 0xe6, 0x17, 0x65, 0x49, 0x1e, 0x85, 0x61, 0x51, 0x3e, 0xdd, 0x2c, 0xf8, 0x5d, 0xe3, 0x07,
	0x48, 0xa4, 0xaa, 0x4a, 0x49, 0x6c, 0x92, 0x47, 0x9b, 0x05, 0x0f, 0xaa, 0xd8, 0x41, 0xc2, 0xc9,
	0x74, 0xd2, 0xe2, 0x27, 0x48, 0x7c, 0x75, 0xff, 0x36, 0xdb, 0x2f, 0xb7, 0x43, 0xc7, 0x5f, 0x0e,
	0xf3, 0x90, 0x2d, 0x9e, 0x22, 0x88, 0x3d, 0xc1, 0x25, 0x4c, 0x6a, 0xc1, 0x22, 0xff, 0xdf, 0xa4,
	0x16, 0x88, 0x30, 0x6b, 0xca, 0xeb, 0xd0, 0xc1, 0xc7, 0x98, 0x43, 0xd6, 0x5e, 0x54, 0x43, 0xbf,
	0xbb, 0xeb, 0x89, 0x34, 0x9b, 0xfa, 0xd4, 0x18, 0xb9, 0x0a, 0x41, 0xa6, 0xd2, 0x75, 0x6b, 0x6b,
	0xd5, 0xb0, 0x59, 0x5f, 0x31, 0x42, 0xf8, 0x19, 0xd2, 0x52, 0x08, 0x4d, 0xc6, 0xb0, 0x38, 0x8f,
	0x36, 0xd9, 0x7e, 0x75, 0x1f, 0xed, 0x7b, 0xcf, 0xf9, 0x50, 0xe0, 0x5c, 0xd4, 0xd7, 0xf2, 0x2f,
	0x19, 0x96, 0xbc, 0x72, 0x71, 0x74, 0x98, 0x87, 0x6c, 0xf1, 0x3f, 0x82, 0x34, 0x3c, 0xc6, 0x02,
	0xde, 0x18, 0xab, 0x89, 0x6c, 0x18, 0xb2, 0x77, 0xf4, 0x82, 0xe1, 0x47, 0x80, 0xa0, 0x1f, 0x0e,
	0x47, 0xc4, 0x79, 0xaf, 0x6a, 0x7b, 0x0b, 0x06, 0x7d, 0x8c, 0xef, 0x21, 0x36, 0xb6, 0xb4, 0x14,
	0x3c, 0xf5, 0x02, 0x19, 0xa4, 0x95, 0xea, 0x1a, 0xab, 0x6f, 0xde, 0xcd, 0x82, 0x0f, 0xd2, 0xf5,
	0x68, 0x95, 0xb1, 0xa5, 0xfc, 0xa1, 0x04, 0xb1, 0xa4, 0xef, 0xf1, 0x20, 0xc5, 0x01, 0x62, 0x6f,
	0x02, 0x57, 0x30, 0xed, 0xb4, 0x0c, 0x73, 0xba, 0xd0, 0x7d, 0x2a, 0xe8, 0x5c, 0x76, 0xd2, 0xfa,
	0xd9, 0xe6, 0x7c, 0x90, 0xfb, 0xaf, 0x90, 0xfe, 0xe9, 0x37, 0x80, 0x3b, 0xc8, 0x7e, 0x92, 0x0d,
	0xca, 0xe0, 0x63, 0x8b, 0xe1, 0x80, 0xd6, 0x6f, 0x47, 0xc4, 0xdd, 0xc4, 0x29, 0xf1, 0xc7, 0x76,
	0x78, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x45, 0x4c, 0xc8, 0x04, 0x83, 0x02, 0x00, 0x00,
}
