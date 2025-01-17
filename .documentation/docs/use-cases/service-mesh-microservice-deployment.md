---
layout: docs
page_title: Using Waypoint to Deploy Microservices to a Service Mesh
description: |-
  HashiCorp Waypoint eases the deployment of microservices to a Consul service mesh
---

# Using Waypoint to Deploy Microservices to a Service Mesh

HashiCorp Waypoint eases the deployment of microservices to a service mesh. Some
of the well-known benefits of a microservices architecture include scalability
and maintainability. However, as some engineers who build, deploy and release
countless microservices can attest to, orchestrating these processes can become
a toilsome task. The task becomes greater when there are multiple platforms to
consider for deploying each of the microservices. This use case shows how Waypoint
can help with that.

A [companion GitHub repo](https://github.com/paladin-devops/waypoint-service-mesh-use-case)
to this use case write-up contains an example deployment of microservices to a
service mesh using HashiCorp Waypoint. HashiCorp Consul is the service mesh
platform, and in this example, there are two Consul clusters peered together.
One cluster, named `nuka-cola`, runs on AWS EC2s, integrated with a HashiCorp
Nomad cluster. The other, named `sunset-sarsaparilla`, runs on Kubernetes,
specifically in an Azure Kubernetes Services (AKS) cluster.

There are two services in this deployment: 1) `counter-api` is deployed to the
`nuka-cola` datacenter and 2) `counter-ui` is deployed to the `sunset-sarsaparilla`
datacenter. Each of these is built and deployed using HashiCorp Waypoint.

This use case does not cover the setup of the AKS cluster, the AWS EC2s, or
the Consul clusters. However, these configurations are kept in the repo to
serve as an example of how such resources may be configured.

This use case does cover the Consul peering configuration of a specific service,
how to configure to apps to communicate between peered clusters, and deployment
of the apps to the two Consul clusters.

## Pre-requisites

1. Kubernetes cluster
   - Local context connected to this cluster
2. Consul Helm chart deployed to Kubernetes cluster
   - Mesh gateway enabled
   - Additional Kubernetes resources applied to cluster (found in the use case repo under `sunset-sarsaparilla/kubernetes-config/`)
     - Service for `counter-ui`
     - Service account for `counter-ui`
     - Consul service default config entry for `counter-ui`
     - Consul mesh config entry
     - Consul mesh gateway proxy default
3. Nomad cluster, integrated with Consul
   - Mesh gateway deployed
   - Additional Consul resources created (found in the use case repo under `nuka-cola/`)
     - Consul service default config entry for `counter-api`
     - Consul mesh config entry
     - Consul mesh gateway proxy default
4. An HCP Waypoint server
   - Local context connected to HCP Waypoint
   - An OSS Waypoint server will also work, as long as the runners in Nomad and Kubernetes can connect to it.
5. A Docker image registry
   - Credentials to push to the registry

## Deploy the Waypoint Runners to Nomad and Kubernetes

Prior to deploying the applications to each cluster, Waypoint runners must be
deployed to each, in order to use remote operations.

```shell
$ waypoint runner install \
  -platform=kubernetes \
  -id=mr-new-vegas \
  -server-addr=api.hashicorp.cloud:443 \
  -- -label=consul-dc=sunset-sarsaparilla

$ waypoint runner install \
  -platform=nomad \
  -id=mr-house \
  -nomad-host-volume=waypoint-runner \
  -server-addr=api.hashicorp.cloud:443 \
  -nomad-dc=nuka-cola
  -- -label=consul-dc=nuka-cola
```

After the installation to of a runner to Nomad and Kubernetes is complete, `waypoint runner list`
will show that there are two runners reporting to the server, waiting for jobs to be queued:

```shell
$ waypoint runner list
              ID             |   STATE    |   KIND    |             LABELS             | LAST REGISTERED
-----------------------------+------------+-----------+--------------------------------+------------------
  mr-house                   | preadopted | remote    | consul-dc:nuka-cola            |
  mr-new-vegas               | preadopted | remote    | consul-dc:sunset-sarsaparilla  |
```

There will also be two runner profiles configured. When the time comes to build,
deploy and release the apps, the "target runner" indicated for each will be used
to filter the runners on which a Waypoint job will be executed. To read more about
remote operations and how they work, check out [the documentation](../docs/projects/remote).

```shell
$ waypoint runner profile list
Runner profiles
           NAME           | PLUGIN TYPE |            OCI URL            |            TARGET RUNNER            | DEFAULT
--------------------------+-------------+-------------------------------+-------------------------------------+----------
  kubernetes-mr-new-vegas | kubernetes  | hashicorp/waypoint-odr:latest | labels:                             |
                          |             |                               | {"consul-dc":"sunset-sarsaparilla"} |
  nomad-mr-house          | nomad       | hashicorp/waypoint-odr:latest | labels:                             |
                          |             |                               | {"consul-dc":"nuka-cola"}           |
```

## Build & Deploy counter-api to Nomad

Because `counter-ui` depends on `counter-api`, the latter will be deployed first.
The Waypoint configuration for `counter-api` is below:

```hcl
app "counter-api" {
  build {
    use "docker" {
      dockerfile = "${path.app}/counter-api/Dockerfile"
      context    = "${path.app}/counter-api/"
    }

    registry {
      use "docker" {
        image = var.api_image_name
        tag   = var.api_tag_name

        auth {
          username = var.username
          password = var.password
        }
      }
    }
  }

  deploy {
    use "nomad-jobspec" {
      jobspec = templatefile("${path.app}/nuka-cola/counter-api.nomad.hcl", {
        count = var.api_replicas
      })
    }
  }
}
```

