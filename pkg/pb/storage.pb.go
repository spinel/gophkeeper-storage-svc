// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/pb/storage.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateEntityRequest struct {
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	TypeID     int64  `protobuf:"varint,2,opt,name=typeID,proto3" json:"typeID,omitempty"`
	// Types that are valid to be assigned to Entity:
	//	*CreateEntityRequest_Account
	//	*CreateEntityRequest_CreditCard
	Entity               isCreateEntityRequest_Entity `protobuf_oneof:"entity"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CreateEntityRequest) Reset()         { *m = CreateEntityRequest{} }
func (m *CreateEntityRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEntityRequest) ProtoMessage()    {}
func (*CreateEntityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b273432c197719ba, []int{0}
}

func (m *CreateEntityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntityRequest.Unmarshal(m, b)
}
func (m *CreateEntityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntityRequest.Marshal(b, m, deterministic)
}
func (m *CreateEntityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntityRequest.Merge(m, src)
}
func (m *CreateEntityRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEntityRequest.Size(m)
}
func (m *CreateEntityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntityRequest proto.InternalMessageInfo

func (m *CreateEntityRequest) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *CreateEntityRequest) GetTypeID() int64 {
	if m != nil {
		return m.TypeID
	}
	return 0
}

type isCreateEntityRequest_Entity interface {
	isCreateEntityRequest_Entity()
}

type CreateEntityRequest_Account struct {
	Account *Account `protobuf:"bytes,3,opt,name=account,proto3,oneof"`
}

type CreateEntityRequest_CreditCard struct {
	CreditCard *CreditCard `protobuf:"bytes,4,opt,name=credit_card,json=creditCard,proto3,oneof"`
}

func (*CreateEntityRequest_Account) isCreateEntityRequest_Entity() {}

func (*CreateEntityRequest_CreditCard) isCreateEntityRequest_Entity() {}

func (m *CreateEntityRequest) GetEntity() isCreateEntityRequest_Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *CreateEntityRequest) GetAccount() *Account {
	if x, ok := m.GetEntity().(*CreateEntityRequest_Account); ok {
		return x.Account
	}
	return nil
}

func (m *CreateEntityRequest) GetCreditCard() *CreditCard {
	if x, ok := m.GetEntity().(*CreateEntityRequest_CreditCard); ok {
		return x.CreditCard
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CreateEntityRequest) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CreateEntityRequest_OneofMarshaler, _CreateEntityRequest_OneofUnmarshaler, _CreateEntityRequest_OneofSizer, []interface{}{
		(*CreateEntityRequest_Account)(nil),
		(*CreateEntityRequest_CreditCard)(nil),
	}
}

func _CreateEntityRequest_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CreateEntityRequest)
	// entity
	switch x := m.Entity.(type) {
	case *CreateEntityRequest_Account:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Account); err != nil {
			return err
		}
	case *CreateEntityRequest_CreditCard:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CreditCard); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CreateEntityRequest.Entity has unexpected type %T", x)
	}
	return nil
}

func _CreateEntityRequest_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CreateEntityRequest)
	switch tag {
	case 3: // entity.account
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Account)
		err := b.DecodeMessage(msg)
		m.Entity = &CreateEntityRequest_Account{msg}
		return true, err
	case 4: // entity.credit_card
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(CreditCard)
		err := b.DecodeMessage(msg)
		m.Entity = &CreateEntityRequest_CreditCard{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CreateEntityRequest_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CreateEntityRequest)
	// entity
	switch x := m.Entity.(type) {
	case *CreateEntityRequest_Account:
		s := proto.Size(x.Account)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CreateEntityRequest_CreditCard:
		s := proto.Size(x.CreditCard)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CreateEntityResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id                   int64    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateEntityResponse) Reset()         { *m = CreateEntityResponse{} }
func (m *CreateEntityResponse) String() string { return proto.CompactTextString(m) }
func (*CreateEntityResponse) ProtoMessage()    {}
func (*CreateEntityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b273432c197719ba, []int{1}
}

func (m *CreateEntityResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEntityResponse.Unmarshal(m, b)
}
func (m *CreateEntityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEntityResponse.Marshal(b, m, deterministic)
}
func (m *CreateEntityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEntityResponse.Merge(m, src)
}
func (m *CreateEntityResponse) XXX_Size() int {
	return xxx_messageInfo_CreateEntityResponse.Size(m)
}
func (m *CreateEntityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEntityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEntityResponse proto.InternalMessageInfo

func (m *CreateEntityResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CreateEntityResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateEntityResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Account struct {
	Login                string   `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_b273432c197719ba, []int{2}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *Account) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreditCard struct {
	Number               string   `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	HolderName           string   `protobuf:"bytes,2,opt,name=holder_name,json=holderName,proto3" json:"holder_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreditCard) Reset()         { *m = CreditCard{} }
func (m *CreditCard) String() string { return proto.CompactTextString(m) }
func (*CreditCard) ProtoMessage()    {}
func (*CreditCard) Descriptor() ([]byte, []int) {
	return fileDescriptor_b273432c197719ba, []int{3}
}

func (m *CreditCard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreditCard.Unmarshal(m, b)
}
func (m *CreditCard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreditCard.Marshal(b, m, deterministic)
}
func (m *CreditCard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreditCard.Merge(m, src)
}
func (m *CreditCard) XXX_Size() int {
	return xxx_messageInfo_CreditCard.Size(m)
}
func (m *CreditCard) XXX_DiscardUnknown() {
	xxx_messageInfo_CreditCard.DiscardUnknown(m)
}

var xxx_messageInfo_CreditCard proto.InternalMessageInfo

func (m *CreditCard) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *CreditCard) GetHolderName() string {
	if m != nil {
		return m.HolderName
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateEntityRequest)(nil), "storage.CreateEntityRequest")
	proto.RegisterType((*CreateEntityResponse)(nil), "storage.CreateEntityResponse")
	proto.RegisterType((*Account)(nil), "storage.Account")
	proto.RegisterType((*CreditCard)(nil), "storage.CreditCard")
}

