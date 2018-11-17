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
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';

const styles = {}

class ViewServiceInstanceDialog extends React.Component {

    state = {
      configurations: [],
      keys: [],
    }

    componentDidMount() {
      this.loadConfigurations();
      this.loadKeys();
    }

    loadConfigurations = () => {
        var url = BASE_API_URL + "/serviceInstanceConfiguration/listForServiceInstance?serviceInstanceName=" + this.props.serviceInstance.name +
        "&serviceInstanceServiceName=" + this.props.serviceInstance.serviceName +
        "&serviceInstanceOrganizationName=" + this.props.serviceInstance.organizationName;
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
                configurations: JSON.parse(json.data)
            });
        });
    }

    loadKeys = () => {
      var url = BASE_API_URL + "/serviceInstanceKey/listForServiceInstance?serviceInstanceName=" + this.props.serviceInstance.name +
      "&serviceInstanceServiceName=" + this.props.serviceInstance.serviceName +
      "&serviceInstanceOrganizationName=" + this.props.serviceInstance.organizationName;
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
              keys: JSON.parse(json.data)
          });
      });
    }

    handleClose = () => {
      this.props.onClose();
    };
  
    render() {
      const { classes, onClose, ...other } = this.props;
  
      return (
        <Dialog onClose={this.handleClose} aria-labelledby="simple-dialog-title" {...other}>
          <DialogTitle>Details for {this.props.serviceInstance.name}</DialogTitle>
          <div>
          <Typography>
            Overview
          </Typography>
          <Table className={classes.table}>
            <TableBody>
              <TableRow key="service">
                <TableCell component="th" scope="row">
                  Service
                </TableCell>
                <TableCell>{this.props.serviceInstance.serviceName}</TableCell>
              </TableRow>
              <TableRow key="region">
                <TableCell component="th" scope="row">
                  Region
                </TableCell>
                <TableCell>{this.props.serviceInstance.regionName}</TableCell>
              </TableRow>
            </TableBody>
          </Table>
          <Divider light />
          <Typography>
            Configurations
          </Typography>
          <Table className={classes.table}>
            <TableBody>
              <TableHead>
              <TableRow>
                <TableCell component="th" scope="row">
                  Key
                </TableCell>
                <TableCell component="th" scope="row">
                  Value
                </TableCell>
              </TableRow>
              </TableHead>

            {this.state.configurations.map(function(configuration, idx){
                    return (
              <TableRow key={configuration.configKey}>
                <TableCell component="th" scope="row">
                  {configuration.configKey}
                </TableCell>
                <TableCell component="th" scope="row">
                  {configuration.data}
                </TableCell>
              </TableRow>)})}
            </TableBody>
          </Table>
          <Divider light />
          <Typography>
            Keys
          </Typography>
          <Table className={classes.table}>
            <TableBody>
            <TableHead>
              <TableRow>
                <TableCell component="th" scope="row">
                  Key
                </TableCell>
                <TableCell component="th" scope="row">
                  Active until
                </TableCell>
              </TableRow>
              </TableHead>
            {this.state.keys.map(function(key, idx){
                    return (
              <TableRow key={key.keyValue}>
                <TableCell component="th" scope="row">
                  {key.keyValue}
                </TableCell>
                <TableCell component="th" scope="row">
                  {key.activeUntil}
                </TableCell>
              </TableRow>)})}
            </TableBody>
          </Table>
          </div>
        </Dialog>
      );
    }
  }
  
  ViewServiceInstanceDialog.propTypes = {
    classes: PropTypes.object.isRequired,
    onClose: PropTypes.func,
    selectedValue: PropTypes.string,
  };
  
  export default withStyles(styles)(ViewServiceInstanceDialog);
