// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: multisig/proposal.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

type VoteOption int32

const (
	VoteOption_YES VoteOption = 0
	VoteOption_NO  VoteOption = 1
)

var VoteOption_name = map[int32]string{
	0: "YES",
	1: "NO",
}

var VoteOption_value = map[string]int32{
	"YES": 0,
	"NO":  1,
}

func (x VoteOption) String() string {
	return proto.EnumName(VoteOption_name, int32(x))
}

func (VoteOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_840e3fa04af7a565, []int{0}
}

// ProposalStatus defines proposal statuses.
type ProposalStatus int32

const (
	// Initial status of a proposal when submitted.
	ProposalStatus_SUBMITTED ProposalStatus = 0
	// Status of a proposal when it passes the group's decision policy.
	ProposalStatus_ACCEPTED ProposalStatus = 1
	// Status of a proposal when it is rejected by the group's decision policy.
	ProposalStatus_REJECTED ProposalStatus = 2
	// Status of a proposal when it is successfully executed by the module.
	ProposalStatus_EXECUTED ProposalStatus = 3
	// Status of a proposal when execution is failed.
	ProposalStatus_FAILED ProposalStatus = 4
)

var ProposalStatus_name = map[int32]string{
	0: "SUBMITTED",
	1: "ACCEPTED",
	2: "REJECTED",
	3: "EXECUTED",
	4: "FAILED",
}

var ProposalStatus_value = map[string]int32{
	"SUBMITTED": 0,
	"ACCEPTED":  1,
	"REJECTED":  2,
	"EXECUTED":  3,
	"FAILED":    4,
}

func (x ProposalStatus) String() string {
	return proto.EnumName(ProposalStatus_name, int32(x))
}

func (ProposalStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_840e3fa04af7a565, []int{1}
}