func init() { proto.RegisterFile("pkg/pb/storage.proto", fileDescriptor_b273432c197719ba) }

var fileDescriptor_b273432c197719ba = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x4e, 0xe3, 0x30,
	0x10, 0x87, 0x9b, 0x66, 0xb7, 0x7f, 0xa6, 0xab, 0xaa, 0x72, 0xab, 0x55, 0x54, 0xed, 0x42, 0x95,
	0x53, 0x0f, 0xa8, 0x91, 0x8a, 0xc4, 0x85, 0x13, 0x2d, 0x95, 0xca, 0x01, 0x0e, 0x2e, 0x27, 0x2e,
	0x95, 0x13, 0x0f, 0xc5, 0xa2, 0x8d, 0x83, 0xed, 0x80, 0xfa, 0x70, 0xbc, 0x1b, 0x8a, 0xe3, 0x86,
	0x82, 0xc4, 0xf1, 0x1b, 0x7b, 0x26, 0x5f, 0x7e, 0x1e, 0x18, 0x64, 0xcf, 0x9b, 0x28, 0x8b, 0x23,
	0x6d, 0xa4, 0x62, 0x1b, 0x9c, 0x64, 0x4a, 0x1a, 0x49, 0x9a, 0x0e, 0xc3, 0x77, 0x0f, 0xfa, 0x73,
	0x85, 0xcc, 0xe0, 0x22, 0x35, 0xc2, 0xec, 0x29, 0xbe, 0xe4, 0xa8, 0x0d, 0x39, 0x01, 0x10, 0x1c,
	0x53, 0x23, 0x1e, 0x05, 0xaa, 0xc0, 0x1b, 0x79, 0xe3, 0x36, 0x3d, 0xaa, 0x90, 0xbf, 0xd0, 0x30,
	0xfb, 0x0c, 0x6f, 0xae, 0x83, 0xfa, 0xc8, 0x1b, 0xfb, 0xd4, 0x11, 0x39, 0x83, 0x26, 0x4b, 0x12,
	0x99, 0xa7, 0x26, 0xf0, 0x47, 0xde, 0xb8, 0x33, 0xed, 0x4d, 0x0e, 0x5f, 0xbe, 0x2a, 0xeb, 0xcb,
	0x1a, 0x3d, 0x5c, 0x21, 0x17, 0xd0, 0x49, 0x14, 0x72, 0x61, 0xd6, 0x09, 0x53, 0x3c, 0xf8, 0x65,
	0x3b, 0xfa, 0x55, 0xc7, 0xdc, 0x9e, 0xcd, 0x99, 0xe2, 0xcb, 0x1a, 0x85, 0xa4, 0xa2, 0x59, 0x0b,
	0x1a, 0x68, 0x75, 0xc3, 0x7b, 0x18, 0x7c, 0xd5, 0xd7, 0x99, 0x4c, 0x35, 0x16, 0x7e, 0xda, 0x30,
	0x93, 0x6b, 0xeb, 0xee, 0x53, 0x47, 0x64, 0x00, 0xbf, 0x51, 0x29, 0xa9, 0xac, 0x76, 0x9b, 0x96,
	0x40, 0xba, 0x50, 0x17, 0xdc, 0x0a, 0xfb, 0xb4, 0x2e, 0x78, 0x78, 0x09, 0x4d, 0x67, 0x5b, 0x34,
	0x6c, 0xe5, 0x46, 0xa4, 0x2e, 0x83, 0x12, 0xc8, 0x10, 0x5a, 0x19, 0xd3, 0xfa, 0x4d, 0x2a, 0xee,
	0x26, 0x55, 0x1c, 0x2e, 0x00, 0x3e, 0xc5, 0x0b, 0x91, 0x34, 0xdf, 0xc5, 0x55, 0x88, 0x8e, 0xc8,
	0x29, 0x74, 0x9e, 0xe4, 0x96, 0xa3, 0x5a, 0xa7, 0x6c, 0x87, 0x6e, 0x08, 0x94, 0xa5, 0x3b, 0xb6,
	0xc3, 0xe9, 0x1a, 0xba, 0xab, 0x32, 0x87, 0x15, 0xaa, 0x57, 0x91, 0x20, 0xb9, 0x85, 0x3f, 0xc7,
	0xff, 0x4a, 0xfe, 0x1d, 0x07, 0xf5, 0xfd, 0x05, 0x87, 0xff, 0x7f, 0x38, 0x2d, 0x03, 0x0a, 0x6b,
	0x33, 0xf2, 0xd0, 0x9b, 0x44, 0xc5, 0x76, 0xb8, 0x7b, 0x51, 0x16, 0xc7, 0x0d, 0xbb, 0x1e, 0xe7,
	0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86, 0x8a, 0x40, 0x62, 0x36, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageServiceClient is the client API for StorageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageServiceClient interface {
	CreateEntity(ctx context.Context, in *CreateEntityRequest, opts ...grpc.CallOption) (*CreateEntityResponse, error)
}

type storageServiceClient struct {
	cc *grpc.ClientConn
}

func NewStorageServiceClient(cc *grpc.ClientConn) StorageServiceClient {
	return &storageServiceClient{cc}
}

func (c *storageServiceClient) CreateEntity(ctx context.Context, in *CreateEntityRequest, opts ...grpc.CallOption) (*CreateEntityResponse, error) {
	out := new(CreateEntityResponse)
	err := c.cc.Invoke(ctx, "/storage.StorageService/CreateEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServiceServer is the server API for StorageService service.
type StorageServiceServer interface {
	CreateEntity(context.Context, *CreateEntityRequest) (*CreateEntityResponse, error)
}

func RegisterStorageServiceServer(s *grpc.Server, srv StorageServiceServer) {
	s.RegisterService(&_StorageService_serviceDesc, srv)
}

func _StorageService_CreateEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEntityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServiceServer).CreateEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storage.StorageService/CreateEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServiceServer).CreateEntity(ctx, req.(*CreateEntityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StorageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storage.StorageService",
	HandlerType: (*StorageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEntity",
			Handler:    _StorageService_CreateEntity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/storage.proto",
}