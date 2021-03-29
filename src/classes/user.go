package classes

import (
	"github.com/gin-gonic/gin"
	"github.com/shijting/goft/src/goft"
)

type UserClass struct {

}

func NewUserClass() *UserClass {
	return &UserClass{}
}

func (this *UserClass) UserList() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, gin.H{
			"result": "userlist",
		})
	}
}

func (this *UserClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/userlist", this.UserList())
}
