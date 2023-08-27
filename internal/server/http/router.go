// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: master-436ef70-IE
// Author: meta-egg
// Generated at: 2023-08-27 20:45

package server

import (
	hdl "meta-egg-layout/internal/handler/http"

	"github.com/gin-gonic/gin"
)

// annotation for swagger
// @title	meta-egg-layout
// @version xxx
func (s *Server) initRouter() {
	router := gin.New()
	router.Use(Logger(), gin.Recovery())
	router.Use(errorHandler())
	authGroup := router.Group("/v1")
	authGroup.Use(authHandler(s.Resource.JWT, s.Cfg))

	handler := hdl.WireHandler(s.Resource)

	// 用户
	authGroup.POST("/users", handler.CreateUser)
	authGroup.GET("/users/:id", handler.GetUserDetail)
	authGroup.GET("/users", handler.GetUserList)
	authGroup.PUT("/users/:id", handler.UpdateUser)
	authGroup.DELETE("/users/:id", handler.DeleteUser)

	s.Router = router
}
