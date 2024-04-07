package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
Navigating code
@Structs
@main
@functions
*/

// _standalone package rather than a library
// _Store album data in memory
// _Defining Data
// album data about record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// _Main
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

// _Get
// getAlbums responds with list of all albums as json
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// GetAlbumByID lcoates album by ID (WOAH)
// param sent by client, return album as response
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	//Loop list of albums
	// find ID match param
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// _POST
// postAlbums adds album from json in request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind received JSON to
	// newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// add new album to slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
