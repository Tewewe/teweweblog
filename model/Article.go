package model

import (
	"gorm.io/gorm"
	"teweweblog/utils/errmsg"
)

type  Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
}


//新增文章
func CreateArticle(data *Article) int {
	err:=db.Create(&data).Error
	if err!=nil{
		return errmsg.ERROR  //500
	}
	return errmsg.SUCCES //200
}

// 查询分类下的所有文章
func GetCatArts(id int ,pageSize int,pageNum int) ([]Article,int,int64) {
	var cateartlist []Article
	var total int64
	offset := (pageNum - 1) * pageSize
	err := db.Preload("Category").Limit(pageSize).Offset(offset).Where("cid=?",id).Find(&cateartlist).Count(&total).Error
	if err!=nil{
		return nil,errmsg.ERROR_CATEGORY_NOT_EXIST,0
	}
	return cateartlist,errmsg.SUCCES,total
}

// 查询单个文章
func GetSinArticle(id int) (Article,int){
	var article Article
	err :=db.Preload("Category").Where("id=?",id).First(&article).Error
	if err!=nil{
		return article,errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return article,errmsg.SUCCES
}

// 查询文章列表
func GetArticle(title string, pageSize int,pageNum int) ([]Article,int,int64) {
	//pageSize 分页：每页个数
	//pageNum 分页：当前页码
	var articlelist []Article
	var total int64
	var err error
	if title==""{
		err=db.Order("Updated_At DESC").Preload("Category").Find(&articlelist).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
		//单独计数
		db.Model(&articlelist).Count(&total)
		if err!=nil && err!=gorm.ErrRecordNotFound{
			return nil,errmsg.ERROR,0
		}
		return articlelist,errmsg.SUCCES,total
	}
	err = db.Order("Updated_At DESC").Preload("Category").Where("title LIKE ?",title+"%").Find(&articlelist).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&articlelist).Where("title LIKE ?",title+"%").Count(&total)
	if err!=nil && err!=gorm.ErrRecordNotFound{
		return nil,errmsg.ERROR,0
	}
	return articlelist,errmsg.SUCCES,total
}

//编辑文章
func EditArticle(id int,data *Article)int {
	var article Article
	var maps=make(map[string]interface{})
	maps["title"]=data.Title
	maps["cid"]=data.Cid
	maps["desc"]=data.Desc
	maps["content"]=data.Content
	maps["img"]=data.Img
	err=db.Model(&article).Where("id=?",id).Updates(maps).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}

//删除文章
func DeleteArticle(id int) int {
	var article Article
	err=db.Where("id=?",id).Delete(&article).Error
	if err!=nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCES
}