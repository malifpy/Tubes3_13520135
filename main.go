package main

import (
	"database/sql"
	"dna-matcher/dbhandler"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {

	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	router := gin.Default()
    router.Use(cors.Default())
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

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

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

	if err := c.BindJSON(&newJenisPenyakit); err != nil {
		return
	}

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

	if err := c.BindJSON(&newHasilPrediksi); err != nil {
		return
	}

	dbhandler.InsertHasilPrediksi(db, newHasilPrediksi)
	c.IndentedJSON(http.StatusCreated, newHasilPrediksi)
}
