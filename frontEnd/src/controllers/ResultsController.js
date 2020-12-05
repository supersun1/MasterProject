import React from 'react'
import ResultsView from '../views/ResultsView'
import ContactFormController from '../controllers/ContactFormController'
import { render } from '@testing-library/react'
import { CONFIG } from '../Constants';


// class ResultsController extends React.Component {
//   //state = {}


//   constructor(props) {
//     super(props)
  
//     this.state = {
//       name: ''
//     }
//   }

//  componentDidMount(){
//     const getResults = {
//         method: 'POST',
//         headers: { 'Content-Type': 'application/json' },
//         body: JSON.stringify(this.state)
//     };
//     fetch(`${CONFIG.APIBaseUrl}`, getResults)
//         .then(response => response.json())
//         .then(data => this.setState({results: data.body}));
//   }

//   render() {
//     return (
     
//     <ResultsView>
//    <message> Approximately accurate </message>
//     </ResultsView>
      
//     )

//   }}
//   export default ResultsController

export default (props) => (
  <ResultsView>
   
   <message> This text is {props.location.state.results} with {props.location.state.confidence_score} confidence </message>
    
  </ResultsView>
)