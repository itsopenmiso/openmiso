---
layout: docs
page_title: Files - Application Configuration
description: |-
  Waypoint can write application configuration to files. This is useful for applications that expect their configuration at a certain path or in a certain format.
---

# Configuration Files

Waypoint can write application configuration to files. This is useful for
applications that expect their configuration at a certain path or in a certain
format. For example, you can write JSON, YAML, XML files for web frameworks
or other services.

## Writing a Configuration File

Application configuration file rendering must be configured in the
`waypoint.hcl` file in the [`config` stanza](../docs/waypoint-hcl/config).
The example below writes a file `config/database.yml` relative to your
deployed and running application:

```hcl
config {
  file = {
    "config/database.yml" = templatefile("${path.project}/database.yml.tpl", {
      host = "db.service"
    })
  }
}
```

The template `database.yml.tpl` may look like this:

```yaml
production:
  host: ${host}
  username: myuser
  password: mypass
```

Notice that since we're using the `templatefile` function, we can template
the file by using variables such as `${host}` within our file. In this example,
we set host to a static string, but the variable may also reference
[internal values](../docs/app-config/internal) which allow files to be built
up of dynamic values as well. When a value
changes at runtime, Waypoint will re-render your configuration file and
send a signal to your application.

## File Change Signal

If the input variables to a configuration file change at runtime, Waypoint
will rerender the file and signal your application. By default, Waypoint will
send the signal `SIGUSR2`. Waypoint expects your application to respond to this
signal by reloading the configuration.

The file change signal can be customized using the
[`file_change_signal` field](../docs/waypoint-hcl/config#file_change_signal)
in the `config` stanza. For example, you `SIGHUP` is a common signal
for reloading configuration. You could configure your application to
receive this signal:

```hcl
config {
  file_change_signal = "HUP"
}
```
