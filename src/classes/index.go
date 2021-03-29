package classes

import (
	"github.com/gin-gonic/gin"
	"github.com/shijting/goft/src/goft"
	"github.com/shijting/goft/src/models"
)

type IndexClass struct {
}

func NewIndexClass() *IndexClass {
	return &IndexClass{}
}

func (this *IndexClass) GetIndex(ctx *gin.Context) string {
	return "abc"
}

func (this *IndexClass) Detail(ctx *gin.Context) goft.Model {
	return &models.IndexModel{
		Id:   10,
		Name: "sldjf",
	}
}

func (this *IndexClass) List(ctx *gin.Context) goft.Models {
	users := []*models.IndexModel{
		&models.IndexModel{
			Id:   1,
			Name: "li",
		},
		&models.IndexModel{
			Id:   2,
			Name: "zh",
		},
	}
	return goft.MakeModels(users)
}

func (this *IndexClass) Build(goft *goft.Goft) {
	goft.Handle("GET", "/detail", this.Detail).
		Handle("GET", "/", this.GetIndex).
		Handle("GET", "/list", this.List)

}
