import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Avatar from '@material-ui/core/Avatar';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import ListItemText from '@material-ui/core/ListItemText';
import DialogTitle from '@material-ui/core/DialogTitle';
import Dialog from '@material-ui/core/Dialog';
import PersonIcon from '@material-ui/icons/Person';
import AddIcon from '@material-ui/icons/Add';
import { BASE_API_URL } from "../config";

const styles = {}

class CollectionPickerDialog extends React.Component {

    state = {
        data: [],
    }

    componentDidMount() {
      this.load();
    }

    load = () => {
      if (this.props.dialog.staticdata) {
        this.setState(state => ({data: this.props.dialog.staticdata}));
        return;
      }
      var self = this;
      fetch(BASE_API_URL + this.props.dialog.dataEndpoint)
      .then(function(response) {
          if (response.status >= 400) {
          throw new Error("Bad response from server");
          }
          return response.json();
      })
      .then(function(json) {
          self.setState(state => ({data: JSON.parse(json.data)}));
      });
    }

    handleOpen = () => {
      this.load();
    }

    handleClose = () => {
      this.props.dialog.onClose();
    };
  
    handleListItemClick = value => {
      this.props.dialog.onClose(value);
    };
  
    render() {
      const { classes, dialog, ...other } = this.props;
      return (
        <Dialog onEnter={this.handleOpen} onClose={this.handleClose} aria-labelledby="simple-dialog-title" {...other}>
          <DialogTitle>{dialog.title}</DialogTitle>
          <div>
            <List>
              {this.state.data.map(function (d, idx) {
                  return (
                <ListItem button onClick={this.handleListItemClick.bind(this, dialog.keyfn(d))} key={dialog.keyfn(d)}>
                  <ListItemText primary={dialog.displayfn(d)} />
                </ListItem>
              )}.bind(this))}
            </List>
          </div>
        </Dialog>
      );
    }
  }
  
  CollectionPickerDialog.propTypes = {
    classes: PropTypes.object.isRequired,
    dialog: PropTypes.object,
  };
  
  export default withStyles(styles)(CollectionPickerDialog);
