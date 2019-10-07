import React from 'react';
import { IconButton, Tooltip, Icon } from '@material-ui/core';
import { Add, CloudUpload } from '@material-ui/icons';
import { withStyles } from '@material-ui/core/styles';
const defaultToolbarStyles = {
  iconButton: {},
};

function CustomToolbar(props) {

  handleClick = () => {
    console.log('clicked on icon!');
  };
  const { openNewDialog, openNewMaturityUpload, classes } = this.props;

  return (
    <React.Fragment>
      <Tooltip title="New Product">
        <IconButton className={classes.iconButton} onClick={openNewDialog}>
          <Add className={classes.deleteIcon} />
        </IconButton>
      </Tooltip>

      <Tooltip title="Maturity Upload">
        <IconButton
          className={classes.iconButton}
          onClick={openNewMaturityUpload}
        >
          <CloudUpload className={classes.deleteIcon} />
        </IconButton>
      </Tooltip>
    </React.Fragment>
  );
}

export default withStyles(defaultToolbarStyles, { name: 'CustomToolbar' })(
  CustomToolbar,
);
