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
    <Card>
      <CardContent>
        <Grid
          container
          direction="row"
          justifyContent="flex-start"
          alignItems="center"
          spacing={2}
        >
          <Grid item>
            <UserAvatar user={user} size={10} sx={{ m: 1 }} />
          </Grid>
          <Grid item>
            <Typography variant="h5" component="div">
              {name} {surname}
            </Typography>
          </Grid>
          <Grid item flexGrow={1}>
            <Typography variant="h5" component="div" color="text.secondary">
              @{login}
            </Typography>
          </Grid>
          <Grid item>
            {children}
          </Grid>
        </Grid>
      </CardContent>
    </Card>
  );
}

export default UserCard;
