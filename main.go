package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("main1")
	select {

	}
	fmt.Println("main2")


	router := gin.Default()
	//v

	router.Use(func(f *gin.Context) {
		//fmt.Println("jsfsf:", f.)
		//f.en
		fmt.Println("ip", f.ClientIP())
		fmt.Println("relatevePath1:", f.RelativePath)
		//f.Params
	})
	v1 := router.Group("/api/v1/todos")
	{
		v1.GET("/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"id": c.Param("id"),
			})
		})
	}

	router.GET("/test/:x/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": c.Param("x"),
		})
	})

	router.GET("/a/:a/b/:b", func(c *gin.Context) {
		fmt.Println("relatevePath2:", c.RelativePath)
		c.JSON(200, gin.H{
			"message": c.Param("a"),
		})
	})
	router.Run()
}

func test() {
	x := 100
	go func(a int) {
		a++
		println(a)
	}(x)

	x++
}

func closureTest() {
	y := 200
	go func() {
		y++
		println(y)
	}()
	y++
}
