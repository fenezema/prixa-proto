// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/diagnostic/v1/EmailNotification.proto

package prixa_emailnotification_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type EmailDiagnosticResultRequest struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Gender               string   `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	AgeRange             string   `protobuf:"bytes,4,opt,name=ageRange,proto3" json:"ageRange,omitempty"`
	WeightInKg           float32  `protobuf:"fixed32,5,opt,name=weightInKg,proto3" json:"weightInKg,omitempty"`
	HeightInCm           float32  `protobuf:"fixed32,6,opt,name=heightInCm,proto3" json:"heightInCm,omitempty"`
	BmiIndex             float32  `protobuf:"fixed32,7,opt,name=bmiIndex,proto3" json:"bmiIndex,omitempty"`
	Subject              string   `protobuf:"bytes,8,opt,name=subject,proto3" json:"subject,omitempty"`
	Preconditions        string   `protobuf:"bytes,9,opt,name=preconditions,proto3" json:"preconditions,omitempty"`
	MainSymptoms         string   `protobuf:"bytes,10,opt,name=mainSymptoms,proto3" json:"mainSymptoms,omitempty"`
	Symptoms             string   `protobuf:"bytes,11,opt,name=symptoms,proto3" json:"symptoms,omitempty"`
	NotKnown             string   `protobuf:"bytes,12,opt,name=notKnown,proto3" json:"notKnown,omitempty"`
	Unknown              string   `protobuf:"bytes,13,opt,name=unknown,proto3" json:"unknown,omitempty"`
	PrixaResults         string   `protobuf:"bytes,14,opt,name=prixaResults,proto3" json:"prixaResults,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailDiagnosticResultRequest) Reset()         { *m = EmailDiagnosticResultRequest{} }
func (m *EmailDiagnosticResultRequest) String() string { return proto.CompactTextString(m) }
func (*EmailDiagnosticResultRequest) ProtoMessage()    {}
func (*EmailDiagnosticResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b1256ac4bae3b857, []int{0}
}

func (m *EmailDiagnosticResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailDiagnosticResultRequest.Unmarshal(m, b)
}
func (m *EmailDiagnosticResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailDiagnosticResultRequest.Marshal(b, m, deterministic)
}
func (m *EmailDiagnosticResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailDiagnosticResultRequest.Merge(m, src)
}
func (m *EmailDiagnosticResultRequest) XXX_Size() int {
	return xxx_messageInfo_EmailDiagnosticResultRequest.Size(m)
}
func (m *EmailDiagnosticResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailDiagnosticResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailDiagnosticResultRequest proto.InternalMessageInfo

func (m *EmailDiagnosticResultRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *EmailDiagnosticResultRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EmailDiagnosticResultRequest) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *EmailDiagnosticResultRequest) GetAgeRange() string {
	if m != nil {
		return m.AgeRange
	}
	return ""
}

func (m *EmailDiagnosticResultRequest) GetWeightInKg() float32 {
	if m != nil {
		return m.WeightInKg
	}
	return 0
}

func (m *EmailDiagnosticResultRequest) GetHeightInCm() float32 {
	if m != nil {
		return m.HeightInCm
	}
	return 0
}

func (m *EmailDiagnosticResultRequest) GetBmiIndex() float32 {
	if m != nil {
		return m.BmiIndex
	}
	return 0
}

func (m *EmailDiagnosticResultRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *EmailDiagnosticResultRequest) GetPreconditions() string {
	if m != nil {
		return m.Preconditions
	}
	return ""
}

func (m *EmailDiagnosticResultRequest) GetMainSymptoms() string {
	if m != nil {
		return m.MainSymptoms
	}
	return ""
}

func (m *EmailDiagnosticResultRequest) GetSymptoms() string {
	if m != nil {
		return m.Symptoms
	}
	return ""
}

func (m *EmailDiagnosticResultRequest) GetNotKnown() string {
	if m != nil {
		return m.NotKnown
	}
	return ""
}

func (m *EmailDiagnosticResultRequest) GetUnknown() string {
	if m != nil {
		return m.Unknown
	}
	return ""
}

func (m *EmailDiagnosticResultRequest) GetPrixaResults() string {
	if m != nil {
		return m.PrixaResults
	}
	return ""
}

func init() {
	proto.RegisterType((*EmailDiagnosticResultRequest)(nil), "prixa.emailnotification.v1.EmailDiagnosticResultRequest")
}

func init() {
	proto.RegisterFile("proto/diagnostic/v1/EmailNotification.proto", fileDescriptor_b1256ac4bae3b857)
}