// Proposal defines a group proposal. Any member of a group can submit a proposal
// for a module to decide upon.
// A proposal consists of a set of `sdk.Msg`s that will be executed if the proposal
// passes as well.
type Proposal struct {
	// Account address of the proposer.
	Proposer string `protobuf:"bytes,1,opt,name=proposer,proto3" json:"proposer,omitempty"`
	// Unique id of the proposal.
	Id uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	// Account address of the group.
	Group string `protobuf:"bytes,3,opt,name=group,proto3" json:"group,omitempty"`
	// Block height when the proposal was submitted.
	SubmitBlock uint64 `protobuf:"varint,5,opt,name=submitBlock,proto3" json:"submitBlock,omitempty"`
	// Status represents the high level position in the life cycle of the proposal. Initial value is Submitted.
	Status ProposalStatus `protobuf:"varint,8,opt,name=status,proto3,enum=core.multisig.ProposalStatus" json:"status,omitempty"`
	// Contains the sums of all votes for this proposal for each vote option.
	// It is empty at submission, and only populated after tallying, at voting end block.
	FinalTallyResult *TallyResult `protobuf:"bytes,9,opt,name=finalTallyResult,proto3" json:"finalTallyResult,omitempty"`
	// Block height before which voting must be done.
	VotingEndBlock uint64 `protobuf:"varint,10,opt,name=votingEndBlock,proto3" json:"votingEndBlock,omitempty"`
	// List of `sdk.Msg`s that will be executed if the proposal passes.
	Messages []*types.Any `protobuf:"bytes,12,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (m *Proposal) Reset()         { *m = Proposal{} }
func (m *Proposal) String() string { return proto.CompactTextString(m) }
func (*Proposal) ProtoMessage()    {}
func (*Proposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_840e3fa04af7a565, []int{0}
}
func (m *Proposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proposal.Merge(m, src)
}
func (m *Proposal) XXX_Size() int {
	return m.Size()
}
func (m *Proposal) XXX_DiscardUnknown() {
	xxx_messageInfo_Proposal.DiscardUnknown(m)
}

var xxx_messageInfo_Proposal proto.InternalMessageInfo

func (m *Proposal) GetProposer() string {
	if m != nil {
		return m.Proposer
	}
	return ""
}

func (m *Proposal) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Proposal) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

func (m *Proposal) GetSubmitBlock() uint64 {
	if m != nil {
		return m.SubmitBlock
	}
	return 0
}

func (m *Proposal) GetStatus() ProposalStatus {
	if m != nil {
		return m.Status
	}
	return ProposalStatus_SUBMITTED
}

func (m *Proposal) GetFinalTallyResult() *TallyResult {
	if m != nil {
		return m.FinalTallyResult
	}
	return nil
}

func (m *Proposal) GetVotingEndBlock() uint64 {
	if m != nil {
		return m.VotingEndBlock
	}
	return 0
}

func (m *Proposal) GetMessages() []*types.Any {
	if m != nil {
		return m.Messages
	}
	return nil
}

// TallyResult represents the sum of votes for each vote option.
type TallyResult struct {
	// Sum of yes votes.
	YesCount uint64 `protobuf:"varint,1,opt,name=yesCount,proto3" json:"yesCount,omitempty"`
	// Sum of no votes.
	NoCount uint64 `protobuf:"varint,3,opt,name=noCount,proto3" json:"noCount,omitempty"`
}

func (m *TallyResult) Reset()         { *m = TallyResult{} }
func (m *TallyResult) String() string { return proto.CompactTextString(m) }
func (*TallyResult) ProtoMessage()    {}
func (*TallyResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_840e3fa04af7a565, []int{1}
}
func (m *TallyResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TallyResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TallyResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TallyResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TallyResult.Merge(m, src)
}
func (m *TallyResult) XXX_Size() int {
	return m.Size()
}
func (m *TallyResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TallyResult.DiscardUnknown(m)
}

var xxx_messageInfo_TallyResult proto.InternalMessageInfo

func (m *TallyResult) GetYesCount() uint64 {
	if m != nil {
		return m.YesCount
	}
	return 0
}

func (m *TallyResult) GetNoCount() uint64 {
	if m != nil {
		return m.NoCount
	}
	return 0
}

// Vote represents a vote for a proposal.
type Vote struct {
	// Unique ID of the proposal.
	ProposalId uint64 `protobuf:"varint,1,opt,name=proposalId,proto3" json:"proposalId,omitempty"`
	// Voter is the account address of the voter.
	Voter string `protobuf:"bytes,2,opt,name=voter,proto3" json:"voter,omitempty"`
	// Option is the voter's choice on the proposal.
	Option VoteOption `protobuf:"varint,3,opt,name=option,proto3,enum=core.multisig.VoteOption" json:"option,omitempty"`
	// Block height when the vote was submitted.
	SubmitBlock uint64 `protobuf:"varint,5,opt,name=submitBlock,proto3" json:"submitBlock,omitempty"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_840e3fa04af7a565, []int{2}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetProposalId() uint64 {
	if m != nil {
		return m.ProposalId
	}
	return 0
}

func (m *Vote) GetVoter() string {
	if m != nil {
		return m.Voter
	}
	return ""
}

func (m *Vote) GetOption() VoteOption {
	if m != nil {
		return m.Option
	}
	return VoteOption_YES
}

func (m *Vote) GetSubmitBlock() uint64 {
	if m != nil {
		return m.SubmitBlock
	}
	return 0
}

func init() {
	proto.RegisterEnum("core.multisig.VoteOption", VoteOption_name, VoteOption_value)
	proto.RegisterEnum("core.multisig.ProposalStatus", ProposalStatus_name, ProposalStatus_value)
	proto.RegisterType((*Proposal)(nil), "core.multisig.Proposal")
	proto.RegisterType((*TallyResult)(nil), "core.multisig.TallyResult")
	proto.RegisterType((*Vote)(nil), "core.multisig.Vote")
}

func init() { proto.RegisterFile("multisig/proposal.proto", fileDescriptor_840e3fa04af7a565) }

var fileDescriptor_840e3fa04af7a565 = []byte{
	// 509 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x6e, 0xda, 0x40,
	0x10, 0x86, 0x59, 0x43, 0x88, 0x19, 0x12, 0x84, 0x56, 0x91, 0xea, 0x20, 0xc5, 0xb2, 0x38, 0x54,
	0x28, 0x52, 0xed, 0x86, 0xaa, 0x0f, 0x00, 0x8e, 0x23, 0x51, 0xb5, 0x0d, 0x5a, 0xa0, 0x6a, 0xab,
	0x5e, 0x0c, 0x6c, 0xdc, 0x6d, 0x8d, 0xd7, 0xf2, 0xae, 0x51, 0x79, 0x8b, 0xf6, 0xa9, 0xda, 0x63,
	0x8e, 0x3d, 0x56, 0xf0, 0x22, 0x95, 0xd7, 0x98, 0x10, 0x7a, 0xe8, 0xf1, 0x9b, 0xf9, 0x67, 0x3c,
	0xfe, 0x7f, 0x2d, 0x3c, 0x59, 0xa4, 0xa1, 0x64, 0x82, 0x05, 0x4e, 0x9c, 0xf0, 0x98, 0x0b, 0x3f,
	0xb4, 0xe3, 0x84, 0x4b, 0x8e, 0x4f, 0x67, 0x3c, 0xa1, 0x76, 0xd1, 0x6d, 0x9d, 0x07, 0x9c, 0x07,
	0x21, 0x75, 0x54, 0x73, 0x9a, 0xde, 0x39, 0x7e, 0xb4, 0xca, 0x95, 0xed, 0x9f, 0x1a, 0xe8, 0xc3,
	0xed, 0x30, 0x6e, 0x81, 0x9e, 0x2f, 0xa2, 0x89, 0x81, 0x2c, 0xd4, 0xa9, 0x91, 0x1d, 0xe3, 0x06,
	0x68, 0x6c, 0x6e, 0x68, 0x16, 0xea, 0x54, 0x88, 0xc6, 0xe6, 0xf8, 0x0c, 0x8e, 0x82, 0x84, 0xa7,
	0xb1, 0x51, 0x56, 0xc2, 0x1c, 0xb0, 0x05, 0x75, 0x91, 0x4e, 0x17, 0x4c, 0xf6, 0x43, 0x3e, 0xfb,
	0x6a, 0x1c, 0x29, 0xf9, 0x7e, 0x09, 0xbf, 0x84, 0xaa, 0x90, 0xbe, 0x4c, 0x85, 0xa1, 0x5b, 0xa8,
	0xd3, 0xe8, 0x5e, 0xd8, 0x8f, 0x6e, 0xb5, 0x8b, 0x63, 0x46, 0x4a, 0x44, 0xb6, 0x62, 0x7c, 0x03,
	0xcd, 0x3b, 0x16, 0xf9, 0xe1, 0xd8, 0x0f, 0xc3, 0x15, 0xa1, 0x22, 0x0d, 0xa5, 0x51, 0xb3, 0x50,
	0xa7, 0xde, 0x6d, 0x1d, 0x2c, 0xd8, 0x53, 0x90, 0x7f, 0x66, 0xf0, 0x53, 0x68, 0x2c, 0xb9, 0x64,
	0x51, 0xe0, 0x45, 0xf3, 0xfc, 0x46, 0x50, 0x37, 0x1e, 0x54, 0xf1, 0x73, 0xd0, 0x17, 0x54, 0x08,
	0x3f, 0xa0, 0xc2, 0x38, 0xb1, 0xca, 0x9d, 0x7a, 0xf7, 0xcc, 0xce, 0x5d, 0xb4, 0x0b, 0x17, 0xed,
	0x5e, 0xb4, 0x22, 0x3b, 0x55, 0xdb, 0x85, 0xfa, 0xfe, 0x87, 0x5a, 0xa0, 0xaf, 0xa8, 0x70, 0x79,
	0x1a, 0x49, 0xe5, 0x65, 0x85, 0xec, 0x18, 0x1b, 0x70, 0x1c, 0xf1, 0xbc, 0x55, 0x56, 0xad, 0x02,
	0xdb, 0x3f, 0x10, 0x54, 0xde, 0x71, 0x49, 0xb1, 0x09, 0x50, 0x64, 0x3a, 0x98, 0x6f, 0x17, 0xec,
	0x55, 0x32, 0xfb, 0x97, 0x5c, 0xd2, 0x44, 0x25, 0x52, 0x23, 0x39, 0xe0, 0x2b, 0xa8, 0xf2, 0x58,
	0x32, 0x1e, 0xa9, 0xbd, 0x8d, 0xee, 0xf9, 0x81, 0x37, 0xd9, 0xea, 0x5b, 0x25, 0x20, 0x5b, 0xe1,
	0xff, 0x13, 0xbb, 0xbc, 0x00, 0x78, 0x98, 0xc3, 0xc7, 0x50, 0xfe, 0xe0, 0x8d, 0x9a, 0x25, 0x5c,
	0x05, 0xed, 0xed, 0x6d, 0x13, 0x5d, 0x4e, 0xa0, 0xf1, 0x38, 0x33, 0x7c, 0x0a, 0xb5, 0xd1, 0xa4,
	0xff, 0x66, 0x30, 0x1e, 0x7b, 0xd7, 0xcd, 0x12, 0x3e, 0x01, 0xbd, 0xe7, 0xba, 0xde, 0x30, 0x23,
	0x94, 0x11, 0xf1, 0x5e, 0x79, 0x6e, 0x46, 0x5a, 0x46, 0xde, 0x7b, 0xcf, 0x9d, 0x64, 0x54, 0xc6,
	0x00, 0xd5, 0x9b, 0xde, 0xe0, 0xb5, 0x77, 0xdd, 0xac, 0xf4, 0x3f, 0xfd, 0x5a, 0x9b, 0xe8, 0x7e,
	0x6d, 0xa2, 0x3f, 0x6b, 0x13, 0x7d, 0xdf, 0x98, 0xa5, 0xfb, 0x8d, 0x59, 0xfa, 0xbd, 0x31, 0x4b,
	0x1f, 0xfb, 0x01, 0x93, 0x9f, 0xd3, 0xa9, 0x3d, 0xe3, 0x0b, 0xa7, 0x9f, 0xb0, 0x79, 0x40, 0x43,
	0x2a, 0xc4, 0xb3, 0x61, 0xc2, 0xbf, 0xd0, 0x99, 0x74, 0xa6, 0x0f, 0xa5, 0xec, 0xe7, 0x9d, 0xe5,
	0x55, 0xd7, 0xf9, 0xe6, 0xec, 0x5e, 0x8a, 0x5c, 0xc5, 0x54, 0x4c, 0xab, 0x2a, 0xc4, 0x17, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x30, 0x8f, 0x87, 0x42, 0x03, 0x00, 0x00,
}

func (m *Proposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Messages[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposal(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if m.VotingEndBlock != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.VotingEndBlock))
		i--
		dAtA[i] = 0x50
	}
	if m.FinalTallyResult != nil {
		{
			size, err := m.FinalTallyResult.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProposal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if m.Status != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x40
	}
	if m.SubmitBlock != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.SubmitBlock))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Group) > 0 {
		i -= len(m.Group)
		copy(dAtA[i:], m.Group)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Group)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Proposer) > 0 {
		i -= len(m.Proposer)
		copy(dAtA[i:], m.Proposer)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Proposer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TallyResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TallyResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TallyResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NoCount != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.NoCount))
		i--
		dAtA[i] = 0x18
	}
	if m.YesCount != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.YesCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Vote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SubmitBlock != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.SubmitBlock))
		i--
		dAtA[i] = 0x28
	}
	if m.Option != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.Option))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Voter) > 0 {
		i -= len(m.Voter)
		copy(dAtA[i:], m.Voter)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Voter)))
		i--
		dAtA[i] = 0x12
	}
	if m.ProposalId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.ProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Proposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Proposer)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovProposal(uint64(m.Id))
	}
	l = len(m.Group)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.SubmitBlock != 0 {
		n += 1 + sovProposal(uint64(m.SubmitBlock))
	}
	if m.Status != 0 {
		n += 1 + sovProposal(uint64(m.Status))
	}
	if m.FinalTallyResult != nil {
		l = m.FinalTallyResult.Size()
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.VotingEndBlock != 0 {
		n += 1 + sovProposal(uint64(m.VotingEndBlock))
	}
	if len(m.Messages) > 0 {
		for _, e := range m.Messages {
			l = e.Size()
			n += 1 + l + sovProposal(uint64(l))
		}
	}
	return n
}

