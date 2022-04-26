import React from 'react';
import axios from 'axios';
import { Endpoints } from '../Api'
import moment from 'moment';
import "./TestDNA.scss"

class TestDNA extends React.Component {
  
  constructor(props) {
    super(props);
    this.state = {
      id: 999,
      nama_pengguna:'',
      rantai_dna: '',
      nama_penyakit:'',
      tanggal_prediksi: '',
      results:[],
      processing: false,
      doneProcess: false
    };
    
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleFile = this.handleFile.bind(this);
    this.getResult = this.getResult.bind(this);
  }
  

  handleFile = (file) => {
    var fileReader = new FileReader()
    fileReader.readAsText(file)
    fileReader.onloadend = (e)=> {
      var content = e.target.result
      const yourDate = new Date()
      const NewDate = moment(yourDate, 'YYYY-MM-DD')
      var tanggal = NewDate.toISOString().split('T')[0]
      console.log(tanggal)
      console.log(content)
      this.setState({
        rantai_dna : content,
        tanggal_prediksi: tanggal
      });
    }
  }

  getResult = () => {

    axios.get(Endpoints.hasilPrediksi)
        .then(res => {
          const records = res.data;
          const record = records[records.length-1]
          console.log(record)
          this.setState({
            results : record,
            doneProcess: true
          });
        })

  }


  handleSubmit = () => {
    
    console.log(this.state.id);
    console.log(this.nama_pengguna.value);
    console.log(this.state.rantai_dna);
    console.log(this.nama_penyakit.value);
    console.log(this.state.tanggal_prediksi);
    
    const formData = new FormData();
    
    // Update the formData object    
    formData.append("id", this.id);
    formData.append("nama_pengguna", this.nama_pengguna.value);
    formData.append("rantai_dna", this.state.rantai_dna);
    formData.append("nama_penyakit", this.nama_penyakit.value);
    formData.append("tanggal_prediksi", this.state.tanggal_prediksi);


    if(formData.values != null){
        
        console.log(Endpoints.testDNA)
        axios({
            method: 'POST',
            url: Endpoints.testDNA,
            headers: {
                'Content-Type': 'application/json'
            },
            data: {
                id: this.id,
                nama_pengguna: this.nama_pengguna.value,
                rantai_dna: this.state.rantai_dna,
                nama_penyakit: this.nama_penyakit.value,
                tanggal_prediksi: this.state.tanggal_prediksi
            }
        }).then((response) => {
          console.log(response);
        }, (error) => {
          console.log(error);
        });

        this.setState({
          processing: true
        })
    
    } else {

      console.log("error");

    } 
  }
  
  render() {
    return (
      <div className = 'card'>
        <div>
          <p>Periksa Rantai DNA</p>
        </div>
        <div>
          <div>
          <p>
            Nama Pengguna
          </p>
          <input type="text" required={true} ref={(ref) => {this.nama_pengguna = ref; }}/>
          </div>
          <div>
          <p>
            Rantai DNA
          </p>
          <input type="file" required={true} accept='.txt' onChange={e => this.handleFile(e.target.files[0])}/>
          </div>
          <div>
            <p>
              Nama Penyakit
            </p>
            <input type="text" required={true} ref={(ref) => {this.nama_penyakit = ref; }}/>
          </div>
          <div>
            <button onClick={this.handleSubmit} >Submit</button>
          </div>

          {this.state.processing ? (
                this.getResult(),
                this.setState({
                  processing: false
                })) : (<></>)
            }

          <div className={this.state.doneProcess ? "result-wrapper" : "result-wrapper-hidden"} >
            { 
              <p>{String(this.state.results.tanggal_prediksi).substring(0,10)}</p>
            }
            {
              <p>{this.state.results.nama_pasien}</p>
            }
            {
              <p>{this.state.results.nama_penyakit}</p>
            }
            {
              <p>{String(this.state.results.tingkat_kemiripan*100) + '%'}</p>
            }
            {
              <p>{String(this.state.results.status_prediksi)}</p>
            }
          </div>
        </div>
      </div>
    );
  }
}
export default TestDNA;