Waypoint will use the Docker plugin to build the `counter-api` as per the
Dockerfile written for the app. Following that, the Nomad jobspec file,
`counter-api.nomad.hcl`, will be used to deploy the Nomad job. Before starting
the build and deployment, the runner profile must be configured to use the
Nomad profile (the `runner` stanza below should be configured in the Waypoint
file, and committed & pushed to the repo):

```hcl
runner {
  enabled = true
  profile = "nomad-mr-house"
}
```

Following this, running `waypoint up` will build the app, push the app image
to a registry, and then deploy it to Nomad.

```shell
# waypoint.wpvars contains various input vars set that are required in the
# Waypoint file
$ waypoint up -p countdash -a counter-api -var-file=waypoint.wpvars
```

When the deployment is complete, there should be a Nomad job up and running,
with a Consul Connect sidecar proxy running Envoy alongside it:

![](/img/use-cases/service-mesh-microservice-deployment/counter-api-nomad-status.png)

## Build & Deploy counter-ui to Kubernetes

With `counter-api` deployed, `counter-ui` can be deployed next. Its Waypoint
configuration is below:

```hcl
app "counter-ui" {
  build {
    use "docker-pull" {
      image = "hashicorpnomad/counter-dashboard"
      tag   = "v1"
    }

    registry {
      use "docker" {
        image = var.ui_image_name
        tag   = var.ui_tag_name

        auth {
          username = var.username
          password = var.password
        }
      }
    }
  }

  deploy {
    use "kubernetes" {
      replicas        = var.ui_replicas
      service_port    = 9002
      service_account = "${app.name}"
      labels = {
        "waypoint-app" = "${app.name}"
      }
      pod {
        container {
          name = "${app.name}"
          static_environment = {
            "COUNTING_SERVICE_URL" = "http://counter-api.virtual.nuka-cola.consul"
            "PORT" = "9002"
          }
        }
      }
      probe_path = "/health"
    }
  }
}
```

Waypoint will use the Docker Pull plugin to pull the UI image from DockerHub,
and push it to a new registry after adding the [Waypoint Entrypoint](../docs/entrypoint).
Following that, the image will be deployed to Kubernetes using the Kubernetes
plugin. `counter-ui` uses the environment variable `COUNTING_SERVICE_URL` to
determine where to reach the API service. Soon, we will see how Consul cluster
peering enables connectivity to the API service at `counter-api.virtual.nuka-cola.consul`.

Some additional configurations of note include the `service_account` and `labels`
for the Kubernetes deployment. Consul on Kubernetes by default expects a service
account to be created matching the name of the service to be registered to Consul.
In this example, the name of the service registered to Consul, as configured in
`counter-ui-service.yaml`, is `counter-ui`.

The `counter-ui` service registered to Consul, as per the pre-requisites, contains
a label selector, to ensure that the pods with the `waypoint-app` label matching
`counter-ui` receive traffic for the service. This is why the `labels` config
is set for the Kubernetes plugin.

Finally, the `probe_path` config is set, so that an HTTP health check is registered
to Kubernetes, and also registered to Consul by the service mesh's transparent proxy.
Consul will therefore automatically expose the health check endpoint to be accessible
by Kubernetes from outside the service mesh. There is more information about
this in the [Consul documentation](/consul/docs/k8s/connect/health).

With these configurations set, the runner profile must be updated once more so
that Waypoint launches a runner in the Kubernetes cluster when building and
deploying `counter-ui`.

```hcl
runner {
  enabled = true
  profile = "kubernetes-mr-new-vegas"
}
```

Now, after changing only the application flag (`-a`) to `counter-ui`, the build
and deployment may commence with the `waypoint up` command below:

```shell
# waypoint.wpvars contains various input vars set that are required in the
# Waypoint file
$ waypoint up -p countdash -a counter-ui -var-file=waypoint.wpvars
```

After the build and deployment are complete, the deployment should be ready in
Kubernetes:

```shell
$ kubectl get deployment
NAME             READY   UP-TO-DATE   AVAILABLE   AGE
counter-ui-v1   1/1     1            1           5m
```

## Service Mesh Connectivity between Waypoint Apps

It's possible at this point to set up an external IP address on the Kubernetes
service, to enable you to access the UI from outside the cluster - but for now,
we'll access it locally using port forwarding.

```shell
$ kubectl port-forward counter-ui-v1-869b8cd67c-v9b5b 9002:9002
```

If you browse to http://localhost:9002, you'll see that the counter dashboard
is visible, but not connected to the API service:

![](/img/use-cases/service-mesh-microservice-deployment/counter-ui-inaccessible.png)

To fix this, a couple of additional things need to be done. Firstly, there needs to be
a Consul peering configuration applied. The configuration below exports the `counter-api`
service from the `nuka-cola` datacenter to the `sunset-sarsaparilla` datacenter:

```hcl
Kind = "exported-services"
Name = "default"
Services = [
  {
    Name      = "counter-api"
    Consumers = [
      {
        Peer = "sunset-sarsaparilla"
      }
    ]
  }
]
```

A Consul intention must also be added, to enable the UI service to access the
API service:

```hcl
Kind      = "service-intentions"
Name      = "counter-api"

Sources = [
  {
    Name   = "counter-ui"
    Peer   = "sunset-sarsaparilla"
    Action = "allow"
  }
]
```

Both of the configurations above can be set using the `consul config write` command.

With this completed, the counter application should be working as expected:

![](/img/use-cases/service-mesh-microservice-deployment/countdash-ui-accessible.png)
