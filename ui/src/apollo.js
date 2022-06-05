import React from 'react';

import {
  ApolloClient,
  InMemoryCache,
  ApolloProvider as Provider,
  createHttpLink,
} from '@apollo/client';

import { relayStylePagination } from '@apollo/client/utilities';
import { setContext } from '@apollo/client/link/context';
import { SessionContext } from './session';

function ApolloProvider({ children }) {
  const httpLink = createHttpLink({
    uri: '/query',
  });

  const [session] = React.useContext(SessionContext);
  const authLink = setContext((_, { headers }) => {
    return {
      headers: {
        ...headers,
        'AuthToken': session ? session.token : '',
      },
    };
  });

  const cache = new InMemoryCache({
    typePolicies: {
      Query: {
        fields: {
          feed: {
            merge(existing, incoming, { readField }) {
              const posts = existing ? { ...existing.posts } : {};
              if (incoming.posts) {
                incoming.posts.forEach((post) => {
                  posts[readField('id', post)] = post;
                });
              }
              return {
                totalCount: incoming.totalCount,
                hasNextPage: incoming.hasNextPage,
                scroll: incoming.scroll,
                posts: posts,
              };
            },

            read(existing) {
              if (existing) {
                return {
                  totalCount: existing.totalCount,
                  hasNextPage: existing.hasNextPage,
                  scroll: existing.scroll,
                  posts: Object.values(existing.posts),
                };
              }
            },
          },
        },
      },
      User: {
        fields: {
          friends: relayStylePagination(),
        },
      },
    },
  });

  const client = new ApolloClient({
    link: authLink.concat(httpLink),
    cache: cache,
  });
  return (
    <Provider client={client}>
      {children}
    </Provider>
  );
};

export default ApolloProvider;
