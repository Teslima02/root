/**
 *
 * AllPosts
 *
 */

import React, { memo } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Helmet } from 'react-helmet';
import { FormattedMessage } from 'react-intl';
import { createStructuredSelector } from 'reselect';
import { compose } from 'redux';

import { useInjectSaga } from 'utils/injectSaga';
import { useInjectReducer } from 'utils/injectReducer';
import { Grid, Paper, TextField, makeStyles, Button } from '@material-ui/core';
import makeSelectAllPosts from './selectors';
import reducer from './reducer';
import saga from './saga';
import messages from './messages';
import { AllPostsList } from './components/AllPostsList';

const useStyles = makeStyles(theme => ({
  textField: {
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
  },
}));

export function AllPosts() {
  const classes = useStyles();
  useInjectReducer({ key: 'allPosts', reducer });
  useInjectSaga({ key: 'allPosts', saga });

  const [values, setValues] = React.useState({
    title: '',
    description: '',
  });

  const handleChange = name => event => {
    setValues({ ...values, [name]: event.target.value });
  };

  return (
    <div>
      <Helmet>
        <title>AllPosts</title>
        <meta name="description" content="Description of AllPosts" />
      </Helmet>

      <Grid container spacing={3}>
        <Grid item xs={12} sm={12} md={12}>
          <div>
            <TextField
              id="standard-title"
              label="Title"
              className={classes.textField}
              value={values.title}
              onChange={handleChange('title')}
              margin="normal"
            />
            <TextField
              id="standard-description"
              label="Description"
              className={classes.textField}
              value={values.description}
              onChange={handleChange('description')}
              margin="normal"
            />
          </div>

          <AllPostsList />
        </Grid>
      </Grid>
    </div>
  );
}

AllPosts.propTypes = {
  // dispatch: PropTypes.func.isRequired,
};

const mapStateToProps = createStructuredSelector({
  allPosts: makeSelectAllPosts(),
});

function mapDispatchToProps(dispatch) {
  return {
    dispatch,
  };
}

const withConnect = connect(
  mapStateToProps,
  mapDispatchToProps,
);

export default compose(
  withConnect,
  memo,
)(AllPosts);
