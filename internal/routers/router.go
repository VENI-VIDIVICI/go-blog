package routers

import (
	v1 "github.com/VENI-VIDIVICI/go-blog/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	tag := v1.NewTag()
	article := v1.NewArticle()
	api := r.Group("/api/v1")
	{
		// tag
		api.POST("/tag", tag.CreateTag)
		api.DELETE("/tags/:id", tag.DeleteTagById)
		api.PUT("/tags/:id", tag.UpdateTagById)
		api.GET("/tags", tag.GetAll)

		//	article
		api.POST("/article", article.CreateArticle)
		api.DELETE("/article/:id", article.DeleteArticleById)
		api.PUT("/article/:id", article.UpdateArticleById)
		api.GET("/article/:id", article.GetArticleById)
		api.GET("/articles", article.GetArticleList)
	}
	return r
}
