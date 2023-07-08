import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import Home from './pages/Home';
import Header from './components/Header';
import Cart from './pages/Cart';

import './App.css'
function App() {
  return (
    <div className="App">
      <Router>
        <Header />
        <Routes>
          <Route path='/cart' Component={Cart} />
          <Route  path='/' Component={Home} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
