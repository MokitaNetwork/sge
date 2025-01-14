// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fury/dvm/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgPubkeysChangeProposalRequest is the type of request for modification of
// public keys.
type MsgSubmitPubkeysChangeProposalRequest struct {
	// creator is the account address of the creator.
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// ticket is the jwt ticket data.
	Ticket string `protobuf:"bytes,2,opt,name=ticket,proto3" json:"ticket,omitempty"`
}

func (m *MsgSubmitPubkeysChangeProposalRequest) Reset()         { *m = MsgSubmitPubkeysChangeProposalRequest{} }
func (m *MsgSubmitPubkeysChangeProposalRequest) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitPubkeysChangeProposalRequest) ProtoMessage()    {}
func (*MsgSubmitPubkeysChangeProposalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78e7f1885428b6a5, []int{0}
}
func (m *MsgSubmitPubkeysChangeProposalRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitPubkeysChangeProposalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitPubkeysChangeProposalRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitPubkeysChangeProposalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitPubkeysChangeProposalRequest.Merge(m, src)
}
func (m *MsgSubmitPubkeysChangeProposalRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitPubkeysChangeProposalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitPubkeysChangeProposalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitPubkeysChangeProposalRequest proto.InternalMessageInfo

func (m *MsgSubmitPubkeysChangeProposalRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSubmitPubkeysChangeProposalRequest) GetTicket() string {
	if m != nil {
		return m.Ticket
	}
	return ""
}

// MsgPubkeysChangeProposalResponse is the type of response for modification of
// public keys.
type MsgSubmitPubkeysChangeProposalResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *MsgSubmitPubkeysChangeProposalResponse) Reset() {
	*m = MsgSubmitPubkeysChangeProposalResponse{}
}
func (m *MsgSubmitPubkeysChangeProposalResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitPubkeysChangeProposalResponse) ProtoMessage()    {}
func (*MsgSubmitPubkeysChangeProposalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78e7f1885428b6a5, []int{1}
}
func (m *MsgSubmitPubkeysChangeProposalResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitPubkeysChangeProposalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitPubkeysChangeProposalResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitPubkeysChangeProposalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitPubkeysChangeProposalResponse.Merge(m, src)
}
func (m *MsgSubmitPubkeysChangeProposalResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitPubkeysChangeProposalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitPubkeysChangeProposalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitPubkeysChangeProposalResponse proto.InternalMessageInfo

func (m *MsgSubmitPubkeysChangeProposalResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

// MsgVotePubkeysChangeRequest is the type of request to vote on the
// modification of public keys proposal.
type MsgVotePubkeysChangeRequest struct {
	// creator is the account address of the creator.
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// ticket is the jwt ticket data.
	Ticket string `protobuf:"bytes,2,opt,name=ticket,proto3" json:"ticket,omitempty"`
	// voter_key_index is the public key index of the voter in the current list
	// of public keys in the vault.
	VoterKeyIndex uint32 `protobuf:"varint,3,opt,name=voter_key_index,json=voterKeyIndex,proto3" json:"voter_key_index,omitempty"`
}

func (m *MsgVotePubkeysChangeRequest) Reset()         { *m = MsgVotePubkeysChangeRequest{} }
func (m *MsgVotePubkeysChangeRequest) String() string { return proto.CompactTextString(m) }
func (*MsgVotePubkeysChangeRequest) ProtoMessage()    {}
func (*MsgVotePubkeysChangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_78e7f1885428b6a5, []int{2}
}
func (m *MsgVotePubkeysChangeRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVotePubkeysChangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVotePubkeysChangeRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVotePubkeysChangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVotePubkeysChangeRequest.Merge(m, src)
}
func (m *MsgVotePubkeysChangeRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgVotePubkeysChangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVotePubkeysChangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVotePubkeysChangeRequest proto.InternalMessageInfo

func (m *MsgVotePubkeysChangeRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgVotePubkeysChangeRequest) GetTicket() string {
	if m != nil {
		return m.Ticket
	}
	return ""
}

func (m *MsgVotePubkeysChangeRequest) GetVoterKeyIndex() uint32 {
	if m != nil {
		return m.VoterKeyIndex
	}
	return 0
}

// MsgVotePubkeysChangeResponse is the type of response vote for public keys
// modification.
type MsgVotePubkeysChangeResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (m *MsgVotePubkeysChangeResponse) Reset()         { *m = MsgVotePubkeysChangeResponse{} }
func (m *MsgVotePubkeysChangeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgVotePubkeysChangeResponse) ProtoMessage()    {}
func (*MsgVotePubkeysChangeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_78e7f1885428b6a5, []int{3}
}
func (m *MsgVotePubkeysChangeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVotePubkeysChangeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVotePubkeysChangeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVotePubkeysChangeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVotePubkeysChangeResponse.Merge(m, src)
}
func (m *MsgVotePubkeysChangeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgVotePubkeysChangeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVotePubkeysChangeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVotePubkeysChangeResponse proto.InternalMessageInfo

func (m *MsgVotePubkeysChangeResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*MsgSubmitPubkeysChangeProposalRequest)(nil), "furynet.furynetwork.dvm.MsgSubmitPubkeysChangeProposalRequest")
	proto.RegisterType((*MsgSubmitPubkeysChangeProposalResponse)(nil), "furynet.furynetwork.dvm.MsgSubmitPubkeysChangeProposalResponse")
	proto.RegisterType((*MsgVotePubkeysChangeRequest)(nil), "furynet.furynetwork.dvm.MsgVotePubkeysChangeRequest")
	proto.RegisterType((*MsgVotePubkeysChangeResponse)(nil), "furynet.furynetwork.dvm.MsgVotePubkeysChangeResponse")
}

func init() { proto.RegisterFile("fury/dvm/tx.proto", fileDescriptor_78e7f1885428b6a5) }

var fileDescriptor_78e7f1885428b6a5 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0xbb, 0x2d, 0x54, 0x5d, 0x28, 0xea, 0x1e, 0x24, 0xb4, 0x12, 0x4a, 0xc0, 0xd2, 0x8b,
	0x89, 0xe8, 0x45, 0x3d, 0x49, 0x3d, 0x89, 0x14, 0x4a, 0x04, 0x41, 0x2f, 0xa5, 0x49, 0x87, 0x6d,
	0x88, 0xc9, 0xc6, 0x9d, 0x4d, 0xda, 0xbc, 0x85, 0x87, 0x3e, 0x94, 0xc7, 0x1e, 0x3d, 0x4a, 0xfb,
	0x22, 0xd2, 0x34, 0x3d, 0x48, 0xad, 0x2d, 0xbd, 0xed, 0x37, 0xcb, 0xfc, 0xe6, 0x9b, 0x3f, 0xf4,
	0x08, 0x39, 0x58, 0xfd, 0x24, 0xb0, 0xd4, 0xc8, 0x8c, 0xa4, 0x50, 0x82, 0x31, 0xe4, 0x10, 0x82,
	0x1a, 0x0a, 0xe9, 0x9b, 0xc8, 0xc1, 0xec, 0x27, 0x81, 0xf1, 0x42, 0xcf, 0xda, 0xc8, 0x9f, 0x62,
	0x27, 0xf0, 0x54, 0x27, 0x76, 0x7c, 0x48, 0xf1, 0x7e, 0xd0, 0x0b, 0x39, 0x74, 0xa4, 0x88, 0x04,
	0xf6, 0xde, 0x6c, 0x78, 0x8f, 0x01, 0x15, 0xd3, 0xe8, 0x9e, 0x2b, 0xa1, 0xa7, 0x84, 0xd4, 0x48,
	0x9d, 0x34, 0x0f, 0xec, 0xa5, 0x64, 0x27, 0xb4, 0xac, 0x3c, 0xd7, 0x07, 0xa5, 0x15, 0xb3, 0x8f,
	0x5c, 0x19, 0x2d, 0xda, 0xd8, 0x84, 0xc6, 0x48, 0x84, 0x08, 0x73, 0x36, 0xc6, 0xae, 0x0b, 0x88,
	0x19, 0x7b, 0xdf, 0x5e, 0x4a, 0x63, 0x48, 0x6b, 0x6d, 0xe4, 0xcf, 0x42, 0xc1, 0x2f, 0xc2, 0xce,
	0xa6, 0x58, 0x83, 0x1e, 0x26, 0x42, 0x81, 0xec, 0xfa, 0x90, 0x76, 0xbd, 0xb0, 0x0f, 0x23, 0xad,
	0x54, 0x27, 0xcd, 0x8a, 0x5d, 0xc9, 0xc2, 0x8f, 0x90, 0x3e, 0xcc, 0x83, 0xc6, 0x35, 0x3d, 0xfd,
	0xbb, 0xf0, 0x26, 0xcb, 0x97, 0xe3, 0x22, 0x2d, 0xb5, 0x91, 0xb3, 0x31, 0xa1, 0xb5, 0x7f, 0x9a,
	0x67, 0x37, 0xe6, 0xea, 0x3a, 0xcc, 0xad, 0x76, 0x51, 0xbd, 0xdd, 0x25, 0x35, 0x37, 0x9e, 0xd0,
	0xe3, 0x95, 0xae, 0x98, 0xb5, 0x06, 0xb8, 0x6e, 0xf0, 0xd5, 0x8b, 0xed, 0x13, 0x16, 0x75, 0x5b,
	0x77, 0x9f, 0x53, 0x9d, 0x4c, 0xa6, 0x3a, 0xf9, 0x9e, 0xea, 0xe4, 0x63, 0xa6, 0x17, 0x26, 0x33,
	0xbd, 0xf0, 0x35, 0xd3, 0x0b, 0xaf, 0x0d, 0xee, 0xa9, 0x41, 0xec, 0x98, 0xae, 0x08, 0x2c, 0xe4,
	0x70, 0x9e, 0x63, 0xe7, 0x6f, 0x6b, 0xb4, 0xb8, 0xe0, 0x34, 0x02, 0x74, 0xca, 0xd9, 0x15, 0x5f,
	0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x99, 0x7f, 0x08, 0xdf, 0xd9, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// PubkeysChangeProposal defines a method to submit a proposal for changing
	// the allowed public keys.
	SubmitPubkeysChangeProposal(ctx context.Context, in *MsgSubmitPubkeysChangeProposalRequest, opts ...grpc.CallOption) (*MsgSubmitPubkeysChangeProposalResponse, error)
	// VotePubkeysChange defines a method to vote for a proposal for changing the
	// allowed public keys.
	VotePubkeysChange(ctx context.Context, in *MsgVotePubkeysChangeRequest, opts ...grpc.CallOption) (*MsgVotePubkeysChangeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitPubkeysChangeProposal(ctx context.Context, in *MsgSubmitPubkeysChangeProposalRequest, opts ...grpc.CallOption) (*MsgSubmitPubkeysChangeProposalResponse, error) {
	out := new(MsgSubmitPubkeysChangeProposalResponse)
	err := c.cc.Invoke(ctx, "/furynet.furynetwork.dvm.Msg/SubmitPubkeysChangeProposal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) VotePubkeysChange(ctx context.Context, in *MsgVotePubkeysChangeRequest, opts ...grpc.CallOption) (*MsgVotePubkeysChangeResponse, error) {
	out := new(MsgVotePubkeysChangeResponse)
	err := c.cc.Invoke(ctx, "/furynet.furynetwork.dvm.Msg/VotePubkeysChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// PubkeysChangeProposal defines a method to submit a proposal for changing
	// the allowed public keys.
	SubmitPubkeysChangeProposal(context.Context, *MsgSubmitPubkeysChangeProposalRequest) (*MsgSubmitPubkeysChangeProposalResponse, error)
	// VotePubkeysChange defines a method to vote for a proposal for changing the
	// allowed public keys.
	VotePubkeysChange(context.Context, *MsgVotePubkeysChangeRequest) (*MsgVotePubkeysChangeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitPubkeysChangeProposal(ctx context.Context, req *MsgSubmitPubkeysChangeProposalRequest) (*MsgSubmitPubkeysChangeProposalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitPubkeysChangeProposal not implemented")
}
func (*UnimplementedMsgServer) VotePubkeysChange(ctx context.Context, req *MsgVotePubkeysChangeRequest) (*MsgVotePubkeysChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VotePubkeysChange not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitPubkeysChangeProposal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitPubkeysChangeProposalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitPubkeysChangeProposal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/furynet.furynetwork.dvm.Msg/SubmitPubkeysChangeProposal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitPubkeysChangeProposal(ctx, req.(*MsgSubmitPubkeysChangeProposalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_VotePubkeysChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgVotePubkeysChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).VotePubkeysChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/furynet.furynetwork.dvm.Msg/VotePubkeysChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).VotePubkeysChange(ctx, req.(*MsgVotePubkeysChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "furynet.furynetwork.dvm.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitPubkeysChangeProposal",
			Handler:    _Msg_SubmitPubkeysChangeProposal_Handler,
		},
		{
			MethodName: "VotePubkeysChange",
			Handler:    _Msg_VotePubkeysChange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fury/dvm/tx.proto",
}

func (m *MsgSubmitPubkeysChangeProposalRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitPubkeysChangeProposalRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitPubkeysChangeProposalRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ticket) > 0 {
		i -= len(m.Ticket)
		copy(dAtA[i:], m.Ticket)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Ticket)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitPubkeysChangeProposalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitPubkeysChangeProposalResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitPubkeysChangeProposalResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgVotePubkeysChangeRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVotePubkeysChangeRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVotePubkeysChangeRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VoterKeyIndex != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.VoterKeyIndex))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Ticket) > 0 {
		i -= len(m.Ticket)
		copy(dAtA[i:], m.Ticket)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Ticket)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgVotePubkeysChangeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVotePubkeysChangeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVotePubkeysChangeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSubmitPubkeysChangeProposalRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Ticket)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitPubkeysChangeProposalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func (m *MsgVotePubkeysChangeRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Ticket)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.VoterKeyIndex != 0 {
		n += 1 + sovTx(uint64(m.VoterKeyIndex))
	}
	return n
}

func (m *MsgVotePubkeysChangeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSubmitPubkeysChangeProposalRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitPubkeysChangeProposalRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitPubkeysChangeProposalRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticket", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ticket = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSubmitPubkeysChangeProposalResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSubmitPubkeysChangeProposalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitPubkeysChangeProposalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgVotePubkeysChangeRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgVotePubkeysChangeRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVotePubkeysChangeRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ticket", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ticket = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoterKeyIndex", wireType)
			}
			m.VoterKeyIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VoterKeyIndex |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgVotePubkeysChangeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgVotePubkeysChangeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVotePubkeysChangeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
			m.Success = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
