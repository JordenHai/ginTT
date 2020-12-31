package demos

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func noneMiddleware(){
	route := gin.New()
	route.Run()
}

func routeWithMiddleware(){
	route := gin.Default()
	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK,gin.H{
			"name":"牛奶",
		})
	})
	route.Run()
}

func CustomRouteMiddleware1(c *gin.Context){
	t := time.Now()
	fmt.Println("http works")
	c.Set("example","CustomRouteMiddleware1")
	c.Next()
	fmt.Println("http finished")
	t2 := time.Since(t)
	fmt.Println(t2)
}

func CustomRouterMiddle2() gin.HandlerFunc{
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("我是自定义中间件第2种定义方式---请求之前")
		//在gin上下文中定义一个变量
		c.Set("example", "CustomRouterMiddle2")
		//请求之前
		c.Next()
		fmt.Println("我是自定义中间件第2种定义方式---请求之后")
		//请求之后
		//计算整个请求过程耗时
		t2 := time.Since(t)
		log.Println(t2)
	}
}

func TestMiddleware() {
	//noneMiddleware()
	//routeWithMiddleware()

	r := gin.New()
	r.Use(CustomRouteMiddleware1)
	r.GET("/test", func(context *gin.Context) {
		example := context.MustGet("example").(string)
		log.Println(example)
	})
	r.Run()
}

func RouteMiddle1(c *gin.Context){
	fmt.Println("我是路由中间件1")
}

func RouteMiddle2(c *gin.Context){
	fmt.Println("我是路由中间件2")
}

func oneRouteMiddleHandle() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("it handler function")
	}
}

func routeMultiHandle()  {
	r := gin.New()
	r.GET("/test",
		RouteMiddle1,
		RouteMiddle2,
		oneRouteMiddleHandle())
	r.Run()
}

func GroupRouterGoodsMiddle1(c *gin.Context)  {
	fmt.Println("我是goods路由组中间件1")
}

func GroupRouterGoodsMiddle2(c *gin.Context) {
	fmt.Println("我是goods路由组中间件2")
}

func GroupRouterOrderMiddle1(c *gin.Context) {
	fmt.Println("我是order路由组中间件1")
}

func GroupRouterOrderMiddle2(c *gin.Context) {
	fmt.Println("我是order路由组中间件2")
}

func Test() {
	//创建一个无中间件路由
	router := gin.New()
	router.Use(gin.Logger())

	//第1种路由组使用方式 可以添加多个处理函数 但是不知道为什么 官方举例的这第一种方式用不了
	router.Group("/goods", GroupRouterGoodsMiddle1, GroupRouterGoodsMiddle2)
	router.GET("/goods/add", func(context *gin.Context) {
		fmt.Println("/goods/add")
	})


	//第2种路由组使用方式
	orderGroup := router.Group("/order")
	orderGroup.Use(GroupRouterOrderMiddle1, GroupRouterOrderMiddle2)
	{
		orderGroup.GET("/add", func(context *gin.Context) {
			fmt.Println("/order/add")

		})

		orderGroup.GET("/del", func(context *gin.Context) {
			fmt.Println("/order/del")

		})

		//orderGroup下再嵌套一个testGroup
		testGroup:=orderGroup.Group("/test", func(context *gin.Context) {
			fmt.Println("order/test下的中间件")
		})

		testGroup.GET("/test1", func(context *gin.Context) {
			fmt.Println("order/test/test1的函数")
		})
	}

	router.Run()
}
