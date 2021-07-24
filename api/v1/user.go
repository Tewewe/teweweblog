package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"teweweblog/model"
	"teweweblog/utils/errmsg"
	"teweweblog/utils/validator"
)

var code int

//添加用户
func AddUser(c*gin.Context)  {
	var data model.User
	var msg string
	_=c.ShouldBindJSON(&data)

	msg,code=validator.Validate(&data)
	if code!=errmsg.SUCCES{
		c.JSON(http.StatusOK,gin.H{
			"status":code,
			"message":msg,
		})
		return
	}

	code=model.CheckUser(data.Username)
	if code==errmsg.SUCCES{
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
//查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetUser(id)
	c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   1,
			"message": errmsg.GetErrMsg(code),
		})
}


//查询用户列表
func GetUsers(c*gin.Context)  {
	pageSize,_ :=strconv.Atoi(c.Query("pageSize"))
	pageNum,_ :=strconv.Atoi(c.Query("pageNum"))
	username := c.Query("username")
	if pageSize==0{
		pageSize=-1
	}
	if pageNum==0{
		pageNum=1
	}
	data,total:=model.GetUsers(username,pageSize,pageNum)
	code=errmsg.SUCCES
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
}
//编辑用户信息
func EditUser(c*gin.Context)  {
	var data model.User
	id,_:=strconv.Atoi(c.Param("id"))
	_=c.ShouldBindJSON(&data)
	code=model.CheckUpUser(id,data.Username)
	if code==errmsg.SUCCES{
		model.EditUser(id,&data)
	}
	if code==errmsg.ERROR_USERNAME_USED{
		c.Abort()
	}
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}

// 修改密码
func ChangeUserPassword(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code = model.ChangePassword(id, &data)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

//删除用户
func DeleteUser(c*gin.Context)  {
	id,_:=strconv.Atoi(c.Param("id"))
	code=model.DeleteUser(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}

