import React from 'react';

import {
  Container,
  Stack,
  Button,
} from '@mui/material';

import InfiniteScroll from 'react-infinite-scroller';
import { gql, useQuery, useMutation } from '@apollo/client';

import UserCard from '../components/UserCard';

const FRIEND_REQUESTS = gql`
  query FriendRequests($cursor: Cursor) {
    friendRequests(first: 10, after: $cursor) {
      pageInfo {
        hasNextPage
        endCursor
      }
      edges {
        cursor
        node {
          id
          from {
            id
            login
            name
            surname
            relation
          }
          to {
            id
            login
            name
            surname
            relation
          }
        }
      }
    }
  }
`;

const ACCEPT_FRIEND_REQUEST = gql`
  mutation AcceptFriendRequest($id: ID!) {
    acceptFriendRequest(id: $id)
  }
`;

const DECLINE_FRIEND_REQUEST = gql`
  mutation DeclineFriendRequest($id: ID!) {
    declineFriendRequest(id: $id)
  }
`;

function RequestButtons(props) {
  const { friendRequest } = props;
  const { id } = friendRequest;

  const [disabled, setDisabled] = React.useState(false);

  const params = {
    variables: {
      id: id,
    },
    update(cache) {
      const normalizedId = cache.identify({ id, __typename: 'FriendRequest' });
      cache.evict({ id: normalizedId });
      cache.gc();
    },
  };

  const [accept] = useMutation(ACCEPT_FRIEND_REQUEST, { ...params });
  const [decline] = useMutation(DECLINE_FRIEND_REQUEST, { ...params });

  const onAccept = (e) => {
    e.preventDefault();
    accept().then((res) => {
      setDisabled(true);
    });
  };
  const onDecline = (e) => {
    e.preventDefault();
    decline().then((res) => {
      setDisabled(true);
    });
  };

  return (
    <div>
      {!disabled &&
        <Stack spacing={2}>
          <Button variant="outlined" onClick={onAccept} >
            Accept
          </Button>
          <Button variant="outlined" onClick={onDecline} >
            Decline
          </Button>
        </Stack>
      }
    </div>
  );
}

function FriendRequests() {
  const { data, loading, error, fetchMore } = useQuery(FRIEND_REQUESTS);

  if (error) {
    console.log(errror);
  }
  if (loading || error) {
    return <div />;
  }

  const nodes = data.friendRequests.edges.map((edge) => edge.node);
  const pageInfo = data.friendRequests.pageInfo;

  const items = nodes.map((n, id) => {
    return (
      <UserCard key={id + 1} user={n.from}>
        <RequestButtons friendRequest={n} />
      </UserCard>
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

export default FriendRequests;
