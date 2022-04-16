import React from 'react';
import { Endpoints } from '../Api'
import "./AddPenyakit.scss"

class AddPenyakit extends React.Component {
  
  constructor(props) {
    super(props);
    this.state = {
      nama_penyakit: '',
      rantai_dna: ''
    };
    
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
      
  handleChange(event) {
    this.setState({value: event.target.value});
  }
    
  handleSubmit = async (e) => {
    e.preventDefault()
    try {
      const res = await fetch(Endpoints.addPenyakit, {
        method: 'POST',
        body: JSON.stringify({ }),
            headers: {
              'Content-Type': 'application/json',
            },
        })
        const { success, errors = [] } = await res.json()
          setErrors(errors)
    } catch (e) {
      setErrors([e.toString()]);
    }
  }
  
  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <ul>
          <label>
            nama penyakit
            <input type="text" required="true"/>
          </label>
        </ul>
        <ul>
          <label>
            rantai dna
            <input type="file" value = {this.state.value} required="true" onChange={this.handleChange} />
          </label>
        </ul>
        <ul>
          <input type="submit" value="Submit" />
        </ul>
      </form>
    );
  }
}
export default AddPenyakit;