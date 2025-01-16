---
layout: plugins
page_title: 'Plugin: AWS Lambda'
description: 'Build, Release, and Deploy on AWS Lambda'
---

# AWS Lambda

## Image Support Only

The AWS Lambda deploy plugin only supports AWS Lambda deployments that utilize ECR images
rather than .zip files.

We recommend that the `docker` build plugin is used to generate these images.

Below is a simple Dockerfile, handler code, and waypoint.hcl for building, deploying, and
releasing using Waypoint.

### handler.rb

```ruby
require 'json'

def handler(event:, context:)
    qs = event.fetch("queryStringParameters", {})

    name = qs.fetch("name", "unknown user")

    STDERR.puts "Handling ALB request for #{name}"

    {
        "statusCode": 200,
        "statusDescription": "200 OK",
        "isBase64Encoded": false,
        "headers": {
            "Content-Type": "text/html"
        },
        "body": "<html><body><h1>Hello there #{name} from Lambda!</h1></body></html>"
    }
end
```

### Dockerfile

```docker
FROM public.ecr.aws/lambda/ruby:2.7

COPY handler.rb /var/task

CMD [ "handler.handler" ]
```

### waypoint.hcl

```hcl
project = "hello"

app "hello-ruby" {
  build {
    use "docker" {}

    registry {
      use "aws-ecr" {
        region = "us-west-2"
        repository = "hello-ruby"
        tag = gitrefpretty()
      }
    }
  }

  deploy {
    use "aws-lambda" {
      region = "us-west-2"
    }
  }

  release {
    use "aws-alb" {

    }
  }
}
```

# Reference

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

## aws-ecr-pull (builder)

Use an existing, pre-built AWS ECR image.

This builder attempts to find an image by repository and tag in the
specified region. If found, it will pass along the image information
to the next step.

This builder will not modify the image. 

If you wish to rename or retag an image, please use the "docker-pull" component
in conjunction with the "aws-ecr" registry option.

### Interface

- Input: **component.Source**
- Output: **ecr.Image**

### Examples

```hcl
build {
  use "aws-ecr-pull" {
    region     = "us-east-1"
    repository = "deno-http"
    tag        = "latest"
  }
}
```

### Required Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### repository

The AWS ECR repository name.

- Type: **string**

#### tag

The tag of the image to pull.

- Type: **string**

### Optional Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### disable_entrypoint

If set, the entrypoint binary won't be injected into the image.

The entrypoint binary is what provides extended functionality such as logs and exec. If it is not injected at build time the expectation is that the image already contains it.

- Type: **bool**
- **Optional**

#### force_architecture

**Note**: This is a temporary field that enables overriding the `architecture` output attribute. Valid values are: `"x86_64"`, `"arm64"`.

- Type: **string**
- **Optional**
- Default: `""`

#### region

The AWS region the ECR repository is in.

If not set uses the environment variable AWS_REGION or AWS_REGION_DEFAULT.

- Type: **string**
- **Optional**

### Output Attributes

Output attributes can be used in your `waypoint.hcl` as [variables](../docs/waypoint-hcl/variables) via [`artifact`](../docs/waypoint-hcl/variables/artifact) or [`deploy`](../docs/waypoint-hcl/variables/deploy).

#### architecture

- Type: **string**

#### image

- Type: **string**

#### tag

- Type: **string**

## aws-ecr (registry)

Store a docker image within an Elastic Container Registry on AWS.

### Interface

- Input: **docker.Image**
- Output: **ecr.Image**

### Examples

```hcl
registry {
    use "aws-ecr" {
      region = "us-east-1"
      tag = "latest"
    }
}
```

### Mappers

#### Allow an ECR Image to be used as a standard docker.Image

- Input: **ecr.Image**
- Output: **docker.Image**

### Required Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### tag

The docker tag to assign to the new image.

- Type: **string**

### Optional Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### region

The AWS region the ECR repository is in.

If not set uses the environment variable AWS_REGION or AWS_REGION_DEFAULT.

- Type: **string**
- **Optional**
- Environment Variable: **AWS_REGION_DEFAULT**

#### repository

The ECR repository to store the image into.

This defaults to waypoint- then the application name. The repository will be automatically created if needed.

- Type: **string**
- **Optional**

### Output Attributes

