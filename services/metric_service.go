package services

import (
	"errors"
	"golang-go/loggers"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	metricMap      = make(map[string]bool)
	metricMapCount = make(map[string]prometheus.Counter)
)

func generateMetrics(name string) error {
	if exist := metricMap[name]; exist {
		metricMapCount[name].Inc()
		return nil
	}
	newMetrics := prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: name,
		},
	)
	if err := prometheus.Register(newMetrics); err != nil {
		loggers.Loggers.Error("cloud not register metrics: %s", name)
		return errors.New("cloud not register metrics")
	}
	metricMap[name] = true
	metricMapCount[name] = newMetrics
	metricMapCount[name].Inc()
	return nil
}
