import React from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import Drawer from '@material-ui/core/Drawer';
import CssBaseline from '@material-ui/core/CssBaseline';
import List from '@material-ui/core/List';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemSecondaryAction from '@material-ui/core/ListItemSecondaryAction';
import DomainIcon from '@material-ui/icons/Domain';
import ShopIcon from '@material-ui/icons/Shop';
import PersonIcon from '@material-ui/icons/Person';
import PeopleIcon from '@material-ui/icons/People';
import ExpandLess from '@material-ui/icons/ExpandLess';
import ExpandMore from '@material-ui/icons/ExpandMore';
import ExitToApp from '@material-ui/icons/ExitToApp';
import Collapse from '@material-ui/core/Collapse';
import ViewModule from '@material-ui/icons/ViewModule';
import Camera from '@material-ui/icons/Camera';
import Computer from '@material-ui/icons/Computer';
import AttachMoney from '@material-ui/icons/AttachMoney';
import CreditCardIcon from '@material-ui/icons/CreditCard';
import {BASE_API_URL} from "../config";
import Card from '@material-ui/core/Card';
import CardActionArea from '@material-ui/core/CardActionArea';
import CardActions from '@material-ui/core/CardActions';
import CardContent from '@material-ui/core/CardContent';
import CardMedia from '@material-ui/core/CardMedia';
import Button from '@material-ui/core/Button';
import Paper from '@material-ui/core/Paper';
import Grid from '@material-ui/core/Grid';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import ViewServiceInstanceDialog from './ViewServiceInstanceDialog';
import ConfirmationDialog from './ConfirmationDialog';
import CollectionPicker from './CollectionPicker';
import CreationDialog from './CreationDialog';
import PlayArrow from '@material-ui/icons/PlayArrow';

const drawerWidth = 300;

const styles = theme => ({
    root: {
        display: 'flex',
    },
    drawer: {
        zIndex: theme.zIndex.appBar - 1,
        width: drawerWidth,
        flexShrink: 0,
    },
    drawerPaper: {
        width: drawerWidth,
    },
    content: {
        flexGrow: 1,
        padding: theme.spacing.unit * 3,
    },
    toolbar: theme.mixins.toolbar,
    nested: {
        paddingLeft: theme.spacing.unit * 4,
    },
    nestedTwice: {
        paddingLeft: theme.spacing.unit * 8,
    },
    card: {
        width: 300,
        minHeight: 300,
        display: 'inline-block',
        margin: theme.spacing.unit * 2,
    },
    media: {
        height: 140,
    },
    accessGroup: {
        width: drawerWidth,
        margin: theme.spacing.unit,
    },
});

const defaultImageUrl = "https://material-ui.com/static/images/cards/contemplative-reptile.jpg";

class ClippedDrawer extends React.Component {
    state = {
        serviceOpen: true,
        organizationOpen: false,
        myServicesOpen: false,
        activePageId: "store",
        activeServiceName: null,
        services: [],
        servicesMap: {},
        regions: [],
        organizationServiceInstances: [],
        organizationVirtualMachines: [],
        organizationActiveSubscriptions: [],
        organizationTransactions: [],
        organizationAccessGroups: [],
        activeSubscriptionsMap: {},
        displayedServiceInstance: null,
        serviceInstancesMap: {},
        accessGroupsMap: {},
        confirmationDialog: null,
        creationDialog: null,
        addingToGroup: null,
    };

    componentDidMount() {
        this.loadRegions();
        this.loadServices();
        this.loadOrganizationVirtualMachines();
        this.loadActiveServices();
        this.loadTransactions();
        this.loadAccessGroups();
        this.loadAccessGroupUsers();
    }

    handleClick = (id, data) => {
        return () => {
            if (id === "store") {
                this.loadServices();
            }
            if (id === "instances") {
                this.setState(state => ({activeServiceName: data}));
                this.loadOrganizationServiceInstances(data);
                this.setState(state => ({serviceOpen: true}));
                this.setState(state => ({myServicesOpen: true}));
            }
            if (id === "virtual-machines") {
                this.loadOrganizationVirtualMachines();
            }
            if (id === "billing") {
                this.loadActiveServices();
                this.loadTransactions();
            }
            if (id === "access-groups"){
                this.loadAccessGroups();
                this.loadAccessGroupUsers();
            }

            if (id === "service") {
                this.setState(state => ({serviceOpen: !state.serviceOpen}));
            }
            else if (id === "organization") {
                this.setState(state => ({organizationOpen: !state.organizationOpen}));
            }
            else if (id === "my-services") {
                this.setState(state => ({myServicesOpen: !state.myServicesOpen}));
            }
            else {
                this.setState(state => ({activePageId: id}));
            }
        }
    };

