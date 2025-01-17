---
layout: docs
page_title: 'Creating an Example Application'
description: |-
  How Waypoint plugins work
---

## Creating an Example Application

Video tutorial below:

<iframe
  width="560"
  height="315"
  src="https://www.youtube.com/embed/4zwZCpqo0j8?start=1262&end=1395"
  frameborder="0"
  allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
  allowfullscreen
></iframe>

Create a new folder called `example_app` in your `$HOME` folder; there is no restriction on the location of Waypoint apps,
`$HOME` is just used for the guide.

```shell-session
$ mkdir $HOME/example_app
```

If you look at the following example configuration, you will see that the `use` block for the `build` has a value
`gobuilder`, which is the plugin's name to be loaded for the build component. Waypoint defines the convention that
plugins are always called `waypoint-plugin-<name>`. When you run the `waypoint build` command, Waypoint will search the
known plugin folders for a file called `waypoint-plugin-gobuilder` it will then use that plugin for the `build`
component.

The configuration also has the values `app` for the `output_name` and the source as `./`, we can put our Go application
in the same location as the plugin config. First, create a file `waypoint.hcl` in the folder `$HOME/example_app/` with
the following contents.

```hcl
project = "guides"

app "example" {

  build {
    use "gobuilder" {
      output_name = "app"
      source = "./"
    }
  }

  deploy {
    use "gobuilder" {}
  }
}
```

Now you can create the source for the example application, let's make the most simple Go application possible. Create a
new file `main.go` in your Waypoint application folder and copy the following to it.

```go
package main

import "fmt"

func main() {
  fmt.Println("Hello Waypoint")
}
```

You should now have two files in your application folder `waypoint.hcl` and `main.go`, let's test your new plugin.

[Next - Testing the Plugin](../docs/extending-waypoint/creating-plugins/testing)
