import React, { Component } from 'react'
import ClippedDrawer from './components/ClippedDrawer'
import LoginPage from './components/LoginPage'
import OrganizationPage from './components/OrganizationPage'
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import CloudQueueIcon from '@material-ui/icons/CloudQueue';
import Typography from '@material-ui/core/Typography';

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      isLoggedIn: false,
      organizationName: null,
      user: null,
    };

    this.login = this.login.bind(this);
    this.logout = this.logout.bind(this);
    this.setUser = this.setUser.bind(this);
    this.setOrganization = this.setOrganization.bind(this);
  }

  login() {
    this.setState(state => ({ isLoggedIn: true }));
  }
  logout() {
    this.setState(state => ({
      isLoggedIn: false,
      user: null,
      organizationName: null}));
  }

  setUser(user) {
    this.setState(state => ({ user: user}));
  }

  setOrganization(organizationName) {
    this.setState(state => ({ organizationName: organizationName}));
  }

  render() {
    return (
      <div>
        <AppBar position="fixed">
          <Toolbar>
            <CloudQueueIcon />
            &nbsp;&nbsp;&nbsp;&nbsp;
            <Typography variant="h6" color="inherit" noWrap>
                {this.state.organizationName !== null? "Design Cloud / " + this.state.organizationName : "Design Cloud"}
            </Typography>
          </Toolbar>
        </AppBar>
        {this.state.isLoggedIn === false && <LoginPage login={this.login} setUser={this.setUser}/>}
        {(this.state.isLoggedIn === true && this.state.organizationName === null) && <OrganizationPage setOrganization={this.setOrganization} user={this.state.user}/>}
        {(this.state.isLoggedIn === true && this.state.organizationName !== null) &&
        <ClippedDrawer setOrganization={this.setOrganization}
                       logout={this.logout}
                       organizationName={this.state.organizationName}
                       user={this.state.user}/>}
      </div>
    )
  }
}

export default App
