// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: errcode.proto

package app_error

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RetCode int32

const (
	RetCode_SUCCEED RetCode = 0 // 通用 error
	// User error
	RetCode_ERR_USER_ALREADY_EXISTS RetCode = 21000 // 用户已存在
	RetCode_ERR_USER_NOT_EXIST      RetCode = 21001 // 用户不存在
	RetCode_ERR_INVALID_USER        RetCode = 21002 // 无效用户
	RetCode_ERR_AUTHSERVICE_ERROR   RetCode = 21003 // 鉴权服务错误
	RetCode_ERR_MISSING_USER_ID     RetCode = 21004 // 缺失用户id
	// Qrcode error
	RetCode_ERR_INVALID_QRCODE  RetCode = 22001 // 无效二维码
	RetCode_ERR_REMAINING_USAGE RetCode = 22002 // 剩余使用次数不足
	RetCode_ERR_NO_HANDLER      RetCode = 22003 // 没有对应的handler
	// Report error
	RetCode_ERR_REPORT_FILE_PATH_INVALID RetCode = 23001 // 上报文件路径无效
	RetCode_ERR_ZIP_HANDLER_FAILED       RetCode = 23002 // 压缩上报文件失败
	RetCode_ERR_SIGN_NOT_MATCH           RetCode = 23003 // 查询上报文件鉴权sign无效
	// Asset error
	RetCode_ERR_GET_LONG_TERM_TOKEN_FAILED   RetCode = 24001 // 对公账户获取长期token失败
	RetCode_ERR_GET_DA_USER_INFO_FAILED      RetCode = 24002 // 获取da用户信息失败
	RetCode_ERR_GET_CONTRACT_INFO_FAILED     RetCode = 24003 // 获取合约信息失败
	RetCode_ERR_UPLOAD_COS_PIC_FAILED        RetCode = 24004 // 上传图片到cos桶失败
	RetCode_ERR_GET_USER_DID_FAILED          RetCode = 24005 // 获取用户did失败
	RetCode_ERR_GET_DA_USER_ID_FAILED        RetCode = 24006 // 获取da的用户id失败
	RetCode_ERR_COMBINED_ISSUE_FAILED        RetCode = 24007 // 组合发行失败
	RetCode_ERR_CALL_RANKS_CONTESTANT_FAILED RetCode = 24008 // 请求排行榜服务记录发行的图片失败
	// Rabbit & 1024 error
	RetCode_ERR_EGG_LOCK_USER_FAILED             RetCode = 50001
	RetCode_ERR_EGG_LOCK_EGG_FAILED              RetCode = 50002
	RetCode_ERR_UPDATE_EGG_COLLECTED_INFO_FAILED RetCode = 50003
	RetCode_ERR_GET_EGG_INFO_FAILED              RetCode = 50004
	RetCode_ERR_NO_ROWS_UPDATED                  RetCode = 50006
	RetCode_ERR_DATA_NOT_FOUND                   RetCode = 50007
	RetCode_ERR_HIT_LIMIT                        RetCode = 50008
	RetCode_ERR_NO_COUPON_LEFT                   RetCode = 50009
	RetCode_ERR_HAS_BEEN_REDEEMED                RetCode = 50010
	RetCode_ERR_NOT_OWNER                        RetCode = 50011
	RetCode_ERR_NOT_ENOUGH_MEDALS                RetCode = 50012
	RetCode_ERR_EMPTY_GIFT                       RetCode = 50013
	RetCode_ERR_ACCESS_DENIED                    RetCode = 50014
	RetCode_ERR_GET_ASSET_DETAIL_FAILED          RetCode = 50016
	RetCode_ERR_INVALID_TIME_PERIOD              RetCode = 50017
	RetCode_ERR_NOT_SUPPORTED_OPERATION          RetCode = 50018
	RetCode_ERR_GET_COMBINED_ISSUE_FAILED        RetCode = 50019
	RetCode_ERR_PRE_UPLOAD_COS_FILES_FAILED      RetCode = 50020
	RetCode_ERR_LOCKED_COUPON_FOUND              RetCode = 50022
	RetCode_ERR_INVALID_IMAGE_URL                RetCode = 50023
	RetCode_ERR_DA_ERROR                         RetCode = 50024
	RetCode_ERR_TOO_MANY_REQUESTS                RetCode = 50025
	// New Year Rabbit
	RetCode_ERR_EMPTY_IMAGE_INFO_FROM_APP_CONFIG RetCode = 51000
	RetCode_ERR_GET_RANDOM_UPGRADE_PARTS_FAILED  RetCode = 51001
)

