/**
 *
 * AllPosts
 *
 */

import React, { memo } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Helmet } from 'react-helmet';
import { createStructuredSelector } from 'reselect';
import { compose } from 'redux';

import { useInjectSaga } from 'utils/injectSaga';
import { useInjectReducer } from 'utils/injectReducer';
import { makeStyles, FormControlLabel, Icon, List } from '@material-ui/core';
import MUIDataTable from 'mui-datatables';
import AddButton from './AddButton';
import { makeSelectPostDialog, makeSelectGetAllPosts } from '../selectors';
import reducer from '../reducer';
import saga from '../saga';
import { openNewPostDialog, closeNewPostDialog, allPosts } from '../actions';
import LoadingIndicator from '../../../components/LoadingIndicator';

const useStyles = makeStyles(theme => ({
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    marginLeft: theme.spacing(1),
    marginRight: theme.spacing(1),
    width: 200,
  },
  dense: {
    marginTop: 19,
  },
  menu: {
    width: 200,
  },
}));

export function AllPostsList({ getAllPosts, loading, error }) {
  const classes = useStyles();
  useInjectReducer({ key: 'allPosts', reducer });
  useInjectSaga({ key: 'allPosts', saga });

  const columns = [
    {
      name: 'Id',
      label: 'S/N',
      options: {
        filter: true,
        customBodyRender: (value, tableMeta) => {
          if (value === '') {
            return '';
          }
          return (
            <FormControlLabel
              label={tableMeta.rowIndex + 1}
              control={<Icon />}
            />
          );
        },
      },
    },
    {
      name: 'title',
      label: 'Tittle',
      options: {
        filter: true,
        sort: false,
      },
    },
    {
      name: 'desc',
      label: 'Description',
      options: {
        filter: true,
        sort: false,
      },
    },
    {
      name: 'content',
      label: 'Contents',
      options: {
        filter: true,
        sort: false,
      },
    },
  ];

  const options = {
    filterType: 'checkbox',
    responsive: 'scrollMaxHeight',
    selectableRows: 'none',
    customToolbar: () => <AddButton />,
  };

  if (loading) {
    return <List component={LoadingIndicator} />;
  }

  return (
    <div>
      <MUIDataTable
        title="All Posts"
        data={getAllPosts}
        columns={columns}
        options={options}
      />
    </div>
  );
}

AllPostsList.propTypes = {
  getAllPosts: PropTypes.array.isRequired,
  loading: PropTypes.bool,
  error: PropTypes.oneOfType([PropTypes.object, PropTypes.bool]),
};

const mapStateToProps = createStructuredSelector({
  getAllPosts: makeSelectGetAllPosts(),
  postDialog: makeSelectPostDialog(),
});

function mapDispatchToProps(dispatch) {
  return {
    openNewPostDialog: () => dispatch(openNewPostDialog()),
    closeNewPostDialog: () => dispatch(closeNewPostDialog()),
    allPostsD: () => dispatch(allPosts()),
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
)(AllPostsList);
