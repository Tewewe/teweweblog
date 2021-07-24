package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"teweweblog/model"
	"teweweblog/utils/errmsg"
)

//添加文章
func AddArticle(c*gin.Context)  {
	var data model.Article
	_=c.ShouldBindJSON(&data)
	code= model.CreateArticle(&data)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
// 查询分类下的所有文章
func GetCatArts(c*gin.Context)  {
	pageSize,_ :=strconv.Atoi(c.Query("pageSize"))
	pageNum,_ :=strconv.Atoi(c.Query("pageNum"))
	if pageSize==0{
		pageSize=-1
	}
	if pageNum==0{
		pageNum=1
	}
	id,_:=strconv.Atoi(c.Param("id"))
	data,code,total:=model.GetCatArts(id,pageSize,pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
}

// 查询单个文章信息
func GetSinArticle(c*gin.Context) {
	id,_:=strconv.Atoi(c.Param("id"))
	data,code:=model.GetSinArticle(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}


// 查询文章列表
func GetArticle(c*gin.Context)  {
	pageSize,_ :=strconv.Atoi(c.Query("pageSize"))
	pageNum,_ :=strconv.Atoi(c.Query("pageNum"))
	title:=c.Query("title")

	if pageSize==0{
		pageSize=-1
	}
	if pageNum==0{
		pageNum=1
	}
	data,code,total:=model.GetArticle(title,pageSize,pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
}
//编辑文章信息
func EditArticle(c*gin.Context)  {
	var data model.Article
	id,_:=strconv.Atoi(c.Param("id"))
	_=c.ShouldBindJSON(&data)
	code=model.EditArticle(id,&data)

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
//删除文章
func DeleteArticle(c*gin.Context)  {
	id,_:=strconv.Atoi(c.Param("id"))
	code=model.DeleteArticle(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}