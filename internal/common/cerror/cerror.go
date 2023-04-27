// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v1.1.0-EE
// Author: meta-egg
// Generated at: 2023-04-27 19:17

package cerror

import (
	"fmt"
	"net/http"
	"strings"

	api "meta-egg-layout/api/meta_egg_layout"
)

type CustomError struct {
	HttpStatus int         `json:"http_status"` // http 状态
	Code       api.ErrCode `json:"code"`        // code 错误码
	Message    string      `json:"msg"`         // msg 消息
	Detail     string      `json:"detail"`      // detail 详细信息
}

// 实现 error 接口
func (e *CustomError) Error() string {
	if e.Detail != "" {
		return fmt.Sprintf("%s: %s", e.Message, e.Detail)
	} else {
		return e.Message
	}
}

type NewCErrorWithDetail func(detail ...string) *CustomError

func NewCError(httpStatus int, code api.ErrCode, message string) NewCErrorWithDetail {
	return func(detail ...string) *CustomError {
		err := &CustomError{
			HttpStatus: httpStatus,
			Code:       code,
			Message:    message,
		}
		if len(detail) > 0 {
			err.Detail = strings.Join(detail, " ")
		}
		return err
	}
}

var (
	Ok = NewCError(http.StatusOK, api.ErrCode_Ok, "ok")

	// Basic error (read comments in proto/error.proto)
	Unknown            = NewCError(http.StatusInternalServerError, api.ErrCode_Unknown, "unknown error")            // 未知错误
	InvalidArgument    = NewCError(http.StatusBadRequest, api.ErrCode_InvalidArgument, "invaild argument")          // 参数错
	NotFound           = NewCError(http.StatusBadRequest, api.ErrCode_NotFound, "not found")                        // 实体不存在
	AlreadyExists      = NewCError(http.StatusBadRequest, api.ErrCode_AlreadyExists, "already exists")              // 创建实体时冲突
	PermissionDenied   = NewCError(http.StatusForbidden, api.ErrCode_PermissionDenied, "permission denied")         // 权限不足
	ResourceExhausted  = NewCError(http.StatusTooManyRequests, api.ErrCode_ResourceExhausted, "resource exhausted") // 资源不足
	FailedPrecondition = NewCError(http.StatusBadRequest, api.ErrCode_FailedPrecondition, "failed precondition")    // 前置条件失败
	Aborted            = NewCError(http.StatusBadRequest, api.ErrCode_Aborted, "aborted")                           // 操作被中止
	OutOfRange         = NewCError(http.StatusBadRequest, api.ErrCode_OutOfRange, "out of range")                   // 超出范围
	Internal           = NewCError(http.StatusInternalServerError, api.ErrCode_Internal, "internal error")          // 内部错误
	Unavailable        = NewCError(http.StatusServiceUnavailable, api.ErrCode_Unavailable, "unavailable")           // 服务不可用，请重试
	DataLoss           = NewCError(http.StatusInternalServerError, api.ErrCode_DataLoss, "data loss")               // 数据丢失或损坏
	Unauthenticated    = NewCError(http.StatusUnauthorized, api.ErrCode_Unauthenticated, "unauthenticated")         // 未认证，客户端未提供凭据或提供的凭据无效

	// 以上是框架生成的, 自行新增的, 请将code值从1001开始
)
