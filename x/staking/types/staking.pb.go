// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: staking.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgClaimValidator struct {
	Moniker    string                                        `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Website    string                                        `protobuf:"bytes,2,opt,name=website,proto3" json:"website,omitempty"`
	Social     string                                        `protobuf:"bytes,3,opt,name=social,proto3" json:"social,omitempty"`
	Identity   string                                        `protobuf:"bytes,4,opt,name=identity,proto3" json:"identity,omitempty"`
	Commission github_com_cosmos_cosmos_sdk_types.Dec        `protobuf:"bytes,5,opt,name=commission,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"commission" yaml:"commission"`
	ValKey     github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,6,opt,name=val_key,json=valKey,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"val_key,omitempty" yaml:"val_key"`
	PubKey     string                                        `protobuf:"bytes,7,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
}

func (m *MsgClaimValidator) Reset()         { *m = MsgClaimValidator{} }
func (m *MsgClaimValidator) String() string { return proto.CompactTextString(m) }
func (*MsgClaimValidator) ProtoMessage()    {}
func (*MsgClaimValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{0}
}
func (m *MsgClaimValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgClaimValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgClaimValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgClaimValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgClaimValidator.Merge(m, src)
}
func (m *MsgClaimValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgClaimValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgClaimValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgClaimValidator proto.InternalMessageInfo

func (m *MsgClaimValidator) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *MsgClaimValidator) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *MsgClaimValidator) GetSocial() string {
	if m != nil {
		return m.Social
	}
	return ""
}

func (m *MsgClaimValidator) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *MsgClaimValidator) GetValKey() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.ValKey
	}
	return nil
}

func (m *MsgClaimValidator) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

type Validator struct {
	Moniker    string                                        `protobuf:"bytes,1,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Website    string                                        `protobuf:"bytes,2,opt,name=website,proto3" json:"website,omitempty"`
	Social     string                                        `protobuf:"bytes,3,opt,name=social,proto3" json:"social,omitempty"`
	Identity   string                                        `protobuf:"bytes,4,opt,name=identity,proto3" json:"identity,omitempty"`
	Commission github_com_cosmos_cosmos_sdk_types.Dec        `protobuf:"bytes,5,opt,name=commission,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"commission" yaml:"commission"`
	ValKey     github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,6,opt,name=val_key,json=valKey,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"val_key,omitempty" yaml:"val_key"`
	PubKey     string                                        `protobuf:"bytes,7,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty" yaml:"consensus_pubkey"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_289e7c8aea278311, []int{1}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *Validator) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *Validator) GetSocial() string {
	if m != nil {
		return m.Social
	}
	return ""
}

func (m *Validator) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *Validator) GetValKey() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.ValKey
	}
	return nil
}

func (m *Validator) GetPubKey() string {
	if m != nil {
		return m.PubKey
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgClaimValidator)(nil), "kira.staking.MsgClaimValidator")
	proto.RegisterType((*Validator)(nil), "kira.staking.Validator")
}

func init() { proto.RegisterFile("staking.proto", fileDescriptor_289e7c8aea278311) }

