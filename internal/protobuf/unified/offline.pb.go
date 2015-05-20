// Code generated by protoc-gen-go.
// source: steammessages_offline.steamclient.proto
// DO NOT EDIT!

package unified

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type COffline_GetOfflineLogonTicket_Request struct {
	Priority         *uint32 `protobuf:"varint,1,opt,name=priority" json:"priority,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *COffline_GetOfflineLogonTicket_Request) Reset() {
	*m = COffline_GetOfflineLogonTicket_Request{}
}
func (m *COffline_GetOfflineLogonTicket_Request) String() string { return proto.CompactTextString(m) }
func (*COffline_GetOfflineLogonTicket_Request) ProtoMessage()    {}

func (m *COffline_GetOfflineLogonTicket_Request) GetPriority() uint32 {
	if m != nil && m.Priority != nil {
		return *m.Priority
	}
	return 0
}

type COffline_GetOfflineLogonTicket_Response struct {
	SerializedTicket []byte `protobuf:"bytes,1,opt,name=serialized_ticket" json:"serialized_ticket,omitempty"`
	Signature        []byte `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *COffline_GetOfflineLogonTicket_Response) Reset() {
	*m = COffline_GetOfflineLogonTicket_Response{}
}
func (m *COffline_GetOfflineLogonTicket_Response) String() string { return proto.CompactTextString(m) }
func (*COffline_GetOfflineLogonTicket_Response) ProtoMessage()    {}

func (m *COffline_GetOfflineLogonTicket_Response) GetSerializedTicket() []byte {
	if m != nil {
		return m.SerializedTicket
	}
	return nil
}

func (m *COffline_GetOfflineLogonTicket_Response) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type COffline_GetUnsignedOfflineLogonTicket_Request struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *COffline_GetUnsignedOfflineLogonTicket_Request) Reset() {
	*m = COffline_GetUnsignedOfflineLogonTicket_Request{}
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Request) String() string {
	return proto.CompactTextString(m)
}
func (*COffline_GetUnsignedOfflineLogonTicket_Request) ProtoMessage() {}

type COffline_OfflineLogonTicket struct {
	Accountid           *uint32 `protobuf:"varint,1,opt,name=accountid" json:"accountid,omitempty"`
	Rtime32CreationTime *uint32 `protobuf:"fixed32,2,opt,name=rtime32_creation_time" json:"rtime32_creation_time,omitempty"`
	XXX_unrecognized    []byte  `json:"-"`
}

func (m *COffline_OfflineLogonTicket) Reset()         { *m = COffline_OfflineLogonTicket{} }
func (m *COffline_OfflineLogonTicket) String() string { return proto.CompactTextString(m) }
func (*COffline_OfflineLogonTicket) ProtoMessage()    {}

func (m *COffline_OfflineLogonTicket) GetAccountid() uint32 {
	if m != nil && m.Accountid != nil {
		return *m.Accountid
	}
	return 0
}

func (m *COffline_OfflineLogonTicket) GetRtime32CreationTime() uint32 {
	if m != nil && m.Rtime32CreationTime != nil {
		return *m.Rtime32CreationTime
	}
	return 0
}

type COffline_GetUnsignedOfflineLogonTicket_Response struct {
	Ticket           *COffline_OfflineLogonTicket `protobuf:"bytes,1,opt,name=ticket" json:"ticket,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *COffline_GetUnsignedOfflineLogonTicket_Response) Reset() {
	*m = COffline_GetUnsignedOfflineLogonTicket_Response{}
}
func (m *COffline_GetUnsignedOfflineLogonTicket_Response) String() string {
	return proto.CompactTextString(m)
}
func (*COffline_GetUnsignedOfflineLogonTicket_Response) ProtoMessage() {}

func (m *COffline_GetUnsignedOfflineLogonTicket_Response) GetTicket() *COffline_OfflineLogonTicket {
	if m != nil {
		return m.Ticket
	}
	return nil
}

func init() {
}
