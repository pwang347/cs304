import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import { BASE_API_URL } from "../config";
import Typography from '@material-ui/core/Typography';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import CreationDialog from './CreationDialog';

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
    button: {
        marginRight: theme.spacing.unit,
        width: 200,
    },
    title: {
        margin: theme.spacing.unit * 10,
    },
    subtitle: {
        marginLeft: theme.spacing.unit * 10,
    },
    textField: {
        width: 400,
        marginTop: -10,
    },
    loginArea: {
        marginTop: theme.spacing.unit * 10,
    }
});

class LoginPage extends React.Component {

    constructor(props) {
        super(props);
    
        this.state = {
            emailAddress: "",
            password: "",
            errorMessage: "",
            creationDialog: null,
        };
    }

    handleLogin() {
        var url = BASE_API_URL + "/user/login?emailAddress=" + this.state.emailAddress
        + "&passwordHash=" + this.state.password;
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
                self.props.setUser(JSON.parse(json.data)[0]);
                self.props.login();
            }
            else {
                throw new Error("Invalid credentials provided.");
            }
        })
        .catch(function(error){
            self.setState({ errorMessage: error.message });
        });
    }

    handleEmailChange(e) {
        this.setState({ emailAddress: e.target.value });
    }

    handlePasswordChange(e) {
        this.setState({ password: e.target.value });
    }

    handleCloseForCreateUser = (result) => {
        this.setState(state => ({
            creationDialog: null,
        }));
        if (!result) return;
        var url = BASE_API_URL + "/user/create?emailAddress=" + result["Email address"]
        + "&firstName=" + result["First name"]
        + "&lastName=" + result["Last name"]
        + "&passwordHash=" + result["Password"]
        + "&twoFactorPhoneNumber=" + result["Phone number"];
        var self = this;
        fetch(url)
            .then(function (response) {
                if (response.status >= 400) {
                    throw new Error("Bad response from server");
                }
                return response.json();
            })
            .then(function(json) {
                if (json.hasOwnProperty("error")) {
                    throw new Error(json.error);
                }
                if (json.affectedRows !== 1) {
                    throw new Error("Could not add user.");
                }
                self.state.emailAddress = result["Email address"];
                self.state.password = result["Password"];
                self.handleLogin();
            });
    }

    handleCreateUser = () => {
        this.setState(state => ({creationDialog: {
            titleText: "Register",
            createText: "Finish",
            fields: [{name: "Email address"}, {name: "First name"}, {name: "Last name"}, {name: "Password"}, {name: "Phone number"}],
            onClose: this.handleCloseForCreateUser.bind(undefined),
        }}));
    }

    render() {
        const { classes } = this.props;

        return (
            <div className={classes.root}>
            <main className={classes.content}>
            <List>
                <ListItem>
                    <Typography className={classes.title} variant="h2">
                    Welcome to Design Cloud.
                    </Typography>
                </ListItem>
                <ListItem>
                    <Typography className={classes.subtitle} variant="h4">
                        We probably run SQL queries and stuff.
                    </Typography>
                </ListItem>
            </List>
            <List className={classes.loginArea}>
                <ListItem>
                <TextField
                    required
                    className={classes.textField}
                    id="emailAddress"
                    label="Email Address"
                    floatinglabeltext="Email Address"
                    margin="normal"
                    value={this.state.emailAddress}
                    onChange={this.handleEmailChange.bind(this)}
                />
                </ListItem>
                <ListItem>
                <TextField
                    className={classes.textField}
                    required
                    type="password"
                    id="password"
                    label="Password"
                    floatinglabeltext="Password"
                    margin="normal"
                    value={this.state.password}
                    onChange={this.handlePasswordChange.bind(this)}
                />
                </ListItem>
                <ListItem>
                <Button className={classes.button} variant="contained" color="primary" onClick={this.handleLogin.bind(this)}>
                    Login
                </Button>
                <Button className={classes.button} variant="contained" color="primary" onClick={this.handleCreateUser.bind(this)}>
                    Register
                </Button>
                </ListItem>
                <ListItem>
                    <Typography>
                        {this.state.errorMessage}
                    </Typography>
                </ListItem>
            </List>
            {this.state.creationDialog !== null &&
                    <CreationDialog dialog={this.state.creationDialog}/>}
            </main>
    </div>);
    }
}

export default withStyles(styles)(LoginPage);