// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/slashing/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState - all slashing state that must be provided at genesis
type GenesisState struct {
	Params       Params                  `protobuf:"bytes,1,opt,name=params,proto3,casttype=Params" json:"params"`
	SigningInfos []ValidatorSigningInfos `protobuf:"bytes,2,rep,name=signing_infos,json=signingInfos,proto3,casttype=ValidatorSigningInfos" json:"signing_infos" yaml:"signing_infos"`
	MissedBlocks []ValidatorMissedBlocks `protobuf:"bytes,3,rep,name=missed_blocks,json=missedBlocks,proto3,casttype=ValidatorMissedBlocks" json:"missed_blocks" yaml:"missed_blocks"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_4742afabdd32b41b, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetSigningInfos() []ValidatorSigningInfos {
	if m != nil {
		return m.SigningInfos
	}
	return nil
}

func (m *GenesisState) GetMissedBlocks() []ValidatorMissedBlocks {
	if m != nil {
		return m.MissedBlocks
	}
	return nil
}

// ValidatorSigningInfos
type ValidatorSigningInfos struct {
	Address      string               `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	SigningInfos ValidatorSigningInfo `protobuf:"bytes,2,opt,name=signing_infos,json=signingInfos,proto3,casttype=ValidatorSigningInfo" json:"signing_infos" yaml:"signing_infos"`
}

func (m *ValidatorSigningInfos) Reset()         { *m = ValidatorSigningInfos{} }
func (m *ValidatorSigningInfos) String() string { return proto.CompactTextString(m) }
func (*ValidatorSigningInfos) ProtoMessage()    {}
func (*ValidatorSigningInfos) Descriptor() ([]byte, []int) {
	return fileDescriptor_4742afabdd32b41b, []int{1}
}
func (m *ValidatorSigningInfos) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSigningInfos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSigningInfos.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSigningInfos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSigningInfos.Merge(m, src)
}
func (m *ValidatorSigningInfos) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSigningInfos) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSigningInfos.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSigningInfos proto.InternalMessageInfo

func (m *ValidatorSigningInfos) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ValidatorSigningInfos) GetSigningInfos() ValidatorSigningInfo {
	if m != nil {
		return m.SigningInfos
	}
	return ValidatorSigningInfo{}
}

//ValidatorMissedBlocks
type ValidatorMissedBlocks struct {
	Address      string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	MissedBlocks []MissedBlock `protobuf:"bytes,2,rep,name=missed_blocks,json=missedBlocks,proto3,casttype=MissedBlock" json:"missed_blocks" yaml:"missed_blocks"`
}

func (m *ValidatorMissedBlocks) Reset()         { *m = ValidatorMissedBlocks{} }
func (m *ValidatorMissedBlocks) String() string { return proto.CompactTextString(m) }
func (*ValidatorMissedBlocks) ProtoMessage()    {}
func (*ValidatorMissedBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_4742afabdd32b41b, []int{2}
}
func (m *ValidatorMissedBlocks) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorMissedBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorMissedBlocks.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorMissedBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorMissedBlocks.Merge(m, src)
}
func (m *ValidatorMissedBlocks) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorMissedBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorMissedBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorMissedBlocks proto.InternalMessageInfo

func (m *ValidatorMissedBlocks) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ValidatorMissedBlocks) GetMissedBlocks() []MissedBlock {
	if m != nil {
		return m.MissedBlocks
	}
	return nil
}

// MissedBlock
type MissedBlock struct {
	Index  int64 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Missed bool  `protobuf:"varint,2,opt,name=missed,proto3" json:"missed,omitempty"`
}

func (m *MissedBlock) Reset()         { *m = MissedBlock{} }
func (m *MissedBlock) String() string { return proto.CompactTextString(m) }
func (*MissedBlock) ProtoMessage()    {}
func (*MissedBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_4742afabdd32b41b, []int{3}
}
func (m *MissedBlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MissedBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MissedBlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MissedBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MissedBlock.Merge(m, src)
}
func (m *MissedBlock) XXX_Size() int {
	return m.Size()
}
func (m *MissedBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_MissedBlock.DiscardUnknown(m)
}

var xxx_messageInfo_MissedBlock proto.InternalMessageInfo

