import React from 'react';
import axios from 'axios';
import { Endpoints } from '../Api'
import "./HasilPrediksi.scss"

class HasilPrediksi extends React.Component {
  
  constructor(props) {
    super(props);
    this.state = {
      search_query:''
    };
    
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  
  handleSubmit = () => {
    
    console.log(this.state.search_query);
    
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
export default HasilPrediksi;