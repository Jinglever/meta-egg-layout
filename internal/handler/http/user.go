// Code generated by meta-egg. CAREFULLY EDIT.
// WILL NOT BE replace after re-generated, unless you confirm it!
// CAREFULLY EDIT.
// Version: v3.3.0-IE-dirty
// Author: meta-egg
// Generated at: 2024-07-02 21:14

package handler

import (
	"context"
	"meta-egg-layout/internal/biz"
	"meta-egg-layout/internal/common/cerror"
	"meta-egg-layout/internal/common/constraint"
	"meta-egg-layout/internal/common/contexts"
	"meta-egg-layout/internal/repo/option"
	"time"

	jgstr "github.com/Jinglever/go-string"
	"github.com/gin-gonic/gin"
)

// 用户详情
type UserDetail struct {
	Id        uint64  `json:"id"`         //
	Name      *string `json:"name"`       // 用户名 (nullable)
	Gender    uint64  `json:"gender"`     // 性别
	Age       uint8   `json:"age"`        // 年龄
	IsOnJob   bool    `json:"is_on_job"`  // 是否在职
	Birthday  *string `json:"birthday"`   // 生日 (nullable)
	CreatedBy *uint64 `json:"created_by"` // 创建者 (nullable)
	CreatedAt string  `json:"created_at"` // 创建时间
	UpdatedBy *uint64 `json:"updated_by"` // 更新者 (nullable)
	UpdatedAt string  `json:"updated_at"` // 更新时间
}

func (h *Handler) ToUserDetail(ctx context.Context, bo *biz.UserBO) (*UserDetail, error) {
	var birthday *string
	if bo.Birthday != nil {
		*birthday = bo.Birthday.Format(constraint.DateFormat)
	}
	return &UserDetail{
		Id:        bo.ID,
		Name:      bo.Name,
		Gender:    bo.Gender,
		Age:       bo.Age,
		IsOnJob:   bo.IsOnJob,
		Birthday:  birthday,
		CreatedBy: bo.CreatedBy,
		CreatedAt: bo.CreatedAt.Format(constraint.SecondTimeFormat),
		UpdatedBy: bo.UpdatedBy,
		UpdatedAt: bo.UpdatedAt.Format(constraint.SecondTimeFormat),
	}, nil
}

type ReqCreateUser struct {
	Name     *string `json:"name" binding:"omitempty,max=64"`                  // 用户名
	Gender   uint64  `json:"gender" binding:"required,gte=1"`                  // 性别
	Age      uint8   `json:"age" binding:"required"`                           // 年龄
	IsOnJob  bool    `json:"is_on_job" binding:"omitempty"`                    // 是否在职
	Birthday *string `json:"birthday" binding:"omitempty,datetime=2006-01-02"` // 生日 格式: 2006-01-02
}

//	@Id			CreateUser
//	@Tags		用户
//	@Summary	创建用户
//	@Description
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string			true	"Bearer <jwt-token>"
//	@Param		body			body		ReqCreateUser	true	"用户"
//	@Success	200				{object}	RspData{data=UserDetail}
//	@Failure	400				{object}	RspBase
//	@Router		/api/v1/users [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var req ReqCreateUser
	err := shouldBind(c, &req)
	if err != nil {
		ResponseFail(c, err)
		return
	}
	ctx := c.Request.Context()
	log := contexts.GetLogger(ctx).
		WithField("req", jgstr.JsonEncode(req))

	var birthday *time.Time
	if req.Birthday != nil {
		t, err := time.ParseInLocation(constraint.DateFormat, *req.Birthday, time.Local)
		if err != nil {
			log.WithError(err).Errorf("fail to parse time: %s", *req.Birthday)
			ResponseFail(c, cerror.InvalidArgument("Birthday"))
			return
		}
		birthday = &t
	}
	userBO := &biz.UserBO{
		Name:     req.Name,
		Gender:   req.Gender,
		Age:      req.Age,
		IsOnJob:  req.IsOnJob,
		Birthday: birthday,
	}
	err = h.BizService.CreateUser(ctx, userBO)
	if err != nil {
		log.WithError(err).Error("BizService.CreateUser failed")
		ResponseFail(c, err)
		return
	}
	d, err := h.ToUserDetail(ctx, userBO)
	if err != nil {
		log.WithError(err).Error("convert UserBO to UserDetail failed")
		ResponseFail(c, err)
		return
	}
	ResponseSuccess(c, d)
}

