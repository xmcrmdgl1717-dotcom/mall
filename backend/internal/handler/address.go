package handler

import (
	"mall/backend/internal/model"
	"mall/backend/internal/pkg/response"
	"mall/backend/internal/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddressHandler struct { repo *repository.AddressRepo }
func NewAddressHandler(repo *repository.AddressRepo) *AddressHandler { return &AddressHandler{repo: repo} }

func (h *AddressHandler) List(c *gin.Context) {
	list, err := h.repo.List(1) // 依然暂时写死 UserID 为 1
	if err != nil { response.Fail(c, 500, "获取失败"); return }
	response.OK(c, list)
}

type AddressReq struct {
	Name      string `json:"name" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Detail    string `json:"detail" binding:"required"`
	IsDefault int8   `json:"is_default"`
}

func (h *AddressHandler) Create(c *gin.Context) {
	var req AddressReq
	if err := c.ShouldBindJSON(&req); err != nil { response.Fail(c, 400, "参数错误"); return }

	addr := &model.Address{
		UserID: 1, Name: req.Name, Phone: req.Phone,
		Province: req.Province, City: req.City, Detail: req.Detail,
		IsDefault: req.IsDefault,
	}
	
	if err := h.repo.Create(addr); err != nil { response.Fail(c, 500, "新增失败"); return }
	response.OK(c, addr)
}

func (h *AddressHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)
	if err := h.repo.Delete(id, 1); err != nil { response.Fail(c, 500, "删除失败"); return }
	response.OK(c, "删除成功")
}
