package controller

import (
	"log"
	"net/http"

	"funboy.top/ginessential/common"
	"funboy.top/ginessential/model"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(ctx *gin.Context) {
	DB := common.GetDB()

	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	// 判断密码是否准确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	// 发放token
	token := "111"

	// 返回结果
	ctx.JSON(200, gin.H{"code": 200, "data": gin.H{"token": token}, "msg": "登录成功"})
}

func Register(ctx *gin.Context) {
	DB := common.GetDB()

	// 获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	if len(name) <= 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户名不能为空"})
		return
	}

	// 判断手机号是否存在
	if isTelephoneExits(DB, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号已经存在"})
		return
	}
	// 创建用户
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}
	DB.Create(&newUser)

	log.Println(newUser)

	// 返回结果
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "注册成功",
	})
}

func isTelephoneExits(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
