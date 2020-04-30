package controller

import (
	"github.com/gin-gonic/gin"

	"strconv"

	"ziogie.top/gin/model"
	"ziogie.top/gin/repository"
	"ziogie.top/gin/response"
	"ziogie.top/gin/vo"
)

//定义接口
type ICategoryController interface {
	RestController
}

type CategoryController struct {
	//DB *gorm.DB
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})

	return  CategoryController{Repository: repository}
}

func (c CategoryController) Create(ctx *gin.Context)  {
	var requestCategory  vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err !=nil {
		response.Fail(ctx,nil, "数据验证错误，分类名称必填!")
		return
	}
	//拿到表单数据
	//ctx.Bind(&requestCategory)

	category,err :=  c.Repository.Create(requestCategory.Name);
	if err != nil {
		//response.Fail(ctx,nil, "创建失败")
		panic(err)
		return
	}

	response.Success(ctx, gin.H{"category": category}, "")

}

func (c CategoryController) Update(ctx *gin.Context)  {
	//body中的参数   模型绑定
	var requestCategory  vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err !=nil {
		response.Fail(ctx,nil, "数据验证错误，分类名称必填!")
		return
	}

//	获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))

	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx,nil,"分类不存在")
		return
	}

//	更新分类  map struct name value
	category,err :=  c.Repository.Update(*updateCategory,requestCategory.Name)
	if err != nil {
		panic(err)
	}

	response.Success(ctx,gin.H{"category": category}, "修改成功")
}
func (c CategoryController) Show(ctx *gin.Context)  {
	//	获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx,nil,"分类不存在")
		return
	}


	response.Success(ctx,gin.H{"category": category}, "")
}
func (c CategoryController) Delete(ctx *gin.Context)  {
	//	获取path中的参数
	categoryId,_ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.Repository.DeleteById(categoryId); err !=nil {
		response.Fail(ctx,nil,"删除失败，请重试")
	}


	response.Success(ctx,nil, "删除成功")
}