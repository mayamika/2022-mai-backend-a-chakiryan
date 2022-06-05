import React from 'react';

import {
  Container,
  Stack,
} from '@mui/material';

import InfiniteScroll from 'react-infinite-scroller';
import { gql, useQuery } from '@apollo/client';

import UserCard from '../components/UserCard';

const FRIENDS = gql`
  query Friends($cursor: Cursor) {
    me {
      id
      friends(first: 10, after: $cursor) {
        edges {
          node {
            id
            login
            name
            surname
          }
          cursor
        }
        pageInfo {
          hasNextPage
          endCursor
        }
      }
    }
  }
`;

function Friends() {
  const { data, loading, error, fetchMore } = useQuery(FRIENDS);

  if (error) {
    console.log(errror);
  }
  if (loading || error) {
    return <div />;
  }

  const nodes = data.me.friends.edges.map((edge) => edge.node);
  const pageInfo = data.me.friends.pageInfo;

  const items = nodes.map((n, id) => {
    return (
      <UserCard key={id + 1} user={n} />
    );
  });
  console.log(pageInfo.hasNextPage);


  return (
    <Container maxWidth='md' sx={{ mt: 5 }}>
      <Stack spacing={2}>
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
          {items}
        </InfiniteScroll>
      </Stack>
    </Container>
  );
}

export default Friends;
