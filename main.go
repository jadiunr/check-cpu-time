package main

import (
	"fmt"
	"reflect"
	"strings"

	corev2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-plugin-sdk/sensu"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/iancoleman/strcase"
)

// Config represents the check plugin config.
type Config struct {
	sensu.PluginConfig
	Critical float64
	Warning float64
	Format string
}

type MetricGroup struct {
	Comment string
	Type string
	Name string
	Metrics []Metric
}

type Metric struct {
	Tags map[string]string
	Value float64
}

func (g *MetricGroup) Output() {
	output := []string{}

	output = append(output, fmt.Sprintf("# HELP %s %s", g.Name, g.Comment))
	output = append(output, fmt.Sprintf("# TYPE %s %s", g.Name, g.Type))

	for _, metric := range g.Metrics {
		tags := []string{}
		for k, v := range metric.Tags {
			tags = append(tags, fmt.Sprintf("%s=\"%s\"", k, v))
		}
		output = append(output, fmt.Sprintf("%s{%s} %.2f", g.Name, strings.Join(tags, ","), metric.Value))
	}

	fmt.Println(strings.Join(output, "\n"))
}

var (
	plugin = Config{
		PluginConfig: sensu.PluginConfig{
			Name:     "check-cpu-time",
			Short:    "Check CPU time",
			Keyspace: "sensu.io/plugins/check-cpu-usage-extended/config",
		},
	}

	options = []sensu.ConfigOption{}
)

func main() {
	check := sensu.NewGoCheck(&plugin.PluginConfig, options, checkArgs, executeCheck, false)
	check.Execute()
}

func checkArgs(event *corev2.Event) (int, error) {
	return sensu.CheckStateOK, nil
}

func executeCheck(event *corev2.Event) (int, error) {
	info, err := cpu.Info()
	if err != nil {
		return sensu.CheckStateCritical, err
	} else if len(info) < 1 {
		return sensu.CheckStateCritical, err
	}
	cores := float64(0)
	for _, i := range info {
		cores += float64(i.Cores)
	}

	timesStats, err := cpu.Times(true)
	if err != nil {
		return sensu.CheckStateCritical, fmt.Errorf("failed to get cpu statistics: %v", err)
	}
	totalTimesStat, err := cpu.Times(false)
	if err != nil {
		return sensu.CheckStateCritical, fmt.Errorf("failed to get cpu statistics: %v", err)
	}
	timesStats = append(timesStats, totalTimesStat...)

	vTimesStats := reflect.ValueOf(timesStats)

	{
		measurement := "Total"
		metricGroup := &MetricGroup{
			Name: fmt.Sprintf("cpu_time_%s", strcase.ToSnake(measurement)),
			Type: "counter",
			Comment: fmt.Sprintf("Static CPU %s time", measurement),
			Metrics: make([]Metric, 0),
		}
		for cpuIdx := 0; cpuIdx < vTimesStats.Len(); cpuIdx++ {
			totalTime := float64(0)
			for fieldIdx := 0; fieldIdx < vTimesStats.Index(0).NumField(); fieldIdx++ {
				if vTimesStats.Index(0).Field(fieldIdx).Type().Name() != "float64" {
					continue
				}
				totalTime += vTimesStats.Index(cpuIdx).Field(fieldIdx).Float()
			}
			metricGroup.Metrics = append(metricGroup.Metrics, Metric{
				Tags: map[string]string{ "cpu": vTimesStats.Index(cpuIdx).FieldByName("CPU").String() },
				Value: totalTime,
			})
		}
		metricGroup.Output()
	}

	for fieldIdx := 0; fieldIdx < vTimesStats.Index(0).NumField(); fieldIdx++ {
		if vTimesStats.Index(0).Field(fieldIdx).Type().Name() != "float64" {
			continue
		}
		measurement := vTimesStats.Index(0).Type().Field(fieldIdx).Name
		metricGroup := &MetricGroup{
			Name: fmt.Sprintf("cpu_time_%s", strcase.ToSnake(measurement)),
			Type: "counter",
			Comment: fmt.Sprintf("Statistic CPU %s time", measurement),
			Metrics: make([]Metric, 0),

		}
		for cpuIdx := 0; cpuIdx < vTimesStats.Len(); cpuIdx++ {
			metricGroup.Metrics = append(metricGroup.Metrics, Metric{
				Tags: map[string]string{ "cpu": vTimesStats.Index(cpuIdx).FieldByName("CPU").String() },
				Value: vTimesStats.Index(cpuIdx).Field(fieldIdx).Float(),
			})
		}
		metricGroup.Output()
	}

	return sensu.CheckStateOK, nil
}