//	@Id			GetUserDetail
//	@Tags		用户
//	@Summary	获取用户详情
//	@Description
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string	true	"Bearer <jwt-token>"
//	@Param		id				path		int		true	"用户ID"
//	@Success	200				{object}	RspData{data=UserDetail}
//	@Failure	400				{object}	RspBase
//	@Router		/api/v1/users/{id} [get]
func (h *Handler) GetUserDetail(c *gin.Context) {
	id := jgstr.UintVal(c.Param("id"))
	ctx := c.Request.Context()
	log := contexts.GetLogger(ctx).WithField("id", id)
	userBO, err := h.BizService.GetUserByID(ctx, id)
	if err != nil {
		log.WithError(err).Error("BizService.GetUserByID failed")
		ResponseFail(c, err)
		return
	}
	d, err := h.ToUserDetail(ctx, userBO)
	if err != nil {
		log.WithError(err).Error("convert UserBO to UserDetail failed")
		ResponseFail(c, err)
		return
	}
	ResponseSuccess(c, d)
}

// 用户列表信息
type UserListInfo struct {
	Id     uint64  `json:"id"`     //
	Name   *string `json:"name"`   // 用户名 (nullable)
	Gender uint64  `json:"gender"` // 性别
}

func (h *Handler) ToUserListInfo(ctx context.Context, objs []*biz.UserListBO) ([]*UserListInfo, error) {
	list := make([]*UserListInfo, 0, len(objs))
	for i := range objs {
		list = append(list, &UserListInfo{
			Id:     objs[i].ID,
			Name:   objs[i].Name,
			Gender: objs[i].Gender,
		})
	}
	return list, nil
}

// 用户列表
type UserList struct {
	List  []*UserListInfo `json:"list"`  // 用户列表
	Total int64           `json:"total"` // 总数
}

type ReqGetUserList struct {
	Page     int `form:"page" binding:"required,gte=1"`      // 页码, 从1开始
	PageSize int `form:"page_size" binding:"required,gte=1"` // 每页数量, 要求大于0
	// 筛选条件
	Gender  *uint64 `form:"gender" binding:"omitempty,gte=1"` // 性别
	IsOnJob *bool   `form:"is_on_job" binding:"omitempty"`    // 是否在职
	// 排序条件
	OrderBy   *string `form:"order_by" binding:"omitempty,oneof=id"`         // 排序字段,可选:id
	OrderType *string `form:"order_type" binding:"omitempty,oneof=asc desc"` // 排序类型,默认desc
}

