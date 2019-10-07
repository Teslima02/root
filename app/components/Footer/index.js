import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Container, Grid, Typography, Paper } from '@material-ui/core';
import Wrapper from './Wrapper';

const drawerWidth = 240;

const useStyles = makeStyles(theme => ({
  root: {
    backgroundColor: '#344655',
    display: 'flex',
    [theme.breakpoints.up('sm')]: {
      // width: '100%',
      marginLeft: theme.spacing(3),
    },
  },
  footer: {
    padding: theme.spacing(2),
    marginTop: 'auto',
  },
  paper: {
    padding: theme.spacing(2),
    color: theme.palette.grey[400],
    boxShadow: theme.shadows[0],
    background: 'none',
  },
  header: {
    color: theme.palette.grey[100],
    padding: theme.spacing(0),
  },
}));

function Footer() {
  const classes = useStyles();

  return (
    <React.Fragment>
      <div className={classes.root}>
        <Container component="main" maxWidth="md">
          <Grid item xs={12} className={classes.footer}>
            <Paper className={classes.paper}>
              <Typography
                variant="body2"
                align="center"
                className={classes.header}
              >
                Copyright © 2019 Blue
              </Typography>
            </Paper>
          </Grid>
        </Container>
      </div>
    </React.Fragment>
  );
}

export default Footer;
