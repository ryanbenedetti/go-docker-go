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

// Logic for REST responses

// GET: getAlbums responds with the list of ALL albums as JSON.
func getAlbums(c *gin.Context) {

	//serialize the struct into JSON and add it to the response.
	c.IndentedJSON(http.StatusOK, albums)
}

// POST: postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// BindJSON binds the received JSON to a newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the albums slice.
	albums = append(albums, newAlbum)

	// Add a 201 status code to the response, along with the album JSON
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	// Initialize a Gin router using Default.
	router := gin.Default()

	// router maps the request path to the response logic

	// GET HTTP method associates /albums path with getAlbums handler
	router.GET("/albums", getAlbums)

	// GET HTTP method associates /albums/id path with getAlbums instance
	router.GET("/albums/:id", getAlbumByID)

	// POST HTTP method associates /albums path with postAlbums handler
	router.POST("/albums", postAlbums)

	// Run attaches the router to an http.Server and starts the server
	router.Run("localhost:8080")
}
