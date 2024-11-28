Copying the docker-compose.yml to your system, ensure the user directives is set
to a non-root user ID.

Create /srv/dockerdata or change the mountpoint for the prom and grafana
containers.

Copy prometheus.yml over to the mointpoint declared in the prometheus
container's volume settings, by default it's `/srv/dockerdata/prometheus.yml`

Start via `docker compose up -d`
