// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fiamma/zkpverify/zkpverify.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type VerifyData struct {
	ProofSystem uint64 `protobuf:"varint,1,opt,name=proofSystem,proto3" json:"proofSystem,omitempty"`
	Proof       []byte `protobuf:"bytes,2,opt,name=proof,proto3" json:"proof,omitempty"`
	PublicInput []byte `protobuf:"bytes,3,opt,name=publicInput,proto3" json:"publicInput,omitempty"`
	Vk          []byte `protobuf:"bytes,4,opt,name=vk,proto3" json:"vk,omitempty"`
}

func (m *VerifyData) Reset()         { *m = VerifyData{} }
func (m *VerifyData) String() string { return proto.CompactTextString(m) }
func (*VerifyData) ProtoMessage()    {}
func (*VerifyData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d31b59671d261aa, []int{0}
}
func (m *VerifyData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerifyData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerifyData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerifyData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyData.Merge(m, src)
}
func (m *VerifyData) XXX_Size() int {
	return m.Size()
}
func (m *VerifyData) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyData.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyData proto.InternalMessageInfo

func (m *VerifyData) GetProofSystem() uint64 {
	if m != nil {
		return m.ProofSystem
	}
	return 0
}

func (m *VerifyData) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

func (m *VerifyData) GetPublicInput() []byte {
	if m != nil {
		return m.PublicInput
	}
	return nil
}

func (m *VerifyData) GetVk() []byte {
	if m != nil {
		return m.Vk
	}
	return nil
}

