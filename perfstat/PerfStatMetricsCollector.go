package prometheus

import (
	"github.com/go-perfstat/go/perfstat"
	"github.com/prometheus/client_golang/prometheus"
)

type PerfStatMetricsCollector struct {
	minTimeMs        *prometheus.Desc
	minTimeSampleMs  *prometheus.Desc
	avgTimeMs        *prometheus.Desc
	avgTimeSampleMs  *prometheus.Desc
	maxTimeMs        *prometheus.Desc
	maxTimeSampleMs  *prometheus.Desc
	totalTimeSec     *prometheus.Desc
	leapsCount       *prometheus.Desc
	leapsCountSample *prometheus.Desc
	peersCount       *prometheus.Desc
}

func NewPerfStatMetricsCollector() *PerfStatMetricsCollector {
	return NewPerfStatMetricsCollectorWithLabels(prometheus.Labels{})
}

func NewPerfStatMetricsCollectorWithLabels(labels prometheus.Labels) *PerfStatMetricsCollector {
	return &PerfStatMetricsCollector{
		minTimeMs:        prometheus.NewDesc("perfstat_minTimeMs", "Global Min(ms)", []string{"type", "name"}, labels),
		minTimeSampleMs:  prometheus.NewDesc("perfstat_minTimeSampleMs", "Min(ms)", []string{"type", "name"}, labels),
		avgTimeMs:        prometheus.NewDesc("perfstat_avgTimeMs", "Global Avg(ms)", []string{"type", "name"}, labels),
		avgTimeSampleMs:  prometheus.NewDesc("perfstat_avgTimeSampleMs", "Avg(ms)", []string{"type", "name"}, labels),
		maxTimeMs:        prometheus.NewDesc("perfstat_maxTimeMs", "Global Max(ms)", []string{"type", "name"}, labels),
		maxTimeSampleMs:  prometheus.NewDesc("perfstat_maxTimeSampleMs", "Max(ms)", []string{"type", "name"}, labels),
		totalTimeSec:     prometheus.NewDesc("perfstat_totalTimeSec", "Total(s)", []string{"type", "name"}, labels),
		leapsCount:       prometheus.NewDesc("perfstat_leapsCount", "Total Leaps", []string{"type", "name"}, labels),
		leapsCountSample: prometheus.NewDesc("perfstat_leapsCountSample", "Leaps", []string{"type", "name"}, labels),
		peersCount:       prometheus.NewDesc("perfstat_peersCount", "Peers", []string{"type", "name"}, labels),
	}
}

func (c *PerfStatMetricsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.minTimeMs
	ch <- c.minTimeSampleMs
	ch <- c.avgTimeMs
	ch <- c.avgTimeSampleMs
	ch <- c.maxTimeMs
	ch <- c.maxTimeSampleMs
	ch <- c.totalTimeSec
	ch <- c.leapsCount
	ch <- c.leapsCountSample
	ch <- c.peersCount
}

func (c *PerfStatMetricsCollector) Collect(ch chan<- prometheus.Metric) {
	stats := perfstat.GetAll()
	for typ, names := range stats {
		for name, stat := range names {
			ch <- prometheus.MustNewConstMetric(c.minTimeMs, prometheus.GaugeValue, stat.GetMinTimeMs(), typ, name)
			ch <- prometheus.MustNewConstMetric(c.minTimeSampleMs, prometheus.GaugeValue, stat.GetMinTimeSampleMs(), typ, name)
			ch <- prometheus.MustNewConstMetric(c.avgTimeMs, prometheus.GaugeValue, stat.GetAvgTimeMs(), typ, name)
			ch <- prometheus.MustNewConstMetric(c.avgTimeSampleMs, prometheus.GaugeValue, stat.GetAvgTimeSampleMs(), typ, name)
			ch <- prometheus.MustNewConstMetric(c.maxTimeMs, prometheus.GaugeValue, stat.GetMaxTimeMs(), typ, name)
			ch <- prometheus.MustNewConstMetric(c.maxTimeSampleMs, prometheus.GaugeValue, stat.GetMaxTimeSampleMs(), typ, name)
			ch <- prometheus.MustNewConstMetric(c.totalTimeSec, prometheus.GaugeValue, stat.GetTotalTimeMs()/1000.0, typ, name)
			ch <- prometheus.MustNewConstMetric(c.leapsCount, prometheus.GaugeValue, float64(stat.GetLeapsCount()), typ, name)
			ch <- prometheus.MustNewConstMetric(c.leapsCountSample, prometheus.GaugeValue, float64(stat.GetLeapsCountSample()), typ, name)
			ch <- prometheus.MustNewConstMetric(c.peersCount, prometheus.GaugeValue, float64(stat.GetPeersCount()), typ, name)
		}
	}
}
