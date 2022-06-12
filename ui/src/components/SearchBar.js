import React from 'react';
import { styled, alpha } from '@mui/material/styles';
import { Box, InputBase } from '@mui/material';
import SearchIcon from '@mui/icons-material/Search';

const Search = styled('div')(({ theme }) => ({
  'position': 'relative',
  'borderRadius': theme.shape.borderRadius,
  'backgroundColor': alpha(theme.palette.common.white, 0.15),
  '&:hover': {
    backgroundColor: alpha(theme.palette.common.white, 0.25),
  },
  'marginRight': theme.spacing(2),
  'marginLeft': 0,
  'width': '100%',
  [theme.breakpoints.up('sm')]: {
    width: 'auto',
  },
}));

const SearchIconWrapper = styled('div')(({ theme }) => ({
  padding: theme.spacing(0, 2),
  height: '100%',
  position: 'absolute',
  pointerEvents: 'none',
  display: 'flex',
  alignItems: 'center',
  justifyContent: 'center',
}));

const StyledInputBase = styled(InputBase)(({ theme }) => ({
  'color': 'inherit',
  '& .MuiInputBase-input': {
    padding: theme.spacing(1, 1, 1, 0),
    // vertical padding + font size from searchIcon
    paddingLeft: `calc(1em + ${theme.spacing(4)})`,
    transition: theme.transitions.create('width'),
  },
}));

function SearchBar({ placeholder, onSubmit, value }) {
  const [query, setQuery] = React.useState(value);

  React.useEffect(() => {
    setQuery(value);
  }, [value]);

  const handleSubmit = (e) => {
    e.preventDefault();
    if (onSubmit) onSubmit(query, setQuery);
  };

  return (
    <Box component="form" noValidate onSubmit={handleSubmit}>
      <Search>
        <SearchIconWrapper>
          <SearchIcon />
        </SearchIconWrapper>
        <StyledInputBase
          fullWidth
          placeholder={placeholder}
          inputProps={{ 'aria-label': 'search' }}
          value={query}
          onChange={(e) => {
            e.preventDefault();
            setQuery(e.target.value);
          }}
          onSubmit={(e) => {
            e.preventDefault();
            console.log('sub');
          }}
        />
      </Search>
    </Box>
  );
}

SearchBar.defaultProps = {
  placeholder: 'Search…',
  value: '',
};

export default SearchBar;