type VerifyResult struct {
	VerifyId       string `protobuf:"bytes,1,opt,name=VerifyId,proto3" json:"VerifyId,omitempty"`
	DataCommitment string `protobuf:"bytes,2,opt,name=dataCommitment,proto3" json:"dataCommitment,omitempty"`
	Result         bool   `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *VerifyResult) Reset()         { *m = VerifyResult{} }
func (m *VerifyResult) String() string { return proto.CompactTextString(m) }
func (*VerifyResult) ProtoMessage()    {}
func (*VerifyResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d31b59671d261aa, []int{1}
}
func (m *VerifyResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerifyResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerifyResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerifyResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyResult.Merge(m, src)
}
func (m *VerifyResult) XXX_Size() int {
	return m.Size()
}
func (m *VerifyResult) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyResult.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyResult proto.InternalMessageInfo

func (m *VerifyResult) GetVerifyId() string {
	if m != nil {
		return m.VerifyId
	}
	return ""
}

func (m *VerifyResult) GetDataCommitment() string {
	if m != nil {
		return m.DataCommitment
	}
	return ""
}

func (m *VerifyResult) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func init() {
	proto.RegisterType((*VerifyData)(nil), "fiamma.zkpverify.VerifyData")
	proto.RegisterType((*VerifyResult)(nil), "fiamma.zkpverify.VerifyResult")
}

func init() { proto.RegisterFile("fiamma/zkpverify/zkpverify.proto", fileDescriptor_5d31b59671d261aa) }

var fileDescriptor_5d31b59671d261aa = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xcb, 0x4c, 0xcc,
	0xcd, 0x4d, 0xd4, 0xaf, 0xca, 0x2e, 0x28, 0x4b, 0x2d, 0xca, 0x4c, 0xab, 0x44, 0xb0, 0xf4, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x20, 0x2a, 0xf4, 0xe0, 0xe2, 0x4a, 0x65, 0x5c, 0x5c, 0x61,
	0x60, 0x96, 0x4b, 0x62, 0x49, 0xa2, 0x90, 0x02, 0x17, 0x77, 0x41, 0x51, 0x7e, 0x7e, 0x5a, 0x70,
	0x65, 0x71, 0x49, 0x6a, 0xae, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0xb2, 0x90, 0x90, 0x08,
	0x17, 0x2b, 0x98, 0x2b, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x13, 0x04, 0xe1, 0x80, 0xf5, 0x95, 0x26,
	0xe5, 0x64, 0x26, 0x7b, 0xe6, 0x15, 0x94, 0x96, 0x48, 0x30, 0x83, 0xe5, 0x90, 0x85, 0x84, 0xf8,
	0xb8, 0x98, 0xca, 0xb2, 0x25, 0x58, 0xc0, 0x12, 0x4c, 0x65, 0xd9, 0x4a, 0x59, 0x5c, 0x3c, 0x10,
	0x7b, 0x83, 0x52, 0x8b, 0x4b, 0x73, 0x4a, 0x84, 0xa4, 0xb8, 0x38, 0x20, 0x7c, 0xcf, 0x14, 0xb0,
	0xb5, 0x9c, 0x41, 0x70, 0xbe, 0x90, 0x1a, 0x17, 0x5f, 0x4a, 0x62, 0x49, 0xa2, 0x73, 0x7e, 0x6e,
	0x6e, 0x66, 0x49, 0x6e, 0x6a, 0x5e, 0x09, 0xd8, 0x72, 0xce, 0x20, 0x34, 0x51, 0x21, 0x31, 0x2e,
	0xb6, 0x22, 0xb0, 0x69, 0x60, 0x07, 0x70, 0x04, 0x41, 0x79, 0x4e, 0x46, 0x27, 0x1e, 0xc9, 0x31,
	0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb,
	0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x01, 0x0d, 0xb1, 0x0a, 0xa4, 0x30, 0x2b, 0xa9, 0x2c,
	0x48, 0x2d, 0x4e, 0x62, 0x03, 0x07, 0x98, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x95, 0x9f, 0xc1,
	0x1b, 0x54, 0x01, 0x00, 0x00,
}

func (m *VerifyData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerifyData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VerifyData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Vk) > 0 {
		i -= len(m.Vk)
		copy(dAtA[i:], m.Vk)
		i = encodeVarintZkpverify(dAtA, i, uint64(len(m.Vk)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PublicInput) > 0 {
		i -= len(m.PublicInput)
		copy(dAtA[i:], m.PublicInput)
		i = encodeVarintZkpverify(dAtA, i, uint64(len(m.PublicInput)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Proof) > 0 {
		i -= len(m.Proof)
		copy(dAtA[i:], m.Proof)
		i = encodeVarintZkpverify(dAtA, i, uint64(len(m.Proof)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProofSystem != 0 {
		i = encodeVarintZkpverify(dAtA, i, uint64(m.ProofSystem))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *VerifyResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerifyResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VerifyResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Result {
		i--
		if m.Result {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.DataCommitment) > 0 {
		i -= len(m.DataCommitment)
		copy(dAtA[i:], m.DataCommitment)
		i = encodeVarintZkpverify(dAtA, i, uint64(len(m.DataCommitment)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.VerifyId) > 0 {
		i -= len(m.VerifyId)
		copy(dAtA[i:], m.VerifyId)
		i = encodeVarintZkpverify(dAtA, i, uint64(len(m.VerifyId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintZkpverify(dAtA []byte, offset int, v uint64) int {
	offset -= sovZkpverify(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VerifyData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProofSystem != 0 {
		n += 1 + sovZkpverify(uint64(m.ProofSystem))
	}
	l = len(m.Proof)
	if l > 0 {
		n += 1 + l + sovZkpverify(uint64(l))
	}
	l = len(m.PublicInput)
	if l > 0 {
		n += 1 + l + sovZkpverify(uint64(l))
	}
	l = len(m.Vk)
	if l > 0 {
		n += 1 + l + sovZkpverify(uint64(l))
	}
	return n
}

func (m *VerifyResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.VerifyId)
	if l > 0 {
		n += 1 + l + sovZkpverify(uint64(l))
	}
	l = len(m.DataCommitment)
	if l > 0 {
		n += 1 + l + sovZkpverify(uint64(l))
	}
	if m.Result {
		n += 2
	}
	return n
}

func sovZkpverify(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozZkpverify(x uint64) (n int) {
	return sovZkpverify(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VerifyData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZkpverify
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
			return fmt.Errorf("proto: VerifyData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifyData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofSystem", wireType)
			}
			m.ProofSystem = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZkpverify
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProofSystem |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZkpverify
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
				return ErrInvalidLengthZkpverify
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthZkpverify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = append(m.Proof[:0], dAtA[iNdEx:postIndex]...)
			if m.Proof == nil {
				m.Proof = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicInput", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZkpverify
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
				return ErrInvalidLengthZkpverify
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthZkpverify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PublicInput = append(m.PublicInput[:0], dAtA[iNdEx:postIndex]...)
			if m.PublicInput == nil {
				m.PublicInput = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vk", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZkpverify
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
				return ErrInvalidLengthZkpverify
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthZkpverify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Vk = append(m.Vk[:0], dAtA[iNdEx:postIndex]...)
			if m.Vk == nil {
				m.Vk = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipZkpverify(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthZkpverify
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
func (m *VerifyResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowZkpverify
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
			return fmt.Errorf("proto: VerifyResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifyResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZkpverify
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
				return ErrInvalidLengthZkpverify
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthZkpverify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerifyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataCommitment", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZkpverify
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
				return ErrInvalidLengthZkpverify
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthZkpverify
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataCommitment = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowZkpverify
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
			m.Result = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipZkpverify(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthZkpverify
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
func skipZkpverify(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowZkpverify
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
					return 0, ErrIntOverflowZkpverify
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
					return 0, ErrIntOverflowZkpverify
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
				return 0, ErrInvalidLengthZkpverify
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupZkpverify
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthZkpverify
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthZkpverify        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowZkpverify          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupZkpverify = fmt.Errorf("proto: unexpected end of group")
)
