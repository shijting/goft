package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserMid struct {

}

func NewUserMid() *UserMid {
	return &UserMid{}
}

func (m *UserMid)OnRequest(ctx *gin.Context) error  {
	fmt.Println("user middleware")
	return nil
}