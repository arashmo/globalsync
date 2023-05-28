package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/arashmo/globalsync/store"
)

type Handler struct {
	Store store.Store
}

func NewHandler(s store.Store) *Handler {
	return &Handler{
		Store: s,
	}
}

type datasetRequest struct {
	Name    string `json:"name"`
	Size    int    `json:"size"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

func (h *Handler) CreateDataset(c *gin.Context) {
	var req datasetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Store.InsertDataset(req.Name, req.Size, req.Version, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
