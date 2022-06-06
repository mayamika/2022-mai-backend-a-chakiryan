import React from 'react';

import {
  AppBar,
  Toolbar,
  Typography,
  Stack,
  Box,
  Popper,
  Grow,
  Paper,
  Button,
  ClickAwayListener,
  MenuList,
  MenuItem,
  IconButton,
} from '@mui/material';
import { Feed, People, PersonAdd } from '@mui/icons-material';

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

  const searchUsers = (query) => {
    const loc = {
      pathname: '/search',
      search: `?${createSearchParams({ query: query })}`,
    };
    navigate(loc);
  };

  const [open, setOpen] = React.useState(false);
  const anchorRef = React.useRef(null);

  const handleClose = (e) => {
    if (anchorRef.current && anchorRef.current.contains(e.target)) {
      return;
    }
    setOpen(false);
  };
  const handleLogout = (e) => {
    handleClose(e);

    setSession({
      token: null,
    });
    navigate('/');
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
            <Feed fontSize="medium" />
          </IconButton>
          {/* <IconButton
            size="large"
            // edge="end"
            color="inherit"
            component={RouterLink}
            to="/chat"
          >
            <ChatBubble fontSize="medium" />
          </IconButton> */}
          <IconButton
            size="large"
            // edge="start"
            color="inherit"
            component={RouterLink}
            to="/friends"
          >
            <People fontSize="medium" />
          </IconButton>
          <IconButton
            size="large"
            // edge="start"
            color="inherit"
            component={RouterLink}
            to="/friend-requests"
          >
            <PersonAdd fontSize="medium" />
          </IconButton>

          <Stack direction="row" alignItems="center">
            <Button
              ref={anchorRef}
              onClick={(e) => {
                setOpen((prevOpen) => !prevOpen);
              }}
              color='inherit'
              size="small"
              startIcon={
                <UserAvatar
                  user={user}
                  sx={{ m: 1 }}
                />
              }
              sx={{
                textTransform: 'none',
              }}
            >
              <Typography
                variant="h6"
                noWrap
                component="div"
              >
                {user.name} {user.surname}
              </Typography>
            </Button>
            <Popper
              open={open}
              anchorEl={anchorRef.current}
              role={undefined}
              placement="bottom-start"
              transition
              disablePortal
            >
              {({ TransitionProps, placement }) => (
                <Grow
                  {...TransitionProps}
                  style={{
                    transformOrigin:
                      placement === 'bottom-start' ? 'left top' : 'left bottom',
                  }}
                >
                  <Paper>
                    <ClickAwayListener onClickAway={handleClose}>
                      <MenuList
                        autoFocusItem={open}
                      >
                        <MenuItem onClick={handleLogout}>Logout</MenuItem>
                      </MenuList>
                    </ClickAwayListener>
                  </Paper>
                </Grow>
              )}
            </Popper>
          </Stack>
        </Toolbar>
      </AppBar>
    </Box >
  );
}

export default Menu;
