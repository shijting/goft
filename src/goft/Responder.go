package goft

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

var ResponderList []Responder

func init() {
	ResponderList = []Responder{new(StringResponder), new(ModelResponder), new(ModelsResponder)}
}

type Responder interface {
	RespondTo() gin.HandlerFunc
}

func Convert(handler interface{}) gin.HandlerFunc {
	h_ref := reflect.ValueOf(handler)
	//fmt.Printf("%#v\n", ResponderList)
	for _, r := range ResponderList {

		r_ref := reflect.ValueOf(r).Elem()
		// 判断两个类型是否相等
		if h_ref.Type().ConvertibleTo(r_ref.Type()) {
			fmt.Println(1)
			r_ref.Set(h_ref)
			// 调用接口
			return r_ref.Interface().(Responder).RespondTo()
		}
	}
	return nil
}

// 业务控制器返回string
type StringResponder func(*gin.Context) string

func (this StringResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, this(ctx))
	}
}

// 业务控制器返回struct
type ModelResponder func(*gin.Context) Model
func (this ModelResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, this(ctx))
	}
}

type ModelsResponder func(*gin.Context) Models
func (this ModelsResponder) RespondTo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteString(string(this(ctx)))
	}
}