func (m *MissedBlock) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *MissedBlock) GetMissed() bool {
	if m != nil {
		return m.Missed
	}
	return false
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "cosmos.slashing.GenesisState")
	proto.RegisterType((*ValidatorSigningInfos)(nil), "cosmos.slashing.ValidatorSigningInfos")
	proto.RegisterType((*ValidatorMissedBlocks)(nil), "cosmos.slashing.ValidatorMissedBlocks")
	proto.RegisterType((*MissedBlock)(nil), "cosmos.slashing.MissedBlock")
}

func init() { proto.RegisterFile("cosmos/slashing/genesis.proto", fileDescriptor_4742afabdd32b41b) }

var fileDescriptor_4742afabdd32b41b = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x2f, 0xce, 0x49, 0x2c, 0xce, 0xc8, 0xcc, 0x4b, 0xd7, 0x4f, 0x4f, 0xcd, 0x4b,
	0x2d, 0xce, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x87, 0x48, 0xeb, 0xc1, 0xa4,
	0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x72, 0xfa, 0x20, 0x16, 0x44, 0x99, 0x94, 0x1c, 0xba,
	0x29, 0x30, 0x06, 0x44, 0x5e, 0xe9, 0x26, 0x13, 0x17, 0x8f, 0x3b, 0xc4, 0xe0, 0xe0, 0x92, 0xc4,
	0x92, 0x54, 0x21, 0x7b, 0x2e, 0xb6, 0x82, 0xc4, 0xa2, 0xc4, 0xdc, 0x62, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0x6e, 0x23, 0x71, 0x3d, 0x34, 0x8b, 0xf4, 0x02, 0xc0, 0xd2, 0x4e, 0x7c, 0x27, 0xee, 0xc9,
	0x33, 0xfc, 0xba, 0x27, 0xcf, 0x06, 0xe1, 0x07, 0x41, 0xb5, 0x09, 0x35, 0x31, 0x72, 0xf1, 0x16,
	0x67, 0xa6, 0xe7, 0x65, 0xe6, 0xa5, 0xc7, 0x67, 0xe6, 0xa5, 0xe5, 0x17, 0x4b, 0x30, 0x29, 0x30,
	0x6b, 0x70, 0x1b, 0xa9, 0x61, 0x18, 0x14, 0x96, 0x98, 0x93, 0x99, 0x92, 0x58, 0x92, 0x5f, 0x14,
	0x0c, 0x51, 0xee, 0x09, 0x52, 0xed, 0x64, 0x0a, 0x32, 0xf7, 0xd3, 0x3d, 0x79, 0x91, 0xca, 0xc4,
	0xdc, 0x1c, 0x2b, 0x25, 0x14, 0xa3, 0x94, 0x7e, 0xdd, 0x93, 0x17, 0xc5, 0xaa, 0x2d, 0x88, 0xa7,
	0x18, 0x89, 0x07, 0x76, 0x44, 0x6e, 0x66, 0x71, 0x71, 0x6a, 0x4a, 0x7c, 0x52, 0x4e, 0x7e, 0x72,
	0x76, 0xb1, 0x04, 0x33, 0x21, 0x47, 0xf8, 0x82, 0x95, 0x3b, 0x81, 0x55, 0xa3, 0x3b, 0x02, 0xc5,
	0x28, 0x54, 0x47, 0x20, 0x6b, 0x0b, 0xe2, 0xc9, 0x45, 0xe2, 0x29, 0x6d, 0x61, 0xe4, 0xc2, 0xee,
	0x58, 0x21, 0x09, 0x2e, 0xf6, 0xc4, 0x94, 0x94, 0xa2, 0xd4, 0x62, 0x48, 0x28, 0x73, 0x06, 0xc1,
	0xb8, 0x42, 0x0d, 0x58, 0x42, 0x0f, 0x14, 0x0d, 0xaa, 0x44, 0x85, 0x9e, 0x93, 0x09, 0x81, 0xc0,
	0x13, 0xc1, 0xa6, 0x0b, 0x35, 0xec, 0x94, 0x16, 0x22, 0x3b, 0x1b, 0xd9, 0x7b, 0x78, 0x9c, 0x9d,
	0x87, 0x1e, 0xdc, 0x90, 0x38, 0x97, 0xc1, 0x70, 0x35, 0x92, 0x79, 0x4e, 0xda, 0x04, 0x02, 0x99,
	0x1b, 0x49, 0x31, 0x5a, 0xd0, 0x5a, 0x73, 0x21, 0x4b, 0x0a, 0x89, 0x70, 0xb1, 0x66, 0xe6, 0xa5,
	0xa4, 0x56, 0x80, 0x9d, 0xc5, 0x1c, 0x04, 0xe1, 0x08, 0x89, 0x71, 0xb1, 0x41, 0x34, 0x81, 0xc3,
	0x90, 0x23, 0x08, 0xca, 0x73, 0xf2, 0x5e, 0xf1, 0x48, 0x8e, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86,
	0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73,
	0xf5, 0xa1, 0x99, 0x07, 0x42, 0xe9, 0x16, 0xa7, 0x64, 0xeb, 0x57, 0x20, 0x72, 0x52, 0x49, 0x65,
	0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0x1f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xad, 0x7b,
	0x90, 0x31, 0xaf, 0x03, 0x00, 0x00,
}

