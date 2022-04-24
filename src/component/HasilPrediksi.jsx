import React from 'react';
import axios from 'axios';
import { Endpoints } from '../Api'
import "./HasilPrediksi.scss"

class HasilPrediksi extends React.Component {
  
  constructor(props) {
    super(props);
    this.state = {
      search_query:'',
      result_query:[]
    };
    
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  
  handleSubmit = () => {
    
    console.log(this.search_query);

    axios.get(Endpoints.hasilPrediksi)
      .then(res => {
        const hasil = res.data;
        this.setState({
          result_query : hasil
        });
      })

    console.log(this.state.result_query);

  }
  
  render() {
    return (
      <div className = 'card'>
        <div>
          <p> Cari Hasil Prediksi</p>
        </div>
        <div>
          <div>
            <p>
            Query
            </p>
            <input type="text" required={true} ref={(ref) => {this.search_query = ref; }}/>
            <div >
              <button onClick={this.handleSubmit}>Cari</button>
            </div>
            <ul>
              {
                this.state.result_query.map(result =>
                  <li key={result.id}>{result.nama_penyakit}</li>
                )
              }
            </ul>
          </div>
        </div>
      </div>
    );
  }
}

export default HasilPrediksi;