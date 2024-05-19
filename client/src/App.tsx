import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Box } from '@mui/material';
import Header from './Header/Header';
import Login from './Login/Login';
import Register from './Register/Register';
import Home from './Home/Home';

function App() {
  return (
      <BrowserRouter>
    <Box className="App">
      <Header />
        <Routes>
          <Route path={'/'} element={<Home />} />
          <Route path={'/login'} element={<Login />} />
          <Route path={'/register'} element={<Register />} />
        </Routes>
    </Box>
      </BrowserRouter>
  );
}

export default App;
