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
    this.setState({
      selectedField: field,
      collectionPickerDialog: {
        title: "Select " + field.name,
        onClose: this.handleSelectFieldClose.bind(this),
        staticdata: field.options,
        displayfn: field.displayfn,
        keyfn: field.keyfn,
      }
    });
  }

  handleOpen = () => {
    if (!this.props.dialog.updateDefaults) {
      return;
    }
    this.setState(state => ({data: this.props.dialog.updateDefaults}));
  }

  render() {
    const { classes, onClose, dialog, ...other } = this.props;
    return (
      <div>
        <Dialog
          onEnter={this.handleOpen}
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
                        <Button onClick={this.handleSelectField.bind(this, field)}>
                            Select {field.name}
                        </Button>
                        <Typography>
                            {this.state.data[field.name]}
                        </Typography>
                    </div>
                    :
                    <TextField
                        required
                        id={field.name}
                        label={field.name}
                        floatinglabeltext={field.name}
                        margin="normal"
                        onChange={this.handleDataChange.bind(this, field)}
                        value={this.state.data[field.name]}
                    />}
                    </ListItem>
            )}.bind(this))}
              </List>
          </DialogContent>
          <DialogActions>
            <Button onClick={this.handleClose.bind(this, this.state.data)} color="primary" autoFocus>
              {dialog.createText || "Create"}
            </Button>
            <Button onClick={this.handleClose.bind(this, null)} color="primary">
              {dialog.cancelText || "Cancel"}
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