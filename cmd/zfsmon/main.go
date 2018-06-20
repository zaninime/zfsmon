package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"zanini.me/zfsmon/clif"
)

type zpoolMetrics struct {
	cmd         *clif.ZpoolCommand
	healthyDesc *prometheus.Desc
}

func newZpoolMetrics(cmd *clif.ZpoolCommand) *zpoolMetrics {
	return &zpoolMetrics{
		cmd:         cmd,
		healthyDesc: prometheus.NewDesc("zfs_pool_healthy", "Check pool health. ONLINE == 1, else 0", []string{"name"}, nil),
	}
}

func (zm *zpoolMetrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- zm.healthyDesc
}

func (zm *zpoolMetrics) Collect(ch chan<- prometheus.Metric) {
	poolsHealth, err := zm.cmd.ListAllByPoolName("health")

	if err != nil {
		log.Fatal(err)
	}

	for poolName, healthStr := range poolsHealth {
		health, err := clif.NewHealthFromCliOutput(healthStr)

		if err != nil {
			log.Fatal(err)
		}

		var gaugeValue float64

		if *health == clif.Online {
			gaugeValue = 1
		} else {
			gaugeValue = 0
		}

		ch <- prometheus.MustNewConstMetric(zm.healthyDesc, prometheus.GaugeValue, gaugeValue, poolName)
	}
}

func main() {
	cmd := clif.NewDefaultZpoolCommand()
	prometheus.MustRegister(newZpoolMetrics(cmd))

	http.Handle("/metrics", promhttp.Handler())
	log.Println("Beginning to serve on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
