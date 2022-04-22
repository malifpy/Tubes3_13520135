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
var pasien = []Pasien{
	{IdPengguna: 1, NamaPengguna: "Zian", RantaiDNA: "ACGTACGTACGTCGTA", PrediksiPenyakit: "siput gila"},
}

func main() {

	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/jenis_penyakit", getJenisPenyakit)
	router.POST("/jenis_penyakit", postJenisPenyakit)
	router.GET("/hasil_prediksi", getHasilPrediksi)
	router.POST("/hasil_prediksi", postHasilPrediksi)
	router.POST("/pasien", postPasien)
	router.GET("/pasien", getPasien)

	router.Run()
	defer db.Close()
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

type Pasien struct {
	IdPengguna       int64  `json:"id"`
	NamaPengguna     string `json:"nama_pengguna"`
	RantaiDNA        string `json:"rantai_dna"`
	PrediksiPenyakit string `json:"prediksi_penyakit"`
}

func postPasien(c *gin.Context) {
	var newPasien Pasien

	if err := c.BindJSON(&newPasien); err != nil {
		return
	}
	pasien = append(pasien, newPasien)
	c.IndentedJSON(http.StatusCreated, newPasien)
}

func getPasien(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pasien)
}

func (p Pasien) periksaPenyakit() {

	validator.IsValid(p.RantaiDNA)

}
