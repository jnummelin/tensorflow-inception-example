// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/graph_transfer_info.proto

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type GraphTransferInfo_Destination int32

const (
	GraphTransferInfo_NOP     GraphTransferInfo_Destination = 0
	GraphTransferInfo_HEXAGON GraphTransferInfo_Destination = 1
)

var GraphTransferInfo_Destination_name = map[int32]string{
	0: "NOP",
	1: "HEXAGON",
}
var GraphTransferInfo_Destination_value = map[string]int32{
	"NOP":     0,
	"HEXAGON": 1,
}

func (x GraphTransferInfo_Destination) String() string {
	return proto.EnumName(GraphTransferInfo_Destination_name, int32(x))
}
func (GraphTransferInfo_Destination) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor6, []int{0, 0}
}

// Protocol buffer representing a handle to a tensorflow resource. Handles are
// not valid across executions, but can be serialized back and forth from within
// a single run.
type GraphTransferInfo struct {
	NodeInfo       []*GraphTransferInfo_NodeInfo       `protobuf:"bytes,1,rep,name=node_info,json=nodeInfo" json:"node_info,omitempty"`
	ConstNodeInfo  []*GraphTransferInfo_ConstNodeInfo  `protobuf:"bytes,2,rep,name=const_node_info,json=constNodeInfo" json:"const_node_info,omitempty"`
	NodeInputInfo  []*GraphTransferInfo_NodeInputInfo  `protobuf:"bytes,3,rep,name=node_input_info,json=nodeInputInfo" json:"node_input_info,omitempty"`
	NodeOutputInfo []*GraphTransferInfo_NodeOutputInfo `protobuf:"bytes,4,rep,name=node_output_info,json=nodeOutputInfo" json:"node_output_info,omitempty"`
	// Input Node parameters of transferred graph
	GraphInputNodeInfo  []*GraphTransferInfo_GraphInputNodeInfo  `protobuf:"bytes,5,rep,name=graph_input_node_info,json=graphInputNodeInfo" json:"graph_input_node_info,omitempty"`
	GraphOutputNodeInfo []*GraphTransferInfo_GraphOutputNodeInfo `protobuf:"bytes,6,rep,name=graph_output_node_info,json=graphOutputNodeInfo" json:"graph_output_node_info,omitempty"`
	// Destination of graph transfer
	Destination GraphTransferInfo_Destination `protobuf:"varint,7,opt,name=destination,enum=tensorflow.GraphTransferInfo_Destination" json:"destination,omitempty"`
}

func (m *GraphTransferInfo) Reset()                    { *m = GraphTransferInfo{} }
func (m *GraphTransferInfo) String() string            { return proto.CompactTextString(m) }
func (*GraphTransferInfo) ProtoMessage()               {}
func (*GraphTransferInfo) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *GraphTransferInfo) GetNodeInfo() []*GraphTransferInfo_NodeInfo {
	if m != nil {
		return m.NodeInfo
	}
	return nil
}

func (m *GraphTransferInfo) GetConstNodeInfo() []*GraphTransferInfo_ConstNodeInfo {
	if m != nil {
		return m.ConstNodeInfo
	}
	return nil
}

func (m *GraphTransferInfo) GetNodeInputInfo() []*GraphTransferInfo_NodeInputInfo {
	if m != nil {
		return m.NodeInputInfo
	}
	return nil
}

func (m *GraphTransferInfo) GetNodeOutputInfo() []*GraphTransferInfo_NodeOutputInfo {
	if m != nil {
		return m.NodeOutputInfo
	}
	return nil
}

func (m *GraphTransferInfo) GetGraphInputNodeInfo() []*GraphTransferInfo_GraphInputNodeInfo {
	if m != nil {
		return m.GraphInputNodeInfo
	}
	return nil
}

func (m *GraphTransferInfo) GetGraphOutputNodeInfo() []*GraphTransferInfo_GraphOutputNodeInfo {
	if m != nil {
		return m.GraphOutputNodeInfo
	}
	return nil
}

func (m *GraphTransferInfo) GetDestination() GraphTransferInfo_Destination {
	if m != nil {
		return m.Destination
	}
	return GraphTransferInfo_NOP
}