var fileDescriptor_b1256ac4bae3b857 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0x86, 0x95, 0xb4, 0x6c, 0x5b, 0xd3, 0xf6, 0x60, 0xa0, 0x32, 0xa1, 0x82, 0x6a, 0xc5, 0xa1,
	0x02, 0x29, 0x56, 0xe1, 0x82, 0xb8, 0x42, 0x0f, 0x55, 0x25, 0x0e, 0xe9, 0x13, 0x78, 0x93, 0xa9,
	0x6b, 0x58, 0x8f, 0x43, 0xec, 0x6c, 0xdb, 0x2b, 0x47, 0xae, 0xbc, 0x01, 0x67, 0xde, 0x86, 0x57,
	0xe0, 0x41, 0x90, 0xc7, 0xc9, 0x92, 0x4a, 0xc0, 0x2d, 0xff, 0xff, 0xd9, 0x33, 0x13, 0xcf, 0xcf,
	0x5e, 0xb6, 0x9d, 0x0b, 0x4e, 0x36, 0x46, 0x69, 0x74, 0x3e, 0x98, 0x5a, 0xae, 0x4e, 0xe4, 0xa9,
	0x55, 0x66, 0xf9, 0xc1, 0x05, 0x73, 0x69, 0x6a, 0x15, 0x8c, 0xc3, 0x92, 0x4e, 0xf1, 0xa2, 0xed,
	0xcc, 0x8d, 0x2a, 0x21, 0x62, 0x9c, 0xe2, 0xd5, 0x49, 0xf1, 0x44, 0x3b, 0xa7, 0x97, 0x20, 0xe9,
	0xe4, 0xa2, 0xbf, 0x94, 0x60, 0xdb, 0x70, 0x9b, 0x2e, 0x16, 0x87, 0x03, 0x54, 0xad, 0x91, 0x0a,
	0xd1, 0x05, 0xba, 0xe6, 0x13, 0x9d, 0xff, 0xd8, 0x60, 0x87, 0xd4, 0xf2, 0xfd, 0x7a, 0x8a, 0x0a,
	0x7c, 0xbf, 0x0c, 0x15, 0x7c, 0xee, 0xc1, 0x07, 0xbe, 0xcf, 0xf2, 0xe0, 0x44, 0x76, 0x94, 0x1d,
	0xef, 0x54, 0x79, 0x70, 0x9c, 0xb3, 0x4d, 0x54, 0x16, 0x44, 0x4e, 0x0e, 0x7d, 0xf3, 0x03, 0x36,
	0xd3, 0x80, 0x0d, 0x74, 0x62, 0x83, 0xdc, 0x41, 0xf1, 0x82, 0x6d, 0x2b, 0x0d, 0x95, 0x42, 0x0d,
	0x62, 0x93, 0xc8, 0x5a, 0xf3, 0xa7, 0x8c, 0x5d, 0x83, 0xd1, 0x57, 0xe1, 0x0c, 0xcf, 0xb5, 0xb8,
	0x77, 0x94, 0x1d, 0xe7, 0xd5, 0xc4, 0x89, 0xfc, 0x6a, 0x50, 0xef, 0xac, 0x98, 0x25, 0xfe, 0xc7,
	0x89, 0xb5, 0x17, 0xd6, 0x9c, 0x61, 0x03, 0x37, 0x62, 0x8b, 0xe8, 0x5a, 0x73, 0xc1, 0xb6, 0x7c,
	0xbf, 0xf8, 0x08, 0x75, 0x10, 0xdb, 0xd4, 0x76, 0x94, 0xfc, 0x39, 0xdb, 0x6b, 0x3b, 0xa8, 0x1d,
	0x36, 0x86, 0x5e, 0x41, 0xec, 0x10, 0xbf, 0x6b, 0xf2, 0x39, 0xdb, 0xb5, 0xca, 0xe0, 0xc5, 0xad,
	0x6d, 0x83, 0xb3, 0x5e, 0x30, 0x3a, 0x74, 0xc7, 0x8b, 0xfd, 0xfd, 0xc8, 0xef, 0xa7, 0x7f, 0xf3,
	0x13, 0x86, 0x2e, 0x9c, 0xa3, 0xbb, 0x46, 0xb1, 0x9b, 0xd8, 0xa8, 0xe3, 0x6c, 0x3d, 0x7e, 0x22,
	0xb4, 0x97, 0x66, 0x1b, 0x64, 0xec, 0x4a, 0x3b, 0x4e, 0xef, 0xef, 0xc5, 0x7e, 0xea, 0x3a, 0xf5,
	0x5e, 0x7d, 0xcf, 0xd8, 0x83, 0x69, 0x38, 0x2e, 0xa0, 0x5b, 0x99, 0x1a, 0xf8, 0xd7, 0x8c, 0x3d,
	0xfa, 0xeb, 0x1a, 0xf9, 0x9b, 0xf2, 0xdf, 0xc1, 0x29, 0xff, 0xb7, 0xf9, 0xe2, 0xa0, 0x4c, 0xc9,
	0x29, 0xc7, 0x58, 0x95, 0xa7, 0x31, 0x56, 0xf3, 0x67, 0x5f, 0x7e, 0xfe, 0xfa, 0x96, 0x3f, 0x9e,
	0x3f, 0x8c, 0x71, 0xa5, 0xe2, 0x92, 0x8a, 0x4b, 0x0f, 0xd8, 0xbc, 0xcd, 0x5e, 0x2c, 0x66, 0x74,
	0xe1, 0xf5, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x64, 0xef, 0x6e, 0x5b, 0xe0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
	EmailDiagnosticResult(ctx context.Context, in *EmailDiagnosticResultRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type notificationServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotificationServiceClient(cc *grpc.ClientConn) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) EmailDiagnosticResult(ctx context.Context, in *EmailDiagnosticResultRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/prixa.emailnotification.v1.NotificationService/EmailDiagnosticResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
type NotificationServiceServer interface {
	EmailDiagnosticResult(context.Context, *EmailDiagnosticResultRequest) (*empty.Empty, error)
}

// UnimplementedNotificationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (*UnimplementedNotificationServiceServer) EmailDiagnosticResult(ctx context.Context, req *EmailDiagnosticResultRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmailDiagnosticResult not implemented")
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_EmailDiagnosticResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailDiagnosticResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).EmailDiagnosticResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prixa.emailnotification.v1.NotificationService/EmailDiagnosticResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).EmailDiagnosticResult(ctx, req.(*EmailDiagnosticResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "prixa.emailnotification.v1.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmailDiagnosticResult",
			Handler:    _NotificationService_EmailDiagnosticResult_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/diagnostic/v1/EmailNotification.proto",
}
