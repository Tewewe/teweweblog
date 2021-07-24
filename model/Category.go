package model

import (
	"gorm.io/gorm"
	"teweweblog/utils/errmsg"
)

type Category struct {
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null"json:"name"`
}

//查询分类是否存在
func CheckCategory(name string)(code int){
	var cate Category
	db.Select("id").Where("Name=?",name).First(&cate)
	if cate.ID>0{
		return errmsg.ERROR_CATEGORY_USED   //3001
	}
	return errmsg.SUCCES
}

//新增分类
func CreateCategory(data *Category) int {
	err:=db.Create(&data).Error
	if err!=nil{
		return errmsg.ERROR  //500
	}
	return errmsg.SUCCES //200
}

// todo 查询分类下的所有文章

//查询分类列表
func GetCategory(pageSize int,pageNum int) ([]Category,int64) {
	//pageSize 分页：每页个数
	//pageNum 分页：当前页码
	var cate []Category
	var total int64
	offset := (pageNum - 1) * pageSize

	err := db.Find(&cate).Count(&total).Limit(pageSize).Offset(offset).Error
	if err!=nil && err!=gorm.ErrRecordNotFound{
		return nil,0
	}
	return cate,total
}

// GetCateInfo 查询单个分类信息
func GetCategoryInfo(id int) (Category,int) {
	var cate Category
	db.Where("id = ?",id).First(&cate)
	return cate,errmsg.SUCCES
}

//编辑分类名
func EditCategory(id int,data *Category)int {
	var cate Category
	var maps=make(map[string]interface{})
	maps["name"]=data.Name
	err=db.Model(&cate).Where("id=?",id).Updates(maps).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}

//删除分类
func DeleteCategory(id int) int {
	var cate Category
	err=db.Where("id=?",id).Delete(&cate).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}