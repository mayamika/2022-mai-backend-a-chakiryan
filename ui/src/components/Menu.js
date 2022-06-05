import React from 'react';

import {
  AppBar,
  Toolbar,
  Typography,
  Stack,
  Box,
  Button,
  IconButton,
} from '@mui/material';
import { Feed, ChatBubble, People } from '@mui/icons-material';

import {
  Link as RouterLink,
  useLocation,
  useNavigate,
  createSearchParams,
} from 'react-router-dom';

import UserAvatar from './UserAvatar';
import SearchBar from './SearchBar';

import { SessionContext } from '../session';
import { gql, useQuery } from '@apollo/client';

const ME = gql`
  query Me {
    me {
      id
      login
      name
      surname
    }
  }
`;

function Menu() {
  const location = useLocation();
  if (location.pathname == '/signin' || location.pathname == '/signup') {
    return <div />;
  }

  const { data, error } = useQuery(ME);
  const user = (!error && data) ? data.me : {
    login: '',
    name: '',
    surname: '',
  };
  console.log(data, error, user);

  const [, setSession] = React.useContext(SessionContext);
  const navigate = useNavigate();

  const logout = () => {
    setSession({
      token: null,
    });

    navigate('/');
  };
  const searchUsers = (query) => {
    const loc = {
      pathname: '/search',
      search: `?${createSearchParams({ query: query })}`,
    };
    navigate(loc);
  };

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static" color='primary' elevation={1}>
        <Toolbar
          sx={{
            justifyContent: 'flex-start',
          }}
        >
          <Typography variant="h6" component="p">
            Sample Name
          </Typography>
          <SearchBar
            placeholder='Search users…'
            onSubmit={searchUsers}
          />
          <Stack
            direction="row"
            alignItems="center"
            spacing={8}
            flexGrow={1}
            justifyContent='center'
          >
          </Stack>
          <IconButton
            size="large"
            // edge="start"
            color="inherit"
            component={RouterLink}
            to="/feed"
          >
            <Feed fontSize="large" />
          </IconButton>
          <IconButton
            size="large"
            // edge="end"
            color="inherit"
            component={RouterLink}
            to="/chat"
          >
            <ChatBubble fontSize="large" />
          </IconButton>
          <IconButton
            size="large"
            // edge="start"
            color="inherit"
            component={RouterLink}
            to="/friends"
          >
            <People fontSize="large" />
          </IconButton>
          <IconButton
            size="large"
            // edge="start"
            color="inherit"
            component={RouterLink}
            to="/friend-requests"
          >
            <People fontSize="small" />
          </IconButton>

          <Stack direction="row" alignItems="center">
            <UserAvatar
              user={user}
              sx={{ m: 1 }}
            />
            <Typography
              variant="h6"
              noWrap
              component="div"
            >
              {user.name} {user.surname}
            </Typography>
            <Button onClick={logout} color="inherit">Logout</Button>
          </Stack>
        </Toolbar>
      </AppBar>
    </Box >
  );
}

export default Menu;
