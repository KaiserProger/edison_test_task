import React from 'react';
import './App.css';
import { Index } from './pages/Index';
import { Routes, Route } from 'react-router-dom';
import { History } from './pages/History';
import { GuessesPage } from './pages/GuessesPage';
import { ExtrasenseTrustPage } from './pages/Trust';

function App() {
  return (
    <div className="App">
      <Routes>
        <Route path="" element={<Index></Index>}></Route>
        <Route path="/poll" element={<GuessesPage></GuessesPage>}></Route>
        <Route path="/history" element={<History></History>}></Route>
        <Route path="/trust" element={<ExtrasenseTrustPage></ExtrasenseTrustPage>}></Route>
      </Routes>
    </div>
  );
}

export default App;
