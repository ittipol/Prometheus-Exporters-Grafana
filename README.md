## Prometheus
- Docker https://hub.docker.com/r/prom/prometheus
- Helm Chart https://artifacthub.io/packages/helm/prometheus-community/prometheus

## Exporters
- https://prometheus.io/docs/instrumenting/exporters/
- https://github.com/prometheus-community/helm-charts/tree/main/charts
- Docker https://hub.docker.com/u/prom

## Grafana
- https://prometheus.io/docs/visualization/grafana/
- Docker https://hub.docker.com/r/grafana/grafana

## Prometheus exporter port numbers
- https://prometheus.io/docs/instrumenting/writing_exporters/
- https://github.com/prometheus/prometheus/wiki/Default-port-allocations

## Grafana - Prometheus data source
1. Go to menu Data sources
2. Click "Add new data source" button
3. Select "Prometheus"
4. Input URL e.g. http://prometheus:9090
5. Click "Save & test" button

## Grafana - Import dashboard ID
- [Grafana Dashboards](https://grafana.com/grafana/dashboards/)

### Search dashboard
1. Go to https://grafana.com/grafana/dashboards/
2. Filter by > select Data Data Source > prometheus
3. Select a dashboard
4. Click button "Copy ID to clipboard"