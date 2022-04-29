import React from 'react';
import axios from 'axios';
import { Endpoints } from '../Api'
import "./HasilPrediksi.scss"

class HasilPrediksi extends React.Component {
  
  constructor(props) {
    super(props);
    this.state = {
      id: 0,
      search_query:'',
      result_query:[],
      doneProcess: false,
      processing:false,
      status: null
    };
    
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  
  getResult = () => {

    axios.get(Endpoints.query)
        .then(res => {
          const records = res.data;
          
          if (res.status === 204) {
            this.setState({
              status: "Data tidak ditemukan !",
              processing: false
            })
          } else {
            this.setState({
              result_query : records,
            });
          }
        })
        if (this.state.status === null){
          this.setState({
            processing: true
          })
        }
  }


  handleSubmit = (e) => {
    
    e.preventDefault()
    axios({
      method: 'POST',
      url: Endpoints.query,
      headers: {
          'Content-Type': 'application/json'
      },
      data: {
          id: this.id,
          query: this.search_query.value
      }
    }).then((response) => {
      console.log(response);
  
      this.getResult()

    }, (error) => {
      console.log(error.message);
    });
  }
  
  render() {
    return (
      <form onSubmit={this.handleSubmit} className = 'card'>
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
              <input type="submit"  value="Submit"/>
            </div>
            {this.state.processing ? (
              (
              <ul>
              {
                this.state.result_query.map(result =>
                  <li key={result.id}>
                    <p>
                      {String(result.tanggal_prediksi).substring(0,10) +"-"+String(result.nama_pasien)+"-"+String(result.nama_penyakit)+"-"+String(result.tingkat_kemiripan*100)+"% -"+String(result.status_prediksi)}
                    </p>
                  </li>
                )
              }
              </ul>
              )
            ):(<p>{this.state.status}</p>)
            }
          </div>
        </div>
      </form>
    );
  }
}

export default HasilPrediksi;