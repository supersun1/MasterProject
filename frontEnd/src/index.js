import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import { BrowserRouter, Switch, Route } from 'react-router-dom'
import App from './App';
import reportWebVitals from './reportWebVitals';
import './styles'
import './scripts'
import IndexView from './views/IndexView'
import ResultsView from './views/ResultsView'
import IndexController from './controllers/IndexController'
import ResultsController from './controllers/ResultsController'
import AboutView from './views/AboutView'
import ApiView from './views/ApiView'
import CommunityView from './views/CommunityView'
import 'bootstrap/dist/css/bootstrap.min.css';



ReactDOM.render(
   <BrowserRouter>
    <Switch>
      <Route exact path="/index.html" component={IndexController} />
      <Route exact path="/results" component={ResultsController} />
      <Route exact path='/about.html' component={AboutView} />
      <Route exact path='/community.html' component={CommunityView} />
      <Route exact path='/api.html' component={ApiView} />
    </Switch>
   </BrowserRouter>
  ,document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();