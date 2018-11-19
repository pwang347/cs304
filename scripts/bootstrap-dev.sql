USE cs304;

INSERT INTO User (emailAddress, firstName, lastName, passwordHash, twoFactorPhoneNumber) VALUES ("a", "Bob", "Joe", "b", "6041234567");

INSERT INTO `Organization` (`name`, `createdTimestamp`, `contactEmailAddress`) VALUES ('Macrohard', '2018-11-15 15:32:50', 'a');
INSERT INTO `Organization` (`name`, `createdTimestamp`, `contactEmailAddress`) VALUES ('test', '2018-11-15 18:28:01', 'a');
INSERT INTO `Organization` (`name`, `createdTimestamp`, `contactEmailAddress`) VALUES ('test2', '2018-11-15 18:34:37', 'a');
INSERT INTO `Organization` (`name`, `createdTimestamp`, `contactEmailAddress`) VALUES ('asd', '2018-11-15 18:35:50', 'a');
INSERT INTO `Organization` (`name`, `createdTimestamp`, `contactEmailAddress`) VALUES ('asdsad', '2018-11-15 18:37:26', 'a');
INSERT INTO `Organization` (`name`, `createdTimestamp`, `contactEmailAddress`) VALUES ('asdasf', '2018-11-15 20:20:08', 'a');

INSERT INTO `UserOrganizationPairs` (`organizationName`, `userEmailAddress`, `isAdmin`) VALUES ('asdasf', 'a', true);
INSERT INTO `UserOrganizationPairs` (`organizationName`, `userEmailAddress`, `isAdmin`) VALUES ('Macrohard', 'a', true);

INSERT INTO `Service` (`name`, `description`, `isPreview`, `isEnabled`, `isVirtualMachineService`, `imageUrl`) VALUES ('EC2', 'This is some cool stuff', 0, 1, 1, "https://upload.wikimedia.org/wikipedia/commons/thumb/b/b9/AWS_Simple_Icons_Compute_Amazon_EC2_Instances.svg/300px-AWS_Simple_Icons_Compute_Amazon_EC2_Instances.svg.png");
INSERT INTO `Service` (`name`, `description`, `isPreview`, `isEnabled`, `isVirtualMachineService`, `imageUrl`) VALUES ('S3', 'ASdasdasd', 0, 1, 0, "https://sylvainleroy.com/wp-content/uploads/2018/02/s3.png");
INSERT INTO `Service` (`name`, `description`, `isPreview`, `isEnabled`, `isVirtualMachineService`, `imageUrl`) VALUES ('Compute', 'asdasdasd', 1, 1, 0, "https://www.stratoscale.com/wp-content/uploads/AWS-Lambda.png");

INSERT INTO `BaseImage` (`os`, `version`) VALUES ('Ubuntu', '14.03');

INSERT INTO `Region` (`name`) VALUES ('us-west');

INSERT INTO `VirtualMachine` (`description`, `ipAddress`, `state`, `cores`, `diskSpace`, `ram`, `baseImageOs`, `baseImageVersion`, `regionName`, `organizationName`, `virtualMachineServiceName`) VALUES ('Nobody toucha mah spaghet!', '0.0.0.0', 1, 2, 3, 4, 'Ubuntu', '14.03', 'us-west', 'Macrohard', 'S3');

INSERT INTO `ServiceInstance` (`name`, `regionName`, `serviceName`, `organizationName`) VALUES ('s1', 'us-west', 'S3', 'Macrohard');

INSERT INTO `ServiceInstanceKey` (`keyValue`, `activeUntil`, `serviceInstanceName`, `serviceInstanceServiceName`, `serviceInstanceOrganizationName`) VALUES ('asfas123213', '2018-11-16 17:48:49', 's1', 'S3', 'Macrohard');

