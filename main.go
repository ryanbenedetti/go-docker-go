package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DATA STRUCTURE: This album struct represents data for our app about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// DATA: albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// Logic to prepare responses

// getAlbums responds with the list of ALL albums as JSON.
func getAlbums(c *gin.Context) {

	//serialize the struct into JSON and add it to the response.
	c.IndentedJSON(http.StatusOK, albums)
}

// Code to map the request path to the response logic

func main() {
	// Initialize a Gin router using Default.
	router := gin.Default()

	// GET HTTP method associates the /albums path with the getAlbums handler
	router.GET("/albums", getAlbums)

	// Run attaches the router to an http.Server and starts the server
	router.Run("localhost:8080")
}
