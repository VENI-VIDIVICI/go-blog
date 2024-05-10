package v1

import "github.com/gin-gonic/gin"

type Tag struct {
}

func NewTag() Tag {
	return Tag{}
}

// 新增标签

func (t Tag) CreateTag(c *gin.Context) {

}

// 删除指定标签

func (t Tag) DeleteTagById(c *gin.Context) {

}

//更新指定标签

func (t Tag) UpdateTagById(c *gin.Context) {

}

//获取全部标签

func (t Tag) GetAll(c *gin.Context) {

}

