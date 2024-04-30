package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Albums list
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func handler_new(c *gin.Context) {
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Catalog of music albums"})
}

// Route declaration
func router() *mux.Router {
	r := mux.NewRouter()
	//r.HandleFunc("/", handler_new)
	return r
}
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
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

// Initiate web server
func main() {

	/*	router := router()
		srv := &http.Server{
			Handler:      router,
			Addr:         "127.0.0.1:9100",
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		} */

	router1 := gin.Default()
	router1.GET("/", handler_new)
	router1.GET("/albums", getAlbums)
	router1.GET("/albums/:id", getAlbumByID)
	router1.POST("/albums", postAlbums)
	router1.Run("localhost:8081")

	//log.Fatal(srv.ListenAndServe())
}
