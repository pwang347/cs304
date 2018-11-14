import React, { Component } from 'react'
import ClippedDrawer from './components/ClippedDrawer'
import LoginPage from './components/LoginPage'
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import CloudQueueIcon from '@material-ui/icons/CloudQueue';
import Typography from '@material-ui/core/Typography';

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      isLoggedIn: false,
    };

    this.login = this.login.bind(this);
    this.logout = this.logout.bind(this);
  }

  login() {
    this.setState(state => ({ isLoggedIn: true }));
  }
  logout() {
    this.setState(state => ({ isLoggedIn: false}));
  }

  render() {
    return (
      <div>
        <AppBar position="fixed">
          <Toolbar>
            <CloudQueueIcon />
            &nbsp;&nbsp;&nbsp;&nbsp;
            <Typography variant="h6" color="inherit" noWrap>
                Design Cloud
            </Typography>
          </Toolbar>
        </AppBar>
        {this.state.isLoggedIn === false && <LoginPage login={this.login}/>}
        {this.state.isLoggedIn === true && <ClippedDrawer logout={this.logout}/>}
      </div>
    )
  }
}

export default App
