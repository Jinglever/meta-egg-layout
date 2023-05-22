// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v1.5.1-EE
// Author: meta-egg
// Generated at: 2023-05-22 16:28

package server

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"meta-egg-layout/internal/common/cerror"
	"meta-egg-layout/internal/common/contexts"
	handler "meta-egg-layout/internal/handler/http"
	log "meta-egg-layout/pkg/log"

	jgjwt "github.com/Jinglever/go-jwt"
	"github.com/gin-gonic/gin"
)

const (
	AuthToken     = "Authorization" // authorization header key
	AuthTokenMark = "Bearer"        // authorization header value mark
)

// 对错误结果统一处理
func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errs := c.Errors.ByType(gin.ErrorTypeAny)
		if len(errs) > 0 {
			err := errs.Last().Err
			var (
				cErr *cerror.CustomError
				ok   bool
			)
			if cErr, ok = err.(*cerror.CustomError); !ok {
				cErr = cerror.Unknown(err.Error())
			}
			c.JSON(
				cErr.HttpStatus,
				handler.RspBase{
					Code:    cErr.Code,
					Message: cErr.Error(),
				},
			)
		}
	}
}

/*
 * 解析jwt token, 获得当前操作人信息
 */
func authHandler(jwt *jgjwt.JWT, cfg *Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenFields := strings.Fields(c.GetHeader(AuthToken))
		if len(tokenFields) != 2 || tokenFields[0] != AuthTokenMark {
			log.Errorf("%s: authorization token is not provided", c.Request.RequestURI)
			c.Error(cerror.Unauthenticated("authorization token is not provided"))
			c.Abort()
			return
		}
		token := tokenFields[1]

		var (
			jwtClaim *jgjwt.Claims
			err      error
		)
		if cfg.VerifyJWT {
			jwtClaim, err = jwt.DecodeHS256Token(token)
		} else {
			jwtClaim, err = jwt.DecodeTokenUnverified(token)
		}
		if err != nil {
			log.Errorf("%s: invalid authorization token", c.Request.RequestURI)
			c.Error(cerror.Unauthenticated("invalid authorization token"))
			c.Abort()
			return
		}

		// get me
		var me contexts.ME
		err = json.Unmarshal([]byte(jwtClaim.Payload), &me)
		if err != nil {
			log.Errorf("%s: invalid authorization token", c.Request.RequestURI)
			c.Error(cerror.Unauthenticated("invalid authorization token"))
			c.Abort()
			return
		}

		newCtx := contexts.SetME(c.Request.Context(), me)
		newCtx = contexts.SetLogger(newCtx, log.WithField("meID", me.ID))
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		var statusColor, methodColor, resetColor string
		if param.IsOutputColor() {
			statusColor = param.StatusCodeColor()
			methodColor = param.MethodColor()
			resetColor = param.ResetColor()
		}

		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		me, ok := contexts.GetME(param.Request.Context())
		if ok {
			return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s | meID=%v |%s %-7s %s %#v\n%s",
				param.TimeStamp.Format("2006/01/02 - 15:04:05"),
				statusColor, param.StatusCode, resetColor,
				param.Latency,
				param.ClientIP,
				me.ID,
				methodColor, param.Method, resetColor,
				param.Path,
				param.ErrorMessage,
			)
		} else {
			return fmt.Sprintf("[GIN] %v |%s %3d %s| %13v | %15s | meID=%v |%s %-7s %s %#v\n%s",
				param.TimeStamp.Format("2006/01/02 - 15:04:05"),
				statusColor, param.StatusCode, resetColor,
				param.Latency,
				param.ClientIP,
				0,
				methodColor, param.Method, resetColor,
				param.Path,
				param.ErrorMessage,
			)
		}
	})
}