var fileDescriptor_289e7c8aea278311 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x93, 0x3f, 0x6b, 0xdb, 0x40,
	0x18, 0xc6, 0x25, 0xd7, 0x95, 0xea, 0xc3, 0x2d, 0x58, 0x94, 0x5a, 0xb8, 0x20, 0x19, 0x0d, 0xc5,
	0x1d, 0x2c, 0x0d, 0xed, 0xe4, 0xad, 0x52, 0x37, 0xd3, 0xc5, 0x83, 0x87, 0x52, 0x30, 0x27, 0xe9,
	0x50, 0x0f, 0xfd, 0x39, 0xa1, 0xf7, 0xe4, 0x56, 0xdf, 0x22, 0x1f, 0x21, 0x1f, 0xc7, 0xa3, 0xc7,
	0x10, 0x88, 0x08, 0xf6, 0x92, 0x21, 0x93, 0xc7, 0x4c, 0x41, 0x7f, 0x9c, 0x98, 0x4c, 0x21, 0x73,
	0xa6, 0xbb, 0xe7, 0x7e, 0x77, 0xef, 0xfb, 0xdc, 0x03, 0x2f, 0x7a, 0x0f, 0x1c, 0x87, 0x34, 0x09,
	0xcc, 0x34, 0x63, 0x9c, 0x29, 0xfd, 0x90, 0x66, 0xd8, 0x6c, 0xcf, 0x46, 0x1f, 0x03, 0x16, 0xb0,
	0x1a, 0x58, 0xd5, 0xae, 0xb9, 0x63, 0x5c, 0x75, 0xd0, 0xe0, 0x17, 0x04, 0x4e, 0x84, 0x69, 0xbc,
	0xc4, 0x11, 0xf5, 0x31, 0x67, 0x99, 0xa2, 0x22, 0x39, 0x66, 0x09, 0x0d, 0x49, 0xa6, 0x8a, 0x63,
	0x71, 0xd2, 0x5b, 0x1c, 0x65, 0x45, 0xfe, 0x11, 0x17, 0x28, 0x27, 0x6a, 0xa7, 0x21, 0xad, 0x54,
	0x3e, 0x21, 0x09, 0x98, 0x47, 0x71, 0xa4, 0xbe, 0xa9, 0x41, 0xab, 0x94, 0x11, 0x7a, 0x47, 0x7d,
	0x92, 0x70, 0xca, 0x0b, 0xb5, 0x5b, 0x93, 0x07, 0xad, 0x78, 0x08, 0x79, 0x2c, 0x8e, 0x29, 0x00,
	0x65, 0x89, 0xfa, 0xb6, 0xa2, 0xb6, 0xb3, 0x29, 0x75, 0xe1, 0xb2, 0xd4, 0xbf, 0x04, 0x94, 0xff,
	0xcd, 0x5d, 0xd3, 0x63, 0xb1, 0xe5, 0x31, 0x88, 0x19, 0xb4, 0xcb, 0x14, 0xfc, 0xd0, 0xe2, 0x45,
	0x4a, 0xc0, 0xfc, 0x49, 0xbc, 0x43, 0xa9, 0x0f, 0x0a, 0x1c, 0x47, 0x33, 0xe3, 0xb1, 0x92, 0xb1,
	0x38, 0x29, 0xab, 0xfc, 0x41, 0xf2, 0x1a, 0x47, 0xab, 0x90, 0x14, 0xaa, 0x34, 0x16, 0x27, 0x7d,
	0xdb, 0x39, 0x94, 0xfa, 0x87, 0xe6, 0x4d, 0x0b, 0x8c, 0xbb, 0x52, 0x9f, 0x3e, 0xa3, 0xdf, 0x12,
	0x47, 0x3f, 0x7c, 0x3f, 0x23, 0x00, 0x0b, 0x69, 0x8d, 0xa3, 0x39, 0x29, 0x94, 0x21, 0x92, 0xd3,
	0xdc, 0xad, 0xab, 0xcb, 0xcd, 0xbf, 0xd3, 0xdc, 0x9d, 0x93, 0x62, 0xd6, 0xbd, 0x39, 0xd7, 0x45,
	0xe3, 0xb6, 0x83, 0x7a, 0xaf, 0xb9, 0xbe, 0x20, 0xd7, 0xef, 0x4f, 0x72, 0xb5, 0x3f, 0x1f, 0x4a,
	0x7d, 0x78, 0x74, 0x94, 0x00, 0x49, 0x20, 0x87, 0x55, 0x9a, 0xbb, 0x55, 0x9b, 0x63, 0xe8, 0xb6,
	0xb3, 0xd9, 0x69, 0xe2, 0x76, 0xa7, 0x89, 0xd7, 0x3b, 0x4d, 0x3c, 0xdb, 0x6b, 0xc2, 0x76, 0xaf,
	0x09, 0x17, 0x7b, 0x4d, 0xf8, 0xfd, 0xf5, 0xc4, 0xc6, 0x9c, 0x66, 0xd8, 0x61, 0x19, 0xb1, 0x80,
	0x84, 0x98, 0x5a, 0xff, 0xad, 0x76, 0x46, 0x1a, 0x37, 0xae, 0x54, 0x8f, 0xc6, 0xb7, 0xfb, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x6a, 0x50, 0xe6, 0x3a, 0x4f, 0x03, 0x00, 0x00,
}

func (this *MsgClaimValidator) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgClaimValidator)
	if !ok {
		that2, ok := that.(MsgClaimValidator)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Moniker != that1.Moniker {
		return false
	}
	if this.Website != that1.Website {
		return false
	}
	if this.Social != that1.Social {
		return false
	}
	if this.Identity != that1.Identity {
		return false
	}
	if !this.Commission.Equal(that1.Commission) {
		return false
	}
	if !bytes.Equal(this.ValKey, that1.ValKey) {
		return false
	}
	if this.PubKey != that1.PubKey {
		return false
	}
	return true
}
func (m *MsgClaimValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgClaimValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgClaimValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ValKey) > 0 {
		i -= len(m.ValKey)
		copy(dAtA[i:], m.ValKey)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.ValKey)))
		i--
		dAtA[i] = 0x32
	}
	{
		size := m.Commission.Size()
		i -= size
		if _, err := m.Commission.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Social) > 0 {
		i -= len(m.Social)
		copy(dAtA[i:], m.Social)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Social)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ValKey) > 0 {
		i -= len(m.ValKey)
		copy(dAtA[i:], m.ValKey)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.ValKey)))
		i--
		dAtA[i] = 0x32
	}
	{
		size := m.Commission.Size()
		i -= size
		if _, err := m.Commission.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintStaking(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Social) > 0 {
		i -= len(m.Social)
		copy(dAtA[i:], m.Social)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Social)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintStaking(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStaking(dAtA []byte, offset int, v uint64) int {
	offset -= sovStaking(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgClaimValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.Social)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = m.Commission.Size()
	n += 1 + l + sovStaking(uint64(l))
	l = len(m.ValKey)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	return n
}

func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.Social)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = m.Commission.Size()
	n += 1 + l + sovStaking(uint64(l))
	l = len(m.ValKey)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovStaking(uint64(l))
	}
	return n
}

func sovStaking(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStaking(x uint64) (n int) {
	return sovStaking(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgClaimValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgClaimValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgClaimValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Social", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Social = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Commission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValKey = append(m.ValKey[:0], dAtA[iNdEx:postIndex]...)
			if m.ValKey == nil {
				m.ValKey = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Social", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Social = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commission", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Commission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValKey = append(m.ValKey[:0], dAtA[iNdEx:postIndex]...)
			if m.ValKey == nil {
				m.ValKey = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStaking
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStaking
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStaking(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStaking
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStaking(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStaking
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStaking
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthStaking
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStaking
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStaking
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStaking        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStaking          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStaking = fmt.Errorf("proto: unexpected end of group")
)
