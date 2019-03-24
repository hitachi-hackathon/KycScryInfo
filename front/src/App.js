import React, { Component } from 'react';
import {
  Route,
  Switch
} from 'react-router-dom';
import NavBar from './nav'
import Footer from './footer'
// import css 
import '../node_modules/bootstrap/dist/css/bootstrap.min.css';
import './CSS/main.css';
import "./CSS/process.css"

// modules
import Home from "./home"
import Certify from "./validator"
import Institute from "./institute"
import User from "./user"
import Verify from "./authVerify"


class App extends Component {

  render() {
    return (
      <div>
          <NavBar />
          <Switch>
            <Route path="/authority/certify" component={Certify} />
            <Route path="/authority/verify" component={Verify} />
            <Route path="/institute" component={Institute} />
            <Route path="/user" component={User} />
            <Route path="/" component={Home} />
          </Switch>
          <Footer />
      </div>
    );
  }
}

export default App;
