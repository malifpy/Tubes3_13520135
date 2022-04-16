package dbhandler

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type JenisPenyakit struct {
	ID        int64  `json:"id"`
	Nama      string `json:"nama"`
	RantaiDNA string `json:"rantai_dna"`
}

type HasilPrediksi struct {
	ID               int64     `json:"id"`
	TanggalPrediksi  time.Time `json:"tanggal_prediksi"`
	NamaPasien       string    `json:"nama_pasien"`
	NamaPenyakit     string    `json:"nama_penyakit"`
	TingkatKemiripan float64   `json:"tingkat_kemiripan"`
	StatusPrediksi   bool      `json:"status_prediksi"`
}

func InsertAlbums(db *sql.DB, alb Album) {

	sqlQuery := `INSERT INTO albums (title, artist, price) VALUES($1, $2, $3) RETURNING id;`
	id := 0
	err := db.QueryRow(sqlQuery, alb.Title, alb.Artist, alb.Price).Scan(&id)
	if err != nil {
		panic(err)
	}
}

func ViewAllAlbums(db *sql.DB) ([]Album, error) {
	var albums []Album
	rows, err := db.Query("SELECT * FROM albums;")
	if err != nil {
		return nil, fmt.Errorf("SOMEHOW ERROR: %v", err)
	}
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("SOMEHOW ERROR: %v", err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("SOMEHOW ERROR: %v", err)
	}
	defer rows.Close()
	return albums, nil
}

func InsertJenisPenyakit(db *sql.DB, rDNA JenisPenyakit) {

	sqlQuery := `INSERT INTO jenis_penyakit (nama, rantai_dna) VALUES($1, $2) RETURNING id;`
	id := 0
	err := db.QueryRow(sqlQuery, rDNA.Nama, rDNA.RantaiDNA).Scan(&id)
	if err != nil {
		panic(err)
	}
}

func ViewAllJenisPenyakit(db *sql.DB) ([]JenisPenyakit, error) {
	var jenisPenyakit []JenisPenyakit
	rows, err := db.Query("SELECT * FROM jenis_penyakit;")
	if err != nil {
		return nil, fmt.Errorf("SOMEHOW ERROR: %v", err)
	}
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var jenis_penyakit JenisPenyakit
		if err := rows.Scan(&jenis_penyakit.ID, &jenis_penyakit.Nama, &jenis_penyakit.RantaiDNA); err != nil {
			return nil, fmt.Errorf("SOMEHOW ERROR: %v", err)
		}
		jenisPenyakit = append(jenisPenyakit, jenis_penyakit)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("SOMEHOW ERROR: %v", err)
	}
	defer rows.Close()
	return jenisPenyakit, nil
}

func InsertHasilPrediksi(db *sql.DB, resPred HasilPrediksi) {

	sqlQuery := `INSERT INTO hasil_prediksi (tanggal_prediksi, nama_pasien, nama_penyakit, tingkat_kemiripan, status_prediksi) VALUES($1, $2, $3, $4, $5) RETURNING id;`
	id := 0
	err := db.QueryRow(sqlQuery, resPred.TanggalPrediksi, resPred.NamaPasien, resPred.NamaPenyakit, resPred.TingkatKemiripan, resPred.StatusPrediksi).Scan(&id)
	if err != nil {
		panic(err)
	}
}

func ViewAllHasilPrediksi(db *sql.DB) ([]HasilPrediksi, error) {
	var hasilPrediksi []HasilPrediksi
	rows, err := db.Query("SELECT * FROM hasil_prediksi;")
	if err != nil {
		return nil, fmt.Errorf("SOMEHOW ERROR: %v", err)
	}
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var hasil_prediksi HasilPrediksi
		if err := rows.Scan(&hasil_prediksi.ID, &hasil_prediksi.TanggalPrediksi, &hasil_prediksi.NamaPasien, &hasil_prediksi.NamaPenyakit, &hasil_prediksi.TingkatKemiripan, &hasil_prediksi.StatusPrediksi); err != nil {
			return nil, fmt.Errorf("SOMEHOW ERROR: %v", err)
		}
		hasilPrediksi = append(hasilPrediksi, hasil_prediksi)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("SOMEHOW ERROR: %v", err)
	}
	defer rows.Close()
	return hasilPrediksi, nil
}
