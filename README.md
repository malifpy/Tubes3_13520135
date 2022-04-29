# Tugas Besar 3 IF2211 Strategi Algoritma Penerapan String Matching dan Regular Expression dalam DNA Pattern Matching

## Deskripsi Singkat

Web aplikasi DNA pattern matching adalah web aplikasi yang berguna untuk memprediksi penyakit yang diderita oleh seseorang dengan memproses rantai DNA dari orang tersebut. Web aplikasi ini memanfaatkan  algoritma pattern matching knuth-morris-pratt untuk menentukan tingkat kesamaan rantai DNA yang dimiliki oleh seseorang dengan rantai DNA suatu penyakit yang terdapat di basis data, selain itu algoritma pattern matching Boyer-Moore juga berhasil di implementasikan dengan baik. 

Web aplikasi ini memiliki beberapa fitur sebagai berikut:
  - Menambahkan jenis penyakit
  - Memprediksi penyakit pengguna
  - Mencari hasil prediksi penyakit berdasarkan tanggal dilakukannya prediksi
  - Mencari hasil prediksi penyakit berdasarkan nama penyakit 

## Requirements

## How to compile and run

### Lokal

Frontend
- Pergi ke folder my-app
- Ubah variable ``` apiURL ``` pada file ``` Api.js ``` dengan 'http://localhost:8080'
- Simpan perubahan
- Pada terminal jalankan command
``` npm install ```
kemudian
``` npm start ```

Backend
- Pergi ke folder backend
- Pada terminal jalankan command
``` go build ```
kemudian
``` export DATABASE_URL=postgres://oyukhvnsxohoku:4e4319499fb3816f3d229893806cbfde848caa6550e31798bd2895be581f9bf3@ec2-34-207-12-160.compute-1.amazonaws.com:5432/de4m84cglucnno ``` 
kemudian
``` ./dna-matcher ```

## Authors

| NIM      | NAMA                          | Pembagian Tugas                                                  |	
|----------|-------------------------------|------------------------------------------------------------------|
| 13520135 | Muhammad Alif Putra Yasa      | Implementasi Algoritma KMP, BM, Regex, Backend, Deploy           |
| 13520158 | Azmi Alfatih Shalahuddin      |                                                                  |
| 13520165 | Ghazian Tsabit Alkamil        | Impelementasi semua fitur pada Frontend, Backend, Database       |
