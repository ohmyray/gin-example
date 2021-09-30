package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ohmyray/gin-example/common"
	"github.com/ohmyray/gin-example/util"

	"github.com/ohmyray/gin-example/dto"
	"github.com/ohmyray/gin-example/model"
	"github.com/ohmyray/gin-example/response"

	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	var bindUser model.User

	if err := ctx.ShouldBind(&bindUser); err != nil {
		response.Fail(ctx, nil, "请求参数错误")
		return
	}
	fmt.Println(bindUser.Name, bindUser.Password)
	//password := ctx.DefaultQuery("password", "17607081307")
	//telephone := ctx.DefaultQuery("telephone", "123456")

	db := common.GetDB()
	var user model.User

	db.Where("telephone = ?", bindUser.Telephone).First(&user)

	if user.ID == 0 {
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": 422,
		//	"msg":  "用户不存在",
		//})

		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(bindUser.Password)); err != nil {
		//ctx.JSON(http.StatusBadRequest, map[string]interface{}{
		//	"code": 400,
		//	"msg":  "用户名或密码错误",
		//})

		response.Response(ctx, http.StatusOK, 400, nil, "用户名或密码错误")
		return
	}

	token, err := common.ReleaseToken(user)

	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
		//	"code": 500,
		//	"msg":  "服务器内部错误",
		//})
		log.Printf("token generate error: %s", err)
		response.Fail(ctx, nil, "服务器内部错误")
		return
	}

	//ctx.JSON(http.StatusOK, map[string]interface{}{
	//	"msg":  "登录成功",
	//	"code": 200,
	//	"data": gin.H{"token": token},
	//})

	response.Success(ctx, gin.H{"token": token}, "登录成功")
}

func Register(ctx *gin.Context) {

	db := common.GetDB()

	var bindUser model.User

	if err := ctx.ShouldBind(&bindUser); err != nil {
		response.Fail(ctx, nil, "请求参数错误")
		return
	}

	name := bindUser.Name
	telephone := bindUser.Telephone
	password := bindUser.Password

	if len(telephone) != 11 {
		//ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
		//	"code": 422,
		//	"msg":  "手机号必须为11位",
		//})

		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		//ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
		//	"code": 422,
		//	"msg":  "密码不能少于6位",
		//})

		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}

	if isTelephoneExist(db, telephone) {
		//ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
		//	"code": 422,
		//	"msg":  telephone + "已注册",
		//})

		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, telephone+"已注册")
		return
	}

	hasedPaddword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		//ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
		//	"code": 450,
		//	"msg":  "加密错误",
		//})

		response.Response(ctx, http.StatusInternalServerError, 450, nil, "加密错误")
		return
	}

	//	如果名称都没有传，给一个10位的随机字符串
	if len(name) < 3 {
		name = util.RandomString(10)
	}

	newUser := model.User{
		Name:      name,
		Password:  string(hasedPaddword),
		Telephone: telephone,
	}
	db.Create(&newUser)

	log.Println(name, telephone, password)
	//ctx.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"msg":  "注册成功",
	//})

	response.Success(ctx, nil, "注册成功")
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	//ctx.JSON(http.StatusOK, map[string]interface{}{
	//	"code": 200,
	//	"data": gin.H{
	//		"user": dto.TransformToUserDto(user.(model.User)),
	//	},
	//})

	response.Success(ctx, gin.H{"user": dto.TransformToUserDto(user.(model.User))}, "查询成功")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
