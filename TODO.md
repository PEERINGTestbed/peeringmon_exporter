# TODO

After the refactoring, the following tasks are still pending:

* Get the Grafana dashboard configuration JSON file from Ezri's existing deployment and add it to the `config` directory.

* Integrate deployment and configuration of the PEERING client that the announcement controller (`peeringmon_controller`) depends on.

* Possibly migrate the controller to use the default client Python module to control announcements and send updates to Prometheus.  This may be needed as there is a comment mentioning some memory leak if upstreams (muxes) send routes to the controller.
