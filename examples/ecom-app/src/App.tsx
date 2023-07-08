import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import Home from './pages/Home';

function App() {
  return (
    <div className="App">
      <Router>
        <Switch>
          <Route>

          </Route>
          <Route  path='/' component={Home} exact/>
        </Switch>
      </Router>
    </div>
  );
}

export default App;
