import React from 'react';
import PropTypes from 'prop-types';
import { withStyles, Grid } from '@material-ui/core';
import CssBaseline from '@material-ui/core/CssBaseline';
import Header2 from '../../Header2';
import Footer from '../../Footer';

const drawerWidth = 240;

const styles = theme => ({
  root: {
    display: 'flex',
  },
  drawer: {
    [theme.breakpoints.up('sm')]: {
      width: drawerWidth,
      flexShrink: 0,
    },
  },
  appBar: {
    marginLeft: drawerWidth,
    [theme.breakpoints.up('sm')]: {
      width: `calc(100% - ${drawerWidth}px)`,
    },
  },
  menuButton: {
    marginRight: theme.spacing(2),
    [theme.breakpoints.up('sm')]: {
      display: 'none',
    },
  },
  toolbar: {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'flex-end',
    padding: theme.spacing(0, 1),
    ...theme.mixins.toolbar,
  },
  content: {
    flexGrow: 1,
    padding: theme.spacing(3),
  },
});

export function Layout1(props) {
  return (
    <React.Fragment>
      <CssBaseline />

      <Grid container>
        <Grid item xs={12} sm={12} md={12}>
          <Header2 />
        </Grid>
      </Grid>

      <Grid container className={props.classes.root}>
        <Grid item xs={12} sm={12} md={12}>
          <main className={props.classes.content}>
            <div className={props.classes.toolbar} />
            {props.children}
          </main>
        </Grid>
      </Grid>
      <Footer />
    </React.Fragment>
  );
}

Layout1.propTypes = {
  classes: PropTypes.object.isRequired,
  children: PropTypes.object.isRequired,
};

export default withStyles(styles)(Layout1);
