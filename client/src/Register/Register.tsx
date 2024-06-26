import { useState } from 'react';
import axios from 'axios';
import { Link, useNavigate } from 'react-router-dom';
import Box from '@mui/material/Box';
import Card from '@mui/material/Card';
import Stack from '@mui/material/Stack';
import Button from '@mui/material/Button';
import Divider from '@mui/material/Divider';
import TextField from '@mui/material/TextField';
import Typography from '@mui/material/Typography';

interface IRegisterData {
  email: string;
  password: string;
  first_name: string;
  last_name: string;
  password_confirm: string;
}

export default function Register() {
  const navigate = useNavigate();
  const [registerData, setRegisterData] = useState<IRegisterData>({email: '', password: '', first_name: '', last_name: '', password_confirm: ''});

  const handleClick = () => {
    axios.post('/register', registerData).then(() => navigate('/login'));
  };

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setRegisterData(state => ({
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
          <Typography variant="h4">Sign up to Ambassador</Typography>

          <Typography variant="body2" sx={{ mt: 2, mb: 5 }}>
            Already have an account?
            <Link to="/login">
              <Typography variant="subtitle2" component="span" sx={{ ml: 0.5 }}>Login</Typography>
            </Link>
          </Typography>

          
          <Stack spacing={3}>
            <TextField name="email" label="Email address" />
            <TextField name="first_name" label="First Name" />
            <TextField name="last_name" label="Last Name" />

            <TextField
              name="password"
              label="Password"
              type="password"
            />
            <TextField
              name="password_confirm"
              label="Confirm Password"
              type="password"
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
            Register
          </Button>
        </Card>
      </Stack>
    </Box>
  );
}