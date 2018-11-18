import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import DialogTitle from '@material-ui/core/DialogTitle';
import Dialog from '@material-ui/core/Dialog';
import DialogContent from '@material-ui/core/DialogContent';
import { BASE_API_URL } from "../config";
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Typography from '@material-ui/core/Typography';

const styles = {}

class DetailViewDialog extends React.Component {

    state = {
      data: {},
    }

    componentDidMount() {
      this.load();
    }

    load = () => {
        var self = this;
        for (var table of this.props.dialog.tables) {
            if (table.staticdata) {
                this.state.data[table.title] = table.staticdata;
                this.setState(state => ({data: this.state.data}));
                continue;
            }
            fetch(BASE_API_URL + table.dataEndpoint)
            .then(function(response) {
                if (response.status >= 400) {
                throw new Error("Bad response from server");
                }
                return response.json();
            })
            .then(function(json) {
                self.state.data[table.title] = JSON.parse(json.data);
                self.setState(state => ({data: self.state.data}));
            });
        }
    }

    handleOpen = () => {
        this.load();
    }

    handleClose = () => {
      this.props.dialog.onClose();
    }
  
    render() {
      const { classes, onClose, dialog, ...other } = this.props;
      return (
        <Dialog onEnter={this.handleOpen} onClose={this.handleClose} aria-labelledby="simple-dialog-title" {...other}>
          <DialogTitle>{dialog.title}</DialogTitle>
          <DialogContent>
              <div>
                {dialog.tables.map(function (table, idx) {
                    console.log(JSON.stringify(this.state.data));
                    console.log(JSON.stringify(table));
                return (
                    <List>
                        <ListItem>
                        <Typography>
                            {table.title}
                        </Typography>
                        </ListItem>
                        <ListItem>
                        <Table>
                            <TableHead>
                            <TableRow>
                            {table.columns.map(function (column, idx) {
                                return (
                                    <TableCell key={column.key}>
                                        {column.name}
                                    </TableCell>)})}
                            </TableRow>
                            </TableHead>
                            <TableBody>
                            {(this.state.data[table.title] || []).map(function (row, rowidx) {
                                return (
                                    <TableRow key={table.title + "row:" + rowidx}>
                                    {table.columns.map(function (column, colidx) {
                                    return (
                                        <TableCell key={table.title + "row:" + rowidx + "col:" + colidx}>
                                            {row[column.key]}
                                        </TableCell>)}.bind(this))}
                                    </TableRow>
                                )
                            }.bind(this))}
                            </TableBody>
                        </Table>
                        </ListItem>
                    </List>
                )}.bind(this))}
              </div>
          </DialogContent>
        </Dialog>
      );
    }
  }
  
  DetailViewDialog.propTypes = {
    classes: PropTypes.object.isRequired,
    dialog: PropTypes.object,
  };
  
  export default withStyles(styles)(DetailViewDialog);
