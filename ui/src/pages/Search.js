import React from 'react';

import {
  useNavigate,
  useSearchParams,
  createSearchParams,
} from 'react-router-dom';

import {
  Container,
  Paper,
  Typography,
  Stack,
  Grid,
  Button,
} from '@mui/material';
import { Box } from '@mui/system';

import { gql, useQuery } from '@apollo/client';

import Post from '../components/Post';
import SearchBar from '../components/SearchBar';
import UserAvatar from '../components/UserAvatar';

const USERS = gql`
  query Users($search: String) {
    users(first: 4, search: $search) {
      totalCount
      edges {
        node {
          id
          login
          name
          surname
          relation
        }
      }
    }
  }
`;

const POSTS = gql`
  query Feed($search: String) {
    feed(first: 5, search: $search) {
      totalCount
      posts {
        id
        from {
          id
          login
          name
          surname
          relation
        }
        text
        createdAt
      }
    }
  }
`;

function shrinkText(text, n) {
  if (text.length > n) {
    text = text.substring(0, n) + '…';
  }
  return text;
}

function SearchSection({ name, onShowAll, children }) {
  return (
    <div>
      <Grid container flexDirection='row'>
        <Grid item flexGrow={1}>
          <Typography variant='h5' component='div'>
            {name}
          </Typography>
        </Grid>
        <Grid item>
          <Button onClick={onShowAll}>
            Show all
          </Button>
        </Grid>
      </Grid>
      <Box mt={1}>
        {children}
      </Box>
    </div>
  );
}

function Users({ query, onShowAll }) {
  const { data } = useQuery(USERS, {
    variables: {
      search: query,
    },
  });

  if (!data) {
    return null;
  }
  const { users } = data;
  if (users.totalCount === 0) {
    return null;
  }

  function User({ user }) {
    const { login, name, surname } = user;
    return (
      <Grid
        container
        direction="column"
        justifyContent="flex-start"
        alignItems="center"
        my={2}
      >
        <Grid item>
          <UserAvatar user={user} size={10} />
        </Grid>
        <Grid item mt={1}>
          <Typography variant='body2' component='div' noWrap>
            {shrinkText(name + ' ' + surname, 14)}
          </Typography>
        </Grid>
        <Grid item>
          <Typography color='text.secondary' noWrap>
            @{shrinkText(login, 12)}
          </Typography>
        </Grid>
      </Grid >
    );
  }
  const items = users.edges.map((e) => {
    return (
      <Grid item key={e.node.id} xs={3}>
        <User user={e.node} />
      </Grid>
    );
  });

  return (
    <SearchSection
      name={`Users (${users.totalCount})`}
      onShowAll={onShowAll}
    >
      <Paper>
        <Grid
          container
          direction="row"
          justifyContent="flex-start"
          alignItems="flex-start"
        >
          {items}
        </Grid>
      </Paper>
    </SearchSection>
  );
}

function Posts({ query, onShowAll }) {
  const { data, error } = useQuery(POSTS, {
    variables: {
      search: (query) ? query : null,
    },
  });
  if (error) {
    console.log(error);
  }

  if (!data) {
    return null;
  }
  console.log(data);

  const { totalCount, posts } = data.feed;
  if (totalCount === 0) {
    return null;
  }

  const items = posts.map((post) => {
    return (
      <Post key={post.id} {...post} />
    );
  });

  return (
    <SearchSection
      name={`Posts (${totalCount})`}
      onShowAll={onShowAll}
    >
      <Stack spacing={2}>
        {items}
      </Stack>
    </SearchSection>
  );
}

function Search() {
  const [searchParams] = useSearchParams();
  const query = searchParams.get('query');

  const navigate = useNavigate();
  const onSubmit = (val) => {
    const loc = {
      pathname: '/search',
      search: `?${createSearchParams({ query: val })}`,
    };
    navigate(loc);
  };

  const showAllUsers = () => {
    const loc = {
      pathname: '/search/users',
      search: `?${createSearchParams({ query: query })}`,
    };
    navigate(loc);
  };
  const showAllPosts = () => {
    const loc = {
      pathname: '/search/posts',
      search: `?${createSearchParams({ query: query })}`,
    };
    navigate(loc);
  };

  return (
    <Container maxWidth='sm' sx={{ mt: 5 }}>
      <Paper
        sx={{ p: 1 }}
      >
        <SearchBar value={query} onSubmit={onSubmit} />
      </Paper>
      <Stack spacing={3} mt={3}>
        <Users query={query} onShowAll={showAllUsers} />
        <Posts query={query} onShowAll={showAllPosts} />
      </Stack>
    </Container>
  );
}

export default Search;
