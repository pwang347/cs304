DROP DATABASE IF EXISTS cs304;
CREATE DATABASE cs304;
USE cs304;

CREATE TABLE User
(
    emailAddress VARCHAR(255),
    firstName VARCHAR(20),
    lastName VARCHAR(20),
    passwordHash VARCHAR(255),
    isAdmin BOOLEAN,
    twoFactorPhoneNumber VARCHAR(20),
    PRIMARY KEY (emailAddress),
    UNIQUE (twoFactorPhoneNumber)
);

CREATE TABLE Organization
(
    name VARCHAR(255),
    createdTimestamp TIMESTAMP,
    contactEmailAddress VARCHAR(255) NOT NULL,
    PRIMARY KEY (name),
    FOREIGN KEY (contactEmailAddress) REFERENCES User(emailAddress)
        ON DELETE NO ACTION
        ON UPDATE CASCADE,
    UNIQUE (createdTimestamp, contactEmailAddress)
);

CREATE TABLE UserOrganizationPairs
(
    organizationName VARCHAR(255),
    userEmailAddress VARCHAR(255),
    PRIMARY KEY (organizationName, userEmailAddress),
    FOREIGN KEY (organizationName) REFERENCES Organization(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (userEmailAddress) REFERENCES User(emailAddress)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE AccessGroup
(
    name VARCHAR(255),
    organizationName VARCHAR(255),
    PRIMARY KEY (name, organizationName),
    FOREIGN KEY (organizationName) REFERENCES Organization(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE UserAccessGroupPairs
(
    accessGroupName VARCHAR(255),
    accessGroupOrganizationName VARCHAR(255),
    userEmailAddress VARCHAR(255),
    PRIMARY KEY (accessGroupName, accessGroupOrganizationName, userEmailAddress),
    FOREIGN KEY (accessGroupName, accessGroupOrganizationName) REFERENCES AccessGroup(name, organizationName)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (userEmailAddress) REFERENCES User(emailAddress)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE Service
(
    name VARCHAR(255),
    description VARCHAR(255),
    isPreview BOOLEAN,
    isEnabled BOOLEAN,
    isVirtualMachineService BOOLEAN,
    imageUrl VARCHAR(255),
    PRIMARY KEY (name)
);

CREATE TABLE Region
(
    name VARCHAR(255),
    PRIMARY KEY (name)
);

CREATE TABLE ServiceInstance
(
    name VARCHAR(255),
    regionName VARCHAR(255) NOT NULL,
    serviceName VARCHAR(255),
    organizationName VARCHAR(255),
    PRIMARY KEY (name, serviceName, organizationName),
    FOREIGN KEY (serviceName) REFERENCES Service(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (organizationName) REFERENCES Organization(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (regionName) REFERENCES Region(name)
        ON DELETE NO ACTION
        ON UPDATE CASCADE
);

CREATE TABLE ServiceInstanceConfiguration
(
    configKey VARCHAR(255),
    serviceInstanceName VARCHAR(255),
    serviceInstanceServiceName VARCHAR(255),
    serviceInstanceOrganizationName VARCHAR(255),
    data TEXT,
    PRIMARY KEY (configKey, serviceInstanceName, serviceInstanceServiceName, serviceInstanceOrganizationName),
    FOREIGN KEY (serviceInstanceName, serviceInstanceServiceName, serviceInstanceOrganizationName) REFERENCES ServiceInstance(name, serviceName, organizationName)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE ServiceInstanceKey
(
    keyValue VARCHAR(255),
    activeUntil TIMESTAMP,
    serviceInstanceName VARCHAR(255),
    serviceInstanceServiceName VARCHAR(255),
    serviceInstanceOrganizationName VARCHAR(255),
    PRIMARY KEY (keyValue, serviceInstanceName, serviceInstanceServiceName, serviceInstanceOrganizationName),
    FOREIGN KEY (serviceInstanceName, serviceInstanceServiceName, serviceInstanceOrganizationName) REFERENCES ServiceInstance(name, serviceName, organizationName)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE ServiceInstanceAccessGroupPermissions
(
    serviceInstanceName VARCHAR(255),
    serviceInstanceServiceName VARCHAR(255),
    serviceInstanceOrganizationName VARCHAR(255),
    accessGroupName VARCHAR(255),
    accessLevel INT,
    PRIMARY KEY (serviceInstanceName, serviceInstanceServiceName, serviceInstanceOrganizationName, accessGroupName),
    FOREIGN KEY (serviceInstanceName, serviceInstanceServiceName, serviceInstanceOrganizationName) REFERENCES ServiceInstance(name, serviceName, organizationName)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (accessGroupName) REFERENCES AccessGroup(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE BaseImage
(
    os VARCHAR(255),
    version VARCHAR(255),
    PRIMARY KEY (os, version)
);

CREATE TABLE VirtualMachine
(
    description TEXT,
    ipAddress VARCHAR(255),
    state INT,
    cores INT,
    diskSpace INT,
    ram INT,
    baseImageOs VARCHAR(255) NOT NULL,
    baseImageVersion VARCHAR(255) NOT NULL,
    regionName VARCHAR(255) NOT NULL,
    organizationName VARCHAR(255) NOT NULL,
    virtualMachineServiceName VARCHAR(255) NOT NULL,
    PRIMARY KEY (ipAddress),
    FOREIGN KEY (baseImageOs, baseImageVersion) REFERENCES BaseImage(os, version)
        ON DELETE NO ACTION
        ON UPDATE CASCADE,
    FOREIGN KEY (regionName) REFERENCES Region(name)
        ON DELETE NO ACTION
        ON UPDATE CASCADE,
    FOREIGN KEY (organizationName) REFERENCES Organization(name)
        ON DELETE NO ACTION
        ON UPDATE CASCADE,
    FOREIGN KEY (virtualMachineServiceName) REFERENCES Service(name)
        ON DELETE NO ACTION
        ON UPDATE CASCADE
);

CREATE TABLE EventLog
(
    logNumber INT,
    timestamp TIMESTAMP,
    data TEXT,
    eventType VARCHAR(255),
    VirtualMachineIpAddress VARCHAR(255),
    PRIMARY KEY (logNumber, VirtualMachineIpAddress),
    FOREIGN KEY (VirtualMachineIpAddress) REFERENCES VirtualMachine(ipAddress)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE VirtualMachineAccessGroupPermissions
(
    VirtualMachineIpAddress VARCHAR(255),
    accessGroupOrganizationName VARCHAR(255),
    accessGroupName VARCHAR(255),
    accessLevel INT,
    PRIMARY KEY (VirtualMachineIpAddress, accessGroupOrganizationName, accessGroupName),
    FOREIGN KEY (VirtualMachineIpAddress) REFERENCES VirtualMachine(ipAddress)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (accessGroupOrganizationName, accessGroupName) REFERENCES AccessGroup(organizationName, name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE ServiceSubscriptionTransaction
(
    type INT,
    serviceName VARCHAR(255),
    organizationName VARCHAR(255),
    description VARCHAR(255),
    activeUntil DATETIME,
    transactionNumber INT,
    amountPaid INT,
    processedTimestamp DATETIME,
    PRIMARY KEY (transactionNumber, organizationName),
    UNIQUE (type, serviceName, organizationName),
    FOREIGN KEY (organizationName) REFERENCES Organization(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE CreditCard
(
    cardNumber CHAR(20),
    cvc CHAR(3),
    expiryDate DATE,
    cardType INT,
    PRIMARY KEY (cardNumber)
);

CREATE TABLE OrganizationCreditCardPairs
(
    organizationName VARCHAR(255),
    creditCardNumber CHAR(20),
    PRIMARY KEY (organizationName, creditCardNumber),
    FOREIGN KEY (organizationName) REFERENCES Organization(name)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    FOREIGN KEY (creditCardNumber) REFERENCES CreditCard(cardNumber)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);
