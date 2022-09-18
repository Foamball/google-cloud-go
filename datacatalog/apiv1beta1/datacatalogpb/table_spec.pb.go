// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.5
// source: google/cloud/datacatalog/v1beta1/table_spec.proto

package datacatalogpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Table source type.
type TableSourceType int32

const (
	// Default unknown type.
	TableSourceType_TABLE_SOURCE_TYPE_UNSPECIFIED TableSourceType = 0
	// Table view.
	TableSourceType_BIGQUERY_VIEW TableSourceType = 2
	// BigQuery native table.
	TableSourceType_BIGQUERY_TABLE TableSourceType = 5
)

// Enum value maps for TableSourceType.
var (
	TableSourceType_name = map[int32]string{
		0: "TABLE_SOURCE_TYPE_UNSPECIFIED",
		2: "BIGQUERY_VIEW",
		5: "BIGQUERY_TABLE",
	}
	TableSourceType_value = map[string]int32{
		"TABLE_SOURCE_TYPE_UNSPECIFIED": 0,
		"BIGQUERY_VIEW":                 2,
		"BIGQUERY_TABLE":                5,
	}
)

func (x TableSourceType) Enum() *TableSourceType {
	p := new(TableSourceType)
	*p = x
	return p
}

func (x TableSourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TableSourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_datacatalog_v1beta1_table_spec_proto_enumTypes[0].Descriptor()
}

func (TableSourceType) Type() protoreflect.EnumType {
	return &file_google_cloud_datacatalog_v1beta1_table_spec_proto_enumTypes[0]
}

func (x TableSourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TableSourceType.Descriptor instead.
func (TableSourceType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1beta1_table_spec_proto_rawDescGZIP(), []int{0}
}

// Describes a BigQuery table.
type BigQueryTableSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The table source type.
	TableSourceType TableSourceType `protobuf:"varint,1,opt,name=table_source_type,json=tableSourceType,proto3,enum=google.cloud.datacatalog.v1beta1.TableSourceType" json:"table_source_type,omitempty"`
	// Output only.
	//
	// Types that are assignable to TypeSpec:
	//
	//	*BigQueryTableSpec_ViewSpec
	//	*BigQueryTableSpec_TableSpec
	TypeSpec isBigQueryTableSpec_TypeSpec `protobuf_oneof:"type_spec"`
}

func (x *BigQueryTableSpec) Reset() {
	*x = BigQueryTableSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1beta1_table_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigQueryTableSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigQueryTableSpec) ProtoMessage() {}

func (x *BigQueryTableSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1beta1_table_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigQueryTableSpec.ProtoReflect.Descriptor instead.
func (*BigQueryTableSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1beta1_table_spec_proto_rawDescGZIP(), []int{0}
}

func (x *BigQueryTableSpec) GetTableSourceType() TableSourceType {
	if x != nil {
		return x.TableSourceType
	}
	return TableSourceType_TABLE_SOURCE_TYPE_UNSPECIFIED
}

func (m *BigQueryTableSpec) GetTypeSpec() isBigQueryTableSpec_TypeSpec {
	if m != nil {
		return m.TypeSpec
	}
	return nil
}

func (x *BigQueryTableSpec) GetViewSpec() *ViewSpec {
	if x, ok := x.GetTypeSpec().(*BigQueryTableSpec_ViewSpec); ok {
		return x.ViewSpec
	}
	return nil
}

func (x *BigQueryTableSpec) GetTableSpec() *TableSpec {
	if x, ok := x.GetTypeSpec().(*BigQueryTableSpec_TableSpec); ok {
		return x.TableSpec
	}
	return nil
}

type isBigQueryTableSpec_TypeSpec interface {
	isBigQueryTableSpec_TypeSpec()
}

type BigQueryTableSpec_ViewSpec struct {
	// Table view specification. This field should only be populated if
	// `table_source_type` is `BIGQUERY_VIEW`.
	ViewSpec *ViewSpec `protobuf:"bytes,2,opt,name=view_spec,json=viewSpec,proto3,oneof"`
}

type BigQueryTableSpec_TableSpec struct {
	// Spec of a BigQuery table. This field should only be populated if
	// `table_source_type` is `BIGQUERY_TABLE`.
	TableSpec *TableSpec `protobuf:"bytes,3,opt,name=table_spec,json=tableSpec,proto3,oneof"`
}

func (*BigQueryTableSpec_ViewSpec) isBigQueryTableSpec_TypeSpec() {}

func (*BigQueryTableSpec_TableSpec) isBigQueryTableSpec_TypeSpec() {}

// Table view specification.
type ViewSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The query that defines the table view.
	ViewQuery string `protobuf:"bytes,1,opt,name=view_query,json=viewQuery,proto3" json:"view_query,omitempty"`
}

