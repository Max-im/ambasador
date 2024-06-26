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
import UserLinks from './UserLinks/UserLinks';
import Products from './Products/Products';
import CreateProduct from './Products/CreateProduct';
import EditProduct from './Products/EditProduct';
import Order from './Order/Order';

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
          <Route path={'/users/:userId/links'} element={<UserLinks />} />
          <Route path={'/products'} element={<Products />} />
          <Route path={'/product/create'} element={<CreateProduct />} />
          <Route path={'/product/:productId/edit'} element={<EditProduct />} />
          <Route path={'/order'} element={<Order />} />
        </Routes>
    </Box>
      </BrowserRouter>
  );
}

export default App;
