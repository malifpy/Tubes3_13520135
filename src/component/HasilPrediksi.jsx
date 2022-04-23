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
      <div class = "card">
        <div>
          <p>
            nama pengguna
          </p>
          <div>
            <input type="text" required={true} ref={(ref) => {this.nama_pengguna = ref; }}/>
          </div>
          <p>
            rantai dna
          </p>
          <div>
            <input type="file" required={true} accept='.txt' onChange={e => this.handleFile(e.target.files[0])}/>
          </div>
          <p>
            nama penyakit
          </p>
          <div>
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
export default HasilPrediksi;