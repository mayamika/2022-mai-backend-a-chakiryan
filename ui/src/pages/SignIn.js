import React, { useContext, useState } from 'react';

import {
  Box,
  Typography,

  Avatar,
  Container,
  Grid,
  TextField,
  Button,
  Link,
} from '@mui/material';
import { LockOutlined } from '@mui/icons-material';

import { gql, useMutation } from '@apollo/client';
import { SessionContext } from '../session';
import { useNavigate } from 'react-router-dom';

const LOGIN = gql`
  mutation Login($input: LoginInput!) {
    login(input: $input) {
      token
    }
  }
`;

function SignIn() {
  const [login, setLogin] = useState('');
  const [pass, setPass] = useState('');

  const change = (setFn) => {
    return ((e) => {
      e.preventDefault();
      setFn(e.target.value);
    });
  };

  const [signIn, { error }] = useMutation(LOGIN, {
    variables: {
      input: {
        login: login,
        password: pass,
      },
    },
  });
  if (error) console.log(`:( ${error.message}`);

  const [, setSession] = useContext(SessionContext);
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();
    signIn().then((res) => {
      const { token } = res.data.login;
      console.log(token);

      setSession({
        token: token,
      });
      navigate('/');
    });
  };

  return (
    <Container maxWidth='xs'>
      <Box
        sx={{
          marginTop: 8,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
          <LockOutlined />
        </Avatar>
        <Typography component="h1" variant="h5">
          Sign in
        </Typography>
        <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 3 }}>
          <Grid container spacing={2}>
            <Grid item xs={12}>
              <TextField
                required
                fullWidth
                id="username"
                label="Username"
                name="username"
                autoComplete="username"
                autoFocus
                value={login}
                onChange={change(setLogin)}
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                required
                fullWidth
                name="password"
                label="Password"
                type="password"
                id="password"
                autoComplete="new-password"
                value={pass}
                onChange={change(setPass)}
              />
            </Grid>
          </Grid>
          <Button
            type="submit"
            fullWidth
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            Sign In
          </Button>
          <Grid container justifyContent="flex-end">
            <Grid item>
              <Link href="/signup" variant="body2">
                {'Don\'t have an account? Sign up'}
              </Link>
            </Grid>
          </Grid>
        </Box>
      </Box>
    </Container>
  );
}

export default SignIn;
