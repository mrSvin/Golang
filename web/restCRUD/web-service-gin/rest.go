package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func RestInfo() {
	router := gin.Default()
	router.GET("/", getHome)
	router.GET("/albums", getAlbums)
	router.GET("/names", getNames)
	router.Run("localhost:8090")
}

func getHome(c *gin.Context) {
	fmt.Println("ok")
	c.IndentedJSON(http.StatusOK, "ok")
}

func getAlbums(c *gin.Context) {
	fmt.Println(albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func getNames(c *gin.Context) {
	fmt.Println(SelectData("names"))
	c.IndentedJSON(http.StatusOK, SelectData("names"))
}
