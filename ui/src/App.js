import React from 'react';
import './App.css';

import {
  BrowserRouter as Router,
  Routes,
  Route,
  useNavigate,
  useLocation,
} from 'react-router-dom';

import {
  CssBaseline,
  Link,
  Box,
  Typography,
} from '@mui/material';

import SessionProvider, { SessionContext } from './session';
import ApolloProvider from './apollo';

import { gql, useQuery } from '@apollo/client';

function Copyright() {
  return (
    <Typography variant="body2" color="text.secondary" align="center" mt={5}>
      {'Copyright © '}
      <Link color="inherit" href="https://github.com/mayamika">
        mayamika
      </Link>{' '}
      {new Date().getFullYear()}
      {'.'}
    </Typography>
  );
}

const ME = gql`
  query Me {
    me {
      login
    }
  }
`;

function LoginRedirect({ children }) {
  const location = useLocation();
  const navigate = useNavigate();
  const [, setSession] = React.useContext(SessionContext);

  const { loading, error } = useQuery(ME);

  React.useEffect(() => {
    if (location.pathname == '/signin' || location.pathname == '/signup') {
      return;
    }
    if (loading) {
      return;
    }

    if (error) {
      setSession({
        token: null,
      });
      navigate('/signin');
    }

    if (location.pathname == '/') {
      navigate('/feed');
    }
  }, [location, loading, error]);

  return children;
}

function Home() {
  return (<div />);
}

import Menu from './components/Menu';

import SignUp from './pages/SignUp';
import SignIn from './pages/SignIn';
import Feed from './pages/Feed';
import User from './pages/User';
import Friends from './pages/Friends';
import SearchUsers from './pages/SearchUsers';
import FriendRequests from './pages/FriendRequests';

function App() {
  return (
    <SessionProvider>
      <ApolloProvider>
        <Router>
          <LoginRedirect>
            <CssBaseline />
            <Box sx={{
              minHeight: '100vh',
              display: 'flex',
              flexDirection: 'column',
              justifyConent: 'flex-end',
              backgroundColor: 'grey.200',
            }}>
              <Box sx={{ flexGrow: 1 }}>
                <Menu />
                <Routes>
                  <Route path="/" element={<Home />} />
                  <Route path="/signup" element={<SignUp />} />
                  <Route path="/signin" element={<SignIn />} />
                  <Route path="/feed" element={<Feed />} />
                  <Route path="/u/:login" element={<User />} />
                  <Route path="/friends" element={<Friends />} />
                  <Route path="/search" element={<SearchUsers />} />
                  <Route path="/friend-requests" element={<FriendRequests />} />
                </Routes>
              </Box>
              <Box sx={{ mb: 1 }}>
                <Copyright />
              </Box>
            </Box>
          </LoginRedirect>
        </Router>
      </ApolloProvider>
    </SessionProvider>
  );
}

export default App;
