package prometheus

import (
	"math"

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

func (this *PerfStatMetricsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- this.minTimeMs
	ch <- this.minTimeSampleMs
	ch <- this.avgTimeMs
	ch <- this.avgTimeSampleMs
	ch <- this.maxTimeMs
	ch <- this.maxTimeSampleMs
	ch <- this.totalTimeSec
	ch <- this.leapsCount
	ch <- this.leapsCountSample
	ch <- this.peersCount
}

func (this *PerfStatMetricsCollector) Collect(ch chan<- prometheus.Metric) {
	stats := perfstat.GetAll()
	for typ, names := range stats {
		for name, stat := range names {
			ch <- prometheus.MustNewConstMetric(this.minTimeMs, prometheus.GaugeValue, stat.GetMinTimeMs(), typ, name)
			ch <- prometheus.MustNewConstMetric(this.minTimeSampleMs, prometheus.GaugeValue, stat.GetMinTimeSampleMs(), typ, name)
			ch <- prometheus.MustNewConstMetric(this.avgTimeMs, prometheus.GaugeValue, stat.GetAvgTimeMs(), typ, name)
			ch <- prometheus.MustNewConstMetric(this.avgTimeSampleMs, prometheus.GaugeValue, stat.GetAvgTimeSampleMs(), typ, name)
			ch <- prometheus.MustNewConstMetric(this.maxTimeMs, prometheus.GaugeValue, stat.GetMaxTimeMs(), typ, name)
			ch <- prometheus.MustNewConstMetric(this.maxTimeSampleMs, prometheus.GaugeValue, stat.GetMaxTimeSampleMs(), typ, name)
			ch <- prometheus.MustNewConstMetric(this.totalTimeSec, prometheus.GaugeValue, math.Round(stat.GetTotalTimeMs()/1000.0*100)/100, typ, name)
			ch <- prometheus.MustNewConstMetric(this.leapsCount, prometheus.GaugeValue, float64(stat.GetLeapsCount()), typ, name)
			ch <- prometheus.MustNewConstMetric(this.leapsCountSample, prometheus.GaugeValue, float64(stat.GetLeapsCountSample()), typ, name)
			ch <- prometheus.MustNewConstMetric(this.peersCount, prometheus.GaugeValue, float64(stat.GetPeersCount()), typ, name)
		}
	}
}
