import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import { BASE_API_URL } from "../config";
import Typography from '@material-ui/core/Typography';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';

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

class LoginPage extends React.Component {

    constructor(props) {
        super(props);
    
        this.state = {
            emailAddress: "",
            password: "",
            errorMessage: "",
            showRegister: false,
            firstName: "",
            lastName: "",
            twoFactorPhoneNumber: "",
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
                self.props.setUser(JSON.parse(json.data)[0].emailAddress);
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

    handleRegister() {
        if (this.state.showRegister === false) {
            this.setState({ showRegister: true });
            return;
        }
        var emailAddress = this.state.emailAddress;
        var url = BASE_API_URL + "/user/create?emailAddress=" + emailAddress
        + "&passwordHash=" + this.state.password
        + "&firstName=" + this.state.firstName
        + "&lastName=" + this.state.lastName
        + "&isAdmin=false"
        + "&twoFactorPhoneNumber=" + this.state.twoFactorPhoneNumber;
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
                self.props.setUser(emailAddress);
                self.props.login();
            }
            else {
                throw new Error("User with the same email or phone number already exists.");
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

    handleFirstNameChange(e) {
        this.setState({ firstName: e.target.value });
    }

    handleLastNameChange(e) {
        this.setState({ lastName: e.target.value });
    }

    handleTwoFactorPhoneNumberChange(e) {
        this.setState({ twoFactorPhoneNumber: e.target.value });
    }

    render() {
        const { classes } = this.props;

        return (
            <div className={classes.root}>
            <main className={classes.content}>
            <List>
                <ListItem>
                <TextField
                    required
                    id="emailAddress"
                    label="Email Address"
                    floatinglabeltext="Email Address"
                    margin="normal"
                    value={this.state.emailAddress}
                    onChange={this.handleEmailChange.bind(this)}
                />
                <TextField
                    required
                    id="password"
                    label="Password"
                    floatinglabeltext="Password"
                    margin="normal"
                    value={this.state.password}
                    onChange={this.handlePasswordChange.bind(this)}
                />
                </ListItem>
                {this.state.showRegister === true && <ListItem>
                <TextField
                    required
                    id="firstName"
                    label="First Name"
                    floatinglabeltext="First Name"
                    margin="normal"
                    value={this.state.firstName}
                    onChange={this.handleFirstNameChange.bind(this)}
                />
                <TextField
                    required
                    id="lastName"
                    label="Last Name"
                    floatinglabeltext="Last Name"
                    margin="normal"
                    value={this.state.lastName}
                    onChange={this.handleLastNameChange.bind(this)}
                />
                <TextField
                    required
                    id="twoFactorPhoneNumber"
                    label="Phone Number"
                    floatinglabeltext="Phone Number"
                    margin="normal"
                    value={this.state.twoFactorPhoneNumber}
                    onChange={this.handleTwoFactorPhoneNumberChange.bind(this)}
                />
                </ListItem>}
                <ListItem>
                <Button variant="contained" color="primary" onClick={this.handleLogin.bind(this)}>
                    Login
                </Button>
                <Button variant="contained" color="primary" onClick={this.handleRegister.bind(this)}>
                    Register
                </Button>
                </ListItem>
                <ListItem>
                    <Typography>
                        {this.state.errorMessage}
                    </Typography>
                </ListItem>
            </List>
            </main>
    </div>);
    }
}

export default withStyles(styles)(LoginPage);