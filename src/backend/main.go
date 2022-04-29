package main

import (
	"database/sql"
	"dna-matcher/dbhandler"
	smalgorithm "dna-matcher/sm_algorithm"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var db *sql.DB
var pasien = []Pasien{
	{IdPengguna: 1, NamaPengguna: "Zian", RantaiDNA: "ACGTACGTACGTCGTA", NamaPenyakit: "siput gila", TanggalPrediksi: "2006-January-02"},
}

var query = []Query{
	{IdQuery: 1, SearchQuery: "2022-04-29"},
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
	router.POST("/data", postQuery)
	router.GET("/data", getData)

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
	err1 := dbhandler.InsertJenisPenyakit(db, newJenisPenyakit)
	if err1 != nil {
		fmt.Println("Rantai DNA yang di input tidak valid")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "KO",
			"message": "rantai dna tidak valid",
			"data":    err1,
		})
		return
	} else {
		c.IndentedJSON(http.StatusCreated, newJenisPenyakit)
	}
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
	IdPengguna      int64  `json:"id"`
	NamaPengguna    string `json:"nama_pengguna"`
	RantaiDNA       string `json:"rantai_dna"`
	NamaPenyakit    string `json:"nama_penyakit"`
	TanggalPrediksi string `json:"tanggal_prediksi"`
}

type Query struct {
	IdQuery     int64  `json:"id"`
	SearchQuery string `json:"query"`
}

func postQuery(c *gin.Context) {
	var newQuery Query

	if err := c.BindJSON(&newQuery); err != nil {
		return
	}

	query = append(query, newQuery)
	c.IndentedJSON(http.StatusCreated, newQuery)

}

func getData(c *gin.Context) {

	queries := query[len(query)-1].SearchQuery

	isDate, _ := smalgorithm.IsDate(queries)

	if len(queries) > 10 {

		date := queries[:10]
		nama := queries[11:]

		hasil_prediksi, err := dbhandler.ViewHasilPrediksiByNameNDate(db, nama, date)
		fmt.Println(hasil_prediksi, err)
		if hasil_prediksi != nil {

			c.IndentedJSON(http.StatusOK, hasil_prediksi)

		} else {

			c.JSON(http.StatusNoContent, gin.H{
				"status":  "KO",
				"message": "Data tidak ditemukan",
				"data":    hasil_prediksi,
			})
		}

	} else {

		if isDate {

			hasil_prediksi, err := dbhandler.ViewHasilPrediksiByDate(db, queries)
			fmt.Println(hasil_prediksi, err)
			if hasil_prediksi != nil {

				c.IndentedJSON(http.StatusOK, hasil_prediksi)

			} else {

				c.JSON(http.StatusNoContent, gin.H{
					"status":  "KO",
					"message": "Data tidak ditemukan",
					"data":    hasil_prediksi,
				})
			}
		} else {

			hasil_prediksi, err := dbhandler.ViewHasilPrediksiByName(db, queries)
			fmt.Println(hasil_prediksi, err)

			if hasil_prediksi != nil {

				c.IndentedJSON(http.StatusOK, hasil_prediksi)

			} else {
				c.JSON(http.StatusNoContent, gin.H{
					"status":  "KO",
					"message": "Data tidak ditemukan",
					"data":    hasil_prediksi,
				})
			}
		}

	}

}

func postPasien(c *gin.Context) {
	var newPasien Pasien

	if err := c.BindJSON(&newPasien); err != nil {
		return
	}
	dnaSanitized, _ := smalgorithm.IsValid(newPasien.RantaiDNA)

	if dnaSanitized {

		pasien = append(pasien, newPasien)
		err := newPasien.periksaPenyakit()

		if err != nil {
			c.JSON(http.StatusNoContent, gin.H{
				"status":  "KO",
				"message": "Data penyakit tidak ditemukan",
				"data":    dnaSanitized,
			})
		} else {

			c.IndentedJSON(http.StatusCreated, newPasien)

		}

	} else {

		// fmt.Println("Rantai DNA yang di input tidak valid")
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "KO",
			"message": "rantai dna tidak valid",
			"data":    dnaSanitized,
		})
	}

}

func getPasien(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, pasien)
}

func (p Pasien) periksaPenyakit() error {

	// kamus lokal
	var rantai_dna string
	var mirip bool
	var err error

	sqlQuery := `SELECT rantai_dna FROM jenis_penyakit WHERE nama = $1;`
	err1 := db.QueryRow(sqlQuery, p.NamaPenyakit).Scan(&rantai_dna)
	if err1 != nil {
		// fmt.Println(err1)
		// panic(err1)
		err = err1

	} else {

		fmt.Println("\n", rantai_dna)

		dnaSanitized, err := smalgorithm.IsValid(p.RantaiDNA)
		if dnaSanitized {

			_, diff, _ := smalgorithm.KMPMod(p.RantaiDNA, rantai_dna)

			if diff >= 0.8 {

				mirip = true

			} else {

				mirip = false

			}

			sqlQuery := `INSERT INTO hasil_prediksi (tanggal_prediksi, nama_pasien, nama_penyakit, tingkat_kemiripan, status_prediksi) VALUES($1, $2, $3, $4, $5) RETURNING id;`
			id := 0
			err = db.QueryRow(sqlQuery, p.TanggalPrediksi, p.NamaPengguna, p.NamaPenyakit, diff, mirip).Scan(&id)
			if err != nil {
				panic(err)
			}

		} else {

			fmt.Println(err)

		}
	}

	return err
}
