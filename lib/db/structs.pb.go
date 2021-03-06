// Code generated by protoc-gen-gogo.
// source: structs.proto
// DO NOT EDIT!

/*
	Package db is a generated protocol buffer package.

	It is generated from these files:
		structs.proto

	It has these top-level messages:
		FileVersion
		VersionList
		FileInfoTruncated
*/
package db

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import protocol "github.com/syncthing/syncthing/lib/protocol"

import github_com_syncthing_syncthing_lib_protocol "github.com/syncthing/syncthing/lib/protocol"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type FileVersion struct {
	Version protocol.Vector `protobuf:"bytes,1,opt,name=version" json:"version"`
	Device  []byte          `protobuf:"bytes,2,opt,name=device,proto3" json:"device,omitempty"`
}

func (m *FileVersion) Reset()                    { *m = FileVersion{} }
func (m *FileVersion) String() string            { return proto.CompactTextString(m) }
func (*FileVersion) ProtoMessage()               {}
func (*FileVersion) Descriptor() ([]byte, []int) { return fileDescriptorStructs, []int{0} }

type VersionList struct {
	Versions []FileVersion `protobuf:"bytes,1,rep,name=versions" json:"versions"`
}

func (m *VersionList) Reset()                    { *m = VersionList{} }
func (*VersionList) ProtoMessage()               {}
func (*VersionList) Descriptor() ([]byte, []int) { return fileDescriptorStructs, []int{1} }

