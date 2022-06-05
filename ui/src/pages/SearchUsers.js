import React from 'react';

import { useSearchParams } from 'react-router-dom';

import {
  Container,
  Stack,
} from '@mui/material';

import InfiniteScroll from 'react-infinite-scroller';
import { gql, useQuery } from '@apollo/client';

import FriendButton from '../components/FriendButton';
import UserCard from '../components/UserCard';

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

function SearchUsers() {
  const [searchParams] = useSearchParams();
  const query = searchParams.get('query');

  const { data, loading, error, fetchMore } = useQuery(USERS, {
    variables: {
      search: (query) ? query : null,
    },
  });

  if (error) {
    console.log(error);
  }
  if (loading || error) {
    return <div />;
  }

  const nodes = data.users.edges.map((edge) => edge.node);
  const pageInfo = data.users.pageInfo;

  const items = nodes.map((n, id) => {
    return (
      <UserCard key={id + 1} user={n}>
        <FriendButton user={n} />
      </UserCard>
    );
  });
  console.log(pageInfo.hasNextPage);

  return (
    <Container maxWidth='md' sx={{ mt: 5 }}>
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
    </Container>
  );
}

export default SearchUsers;
