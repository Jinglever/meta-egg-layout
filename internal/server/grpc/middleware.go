// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v3.2.0-IE
// Author: meta-egg
// Generated at: 2024-01-01 22:48

package server

import (
	"context"
	"encoding/json"
	"strings"

	"meta-egg-layout/internal/common/cerror"
	"meta-egg-layout/internal/common/contexts"
	log "meta-egg-layout/pkg/log"

	jgjwt "github.com/Jinglever/go-jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	AuthToken     = "authorization" // "authorization" header key in metadata
	AuthTokenMark = "Bearer"        // authorization header value mark
)

// 对错误结果统一处理
func errorInterceptor(cfg *Config) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			var (
				cErr *cerror.CustomError
				ok   bool
			)
			if cErr, ok = err.(*cerror.CustomError); !ok {
				if cfg.ReturnErrorDetail {
					cErr = cerror.Unknown(err.Error())
				} else {
					cErr = cerror.Unknown()
				}
			} else if !cfg.ReturnErrorDetail {
				cErr.Detail = ""
			}
			err = status.Errorf(codes.Code(cErr.Code), cErr.Error())
		}
		return resp, err
	}
}

/*
 * 解析access token jwt token, 获得当前操作人信息
 */
func authInterceptor(jwt *jgjwt.JWT, cfg *Config) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		skipMethods := map[string]bool{
			// "/metaegglayout.MetaEggLayout/Getxxx": true,
		}
		if skipMethods[info.FullMethod] {
			return handler(ctx, req)
		}

		// Extract the metadata from the context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			log.Errorf("%s: metadata is not provided", info.FullMethod)
			return nil, cerror.Unauthenticated("metadata is not provided")
		}

		// Get the value of the "authorization" header
		authHeader := md.Get(AuthToken)
		if len(authHeader) == 0 {
			log.Errorf("%s: authorization token is not provided", info.FullMethod)
			return nil, cerror.Unauthenticated("authorization token is not provided")
		}
		tokenFields := strings.Fields(authHeader[0])
		if len(tokenFields) != 2 || tokenFields[0] != AuthTokenMark {
			log.Errorf("%s: authorization token is not provided", info.FullMethod)
			return nil, cerror.Unauthenticated("authorization token is not provided")
		}
		token := tokenFields[1]

		// Perform authentication
		var (
			jwtClaim *jgjwt.Claims
			err      error
		)
		if cfg.VerifyAccessToken {
			jwtClaim, err = jwt.DecodeHS256Token(token)
		} else {
			jwtClaim, err = jwt.DecodeTokenUnverified(token)
		}
		if err != nil {
			log.Errorf("%s: invalid authorization token", info.FullMethod)
			return nil, cerror.Unauthenticated("invalid authorization token")
		}

		// get me
		var me contexts.ME
		err = json.Unmarshal([]byte(jwtClaim.Payload), &me)
		if err != nil {
			log.Errorf("%s: invalid authorization token", info.FullMethod)
			return nil, cerror.Unauthenticated("invalid authorization token")
		}

		newCtx := contexts.SetME(ctx, me)
		newCtx = contexts.SetLogger(newCtx, log.WithField("meID", me.ID))
		return handler(newCtx, req)
	}
}
