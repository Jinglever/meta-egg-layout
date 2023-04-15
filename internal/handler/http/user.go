package handler

import (
	jgstr "github.com/Jinglever/go-string"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// 用户详情
type UserDetail struct {
	ID        uint64  `json:"id"`         // 用户ID
	Name      *string `json:"name"`       // 用户名
	Gender    uint64  `json:"gender"`     // 性别
	CreatedBy *uint64 `json:"created_by"` // 创建人
	CreatedAt int64   `json:"created_at"` // 创建时间
	UpdatedBy *uint64 `json:"updated_by"` // 更新人
	UpdatedAt int64   `json:"updated_at"` // 更新时间
}

// @Id	GetUserDetail
// @Tags	User
// @Summary	获取用户详情
// @Description
// @Accept	json
// @Produce	json
// @Param	X-Auth-Token	header	string	true	"jwt access token"
// @Param	id	path	int	true	"用户ID"
// @Success	200	{object}	RspData{data=UserDetail}
// @Failure	40x	{object}	RspBase
// @Router	/v1/users/:id [get]
func (h *Handler) GetUserDetail(c *gin.Context) {
	id := jgstr.UintVal(c.Param("id"))
	mUser, err := h.DomainUsecase.GetUserByID(c.Request.Context(), id)
	if err != nil {
		logrus.Errorf("domainUsecase.GetUserByID failed, err: %v", err)
		ResponseFail(c, err)
		return
	}
	ResponseSuccess(c, UserDetail{
		ID:        mUser.ID,
		Name:      mUser.Name,
		Gender:    mUser.Gender,
		CreatedBy: mUser.CreatedBy,
		CreatedAt: mUser.CreatedAt.Unix(),
		UpdatedBy: mUser.UpdatedBy,
		UpdatedAt: mUser.UpdatedAt.Unix(),
	})
}
