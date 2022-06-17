import React from 'react';

import {
  Button,
  Container,
  Grid,
  Stack,
  TextField,
  Paper,
  // IconButton,
} from '@mui/material';
// import { Send } from '@mui/icons-material';
import { useTheme } from '@mui/system';

import InfiniteScroll from 'react-infinite-scroller';
import { gql, useQuery, useMutation } from '@apollo/client';

import Post from '../components/Post';
import Gallery from 'react-grid-gallery';

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
        images
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
      images
    }
  }
`;

const UPLOAD_IMAGE = gql`
  mutation UploadImage($file: Upload!) {
    uploadImage(file: $file) {
      name
    }
  }
`;

function PostInput() {
  const [text, setText] = React.useState('');
  const [images, setImages] = React.useState([]);

  const theme = useTheme();

  const galleryImages = images.map(({ blob }) => {
    const url = window.URL.createObjectURL(blob);
    const image = {
      src: url,
      thumbnail: url,
    };
    return image;
  });

  const [publish] = useMutation(PUBLISH_POST, {
    variables: {
      input: {
        text: text,
        images: images.map(({ name }) => name),
      },
    },
    refetchQueries: [
      { query: FEED },
      'Feed',
    ],
  });
  const [upload] = useMutation(UPLOAD_IMAGE);

  const handleChange = (e) => {
    console.log(e);
    setText(e.target.value);
  };

  const disabled = !(text || images.length);
  const handleClick = (e) => {
    e.preventDefault();
    publish().then((res) => {
      setText('');
      setImages([]);
    });
  };

  const handlePaste = (e) => {
    navigator.clipboard.read().
      then((items) => {
        for (const item of items) {
          if (!item.types.includes('image/png')) {
            continue;
          }

          item.getType('image/png').
            then((blob) => {
              upload({
                variables: {
                  file: blob,
                },
              }).
                then((res) => {
                  const { name } = res.data.uploadImage;
                  setImages((images) => {
                    return [...images, {
                      name: name,
                      blob: blob,
                    }];
                  });
                });
            });
        }
      });
  };

  return (
    <Paper sx={{ p: 1 }}>
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
            onChange={handleChange}
            onPaste={handlePaste}
            sx={{
              ml: 1,
              backgroudColor: 'white.100',
            }}
            InputProps={{
              disableUnderline: true,
            }}
          />
        </Grid>
        <Grid item>
          <Button size='small' onClick={handleClick} disabled={disabled}>
            Publish
          </Button>
        </Grid>
      </Grid>
      {galleryImages.length > 0 &&
        <Gallery
          images={galleryImages}
          enableImageSelection={false}
          backdropClosesModal={true}
          margin={theme.spacing(1)}
        />
      }
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

  const items = posts.map((post) => {
    return (
      <Post key={post.id} {...post} />
    );
  });

  return (
    <Container maxWidth='sm' sx={{ mt: 5 }}>
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
          <PostInput />
          {items}
        </Stack>
      </InfiniteScroll>
    </Container >
  );
}

export default Feed;
