import React from 'react';
import {Switch,Route,HashRouter as Router} from 'react-router-dom';
import Login from "./pages/login" 
import "./App.css"
import Chat from './pages/chat';


function App() {
  return (
    <Router>
      <Switch>
        {/* <Route path="/" component={Chat}/> */}
        <Route path="/login" component={Login} />
        <Route path="/chat" component={Chat} />
      </Switch>
    </Router>
  );
}

export default App;