type GraphTransferInfo_NodeInput struct {
	NodeId     int32 `protobuf:"varint,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	OutputPort int32 `protobuf:"varint,2,opt,name=output_port,json=outputPort" json:"output_port,omitempty"`
}

func (m *GraphTransferInfo_NodeInput) Reset()                    { *m = GraphTransferInfo_NodeInput{} }
func (m *GraphTransferInfo_NodeInput) String() string            { return proto.CompactTextString(m) }
func (*GraphTransferInfo_NodeInput) ProtoMessage()               {}
func (*GraphTransferInfo_NodeInput) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0, 0} }

func (m *GraphTransferInfo_NodeInput) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *GraphTransferInfo_NodeInput) GetOutputPort() int32 {
	if m != nil {
		return m.OutputPort
	}
	return 0
}

type GraphTransferInfo_NodeInfo struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	NodeId      int32  `protobuf:"varint,2,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	TypeName    string `protobuf:"bytes,3,opt,name=type_name,json=typeName" json:"type_name,omitempty"`
	SocOpId     int32  `protobuf:"varint,4,opt,name=soc_op_id,json=socOpId" json:"soc_op_id,omitempty"`
	PaddingId   int32  `protobuf:"varint,5,opt,name=padding_id,json=paddingId" json:"padding_id,omitempty"`
	InputCount  int32  `protobuf:"varint,6,opt,name=input_count,json=inputCount" json:"input_count,omitempty"`
	OutputCount int32  `protobuf:"varint,7,opt,name=output_count,json=outputCount" json:"output_count,omitempty"`
}

func (m *GraphTransferInfo_NodeInfo) Reset()                    { *m = GraphTransferInfo_NodeInfo{} }
func (m *GraphTransferInfo_NodeInfo) String() string            { return proto.CompactTextString(m) }
func (*GraphTransferInfo_NodeInfo) ProtoMessage()               {}
func (*GraphTransferInfo_NodeInfo) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0, 1} }

func (m *GraphTransferInfo_NodeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GraphTransferInfo_NodeInfo) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *GraphTransferInfo_NodeInfo) GetTypeName() string {
	if m != nil {
		return m.TypeName
	}
	return ""
}

func (m *GraphTransferInfo_NodeInfo) GetSocOpId() int32 {
	if m != nil {
		return m.SocOpId
	}
	return 0
}

func (m *GraphTransferInfo_NodeInfo) GetPaddingId() int32 {
	if m != nil {
		return m.PaddingId
	}
	return 0
}

func (m *GraphTransferInfo_NodeInfo) GetInputCount() int32 {
	if m != nil {
		return m.InputCount
	}
	return 0
}

func (m *GraphTransferInfo_NodeInfo) GetOutputCount() int32 {
	if m != nil {
		return m.OutputCount
	}
	return 0
}

type GraphTransferInfo_ConstNodeInfo struct {
	Name   string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	NodeId int32    `protobuf:"varint,2,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	Shape  []int64  `protobuf:"varint,3,rep,packed,name=shape" json:"shape,omitempty"`
	Data   []byte   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Dtype  DataType `protobuf:"varint,5,opt,name=dtype,enum=tensorflow.DataType" json:"dtype,omitempty"`
}

func (m *GraphTransferInfo_ConstNodeInfo) Reset()         { *m = GraphTransferInfo_ConstNodeInfo{} }
func (m *GraphTransferInfo_ConstNodeInfo) String() string { return proto.CompactTextString(m) }
func (*GraphTransferInfo_ConstNodeInfo) ProtoMessage()    {}
func (*GraphTransferInfo_ConstNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{0, 2}
}

func (m *GraphTransferInfo_ConstNodeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GraphTransferInfo_ConstNodeInfo) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *GraphTransferInfo_ConstNodeInfo) GetShape() []int64 {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *GraphTransferInfo_ConstNodeInfo) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GraphTransferInfo_ConstNodeInfo) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

type GraphTransferInfo_NodeInputInfo struct {
	NodeId    int32                          `protobuf:"varint,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	NodeInput []*GraphTransferInfo_NodeInput `protobuf:"bytes,2,rep,name=node_input,json=nodeInput" json:"node_input,omitempty"`
}

func (m *GraphTransferInfo_NodeInputInfo) Reset()         { *m = GraphTransferInfo_NodeInputInfo{} }
func (m *GraphTransferInfo_NodeInputInfo) String() string { return proto.CompactTextString(m) }
func (*GraphTransferInfo_NodeInputInfo) ProtoMessage()    {}
func (*GraphTransferInfo_NodeInputInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{0, 3}
}