func (this *GenesisState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GenesisState)
	if !ok {
		that2, ok := that.(GenesisState)
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
	if !this.Params.Equal(&that1.Params) {
		return false
	}
	if len(this.SigningInfos) != len(that1.SigningInfos) {
		return false
	}
	for i := range this.SigningInfos {
		if !this.SigningInfos[i].Equal(&that1.SigningInfos[i]) {
			return false
		}
	}
	if len(this.MissedBlocks) != len(that1.MissedBlocks) {
		return false
	}
	for i := range this.MissedBlocks {
		if !this.MissedBlocks[i].Equal(&that1.MissedBlocks[i]) {
			return false
		}
	}
	return true
}
func (this *ValidatorSigningInfos) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorSigningInfos)
	if !ok {
		that2, ok := that.(ValidatorSigningInfos)
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
	if this.Address != that1.Address {
		return false
	}
	if !this.SigningInfos.Equal(&that1.SigningInfos) {
		return false
	}
	return true
}
func (this *ValidatorMissedBlocks) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorMissedBlocks)
	if !ok {
		that2, ok := that.(ValidatorMissedBlocks)
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
	if this.Address != that1.Address {
		return false
	}
	if len(this.MissedBlocks) != len(that1.MissedBlocks) {
		return false
	}
	for i := range this.MissedBlocks {
		if !this.MissedBlocks[i].Equal(&that1.MissedBlocks[i]) {
			return false
		}
	}
	return true
}
func (this *MissedBlock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MissedBlock)
	if !ok {
		that2, ok := that.(MissedBlock)
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
	if this.Index != that1.Index {
		return false
	}
	if this.Missed != that1.Missed {
		return false
	}
	return true
}
func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MissedBlocks) > 0 {
		for iNdEx := len(m.MissedBlocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MissedBlocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SigningInfos) > 0 {
		for iNdEx := len(m.SigningInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SigningInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ValidatorSigningInfos) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSigningInfos) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSigningInfos) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SigningInfos.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorMissedBlocks) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorMissedBlocks) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorMissedBlocks) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MissedBlocks) > 0 {
		for iNdEx := len(m.MissedBlocks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MissedBlocks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MissedBlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MissedBlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MissedBlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Missed {
		i--
		if m.Missed {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Index != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.SigningInfos) > 0 {
		for _, e := range m.SigningInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.MissedBlocks) > 0 {
		for _, e := range m.MissedBlocks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *ValidatorSigningInfos) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.SigningInfos.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *ValidatorMissedBlocks) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.MissedBlocks) > 0 {
		for _, e := range m.MissedBlocks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *MissedBlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovGenesis(uint64(m.Index))
	}
	if m.Missed {
		n += 2
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigningInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SigningInfos = append(m.SigningInfos, ValidatorSigningInfos{})
			if err := m.SigningInfos[len(m.SigningInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedBlocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MissedBlocks = append(m.MissedBlocks, ValidatorMissedBlocks{})
			if err := m.MissedBlocks[len(m.MissedBlocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ValidatorSigningInfos) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ValidatorSigningInfos: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSigningInfos: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigningInfos", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SigningInfos.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *ValidatorMissedBlocks) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: ValidatorMissedBlocks: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorMissedBlocks: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedBlocks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MissedBlocks = append(m.MissedBlocks, MissedBlock{})
			if err := m.MissedBlocks[len(m.MissedBlocks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *MissedBlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: MissedBlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MissedBlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Missed", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Missed = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenesis
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
