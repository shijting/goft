package goft

import "github.com/gin-gonic/gin"

// Goft implement
type Goft struct {
	*gin.Engine
	g *gin.RouterGroup
	Port string
}

// Ignite implements Goft server
func Ignite(post string) *Goft {
	return &Goft{Engine: gin.New(), Port: post}
}

// Mont handle http router.
func (this *Goft) Mont(group string,classes ...IClass) *Goft {
	this.g = this.Group(group)
	for _, class := range classes {
		class.Build(this)
	}
	return this
}

func (this *Goft) Handle(httpMethod , relativePath string, handler interface{}) *Goft {
	//if h,ok := handler.(func(ctx *gin.Context) string);ok {
	//	this.g.Handle(httpMethod, relativePath, func(context *gin.Context) {
	//		context.String(200, h(context))
	//	})
	//}

	if h := Convert(handler); h !=nil {
		this.g.Handle(httpMethod, relativePath, h)
	}
	return this
}

// Launch runs
func (this *Goft) Launch() {
	this.Run(this.Port)
}

func (this *Goft)Attach(fs ...Fairing) *Goft  {
	this.Use(func(ctx *gin.Context) {
		for _, f := range fs {
			if err := f.OnRequest(ctx); err != nil {
				ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
			} else {
				ctx.Next()
			}
		}
	})

	return this
}
