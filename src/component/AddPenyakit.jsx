import React from 'react';
import axios from 'axios';
// import { Endpoints } from '../Api'
import "./AddPenyakit.scss"

class AddPenyakit extends React.Component {
  
  constructor(props) {
    super(props);
    this.state = {
      nama_penyakit:'',
      rantai_dna:''
    };
    
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
      
  handleChange (event) {
    this.setState({nama_penyakit: event.target.nama_penyakit});
    this.setState({rantai_dna: event.target.rantai_dna});
  }
    
  handleSubmit = () => {
    const formData = new FormData();
    
    // Update the formData object
    console.log(this.nama_penyakit);
    console.log(this.rantai_dna);
    
    formData.append(
      this.nama_penyakit,
      this.rantai_dna
    );
    
    // Details of the uploaded file
    console.log(this.state.nama_penyakit);
    console.log(this.state.rantai_dna);
    
    // Request made to the backend api
    // Send formData object
    axios.post(`https://tubes3-dna-matcher.herokuapp.com/jenis_penyakit`, formData);
  }
  
  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <ul>
          <label>
            nama penyakit
            <input type="text" required={true} value = {this.nama_penyakit} onChange={this.handleChange}/>
          </label>
        </ul>
        <ul>
          <label>
            rantai dna
            <input type="text" required={true} value = {this.rantai_dna} onChange={this.handleChange} />
          </label>
        </ul>
        <ul>
          <button onClick={this.handleSubmit}>Submit</button>
        </ul>
      </form>
    );
  }
}
export default AddPenyakit;