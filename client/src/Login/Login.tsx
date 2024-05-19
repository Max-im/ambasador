import React, { useState } from 'react';
import axios from 'axios';
import { Link, useNavigate } from 'react-router-dom';
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import Stack from '@mui/material/Stack';
import Button from '@mui/material/Button';
import Divider from '@mui/material/Divider';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';


interface ILoginData {
  email: string;
  password: string;
}

export default function Login() {
  const navigate = useNavigate();
  const [loginData, setLoginData] = useState<ILoginData>({email: '', password: ''});

  const handleClick = () => {
    axios.post('http://localhost:5000/login', loginData).then(() => navigate('/'));
  };

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setLoginData(state => ({
      ...state,
      [name]: value,
    }));
  };

  return (
    <Box sx={{height: 1}}>
      <Stack alignItems="center" justifyContent="center" sx={{ height: 1 }}>
        <Card
          sx={{
            p: 5,
            width: 1,
            maxWidth: 420,
          }}
        >
          <Typography variant="h4">Sign in to Ambassador</Typography>

          <Typography variant="body2" sx={{ mt: 2, mb: 5 }}>
            Don’t have an account?
            <Link to="/register">
              <Typography variant="subtitle2" component="span" sx={{ ml: 0.5 }}>Register</Typography>
            </Link>
          </Typography>

          
          <Stack spacing={3}>
            <TextField name="email" label="Email address" onChange={onChange}/>

            <TextField
              name="password"
              label="Password"
              type="password"
              onChange={onChange}
            />
          </Stack>

          <Divider sx={{mt:2, mb:2}} />

          <Button
            fullWidth
            size="large"
            type="submit"
            variant="contained"
            color="primary"
            onClick={handleClick}
          >
            Login
          </Button>
        </Card>
      </Stack>
    </Box>
  );
}