import React from 'react';

import {
  Card,
  CardContent,
  CardHeader,
  CardMedia,
  Typography,
} from '@mui/material';
import { useTheme } from '@mui/system';

import UserAvatar from './UserAvatar';
import Gallery from 'react-grid-gallery';

function humanizeDate(rawDate) {
  const options = { year: 'numeric', month: 'long', day: 'numeric' };
  return new Date(rawDate).toLocaleDateString(undefined, options);
}

function Post(props) {
  const { from, text, createdAt, images } = props;
  const { login, name, surname } = from;

  const theme = useTheme();

  const galleryImages = images.map((src) => {
    const url = '/images/' + src;
    const image = {
      src: url,
      thumbnail: url,
    };
    return image;
  });

  return (
    <Card sx={{ display: 'flex', flexDirection: 'column' }}>
      <CardHeader
        avatar={
          <UserAvatar user={from} />
        }
        title={`${name} ${surname} @${login}`}
        subheader={humanizeDate(createdAt)}
      />
      {text &&
        <CardContent>
          <Typography variant="body1" component="div" color="text.primary">
            {text}
          </Typography>
        </CardContent>
      }
      {galleryImages.length > 0 &&
        <CardMedia sx={{ m: 1, mt: -1 }}>
          <Gallery
            images={galleryImages}
            enableImageSelection={false}
            backdropClosesModal={true}
            margin={theme.spacing(1)}
          />
        </CardMedia>
      }
    </Card >
  );
}

export default Post;