//	@Id			GetUserList
//	@Tags		用户
//	@Summary	获取用户列表
//	@Description
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string	true	"Bearer <jwt-token>"
//	@Param		page			query		int		true	"页码, 从1开始"
//	@Param		page_size		query		int		true	"每页数量, 要求大于0"
//	@Param		gender			query		uint64	false	"性别"
//	@Param		is_on_job		query		bool	false	"是否在职"
//	@Param		order_by		query		string	false	"排序字段, 可选: id"
//	@Param		order_type		query		string	false	"排序类型,默认desc"
//	@Success	200				{object}	RspData{data=UserList}
//	@Failure	400				{object}	RspBase
//	@Router		/api/v1/users [get]
func (h *Handler) GetUserList(c *gin.Context) {
	var req ReqGetUserList
	err := shouldBind(c, &req)
	if err != nil {
		ResponseFail(c, err)
		return
	}
	ctx := c.Request.Context()
	log := contexts.GetLogger(ctx).
		WithField("req", jgstr.JsonEncode(req))
	opt := &biz.UserListOption{
		Pagination: &option.PaginationOption{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		Order: &option.OrderOption{
			OrderBy:   req.OrderBy,
			OrderType: req.OrderType,
		},
		Filter: &biz.UserFilterOption{
			Gender:  req.Gender,
			IsOnJob: req.IsOnJob,
		},
	}
	userBOs, total, err := h.BizService.GetUserList(ctx, opt)
	if err != nil {
		log.WithError(err).Error("BizService.GetUserList failed")
		ResponseFail(c, err)
		return
	}
	list, err := h.ToUserListInfo(ctx, userBOs)
	if err != nil {
		log.WithError(err).Error("convert UserListBO to UserListInfo failed")
		ResponseFail(c, err)
		return
	}
	ResponseSuccess(c, UserList{
		List:  list,
		Total: total,
	})
}

type ReqUpdateUser struct {
	Name     *string `json:"name" binding:"omitempty,max=64"`                  // 用户名
	Gender   *uint64 `json:"gender" binding:"omitempty,gte=1"`                 // 性别
	Age      *uint8  `json:"age" binding:"omitempty"`                          // 年龄
	IsOnJob  *bool   `json:"is_on_job" binding:"omitempty"`                    // 是否在职
	Birthday *string `json:"birthday" binding:"omitempty,datetime=2006-01-02"` // 生日 格式: 2006-01-02
}

//	@Id			UpdateUser
//	@Tags		用户
//	@Summary	更新用户
//	@Description
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string			true	"Bearer <jwt-token>"
//	@Param		id				path		int				true	"用户ID"
//	@Param		body			body		ReqUpdateUser	true	"请求体"
//	@Success	200				{object}	RspBase
//	@Failure	400				{object}	RspBase
//	@Router		/api/v1/users/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	id := jgstr.UintVal(c.Param("id"))
	var req ReqUpdateUser
	err := shouldBind(c, &req)
	if err != nil {
		ResponseFail(c, err)
		return
	}
	ctx := c.Request.Context()
	log := contexts.GetLogger(ctx).
		WithField("id", id).
		WithField("req", jgstr.JsonEncode(req))
	var birthday *time.Time
	if req.Birthday != nil {
		t, err := time.ParseInLocation(constraint.DateFormat, *req.Birthday, time.Local)
		if err != nil {
			log.WithError(err).Errorf("fail to parse time: %s", *req.Birthday)
			ResponseFail(c, cerror.InvalidArgument("Birthday"))
			return
		}
		birthday = &t
	}
	setOpt := &biz.UserSetOption{
		Name:     req.Name,
		Gender:   req.Gender,
		Age:      req.Age,
		IsOnJob:  req.IsOnJob,
		Birthday: birthday,
	}
	err = h.BizService.UpdateUserByID(ctx, id, setOpt)
	if err != nil {
		log.WithError(err).Error("BizService.UpdateUserByID failed")
		ResponseFail(c, err)
		return
	}
	ResponseSuccess(c, nil)
}

//	@Id			DeleteUser
//	@Tags		用户
//	@Summary	删除用户
//	@Description
//	@Accept		json
//	@Produce	json
//	@Param		Authorization	header		string	true	"Bearer <jwt-token>"
//	@Param		id				path		int		true	"用户ID"
//	@Success	200				{object}	RspBase
//	@Failure	400				{object}	RspBase
//	@Router		/api/v1/users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	id := jgstr.UintVal(c.Param("id"))
	ctx := c.Request.Context()
	log := contexts.GetLogger(ctx).WithField("id", id)
	err := h.BizService.DeleteUserByID(ctx, id)
	if err != nil {
		log.WithError(err).Error("BizService.DeleteUserByID failed")
		ResponseFail(c, err)
		return
	}
	ResponseSuccess(c, nil)
}
