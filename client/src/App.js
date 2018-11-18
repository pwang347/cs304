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
      isLoggedIn: true,
      organizationName: "Macrohard",
      userEmailAddress: "a",
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
      userEmailAddress: null,
      organizationName: null}));
  }

  setUser(userEmailAddress) {
    this.setState(state => ({ userEmailAddress: userEmailAddress}));
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
                Design Cloud
            </Typography>
          </Toolbar>
        </AppBar>
        {this.state.isLoggedIn === false && <LoginPage login={this.login} setUser={this.setUser}/>}
        {(this.state.isLoggedIn === true && this.state.organizationName === null) && <OrganizationPage setOrganization={this.setOrganization} userEmailAddress={this.state.userEmailAddress}/>}
        {(this.state.isLoggedIn === true && this.state.organizationName !== null) && <ClippedDrawer setOrganization={this.setOrganization} userEmailAddress={this.state.userEmailAddress} logout={this.logout} organizationName={this.state.organizationName}/>}
      </div>
    )
  }
}

export default App
