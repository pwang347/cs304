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

class OrganizationPicker extends React.Component {

    state = {
        organizations: [],
    }

    componentDidMount() {
        this.loadOrganizations();
    }

    loadOrganizations = () => {
        var url = BASE_API_URL + "/organization/listUser?userEmailAddress=" + this.props.userEmailAddress;
        var self = this;
        fetch(url)
        .then(function(response) {
            if (response.status >= 400) {
            throw new Error("Bad response from server");
            }
            return response.json();
        })
        .then(function(json) {
            self.setState({
                organizations: JSON.parse(json.data)
            });
        });
    }

    handleClose = () => {
      this.props.onClose(this.props.selectedValue);
    };
  
    handleListItemClick = value => {
      this.props.onClose(value);
    };
  
    render() {
      const { classes, onClose, selectedValue, userEmailAddress, ...other } = this.props;
  
      return (
        <Dialog onClose={this.handleClose} aria-labelledby="simple-dialog-title" {...other}>
          <DialogTitle id="simple-dialog-title">Pick an organization</DialogTitle>
          <div>
            <List>
              {this.state.organizations.map(organization => (
                <ListItem button onClick={() => this.handleListItemClick(organization.organizationName)} key={organization}>
                  <ListItemText primary={organization.organizationName} />
                </ListItem>
              ))}
            </List>
          </div>
        </Dialog>
      );
    }
  }
  
  OrganizationPicker.propTypes = {
    classes: PropTypes.object.isRequired,
    onClose: PropTypes.func,
    selectedValue: PropTypes.string,
  };
  
  export default withStyles(styles)(OrganizationPicker);
