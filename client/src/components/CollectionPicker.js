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

class CollectionPicker extends React.Component {

    state = {
        data: [],
    }

    componentDidMount() {
      this.load();
    }

    load = () => {
      var self = this;
      fetch(BASE_API_URL + this.props.dataEndpoint)
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
      this.props.onClose();
    };
  
    handleListItemClick = value => {
      this.props.onClose(value);
    };
  
    render() {
      const { classes, onClose, keyfn, titleText, displayfn, ...other } = this.props;
      return (
        <Dialog onEnter={this.handleOpen} onClose={this.handleClose} aria-labelledby="simple-dialog-title" {...other}>
          <DialogTitle>{titleText}</DialogTitle>
          <div>
            <List>
              {this.state.data.map(function (d, idx) {
                  return (
                <ListItem button onClick={this.handleListItemClick.bind(this, keyfn(d))} key={keyfn(d)}>
                  <ListItemText primary={displayfn(d)} />
                </ListItem>
              )}.bind(this))}
            </List>
          </div>
        </Dialog>
      );
    }
  }
  
  CollectionPicker.propTypes = {
    classes: PropTypes.object.isRequired,
    onClose: PropTypes.func,
    keyfn: PropTypes.func,
    displayfn: PropTypes.func,
    dataEndpoint: PropTypes.string,
    titleText: PropTypes.string,
  };
  
  export default withStyles(styles)(CollectionPicker);
