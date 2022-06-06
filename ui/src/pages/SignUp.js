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

const REGISTER = gql`
  mutation Register($input: RegisterInput!) {
    register(input: $input) {
      token
    }
  }
`;

function SignUp() {
  const [login, setLogin] = useState('');
  const [password, setPassword] = useState('');
  const [email, setEmail] = useState('');
  const [name, setName] = useState('');
  const [surname, setSurname] = useState('');

  const loginRe = /^[a-zA-Z][a-zA-Z0-9]*$/;
  const loginValid = loginRe.test(login);

  const nameRe = /^([A-ZА-Я][a-zа-я]*)+$/u;
  const nameValid = nameRe.test(name);
  const surnameValid = nameRe.test(surname);

  /* eslint-disable max-len */
  const emailRe = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
  /* eslint-enable max-len */
  const emailValid = emailRe.test(email);

  const filled = nameValid && surnameValid && emailValid && password;

  const onChange = (setFn) => {
    return ((e) => {
      e.preventDefault();
      setFn(e.target.value);
    });
  };

  const [signUp, { error }] = useMutation(REGISTER, {
    variables: {
      input: {
        login: login,
        password: password,
        email: email,
        name: name,
        surname: surname,
      },
    },
  });
  if (error) {
    console.log(error);
  }

  const [, setSession] = useContext(SessionContext);
  const navigate = useNavigate();

  const handleSubmit = (e) => {
    e.preventDefault();
    signUp().then((res) => {
      const { token } = res.data.register;

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
          Sign up
        </Typography>
        <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 3 }}>
          <Grid container spacing={2}>
            <Grid item xs={12} sm={6}>
              <TextField
                autoComplete="given-name"
                name="firstName"
                required
                fullWidth
                id="firstName"
                label="First Name"
                autoFocus
                value={name}
                error={!nameValid && name}
                onChange={onChange(setName)}
              />
            </Grid>
            <Grid item xs={12} sm={6}>
              <TextField
                required
                fullWidth
                id="lastName"
                label="Last Name"
                name="lastName"
                autoComplete="family-name"
                value={surname}
                error={!surnameValid && surname}
                onChange={onChange(setSurname)}
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                required
                fullWidth
                id="username"
                label="Username"
                name="username"
                autoComplete="username"
                value={login}
                error={!loginValid && login}
                onChange={onChange(setLogin)}
              />
            </Grid>
            <Grid item xs={12}>
              <TextField
                required
                fullWidth
                id="email"
                label="Email Address"
                name="email"
                autoComplete="email"
                value={email}
                error={!emailValid && email}
                onChange={onChange(setEmail)}
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
                value={password}
                onChange={onChange(setPassword)}
              />
            </Grid>
          </Grid>
          <Button
            type="submit"
            fullWidth
            disabled={!filled}
            variant="contained"
            sx={{ mt: 3, mb: 2 }}
          >
            Sign Up
          </Button>
          <Grid container justifyContent="flex-end">
            <Grid item>
              <Link href="/signin" variant="body2">
                Already have an account? Sign in
              </Link>
            </Grid>
          </Grid>
        </Box>
      </Box>
    </Container>
  );
}

export default SignUp;
