import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import Home from './pages/Home';

import './App.css'
function App() {
  return (
    <div className="App">
      <Router>
        <Routes>
          <Route>

          </Route>
          <Route  path='/' Component={Home} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
