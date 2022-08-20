// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: parrots/models.proto

package types

import (
	fmt "fmt"
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

type Profile struct {
	Id             uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator        string   `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Username       string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	DisplayName    string   `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Description    string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	RespectedBeaks []uint64 `protobuf:"varint,6,rep,packed,name=respectedBeaks,proto3" json:"respectedBeaks,omitempty"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f4522fb09def05d, []int{0}
}
func (m *Profile) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return m.Size()
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Profile) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Profile) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Profile) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Profile) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Profile) GetRespectedBeaks() []uint64 {
	if m != nil {
		return m.RespectedBeaks
	}
	return nil
}

type Beak struct {
	Id              uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creator         string   `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	FileIndex       string   `protobuf:"bytes,3,opt,name=file_index,json=fileIndex,proto3" json:"file_index,omitempty"`
	Name            string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	CreatorUsername string   `protobuf:"bytes,5,opt,name=creator_username,json=creatorUsername,proto3" json:"creator_username,omitempty"`
	Description     string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	License         string   `protobuf:"bytes,7,opt,name=license,proto3" json:"license,omitempty"`
	RespectCount    uint64   `protobuf:"varint,8,opt,name=respect_count,json=respectCount,proto3" json:"respect_count,omitempty"`
	LinkedBeaks     []uint64 `protobuf:"varint,9,rep,packed,name=linked_beaks,json=linkedBeaks,proto3" json:"linked_beaks,omitempty"`
	Tags            []string `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (m *Beak) Reset()         { *m = Beak{} }
func (m *Beak) String() string { return proto.CompactTextString(m) }
func (*Beak) ProtoMessage()    {}
func (*Beak) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f4522fb09def05d, []int{1}
}
func (m *Beak) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Beak) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Beak.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Beak) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Beak.Merge(m, src)
}
func (m *Beak) XXX_Size() int {
	return m.Size()
}
func (m *Beak) XXX_DiscardUnknown() {
	xxx_messageInfo_Beak.DiscardUnknown(m)
}

var xxx_messageInfo_Beak proto.InternalMessageInfo

func (m *Beak) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Beak) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Beak) GetFileIndex() string {
	if m != nil {
		return m.FileIndex
	}
	return ""
}

func (m *Beak) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Beak) GetCreatorUsername() string {
	if m != nil {
		return m.CreatorUsername
	}
	return ""
}

func (m *Beak) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Beak) GetLicense() string {
	if m != nil {
		return m.License
	}
	return ""
}

func (m *Beak) GetRespectCount() uint64 {
	if m != nil {
		return m.RespectCount
	}
	return 0
}

func (m *Beak) GetLinkedBeaks() []uint64 {
	if m != nil {
		return m.LinkedBeaks
	}
	return nil
}

func (m *Beak) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*Profile)(nil), "parrots.parrots.Profile")
	proto.RegisterType((*Beak)(nil), "parrots.parrots.Beak")
}

func init() { proto.RegisterFile("parrots/models.proto", fileDescriptor_2f4522fb09def05d) }

var fileDescriptor_2f4522fb09def05d = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0x9b, 0x34, 0x5f, 0xd3, 0x4e, 0xfb, 0xb5, 0x1f, 0xcb, 0x07, 0x2e, 0x82, 0x21, 0x56,
	0x90, 0x7a, 0x69, 0x11, 0xdf, 0xa0, 0x9e, 0xbc, 0x88, 0x14, 0xbc, 0x78, 0x09, 0x69, 0x76, 0x94,
	0xa5, 0x69, 0x36, 0xec, 0x6e, 0xa1, 0x7d, 0x0b, 0x9f, 0xc2, 0xa7, 0xf0, 0x01, 0x3c, 0xf6, 0xe8,
	0x51, 0xda, 0x17, 0x91, 0xdd, 0x6e, 0xaa, 0xd4, 0x93, 0xa7, 0xcc, 0xfc, 0x66, 0x03, 0xff, 0xdf,
	0x30, 0xf0, 0xbf, 0x4c, 0xa5, 0x14, 0x5a, 0x8d, 0xe6, 0x82, 0x61, 0xae, 0x86, 0xa5, 0x14, 0x5a,
	0x90, 0x9e, 0xa3, 0x43, 0xf7, 0xed, 0xbf, 0x7a, 0x10, 0xde, 0x49, 0xf1, 0xc8, 0x73, 0x24, 0x5d,
	0xf0, 0x39, 0xa3, 0x5e, 0xec, 0x0d, 0x82, 0x89, 0xcf, 0x19, 0xa1, 0x10, 0x66, 0x12, 0x53, 0x2d,
	0x24, 0xf5, 0x63, 0x6f, 0xd0, 0x9a, 0x54, 0x2d, 0x39, 0x86, 0xe6, 0x42, 0xa1, 0x2c, 0xd2, 0x39,
	0xd2, 0xba, 0x1d, 0xed, 0x7b, 0x72, 0x0a, 0x1d, 0xc6, 0x55, 0x99, 0xa7, 0xab, 0xc4, 0xce, 0x03,
	0x3b, 0x6f, 0x3b, 0x76, 0x6b, 0x9e, 0xc4, 0xd0, 0x66, 0xa8, 0x32, 0xc9, 0x4b, 0xcd, 0x45, 0x41,
	0xff, 0xb8, 0x17, 0x5f, 0x88, 0x9c, 0x43, 0x57, 0xa2, 0x2a, 0x31, 0xd3, 0xc8, 0xc6, 0x98, 0xce,
	0x14, 0x6d, 0xc4, 0xf5, 0x41, 0x30, 0x39, 0xa0, 0xfd, 0x17, 0x1f, 0x02, 0x53, 0xfd, 0x22, 0xfb,
	0x09, 0x80, 0xb1, 0x4d, 0x78, 0xc1, 0x70, 0xe9, 0xd2, 0xb7, 0x0c, 0xb9, 0x31, 0x80, 0x10, 0x08,
	0xbe, 0xc5, 0xb6, 0x35, 0xb9, 0x80, 0x7f, 0xee, 0xef, 0x64, 0xaf, 0xbd, 0x0b, 0xdd, 0x73, 0xfc,
	0xbe, 0xb2, 0x3f, 0x50, 0x6b, 0xfc, 0x54, 0xa3, 0x10, 0xe6, 0x3c, 0xc3, 0x42, 0x21, 0x0d, 0x77,
	0xc9, 0x5c, 0x4b, 0xce, 0xe0, 0xaf, 0xd3, 0x4b, 0x32, 0xb1, 0x28, 0x34, 0x6d, 0x5a, 0x9d, 0x8e,
	0x83, 0xd7, 0x86, 0x99, 0xf5, 0xe6, 0xbc, 0x98, 0x21, 0x4b, 0xa6, 0x76, 0x2f, 0x2d, 0xbb, 0x97,
	0xf6, 0x8e, 0xd9, 0xa5, 0x18, 0x05, 0x9d, 0x3e, 0x29, 0x0a, 0x71, 0xdd, 0x28, 0x98, 0x7a, 0x7c,
	0xf9, 0xb6, 0x89, 0xbc, 0xf5, 0x26, 0xf2, 0x3e, 0x36, 0x91, 0xf7, 0xbc, 0x8d, 0x6a, 0xeb, 0x6d,
	0x54, 0x7b, 0xdf, 0x46, 0xb5, 0x87, 0xa3, 0xea, 0x50, 0x96, 0xa3, 0xaa, 0xd2, 0xab, 0x12, 0xd5,
	0xb4, 0x61, 0x4f, 0xe6, 0xea, 0x33, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x90, 0x8a, 0xab, 0x4a, 0x02,
	0x00, 0x00,
}

func (m *Profile) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Profile) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Profile) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RespectedBeaks) > 0 {
		dAtA2 := make([]byte, len(m.RespectedBeaks)*10)
		var j1 int
		for _, num := range m.RespectedBeaks {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintModels(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DisplayName) > 0 {
		i -= len(m.DisplayName)
		copy(dAtA[i:], m.DisplayName)
		i = encodeVarintModels(dAtA, i, uint64(len(m.DisplayName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintModels(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Beak) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Beak) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Beak) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tags) > 0 {
		for iNdEx := len(m.Tags) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tags[iNdEx])
			copy(dAtA[i:], m.Tags[iNdEx])
			i = encodeVarintModels(dAtA, i, uint64(len(m.Tags[iNdEx])))
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.LinkedBeaks) > 0 {
		dAtA4 := make([]byte, len(m.LinkedBeaks)*10)
		var j3 int
		for _, num := range m.LinkedBeaks {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintModels(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0x4a
	}
	if m.RespectCount != 0 {
		i = encodeVarintModels(dAtA, i, uint64(m.RespectCount))
		i--
		dAtA[i] = 0x40
	}
	if len(m.License) > 0 {
		i -= len(m.License)
		copy(dAtA[i:], m.License)
		i = encodeVarintModels(dAtA, i, uint64(len(m.License)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.CreatorUsername) > 0 {
		i -= len(m.CreatorUsername)
		copy(dAtA[i:], m.CreatorUsername)
		i = encodeVarintModels(dAtA, i, uint64(len(m.CreatorUsername)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.FileIndex) > 0 {
		i -= len(m.FileIndex)
		copy(dAtA[i:], m.FileIndex)
		i = encodeVarintModels(dAtA, i, uint64(len(m.FileIndex)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintModels(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintModels(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintModels(dAtA []byte, offset int, v uint64) int {
	offset -= sovModels(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Profile) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovModels(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	if len(m.RespectedBeaks) > 0 {
		l = 0
		for _, e := range m.RespectedBeaks {
			l += sovModels(uint64(e))
		}
		n += 1 + sovModels(uint64(l)) + l
	}
	return n
}

func (m *Beak) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovModels(uint64(m.Id))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.FileIndex)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.CreatorUsername)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	l = len(m.License)
	if l > 0 {
		n += 1 + l + sovModels(uint64(l))
	}
	if m.RespectCount != 0 {
		n += 1 + sovModels(uint64(m.RespectCount))
	}
	if len(m.LinkedBeaks) > 0 {
		l = 0
		for _, e := range m.LinkedBeaks {
			l += sovModels(uint64(e))
		}
		n += 1 + sovModels(uint64(l)) + l
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovModels(uint64(l))
		}
	}
	return n
}

func sovModels(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozModels(x uint64) (n int) {
	return sovModels(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Profile) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModels
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
			return fmt.Errorf("proto: Profile: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Profile: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowModels
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.RespectedBeaks = append(m.RespectedBeaks, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowModels
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthModels
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthModels
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.RespectedBeaks) == 0 {
					m.RespectedBeaks = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowModels
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.RespectedBeaks = append(m.RespectedBeaks, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field RespectedBeaks", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModels
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
func (m *Beak) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowModels
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
			return fmt.Errorf("proto: Beak: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Beak: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatorUsername", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreatorUsername = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field License", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.License = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RespectCount", wireType)
			}
			m.RespectCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RespectCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowModels
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.LinkedBeaks = append(m.LinkedBeaks, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowModels
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthModels
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthModels
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.LinkedBeaks) == 0 {
					m.LinkedBeaks = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowModels
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.LinkedBeaks = append(m.LinkedBeaks, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field LinkedBeaks", wireType)
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowModels
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
				return ErrInvalidLengthModels
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthModels
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipModels(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthModels
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
func skipModels(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowModels
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
					return 0, ErrIntOverflowModels
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
					return 0, ErrIntOverflowModels
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
				return 0, ErrInvalidLengthModels
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupModels
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthModels
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthModels        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowModels          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupModels = fmt.Errorf("proto: unexpected end of group")
)
