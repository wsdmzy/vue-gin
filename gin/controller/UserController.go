package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"ziogie.top/gin/common"
	"ziogie.top/gin/dto"
	"ziogie.top/gin/model"
	"ziogie.top/gin/response"
	"ziogie.top/gin/util"
)

//注册
func Register(ctx *gin.Context) {
	DB := common.GetDB()
	//	1. 获取参数
	//使用map获取请求的参数
	//var requestMap = make(map[string]string)
	//json.NewDecoder(ctx.Request.Body).Decode(&requestMap)
	//使用结构体和gin提供的Bind
	var requestUser = model.User{}
	ctx.Bind(&requestUser)

	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	//2. 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")

		return
	}

	//如果没有传名称  给一个10位随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	//3. 判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号已经注册")
		return
	}
	//4. 创建用户
	haedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(haedPassword),
	}
	DB.Create(&newUser)
	//5. 返回结果
	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "token生成失败")
		log.Printf("token generate error: %v", err)
		return
	}
	//返回结果

	response.Success(ctx, gin.H{"token": token}, "注册成功")
}

//登陆
func Login(ctx *gin.Context) {
	DB := common.GetDB()
	//	获取参数
	var requestUser = model.User{}
	ctx.Bind(&requestUser)

	telephone := requestUser.Telephone
	password := requestUser.Password

	//数据验证
	if len(telephone) != 11 {

		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}
	//判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(ctx, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "token生成失败")
		log.Printf("token generate error: %v", err)
		return
	}
	//返回结果

	response.Success(ctx, gin.H{"token": token}, "登陆成功")
}

//获取用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
	//response.Success(ctx, gin.H{"user": dto.ToUserDto(user.(model.User))},"")
}

//查询手机号存在
func isTelephoneExist(DB *gorm.DB, telephone string) bool {
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
