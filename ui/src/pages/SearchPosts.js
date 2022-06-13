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

import Post from '../components/Post';
import SearchBar from '../components/SearchBar';

const POSTS = gql`
  query Feed($after: String, $search: String) {
    feed(first: 10, after: $after, search: $search) {
      totalCount
      hasNextPage
      scroll
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
        images
      }
    }
  }
`;

function Posts({ query }) {
  const { data, fetchMore } = useQuery(POSTS, {
    variables: {
      search: (query) ? query : null,
    },
  });
  if (!data) {
    return null;
  }

  const feed = data.feed;
  const posts = feed.posts;

  const items = posts.map((post) => {
    return (
      <Post key={post.id} {...post} />
    );
  });

  return (
    <InfiniteScroll
      pageStart={0}
      loadMore={() => {
        fetchMore({
          variables: {
            after: feed.scroll,
          },
        });
      }}
      hasMore={feed.hasNextPage}
    >
      <Stack spacing={2}>
        {items}
      </ Stack >
    </InfiniteScroll>
  );
}

function SearchPosts() {
  const [searchParams] = useSearchParams();
  const query = searchParams.get('query');

  const navigate = useNavigate();
  const onSubmit = (val) => {
    const loc = {
      pathname: '/search/posts',
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
        <Posts query={query} />
      </Stack>
    </Container >
  );
}

export default SearchPosts;
