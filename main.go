package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

var posts = []Post{
	{Id: 1, UserId: 1, Title: "First post", Body: "First post body"},
	{Id: 2, UserId: 1, Title: "Second post", Body: "Second post body"},
	{Id: 3, UserId: 1, Title: "Third post", Body: "Third post body"},
	{Id: 4, UserId: 1, Title: "Fourth post", Body: "Fourth post body"},
}

func homeRoute(context *gin.Context) {
	context.String(http.StatusOK, "Welcome to Go Modular")
}

func postsRoute(context *gin.Context) {
	context.JSON(http.StatusOK, posts)
}

func main() {
	router := gin.Default()

	router.GET("/", homeRoute)
	router.GET("/posts", postsRoute)

	err := router.Run(":3000")
	if err != nil {
		panic(err)
	}
}
