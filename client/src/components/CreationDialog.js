import React from 'react';
import Button from '@material-ui/core/Button';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle';
import PropTypes from 'prop-types';
import TextField from '@material-ui/core/TextField';
import CollectionPickerDialog from './CollectionPickerDialog';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import Typography from '@material-ui/core/Typography';

class CreationDialog extends React.Component {

  state = {
      data: {},
      selectedField: null,
      collectionPickerDialog: null,
  }

  handleClose = (result) => {
    this.props.dialog.onClose(result);
  };

  handleDataChange = (field, e) => {
    this.state.data[field.name] = e.target.value;
  }

  handleSelectFieldClose = (data) => {
    this.state.data[this.state.selectedField.name] = data;
    this.setState(state => ({selectedField: null}));
    this.forceUpdate();
  }

  handleSelectField = (field) => {
    this.setState(state => ({selectedField: field,
    collectionPickerDialog: {
      title: "Select " + this.state.selectedField.name,
      onClose: this.handleSelectFieldClose.bind(this),
      staticdata: this.state.selectedField.options,
      displayfn: this.state.selectedField.displayfn,
      keyfn: this.state.selectedField.keyfn,
    }}));
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
              <List>
              {dialog.fields.map(function (field, idx) {
                return (
                    <ListItem key={field.name}>
                        {field.hasOwnProperty("options")?
                    <div>
                        <Typography>
                            {this.state.data[field.name]}
                        </Typography>
                        <Button onClick={this.handleSelectField.bind(this, field)}>
                            Select {field.name}
                        </Button>
                    </div>
                    :
                    <TextField
                        required
                        id={field.name}
                        label={field.name}
                        floatinglabeltext={field.name}
                        margin="normal"
                        onChange={this.handleDataChange.bind(this, field)}
                    />}
                    </ListItem>
            )}.bind(this))}
              </List>
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
        {this.state.selectedField !== null && <CollectionPickerDialog open={this.state.selectedField !== null} dialog={this.state.collectionPickerDialog}/>}
      </div>
    );
  }
}

CreationDialog.propTypes = {
    dialog: PropTypes.object.isRequired,
};

export default CreationDialog;