// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v1.0.0-EE
// Author: meta-egg
// Generated at: 2023-04-16 18:27
package server

import (
	"encoding/json"
	"strings"

	"meta-egg-layout/internal/common/cerror"
	"meta-egg-layout/internal/common/contexts"
	handler "meta-egg-layout/internal/handler/http"

	jgjwt "github.com/Jinglever/go-jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
			logrus.Errorf("%s: authorization token is not provided", c.Request.RequestURI)
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
			logrus.Errorf("%s: invalid authorization token", c.Request.RequestURI)
			c.Error(cerror.Unauthenticated("invalid authorization token"))
			c.Abort()
			return
		}

		// get me
		var me contexts.ME
		err = json.Unmarshal([]byte(jwtClaim.Payload), &me)
		if err != nil {
			logrus.Errorf("%s: invalid authorization token", c.Request.RequestURI)
			c.Error(cerror.Unauthenticated("invalid authorization token"))
			c.Abort()
			return
		}

		newCtx := contexts.SetME(c.Request.Context(), me)
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}
