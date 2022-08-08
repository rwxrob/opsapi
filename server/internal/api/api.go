package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Handle(c *gin.Context, op string) {
	log.Println("would handle %v", op)
}
