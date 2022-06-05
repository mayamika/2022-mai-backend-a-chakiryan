import React from 'react';

import {
  Container,
  Grid,
  Stack,
  TextField,
  Paper,
  IconButton,
} from '@mui/material';
import { Send } from '@mui/icons-material';

import InfiniteScroll from 'react-infinite-scroller';
import { gql, useQuery, useMutation } from '@apollo/client';

import Post from '../components/Post';

const FEED = gql`
  query Feed($after: String) {
    feed(first: 10, after: $after) {
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
      }
    }
  }
`;

const PUBLISH_POST = gql`
  mutation PublishPost($input: PostInput!) {
    publishPost(input: $input) {
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
`;

function PostInput() {
  const [text, setText] = React.useState('');

  const [publish] = useMutation(PUBLISH_POST, {
    variables: {
      input: {
        text: text,
      },
    },
    refetchQueries: [
      { query: FEED },
      'Feed',
    ],
  });

  const onChange = (e) => {
    setText(e.target.value);
  };
  const onClick = (e) => {
    e.preventDefault();
    publish().then((res) => {
      setText('');
    });
  };

  return (
    <Paper>
      <Grid
        container
        direction="row"
        justifyContent="flex-start"
        alignItems="flex-end"
        spacing={2}
      >
        <Grid item flexGrow={1}>
          <TextField
            multiline
            fullWidth
            variant="standard"
            placeholder='New post'
            value={text}
            onChange={onChange}
            sx={{
              m: 2,
              backgroudColor: 'white.100',
            }}
            InputProps={{
              disableUnderline: true,
            }}
          />
        </Grid>
        <Grid item>
          <IconButton sx={{ m: 1 }} onClick={onClick}>
            <Send />
          </IconButton>
        </Grid>
      </Grid>
    </Paper>
  );
}

function Feed() {
  const { data, loading, error, fetchMore } = useQuery(FEED);

  if (error) {
    console.log(error);
  }
  if (loading || error) {
    return <div />;
  }

  const feed = data.feed;
  const posts = feed.posts;

  const items = posts.map((post, id) => {
    return (
      <Post key={id + 1} {...post} />
    );
  });
  console.log('feed', feed);

  return (
    <Container maxWidth='sm' sx={{ mt: 5 }}>
      <InfiniteScroll
        pageStart={0}
        loadMore={() => {
          console.log('feed scroll', feed.scroll);
          fetchMore({
            variables: {
              after: feed.scroll,
            },
          });
        }}
        hasMore={feed.hasNextPage}
      >
        <Stack spacing={2}>
          <PostInput />
          {items}
        </Stack>
      </InfiniteScroll>
    </Container >
  );
}

export default Feed;