INSERT INTO `ServiceInstanceConfiguration` (`configKey`, `serviceInstanceName`, `serviceInstanceServiceName`, `serviceInstanceOrganizationName`, `data`) VALUES ('PORT', 's1', 'S3', 'Macrohard', '8080');

INSERT INTO `AccessGroup` (`name`, `organizationName`) VALUES ('Admins', 'Macrohard');
INSERT INTO `AccessGroup` (`name`, `organizationName`) VALUES ('Merp', 'Macrohard');
INSERT INTO `AccessGroup` (`name`, `organizationName`) VALUES ('Testers', 'Macrohard');

INSERT INTO `UserAccessGroupPairs` (`accessGroupName`, `accessGroupOrganizationName`, `userEmailAddress`) VALUES ('Admins', 'Macrohard', 'a');
INSERT INTO `UserAccessGroupPairs` (`accessGroupName`, `accessGroupOrganizationName`, `userEmailAddress`) VALUES ('Testers', 'Macrohard', 'a');

INSERT INTO `CreditCard` (`cardNumber`, `cvc`, `expiryDate`, `cardType`) VALUE ('5555111122224444', '321', NOW(), 505);

INSERT INTO `OrganizationCreditCardPairs` (`organizationName`, `creditCardNumber`) VALUES ('MacroHard', '5555111122224444');

INSERT INTO `EventLog` (`logNumber`, `timestamp`, `data`, `eventType`, `VirtualMachineIpAddress`) VALUES (1, '2018-10-02 15:04:05', 'something happened', 'info', '0.0.0.0');
INSERT INTO `EventLog` (`logNumber`, `timestamp`, `data`, `eventType`, `VirtualMachineIpAddress`) VALUES (2, '2018-11-10 15:04:05', 'something went wrong', 'error', '0.0.0.0');

INSERT INTO `ServiceType` (`type`, `serviceTypeName`) VALUES (1, 'One time');
INSERT INTO `ServiceType` (`type`, `serviceTypeName`) VALUES (2, 'One month');
INSERT INTO `ServiceType` (`type`, `serviceTypeName`) VALUES (3, 'Six month');
INSERT INTO `ServiceType` (`type`, `serviceTypeName`) VALUES (4, 'One year');
INSERT INTO `ServiceServiceTypePairs` (`serviceType`, `serviceName`, `price`) VALUES (1, 'EC2', 1000);
INSERT INTO `ServiceServiceTypePairs` (`serviceType`, `serviceName`, `price`) VALUES (2, 'EC2', 100);
INSERT INTO `ServiceServiceTypePairs` (`serviceType`, `serviceName`, `price`) VALUES (3, 'EC2', 400);
INSERT INTO `ServiceServiceTypePairs` (`serviceType`, `serviceName`, `price`) VALUES (4, 'EC2', 700);
INSERT INTO `ServiceServiceTypePairs` (`serviceType`, `serviceName`, `price`) VALUES (3, 'Compute', 2500);
INSERT INTO `ServiceServiceTypePairs` (`serviceType`, `serviceName`, `price`) VALUES (4, 'Compute', 4500);
INSERT INTO `ServiceServiceTypePairs` (`serviceType`, `serviceName`, `price`) VALUES (1, 'S3', 50000);
INSERT INTO `ServiceServiceTypePairs` (`serviceType`, `serviceName`, `price`) VALUES (2, 'S3', 12000);
INSERT INTO `ServiceServiceTypePairs` (`serviceType`, `serviceName`, `price`) VALUES (3, 'S3', 15000);
INSERT INTO `ServiceServiceTypePairs` (`serviceType`, `serviceName`, `price`) VALUES (4, 'S3', 19999);

INSERT INTO `ServiceSubscriptionTransaction` (`type`, `serviceName`, `organizationName`, `description`, `activeUntil`, `amountPaid`, `processedTimestamp`) VALUES (2, 'EC2', 'MacroHard', 'hello', '2008-01-20 00:00:00', 100, '2007-12-20 00:00:00');