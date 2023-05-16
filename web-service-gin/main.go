package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Struct tags such as json:"artist" specify what a field's name should be when the struct's contents are serialized into JSON
// Without them, the JSON would use the struct's capitalized field names - a style not as common in JSON

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// Initialize a Gin router using Default
	router := gin.Default()

	// assign the handler function to an endpoint path
	router.GET("/albums", getAlbums)

	router.GET("/albums/:id", getAlbumByID)

	router.POST("/albums", postAlbums)

	// Use the Run function to attach the router to an http.Server and start the server
	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON
// gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
	// Note that you can replace Context.IndentedJSON with a call to Context.JSON to send more compact JSON
	// In practice, the indented form is much easier to work with when debugging and the size difference is usually small
}

// postAlbums adds an album from JSON received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response
func getAlbumByID(c *gin.Context) {
	// User Context.Param to retrieve the id path parameter from the URL
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	// A real-world service would likely use a database query to perform this lookup

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
