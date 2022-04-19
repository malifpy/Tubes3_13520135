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
      rantai_dna:''
    };
    
    this.handleSubmit = this.handleSubmit.bind(this);
  }
    
  handleSubmit = () => {
    
    console.log(this.nama_penyakit.value);
    console.log(this.rantai_dna.value);
    
    const formData = new FormData();
    
    // Update the formData object    
    formData.append("id", this.id);
    formData.append("nama", this.nama_penyakit);
    formData.append("rantai_dna", this.rantai_dna);


    if(formData.values != null){
      
      axios.post(Endpoints.addPenyakit, formData).then((response) => {
        console.log(response);
      }, (error) => {
        console.log(error);
      });
    
    } else {

      console.log("error");

    }
    // Request made to the backend api
    // Send formData object
  }
  
  render() {
    return (
      <div>
        <ul>
          <label>
            nama penyakit
            <input type="text" required={true} ref={(ref) => {this.nama_penyakit = ref; }}/>
          </label>
        </ul>
        <ul>
          <label>
            rantai dna
            <input type="text" required={true} ref={(ref) => {this.rantai_dna = ref; }}/>
          </label>
        </ul>
        <ul>
          <button onClick={this.handleSubmit}>Submit</button>
        </ul>
      </div>
    );
  }
}
export default AddPenyakit;