import React from 'react';

import "./AddPenyakit.scss"

class AddPenyakit extends React.Component {
  
    constructor(props) {
        super(props);
        this.state = {value: ''};
    
        this.handleChange = this.handleChange.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
      }
    
      handleChange(event) {
        this.setState({value: event.target.value});
      }
    
      handleSubmit(event) {
        alert('A name was submitted: ' + this.state.value);
        event.preventDefault();
      }
    
      render() {
        return (
          <form onSubmit={this.handleSubmit}>
            <ul>
              <label>
                nama penyakit
                <input type="text"/>
              </label>
            </ul>
            <ul>
              <label>
                rantai dna
                <input type="file" value = {this.state.value} onChange={this.handleChange} />
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