    loadRegions = () => {
        var url = BASE_API_URL + "/region/list";
        var self = this;
        fetch(url)
            .then(function (response) {
                if (response.status >= 400) {
                    throw new Error("Bad response from server");
                }
                return response.json();
            })
            .then(function (json) {
                if (json.affectedRows > 0) {
                    self.setState({
                        regions: JSON.parse(json.data)
                    });
                }
            });
    }

    loadActiveServices = () => {
        var url = BASE_API_URL + "/serviceSubscriptionTransaction/listActiveSubscriptions?organizationName=" + this.props.organizationName;
        var self = this;
        fetch(url)
            .then(function (response) {
                if (response.status >= 400) {
                    throw new Error("Bad response from server");
                }
                return response.json();
            })
            .then(function (json) {
                if (json.affectedRows > 0) {
                    var organizationActiveSubscriptions =  JSON.parse(json.data);
                    self.setState({
                        organizationActiveSubscriptions: organizationActiveSubscriptions
                    });
                    for (var activeSubscription of organizationActiveSubscriptions) {
                        self.state.activeSubscriptionsMap[activeSubscription.serviceName] = activeSubscription;
                    }
                }
            });
    }

    loadTransactions = () => {
        var url = BASE_API_URL + "/serviceSubscriptionTransaction/listTransactions?organizationName=" + this.props.organizationName;
        var self = this;
        fetch(url)
            .then(function (response) {
                if (response.status >= 400) {
                    throw new Error("Bad response from server");
                }
                return response.json();
            })
            .then(function (json) {
                if (json.affectedRows > 0) {
                    self.setState({
                        organizationTransactions: JSON.parse(json.data)
                    });
                }
            });
    }

    loadOrganizationServiceInstances = (serviceName) => {
        var url = BASE_API_URL + "/serviceInstance/listServiceOrganization?organizationName=" + this.props.organizationName
        + "&serviceName=" + serviceName;
        var self = this;
        fetch(url)
            .then(function (response) {
                if (response.status >= 400) {
                    throw new Error("Bad response from server");
                }
                return response.json();
            })
            .then(function (json) {
                var organizationServiceInstances = JSON.parse(json.data);
                self.setState({
                    organizationServiceInstances: organizationServiceInstances
                });
                for (var serviceInstance of organizationServiceInstances) {
                    self.state.serviceInstancesMap[serviceInstance.name] = serviceInstance;
                }
            });
    }

    loadOrganizationVirtualMachines = () => {
        var url = BASE_API_URL + "/virtualMachine/listOrganization?organizationName=" + this.props.organizationName;
        var self = this;
        fetch(url)
            .then(function (response) {
                if (response.status >= 400) {
                    throw new Error("Bad response from server");
                }
                return response.json();
            })
            .then(function (json) {
                self.setState({
                    organizationVirtualMachines: JSON.parse(json.data)
                });
            });
    }

    loadAccessGroups = () => {
        var url = BASE_API_URL + "/accessGroup/listOrganization?organizationName=" + this.props.organizationName;
        var self = this;
        fetch(url)
            .then(function (response) {
                if (response.status >= 400) {
                    throw new Error("Bad response from server");
                }
                return response.json();
            })
            .then(function (json) {
                self.setState({
                    organizationAccessGroups: JSON.parse(json.data)
                });
            });
    }

    loadAccessGroupUsers = () => {
        var url = BASE_API_URL + "/accessGroup/listUsersForOrganization?organizationName=" + this.props.organizationName;
        var self = this;
        fetch(url)
            .then(function (response) {
                if (response.status >= 400) {
                    throw new Error("Bad response from server");
                }
                return response.json();
            })
            .then(function (json) {
                for (var member in self.state.accessGroupsMap) delete self.state.accessGroupsMap[member];
                var userGroupPairings = JSON.parse(json.data);
                for (var pairing of userGroupPairings) {
                    if(!self.state.accessGroupsMap.hasOwnProperty(pairing.accessGroupName)) {
                        self.state.accessGroupsMap[pairing.accessGroupName] = [];
                    }
                    self.state.accessGroupsMap[pairing.accessGroupName].push(pairing.userEmailAddress);
                }
            });
    }

