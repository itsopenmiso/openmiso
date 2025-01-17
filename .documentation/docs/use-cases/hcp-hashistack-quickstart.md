---
layout: docs
page_title: A Complete Cloud Ecosystem built with Terraform, HCP Waypoint, HCP Vault, HCP Consul, and Nomad
description: |-
  Your enterprise application is only a 'terraform apply' and 'waypoint up' away from life on the cloud.
---

# A Complete Cloud Ecosystem built with Terraform, HCP Waypoint, HCP Vault, HCP Consul, and Nomad

In this use case, we will spin up our entire infrastructure with `terraform apply`,
then launch the application to an EC2 container within our secured internal network with `waypoint up`.
Welcome to the HashiCorp stack magic.

The ingredients for this special dish are [Waypoint](https://www.waypointproject.io/),
[Terraform](https://www.terraform.io/), [Consul](https://www.consul.io/), [Nomad](https://www.nomadproject.io/), and
[Vault](https://www.vaultproject.io/).
We chose to use HashiCorp managed versions of Waypoint, Consul, and Vault for one convenient HCP UI and
to showcase the [HashiCorp Virtual Network (HVN)](/hcp/docs/hcp/network).
A similar experience is of course possible with self-managed versions,
and equally possible with Terraform automation, although it may require more familiarity with Terraform modules.

## Prerequisites

To follow along with this use case on your own machine, you will need:

- A [HashiCorp Cloud Platform (HCP)](https://portal.cloud.hashicorp.com/sign-in) account.
- An [AWS account](https://aws.amazon.com/account/) with AWS Access Credentials configured locally:

```shell
export AWS_ACCESS_KEY_ID=<your AWS access key ID>
export AWS_SECRET_ACCESS_KEY=<your AWS secret access key>
export AWS_SESSION_TOKEN=<your AWS session token>
```

- Docker registry
- Credentials to push to the Docker registry

As well as the following binaries installed locally:

- [Waypoint CLI](../downloads)
- [Terraform CLI](/terraform/tutorials/aws-get-started/install-cli)

It will help to have some familiarity with HashiCorp Configuration Language (HCL), which is the syntax used to
write configuration files for [Terraform](/terraform/language) and [Waypoint](../docs/waypoint-hcl/syntax).

In this use case, we will configure our infrastructure as code, which Terraform will run to provision the following:

- [AWS Virtual Private Cloud](https://aws.amazon.com/pt/vpc/) and peer it with our [HashiCorp Virtual Network (HVN)](/hcp/docs/hcp/network)
- [HCP Consul](https://cloud.hashicorp.com/products/consul) service mesh to manage services within the secure private network
- [HCP Vault](https://cloud.hashicorp.com/products/vault) cluster to securely store credentials
- [HashiCorp Nomad](/nomad/tutorials/get-started/get-started-intro) cluster on AWS EC2 instance to manage application containers

We will write application deployment configurations as code, which [HCP Waypoint](https://cloud.hashicorp.com/products/waypoint)
will run to build the container image and deploy the application to the infrastructure we set up with Terraform.

All it takes are two commands:

- `terraform apply`
- `waypoint up`

The Terraform and Waypoint configuration files and a sample application live in [this repository](https://github.com/xiaolin-ninja/2048_hcp_hashistack).
We will deploy a custom HashiCorp version of the popular game "2048".
Please clone the repo if you'd like to follow along the tutorial on your local machine.

## Create service principal and key

To leverage the Terraform integration and be able to deploy your HCP Consul using
Terraform, you have to create a Service Principal and a key associated to it.

From the left menu select _Access Control (IAM)_.
On the left menu of the _Access Control (IAM)_ page, click on the _Service Principals_ tab, and on
the page click on the create link.

Use the name you prefer to create the Service Principal (we used `xx-2048` for
this use case) but remember to use `Contributor` as the role.

Once the Service Principal is created, click on the service principal name on
the page to view its details. From the detail page, click on **+ Generate key**.

~> **Note:** Remember to copy the Client ID and secret; you won't be able to
retrieve the secret later once you close the pop up.

Save the client ID and secret as the environment variables `HCP_CLIENT_ID` and `HCP_CLIENT_SECRET`.
Or, configure the provider with the client ID and secret by copy-pasting the values directly into provider config.

```terraform
// Credentials can be set explicitly or via the environment variables HCP_CLIENT_ID and HCP_CLIENT_SECRET
provider "hcp" {
  client_id     = "service-principal-key-client-id"
  client_secret = "service-principal-key-client-secret"
}
```

```shell
export HCP_CLIENT_ID="service-principal-key-client-id"
export HCP_CLIENT_SECRET="service-principal-key-client-secret"
```

You can view existing keys and generate new ones on this page:
![HCP Service Principal](/img/use-cases/hcp-hashistack-quickstart/service-principal.png)

When client credentials are set, they are always used by the HCP Provider client, regardless of an existing user session.

## Define the Virtual Network

### VPC

You can find VPC configurations in [`vpc.tf`](https://github.com/xiaolin-ninja/2048_hcp_hashistack/blob/2da3a78f24e95901e49926ec224edf2237057d24/vpc.tf). There is more than one way to create this VPC, including manually in AWS.
The Amazon VPC can be managed with the [AWS provider](https://registry.terraform.io/providers/hashicorp/aws/latest/docs).

```terraform
provider "aws" {
  region = var.region
}

module "vpc" {
  source         = "terraform-aws-modules/vpc/aws"
  name           = var.name
  cidr           = var.cidr
  public_subnets = var.public_subnets
  public_subnet_tags = {
    Name = "${var.name}-public"
  }

  azs                             = slice(data.aws_availability_zones.available.*.names[0], 0, 3)
  enable_dns_hostnames            = true
  enable_ipv6                     = true
  assign_ipv6_address_on_creation = true
  public_subnet_ipv6_prefixes     = [0, 1, 2]

  tags = {
    Terraform = "true"
    Name      = var.name
  }
}
```

### HVN

In the [`hcp.tf`](https://github.com/xiaolin-ninja/2048_hcp_hashistack/blob/2da3a78f24e95901e49926ec224edf2237057d24/hcp.tf) configuration file,
we define a HashiCorp Virtual Network (HVN) named "xx-2048" and peer it to the VPC.
HVNs enable you to deploy HashiCorp Cloud products without having to manage the networking details.

Creating peering connections from your HVN allows you to connect and launch AWS resources from your HCP account.
Each peering connection requires an HVN route to set its CIDR block.

Once the VPC is peered with the HVN, the resources deployed inside the VPC will be able to reach each other.

```terraform
provider "hcp" {}

resource "hcp_hvn" "xx_2048_hvn" {
  hvn_id         = "xx-2048"
  cloud_provider = "aws"
  region         = var.region
}

resource "hcp_aws_network_peering" "peer" {
  hvn_id          = hcp_hvn.xx_2048_hvn.hvn_id
  peering_id      = "hcp-hashistack"
  peer_vpc_id     = module.vpc.vpc_id
  peer_account_id = module.vpc.vpc_owner_id
  peer_vpc_region = var.region
}

resource "hcp_hvn_route" "hvn-to-vpc" {
  hvn_link         = hcp_hvn.xx_2048_hvn.self_link
  hvn_route_id     = "hvn-to-vpc"
  destination_cidr = module.vpc.vpc_cidr_block
  target_link      = hcp_aws_network_peering.peer.self_link
}

resource "aws_vpc_peering_connection_accepter" "peer" {
  vpc_peering_connection_id = hcp_aws_network_peering.peer.provider_peering_id
  auto_accept               = true
  accepter {
    allow_remote_vpc_dns_resolution = true
  }
}

resource "aws_route" "hvn-peering" {
  route_table_id            = module.vpc.public_route_table_ids[0]
  destination_cidr_block    = hcp_hvn.xx_2048_hvn.cidr_block
  vpc_peering_connection_id = aws_vpc_peering_connection_accepter.peer.id
}
```

## Define HCP Consul

HCP Consul enables you to quickly deploy Consul servers in AWS while offloading the operations burden to the SRE experts at HashiCorp.

Provisioning HCP Consul with Terraform is incredibly easy.
You can define the cluster in Terraform using the [`hcp_consul_cluster`](https://registry.terraform.io/providers/hashicorp/hcp/latest/docs/resources/consul_cluster) resource,
as shown in `hcp.tf`:

```terraform
resource "hcp_consul_cluster" "xx_2048_consul" {
  hvn_id          = hcp_hvn.xx_2048_hvn.hvn_id
  cluster_id      = "xx-2048-consul"
  tier            = "development"
  public_endpoint = true
}
```

You can also [provision a Consul cluster via the HCP UI](/consul/tutorials/get-started-hcp/hcp-gs-deploy).

## Define HCP Vault

Defining an HCP Vault cluster via Terraform also requires minimal configuration,
using the [`hcp_vault_cluster`](https://registry.terraform.io/providers/hashicorp/hcp/latest/docs/resources/vault_cluster) resource.

~> **Note:** We enabled the public endpoint for our Vault cluster for ease of demo,
but deploying Vault in a publicly accessible way should be avoided if possible to reduce security exposure.

```terraform
resource "hcp_vault_cluster" "xx_2048_vault" {
  hvn_id          = hcp_hvn.xx_2048_hvn.hvn_id
  cluster_id      = "xx-2048-vault"
  public_endpoint = true
}
```

### Vault Cluster Configurations

We will further configure Vault to communicate with Nomad.
The [Vault provider](https://registry.terraform.io/providers/hashicorp/vault/latest/docs) allows Terraform to read from, write to, and configure HashiCorp Vault:

```terraform
provider "vault" {
  address   = hcp_vault_cluster.xx_2048_vault.vault_public_endpoint_url
  namespace = "admin"
  token     = hcp_vault_cluster_admin_token.xx_2048_vault_token.token
}
```

You can find the detailed Vault configurations in [`vault-config.tf`](https://github.com/xiaolin-ninja/2048_hcp_hashistack/blob/2da3a78f24e95901e49926ec224edf2237057d24/vault-config.tf)

## Define Nomad Configurations

In [`nomad.tf`](https://github.com/xiaolin-ninja/2048_hcp_hashistack/blob/2da3a78f24e95901e49926ec224edf2237057d24/nomad.tf),
we configure the AWS Security Group, find a demo AMI, and configure the Nomad server, client, and auto-scaling.
The scripts `user-data-client.sh` and `user-data-server.sh` are run directly in the EC2 instances spun up by Terraform to provision Nomad.

~> **Note:** Since this is a use-case demo, we are not following Security best-practices.
For example, my AWS Security Group allows all ingress and egress traffic.
We also provision only one Nomad server and client cluster each, which does not insulate against downtime.
For real production use cases, please properly configure your security group and have multiple clusters.

The launch template looks like this:

```terraform
resource "aws_launch_template" "nomad-clients" {
  name = "XX_2048_Nomad_Clients"

  block_device_mappings {
    device_name = "/dev/sda1"

    ebs {
      volume_size = 16
    }
  }

  image_id = data.aws_ami.base.image_id

  instance_initiated_shutdown_behavior = "terminate"

  instance_type = "t3.medium"

  metadata_options {
    http_endpoint               = "enabled"
    http_tokens                 = "required"
    http_put_response_hop_limit = 1
    instance_metadata_tags      = "enabled"
  }

  monitoring {
    enabled = true
  }

  vpc_security_group_ids = [aws_security_group.xx_2048.id]

  tag_specifications {
    resource_type = "instance"

    tags = {
      Name = "Nomad-Client"
    }
  }

  user_data = base64encode(templatefile("user-data-client.sh", {
    nomad_region                = var.region,
    nomad_datacenter            = var.cluster_name,
    consul_ca_file              = base64decode(hcp_consul_cluster.xx_2048_consul.consul_ca_file),
    consul_gossip_encrypt_key   = jsondecode(base64decode(hcp_consul_cluster.xx_2048_consul.consul_config_file)).encrypt,
    consul_acl_token            = hcp_consul_cluster.xx_2048_consul.consul_root_token_secret_id,
    consul_private_endpoint_url = hcp_consul_cluster.xx_2048_consul.consul_private_endpoint_url,
    vault_endpoint              = hcp_vault_cluster.xx_2048_vault.vault_private_endpoint_url,
  }))
}
```

This is not the only way to define Nomad clusters in Terraform.
For more practice deploying Nomad to EC2 with Terraform, [follow this tutorial](/nomad/tutorials/cluster-setup/cluster-setup-aws).

## Deploy the Infrastructure!

Once you have finished defining infrastructure configurations
(or cloned them from the [repo](https://github.com/xiaolin-ninja/2048_hcp_hashistack)),
and set the environment variables for HCP and AWS access, you are now ready to deploy your infrastructure!

```shell
$ terraform init
Initializing the backend...
Initializing provider plugins...
...
Terraform has been successfully initialized!
...
```

Once Terraform has been initialized, you can verify that the resources will
be created using the `terraform plan` command.

```shell
$ terraform plan
An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create
 <= read (data resources)
Terraform will perform the following actions:
...
...
Plan: 11 to add, 0 to change, 0 to destroy.
```

Finally, this is the magic moment we promised:

```shell
$ terraform apply
...
Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.
  Enter a value:
```

Remember to confirm the run by entering `yes`!

## Set up HCP Waypoint Context

HCP Waypoint is currently in Beta, and a Terraform module to provision Waypoint is still under development.

Head into your HCP UI and [follow this tutorial](../tutorials/runners/hcp-runners-on-your-infra#set-up-the-environment)
to set up your local environment.

Navigate to Waypoint in the HCP dashboard:
![HCP UI](/img/use-cases/hcp-hashistack-quickstart/hcp.png)

Then click on "Install New Runner", which will generate terminal inputs to set up the Waypoint context and install a Runner on your chosen platform.
In this case, we choose Nomad with a host volume named "wp_runner".
![Waypoint Install New Runner](/img/use-cases/hcp-hashistack-quickstart/waypoint-runner-install.png)

You will need to export the Nomad server address, which you can find on the AWS EC2 instance:
![Nomad EC2 Instances](/img/use-cases/hcp-hashistack-quickstart/nomad-ec2.png)

```shell
export NOMAD_ADDR=<nomad_private_addr>
```

![Waypoint Install New Runner CLI](/img/use-cases/hcp-hashistack-quickstart/wp-runner-install.png)

Export the nomad address to your local environment and install the runner after installing the HCP Waypoint context,
and you're all set to use HCP Waypoint.

The `waypoint.hcl` file defines application deployment configurations.
For this use case, we have a simple "build", "deploy", and "release" flow:

```hcl
project = "2048-hashistack"

app "nomad" {
  build {
    use "docker" {}
    registry {
      use "docker" {
        ...
      }
    }
  }

  deploy {
    use "nomad" {
      datacenter = "hcp-2048"
    }
  }
}
```

You can [read more about Waypoint HCL](../docs/waypoint-hcl) fundamentals,
or explore [advanced deployment workflows using custom pipelines](../docs/pipelines).

Now, run `waypoint init`:

```shell
$ waypoint init
✓ Configuration file appears valid
✓ Connection to Waypoint server was successful
✓ Project "2048-hashistack" and all apps are registered with the server.
```

And the moment of magic...

`waypoint up`

```shell
$ waypoint up

» Operation is queued waiting for job "01GRM4JZAFWRC4RA8M8BBR4NV8". Waiting for runner assignment...
  If you interrupt this command, the job will still run in the background.
  Performing operation on "nomad" with runner profile "nomad-01GQTC242045AME1DMSK1KPCDS"

...
...

The deploy was successful!

The release did not provide a URL and the URL service is disabled on the
server, so no further URL information can be automatically provided. If
this is unexpected, please ensure the Waypoint server has both the URL service
enabled and advertise addresses set.
```

Congratulations, your application is live!

HCP Waypoint does not provide a URL service.
You can also deploy the application by [installing self-hosted Waypoint](../commands/install)
and running `waypoint up` with a different context.

## Easy to Monitor Dashboards

After our `terraform apply` and `waypoint up`, you should have a full HCP HashiCorp stack set up automatically and application deployed.

### HashiCorp Virtual Network

You can find "HashiCorp Virtual Networks" on the left sidebar and see that our HCP Consul and HCP Vault resources have
been automatically provisioned via Terraform in the same VPC, and associated with the network. Magic!

![HVN UI](/img/use-cases/hcp-hashistack-quickstart/hvn.png)

### Consul Dashboard

You can access the HCP Consul UI by navigating to the Consul tab in HCP, and clicking "Access Consul".
For this use case, we configured a public IP address for the cluster, so you can access Consul with the public address.
This is not recommended for a production use case.

![Consul UI](/img/use-cases/hcp-hashistack-quickstart/consul.png)

As you can see, Consul service mesh keeps track of services in our VPC.

### Nomad Dashboard

Nomad is a workload orchestration program, so services in the HVN are reflected as jobs.
You can see the Nomad server, client, and Waypoint runner.

![Nomad UI](/img/use-cases/hcp-hashistack-quickstart/nomad-ui.png)

Have fun exploring! You can find more guidance in the [Nomad UI documentation](/nomad/tutorials/web-ui).

### HashiCorp stack 2048

You can find a persistent version of the application we deployed here: [HashiCorp Stack 2048 Game](https://tunzor.github.io/hashi-2048/)

![2048 Game](/img/use-cases/hcp-hashistack-quickstart/hcp-2048.png)

## Summary

Waypoint and Terraform
Cloud automation magic,
Easy deployments!

Our Engineers aren't paid to write haikus, but infrastructure automation inspires innovation.

HashiCorp customers already know how easy it is to provision and manage cloud infrastructure with HashiCorp tools,
and Waypoint expands the HashiCorp stack by enabling deployment automation with custom pipelines.
Hopefully, this use case showcases the synergy between HashiCorp products, and how easy they are to configure.
Our mission is to help increase productivity, reduce risk and capital expenditure,
and drive innovation within your organization.
Thanks for taking the time to follow along this use case, and we are excited to see how you innovate with our tools!

## Dive Deeper into the HashiCorp stack

We covered many HashiCorp tools in this use case. We focused on HCP product offerings, where HashiCorp manages the tooling for you.
Provisioning is easy with the UI or Terraform, as shown in this demo.
Of course, all our tools have Open Source versions and are available for self-management as well.
We include below a few tutorials that dive deeper into pieces of the workflow shown in this use case,
so you can explore this sweet HashiCorp stack synergy with self-managed Waypoint, Consul, Nomad, and Vault as well!

- [HashiCorp Cloud Platform](https://cloud.hashicorp.com/)
- [Consul Service Mesh](/consul/docs/connect)
- [Deploy an Application to Nomad with Waypoint](../tutorials/get-started-nomad/get-started-nomad)
- [Set Up a Nomad Cluster on AWS](/nomad/tutorials/cluster-setup/cluster-setup-aws)
- [Blue-Green Deployments with Waypoint, Nomad, and Consul](../docs/use-cases/blue-green-deployment)
- [Consul Deployment Guide](/consul/tutorials/production-deploy/deployment-guide)
- [Advanced Deployment Pipelines with Waypoint](../docs/pipelines)
