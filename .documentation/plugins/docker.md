---
layout: plugins
page_title: 'Plugin: Docker'
description: 'Build and Deploy on Docker'
---

# Docker

-> For a step by step tutorial, visit [HashiCorp
Learn](../tutorials/get-started-docker).

## docker (builder)

Build a Docker image from a Dockerfile.

If a Docker server is available (either locally or via environment variables
such as "DOCKER_HOST"), then "docker build" will be used to build an image
from a Dockerfile.

### Dockerless Builds

Many hosted environments, such as Kubernetes clusters, don't provide access
to a Docker server. In these cases, it is desirable to perform what is called
a "dockerless" build: building a Docker image without access to a Docker
daemon. Waypoint supports dockerless builds.

Waypoint performs Dockerless builds by leveraging
[Kaniko](https://github.com/GoogleContainerTools/kaniko)
within on-demand launched runners. This should work in all supported
Waypoint installation environments by default and you should not have
to specify any additional configuration.

### Interface

- Input: **component.Source**
- Output: **docker.Image**

### Examples

```hcl
build {
  use "docker" {
	buildkit    = false
	disable_entrypoint = false
  }
}
```

### Required Parameters

This plugin has no required parameters.

### Optional Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### auth (category)

The authentication information to log into the docker repository.


- **Optional**

##### auth.auth

- Type: **string**
- **Optional**

##### auth.email

- Type: **string**
- **Optional**

##### auth.hostname

Hostname of Docker registry.

- Type: **string**
- **Optional**

##### auth.identityToken

Token used to authenticate user.

- Type: **string**
- **Optional**

##### auth.password

Password of Docker registry account.

- Type: **string**
- **Optional**

##### auth.registryToken

Bearer tokens to be sent to Docker registry.

- Type: **string**
- **Optional**

##### auth.serverAddress

Address of Docker registry.

- Type: **string**
- **Optional**

##### auth.username

Username of Docker registry account.

- Type: **string**
- **Optional**

#### build_args

Build args to pass to docker for the build step.

A map of key/value pairs passed as build-args to docker for the build step.

- Type: **map of string to string**
- **Optional**

#### buildkit

If set, use the buildkit builder from Docker.

- Type: **bool**
- **Optional**

#### context

Build context path.

- Type: **string**
- **Optional**

#### disable_entrypoint

If set, the entrypoint binary won't be injected into the image.

The entrypoint binary is what provides extended functionality such as logs and exec. If it is not injected at build time the expectation is that the image already contains it.

- Type: **bool**
- **Optional**

#### dockerfile

The path to the Dockerfile.

Set this when the Dockerfile is not APP-PATH/Dockerfile.

- Type: **string**
- **Optional**

#### no_cache

Do not use cache when building the image.

Ensures a clean image build.

- Type: **bool**
- **Optional**

#### platform

Set target platform to build container if server is multi-platform capable.

Must enable Docker buildkit to use the 'platform' flag.

- Type: **string**
- **Optional**

#### target

The target build stage in a multi-stage Dockerfile.

If buildkit is enabled unused stages will be skipped.

- Type: **string**
- **Optional**

### Output Attributes

Output attributes can be used in your `waypoint.hcl` as [variables](../docs/waypoint-hcl/variables) via [`artifact`](../docs/waypoint-hcl/variables/artifact) or [`deploy`](../docs/waypoint-hcl/variables/deploy).

#### architecture

- Type: **string**

#### image

- Type: **string**

#### location

- Type: **docker.isImage_Location**

#### tag

- Type: **string**

## docker-pull (builder)

Use an existing, pre-built Docker image.

This builder will automatically inject the Waypoint entrypoint. You
can disable this with the "disable_entrypoint" configuration.

If you wish to rename or retag an image, use this along with the
"docker" registry option which will rename/retag the image and then
push it to the specified registry.

If Docker isn't available (the Docker daemon isn't running or a DOCKER_HOST
isn't set), a daemonless solution will be used instead.

If "disable_entrypoint" is set to true and the Waypoint configuration
has no registry, this builder will not physically pull the image. This enables
Waypoint to work in environments where the image is built outside of Waypoint
(such as in a CI pipeline).

### Interface

- Input: **component.Source**
- Output: **docker.Image**

### Examples

```hcl
build {
  use "docker-pull" {
    image = "gcr.io/my-project/my-image"
    tag   = "abcd1234"
  }
}
```

### Required Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### image

The image to pull.

This should NOT include the tag (the value following the ':' in a Docker image). Use `tag` to define the image tag.

- Type: **string**

#### tag

The tag of the image to pull.

- Type: **string**

### Optional Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### auth (category)

The authentication information to log into the docker repository.


- **Optional**

##### auth.auth

- Type: **string**
- **Optional**

##### auth.email

- Type: **string**
- **Optional**

##### auth.hostname

Hostname of Docker registry.

- Type: **string**
- **Optional**

##### auth.identityToken

Token used to authenticate user.

- Type: **string**
- **Optional**

##### auth.password

Password of Docker registry account.

- Type: **string**
- **Optional**

##### auth.registryToken

Bearer tokens to be sent to Docker registry.

- Type: **string**
- **Optional**

##### auth.serverAddress

Address of Docker registry.

- Type: **string**
- **Optional**

##### auth.username

Username of Docker registry account.

- Type: **string**
- **Optional**

#### disable_entrypoint

If set, the entrypoint binary won't be injected into the image.

The entrypoint binary is what provides extended functionality such as logs and exec. If it is not injected at build time the expectation is that the image already contains it.

- Type: **bool**
- **Optional**

#### encoded_auth

The authentication information to log into the docker repository.

WARNING: be very careful to not leak the authentication information by hardcoding it here. Use a helper function like `file()` to read the information from a file not stored in VCS.

- Type: **string**
- **Optional**

### Output Attributes

Output attributes can be used in your `waypoint.hcl` as [variables](../docs/waypoint-hcl/variables) via [`artifact`](../docs/waypoint-hcl/variables/artifact) or [`deploy`](../docs/waypoint-hcl/variables/deploy).

#### architecture

- Type: **string**

#### image

- Type: **string**

#### location

- Type: **docker.isImage_Location**

#### tag

- Type: **string**

## docker-ref (builder)

Use an existing, pre-built Docker image without modifying it.

### Interface

- Input: **component.Source**
- Output: **docker.Image**

### Examples

```hcl
build {
  use "docker-ref" {
    image = "gcr.io/my-project/my-image"
    tag   = "abcd1234"
  }
}
```

### Required Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### image

The image to pull.

This should NOT include the tag (the value following the ':' in a Docker image). Use `tag` to define the image tag.

- Type: **string**

#### tag

The tag of the image to pull.

- Type: **string**

### Optional Parameters

This plugin has no optional parameters.

### Output Attributes

Output attributes can be used in your `waypoint.hcl` as [variables](../docs/waypoint-hcl/variables) via [`artifact`](../docs/waypoint-hcl/variables/artifact) or [`deploy`](../docs/waypoint-hcl/variables/deploy).

#### architecture

- Type: **string**

#### image

- Type: **string**

#### location

- Type: **docker.isImage_Location**

#### tag

- Type: **string**

## docker (registry)

Push a Docker image to a Docker compatible registry.

### Interface

- Input: **docker.Image**
- Output: **docker.Image**

### Examples

```hcl
build {
  use "docker" {}
  registry {
    use "docker" {
      image = "hashicorp/http-echo"
      tag   = "latest"
    }
  }
}
```

### Required Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### image

The image to push the local image to, fully qualified.

This value must be the fully qualified name to the image. for example: gcr.io/waypoint-demo/demo.

- Type: **string**

#### tag

The tag for the new image.

This is added to image to provide the full image reference.

- Type: **string**

### Optional Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### auth (category)

The authentication information to log into the docker repository.


- **Optional**

##### auth.auth

- Type: **string**
- **Optional**

##### auth.email

- Type: **string**
- **Optional**

##### auth.hostname

Hostname of Docker registry.

- Type: **string**
- **Optional**

##### auth.identityToken

Token used to authenticate user.

- Type: **string**
- **Optional**

##### auth.password

Password of Docker registry account.

- Type: **string**
- **Optional**

##### auth.registryToken

Bearer tokens to be sent to Docker registry.

- Type: **string**
- **Optional**

##### auth.serverAddress

Address of Docker registry.

- Type: **string**
- **Optional**

##### auth.username

Username of Docker registry account.

- Type: **string**
- **Optional**

#### encoded_auth

The authentication information to log into the docker repository.

The format of this is base64-encoded JSON. The structure is the [`AuthConfig`](https://pkg.go.dev/github.com/docker/cli/cli/config/types#AuthConfig) structure used by Docker.
  WARNING: be very careful to not leak the authentication information by hardcoding it here. Use a helper function like `file()` to read the information from a file not stored in VCS.

- Type: **string**
- **Optional**

#### insecure

Access the registry via http rather than https.

This indicates that the registry should be accessed via http rather than https. Not recommended for production usage.

- Type: **bool**
- **Optional**

#### local

If set, the image will only be tagged locally and not pushed to a remote repository.

- Type: **bool**
- **Optional**

#### password

Password associated with username on the registry.

This optional conflicts with encoded_auth and thusly only one can be used at a time. If both are used, encoded_auth takes precedence.

- Type: **string**
- **Optional**

#### username

Username to authenticate with the registry.

This optional conflicts with encoded_auth and thusly only one can be used at a time. If both are used, encoded_auth takes precedence.

- Type: **string**
- **Optional**

### Output Attributes

Output attributes can be used in your `waypoint.hcl` as [variables](../docs/waypoint-hcl/variables) via [`artifact`](../docs/waypoint-hcl/variables/artifact) or [`deploy`](../docs/waypoint-hcl/variables/deploy).

#### architecture

- Type: **string**

#### image

- Type: **string**

#### location

- Type: **docker.isImage_Location**

#### tag

- Type: **string**

## docker (platform)

Deploy a container to Docker, local or remote.

### Interface

- Input: **docker.Image**
- Output: **docker.Deployment**

### Examples

```hcl
deploy {
  use "docker" {
	command      = ["ps"]
	service_port = 3000
	static_environment = {
	  "environment": "production",
	  "LOG_LEVEL": "debug"
	}
  }
}
```

### Required Parameters

This plugin has no required parameters.

### Optional Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### auth (category)

The authentication information to log into the docker repository.


- **Optional**

##### auth.auth

- Type: **string**
- **Optional**

##### auth.email

- Type: **string**
- **Optional**

##### auth.hostname

Hostname of Docker registry.

- Type: **string**
- **Optional**

##### auth.identityToken

Token used to authenticate user.

- Type: **string**
- **Optional**

##### auth.password

Password of Docker registry account.

- Type: **string**
- **Optional**

##### auth.registryToken

Bearer tokens to be sent to Docker registry.

- Type: **string**
- **Optional**

##### auth.serverAddress

Address of Docker registry.

- Type: **string**
- **Optional**

##### auth.username

Username of Docker registry account.

- Type: **string**
- **Optional**

#### binds

A 'source:destination' list of folders to mount onto the container from the host.

A list of folders to mount onto the container from the host. The expected format for each string entry in the list is `source:destination`. So for example: `binds: ["host_folder/scripts:/scripts"].

- Type: **list of string**
- **Optional**

#### client_config

Client config for remote Docker engine.

This config block can be used to configure a remote Docker engine. By default Waypoint will attempt to discover this configuration using the environment variables: `DOCKER_HOST` to set the url to the docker server. `DOCKER_API_VERSION` to set the version of the API to reach, leave empty for latest. `DOCKER_CERT_PATH` to load the TLS certificates from. `DOCKER_TLS_VERIFY` to enable or disable TLS verification, off by default.

- Type: **docker.ClientConfig**
- **Optional**

#### command

The command to run to start the application in the container.

- Type: **list of string**
- **Optional**
- Default: the image entrypoint

#### extra_ports

Additional TCP ports the application is listening on to expose on the container.

Used to define and expose multiple ports that the application is listening on for the container in use. These ports will get merged with service_port when creating the container if defined.

- Type: **list of uint**
- **Optional**

#### force_pull

Always pull the docker container from the registry.

- Type: **bool**
- **Optional**
- Default: false

#### labels

A map of key/value pairs to label the docker container with.

A map of key/value pair(s), stored in docker as a string. Each key/value pair must be unique. Validiation occurs at the docker layer, not in Waypoint. Label keys are alphanumeric strings which may contain periods (.) and hyphens (-).

- Type: **map of string to string**
- **Optional**

#### networks

An list of strings with network names to connect the container to.

A list of networks to connect the container to. By default the container will always connect to the `waypoint` network.

- Type: **list of string**
- **Optional**
- Default: waypoint

#### resources

A map of resources to configure the container with, such as memory or cpu limits.

These options are used to configure the container used when deploying with docker. Currently, the supported resources are 'memory' and 'cpu' limits. The field 'memory' is expected to be defined as "512MB", "44kB", etc.

- Type: **map of string to string**
- **Optional**

#### scratch_path

A path within the container to store temporary data.

Docker will mount a tmpfs at this path.

- Type: **string**
- **Optional**

#### service_port

Port that your service is running on in the container.

- Type: **uint**
- **Optional**
- Default: 3000

#### static_environment

Environment variables to expose to the application.

These environment variables should not be run of the mill configuration variables, use waypoint config for that. These variables are used to control over all container modes, such as configuring it to start a web app vs a background worker.

- Type: **map of string to string**
- **Optional**

### Output Attributes

Output attributes can be used in your `waypoint.hcl` as [variables](../docs/waypoint-hcl/variables) via [`artifact`](../docs/waypoint-hcl/variables/artifact) or [`deploy`](../docs/waypoint-hcl/variables/deploy).

#### container

- Type: **string**

#### id

- Type: **string**

#### name

- Type: **string**

#### resource_state

- Type: **opaqueany.Any**

## docker (task)

Launch a Docker container as a task.

If a Docker server is available (either locally or via environment variables
such as "DOCKER_HOST"), then it will be used to start the container.

### Interface

### Examples

```hcl
task {
  use "docker" {
		force_pull = true
  }
}
```

### Required Parameters

This plugin has no required parameters.

### Optional Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### binds

A 'source:destination' list of folders to mount onto the container from the host.

A list of folders to mount onto the container from the host. The expected format for each string entry in the list is `source:destination`. So for example: `binds: ["host_folder/scripts:/scripts"]`.

- Type: **list of string**
- **Optional**

#### client_config

- Type: **docker.ClientConfig**
- **Optional**

#### debug_containers

- Type: **bool**
- **Optional**

#### force_pull

- Type: **bool**
- **Optional**

#### labels

A map of key/value pairs to label the docker container with.

A map of key/value pair(s), stored in docker as a string. Each key/value pair must be unique. Validiation occurs at the docker layer, not in Waypoint. Label keys are alphanumeric strings which may contain periods (.) and hyphens (-).

- Type: **map of string to string**
- **Optional**

#### networks

A list of strings with network names to connect the container to.

A list of networks to connect the container to. By default the container will always connect to the `waypoint` network.

- Type: **list of string**
- **Optional**
- Default: waypoint

#### resources (category)

The resources that the tasks should use.


- **Optional**

##### resources.cpu

The cpu shares that the tasks should use.

- Type: **int64**
- **Optional**

##### resources.memory

The amount of memory to use. Defined as '512MB', '44kB', etc.

- Type: **string**
- **Optional**

#### static_environment

Environment variables to expose to the application.

These variables are used to control all of a container's modes, such as configuring it to start a web app vs a background worker. These environment variables should not be common configuration variables normally set in `waypoint config`.

- Type: **map of string to string**
- **Optional**

### Output Attributes

Output attributes can be used in your `waypoint.hcl` as [variables](../docs/waypoint-hcl/variables) via [`artifact`](../docs/waypoint-hcl/variables/artifact) or [`deploy`](../docs/waypoint-hcl/variables/deploy).

#### id

- Type: **string**
