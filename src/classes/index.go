package classes

import (
	"github.com/gin-gonic/gin"
	"github.com/shijting/goft/src/goft"
)

type IndexClass struct {
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}


func (this *IndexClass) GetIndex(ctx *gin.Context) string {
	return "abc"
}

func (this *IndexClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/", this.GetIndex)
}
