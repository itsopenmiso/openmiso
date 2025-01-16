---
layout: plugins
page_title: 'Plugin: AWS EC2'
description: 'Build, Deploy, and Release on AWS EC2'
---

# AWS EC2

## aws-ami (builder)

Search for and return an existing AMI.

### Interface

- Input: **component.Source**
- Output: **ami.Image**

### Required Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### region

The AWS region to search in.

- Type: **string**

### Optional Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### filters

DescribeImage specific filters to search with.

The filters are always name => [value].

- Type: **map of string to list of string**
- **Optional**

#### name

The name of the AMI to search for, supports wildcards.

- Type: **string**
- **Optional**

#### owners

The set of AMI owners to restrict the search to.

- Type: **list of string**
- **Optional**

### Output Attributes

Output attributes can be used in your `waypoint.hcl` as [variables](../docs/waypoint-hcl/variables) via [`artifact`](../docs/waypoint-hcl/variables/artifact) or [`deploy`](../docs/waypoint-hcl/variables/deploy).

#### image

- Type: **string**

## aws-ec2 (platform)

Deploy the application into an AutoScaling Group on EC2.

### Interface

- Input: **ami.Image**
- Output: **ec2.Deployment**

### Required Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### instance_type

The EC2 instance type to deploy.

- Type: **string**

#### region

The AWS region to deploy into.

- Type: **string**

#### service_port

The TCP port on the instances that the app will be running on.

- Type: **int**

### Optional Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### count

How many EC2 instances to configure the ASG with.

The fields here (desired, min, max) map directly to the typical ASG configuration.

- Type: **ec2.countConfig**
- **Optional**

#### extra_ports

Additional TCP ports to allow into the EC2 instances.

These additional ports are usually used to allow secondary services, such as ssh.

- Type: **list of int**
- **Optional**

#### key

The name of an SSH Key to associate with the instances, as preconfigured in EC2.

- Type: **string**
- **Optional**

#### security_groups

Additional security groups to attached to the EC2 instances.

This plugin creates security groups that match the above ports by default. this field allows additional security groups to be specified for the instances.

- Type: **list of string**
- **Optional**

#### subnet

The subnet to place the instances into.

- Type: **string**
- **Optional**
- Default: a public subnet in the dafault VPC

### Output Attributes

Output attributes can be used in your `waypoint.hcl` as [variables](../docs/waypoint-hcl/variables) via [`artifact`](../docs/waypoint-hcl/variables/artifact) or [`deploy`](../docs/waypoint-hcl/variables/deploy).

#### public_dns

- Type: **string**

#### public_ip

- Type: **string**

#### region

- Type: **string**

#### service_name

- Type: **string**

#### target_group_arn

- Type: **string**

## aws-alb (releasemanager)

Release target groups by attaching them to an ALB.

### Interface

- Input: **alb.TargetGroup**
- Output: **alb.Release**

### Mappers

#### Allow EC2 Deployments to be hooked up to an ALB

- Input: **ec2.Deployment**
- Output: **alb.TargetGroup**
#### Allow Lambda Deployments to be hooked up to an ALB

- Input: **lambda.Deployment**
- Output: **alb.TargetGroup**

### Required Parameters

This plugin has no required parameters.

### Optional Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### certificate

ARN for the certificate to install on the ALB listener.

When this is set, the port automatically changes to 443 unless overriden in this configuration.

- Type: **string**
- **Optional**

#### domain_name

Fully qualified domain name to set for the ALB.

Set along with zone_id to have DNS automatically setup for the ALB. this value should include the full hostname and domain name, for instance app.example.com.

- Type: **string**
- **Optional**

#### listener_arn

The ARN on an existing ALB to configure.

When this is set, no ALB or Listener is created. Instead the application is configured by manipulating this existing Listener. This allows users to configure their ALB outside waypoint but still have waypoint hook the application to that ALB.

- Type: **string**
- **Optional**

#### name

The name to assign the ALB.

Names have to be unique per region.

- Type: **string**
- **Optional**
- Default: derived from application name

#### port

The TCP port to configure the ALB to listen on.

- Type: **int**
- **Optional**
- Default: 80 for HTTP, 443 for HTTPS

#### security_group_ids

The existing security groups to add to the ALB.

A set of existing security groups to add to the ALB.

- Type: **list of string**
- **Optional**

#### subnets

The subnet ids to allow the ALB to run in.

- Type: **list of string**
- **Optional**
- Default: public subnets in the account default VPC

#### zone_id

Route53 ZoneID to create a DNS record into.

Set along with domain_name to have DNS automatically setup for the ALB.

- Type: **string**
- **Optional**