Output attributes can be used in your `waypoint.hcl` as [variables](../docs/waypoint-hcl/variables) via [`artifact`](../docs/waypoint-hcl/variables/artifact) or [`deploy`](../docs/waypoint-hcl/variables/deploy).

#### architecture

- Type: **string**

#### image

- Type: **string**

#### tag

- Type: **string**

## aws-lambda (platform)

Deploy functions as OCI Images to AWS Lambda.

### Interface

- Input: **ecr.Image**
- Output: **lambda.Deployment**

### Examples

```hcl
deploy {
	use "aws-lambda" {
		region = "us-east-1"
		memory = 512
	}
}
```

### Required Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### region

The AWS region for the ECS cluster.

- Type: **string**

### Optional Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### architecture

The instruction set architecture that the function supports. Valid values are: "x86_64", "arm64".

- Type: **string**
- **Optional**
- Default: x86_64

#### iam_role

An IAM Role specified by ARN that will be used by the Lambda at execution time.

- Type: **string**
- **Optional**
- Default: created automatically

#### memory

The amount of memory, in megabytes, to assign the function.

- Type: **int**
- **Optional**
- Default: 256

#### static_environment

Environment variables to expose to the lambda function.

Environment variables that are meant to configure the application in a static way. This might be to control an image that has multiple modes of operation, selected via environment variable. Most configuration should use the waypoint config commands.

- Type: **map of string to string**
- **Optional**

#### storagemb

The storage size (in MB) of the Lambda function's `/tmp` directory. Must be a value between 512 and 10240.

- Type: **int**
- **Optional**
- Default: 512

#### timeout

The number of seconds a function has to return a result.

- Type: **int**
- **Optional**
- Default: 60

### Output Attributes

Output attributes can be used in your `waypoint.hcl` as [variables](../docs/waypoint-hcl/variables) via [`artifact`](../docs/waypoint-hcl/variables/artifact) or [`deploy`](../docs/waypoint-hcl/variables/deploy).

#### func_arn

- Type: **string**

#### id

- Type: **string**

#### region

- Type: **string**

#### storage

- Type: **int64**

#### target_group_arn

- Type: **string**

#### ver_arn

- Type: **string**

#### version

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

## lambda-function-url (releasemanager)

Create an AWS Lambda function URL.

### Interface

- Input: **lambda.Deployment**
- Output: **lambda.Release**

### Examples

```hcl
release {
	use "lambda-function-url" {
		auth_type = "NONE"
		cors {
			allow_methods = ["*"]
		}
	}
}
```

### Required Parameters

This plugin has no required parameters.

### Optional Parameters

These parameters are used in the [`use` stanza](../docs/waypoint-hcl/use) for this plugin.

#### auth_type

The Lambda function URL auth type.

The AuthType parameter determines how Lambda authenticates or authorizes requests to your function URL. Must be either `AWS_IAM` or `NONE`.

- Type: **string**
- **Optional**
- Default: NONE

#### cors (category)

CORS configuration for the function URL.


- **Optional**
- Default: NONE

##### cors.allow_credentials

Whether to allow cookies or other credentials in requests to your function URL.

- Type: **bool**
- **Optional**
- Default: false

##### cors.allow_headers

The HTTP headers that origins can include in requests to your function URL. For example: Date, Keep-Alive, X-Custom-Header.

- Type: **list of string**
- **Optional**
- Default: []

##### cors.allow_methods

The HTTP methods that are allowed when calling your function URL. For example: GET, POST, DELETE, or the wildcard character (*).

- Type: **list of string**
- **Optional**
- Default: []

##### cors.allow_origins

The origins that can access your function URL. You can list any number of specific origins, separated by a comma. You can grant access to all origins using the wildcard character (*).

- Type: **list of string**
- **Optional**
- Default: []

##### cors.expose_headers

The HTTP headers in your function response that you want to expose to origins that call your function URL. For example: Date, Keep-Alive, X-Custom-Header.

- Type: **list of string**
- **Optional**
- Default: []

##### cors.max_age

The maximum amount of time, in seconds, that web browsers can cache results of a preflight request.

- Type: **int64**
- **Optional**
- Default: 0

#### principal

The principal to use when auth_type is `AWS_IAM`.

The Principal parameter specifies the principal that is allowed to invoke the function.

- Type: **string**
- **Optional**
- Default: *