// Enum value maps for RetCode.
var (
	RetCode_name = map[int32]string{
		0:     "SUCCEED",
		21000: "ERR_USER_ALREADY_EXISTS",
		21001: "ERR_USER_NOT_EXIST",
		21002: "ERR_INVALID_USER",
		21003: "ERR_AUTHSERVICE_ERROR",
		21004: "ERR_MISSING_USER_ID",
		22001: "ERR_INVALID_QRCODE",
		22002: "ERR_REMAINING_USAGE",
		22003: "ERR_NO_HANDLER",
		23001: "ERR_REPORT_FILE_PATH_INVALID",
		23002: "ERR_ZIP_HANDLER_FAILED",
		23003: "ERR_SIGN_NOT_MATCH",
		24001: "ERR_GET_LONG_TERM_TOKEN_FAILED",
		24002: "ERR_GET_DA_USER_INFO_FAILED",
		24003: "ERR_GET_CONTRACT_INFO_FAILED",
		24004: "ERR_UPLOAD_COS_PIC_FAILED",
		24005: "ERR_GET_USER_DID_FAILED",
		24006: "ERR_GET_DA_USER_ID_FAILED",
		24007: "ERR_COMBINED_ISSUE_FAILED",
		24008: "ERR_CALL_RANKS_CONTESTANT_FAILED",
		50001: "ERR_EGG_LOCK_USER_FAILED",
		50002: "ERR_EGG_LOCK_EGG_FAILED",
		50003: "ERR_UPDATE_EGG_COLLECTED_INFO_FAILED",
		50004: "ERR_GET_EGG_INFO_FAILED",
		50006: "ERR_NO_ROWS_UPDATED",
		50007: "ERR_DATA_NOT_FOUND",
		50008: "ERR_HIT_LIMIT",
		50009: "ERR_NO_COUPON_LEFT",
		50010: "ERR_HAS_BEEN_REDEEMED",
		50011: "ERR_NOT_OWNER",
		50012: "ERR_NOT_ENOUGH_MEDALS",
		50013: "ERR_EMPTY_GIFT",
		50014: "ERR_ACCESS_DENIED",
		50016: "ERR_GET_ASSET_DETAIL_FAILED",
		50017: "ERR_INVALID_TIME_PERIOD",
		50018: "ERR_NOT_SUPPORTED_OPERATION",
		50019: "ERR_GET_COMBINED_ISSUE_FAILED",
		50020: "ERR_PRE_UPLOAD_COS_FILES_FAILED",
		50022: "ERR_LOCKED_COUPON_FOUND",
		50023: "ERR_INVALID_IMAGE_URL",
		50024: "ERR_DA_ERROR",
		50025: "ERR_TOO_MANY_REQUESTS",
		51000: "ERR_EMPTY_IMAGE_INFO_FROM_APP_CONFIG",
		51001: "ERR_GET_RANDOM_UPGRADE_PARTS_FAILED",
	}
	RetCode_value = map[string]int32{
		"SUCCEED":                              0,
		"ERR_USER_ALREADY_EXISTS":              21000,
		"ERR_USER_NOT_EXIST":                   21001,
		"ERR_INVALID_USER":                     21002,
		"ERR_AUTHSERVICE_ERROR":                21003,
		"ERR_MISSING_USER_ID":                  21004,
		"ERR_INVALID_QRCODE":                   22001,
		"ERR_REMAINING_USAGE":                  22002,
		"ERR_NO_HANDLER":                       22003,
		"ERR_REPORT_FILE_PATH_INVALID":         23001,
		"ERR_ZIP_HANDLER_FAILED":               23002,
		"ERR_SIGN_NOT_MATCH":                   23003,
		"ERR_GET_LONG_TERM_TOKEN_FAILED":       24001,
		"ERR_GET_DA_USER_INFO_FAILED":          24002,
		"ERR_GET_CONTRACT_INFO_FAILED":         24003,
		"ERR_UPLOAD_COS_PIC_FAILED":            24004,
		"ERR_GET_USER_DID_FAILED":              24005,
		"ERR_GET_DA_USER_ID_FAILED":            24006,
		"ERR_COMBINED_ISSUE_FAILED":            24007,
		"ERR_CALL_RANKS_CONTESTANT_FAILED":     24008,
		"ERR_EGG_LOCK_USER_FAILED":             50001,
		"ERR_EGG_LOCK_EGG_FAILED":              50002,
		"ERR_UPDATE_EGG_COLLECTED_INFO_FAILED": 50003,
		"ERR_GET_EGG_INFO_FAILED":              50004,
		"ERR_NO_ROWS_UPDATED":                  50006,
		"ERR_DATA_NOT_FOUND":                   50007,
		"ERR_HIT_LIMIT":                        50008,
		"ERR_NO_COUPON_LEFT":                   50009,
		"ERR_HAS_BEEN_REDEEMED":                50010,
		"ERR_NOT_OWNER":                        50011,
		"ERR_NOT_ENOUGH_MEDALS":                50012,
		"ERR_EMPTY_GIFT":                       50013,
		"ERR_ACCESS_DENIED":                    50014,
		"ERR_GET_ASSET_DETAIL_FAILED":          50016,
		"ERR_INVALID_TIME_PERIOD":              50017,
		"ERR_NOT_SUPPORTED_OPERATION":          50018,
		"ERR_GET_COMBINED_ISSUE_FAILED":        50019,
		"ERR_PRE_UPLOAD_COS_FILES_FAILED":      50020,
		"ERR_LOCKED_COUPON_FOUND":              50022,
		"ERR_INVALID_IMAGE_URL":                50023,
		"ERR_DA_ERROR":                         50024,
		"ERR_TOO_MANY_REQUESTS":                50025,
		"ERR_EMPTY_IMAGE_INFO_FROM_APP_CONFIG": 51000,
		"ERR_GET_RANDOM_UPGRADE_PARTS_FAILED":  51001,
	}
)

