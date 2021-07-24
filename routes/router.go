package routes

import (
	"github.com/gin-gonic/gin"
	v1 "teweweblog/api/v1"
	"teweweblog/middleware"
	"teweweblog/utils"
)

func InitRouter(){
	gin.SetMode(utils.AppMode)
	r:=gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("static/admin/index.html")
	r.Static("admin/static","static/admin/static")
	r.StaticFile("admin/favicon.ico", "static/admin/favicon.ico")

	r.GET("admin", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	authv1:=r.Group("api/v1")
	authv1.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		authv1.PUT("user/:id",v1.EditUser)
		authv1.DELETE("user/:id",v1.DeleteUser)

		//修改密码
		authv1.PUT("admin/changepw/:id", v1.ChangeUserPassword)

		//分类模块的路由接口
		authv1.POST("category/add",v1.AddCategory)
		authv1.PUT("category/:id",v1.EditCategory)
		authv1.DELETE("category/:id",v1.DeleteCategory)

		//文章模块的路由接口
		authv1.POST("article/add",v1.AddArticle)
		authv1.PUT("article/:id",v1.EditArticle)
		authv1.DELETE("article/:id",v1.DeleteArticle)

		//上传文件
		authv1.POST("upload",v1.UpLoad)

		// 更新个人设置
		authv1.PUT("profile/:id", v1.UpdateProfile)
	}
	routerv1:=r.Group("api/v1")
	{
		routerv1.POST("user/add",v1.AddUser)
		routerv1.GET("users",v1.GetUsers)
		routerv1.GET("users/:id",v1.GetUserInfo)


		routerv1.GET("category",v1.GetCategory)
		routerv1.GET("article/catartlist/:id",v1.GetCatArts)
		routerv1.GET("category/:id",v1.GetCategoryInfo)

		routerv1.GET("article/singleart/:id",v1.GetSinArticle)
		routerv1.GET("article",v1.GetArticle)
		routerv1.POST("login",v1.Login)

		// 获取个人设置信息
		routerv1.GET("profile/:id", v1.GetProfile)
	}

	r.Run(utils.HttpPort)
}