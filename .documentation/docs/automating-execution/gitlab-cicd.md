---
layout: docs
page_title: Integrating Waypoint with GitLab CI/CD
description: |-
  How to utilize Waypoint in a GitLab development workflow and
  with GitLab CI/CD
---

# Integrating Waypoint with GitLab CI/CD

Using Waypoint to deploy an application from within GitLab CI/CD is similar to how you might deploy an application from your own workspace.

The example demonstrates the main steps:

1. **Set-up the dependencies Waypoint might use.** This could be
   a Kubernetes context for a more advanced application, or in the
   below example, a Docker daemon to run applications on.
2. **Install Waypoint from the official source.** If you are
   using a Docker custom image executor with your other deploy
   dependencies, you could utilize the public Docker image in a multi-stage build.
3. **Run `waypoint init`.** This depends on the environment variables listed
   below and documented in the [Automating Execution](../docs/automating-execution#init)
   overview.
4. **Run the build, deploy, and release.** In this case, instead of using `waypoint up`,
   it breaks out each stage as a separate command to be easier to read and
   filter through in the GitLab UI.

## Workspaces

This example assumes the use of a single default workspace.

You can also use defined GitLab environments:

```yaml
run: waypoint build -workspace $CI_ENVIRONMENT_NAME
```

If this was in a job triggered by a Git commit or merge request and may be an ephemeral development environment, you may want to interpolate the relevant Git ref for the workspace parameter, as demonstrated below:

```yaml
run: waypoint build -workspace $CI_COMMIT_BRANCH
```

See the GitLab CI/CD [predefined environment variables](https://docs.gitlab.com/ee/ci/variables/predefined_variables.html) page for a full list of variables that could be utilized in this way.

## Example

```yaml
waypoint:
  # GitLab created images, which basically rewraps HashiCorp's officially released binaries and may not match HashiCorp's checksums, See https://gitlab.com/gitlab-org/waypoint-images for more details about what goes into the images.
  image: registry.gitlab.com/gitlab-org/waypoint-images:latest
  stage: build
  services:
    - docker:dind
  variables:
    WAYPOINT_SERVER_ADDR: ''
    WAYPOINT_SERVER_TOKEN: ''
    WAYPOINT_SERVER_TLS: '1'
    WAYPOINT_SERVER_TLS_SKIP_VERIFY: '1'
  script:
    - waypoint init
    - waypoint build
    - waypoint deploy
    - waypoint release
```
