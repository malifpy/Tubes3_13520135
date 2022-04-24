import React from 'react';
import axios from 'axios';
import { Endpoints } from '../Api'
import "./TestDNA.scss"

class TestDNA extends React.Component {
  
  constructor(props) {
    super(props);
    this.state = {
      id: 999,
      nama_pengguna:'',
      rantai_dna: '',
      nama_penyakit:'',
      tanggal_prediksi: ''
    };
    
    this.handleSubmit = this.handleSubmit.bind(this);
    this.handleFile = this.handleFile.bind(this);
  }
  

  handleFile = (file) => {
    var fileReader = new FileReader()
    fileReader.readAsText(file)
    fileReader.onloadend = (e)=> {
      var content = e.target.result
      console.log(content)
      this.setState({
        rantai_dna : content
      });
    }
  }

  handleSubmit = () => {
    
    console.log(this.state.id.value);
    console.log(this.nama_pengguna.value);
    console.log(this.state.rantai_dna);
    console.log(this.nama_penyakit.value);
    console.log(this.state.tanggal_prediksi);
    
    const formData = new FormData();
    
    // Update the formData object    
    formData.append("id", this.id);
    formData.append("nama_pengguna", this.nama_pengguna);
    formData.append("rantai_dna", this.state.rantai_dna);
    formData.append("nama_penyakit", this.nama_penyakit);
    formData.append("tanggal_prediksi", this.state.tanggal_prediksi);


    if(formData.values != null){
        console.log(Endpoints.addPenyakit)
      
        axios({
            method: 'POST',
            url: Endpoints.testDNA,
            headers: {
                'Content-Type': 'application/json'
            },
            data: {
                id: this.id,
                nama_pengguna: this.nama_pengguna,
                rantai_dna: this.state.rantai_dna,
                nama_penyakit: this.nama_penyakit,
                tanggal_prediksi: this.tanggal_prediksi
            }

        }).then((response) => {
        console.log(response);
      }, (error) => {
        console.log(error);
      });
    
    } else {

      console.log("error");

    }
    
  }
  
  render() {
    return (
      <div class = "card">
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
            <button onClick={this.handleSubmit}>Submit</button>
          </div>
        </div>
      </div>
    );
  }
}
export default TestDNA;