    loadServices = () => {
        var url = BASE_API_URL + "/service/list";
        var self = this;
        fetch(url)
            .then(function (response) {
                if (response.status >= 400) {
                    throw new Error("Bad response from server");
                }
                return response.json();
            })
            .then(function (json) {
                var services = JSON.parse(json.data);
                self.setState({
                    services: services
                });
                for (var service of services) {
                    self.state.servicesMap[service.name] = service;
                }
            });
    }

    handleServiceInstanceDetails = (serviceInstanceName) => {
        this.setState(state => ({displayedServiceInstance: this.state.serviceInstancesMap[serviceInstanceName]}));
    }

    handleServiceInstanceDetailsClose = () => {
        this.setState(state => ({displayedServiceInstance: null}));
    }

    handleCloseForDeleteServiceInstance = (serviceInstanceName, result) => {
        this.setState(state => ({
            confirmationDialog: null,
        }));
        if (!result) return;
        this.setState(state => ({deleteServiceInstanceConfirmed: false}));
        var serviceInstance = this.state.serviceInstancesMap[serviceInstanceName];
        var url = BASE_API_URL + "/serviceInstance/delete?name=" + serviceInstance.name
        + "&serviceName=" + serviceInstance.serviceName
        + "&organizationName=" + this.props.organizationName;
        var self = this;
        fetch(url)
            .then(function (response) {
                if (response.status >= 400) {
                    throw new Error("Bad response from server");
                }
                return response.json();
            })
            .then(function (json) {
                if (json.affectedRows !== 1) {
                    throw new Error("Nothing to delete on server");
                }
        });
        var index = this.state.organizationServiceInstances.indexOf(serviceInstance);
        if (index > -1) {
            this.state.organizationServiceInstances.splice(index, 1);
        }
        delete this.state.serviceInstancesMap[serviceInstanceName];
    }

    handleDeleteServiceInstance = (serviceInstanceName) => {
        if (this.state.confirmationDialog === null && !this.state.deleteServiceInstanceConfirmed) {
            this.setState(state => ({confirmationDialog: {
                titleText: "Are you sure you want to delete " + serviceInstanceName + "?",
                contentText: "This action cannot be reversed.",
                yesText: "Yes",
                noText: "No",
                onClose: this.handleCloseForDeleteServiceInstance.bind(undefined, serviceInstanceName)
            }}));
            return;
        }
    }

    handleClickChangeOrganization = () => {
        this.props.setOrganization(null);
    }

    handleRemoveUserFromAccessGroup = (accessGroupName, userEmailAddress) => {
        var url = BASE_API_URL + "/accessGroup/removeUser?accessGroupName=" + accessGroupName
        + "&userEmailAddress=" + userEmailAddress
        + "&accessGroupOrganizationName=" + this.props.organizationName;
        var self = this;
        fetch(url)
            .then(function (response) {
                if (response.status >= 400) {
                    throw new Error("Bad response from server");
                }
                return response.json();
            })
            .then(function (json) {
                if (json.affectedRows !== 1) {
                    throw new Error("Nothing to delete on server");
                }
        });
        var index = this.state.accessGroupsMap[accessGroupName].indexOf(userEmailAddress);
        if (index > -1) {
            this.state.accessGroupsMap[accessGroupName].splice(index, 1);
        }
        // Trigger update
        this.setState(state => ({accessGroupsMap: this.state.accessGroupsMap}));
    }

    handleAddUserToAccessGroup = (accessGroupName) => {
        this.setState({addingToGroup: accessGroupName});
    }
    
    handleAddingToGroupClose = (value) => {
        var groupName = this.state.addingToGroup;
        if (value) {
            var url = BASE_API_URL + "/accessGroup/addUser?accessGroupName=" + groupName
            + "&accessGroupOrganizationName=" + this.props.organizationName
            + "&userEmailAddress=" + value;
            var self = this;
            fetch(url)
            .then(function(response) {
                return response.json();
            })
            .then(function(json) {
                if (json.hasOwnProperty("error")) {
                    throw new Error(json.error);
                }
                if (json.affectedRows !== 1) {
                    throw new Error("Could not add user.");
                }
                if (!self.state.accessGroupsMap.hasOwnProperty(groupName)) {
                    self.state.accessGroupsMap[groupName] = [];
                }
                self.state.accessGroupsMap[groupName].push(value);
                // Trigger update
                self.setState(state => ({accessGroupsMap: self.state.accessGroupsMap}));
            });
        }
        this.setState({addingToGroup: null});
    }

