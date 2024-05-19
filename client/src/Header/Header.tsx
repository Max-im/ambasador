import React from 'react'
import { Link } from 'react-router-dom';
import { Box, Stack, Typography } from '@mui/material';

export default function Header() {
  return (
    <Box sx={{
      display: 'flex',
      justifyContent: 'space-between',
      alignItems: 'center',
      backgroundColor: '#333',
      color: '#fff',
      padding: '1rem',
      '& a': {
        color: '#fff',
        textDecoration: 'none',
        '&:hover': {
          backgroundColor: '#555',
        }
      }
    }}>
      <Typography variant="h6">Header</Typography>
      <Stack direction="row" spacing={2}>
        <Link to="/login">Login</Link>
        <Link to="/users">Users</Link>
        <Link to="/logout">Logout</Link>
        <Link to="/products">Products</Link>
      </Stack>
    </Box>
  )
}
