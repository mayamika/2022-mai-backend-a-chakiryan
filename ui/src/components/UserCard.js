import React from 'react';

import {
  Card,
  CardContent,
  Typography,
  Grid,
} from '@mui/material';

import UserAvatar from './UserAvatar';

function UserCard(props) {
  const { user, children } = props;
  const { login, name, surname } = user;

  return (
    <Card >
      <CardContent sx={{
        'p': 2,
        '&:last-child': {
          paddingBottom: 2,
        },
      }}>
        <Grid
          container
          direction="row"
          justifyContent="flex-start"
          alignItems="center"
        >
          <Grid item >
            <UserAvatar user={user} size={8} />
          </Grid>
          <Grid item ml={2} flexGrow={1}>
            <Typography variant="body1" component="div">
              {name} {surname}
            </Typography>
            <Typography variant="body1" component="div" color="text.secondary">
              @{login}
            </Typography>
          </Grid>
          <Grid item>
            {children}
          </Grid>
        </Grid>
      </CardContent>
    </Card >
  );
}

export default UserCard;
