package model

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
	"teweweblog/utils/errmsg"
)

type  User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null "json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null "json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role int `gorm:"type:int;DEFAULT:2"json:"role" validate:"required,gte=1" label:"角色码"`
}

//查询用户是否存在
func CheckUser(name string)(code int){
	var users User
	db.Select("id").Where("username=?",name).First(&users)
	if users.ID>0{
		return errmsg.ERROR_USERNAME_USED   //1001
	}
	return errmsg.SUCCES
}

// CheckUpUser 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCES
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCES
}

//新增用户
func CreateUser(data *User) int {
	//data.Password=ScryptPw(data.Password)
	err:=db.Create(&data).Error
	if err!=nil{
		return errmsg.ERROR  //500
	}
	return errmsg.SUCCES //200
}

//查询单个用户
// 查询用户
func GetUser(id int) (User, int) {
	var user User
	err = db.Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCES
}

//查询用户列表
func GetUsers(username string,pageSize int,pageNum int) ([]User,int64) {
	//pageSize 分页：每页个数
	//pageNum 分页：当前页码
	var users []User
	var total int64

	if username!=""{
		db.Select("id,username,role").Where("username LIKE ?", username+"%").Find(&users).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize)
		return users,total
	}
	db.Select("id,username,role").Find(&users).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize)

	if err == gorm.ErrRecordNotFound{
		return users,0
	}
	return users,total
}

//编辑用户信息
func EditUser(id int,data *User)int {
	var user User
	var maps=make(map[string]interface{})
	maps["username"]=data.Username
	maps["role"]=data.Role
	err=db.Model(&user).Where("id=?",id).Updates(maps).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}

//删除用户
func DeleteUser(id int) int {
	var user User
	err=db.Where("id=?",id).Delete(&user).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}

// ChangePassword 修改密码
func ChangePassword(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["password"] = ScryptPw(data.Password)
	err=db.Model(&user).Where("id=?",id).Updates(maps).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}

//密码加密
func (u *User)BeforeCreate(tx *gorm.DB) (err error){
	u.Password=ScryptPw(u.Password)
	return
}

func ScryptPw(password string)string  {
	const KeyLen  =10
	salt:=make([]byte,8)
	salt=[]byte{12,32,4,6,66,22,222,11}
	HashPw,err:=scrypt.Key([]byte(password),salt,16384,8,1,KeyLen)
	if err!=nil{
		log.Fatal(err)
	}
	fpw:=base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}


//登录验证
func CheckLogin(username string,password string) int {
	var user User
	db.Where("username=?",username).First(&user)
	if user.ID==0{
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if ScryptPw(password)!=user.Password{
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role!=1{
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCES
}