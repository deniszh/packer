// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/compute/v1/snapshot.proto

package compute // import "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Snapshot_Status int32

const (
	Snapshot_STATUS_UNSPECIFIED Snapshot_Status = 0
	// Snapshot is being created.
	Snapshot_CREATING Snapshot_Status = 1
	// Snapshot is ready to use.
	Snapshot_READY Snapshot_Status = 2
	// Snapshot encountered a problem and cannot operate.
	Snapshot_ERROR Snapshot_Status = 3
	// Snapshot is being deleted.
	Snapshot_DELETING Snapshot_Status = 4
)

var Snapshot_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "CREATING",
	2: "READY",
	3: "ERROR",
	4: "DELETING",
}
var Snapshot_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"CREATING":           1,
	"READY":              2,
	"ERROR":              3,
	"DELETING":           4,
}

func (x Snapshot_Status) String() string {
	return proto.EnumName(Snapshot_Status_name, int32(x))
}
func (Snapshot_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_snapshot_451e295e1dfe1b62, []int{0, 0}
}

// A Snapshot resource. For more information, see [Snapshots](/docs/compute/concepts/snapshot).
type Snapshot struct {
	// ID of the snapshot.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the snapshot belongs to.
	FolderId  string               `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the snapshot. 1-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the snapshot. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `key:value` pairs. Maximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Size of the snapshot, specified in bytes.
	StorageSize int64 `protobuf:"varint,7,opt,name=storage_size,json=storageSize,proto3" json:"storage_size,omitempty"`
	// Size of the disk when the snapshot was created, specified in bytes.
	DiskSize int64 `protobuf:"varint,8,opt,name=disk_size,json=diskSize,proto3" json:"disk_size,omitempty"`
	// License IDs that indicate which licenses are attached to this resource.
	// License IDs are used to calculate additional charges for the use of the virtual machine.
	//
	// The correct license ID is generated by Yandex.Cloud. IDs are inherited by new resources created from this resource.
	//
	// If you know the license IDs, specify them when you create the image.
	// For example, if you create a disk image using a third-party utility and load it into Yandex Object Storage, the license IDs will be lost.
	// You can specify them in the [yandex.cloud.compute.v1.ImageService.Create] request.
	ProductIds []string `protobuf:"bytes,9,rep,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`
	// Current status of the snapshot.
	Status Snapshot_Status `protobuf:"varint,10,opt,name=status,proto3,enum=yandex.cloud.compute.v1.Snapshot_Status" json:"status,omitempty"`
	// ID of the source disk used to create this snapshot.
	SourceDiskId         string   `protobuf:"bytes,11,opt,name=source_disk_id,json=sourceDiskId,proto3" json:"source_disk_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}
func (*Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_snapshot_451e295e1dfe1b62, []int{0}
}
func (m *Snapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Snapshot.Unmarshal(m, b)
}
func (m *Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Snapshot.Marshal(b, m, deterministic)
}
func (dst *Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snapshot.Merge(dst, src)
}
func (m *Snapshot) XXX_Size() int {
	return xxx_messageInfo_Snapshot.Size(m)
}
func (m *Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_Snapshot proto.InternalMessageInfo

func (m *Snapshot) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Snapshot) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *Snapshot) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Snapshot) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Snapshot) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Snapshot) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Snapshot) GetStorageSize() int64 {
	if m != nil {
		return m.StorageSize
	}
	return 0
}

func (m *Snapshot) GetDiskSize() int64 {
	if m != nil {
		return m.DiskSize
	}
	return 0
}

func (m *Snapshot) GetProductIds() []string {
	if m != nil {
		return m.ProductIds
	}
	return nil
}

func (m *Snapshot) GetStatus() Snapshot_Status {
	if m != nil {
		return m.Status
	}
	return Snapshot_STATUS_UNSPECIFIED
}

func (m *Snapshot) GetSourceDiskId() string {
	if m != nil {
		return m.SourceDiskId
	}
	return ""
}

func init() {
	proto.RegisterType((*Snapshot)(nil), "yandex.cloud.compute.v1.Snapshot")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.compute.v1.Snapshot.LabelsEntry")
	proto.RegisterEnum("yandex.cloud.compute.v1.Snapshot_Status", Snapshot_Status_name, Snapshot_Status_value)
}

func init() {
	proto.RegisterFile("yandex/cloud/compute/v1/snapshot.proto", fileDescriptor_snapshot_451e295e1dfe1b62)
}

var fileDescriptor_snapshot_451e295e1dfe1b62 = []byte{
	// 484 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x8f, 0x9b, 0x3e,
	0x10, 0xc5, 0xff, 0x84, 0x24, 0xff, 0x30, 0x44, 0x11, 0xb2, 0xaa, 0x16, 0xa5, 0x87, 0xa5, 0xab,
	0xaa, 0xe2, 0x12, 0xd0, 0xa6, 0x97, 0x6e, 0x7b, 0x69, 0x9a, 0xd0, 0x0a, 0x69, 0xb5, 0xad, 0x4c,
	0xf6, 0xd0, 0x5e, 0x10, 0xc1, 0x5e, 0xd6, 0x0a, 0xc1, 0x08, 0x9b, 0xa8, 0xd9, 0x2f, 0xd8, 0xaf,
	0x55, 0x61, 0x1c, 0x69, 0x2f, 0xab, 0xde, 0x86, 0x37, 0xbf, 0xf1, 0xf3, 0xc3, 0x03, 0xef, 0x4e,
	0x59, 0x45, 0xe8, 0xef, 0x30, 0x2f, 0x79, 0x4b, 0xc2, 0x9c, 0x1f, 0xea, 0x56, 0xd2, 0xf0, 0x78,
	0x15, 0x8a, 0x2a, 0xab, 0xc5, 0x03, 0x97, 0x41, 0xdd, 0x70, 0xc9, 0xd1, 0xab, 0x9e, 0x0b, 0x14,
	0x17, 0x68, 0x2e, 0x38, 0x5e, 0xcd, 0x2f, 0x0a, 0xce, 0x8b, 0x92, 0x86, 0x0a, 0xdb, 0xb5, 0xf7,
	0xa1, 0x64, 0x07, 0x2a, 0x64, 0x76, 0xa8, 0xfb, 0xc9, 0xcb, 0x3f, 0x43, 0x98, 0x24, 0xfa, 0x30,
	0x34, 0x83, 0x01, 0x23, 0xae, 0xe1, 0x19, 0xbe, 0x85, 0x07, 0x8c, 0xa0, 0xd7, 0x60, 0xdd, 0xf3,
	0x92, 0xd0, 0x26, 0x65, 0xc4, 0x1d, 0x28, 0x79, 0xd2, 0x0b, 0x31, 0x41, 0xd7, 0x00, 0x79, 0x43,
	0x33, 0x49, 0x49, 0x9a, 0x49, 0xd7, 0xf4, 0x0c, 0xdf, 0x5e, 0xce, 0x83, 0xde, 0x2f, 0x38, 0xfb,
	0x05, 0xdb, 0xb3, 0x1f, 0xb6, 0x34, 0xbd, 0x92, 0x08, 0xc1, 0xb0, 0xca, 0x0e, 0xd4, 0x1d, 0xaa,
	0x23, 0x55, 0x8d, 0x3c, 0xb0, 0x09, 0x15, 0x79, 0xc3, 0x6a, 0xc9, 0x78, 0xe5, 0x8e, 0x54, 0xeb,
	0xa9, 0x84, 0x22, 0x18, 0x97, 0xd9, 0x8e, 0x96, 0xc2, 0x1d, 0x7b, 0xa6, 0x6f, 0x2f, 0x17, 0xc1,
	0x33, 0xa9, 0x83, 0x73, 0xa0, 0xe0, 0x46, 0xf1, 0x51, 0x25, 0x9b, 0x13, 0xd6, 0xc3, 0xe8, 0x0d,
	0x4c, 0x85, 0xe4, 0x4d, 0x56, 0xd0, 0x54, 0xb0, 0x47, 0xea, 0xfe, 0xef, 0x19, 0xbe, 0x89, 0x6d,
	0xad, 0x25, 0xec, 0x91, 0x76, 0xb9, 0x09, 0x13, 0xfb, 0xbe, 0x3f, 0x51, 0xfd, 0x49, 0x27, 0xa8,
	0xe6, 0x05, 0xd8, 0x75, 0xc3, 0x49, 0x9b, 0xcb, 0x94, 0x11, 0xe1, 0x5a, 0x9e, 0xe9, 0x5b, 0x18,
	0xb4, 0x14, 0x13, 0x81, 0x3e, 0xc3, 0x58, 0xc8, 0x4c, 0xb6, 0xc2, 0x05, 0xcf, 0xf0, 0x67, 0x4b,
	0xff, 0xdf, 0xf7, 0x4c, 0x14, 0x8f, 0xf5, 0x1c, 0x7a, 0x0b, 0x33, 0xc1, 0xdb, 0x26, 0xa7, 0xa9,
	0xba, 0x06, 0x23, 0xae, 0xad, 0x7e, 0xc7, 0xb4, 0x57, 0x37, 0x4c, 0xec, 0x63, 0x32, 0xbf, 0x06,
	0xfb, 0x49, 0x3e, 0xe4, 0x80, 0xb9, 0xa7, 0x27, 0xfd, 0x7a, 0x5d, 0x89, 0x5e, 0xc0, 0xe8, 0x98,
	0x95, 0x2d, 0xd5, 0x4f, 0xd7, 0x7f, 0x7c, 0x1c, 0x7c, 0x30, 0x2e, 0x31, 0x8c, 0x7b, 0x4b, 0xf4,
	0x12, 0x50, 0xb2, 0x5d, 0x6d, 0xef, 0x92, 0xf4, 0xee, 0x36, 0xf9, 0x11, 0xad, 0xe3, 0xaf, 0x71,
	0xb4, 0x71, 0xfe, 0x43, 0x53, 0x98, 0xac, 0x71, 0xb4, 0xda, 0xc6, 0xb7, 0xdf, 0x1c, 0x03, 0x59,
	0x30, 0xc2, 0xd1, 0x6a, 0xf3, 0xd3, 0x19, 0x74, 0x65, 0x84, 0xf1, 0x77, 0xec, 0x98, 0x1d, 0xb3,
	0x89, 0x6e, 0x22, 0xc5, 0x0c, 0xbf, 0x44, 0xbf, 0xd6, 0x05, 0x93, 0x0f, 0xed, 0xae, 0x4b, 0x18,
	0xf6, 0x91, 0x17, 0xfd, 0xe2, 0x16, 0x7c, 0x51, 0xd0, 0x4a, 0xed, 0x44, 0xf8, 0xcc, 0x46, 0x7f,
	0xd2, 0xe5, 0x6e, 0xac, 0xb0, 0xf7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xec, 0x41, 0x99,
	0xfb, 0x02, 0x00, 0x00,
}