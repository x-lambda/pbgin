package sqlx

import (
	"github.com/prometheus/client_golang/prometheus"
)

var defBuckets = []float64{.005, .01, .025, .05, .1, .25, .5, 1}

// sqlDurations 上报prometheus
var sqlDurations = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Namespace: "aidi",
	Subsystem: "sqlx",
	Name:      "sql_durations_seconds",
	Help:      "sql latency distributions",
	Buckets:   defBuckets,
}, []string{"db_name", "table", "cmd"})

func init() {
	prometheus.MustRegister(sqlDurations)
}
