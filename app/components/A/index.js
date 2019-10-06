/**
 * A link to a certain page, an anchor tag
 */
import React from 'react';
import { Button } from '@material-ui/core';
import { makeStyles } from '@material-ui/core/styles';

const useStyles = makeStyles({
  root: {
    color: 'red',
  },
});

const A = ({ children }) => {
  const classes = useStyles();
  return <Button className={classes.root}>{children}</Button>;
};

export default A;
