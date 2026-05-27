# go-perfstat/prometheus

Prometheus exporter for [go-perfstat](https://github.com/go-perfstat/go).

The collector exposes all registered `go-perfstat` statistics as Prometheus
metrics with `type` and `name` labels.

## Register collector

```go
prometheus.MustRegister(perfstat.NewPerfStatMetricsCollector())
```

With custom constant labels:

```go
prometheus.MustRegister(
	perfstat.NewPerfStatMetricsCollectorWithLabels(prometheus.Labels{
		"application": "my-service",
	}),
)
```

> Default aggregation period is 15s.

## Expose `/metrics`

```go
http.Handle("/metrics", promhttp.Handler())
log.Fatal(http.ListenAndServe(":9090", nil))
```

## Metrics

The collector exports global and aggregation-period statistics:

```text
perfstat_minTimeMs
perfstat_minTimeSampleMs
perfstat_avgTimeMs
perfstat_avgTimeSampleMs
perfstat_maxTimeMs
perfstat_maxTimeSampleMs
perfstat_totalTimeSec
perfstat_leapsCount
perfstat_leapsCountSample
perfstat_peersCount
```

Each metric has the following labels:

```text
type
name
```

## Grafana dashboard

A ready-to-import Grafana dashboard is available here:

https://github.com/go-perfstat/prometheus/blob/main/dashboard/PerfStat.json