func (m *GraphTransferInfo_NodeInputInfo) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *GraphTransferInfo_NodeInputInfo) GetNodeInput() []*GraphTransferInfo_NodeInput {
	if m != nil {
		return m.NodeInput
	}
	return nil
}

type GraphTransferInfo_NodeOutputInfo struct {
	NodeId      int32   `protobuf:"varint,1,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	MaxByteSize []int32 `protobuf:"varint,2,rep,packed,name=max_byte_size,json=maxByteSize" json:"max_byte_size,omitempty"`
}

func (m *GraphTransferInfo_NodeOutputInfo) Reset()         { *m = GraphTransferInfo_NodeOutputInfo{} }
func (m *GraphTransferInfo_NodeOutputInfo) String() string { return proto.CompactTextString(m) }
func (*GraphTransferInfo_NodeOutputInfo) ProtoMessage()    {}
func (*GraphTransferInfo_NodeOutputInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{0, 4}
}

func (m *GraphTransferInfo_NodeOutputInfo) GetNodeId() int32 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *GraphTransferInfo_NodeOutputInfo) GetMaxByteSize() []int32 {
	if m != nil {
		return m.MaxByteSize
	}
	return nil
}

type GraphTransferInfo_GraphInputNodeInfo struct {
	Name  string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Shape []int64  `protobuf:"varint,2,rep,packed,name=shape" json:"shape,omitempty"`
	Dtype DataType `protobuf:"varint,3,opt,name=dtype,enum=tensorflow.DataType" json:"dtype,omitempty"`
}

func (m *GraphTransferInfo_GraphInputNodeInfo) Reset()         { *m = GraphTransferInfo_GraphInputNodeInfo{} }
func (m *GraphTransferInfo_GraphInputNodeInfo) String() string { return proto.CompactTextString(m) }
func (*GraphTransferInfo_GraphInputNodeInfo) ProtoMessage()    {}
func (*GraphTransferInfo_GraphInputNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{0, 5}
}

func (m *GraphTransferInfo_GraphInputNodeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GraphTransferInfo_GraphInputNodeInfo) GetShape() []int64 {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *GraphTransferInfo_GraphInputNodeInfo) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

type GraphTransferInfo_GraphOutputNodeInfo struct {
	Name  string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Shape []int64  `protobuf:"varint,2,rep,packed,name=shape" json:"shape,omitempty"`
	Dtype DataType `protobuf:"varint,3,opt,name=dtype,enum=tensorflow.DataType" json:"dtype,omitempty"`
}

func (m *GraphTransferInfo_GraphOutputNodeInfo) Reset()         { *m = GraphTransferInfo_GraphOutputNodeInfo{} }
func (m *GraphTransferInfo_GraphOutputNodeInfo) String() string { return proto.CompactTextString(m) }
func (*GraphTransferInfo_GraphOutputNodeInfo) ProtoMessage()    {}
func (*GraphTransferInfo_GraphOutputNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor6, []int{0, 6}
}

func (m *GraphTransferInfo_GraphOutputNodeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GraphTransferInfo_GraphOutputNodeInfo) GetShape() []int64 {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *GraphTransferInfo_GraphOutputNodeInfo) GetDtype() DataType {
	if m != nil {
		return m.Dtype
	}
	return DataType_DT_INVALID
}

func init() {
	proto.RegisterType((*GraphTransferInfo)(nil), "tensorflow.GraphTransferInfo")
	proto.RegisterType((*GraphTransferInfo_NodeInput)(nil), "tensorflow.GraphTransferInfo.NodeInput")
	proto.RegisterType((*GraphTransferInfo_NodeInfo)(nil), "tensorflow.GraphTransferInfo.NodeInfo")
	proto.RegisterType((*GraphTransferInfo_ConstNodeInfo)(nil), "tensorflow.GraphTransferInfo.ConstNodeInfo")
	proto.RegisterType((*GraphTransferInfo_NodeInputInfo)(nil), "tensorflow.GraphTransferInfo.NodeInputInfo")
	proto.RegisterType((*GraphTransferInfo_NodeOutputInfo)(nil), "tensorflow.GraphTransferInfo.NodeOutputInfo")
	proto.RegisterType((*GraphTransferInfo_GraphInputNodeInfo)(nil), "tensorflow.GraphTransferInfo.GraphInputNodeInfo")
	proto.RegisterType((*GraphTransferInfo_GraphOutputNodeInfo)(nil), "tensorflow.GraphTransferInfo.GraphOutputNodeInfo")
	proto.RegisterEnum("tensorflow.GraphTransferInfo_Destination", GraphTransferInfo_Destination_name, GraphTransferInfo_Destination_value)
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/graph_transfer_info.proto", fileDescriptor6)
}

var fileDescriptor6 = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x95, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x71, 0xd2, 0x24, 0xf5, 0x71, 0x13, 0xca, 0xb4, 0x14, 0xcb, 0x08, 0x51, 0x8a, 0x80,
	0x72, 0x51, 0x0a, 0xed, 0x82, 0x35, 0xbd, 0x50, 0x2a, 0x44, 0x12, 0xb9, 0x15, 0x62, 0x67, 0x4d,
	0xed, 0x71, 0x6a, 0x4a, 0x66, 0x46, 0xf6, 0x44, 0x6d, 0xfa, 0x18, 0xbc, 0x1a, 0xcf, 0xc0, 0x3b,
	0xb0, 0x44, 0x73, 0xc6, 0xb2, 0x1d, 0xa5, 0x37, 0x16, 0xec, 0x66, 0x8e, 0xe7, 0xff, 0xce, 0x4d,
	0xbf, 0x0c, 0x5b, 0x8a, 0xf1, 0x4c, 0xa4, 0xf1, 0x0f, 0x71, 0xb6, 0x11, 0x8a, 0x94, 0x6d, 0xc4,
	0x29, 0x1d, 0xb1, 0x33, 0x91, 0x9e, 0x6e, 0x0c, 0x53, 0x2a, 0x4f, 0x02, 0x95, 0x52, 0x9e, 0xc5,
	0x2c, 0x0d, 0x12, 0x1e, 0x8b, 0xae, 0x4c, 0x85, 0x12, 0x04, 0x4a, 0x91, 0xf7, 0xec, 0x6a, 0x80,
	0x9a, 0x48, 0x96, 0x19, 0xc9, 0xda, 0x6f, 0x07, 0xee, 0xed, 0x6b, 0xe0, 0x51, 0xce, 0x3b, 0xe0,
	0xb1, 0x20, 0x3b, 0x60, 0x73, 0x11, 0x31, 0x64, 0xbb, 0xd6, 0x6a, 0x7d, 0xdd, 0xd9, 0x7c, 0xde,
	0x2d, 0x81, 0xdd, 0x19, 0x45, 0xb7, 0x27, 0x22, 0xa6, 0x0f, 0xfe, 0x3c, 0xcf, 0x4f, 0xe4, 0x10,
	0xee, 0x86, 0x82, 0x67, 0x2a, 0x28, 0x51, 0x35, 0x44, 0xbd, 0xbe, 0x1e, 0xb5, 0xa3, 0x45, 0x05,
	0xaf, 0x1d, 0x56, 0xaf, 0x1a, 0x9a, 0xe3, 0xe4, 0x58, 0x19, 0x68, 0xfd, 0x36, 0x50, 0x03, 0x90,
	0x63, 0x65, 0xa0, 0xbc, 0x7a, 0x25, 0x5f, 0x61, 0x11, 0xa1, 0x62, 0xac, 0x0a, 0xea, 0x1c, 0x52,
	0xdf, 0xdc, 0x4c, 0xed, 0xa3, 0x08, 0xb1, 0x1d, 0x3e, 0x75, 0x27, 0x21, 0xdc, 0x37, 0xcb, 0x32,
	0xd5, 0x96, 0x73, 0x68, 0x20, 0xfc, 0xed, 0xf5, 0x70, 0x8c, 0x60, 0x91, 0xc5, 0x30, 0xc8, 0x70,
	0x26, 0x46, 0x62, 0x58, 0x31, 0x49, 0xf2, 0xea, 0xcb, 0x2c, 0x4d, 0xcc, 0xf2, 0xee, 0x16, 0x59,
	0x4c, 0xcd, 0x45, 0x9a, 0xa5, 0xe1, 0x6c, 0x90, 0x7c, 0x06, 0x27, 0x62, 0x99, 0x4a, 0x38, 0x55,
	0x89, 0xe0, 0x6e, 0x6b, 0xd5, 0x5a, 0xef, 0x6c, 0xbe, 0xbc, 0x1e, 0xbe, 0x5b, 0x0a, 0xfc, 0xaa,
	0xda, 0xdb, 0x03, 0xbb, 0xd8, 0x08, 0x79, 0x00, 0x2d, 0x53, 0x74, 0xe4, 0x5a, 0xab, 0xd6, 0x7a,
	0xc3, 0x6f, 0xe2, 0x7a, 0x22, 0xf2, 0x18, 0x9c, 0xbc, 0x29, 0x29, 0x52, 0xe5, 0xd6, 0xf0, 0x23,
	0x98, 0xd0, 0x40, 0xa4, 0xca, 0xfb, 0x65, 0xc1, 0x7c, 0x51, 0x20, 0x81, 0x39, 0x4e, 0x47, 0x0c,
	0x19, 0xb6, 0x8f, 0xe7, 0x2a, 0xba, 0x36, 0x85, 0x7e, 0x08, 0xb6, 0xb6, 0x41, 0x80, 0x8a, 0x3a,
	0x2a, 0xe6, 0x75, 0xa0, 0xa7, 0x55, 0x1e, 0xd8, 0x99, 0x08, 0x03, 0x21, 0xb5, 0x6e, 0x0e, 0x75,
	0xad, 0x4c, 0x84, 0x7d, 0x79, 0x10, 0x91, 0x47, 0x00, 0x92, 0x46, 0x51, 0xc2, 0x87, 0xfa, 0x63,
	0x03, 0x3f, 0xda, 0x79, 0xc4, 0x94, 0x6c, 0x96, 0x1d, 0x8a, 0x31, 0x57, 0x6e, 0xd3, 0x94, 0x8c,
	0xa1, 0x1d, 0x1d, 0x21, 0x4f, 0x60, 0x21, 0xef, 0xc9, 0xbc, 0x68, 0xe1, 0x8b, 0xbc, 0x4f, 0x7c,
	0xe2, 0xfd, 0xb4, 0xa0, 0x3d, 0x65, 0x82, 0x7f, 0x6b, 0x6d, 0x19, 0x1a, 0xd9, 0x09, 0x95, 0x0c,
	0x8d, 0x51, 0xf7, 0xcd, 0x45, 0x23, 0x22, 0xaa, 0x28, 0xb6, 0xb3, 0xe0, 0xe3, 0x99, 0xbc, 0x82,
	0x46, 0xa4, 0x9b, 0xc6, 0x36, 0x3a, 0x9b, 0xcb, 0xd5, 0x65, 0xee, 0x52, 0x45, 0x8f, 0x26, 0x92,
	0xf9, 0xe6, 0x89, 0x27, 0xa1, 0x3d, 0xe5, 0xa1, 0xab, 0xb7, 0xf6, 0x11, 0xa0, 0xb4, 0x68, 0x6e,
	0xf9, 0x17, 0xb7, 0x74, 0xa7, 0x6f, 0x17, 0xce, 0xf4, 0xbe, 0x40, 0x67, 0xda, 0x5f, 0x57, 0xa7,
	0x5c, 0x83, 0xf6, 0x88, 0x9e, 0x07, 0xc7, 0x13, 0xc5, 0x82, 0x2c, 0xb9, 0x60, 0x98, 0xb5, 0xe1,
	0x3b, 0x23, 0x7a, 0xbe, 0x3d, 0x51, 0xec, 0x30, 0xb9, 0x60, 0xde, 0x77, 0x20, 0xb3, 0x8e, 0xba,
	0x74, 0xb2, 0xc5, 0x00, 0x6b, 0xd5, 0x01, 0x16, 0xc3, 0xaa, 0xdf, 0x3c, 0xac, 0x53, 0x58, 0xba,
	0xc4, 0x57, 0xff, 0x27, 0xd9, 0xda, 0x53, 0x70, 0x2a, 0x3e, 0x23, 0x2d, 0xa8, 0xf7, 0xfa, 0x83,
	0xc5, 0x3b, 0xc4, 0x81, 0xd6, 0xa7, 0xbd, 0x6f, 0x1f, 0xf6, 0xfb, 0xbd, 0x45, 0x6b, 0xfb, 0x3d,
	0xb8, 0x22, 0x1d, 0x56, 0x31, 0xc5, 0xff, 0x60, 0x7b, 0x65, 0x66, 0x21, 0x03, 0xfd, 0x6f, 0x18,
	0x58, 0x7f, 0x2c, 0xeb, 0xb8, 0x89, 0xff, 0x89, 0xad, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb4,
	0x1f, 0x2b, 0xd3, 0x91, 0x06, 0x00, 0x00,
}