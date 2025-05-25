package handler

import (
	"fmt"
	"github.com/ekideno/gosnip/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createRequest struct {
	OriginalURL string `json:"original_url" binding:"required,url"`
	Slug        string `json:"slug"`
}
type URLHandler struct {
	URLService domain.URLService
}

func NewURLHandler(URLService domain.URLService) *URLHandler {
	return &URLHandler{URLService: URLService}
}

func (h *URLHandler) Create(c *gin.Context) {
	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	fmt.Println(req)

	url, err := h.URLService.Create(req.OriginalURL, req.Slug)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"url": url.Original, "slug": url.Slug})
}

func (h *URLHandler) GetBySlug(c *gin.Context) {
	slug := c.Param("slug")

	url, err := h.URLService.GetBySlug(slug)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusFound, url)
}
