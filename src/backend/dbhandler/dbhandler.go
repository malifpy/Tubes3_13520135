package dbhandler

import (
	"database/sql"
	smalgorithm "dna-matcher/sm_algorithm"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type JenisPenyakit struct {
	ID        int64  `json:"id"`
	Nama      string `json:"nama"`
	RantaiDNA string `json:"rantai_dna"`
}

type HasilPrediksi struct {
	ID               int64   `json:"id"`
	TanggalPrediksi  string  `json:"tanggal_prediksi"`
	NamaPasien       string  `json:"nama_pasien"`
	NamaPenyakit     string  `json:"nama_penyakit"`
	TingkatKemiripan float64 `json:"tingkat_kemiripan"`
	StatusPrediksi   bool    `json:"status_prediksi"`
}

func InsertJenisPenyakit(db *sql.DB, rDNA JenisPenyakit) error {

	var err1 error
	sqlQuery := `INSERT INTO jenis_penyakit (nama, rantai_dna) VALUES($1, $2) RETURNING id;`
	id := 0
	dnaSanitized, _ := smalgorithm.IsValid(rDNA.RantaiDNA)

	if dnaSanitized {
		err := db.QueryRow(sqlQuery, rDNA.Nama, rDNA.RantaiDNA).Scan(&id)
		if err != nil {

			return fmt.Errorf("SOMEHOW ERROR: %v", err)

		}
	} else {

		err1 = errors.New("rantai dna tidak valid")
	}

	return err1

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

func ViewHasilPrediksiByDate(db *sql.DB, Date string) ([]HasilPrediksi, error) {
	var hasilPrediksi []HasilPrediksi
	sqlQuery := `SELECT * FROM hasil_prediksi WHERE tanggal_prediksi = $1;`
	rows, err := db.Query(sqlQuery, Date)
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

func ViewHasilPrediksiByName(db *sql.DB, Nama string) ([]HasilPrediksi, error) {
	var hasilPrediksi []HasilPrediksi
	sqlQuery := `SELECT * FROM hasil_prediksi WHERE nama_penyakit = $1;`
	rows, err := db.Query(sqlQuery, Nama)
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

func ViewHasilPrediksiByNameNDate(db *sql.DB, Nama string, Tanggal string) ([]HasilPrediksi, error) {
	var hasilPrediksi []HasilPrediksi
	sqlQuery := `SELECT * FROM hasil_prediksi WHERE nama_penyakit = $1 and tanggal_prediksi = $2;`
	rows, err := db.Query(sqlQuery, Nama, Tanggal)
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
