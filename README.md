### Register in Prometheus

```go
prometheus.MustRegister(perfstat.NewPerfStatMetricsCollector())
```

> Default aggregation period is 15s

### Expose as Grafana dashboard

[PerfStat Grafana Dashboard](https://github.com/go-perfstat/prometheus/blob/main/dashboard/PerfStat.json)
