import { Box, Button, Card, Stack, TextField, Typography } from '@mui/material';
import axios from 'axios';
import React, { useEffect, useState } from 'react';

export default function Account() {
  const [user, setUser] = useState({ first_name: '', last_name: ''});
  const [pass, setPass] = useState({ password: '', password_confirm: '' });

  useEffect(() => {
    axios.get('/me').then(({ data }) => setUser(data));
  }, []);


  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setUser(state => ({
      ...state,
      [name]: value,
    }));
  };

  const onChangePass = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setPass(state => ({
      ...state,
      [name]: value,
    }));
  };

  const onUpdatePassword = () => {
    axios.put('/user/password', pass);
  }

  const onUpdateUser = () => {
    axios.put('/user/info', user);
  }

  
  return (
    <Box sx={{ height: 1 }}>
      <Stack alignItems="center" justifyContent="center" sx={{ height: 1 }}>
        <Card
          sx={{
            p: 5,
            width: 1,
            maxWidth: 600,
          }}
        >
          <Typography variant="h4">Your Profile</Typography>

          <Stack direction="row" alignItems="center" justifyContent="space-between" spacing={2} sx={{ mt: 2, mb: 5 }}>
            <Typography variant="body2">
            First name:
            </Typography>
            <TextField
              value={user.first_name}
              name="first_name"
              label="First name"
              onChange={onChange}
              fullWidth
              variant="outlined"
              margin="normal"
              InputLabelProps={{
                shrink: true,
              }}
              sx={{
                width: '100%',
                maxWidth: '450px'
              }}
            />
          </Stack>

          <Stack direction="row" alignItems="center" justifyContent="space-between" spacing={2} sx={{ mt: 2, mb: 5 }}>
            <Typography variant="body2">
              Last Name:
            </Typography>
            <TextField
              value={user.last_name}
              name="last_name"
              label="Last Name"
              onChange={onChange}
              fullWidth
              variant="outlined"
              margin="normal"
              InputLabelProps={{
                shrink: true,
              }}
              sx={{
                width: '100%',
                maxWidth: '450px'
              }}
            />
          </Stack>


          <Stack spacing={3}>
            <Button
              fullWidth
              size="large"
              type="submit"
              variant="contained"
              color="primary"
              onClick={onUpdateUser}
            >
              Edit Profile
            </Button>
          </Stack>
        </Card>

        <Card
          sx={{
            p: 5,
            width: 1,
            maxWidth: 600,
          }}
        >
          <Typography variant="h4">Update Password</Typography>

          <Stack spacing={3}>
            <TextField
                name="password"
                label="Password"
                type="password"
                onChange={onChangePass}
              />
            
          </Stack>

          <Stack spacing={3}>
          <TextField
                name="password_confirm"
                label="Confirm Passworf"
                type="password"
                onChange={onChangePass}
                />
                </Stack>
            

          <Stack spacing={3}>
            <Button
              fullWidth
              size="large"
              type="submit"
              variant="contained"
              color="primary"
              onClick={onUpdatePassword}
            >
              Edit Profile
            </Button>
          </Stack>
        </Card>
      </Stack>
    </Box>
  );

}
