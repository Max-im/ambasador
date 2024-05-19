import React from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { Box } from '@mui/material';
import Header from './Header/Header';
import Login from './Login/Login';
import Register from './Register/Register';
import Home from './Home/Home';
import Account from './Account/Account';
import Logout from './Logout/Logout';
import Users from './Users/Users';

function App() {
  return (
      <BrowserRouter>
    <Box className="App">
      <Header />
        <Routes>
          <Route path={'/'} element={<Home />} />
          <Route path={'/login'} element={<Login />} />
          <Route path={'/register'} element={<Register />} />
          <Route path={'/account'} element={<Account />} />
          <Route path={'/logout'} element={<Logout />} />
          <Route path={'/users'} element={<Users />} />
        </Routes>
    </Box>
      </BrowserRouter>
  );
}

export default App;
