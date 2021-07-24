package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"teweweblog/utils"
	"teweweblog/utils/errmsg"
	"time"
)

var JwtKey=[]byte(utils.JwtKey)
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

//生成token
func SetToken(username string) (string,int) {
	expireTime:=time.Now().Add(10*time.Hour)
	SetClaims:=MyClaims{
		Username:username,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "teweweblog",
		},
	}
	reqClaim:=jwt.NewWithClaims(jwt.SigningMethodHS256,SetClaims)
	token,err:=reqClaim.SignedString(JwtKey)
	if err!=nil{
		return "",errmsg.ERROR
	}
	return token,errmsg.SUCCES
}

//验证token
func CheckToken(token string) (*MyClaims,int) {
	setToken,_:=jwt.ParseWithClaims(token,&MyClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return JwtKey,nil
	})
	if key,_:=setToken.Claims.(*MyClaims);setToken.Valid{
		return key,errmsg.SUCCES
	}
	return nil,errmsg.ERROR
}

//jwt中间件
func JwtToken()gin.HandlerFunc  {
	return func(c *gin.Context) {
		tokenHeader:=c.Request.Header.Get("Authorization")
		var code int
		if tokenHeader==""{
			code=errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK,gin.H{
				"status":code,
				"message":errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken:=strings.SplitN(tokenHeader," ",2)
		if (len(checkToken)!=2&&checkToken[0]!="Bearer"){
			code=errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK,gin.H{
				"status":code,
				"message":errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key,turnCode:=CheckToken(checkToken[1])
		if turnCode==errmsg.ERROR{
			code=errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK,gin.H{
				"status":code,
				"message":errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix()>key.ExpiresAt{
			code=errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK,gin.H{
				"status":code,
				"message":errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username",key.Username)
		c.Next()
	}
}