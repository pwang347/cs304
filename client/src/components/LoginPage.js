import React from 'react';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';

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

    render() {
        const { classes } = this.props;

        return (
            <div className={classes.root}>
            <main className={classes.content}>
                <TextField
                    required
                    id="standard-required"
                    label="Email Address"
                    floatinglabeltext="Email Address"
                    margin="normal"
                />
                <br/>
                <TextField
                    required
                    id="standard-required"
                    label="Password"
                    floatinglabeltext="Password"
                    margin="normal"
                />
                <Button variant="contained" color="primary" onClick={this.props.login}>
                    Login
                </Button>
                <Button variant="contained" color="primary">
                    Register
                </Button>
            </main>
    </div>);
    }
}

export default withStyles(styles)(LoginPage);