func (x *ViewSpec) Reset() {
	*x = ViewSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1beta1_table_spec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ViewSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ViewSpec) ProtoMessage() {}

func (x *ViewSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1beta1_table_spec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ViewSpec.ProtoReflect.Descriptor instead.
func (*ViewSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1beta1_table_spec_proto_rawDescGZIP(), []int{1}
}

func (x *ViewSpec) GetViewQuery() string {
	if x != nil {
		return x.ViewQuery
	}
	return ""
}

// Normal BigQuery table spec.
type TableSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. If the table is a dated shard, i.e., with name pattern `[prefix]YYYYMMDD`,
	// `grouped_entry` is the Data Catalog resource name of the date sharded
	// grouped entry, for example,
	// `projects/{project_id}/locations/{location}/entrygroups/{entry_group_id}/entries/{entry_id}`.
	// Otherwise, `grouped_entry` is empty.
	GroupedEntry string `protobuf:"bytes,1,opt,name=grouped_entry,json=groupedEntry,proto3" json:"grouped_entry,omitempty"`
}

func (x *TableSpec) Reset() {
	*x = TableSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1beta1_table_spec_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableSpec) ProtoMessage() {}

func (x *TableSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1beta1_table_spec_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableSpec.ProtoReflect.Descriptor instead.
func (*TableSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1beta1_table_spec_proto_rawDescGZIP(), []int{2}
}

func (x *TableSpec) GetGroupedEntry() string {
	if x != nil {
		return x.GroupedEntry
	}
	return ""
}

// Spec for a group of BigQuery tables with name pattern `[prefix]YYYYMMDD`.
// Context:
// https://cloud.google.com/bigquery/docs/partitioned-tables#partitioning_versus_sharding
type BigQueryDateShardedSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. The Data Catalog resource name of the dataset entry the current table
	// belongs to, for example,
	// `projects/{project_id}/locations/{location}/entrygroups/{entry_group_id}/entries/{entry_id}`.
	Dataset string `protobuf:"bytes,1,opt,name=dataset,proto3" json:"dataset,omitempty"`
	// Output only. The table name prefix of the shards. The name of any given shard is
	// `[table_prefix]YYYYMMDD`, for example, for shard `MyTable20180101`, the
	// `table_prefix` is `MyTable`.
	TablePrefix string `protobuf:"bytes,2,opt,name=table_prefix,json=tablePrefix,proto3" json:"table_prefix,omitempty"`
	// Output only. Total number of shards.
	ShardCount int64 `protobuf:"varint,3,opt,name=shard_count,json=shardCount,proto3" json:"shard_count,omitempty"`
}

func (x *BigQueryDateShardedSpec) Reset() {
	*x = BigQueryDateShardedSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1beta1_table_spec_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigQueryDateShardedSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigQueryDateShardedSpec) ProtoMessage() {}

func (x *BigQueryDateShardedSpec) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1beta1_table_spec_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigQueryDateShardedSpec.ProtoReflect.Descriptor instead.
func (*BigQueryDateShardedSpec) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1beta1_table_spec_proto_rawDescGZIP(), []int{3}
}

func (x *BigQueryDateShardedSpec) GetDataset() string {
	if x != nil {
		return x.Dataset
	}
	return ""
}

func (x *BigQueryDateShardedSpec) GetTablePrefix() string {
	if x != nil {
		return x.TablePrefix
	}
	return ""
}

func (x *BigQueryDateShardedSpec) GetShardCount() int64 {
	if x != nil {
		return x.ShardCount
	}
	return 0
}

var File_google_cloud_datacatalog_v1beta1_table_spec_proto protoreflect.FileDescriptor

