import React from 'react';
import axios from 'axios';
import { Endpoints } from '../Api'
import "./TestDNA.scss"

class TestDNA extends React.Component {
  
  constructor(props) {
    super(props);
    this.state = {
      id: 0,
      nama_pengguna:'',
      rantai_dna: '',
      nama_penyakit:''
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
    
    console.log(this.nama_penyakit.value);
    console.log(this.state.rantai_dna);
    
    const formData = new FormData();
    
    // Update the formData object    
    formData.append("id", this.id);
    formData.append("nama", this.nama_penyakit);
    formData.append("rantai_dna", this.state.rantai_dna);


    if(formData.values != null){
        console.log(Endpoints.addPenyakit)
      
        axios({
            method: 'POST',
            url: Endpoints.addPenyakit,
            headers: {
                'Content-Type': 'application/json'
            },
            data: {
                id: this.id,
                nama: this.nama_penyakit.value,
                rantai_dna: this.state.rantai_dna
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
      <div>
        <ul>
          <label>
            nama pengguna
            <input type="text" required={true} ref={(ref) => {this.nama_pengguna = ref; }}/>
          </label>
        </ul>
        <ul>
          <label>
            rantai dna
            <input type="file" required={true} accept='.txt' onChange={e => this.handleFile(e.target.files[0])}/>
          </label>
        </ul>
        <ul>
          <label>
            nama penyakit
            <input type="text" required={true} ref={(ref) => {this.nama_penyakit = ref; }}/>
          </label>
        </ul>
        <ul>
          <button onClick={this.handleSubmit}>Submit</button>
        </ul>
      </div>
    );
  }
}
export default TestDNA;