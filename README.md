# PEERING Prefix Monitor

This repository implements a Prometheus exporter to get prefix visibility status from RIPE RIS via the [RIPEStat API][ripestat].

[ripestat]: https://stat.ripe.net/docs/02.data-api/

## Deploying

The `docker-compose.yml` defines services for the exporter, controller, Prometheus, and Grafana.  The exporter collects BGP routes from RIPE RIS, while the controller makes prefix announcements from PEERING and exports the corresponding information to Prometheus.

> Despite the phrasing above, it is worth noting that our Prometheus deployment works in a pull-based fashion.  So both the exporter and controller only store information, which is pulled by Prometheus periodically.  Grafana does pull from Prometheus to show data in the dashboard.

Configurable parameters are split into `docker-compose-config.yml`.  Two things need to be configured:

1. The containers need to store information across restarts, so we create volumes from `data/grafana` and `data/prometheus`, which are mapped inside the containers.  For the containers to work, you should change the `user` used to run the containers to the ID of the user that owns the `data` directory.  Run `id -u <username>` to get a user's ID.

2. The containers also need configuration files, which are mounted from `config/prometheus.yml` and `config/grafana.yml`.  You may edit the configuration files if necessary to suit your needs, but the provided files should work out of the box.

## Other References

* [RIPEStat ToS](https://www.ripe.net/about-us/legal/ripestat-service-terms-and-conditions/)

* [RIPE RIS Peer List](https://www.ris.ripe.net/peerlist/all.shtml)
