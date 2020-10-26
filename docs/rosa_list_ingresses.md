## rosa list ingresses

List cluster Ingresses

### Synopsis

List API and ingress endpoints for a cluster.

```
rosa list ingresses [flags]
```

### Examples

```
  # List all routes on a cluster named "mycluster"
  rosa list ingresses --cluster=mycluster
```

### Options

```
  -c, --cluster string   Name or ID of the cluster to list the routes of (required).
  -h, --help             help for ingresses
```

### Options inherited from parent commands

```
      --debug            Enable debug mode.
      --profile string   Use a specific AWS profile from your credential file.
  -v, --v Level          log level for V logs
```

### SEE ALSO

* [rosa list](rosa_list.md)	 - List all resources of a specific type