func (m *TallyResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.YesCount != 0 {
		n += 1 + sovProposal(uint64(m.YesCount))
	}
	if m.NoCount != 0 {
		n += 1 + sovProposal(uint64(m.NoCount))
	}
	return n
}

func (m *Vote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalId != 0 {
		n += 1 + sovProposal(uint64(m.ProposalId))
	}
	l = len(m.Voter)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	if m.Option != 0 {
		n += 1 + sovProposal(uint64(m.Option))
	}
	if m.SubmitBlock != 0 {
		n += 1 + sovProposal(uint64(m.SubmitBlock))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Proposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: Proposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Group = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmitBlock", wireType)
			}
			m.SubmitBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubmitBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= ProposalStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FinalTallyResult", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FinalTallyResult == nil {
				m.FinalTallyResult = &TallyResult{}
			}
			if err := m.FinalTallyResult.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingEndBlock", wireType)
			}
			m.VotingEndBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingEndBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages, &types.Any{})
			if err := m.Messages[len(m.Messages)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *TallyResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: TallyResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TallyResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field YesCount", wireType)
			}
			m.YesCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.YesCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NoCount", wireType)
			}
			m.NoCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NoCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *Vote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: Vote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalId", wireType)
			}
			m.ProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Voter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Voter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Option", wireType)
			}
			m.Option = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Option |= VoteOption(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmitBlock", wireType)
			}
			m.SubmitBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubmitBlock |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
