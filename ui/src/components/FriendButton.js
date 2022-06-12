import React from 'react';

import { Button } from '@mui/material';
import { gql, useMutation } from '@apollo/client';

const ADD_FRIEND_REQUEST = gql`
  mutation AddFriendRequest($to: ID!) {
    addFriendRequest(to: $to) {
      id
    }
  }
`;

function AddButton(props) {
  const { id } = props.user;

  const [add] = useMutation(ADD_FRIEND_REQUEST, {
    variables: {
      to: id,
    },
  });

  const [disabled, setDisabled] = React.useState(false);
  const [text, setText] = React.useState('Add to friends');

  const onClick = (e) => {
    e.preventDefault();
    add().then((res) => {
      setDisabled(true);
      setText('Request sent');
    });
  };

  return (
    <Button size='small' variant="outlined"
      disabled={disabled}
      onClick={onClick}
    >
      {text}
    </Button>
  );
}

function DeleteButton(props) {
  // TODO: Add gql query && implement.
  return (
    <Button size='small' variant="outlined">
      Delete friend
    </Button>
  );
}

function SentButton(props) {
  return (
    <Button size='small' variant="outlined" disabled>
      {'Request sent'}
    </Button>
  );
}

function YouButton(props) {
  return (
    <Button size='small' variant="outlined" disabled>
      {'It\'s you'}
    </Button>
  );
}

function FriendButton(props) {
  const { user } = props;
  const { relation } = user;

  switch (relation) {
    case 'YOU':
      return <YouButton user={user} />;
    case 'FRIEND':
      return <DeleteButton user={user} />;
    case 'FRIEND_REQUEST_SENT':
      return <SentButton user={user} />;
    default:
      return <AddButton user={user} />;
  }
}

export default FriendButton;
