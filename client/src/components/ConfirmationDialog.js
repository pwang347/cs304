import React from 'react';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import PropTypes from 'prop-types';

class ConfirmationDialog extends React.Component {

  handleClose = (result) => {
    this.props.dialog.onClose(result);
  };

  render() {
    const { classes, onClose, dialog, ...other } = this.props;

    return (
      <div>
        <Dialog
          onClose={this.handleClose.bind(this, false)}
          aria-labelledby="alert-dialog-title"
          aria-describedby="alert-dialog-description"
          open={true}
          {...other}
        >
          <DialogTitle id="alert-dialog-title">{dialog.titleText}</DialogTitle>
          <DialogContent>
            <DialogContentText id="alert-dialog-description">
              {dialog.contentText}
            </DialogContentText>
          </DialogContent>
          <DialogActions>
            <Button onClick={this.handleClose.bind(this, true)} color="primary">
              {dialog.yesText}
            </Button>
            <Button onClick={this.handleClose.bind(this, false)} color="primary" autoFocus>
              {dialog.noText}
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }
}

ConfirmationDialog.propTypes = {
    dialog: PropTypes.object.isRequired,
};

export default ConfirmationDialog;