package routers

import (
	_ "github.com/VENI-VIDIVICI/go-blog/docs"
	"github.com/VENI-VIDIVICI/go-blog/internal/middleware"
	v1 "github.com/VENI-VIDIVICI/go-blog/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(middleware.Transition())
	tag := v1.NewTag()
	article := v1.NewArticle()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
