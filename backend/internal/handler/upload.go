package handler

import (
	"fmt"
	"math/rand"
	"path/filepath"
	"time"

	"mall/backend/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type UploadHandler struct{}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}

// UploadImage 处理图片上传
func (h *UploadHandler) UploadImage(c *gin.Context) {
	// 1. 接收文件
	file, err := c.FormFile("file")
	if err != nil {
		response.Fail(c, 400, "获取文件失败")
		return
	}

	// 2. 生成唯一文件名 (防止重名覆盖)
	ext := filepath.Ext(file.Filename)
	// 使用时间戳+随机数
	fileName := fmt.Sprintf("%d_%d%s", time.Now().Unix(), rand.Intn(1000), ext)
	
	// 3. 保存路径 (Render 环境下存在根目录的 uploads 文件夹)
	savePath := filepath.Join("uploads", fileName)

	// 4. 保存文件
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		response.Fail(c, 500, "保存文件失败")
		return
	}

	// 5. 拼接并返回完整的访问 URL
	scheme := "http"
	if c.Request.TLS != nil || c.GetHeader("X-Forwarded-Proto") == "https" {
		scheme = "https"
	}
	host := c.Request.Host
	// 返回给前端可以直接使用的完整图片地址
	fullURL := fmt.Sprintf("%s://%s/uploads/%s", scheme, host, fileName)

	response.OK(c, gin.H{"url": fullURL})
}
