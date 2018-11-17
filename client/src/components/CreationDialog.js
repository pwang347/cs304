import React from 'react';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import PropTypes from 'prop-types';
import TextField from '@material-ui/core/TextField';

class CreationDialog extends React.Component {

  state = {
      data: {},
  }

  handleClose = (result) => {
    this.props.dialog.onClose(result);
  };

  handleDataChange = (field, e) => {
    this.state.data[field] = e.target.value;
  }

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
          {dialog.fields.map(function (field, idx) {
            return (
                <TextField
                    required
                    id={field}
                    label={field}
                    floatinglabeltext={field}
                    margin="normal"
                    value={this.state.data[field]}
                    onChange={this.handleDataChange.bind(this, field)}
                    key={field}
                />
            )}.bind(this))}
          </DialogContent>
          <DialogActions>
            <Button onClick={this.handleClose.bind(this, this.state.data)} color="primary" autoFocus>
              Create
            </Button>
            <Button onClick={this.handleClose.bind(this, null)} color="primary">
              Cancel
            </Button>
          </DialogActions>
        </Dialog>
      </div>
    );
  }
}

CreationDialog.propTypes = {
    dialog: PropTypes.object.isRequired,
};

export default CreationDialog;