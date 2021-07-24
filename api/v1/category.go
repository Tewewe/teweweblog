package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"teweweblog/model"
	"teweweblog/utils/errmsg"
)

//添加分类
func AddCategory(c*gin.Context)  {
	var data model.Category
	_=c.ShouldBindJSON(&data)
	code=model.CheckCategory(data.Name)
	if code==errmsg.SUCCES{
		model.CreateCategory(&data)
	}
	if code == errmsg.ERROR_CATEGORY_USED {
		code = errmsg.ERROR_CATEGORY_USED
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
//todo 查询分类下的所有文章

//查询分类列表
func GetCategory(c*gin.Context)  {
	pageSize,_ :=strconv.Atoi(c.Query("pageSize"))
	pageNum,_ :=strconv.Atoi(c.Query("pageNum"))
	if pageSize==0{
		pageSize=-1
	}
	if pageNum==0{
		pageNum=1
	}
	data,total:=model.GetCategory(pageSize,pageNum)
	code=errmsg.SUCCES
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
}

// GetCateInfo 查询分类信息
func GetCategoryInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetCategoryInfo(id)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)

}


//编辑分类名
func EditCategory(c*gin.Context)  {
	var data model.Category
	id,_:=strconv.Atoi(c.Param("id"))
	_=c.ShouldBindJSON(&data)
	code=model.CheckCategory(data.Name)
	if code==errmsg.SUCCES{
		model.EditCategory(id,&data)
	}
	if code==errmsg.ERROR_CATEGORY_USED{
		c.Abort()
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
//删除分类
func DeleteCategory(c*gin.Context)  {
	id,_:=strconv.Atoi(c.Param("id"))
	code=model.DeleteCategory(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}