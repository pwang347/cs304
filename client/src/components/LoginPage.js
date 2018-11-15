import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';
import { BASE_API_URL } from "../config";
import Typography from '@material-ui/core/Typography';

const styles = theme => ({
    root: {
        display: 'flex',
      },
    content: {
        flexGrow: 1,
        padding: theme.spacing.unit * 10,
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
            errorMessage: null,
        };
    
        this.handleLogin = this.handleLogin.bind(this);
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
            if (json.error !== null) {
                throw new Error(json.error);
            }
            console.log(json);
            if (json.affectedRows > 0) {
                console.log(JSON.parse(json.data)[0]);
                this.props.setUser(JSON.parse(json.data)[0].emailAddress);
                this.props.login();
            }
        })
        .catch(function(error){
            console.log(error.message)
            self.state.errorMessage = error.message;
        });
    }

    handleEmailChange(e) {
        this.setState({ emailAddress: e.target.value });
    }

    handlePasswordChange(e) {
        this.setState({ password: e.target.value });
    }

    render() {
        const { classes } = this.props;

        return (
            <div className={classes.root}>
            <main className={classes.content}>
                <TextField
                    required
                    id="emailAddress"
                    label="Email Address"
                    floatinglabeltext="Email Address"
                    margin="normal"
                    value={this.state.emailAddress}
                    onChange={this.handleEmailChange.bind(this)} 
                />
                <br/>
                <TextField
                    required
                    id="password"
                    label="Password"
                    floatinglabeltext="Password"
                    margin="normal"
                    value={this.state.password}
                    onChange={this.handlePasswordChange.bind(this)}
                />
                <Button variant="contained" color="primary" onClick={this.handleLogin}>
                    Login
                </Button>
                <Button variant="contained" color="primary">
                    Register
                </Button>
                {this.state.errorMessage !== null && <Typography>
                    {this.state.errorMessage}
                </Typography>}
            </main>
    </div>);
    }
}

export default withStyles(styles)(LoginPage);