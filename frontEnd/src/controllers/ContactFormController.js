import React, { useState } from 'react'
import ContactFormView from '../views/ContactFormView'
import axios from 'axios'
import { CONFIG } from '../Constants';
import * as ReactBootStrap from 'react-bootstrap';

class ContactFormController extends React.Component {
//  state = {}

constructor(props) {
  super(props)

  this.state = {
    name: '',
    isLoading: true,
    results: '',
    confidence_score: ''
  }
}

setName = e => {
  this.setState({ [e.target.name]: e.target.value })
}

setResults = e => {
  this.setState({ [e.target.confidence_score]: e.target.value })
}

submit = async e => {
  
  e.preventDefault();
  console.log(this.state)
  
 

  const settings = {
    method: 'POST',
    body: JSON.stringify(this.state),
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json',
    }
  }
 
  const getData = await fetch(`${CONFIG.APIBaseUrl}`, settings);
  const data = await getData.json();
  //const [loading, setLoading] = useState(false);

  this.setState({
    confidence_score: data.body.confidence_score,
    results: data.body.text_determination
  });
  
 
  console.log(this.state.confidence_score); 
  {
    this.props.history.push('/results',{
      confidence_score: this.state.confidence_score,
      results: this.state.results
     })}
} 


//    this.props.history.push('/results',{
//     results: this.state.results
   
    
//  })
// }

//   componentDidMount(){
//     const getResults = {
//         method: 'POST',
//         headers: { 'Content-Type': 'application/json' },
//         body: JSON.stringify(this.state)
//     };
//     fetch(`${CONFIG.APIBaseUrl}`, getResults)
//         .then(response => response.json())
//         .then(data => this.setState({results: data.body}));
            
  // }

  render() {
    return (
      <ContactFormView>
       <name onChange={this.setName} />
        <submit onClick={this.submit} />
        <results onChange={this.setResults} />
      </ContactFormView>

    )
    }}
  
export default ContactFormController