import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import { BASE_API_URL } from "../config";
import Typography from '@material-ui/core/Typography';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import CollectionPicker from './CollectionPicker';

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
            newOrganizationName: null,
            errorMessage: "",
            showCreate: false,
            collectionPickerDialog: null,
        };
    }

    handleClickOpen = () => {
        this.setState({
            collectionPickerDialog: {
                title: "Select an organization",
                onClose: this.handleClose.bind(this),
                dataEndpoint: "/organization/listUser?userEmailAddress=" + this.props.userEmailAddress,
                displayfn: (pair) => pair.organizationName,
                keyfn: (pair) => pair.organizationName,
            },
          open: true,
        });
      };
    
    handleClose = value => {
        this.setState({open: false});
        if (value) {
            this.props.setOrganization(value);
        }
    };

    handleCreate = () => {
        if (this.state.showCreate === false) {
            this.setState({ showCreate: true });
            return;
        }
        var newOrganizationName = this.state.newOrganizationName;
        var url = BASE_API_URL + "/organization/create?name=" + newOrganizationName
        + "&contactEmailAddress=" + this.props.userEmailAddress;
        var self = this;
        fetch(url)
        .then(function(response) {
            return response.json();
        })
        .then(function(json) {
            if (json.hasOwnProperty("error")) {
                throw new Error(json.error);
            }
            if (json.affectedRows > 0) {
                self.props.setOrganization(newOrganizationName);
            }
            else {
                throw new Error("Organization with the same name already exists.");
            }
        })
        .catch(function(error){
            self.setState({ errorMessage: error.message });
        });
    }

    handleOrganizationNameChange(e) {
        this.setState({ newOrganizationName: e.target.value });
    }

    render() {
        const { classes } = this.props;

        return (
            <div className={classes.root}>
            <main className={classes.content}>
            <List>
                <ListItem>
                    <Button onClick={this.handleClickOpen}>Select an organization</Button>
                    {this.state.collectionPickerDialog && <CollectionPicker open={this.state.open} dialog={this.state.collectionPickerDialog}/>}
                    <Button variant="contained" color="primary" onClick={this.handleCreate}>
                        Create new organization
                    </Button>
                </ListItem>
                {this.state.showCreate === true && <ListItem>
                <TextField
                    required
                    id="organizationName"
                    label="Organization Name"
                    floatinglabeltext="Organization Name"
                    margin="normal"
                    value={this.state.firstName}
                    onChange={this.handleOrganizationNameChange.bind(this)}
                />
                </ListItem>}
                <ListItem>
                    <Typography>
                        {this.state.errorMessage}
                    </Typography>
                </ListItem>
            </List>
            </main>
          </div>
        );
      }
}

export default withStyles(styles)(OrganizationPage);