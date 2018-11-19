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

const styles = theme => ({
});

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
                var staticdata = table.staticdata.map((d) => Object.assign({}, d));
                for (var idx in staticdata) {
                    for (var key in staticdata[idx]) {
                        staticdata[idx][key] = typeof staticdata[idx][key] === "function" ? staticdata[idx][key]() : staticdata[idx][key];
                    }
                }
                this.state.data[table.title] = staticdata;
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

    handleUpdate = (table, row) => {
        table.updateFn(row, this.load.bind(this));
    }

    handleAction = (action, row) => {
        action(row, this.load.bind(this));
    }
  
    render() {
      const { classes, onClose, dialog, ...other } = this.props;
      return (
        <Dialog onEnter={this.handleOpen} onClose={this.handleClose} aria-labelledby="simple-dialog-title" className={classes.root} {...other}>
          <DialogTitle>{dialog.title}</DialogTitle>
          <DialogContent>
              <div>
                {dialog.tables.map(function (table, idx) {
                return (
                    <List key={idx} width="100%">
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
                            {table.hasOwnProperty("updateFn") && <TableCell></TableCell>}
                            {table.hasOwnProperty("customActions") && table.customActions.map(function(action, idx){
                                        return (
                                            <TableCell key={"action " + idx}></TableCell>
                                        )
                                    })}
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
                                    {table.hasOwnProperty("updateFn") && <TableCell><Button color="primary" onClick={this.handleUpdate.bind(this, table, row)}>Update</Button></TableCell>}
                                    {table.hasOwnProperty("customActions") && table.customActions.map(function(action, idx){
                                        return (
                                            <TableCell key={"action" +idx}><Button color="primary"
                                            onClick={this.handleAction.bind(this, action.hasOwnProperty("nestedFn")? action.fn() : action.fn, row)}>
                                            {typeof action.name === "function"? action.name() : action.name}</Button></TableCell>
                                        )
                                    }.bind(this))}
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
