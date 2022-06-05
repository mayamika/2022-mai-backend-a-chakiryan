import React from 'react';

import { Avatar } from '@mui/material';

function stringToColor(str) {
  let hash = 0;
  let i;

  /* eslint-disable no-bitwise */
  for (i = 0; i < str.length; i += 1) {
    hash = str.charCodeAt(i) + ((hash << 5) - hash);
  }

  let color = '#';

  for (i = 0; i < 3; i += 1) {
    const value = (hash >> (i * 8)) & 0xff;
    color += `00${value.toString(16)}`.slice(-2);
  }
  /* eslint-enable no-bitwise */

  return color;
}

function stringAvatar(user, size, sx) {
  const { name, surname } = user;
  const text = `${name[0] && name[0].toUpperCase()}` +
    `${surname[0] && surname[0].toUpperCase()}`;
  return {
    sx: {
      ...sx,
      bgcolor: stringToColor(`${name} ${surname}`),
      height: (theme) => theme.spacing(size),
      width: (theme) => theme.spacing(size),
      fontSize: (theme) => theme.spacing(0.4 * size),
    },
    children: text,
  };
}

function UserAvatar(props) {
  const { user, size, sx } = props;
  return (
    <Avatar {...stringAvatar(user, size, sx)} />
  );
}

UserAvatar.defaultProps = {
  size: 6,
};

export default UserAvatar;
