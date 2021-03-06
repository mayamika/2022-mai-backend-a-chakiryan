import React from 'react';

import {
  useNavigate,
  useSearchParams,
  createSearchParams,
} from 'react-router-dom';

import {
  Container,
  Stack,
  Paper,
} from '@mui/material';

import InfiniteScroll from 'react-infinite-scroller';
import { gql, useQuery } from '@apollo/client';

import FriendButton from '../components/FriendButton';
import UserCard from '../components/UserCard';
import SearchBar from '../components/SearchBar';

const USERS = gql`
  query Users($cursor: Cursor, $search: String) {
    users(first: 10, after: $cursor, search: $search) {
      edges {
        node {
          id
          login
          name
          surname
          relation
        }
        cursor
      }
      pageInfo {
        hasNextPage
        endCursor
      }
    }
  }
`;

function Users({ query }) {
  const { data, fetchMore } = useQuery(USERS, {
    variables: {
      search: (query) ? query : null,
    },
  });

  if (!data) {
    return null;
  }

  const nodes = data.users.edges.map((edge) => edge.node);
  const pageInfo = data.users.pageInfo;

  const items = nodes.map((n) => {
    return (
      <UserCard key={n.id} user={n}>
        <FriendButton user={n} />
      </UserCard>
    );
  });

  return (
    <InfiniteScroll
      pageStart={0}
      loadMore={() => {
        fetchMore({
          variables: {
            cursor: pageInfo.endCursor,
          },
        });
      }}
      hasMore={pageInfo.hasNextPage}
      loader={<div className="loader" key={0}>Loading ...</div>}
    >
      <Stack spacing={2}>
        {items}
      </Stack>
    </InfiniteScroll>
  );
}

function SearchUsers() {
  const [searchParams] = useSearchParams();
  const query = searchParams.get('query');

  const navigate = useNavigate();
  const onSubmit = (val) => {
    const loc = {
      pathname: '/search/users',
      search: `?${createSearchParams({ query: val })}`,
    };
    navigate(loc);
  };

  return (
    <Container maxWidth='sm' sx={{ mt: 5 }}>
      <Stack spacing={2}>
        <Paper
          sx={{ p: 1 }}
        >
          <SearchBar value={query} onSubmit={onSubmit} />
        </Paper>
        <Users query={query} />
      </Stack>
    </Container>
  );
}

export default SearchUsers;
