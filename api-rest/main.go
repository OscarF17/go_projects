package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// albun Estructura de albun
type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

// Lista de albuns
var albums = []album{
	{ID: "1", Title: "Familia", Artist: "Camila Cabello", Year: 2022},
	{ID: "2", Title: "21", Artist: "Adele", Year: 2011},
	{ID: "3", Title: "The Eminem Show", Artist: "Eminem", Year: 2002},
	{ID: "4", Title: "Meteora", Artist: "Linkin Park", Year: 2003},
	{ID: "5", Title: "25", Artist: "Adele", Year: 2015},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	c.BindJSON(&newAlbum)

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8081")
}
