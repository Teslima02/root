import React from 'react';
import { withStyles } from '@material-ui/core';
import Header from '../../Header';
import Footer from '../../Footer';

const styles = theme => ({
  root: {
    backgroundColor: theme.palette.common.white,
  },
});

class Layout2 extends React.Component {
  render() {
    const { classes } = this.props;

    return (
      <React.Fragment>
        <Header />
        {this.props.children}
        <Footer />
      </React.Fragment>
    );
  }
}

export default withStyles(styles)(Layout2);
