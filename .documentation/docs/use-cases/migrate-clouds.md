---
layout: docs
page_title: Seamless Cloud Migration With Waypoint
description: |-
  Waypoint enables easy migration of applications between different cloud providers.
---

# Seamless Cloud Migration With Waypoint

A time may come where you need to move your application from one cloud provider to
another, or one data center to another, but doing so can be cumbersome. There may be
no standards between application tools for building, you may have special scripts for
deploying, and releasing may prove to be difficult since you can't release across
platforms in the same manner. Waypoint can solve those problems. With Waypoint, it is
as easy as deploying a runner to your new cluster, changing kubectl context, and
running `waypoint up` to migrate your application across multiple cloud providers.

This use case starts with an application deployed in AWS EKS and showcases migrating
that same application to GCP GKE, and then Azure AKS.

## Pre-Requisites

If experimenting with this use case, it will be helpful to have some familiarity
with the tools and concepts listed below.

- [Waypoint CLI](../downloads)
  - [Context created](../docs/server/auth#logging-in) and connected to local Waypoint server or HCP Waypoint server
- Waypoint [runner](../docs/runner/run-manual#installing-a-runner)
  - Must be able to connect to your Kubernetes clusters
- Waypoint [runner configuration](../docs/runner/config)
  - Needs credentials to your registry
- Docker registry
  - Credentials to push a build to the registry
- [AWS EKS](https://docs.aws.amazon.com/eks/latest/userguide/clusters.html) cluster
- [GCP GKE](https://cloud.google.com/kubernetes-engine) cluster
- [Azure AKS](https://azure.microsoft.com/en-us/products/kubernetes-service/) cluster

## Migrate from AWS EKS to GCP GKE

##### Step 1. Start with an existing Kubernetes cluster and a deployed application

If you don't already have an application deployed by Waypoint, follow our [Kubernetes Getting Started Guide](../tutorials/get-started-kubernetes).

##### Step 2. Configure your kubectl context to point to your new cluster you are migrating to

```shell
# Use the context of your cluster you would like to migrate to
$ kubectl config use-context gcp-context
```

##### Step 3. Waypoint runner install into your new cluster

Install a runner into your cluster you are migrating to and configure the runner
profile with your registry credentials.

```shell
$ waypoint runner install -platform=kubernetes -server-addr=<server_addr>  -k8s-runner-image=hashicorp/waypoint:latest -id=gcp -- -label=cloud=gcp
✓ Finished connecting to: <server_addr>
⠧ Installing runner...
✓ Finished connecting to: <server_addr>
✓ Runner "gcp" installed successfully to kubernetes
✓ Runner profile "kubernetes-gcp" created successfully.
✓ Waypoint runner installed with Helm!
✓ Runner "gcp" adopted successfully.
```

After a successful `waypoint runner install` to your cluster you will see the
following resources. Now, the runner and any on-demand runners have cluster
edit permissions within your cluster. Note, the `waypoint runner install` has
installed a runner profile named `kuberenetes-gcp`, and did not change the
`default` runner profile if you already one.

```shell
$ kubectl get po
NAME                                READY   STATUS    RESTARTS   AGE
waypoint-gcp-runner-0               1/1     Running   0          3d1h

$ kubectl get serviceaccount
NAME                  SECRETS   AGE
default               1         29d
waypoint-runner       1         14d
waypoint-runner-odr   1         14d

$ kubectl get rolebinding
NAME                                       ROLE               AGE
waypoint-gcp-runner-odr-rolebinding-edit   ClusterRole/edit   14d
waypoint-gcp-runner-rolebinding            ClusterRole/edit   14d
```

##### Step 4. Write a waypoint.hcl with the following specifications, if you have not already

In the waypoint.hcl, specify the runner profile that was created with the
`waypoint runner install` command. The `runner` stanza configures a project
for remote operations and which runner profile to use.

```hcl
project = "hashitalk-deploy-gcp"

app "hello-app-gcp" {
  runner {
    profile = "kubernetes-gcp"
  }
...
```

The `build` stanza of the "app" is configured later in the Waypoint configuration.
This build uses the docker plugin to build an image and then pushes it to a Docker
registry.

```hcl
...
  build {
    use "docker" {}
    registry {
      use "docker" {
        image = var.image
        // Returns a humanized version of the git hash, taking into account tags and changes
        tag = gitrefpretty()
        // Credentials for authentication to push to docker registry
        auth {
          username = var.username
          password = var.password
        }
      }
    }
  }
...
```

The `deploy` stanza of the "app" describes how to deploy your application.
The deployment takes the previously built artifact and stages it onto the cluster.
Waypoint will deploy your application to the Kubernetes cluster that your
current context is referencing.

```hcl
...
  deploy {
    use "kubernetes" {
      service_port = 5300
      namespace = "default"
    }
  }
```

The `release` stanza of the "app" describes how to release your application.
The release activates the previously staged deployment and exposes it to traffic.

```hcl
...
  release {
    use "kubernetes" {
      port = 5300
    }
  }
...
```

The variables used in the example are defined below.

Note, that the variable `password` is populated from an environment variable
`DOCKER_PASSWORD` that should be defined locally. Specifically, the runners have
access to this environment variable from the `waypoint config set -runner` command.
In a production use case, a better approach would be using a dynamic config
variable sourced directly from Vault. To view various ways to authenticate for
pushing and pulling from private registries, follow our documentation on
[Scenarios for a Registry](../docs/lifecycle/build#scenarios-for-a-registry).

```hcl

variable "image" {
  type = string
  default = "hashicassie/hashitalk-deploy"
}
variable "username" {
  type = string
  default = "hashicassie"
}
variable "password" {
  type = string
  env = ["DOCKER_PASSWORD"]
  sensitive = true
}
```

##### Step 5. Build, Deploy, and Release Your App

Now the waypoint.hcl is complete. To get your application and project initialized
and built, deployed, and released to your Kubernetes cluster simply run the following:

```shell
$ waypoint init
$ waypoint project apply -data-source="git" -git-url="https://github.com/cicoyle/hashitalk-deploy" -git-ref=main -waypoint-hcl=waypoint.hcl hashitalk-deploy-gcp
$ waypoint config set -runner -scope=project DOCKER_PASSWORD=$DOCKER_PASSWORD
$ waypoint config get -runner
$ waypoint up
```

`waypoint init` is used to initialize and validate a project. `waypoint project apply`
is to update the project to include the remote `-git-url` and `-git-ref` such that it
will reference the latest code changes pushed to the main branch in the repo provided.
The `waypoint config set -runner` exposes the environment variables to the runner,
plus all the on demand runners. The scope can be `global`, `project`, or `app`. For this
use case, the configuration is project scoped meaning that this will apply only to
runners operating on jobs for the specific project. Use the `waypoint config get -runner`
to view and confirm the config is set for your runners. You will see that `waypoint up`
encapsulates the build, deploy, release logic.

See the application in your cluster.

```shell
$ kubectl get po
NAME                                READY   STATUS    RESTARTS   AGE
hello-app-gcp-v1-5bc8b5c944-z8922   1/1     Running   0          3d1h
waypoint-gcp-runner-0               1/1     Running   0          3d2h

$ kubectl get deploy
NAME               READY   UP-TO-DATE   AVAILABLE   AGE
hello-app-gcp-v1   1/1     1            1           13d

$ kubectl get rs
NAME                          DESIRED   CURRENT   READY   AGE
hello-app-gcp-v1-5bc8b5c944   1         1         1       13d
```

See the project and application displayed in the HCP UI below.

![Project and App in HCP UI](/img/use-cases/migrate-clouds/display-app-in-project.png)

See the deployment in the HCP UI with some operation logs below.

![Display Deployment in HCP UI](/img/use-cases/migrate-clouds/display-deployed-app.png)

## Migrate from GCP GKE to Azure AKS

Follow the same simple steps from above.

##### Step 1. Start with an existing Kubernetes cluster and a deployed application

If you don't already have an application deployed by Waypoint, follow our [Kubernetes Getting Started Guide](../tutorials/get-started-kubernetes).

##### Step 2. Configure your kubectl context to point to your new cluster you are migrating to

```shell
# Use the context of your cluster you would like to migrate to
$ kubectl config use-context azure-context
```

##### Step 3. Waypoint runner install into your new cluster

Install a runner into your cluster you are migrating to and configure the runner
profile with your registry credentials.

```shell
$ waypoint runner install -platform=kubernetes -server-addr=<server_addr>  -k8s-runner-image=hashicorp/waypoint:latest -id=azure -- -label=cloud=azure
✓ Finished connecting to: <server_addr>
⠧ Installing runner...
✓ Finished connecting to: <server_addr>
✓ Runner "azure" installed successfully to kubernetes
✓ Runner profile "kubernetes-azure" created successfully.
✓ Waypoint runner installed with Helm!
✓ Runner "azure" adopted successfully.
```

##### Step 4. Update the waypoint.hcl

In the waypoint.hcl, specify the runner profile that was created with the
`waypoint runner install` command. The `runner` stanza configures a project
for remote operations, and which runner profile to use.
Specifically, update the following:

1. Project and/or app name
2. Runner profile

```hcl
project = "hashitalk-deploy-azure"

app "hello-app-azure" {
  runner {
    profile = "kubernetes-azure"
  }
...
```

The rest of the waypoint.hcl will remain the same.

##### Step 5. Build, Deploy, and Release Your App

## Summary

At this point, hopefully you see how easy it is to migrate your existing
application to another cloud provider using Waypoint. Whether you are migrating
your existing application from AWS EKS to GCP GKE, or from GCP GKE to Azure AKS,
with a few minor tweaks to your waypoint.hcl, you can get your application migrated
in no time.
