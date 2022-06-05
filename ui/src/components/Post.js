import React from 'react';

import {
  Card,
  CardContent,
  CardHeader,
  Typography,
} from '@mui/material';

import UserAvatar from './UserAvatar';

function humanizeDate(rawDate) {
  const options = { year: 'numeric', month: 'long', day: 'numeric' };
  return new Date(rawDate).toLocaleDateString(undefined, options);
}

function Post(props) {
  const { from, text, createdAt } = props;
  const { login, name, surname } = from;

  return (
    <Card>
      <CardHeader
        avatar={
          <UserAvatar user={from} />
        }
        title={`${name} ${surname} @${login}`}
        subheader={humanizeDate(createdAt)}
      />
      <CardContent>
        <Typography variant="body2" component="div" color="text.primary">
          {text}
        </Typography>
      </CardContent >
    </Card >
  );
}

export default Post;
