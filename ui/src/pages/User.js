import React from 'react';

import { Container } from '@mui/material';
import { useParams } from 'react-router-dom';

function User() {
  const { login } = useParams();

  const [user] = React.useState({
    name: 'Hello',
    surname: 'Pediks',
    login: login,
  });

  return (
    <Container>
      {user.name}
    </Container>
  );
}

export default User;
