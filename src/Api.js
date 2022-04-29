//export const apiURl = 'https://cors-anywhere.herokuapp.com/https://tubes3-dna-matcher.herokuapp.com'
// export const apiURL = 'http://localhost:8080'
export const apiURL = 'https://tubes3-dna-matcher.herokuapp.com'

export const Endpoints = {
  addPenyakit: `${apiURL}/jenis_penyakit`,
  testDNA:  `${apiURL}/pasien`,
  hasilPrediksi: `${apiURL}/hasil_prediksi`,
  query: `${apiURL}/data`
}
