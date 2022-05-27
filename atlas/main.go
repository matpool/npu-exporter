package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	namespace = "atlas"
	port      = ":9403"
)

type Exporter struct {
	up           prometheus.Gauge
	info         *prometheus.GaugeVec
	deviceCount  prometheus.Gauge
	temperatures *prometheus.GaugeVec
	deviceInfo   *prometheus.GaugeVec
	powerUsage   *prometheus.GaugeVec
	memoryTotal  *prometheus.GaugeVec
	memoryUsed   *prometheus.GaugeVec
	aiCore       *prometheus.GaugeVec
}

func main() {
	var (
		listenAddress = flag.String("web.listen-address", port, "Address to listen on for web interface and telemetry.")
		metricsPath   = flag.String("web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	)
	flag.Parse()

	prometheus.MustRegister(NewExporter())

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>NPU Exporter</title></head>
             <body>
             <h1>NPU Exporter</h1>
             <p><a href='` + *metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})
	fmt.Println("Starting HTTP server on", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}

func NewExporter() *Exporter {
	return &Exporter{
		up: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "up",
				Help:      "NPU Metric Collection Operational",
			},
		),
		info: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "driver_info",
				Help:      "NPU Info",
			},
			[]string{"version"},
		),
		deviceCount: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "device_count",
				Help:      "Count of found NPU devices",
			},
		),
		deviceInfo: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "info",
				Help:      "Info as reported by the device",
			},
			[]string{"deviceid", "chipid", "npuid", "name"},
		),
		temperatures: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "temperatures",
				Help:      "Temperature as reported by the device",
			},
			[]string{"deviceid"},
		),
		powerUsage: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "power_usage",
				Help:      "Power usage as reported by the device",
			},
			[]string{"deviceid"},
		),
		memoryTotal: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "memory_total",
				Help:      "Total memory as reported by the device",
			},
			[]string{"deviceid"},
		),
		memoryUsed: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "memory_used",
				Help:      "Used memory as reported by the device",
			},
			[]string{"deviceid"},
		),
		aiCore: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: namespace,
				Name:      "ai_core",
				Help:      "AICore as reported by the device",
			},
			[]string{"deviceid"},
		),
	}
}

func (e *Exporter) Collect(metrics chan<- prometheus.Metric) {
	data, err := collectMetrics()
	if err != nil {
		log.Printf("Failed to collect metrics: %s\n", err)
		e.up.Set(0)
		e.up.Collect(metrics)
		return
	}

	e.up.Set(1)
	e.info.WithLabelValues(data.Version).Set(1)
	e.deviceCount.Set(float64(len(data.Devices)))

	for i := 0; i < len(data.Devices); i++ {
		d := data.Devices[i]
		memUsage, memTotal := parseMemoryByte(d.MemoryUsageMB)
		e.deviceInfo.WithLabelValues(d.Device, d.ChipID, d.NpuID, d.Name).Set(1)
		e.memoryTotal.WithLabelValues(d.Device).Set(memTotal)
		e.memoryUsed.WithLabelValues(d.Device).Set(memUsage)
		e.powerUsage.WithLabelValues(d.Device).Set(parseValueFloat(d.PowerUsage))
		e.temperatures.WithLabelValues(d.Device).Set(parseValueFloat(d.Temperature))
		e.aiCore.WithLabelValues(d.Device).Set(parseValueFloat(d.AICore))
	}

	e.deviceCount.Collect(metrics)
	e.deviceInfo.Collect(metrics)
	e.info.Collect(metrics)
	e.memoryTotal.Collect(metrics)
	e.memoryUsed.Collect(metrics)
	e.powerUsage.Collect(metrics)
	e.temperatures.Collect(metrics)
	e.aiCore.Collect(metrics)
	e.up.Collect(metrics)
}

func (e *Exporter) Describe(descs chan<- *prometheus.Desc) {
	e.deviceCount.Describe(descs)
	e.deviceInfo.Describe(descs)
	e.info.Describe(descs)
	e.memoryTotal.Describe(descs)
	e.memoryUsed.Describe(descs)
	e.powerUsage.Describe(descs)
	e.temperatures.Describe(descs)
	e.up.Describe(descs)
	e.aiCore.Describe(descs)
}
