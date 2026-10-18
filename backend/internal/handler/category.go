package handler

import (
	"mall/backend/internal/model"
	"mall/backend/internal/pkg/response"
	"mall/backend/internal/repository"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct { repo *repository.CategoryRepo }
func NewCategoryHandler(repo *repository.CategoryRepo) *CategoryHandler { return &CategoryHandler{repo: repo} }

func (h *CategoryHandler) List(c *gin.Context) {
	list, err := h.repo.ListAll()
	if err != nil {
		response.Fail(c, 500, "query failed")
		return
	}
	response.OK(c, list)
}

type CategoryReq struct {
	Name string `json:"name" binding:"required"`
	Sort int    `json:"sort"`
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var req CategoryReq
	if err := c.ShouldBindJSON(&req); err != nil { response.Fail(c, 400, "参数错误"); return }
	
	cat := &model.Category{Name: req.Name, Sort: req.Sort}
	if err := h.repo.Create(cat); err != nil { response.Fail(c, 500, "创建失败"); return }
	response.OK(c, cat)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	// 简单解析 ID... 实际项目需用 strconv
	idStr := c.Param("id")
	var id uint64
	for _, char := range idStr { if char >= '0' && char <= '9' { id = id*10 + uint64(char-'0') } }
	
	if err := h.repo.Delete(id); err != nil { response.Fail(c, 500, "删除失败"); return }
	response.OK(c, "删除成功")
}