func (x RetCode) Enum() *RetCode {
	p := new(RetCode)
	*p = x
	return p
}

func (x RetCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RetCode) Descriptor() protoreflect.EnumDescriptor {
	return file_errcode_proto_enumTypes[0].Descriptor()
}

func (RetCode) Type() protoreflect.EnumType {
	return &file_errcode_proto_enumTypes[0]
}

func (x RetCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RetCode.Descriptor instead.
func (RetCode) EnumDescriptor() ([]byte, []int) {
	return file_errcode_proto_rawDescGZIP(), []int{0}
}

var File_errcode_proto protoreflect.FileDescriptor

var file_errcode_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x72, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x61, 0x70, 0x70, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0xbe, 0x0a, 0x0a, 0x07, 0x52,
	0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x88,
	0xa4, 0x01, 0x12, 0x18, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x10, 0x89, 0xa4, 0x01, 0x12, 0x16, 0x0a, 0x10,
	0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x55, 0x53, 0x45, 0x52,
	0x10, 0x8a, 0xa4, 0x01, 0x12, 0x1b, 0x0a, 0x15, 0x45, 0x52, 0x52, 0x5f, 0x41, 0x55, 0x54, 0x48,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x8b, 0xa4,
	0x01, 0x12, 0x19, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47,
	0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x10, 0x8c, 0xa4, 0x01, 0x12, 0x18, 0x0a, 0x12,
	0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x51, 0x52, 0x43, 0x4f,
	0x44, 0x45, 0x10, 0xf1, 0xab, 0x01, 0x12, 0x19, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x5f, 0x52, 0x45,
	0x4d, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x53, 0x41, 0x47, 0x45, 0x10, 0xf2, 0xab,
	0x01, 0x12, 0x14, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x5f, 0x4e, 0x4f, 0x5f, 0x48, 0x41, 0x4e, 0x44,
	0x4c, 0x45, 0x52, 0x10, 0xf3, 0xab, 0x01, 0x12, 0x22, 0x0a, 0x1c, 0x45, 0x52, 0x52, 0x5f, 0x52,
	0x45, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x50, 0x41, 0x54, 0x48, 0x5f,
	0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0xd9, 0xb3, 0x01, 0x12, 0x1c, 0x0a, 0x16, 0x45,
	0x52, 0x52, 0x5f, 0x5a, 0x49, 0x50, 0x5f, 0x48, 0x41, 0x4e, 0x44, 0x4c, 0x45, 0x52, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xda, 0xb3, 0x01, 0x12, 0x18, 0x0a, 0x12, 0x45, 0x52, 0x52,
	0x5f, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10,
	0xdb, 0xb3, 0x01, 0x12, 0x24, 0x0a, 0x1e, 0x45, 0x52, 0x52, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x4c,
	0x4f, 0x4e, 0x47, 0x5f, 0x54, 0x45, 0x52, 0x4d, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xc1, 0xbb, 0x01, 0x12, 0x21, 0x0a, 0x1b, 0x45, 0x52, 0x52,
	0x5f, 0x47, 0x45, 0x54, 0x5f, 0x44, 0x41, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46,
	0x4f, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xc2, 0xbb, 0x01, 0x12, 0x22, 0x0a, 0x1c,
	0x45, 0x52, 0x52, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x41, 0x43, 0x54,
	0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xc3, 0xbb, 0x01,
	0x12, 0x1f, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x43,
	0x4f, 0x53, 0x5f, 0x50, 0x49, 0x43, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xc4, 0xbb,
	0x01, 0x12, 0x1d, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x55, 0x53, 0x45,
	0x52, 0x5f, 0x44, 0x49, 0x44, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xc5, 0xbb, 0x01,
	0x12, 0x1f, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x44, 0x41, 0x5f, 0x55,
	0x53, 0x45, 0x52, 0x5f, 0x49, 0x44, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xc6, 0xbb,
	0x01, 0x12, 0x1f, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x4d, 0x42, 0x49, 0x4e, 0x45,
	0x44, 0x5f, 0x49, 0x53, 0x53, 0x55, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xc7,
	0xbb, 0x01, 0x12, 0x26, 0x0a, 0x20, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x52,
	0x41, 0x4e, 0x4b, 0x53, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x53, 0x54, 0x41, 0x4e, 0x54, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xc8, 0xbb, 0x01, 0x12, 0x1e, 0x0a, 0x18, 0x45, 0x52,
	0x52, 0x5f, 0x45, 0x47, 0x47, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xd1, 0x86, 0x03, 0x12, 0x1d, 0x0a, 0x17, 0x45, 0x52,
	0x52, 0x5f, 0x45, 0x47, 0x47, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x45, 0x47, 0x47, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xd2, 0x86, 0x03, 0x12, 0x2a, 0x0a, 0x24, 0x45, 0x52, 0x52,
	0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x45, 0x47, 0x47, 0x5f, 0x43, 0x4f, 0x4c, 0x4c,
	0x45, 0x43, 0x54, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0xd3, 0x86, 0x03, 0x12, 0x1d, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x5f, 0x47, 0x45, 0x54,
	0x5f, 0x45, 0x47, 0x47, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44,
	0x10, 0xd4, 0x86, 0x03, 0x12, 0x19, 0x0a, 0x13, 0x45, 0x52, 0x52, 0x5f, 0x4e, 0x4f, 0x5f, 0x52,
	0x4f, 0x57, 0x53, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0xd6, 0x86, 0x03, 0x12,
	0x18, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xd7, 0x86, 0x03, 0x12, 0x13, 0x0a, 0x0d, 0x45, 0x52, 0x52,
	0x5f, 0x48, 0x49, 0x54, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0xd8, 0x86, 0x03, 0x12, 0x18,
	0x0a, 0x12, 0x45, 0x52, 0x52, 0x5f, 0x4e, 0x4f, 0x5f, 0x43, 0x4f, 0x55, 0x50, 0x4f, 0x4e, 0x5f,
	0x4c, 0x45, 0x46, 0x54, 0x10, 0xd9, 0x86, 0x03, 0x12, 0x1b, 0x0a, 0x15, 0x45, 0x52, 0x52, 0x5f,
	0x48, 0x41, 0x53, 0x5f, 0x42, 0x45, 0x45, 0x4e, 0x5f, 0x52, 0x45, 0x44, 0x45, 0x45, 0x4d, 0x45,
	0x44, 0x10, 0xda, 0x86, 0x03, 0x12, 0x13, 0x0a, 0x0d, 0x45, 0x52, 0x52, 0x5f, 0x4e, 0x4f, 0x54,
	0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10, 0xdb, 0x86, 0x03, 0x12, 0x1b, 0x0a, 0x15, 0x45, 0x52,
	0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x4e, 0x4f, 0x55, 0x47, 0x48, 0x5f, 0x4d, 0x45, 0x44,
	0x41, 0x4c, 0x53, 0x10, 0xdc, 0x86, 0x03, 0x12, 0x14, 0x0a, 0x0e, 0x45, 0x52, 0x52, 0x5f, 0x45,
	0x4d, 0x50, 0x54, 0x59, 0x5f, 0x47, 0x49, 0x46, 0x54, 0x10, 0xdd, 0x86, 0x03, 0x12, 0x17, 0x0a,
	0x11, 0x45, 0x52, 0x52, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x45, 0x4e, 0x49,
	0x45, 0x44, 0x10, 0xde, 0x86, 0x03, 0x12, 0x21, 0x0a, 0x1b, 0x45, 0x52, 0x52, 0x5f, 0x47, 0x45,
	0x54, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x44, 0x45, 0x54, 0x41, 0x49, 0x4c, 0x5f, 0x46,
	0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xe0, 0x86, 0x03, 0x12, 0x1d, 0x0a, 0x17, 0x45, 0x52, 0x52,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x50, 0x45,
	0x52, 0x49, 0x4f, 0x44, 0x10, 0xe1, 0x86, 0x03, 0x12, 0x21, 0x0a, 0x1b, 0x45, 0x52, 0x52, 0x5f,
	0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f, 0x4f, 0x50,
	0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xe2, 0x86, 0x03, 0x12, 0x23, 0x0a, 0x1d, 0x45,
	0x52, 0x52, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x43, 0x4f, 0x4d, 0x42, 0x49, 0x4e, 0x45, 0x44, 0x5f,
	0x49, 0x53, 0x53, 0x55, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xe3, 0x86, 0x03,
	0x12, 0x25, 0x0a, 0x1f, 0x45, 0x52, 0x52, 0x5f, 0x50, 0x52, 0x45, 0x5f, 0x55, 0x50, 0x4c, 0x4f,
	0x41, 0x44, 0x5f, 0x43, 0x4f, 0x53, 0x5f, 0x46, 0x49, 0x4c, 0x45, 0x53, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0xe4, 0x86, 0x03, 0x12, 0x1d, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x5f, 0x4c,
	0x4f, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x43, 0x4f, 0x55, 0x50, 0x4f, 0x4e, 0x5f, 0x46, 0x4f, 0x55,
	0x4e, 0x44, 0x10, 0xe6, 0x86, 0x03, 0x12, 0x1b, 0x0a, 0x15, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x55, 0x52, 0x4c, 0x10,
	0xe7, 0x86, 0x03, 0x12, 0x12, 0x0a, 0x0c, 0x45, 0x52, 0x52, 0x5f, 0x44, 0x41, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0xe8, 0x86, 0x03, 0x12, 0x1b, 0x0a, 0x15, 0x45, 0x52, 0x52, 0x5f, 0x54,
	0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x53,
	0x10, 0xe9, 0x86, 0x03, 0x12, 0x2a, 0x0a, 0x24, 0x45, 0x52, 0x52, 0x5f, 0x45, 0x4d, 0x50, 0x54,
	0x59, 0x5f, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x46, 0x52, 0x4f,
	0x4d, 0x5f, 0x41, 0x50, 0x50, 0x5f, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0xb8, 0x8e, 0x03,
	0x12, 0x29, 0x0a, 0x23, 0x45, 0x52, 0x52, 0x5f, 0x47, 0x45, 0x54, 0x5f, 0x52, 0x41, 0x4e, 0x44,
	0x4f, 0x4d, 0x5f, 0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x45, 0x5f, 0x50, 0x41, 0x52, 0x54, 0x53,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xb9, 0x8e, 0x03, 0x42, 0x49, 0x5a, 0x47, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x67, 0x61, 0x6c, 0x61,
	0x78, 0x79, 0x2f, 0x77, 0x65, 0x67, 0x61, 0x6c, 0x61, 0x78, 0x79, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x77, 0x65, 0x67, 0x61, 0x6c, 0x61, 0x78, 0x79, 0x2d, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x3b, 0x61, 0x70, 0x70,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errcode_proto_rawDescOnce sync.Once
	file_errcode_proto_rawDescData = file_errcode_proto_rawDesc
)

func file_errcode_proto_rawDescGZIP() []byte {
	file_errcode_proto_rawDescOnce.Do(func() {
		file_errcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_errcode_proto_rawDescData)
	})
	return file_errcode_proto_rawDescData
}

var file_errcode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errcode_proto_goTypes = []interface{}{
	(RetCode)(0), // 0: app_error.RetCode
}
var file_errcode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errcode_proto_init() }
func file_errcode_proto_init() {
	if File_errcode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errcode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errcode_proto_goTypes,
		DependencyIndexes: file_errcode_proto_depIdxs,
		EnumInfos:         file_errcode_proto_enumTypes,
	}.Build()
	File_errcode_proto = out.File
	file_errcode_proto_rawDesc = nil
	file_errcode_proto_goTypes = nil
	file_errcode_proto_depIdxs = nil
}