// Must be the same as FileInfo but without the blocks field
type FileInfoTruncated struct {
	Name          string                                              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type          protocol.FileInfoType                               `protobuf:"varint,2,opt,name=type,proto3,enum=protocol.FileInfoType" json:"type,omitempty"`
	Size          int64                                               `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	Permissions   uint32                                              `protobuf:"varint,4,opt,name=permissions,proto3" json:"permissions,omitempty"`
	ModifiedS     int64                                               `protobuf:"varint,5,opt,name=modified_s,json=modifiedS,proto3" json:"modified_s,omitempty"`
	ModifiedNs    int32                                               `protobuf:"varint,11,opt,name=modified_ns,json=modifiedNs,proto3" json:"modified_ns,omitempty"`
	ModifiedBy    github_com_syncthing_syncthing_lib_protocol.ShortID `protobuf:"varint,12,opt,name=modified_by,json=modifiedBy,proto3,customtype=github.com/syncthing/syncthing/lib/protocol.ShortID" json:"modified_by"`
	Deleted       bool                                                `protobuf:"varint,6,opt,name=deleted,proto3" json:"deleted,omitempty"`
	Invalid       bool                                                `protobuf:"varint,7,opt,name=invalid,proto3" json:"invalid,omitempty"`
	NoPermissions bool                                                `protobuf:"varint,8,opt,name=no_permissions,json=noPermissions,proto3" json:"no_permissions,omitempty"`
	Version       protocol.Vector                                     `protobuf:"bytes,9,opt,name=version" json:"version"`
	Sequence      int64                                               `protobuf:"varint,10,opt,name=sequence,proto3" json:"sequence,omitempty"`
	SymlinkTarget string                                              `protobuf:"bytes,17,opt,name=symlink_target,json=symlinkTarget,proto3" json:"symlink_target,omitempty"`
}

func (m *FileInfoTruncated) Reset()                    { *m = FileInfoTruncated{} }
func (*FileInfoTruncated) ProtoMessage()               {}
func (*FileInfoTruncated) Descriptor() ([]byte, []int) { return fileDescriptorStructs, []int{2} }

func init() {
	proto.RegisterType((*FileVersion)(nil), "db.FileVersion")
	proto.RegisterType((*VersionList)(nil), "db.VersionList")
	proto.RegisterType((*FileInfoTruncated)(nil), "db.FileInfoTruncated")
}
func (m *FileVersion) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileVersion) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintStructs(dAtA, i, uint64(m.Version.ProtoSize()))
	n1, err := m.Version.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.Device) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintStructs(dAtA, i, uint64(len(m.Device)))
		i += copy(dAtA[i:], m.Device)
	}
	return i, nil
}

func (m *VersionList) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Versions) > 0 {
		for _, msg := range m.Versions {
			dAtA[i] = 0xa
			i++
			i = encodeVarintStructs(dAtA, i, uint64(msg.ProtoSize()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *FileInfoTruncated) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileInfoTruncated) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStructs(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.Type != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintStructs(dAtA, i, uint64(m.Type))
	}
	if m.Size != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintStructs(dAtA, i, uint64(m.Size))
	}
	if m.Permissions != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintStructs(dAtA, i, uint64(m.Permissions))
	}
	if m.ModifiedS != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintStructs(dAtA, i, uint64(m.ModifiedS))
	}
	if m.Deleted {
		dAtA[i] = 0x30
		i++
		if m.Deleted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Invalid {
		dAtA[i] = 0x38
		i++
		if m.Invalid {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.NoPermissions {
		dAtA[i] = 0x40
		i++
		if m.NoPermissions {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	dAtA[i] = 0x4a
	i++
	i = encodeVarintStructs(dAtA, i, uint64(m.Version.ProtoSize()))
	n2, err := m.Version.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.Sequence != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintStructs(dAtA, i, uint64(m.Sequence))
	}
	if m.ModifiedNs != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintStructs(dAtA, i, uint64(m.ModifiedNs))
	}
	if m.ModifiedBy != 0 {
		dAtA[i] = 0x60
		i++
		i = encodeVarintStructs(dAtA, i, uint64(m.ModifiedBy))
	}
	if len(m.SymlinkTarget) > 0 {
		dAtA[i] = 0x8a
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintStructs(dAtA, i, uint64(len(m.SymlinkTarget)))
		i += copy(dAtA[i:], m.SymlinkTarget)
	}
	return i, nil
}

func encodeFixed64Structs(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Structs(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintStructs(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FileVersion) ProtoSize() (n int) {
	var l int
	_ = l
	l = m.Version.ProtoSize()
	n += 1 + l + sovStructs(uint64(l))
	l = len(m.Device)
	if l > 0 {
		n += 1 + l + sovStructs(uint64(l))
	}
	return n
}

func (m *VersionList) ProtoSize() (n int) {
	var l int
	_ = l
	if len(m.Versions) > 0 {
		for _, e := range m.Versions {
			l = e.ProtoSize()
			n += 1 + l + sovStructs(uint64(l))
		}
	}
	return n
}

func (m *FileInfoTruncated) ProtoSize() (n int) {
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovStructs(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovStructs(uint64(m.Type))
	}
	if m.Size != 0 {
		n += 1 + sovStructs(uint64(m.Size))
	}
	if m.Permissions != 0 {
		n += 1 + sovStructs(uint64(m.Permissions))
	}
	if m.ModifiedS != 0 {
		n += 1 + sovStructs(uint64(m.ModifiedS))
	}
	if m.Deleted {
		n += 2
	}
	if m.Invalid {
		n += 2
	}
	if m.NoPermissions {
		n += 2
	}
	l = m.Version.ProtoSize()
	n += 1 + l + sovStructs(uint64(l))
	if m.Sequence != 0 {
		n += 1 + sovStructs(uint64(m.Sequence))
	}
	if m.ModifiedNs != 0 {
		n += 1 + sovStructs(uint64(m.ModifiedNs))
	}
	if m.ModifiedBy != 0 {
		n += 1 + sovStructs(uint64(m.ModifiedBy))
	}
	l = len(m.SymlinkTarget)
	if l > 0 {
		n += 2 + l + sovStructs(uint64(l))
	}
	return n
}

func sovStructs(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStructs(x uint64) (n int) {
	return sovStructs(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FileVersion) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStructs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FileVersion: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileVersion: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStructs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Version.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Device", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthStructs
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Device = append(m.Device[:0], dAtA[iNdEx:postIndex]...)
			if m.Device == nil {
				m.Device = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStructs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStructs
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
func (m *VersionList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStructs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VersionList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Versions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStructs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Versions = append(m.Versions, FileVersion{})
			if err := m.Versions[len(m.Versions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStructs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStructs
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
func (m *FileInfoTruncated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStructs
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FileInfoTruncated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileInfoTruncated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStructs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (protocol.FileInfoType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Size", wireType)
			}
			m.Size = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Size |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Permissions", wireType)
			}
			m.Permissions = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Permissions |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModifiedS", wireType)
			}
			m.ModifiedS = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ModifiedS |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deleted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Deleted = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Invalid", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Invalid = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoPermissions", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NoPermissions = bool(v != 0)
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthStructs
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Version.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModifiedNs", wireType)
			}
			m.ModifiedNs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ModifiedNs |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModifiedBy", wireType)
			}
			m.ModifiedBy = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ModifiedBy |= (github_com_syncthing_syncthing_lib_protocol.ShortID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SymlinkTarget", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStructs
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SymlinkTarget = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStructs(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStructs
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
func skipStructs(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStructs
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
					return 0, ErrIntOverflowStructs
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStructs
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthStructs
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStructs
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipStructs(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthStructs = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStructs   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("structs.proto", fileDescriptorStructs) }

var fileDescriptorStructs = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4f, 0x6b, 0xdb, 0x4e,
	0x10, 0xf5, 0xc6, 0x8a, 0xff, 0xac, 0xe2, 0xfc, 0x7e, 0x59, 0x4a, 0x58, 0x0c, 0x95, 0x85, 0xa1,
	0x20, 0x0a, 0x95, 0x5b, 0x87, 0x5e, 0xda, 0x9b, 0x29, 0x81, 0x40, 0x29, 0x45, 0x09, 0xe9, 0xa5,
	0x60, 0x2c, 0x69, 0x2c, 0x2f, 0x95, 0x76, 0x55, 0xed, 0xca, 0xa0, 0x7e, 0x92, 0x1e, 0xf3, 0x71,
	0x7c, 0xec, 0xb9, 0x87, 0xd0, 0xba, 0x9f, 0xa3, 0x50, 0xb4, 0x92, 0x15, 0x1d, 0xdb, 0xdb, 0xbc,
	0xd9, 0xf7, 0xe6, 0xbd, 0x61, 0x16, 0x8f, 0xa4, 0xca, 0xf2, 0x40, 0x49, 0x37, 0xcd, 0x84, 0x12,
	0xe4, 0x28, 0xf4, 0xc7, 0xcf, 0x22, 0xa6, 0x36, 0xb9, 0xef, 0x06, 0x22, 0x99, 0x45, 0x22, 0x12,
	0x33, 0xfd, 0xe4, 0xe7, 0x6b, 0x8d, 0x34, 0xd0, 0x55, 0x25, 0x19, 0xbf, 0x6c, 0xd1, 0x65, 0xc1,
	0x03, 0xb5, 0x61, 0x3c, 0x6a, 0x55, 0x31, 0xf3, 0xab, 0x09, 0x81, 0x88, 0x67, 0x3e, 0xa4, 0x95,
	0x6c, 0xfa, 0x01, 0x9b, 0x97, 0x2c, 0x86, 0x5b, 0xc8, 0x24, 0x13, 0x9c, 0x3c, 0xc7, 0xfd, 0x6d,
	0x55, 0x52, 0x64, 0x23, 0xc7, 0x9c, 0xff, 0xef, 0x1e, 0x44, 0xee, 0x2d, 0x04, 0x4a, 0x64, 0x0b,
	0x63, 0x77, 0x3f, 0xe9, 0x78, 0x07, 0x1a, 0x39, 0xc7, 0xbd, 0x10, 0xb6, 0x2c, 0x00, 0x7a, 0x64,
	0x23, 0xe7, 0xc4, 0xab, 0xd1, 0xf4, 0x12, 0x9b, 0xf5, 0xd0, 0xb7, 0x4c, 0x2a, 0xf2, 0x02, 0x0f,
	0x6a, 0x85, 0xa4, 0xc8, 0xee, 0x3a, 0xe6, 0xfc, 0x3f, 0x37, 0xf4, 0xdd, 0x96, 0x77, 0x3d, 0xb8,
	0xa1, 0xbd, 0x32, 0xbe, 0xde, 0x4d, 0x3a, 0xd3, 0xdf, 0x5d, 0x7c, 0x56, 0xb2, 0xae, 0xf8, 0x5a,
	0xdc, 0x64, 0x39, 0x0f, 0x56, 0x0a, 0x42, 0x42, 0xb0, 0xc1, 0x57, 0x09, 0xe8, 0x90, 0x43, 0x4f,
	0xd7, 0xe4, 0x29, 0x36, 0x54, 0x91, 0x56, 0x39, 0x4e, 0xe7, 0xe7, 0x0f, 0xc1, 0x1b, 0x79, 0x91,
	0x82, 0xa7, 0x39, 0xa5, 0x5e, 0xb2, 0x2f, 0x40, 0xbb, 0x36, 0x72, 0xba, 0x9e, 0xae, 0x89, 0x8d,
	0xcd, 0x14, 0xb2, 0x84, 0xc9, 0x2a, 0xa5, 0x61, 0x23, 0x67, 0xe4, 0xb5, 0x5b, 0xe4, 0x31, 0xc6,
	0x89, 0x08, 0xd9, 0x9a, 0x41, 0xb8, 0x94, 0xf4, 0x58, 0x6b, 0x87, 0x87, 0xce, 0x35, 0xa1, 0xb8,
	0x1f, 0x42, 0x0c, 0x0a, 0x42, 0xda, 0xb3, 0x91, 0x33, 0xf0, 0x0e, 0xb0, 0x7c, 0x61, 0x7c, 0xbb,
	0x8a, 0x59, 0x48, 0xfb, 0xd5, 0x4b, 0x0d, 0xc9, 0x13, 0x7c, 0xca, 0xc5, 0xb2, 0xed, 0x3b, 0xd0,
	0x84, 0x11, 0x17, 0xef, 0x5b, 0xce, 0xad, 0xbb, 0x0c, 0xff, 0xee, 0x2e, 0x63, 0x3c, 0x90, 0xf0,
	0x39, 0x07, 0x1e, 0x00, 0xc5, 0x3a, 0x69, 0x83, 0xc9, 0x04, 0x9b, 0xcd, 0x1e, 0x5c, 0x52, 0xd3,
	0x46, 0xce, 0xb1, 0xd7, 0xac, 0xf6, 0x4e, 0x92, 0x8f, 0x2d, 0x82, 0x5f, 0xd0, 0x13, 0x1b, 0x39,
	0xc6, 0xe2, 0x75, 0x69, 0xf0, 0xfd, 0x7e, 0x72, 0xf1, 0x0f, 0x3f, 0xcd, 0xbd, 0xde, 0x88, 0x4c,
	0x5d, 0xbd, 0x79, 0x98, 0xbe, 0x28, 0xca, 0x9d, 0x65, 0x91, 0xc4, 0x8c, 0x7f, 0x5a, 0xaa, 0x55,
	0x16, 0x81, 0xa2, 0x67, 0xfa, 0x8c, 0xa3, 0xba, 0x7b, 0xa3, 0x9b, 0xd5, 0xfd, 0x17, 0x8f, 0x76,
	0x3f, 0xad, 0xce, 0x6e, 0x6f, 0xa1, 0x6f, 0x7b, 0x0b, 0xfd, 0xd8, 0x5b, 0x9d, 0xbb, 0x5f, 0x16,
	0xf2, 0x7b, 0xda, 0xe0, 0xe2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xcd, 0xe3, 0xfd, 0x38,
	0x03, 0x00, 0x00,
}
