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
import Button from '@material-ui/core/Button';

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
            const title = table.title;
            fetch(BASE_API_URL + table.dataEndpoint)
            .then(function(response) {
                if (response.status >= 400) {
                throw new Error("Bad response from server");
                }
                return response.json();
            })
            .then(function(json) {
                self.state.data[title] = JSON.parse(json.data);
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

    handleAdd = (table) => {
        table.addFn(this.load.bind(this));
    }

    handleDelete = (table, row) => {
        table.deleteFn(row, this.load.bind(this));
    }
  
    render() {
      const { classes, onClose, dialog, ...other } = this.props;
      return (
        <Dialog onEnter={this.handleOpen} onClose={this.handleClose} aria-labelledby="simple-dialog-title" {...other}>
          <DialogTitle>{dialog.title}</DialogTitle>
          <DialogContent>
              <div>
                {dialog.tables.map(function (table, idx) {
                return (
                    <List key={idx}>
                        <ListItem>
                        <Typography>
                            {table.title}
                        </Typography>
                        {table.hasOwnProperty("addFn") && <Button color="primary" onClick={this.handleAdd.bind(this, table)}>Add</Button>}
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
                            {table.hasOwnProperty("deleteFn") && <TableCell></TableCell>}
                            </TableRow>
                            </TableHead>
                            <TableBody>
                            {((this.state.data[table.title] != null && this.state.data[table.title].length > 0)?
                               this.state.data[table.title] : (table.emptyValue? table.emptyValue : [])).map(function (row, rowidx) {
                                return (
                                    <TableRow key={table.title + "row:" + rowidx}>
                                    {table.columns.map(function (column, colidx) {
                                    return (
                                        <TableCell key={table.title + "row:" + rowidx + "col:" + colidx}>
                                            {row[column.key]}
                                        </TableCell>)}.bind(this))}
                                    {table.hasOwnProperty("deleteFn") && <TableCell><Button color="primary" onClick={this.handleDelete.bind(this, table, row)}>Delete</Button></TableCell>}
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
