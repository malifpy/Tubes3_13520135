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
          <p> Cari Hasil Prediksi</p>
        </div>
        <div>
          <div>
            <p>
            Query
            </p>
            <input type="text" required={true} ref={(ref) => {this.nama_penyakit = ref; }}/>
            <div >
              <button onClick={this.handleSubmit}>Cari</button>
            </div>
          </div>
        </div>
      </div>
    );
  }
}

export default HasilPrediksi;