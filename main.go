package main

import (
	"database/sql"
	"dna-matcher/dbhandler"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

// // albums slice to seed record album data.
// var albums = []dbhandler.Album{
//     {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
//     {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
//     {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

func main() {

	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/jenis_penyakit", getJenisPenyakit)
	router.POST("/jenis_penyakit", postJenisPenyakit)
	router.GET("/hasil_prediksi", getHasilPrediksi)
	router.POST("/hasil_prediksi", postHasilPrediksi)

	router.Run()
	defer db.Close()
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	albums, err := dbhandler.ViewAllAlbums(db)
	fmt.Println(albums, err)
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum dbhandler.Album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	// albums = append(albums, newAlbum)
	dbhandler.InsertAlbums(db, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getJenisPenyakit(c *gin.Context) {
	jenis_penyakit, err := dbhandler.ViewAllJenisPenyakit(db)
	fmt.Println(jenis_penyakit, err)
	c.IndentedJSON(http.StatusOK, jenis_penyakit)
}

func postJenisPenyakit(c *gin.Context) {
	var newJenisPenyakit dbhandler.JenisPenyakit

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newJenisPenyakit); err != nil {
		return
	}

	// Add the new album to the slice.
	// albums = append(albums, newAlbum)
	dbhandler.InsertJenisPenyakit(db, newJenisPenyakit)
	c.IndentedJSON(http.StatusCreated, newJenisPenyakit)
}

func getHasilPrediksi(c *gin.Context) {
	hasil_prediksi, err := dbhandler.ViewAllHasilPrediksi(db)
	fmt.Println(hasil_prediksi, err)
	c.IndentedJSON(http.StatusOK, hasil_prediksi)
}

func postHasilPrediksi(c *gin.Context) {
	var newHasilPrediksi dbhandler.HasilPrediksi

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newHasilPrediksi); err != nil {
		return
	}

	// Add the new album to the slice.
	// albums = append(albums, newAlbum)
	dbhandler.InsertHasilPrediksi(db, newHasilPrediksi)
	c.IndentedJSON(http.StatusCreated, newHasilPrediksi)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
// func getAlbumByID(c *gin.Context) {
//     id := c.Param("id")
//
//     // Loop through the list of albums, looking for
//     // an album whose ID value matches the parameter.
//     for _, a := range albums {
//         if a.ID == id {
//             c.IndentedJSON(http.StatusOK, a)
//             return
//         }
//     }
//     c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
// }
