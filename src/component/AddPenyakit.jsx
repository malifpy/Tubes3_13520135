import React from 'react';
import axios from 'axios';
import { Endpoints } from '../Api'
import "./AddPenyakit.scss"

class AddPenyakit extends React.Component {
  
  constructor(props) {
    super(props);
    this.state = {
      id: 0,
      nama_penyakit:'',
      rantai_dna: '',
      status: "data penyakit berhasil ditambahkan",
      doneProcess: false
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
        console.log(error.message);
        this.setState({
          status: "Rantai DNA tidak valid"
        })
      });
    
    } else {

      console.log("error");

    }

    this.setState({
      doneProcess: true
    })
    
  }
  
  render() {
    return (
      <div className='card'> 
        <div>
          <p> Tambah Penyakit </p>
        </div>
        <div>
          <div>
          <p>
            Nama Penyakit          
          </p>
          <input type="text" required={true} ref={(ref) => {this.nama_penyakit = ref; }}/>
          </div>
          <div>
            <p>
              Rantai DNA
            </p>
            <input type="file" required={true} accept='.txt' onChange={e => this.handleFile(e.target.files[0])}/>
          </div>
          <div>
            {/* <input type="submit" value="Submit" /> */}
            <button onClick={this.handleSubmit}>Submit</button>
          </div>
          <div className={this.state.doneProcess ? "result-wrapper" : "result-wrapper-hidden"}>
            <p>{this.state.status}</p>
          </div>
        </div>
      </div>
    );
  }
}
export default AddPenyakit;