    handleCloseForDeleteAccessGroup = (accessGroupName, result) => {
        this.setState(state => ({
            confirmationDialog: null,
        }));
        if (!result) return;
        this.setState(state => ({deleteAccessGroupConfirmed: false}));
        var url = BASE_API_URL + "/accessGroup/delete?name=" + accessGroupName
        + "&organizationName=" + this.props.organizationName;
        var self = this;
        fetch(url)
            .then(function (response) {
                if (response.status >= 400) {
                    throw new Error("Bad response from server");
                }
                return response.json();
            })
            .then(function (json) {
                if (json.affectedRows !== 1) {
                    throw new Error("Nothing to delete on server");
                }
        });
        var index = this.state.organizationAccessGroups.map(g => g.name).indexOf(accessGroupName);
        if (index > -1) {
            this.state.organizationAccessGroups.splice(index, 1);
        }
        delete this.state.accessGroupsMap[accessGroupName];
    }

    handleDeleteAccessGroup = (accessGroupName) => {
        if (this.state.confirmationDialog === null && !this.state.deleteAccessGroupConfirmed) {
            this.setState(state => ({confirmationDialog: {
                titleText: "Are you sure you want to delete " + accessGroupName + "?",
                contentText: "This action cannot be reversed.",
                yesText: "Yes",
                noText: "No",
                onClose: this.handleCloseForDeleteAccessGroup.bind(undefined, accessGroupName)
            }}));
            return;
        }
    }

    handleCloseForCreateAccessGroup = (result) => {
        this.setState(state => ({
            creationDialog: null,
        }));
        if (!result) return;
        var url = BASE_API_URL + "/accessGroup/create?name=" + result.Name
        + "&organizationName=" + this.props.organizationName;
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
                    throw new Error("Could not add access group.");
                }
                self.state.accessGroupsMap[result.Name] = [];
                var insertIdx = 0;
                for (var i = 0; i<self.state.organizationAccessGroups.length; i++){
                    insertIdx = i;
                    if (self.state.organizationAccessGroups[i].name.toLowerCase() > result.Name.toLowerCase()) {
                        break;
                    }
                    if (i === self.state.organizationAccessGroups.length - 1) {
                        insertIdx++;
                    }
                }
                if (insertIdx === self.state.organizationAccessGroups.length) {
                    self.state.organizationAccessGroups.push({name: result.Name});
                } else {
                    self.state.organizationAccessGroups.splice(insertIdx, 0, {name: result.Name});
                }
                
