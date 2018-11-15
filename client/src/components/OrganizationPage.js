import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import { BASE_API_URL } from "../config";
import Typography from '@material-ui/core/Typography';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import OrganizationPicker from './OrganizationPicker';

const styles = theme => ({
    root: {
        display: 'flex',
      },
    content: {
        flexGrow: 1,
        padding: theme.spacing.unit * 30,
        justifyContent: 'center',
        alignItems:'center',
        display: 'flex',
    },
});

class OrganizationPage extends React.Component {

    constructor(props) {
        super(props);
    
        this.state = {
            open: false,
            organizationName: null,
        };
    }

    handleClickOpen = () => {
        this.setState({
          open: true,
        });
      };
    
    handleClose = value => {
        this.setState({ selectedValue: value, open: false });
        this.props.setOrganization(value);
    };

    render() {
        const { classes } = this.props;

        return (
            <div className={classes.root}>
            <main className={classes.content}>
            <Typography variant="subtitle1">Selected: {this.state.organizationName}</Typography>
            <br />
            <Button onClick={this.handleClickOpen}>Open simple dialog</Button>
            <OrganizationPicker
              selectedValue={this.state.organizationName}
              open={this.state.open}
              onClose={this.handleClose}
              userEmailAddress={this.props.userEmailAddress}
            />
            </main>
          </div>
        );
      }
}

export default withStyles(styles)(OrganizationPage);