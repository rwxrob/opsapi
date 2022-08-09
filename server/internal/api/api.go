package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rwxrob/opsapi/go/model"
)

func Handle(c *gin.Context, op string) {
	u := model.User{}
	log.Printf("would handle %v user %v", op, u)
}
