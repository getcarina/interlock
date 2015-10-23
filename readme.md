# Interlock
Dynamic, event-driven Docker plugin system using [Swarm](https://github.com/docker/swarm).

This fork relies on [libcarina](https://github.com/getcarina/libcarina) to connect
directly to Carina for Swarm cluster access.

# Usage
Run `docker run carina/interlock list-plugins` to show available plugins.

Example:

```
docker run carina/interlock --username $USERNAME \
                            --apikey $APIKEY \
                            --clustername boatie \
                            --plugin example start
```

# Commandline options

- `--username`: username for Carina
- `--api-key`: API key for Carina
- `--clustername`: swarm cluster to connect interlock to
- `--plugin`: enable plugin
- `--debug`: enable debug output
- `--version`: show version and exit

# Plugins
See the [Plugins](https://github.com/getcarina/interlock/tree/master/plugins)
directory for available plugins and their corresponding readme.md for usage.

| Name | Description |
|-----|-----|
| [Example](https://github.com/getcarina/interlock/tree/master/plugins/example) | Example Plugin for Reference|
| [HAProxy](https://github.com/getcarina/interlock/tree/master/plugins/haproxy) | [HAProxy](http://www.haproxy.org/) Load Balancer |
| [Nginx](https://github.com/getcarina/interlock/tree/master/plugins/nginx) | [Nginx](http://nginx.org) Load Balancer |
| [Stats](https://github.com/getcarina/interlock/tree/master/plugins/stats) | Container stat forwarding to [Carbon](http://graphite.wikidot.com/carbon) |

# License
Licensed under the Apache License, Version 2.0. See LICENSE for full license text.