                self.setState(state => ({organizationAccessGroups: self.state.organizationAccessGroups}));
                // Trigger update
                self.setState(state => ({accessGroupsMap: self.state.accessGroupsMap}));
            });
    }

    handleCreateAccessGroup = () => {
        this.setState(state => ({creationDialog: {
            titleText: "Create a new access group",
            fields: [{name: "Name"}],
            onClose: this.handleCloseForCreateAccessGroup.bind(undefined),
        }}));
    }

    handleCloseForCreateServiceInstance = (serviceName, result) => {
        this.setState(state => ({
            creationDialog: null,
        }));
        if (!result) return;
        var url = BASE_API_URL + "/serviceInstance/create?name=" + result.Name
        + "&regionName=" + result.Region
        + "&serviceName=" + serviceName
        + "&organizationName=" + this.props.organizationName;
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
                    throw new Error("Could not add service instance.");
                }
                self.loadOrganizationServiceInstances(self.state.activeServiceName);
            });
    }

    handleCreateServiceInstance = () => {
        this.setState(state => ({creationDialog: {
            titleText: "Create a new instance for " + this.state.activeServiceName,
            fields: [{name: "Name"}, {name: "Region", options: this.state.regions, keyfn: r => r.name, displayfn: r => r.name}],
            onClose: this.handleCloseForCreateServiceInstance.bind(undefined, this.state.activeServiceName),
        }}));
    }

    render() {
        const {classes} = this.props;

        return (
            <div className={classes.root}>
                <CssBaseline/>
                <Drawer
                    className={classes.drawer}
                    variant="permanent"
                    classes={{
                        paper: classes.drawerPaper,
                    }}
                >
                    <div className={classes.toolbar}/>
                    <List>
                        <ListItem button onClick={this.handleClick("service")}>
                            <ListItemIcon>
                                <Camera/>
                            </ListItemIcon>
                            <ListItemText inset primary="Services"/>
                            {this.state.serviceOpen ? <ExpandLess/> : <ExpandMore/>}
                        </ListItem>
                        <Collapse in={this.state.serviceOpen} timeout="auto" unmountOnExit>
                            <List component="div" disablePadding>
                                <ListItem button onClick={this.handleClick("store")} className={classes.nested}
                                          selected={this.state.activePageId === "store"}>
                                    <ListItemIcon>
                                        <ShopIcon/>
                                    </ListItemIcon>
                                    <ListItemText inset primary="Store"/>
                                </ListItem>
                                <ListItem button onClick={this.handleClick("my-services")} className={classes.nested}
                                          selected={this.state.activePageId === "my-services"}>
                                    <ListItemIcon>
                                        <ViewModule/>
                                    </ListItemIcon>
                                    <ListItemText inset primary="My services"/>
                                    {this.state.myServicesOpen ? <ExpandLess/> : <ExpandMore/>}
                                </ListItem>
                                <Collapse in={this.state.myServicesOpen} timeout="auto" unmountOnExit>
                                    <List>
                                        {this.state.organizationActiveSubscriptions.map(function(service, idx){
                                            return (
                                                <ListItem button onClick={this.handleClick("instances", service.name)} className={classes.nestedTwice} key={service.name}
                                                    selected={this.state.activePageId === "instances" && this.state.activeServiceName === service.name}>
                                                    <ListItemIcon>
                                                        <PlayArrow/>
                                                    </ListItemIcon>
                                                    <ListItemText inset primary={service.name}/>
                                                </ListItem>
                                            )}.bind(this))}
                                    </List>
                                </Collapse>
                                <ListItem button onClick={this.handleClick("virtual-machines")}
                                          className={classes.nested}
                                          selected={this.state.activePageId === "virtual-machines"}>
                                    <ListItemIcon>
                                        <Computer/>
                                    </ListItemIcon>
                                    <ListItemText inset primary="Virtual Machines"/>
                                </ListItem>
                            </List>
                        </Collapse>
                        <ListItem button onClick={this.handleClick("organization")}>
                            <ListItemIcon>
                                <DomainIcon/>
                            </ListItemIcon>
                            <ListItemText inset primary="Organization"/>
                            {this.state.organizationOpen ? <ExpandLess/> : <ExpandMore/>}
                        </ListItem>
                        <Collapse in={this.state.organizationOpen} timeout="auto" unmountOnExit>
                            <List component="div" disablePadding>
                                <ListItem button onClick={this.handleClick("billing")} className={classes.nested}
                                          selected={this.state.activePageId === "billing"}>
                                    <ListItemIcon>
                                        <AttachMoney/>
                                    </ListItemIcon>
                                    <ListItemText inset primary="Billing"/>
                                </ListItem>
                                <ListItem button onClick={this.handleClick("credit-cards")} className={classes.nested}
                                          selected={this.state.activePageId === "credit-cards"}>
                                    <ListItemIcon>
                                        <CreditCardIcon/>
                                    </ListItemIcon>
                                    <ListItemText inset primary="Credit Cards"/>
                                </ListItem>
                                <ListItem button onClick={this.handleClick("access-groups")} className={classes.nested}
                                          selected={this.state.activePageId === "access-groups"}>
                                    <ListItemIcon>
                                        <PeopleIcon/>
                                    </ListItemIcon>
                                    <ListItemText inset primary="Access Groups"/>
                                </ListItem>
                            </List>
                        </Collapse>
                    </List>
                    <Divider/>
                    <List>
                        <ListItem button key='My profile' onClick={this.handleClick("my-profile")}>
                            <ListItemIcon>
                                <PersonIcon/>
                            </ListItemIcon>
                            <ListItemText primary='My profile'/>
                        </ListItem>
                        <ListItem button key='Sign out' onClick={this.props.logout}>
                            <ListItemIcon>
                                <ExitToApp/>
                            </ListItemIcon>
                            <ListItemText primary='Sign out'/>
                        </ListItem>
                    </List>
                </Drawer>
                <main className={classes.content}>
                    <div className={classes.toolbar}/>
                    {this.state.activePageId === "store" && <div>
                        {this.state.services.map(function (service, idx) {
                            return (<Card className={classes.card} key={service.name}>
                                <CardActionArea>
                                    <CardMedia
                                        className={classes.media}
                                        image={(!service.imageUrl) ? defaultImageUrl : service.imageUrl}
                                        title="Contemplative Reptile"
                                    />
                                    <CardContent>
                                        <Typography gutterBottom variant="h5" component="h2">
                                            {service.name}
                                        </Typography>
                                        <Typography component="p">
                                            {service.description}
                                        </Typography>
                                    </CardContent>
                                </CardActionArea>
                                <CardActions>
                                    {this.state.activeSubscriptionsMap.hasOwnProperty(service.name)?
                                    <Button size="small" color="primary" onClick={this.handleClick("instances", service.name)}>
                                        View
                                    </Button>
                                    :
                                    <Button size="small" color="primary">
                                        Purchase
                                    </Button>}
                                </CardActions>
                            </Card>)
                        }.bind(this))}
                    </div>}
                    {this.state.activePageId === "instances" && <div>
                        {this.state.organizationServiceInstances.map(function (serviceInstance, idx) {
                            return (<Card className={classes.card} key={serviceInstance.name}>
                                <CardActionArea>
                                    <CardMedia
                                        className={classes.media}
                                        image={(!this.state.servicesMap || !this.state.servicesMap.hasOwnProperty(serviceInstance.serviceName)
                                            || !this.state.servicesMap[serviceInstance.serviceName].imageUrl) ? defaultImageUrl : this.state.servicesMap[serviceInstance.serviceName].imageUrl}
                                        title="Contemplative Reptile"
                                    />
                                    <CardContent>
                                        <Typography gutterBottom variant="h5" component="h2">
                                            {serviceInstance.name}
                                        </Typography>
                                        <Typography component="p">
                                            {serviceInstance.serviceName}
                                        </Typography>
                                    </CardContent>
                                </CardActionArea>
                                <CardActions>
                                    <Button size="small" color="primary"
                                            onClick={this.handleServiceInstanceDetails.bind(this, serviceInstance.name)}>
                                        View details
                                    </Button>
                                    <Button size="small" color="primary" onClick={this.handleDeleteServiceInstance.bind(this, serviceInstance.name)}>
                                        Terminate
                                    </Button>
                                </CardActions>
                            </Card>)
                        }.bind(this))}
                        <br/>
                        <Button variant="contained" color="primary" onClick={this.handleCreateServiceInstance}>
                            Create new instance
                        </Button>
                    </div>}
                    {this.state.activePageId === "virtual-machines" && <div>
                        {this.state.organizationVirtualMachines.map(function (virtualMachine, idx) {
                            return (<Card className={classes.card} key={virtualMachine.ipAddress}>
                                <CardActionArea>
                                    <CardMedia
                                        className={classes.media}
                                        image={defaultImageUrl}
                                        title="Contemplative Reptile"
                                    />
                                    <CardContent>
                                        <Typography gutterBottom variant="h5" component="h2">
                                            {virtualMachine.ipAddress}
                                        </Typography>
                                        <Typography component="p">
                                            {virtualMachine.description}
                                        </Typography>
                                    </CardContent>
                                </CardActionArea>
                                <CardActions>
                                    <Button size="small" color="primary">
                                        View details
                                    </Button>
                                    <Button size="small" color="primary">
                                        Terminate
                                    </Button>
                                </CardActions>
                            </Card>)
                        })}
                    </div>}
                    {this.state.activePageId === "billing" &&
                        <Grid container spacing={24}>
                            <Grid item xs={12}>
                                <Typography variant="headline" gutterBottom>Active Subscriptions</Typography>
                                <Paper className={classes.paper}>
                                    <Table className={classes.table}>
                                        <TableHead>
                                            <TableRow>
                                                <TableCell>Service Name</TableCell>
                                                <TableCell numeric>Service Type</TableCell>
                                                <TableCell>Description</TableCell>
                                                <TableCell>Active Until</TableCell>
                                            </TableRow>
                                        </TableHead>
                                        <TableBody>
                                            {this.state.organizationActiveSubscriptions.map(function (activeSub, idx) {
                                                return (
                                                    <TableRow key={activeSub.serviceName}>
                                                        <TableCell component="th" scope="row">
                                                            {activeSub.serviceName}
                                                        </TableCell>
                                                        <TableCell numeric>{activeSub.type}</TableCell>
                                                        <TableCell>{activeSub.description}</TableCell>
                                                        <TableCell>{activeSub.activeUntil}</TableCell>
                                                    </TableRow>
                                                )
                                            })}
                                        </TableBody>
                                    </Table>
                                </Paper>
                            </Grid>
                            <Grid item xs={12} sm={12}>
                                <Typography variant="headline" gutterBottom>Processed Transactions</Typography>
                                <Paper className={classes.paper}>
                                    <Table className={classes.table}>
                                        <TableHead>
                                            <TableRow>
                                                <TableCell>Service Name</TableCell>
                                                <TableCell>Transaction Number</TableCell>
                                                <TableCell numeric>Amount Paid</TableCell>
                                                <TableCell>Processed Date</TableCell>
                                            </TableRow>
                                        </TableHead>
                                        <TableBody>
                                            {this.state.organizationTransactions.map(function (tran, idx) {
                                                return (
                                                    <TableRow key={tran.transactionNumber}>
                                                        <TableCell component="th" scope="row">
                                                            {tran.serviceName}
                                                        </TableCell>
                                                        <TableCell>
                                                            {tran.transactionNumber}
                                                        </TableCell>
                                                        <TableCell numeric>{tran.amountPaid}</TableCell>
                                                        <TableCell>{tran.processedTimestamp}</TableCell>
                                                    </TableRow>
                                                )
                                            })}
                                        </TableBody>
                                    </Table>
                                </Paper>
                            </Grid>
                        </Grid>}
                    {this.state.activePageId === "access-groups" &&
                    <div>
                        <Typography variant="headline" gutterBottom>Access groups</Typography>
                        {this.state.organizationAccessGroups.map(function (accessGroup, idx) {
                        return (
                            <div key={accessGroup.name}>
                            <div>
                            <Typography variant="headline">{accessGroup.name}</Typography>
                            <Button onClick={this.handleDeleteAccessGroup.bind(this, accessGroup.name)}>Delete</Button>
                            </div>
                            <Paper className={classes.accessGroup}> 
                                <List>
                                {(this.state.accessGroupsMap[accessGroup.name] || []).map(function (userEmailAddress, idx){
                                    return (
                                    <ListItem key={accessGroup.name + userEmailAddress}>
                                    <ListItemIcon>
                                        <PersonIcon/>
                                    </ListItemIcon>
                                    <ListItemSecondaryAction>
                                        <Button onClick={this.handleRemoveUserFromAccessGroup.bind(this, accessGroup.name, userEmailAddress)}>x</Button>
                                    </ListItemSecondaryAction>
                                    <ListItemText inset primary={userEmailAddress}/>
                                </ListItem>
                                )}.bind(this))}
                                <ListItem>
                                    <Button onClick={this.handleAddUserToAccessGroup.bind(this, accessGroup.name)}>Add user</Button>
                                </ListItem>
                                </List>
                            </Paper>
                            </div>)}.bind(this))}
                            <Button variant="contained" color="primary" onClick={this.handleCreateAccessGroup}>
                                Create new access group
                            </Button>
                    </div>}
                    {this.state.activePageId === "my-profile" && <Typography paragraph>
                        <Button onClick={this.handleClickChangeOrganization}>
                            Change organization
                        </Button>
                    </Typography>}
                    {this.state.displayedServiceInstance !== null &&
                    <ViewServiceInstanceDialog open={this.state.displayedServiceInstance}
                                               onClose={this.handleServiceInstanceDetailsClose}
                                               serviceInstance={this.state.displayedServiceInstance}/>
                    }
                    {this.state.confirmationDialog !== null &&
                    <ConfirmationDialog dialog={this.state.confirmationDialog}/>
                    }
                    {this.state.creationDialog !== null &&
                    <CreationDialog dialog={this.state.creationDialog}/>}
                    <CollectionPicker  titleText="Select a user"
                        open={this.state.addingToGroup !== null}
                        onClose={this.handleAddingToGroupClose.bind(this)}
                        dataEndpoint={"/organization/listUsersInOrganizationNotInGroup?organizationName=" + this.props.organizationName
                                    + "&accessGroupName=" + this.state.addingToGroup}
                        displayfn={(user) => user.emailAddress}
                        keyfn={(user) => user.emailAddress}>
                    </CollectionPicker>
                    </main>
            </div>
        );
    }
}

ClippedDrawer.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(ClippedDrawer);