var file_google_cloud_datacatalog_v1beta1_table_spec_proto_rawDesc = []byte{
	0x0a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x9d, 0x02, 0x0a, 0x11, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x62, 0x0a, 0x11, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0f, 0x74, 0x61, 0x62, 0x6c,
	0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x49, 0x0a, 0x09, 0x76,
	0x69, 0x65, 0x77, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x08, 0x76, 0x69,
	0x65, 0x77, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4c, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61,
	0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x48, 0x00, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x53, 0x70, 0x65, 0x63, 0x42, 0x0b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x22, 0x2e, 0x0a, 0x08, 0x56, 0x69, 0x65, 0x77, 0x53, 0x70, 0x65, 0x63, 0x12, 0x22, 0x0a,
	0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x22, 0x5a, 0x0a, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x4d,
	0x0a, 0x0d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xe0, 0x41, 0x03, 0xfa, 0x41, 0x22, 0x0a, 0x20, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0xab, 0x01,
	0x0a, 0x17, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x61, 0x72, 0x64, 0x65, 0x64, 0x53, 0x70, 0x65, 0x63, 0x12, 0x42, 0x0a, 0x07, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xe0, 0x41, 0x03, 0xfa,
	0x41, 0x22, 0x0a, 0x20, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x12, 0x26, 0x0a,
	0x0c, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x50,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x24, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x72, 0x64, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52,
	0x0a, 0x73, 0x68, 0x61, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x2a, 0x5b, 0x0a, 0x0f, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21,
	0x0a, 0x1d, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x49, 0x47, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x56, 0x49,
	0x45, 0x57, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x42, 0x49, 0x47, 0x51, 0x55, 0x45, 0x52, 0x59,
	0x5f, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x05, 0x42, 0xe4, 0x01, 0x0a, 0x24, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0xf8, 0x01, 0x01, 0xaa, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x56,
	0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x20, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c,
	0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x43,
	0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_datacatalog_v1beta1_table_spec_proto_rawDescOnce sync.Once
	file_google_cloud_datacatalog_v1beta1_table_spec_proto_rawDescData = file_google_cloud_datacatalog_v1beta1_table_spec_proto_rawDesc
)

func file_google_cloud_datacatalog_v1beta1_table_spec_proto_rawDescGZIP() []byte {
	file_google_cloud_datacatalog_v1beta1_table_spec_proto_rawDescOnce.Do(func() {
		file_google_cloud_datacatalog_v1beta1_table_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_datacatalog_v1beta1_table_spec_proto_rawDescData)
	})
	return file_google_cloud_datacatalog_v1beta1_table_spec_proto_rawDescData
}

var file_google_cloud_datacatalog_v1beta1_table_spec_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_datacatalog_v1beta1_table_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_google_cloud_datacatalog_v1beta1_table_spec_proto_goTypes = []interface{}{
	(TableSourceType)(0),            // 0: google.cloud.datacatalog.v1beta1.TableSourceType
	(*BigQueryTableSpec)(nil),       // 1: google.cloud.datacatalog.v1beta1.BigQueryTableSpec
	(*ViewSpec)(nil),                // 2: google.cloud.datacatalog.v1beta1.ViewSpec
	(*TableSpec)(nil),               // 3: google.cloud.datacatalog.v1beta1.TableSpec
	(*BigQueryDateShardedSpec)(nil), // 4: google.cloud.datacatalog.v1beta1.BigQueryDateShardedSpec
}
var file_google_cloud_datacatalog_v1beta1_table_spec_proto_depIdxs = []int32{
	0, // 0: google.cloud.datacatalog.v1beta1.BigQueryTableSpec.table_source_type:type_name -> google.cloud.datacatalog.v1beta1.TableSourceType
	2, // 1: google.cloud.datacatalog.v1beta1.BigQueryTableSpec.view_spec:type_name -> google.cloud.datacatalog.v1beta1.ViewSpec
	3, // 2: google.cloud.datacatalog.v1beta1.BigQueryTableSpec.table_spec:type_name -> google.cloud.datacatalog.v1beta1.TableSpec
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_cloud_datacatalog_v1beta1_table_spec_proto_init() }
func file_google_cloud_datacatalog_v1beta1_table_spec_proto_init() {
	if File_google_cloud_datacatalog_v1beta1_table_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_datacatalog_v1beta1_table_spec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigQueryTableSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_datacatalog_v1beta1_table_spec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ViewSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_datacatalog_v1beta1_table_spec_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_cloud_datacatalog_v1beta1_table_spec_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigQueryDateShardedSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_google_cloud_datacatalog_v1beta1_table_spec_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BigQueryTableSpec_ViewSpec)(nil),
		(*BigQueryTableSpec_TableSpec)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_datacatalog_v1beta1_table_spec_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_datacatalog_v1beta1_table_spec_proto_goTypes,
		DependencyIndexes: file_google_cloud_datacatalog_v1beta1_table_spec_proto_depIdxs,
		EnumInfos:         file_google_cloud_datacatalog_v1beta1_table_spec_proto_enumTypes,
		MessageInfos:      file_google_cloud_datacatalog_v1beta1_table_spec_proto_msgTypes,
	}.Build()
	File_google_cloud_datacatalog_v1beta1_table_spec_proto = out.File
	file_google_cloud_datacatalog_v1beta1_table_spec_proto_rawDesc = nil
	file_google_cloud_datacatalog_v1beta1_table_spec_proto_goTypes = nil
	file_google_cloud_datacatalog_v1beta1_table_spec_proto_depIdxs = nil
}
