// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/model/pdata"
)

// Type is the component type name.
const Type config.Type = "disk"

// MetricIntf is an interface to generically interact with generated metric.
type MetricIntf interface {
	Name() string
	New() pdata.Metric
	Init(metric pdata.Metric)
}

// Intentionally not exposing this so that it is opaque and can change freely.
type metricImpl struct {
	name     string
	initFunc func(pdata.Metric)
}

// Name returns the metric name.
func (m *metricImpl) Name() string {
	return m.name
}

// New creates a metric object preinitialized.
func (m *metricImpl) New() pdata.Metric {
	metric := pdata.NewMetric()
	m.Init(metric)
	return metric
}

// Init initializes the provided metric object.
func (m *metricImpl) Init(metric pdata.Metric) {
	m.initFunc(metric)
}

type metricStruct struct {
	SystemDiskIo                MetricIntf
	SystemDiskIoTime            MetricIntf
	SystemDiskMerged            MetricIntf
	SystemDiskOperationTime     MetricIntf
	SystemDiskOperations        MetricIntf
	SystemDiskPendingOperations MetricIntf
	SystemDiskWeightedIoTime    MetricIntf
}

// Names returns a list of all the metric name strings.
func (m *metricStruct) Names() []string {
	return []string{
		"system.disk.io",
		"system.disk.io_time",
		"system.disk.merged",
		"system.disk.operation_time",
		"system.disk.operations",
		"system.disk.pending_operations",
		"system.disk.weighted_io_time",
	}
}

var metricsByName = map[string]MetricIntf{
	"system.disk.io":                 Metrics.SystemDiskIo,
	"system.disk.io_time":            Metrics.SystemDiskIoTime,
	"system.disk.merged":             Metrics.SystemDiskMerged,
	"system.disk.operation_time":     Metrics.SystemDiskOperationTime,
	"system.disk.operations":         Metrics.SystemDiskOperations,
	"system.disk.pending_operations": Metrics.SystemDiskPendingOperations,
	"system.disk.weighted_io_time":   Metrics.SystemDiskWeightedIoTime,
}

func (m *metricStruct) ByName(n string) MetricIntf {
	return metricsByName[n]
}

// Metrics contains a set of methods for each metric that help with
// manipulating those metrics.
var Metrics = &metricStruct{
	&metricImpl{
		"system.disk.io",
		func(metric pdata.Metric) {
			metric.SetName("system.disk.io")
			metric.SetDescription("Disk bytes transferred.")
			metric.SetUnit("By")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.disk.io_time",
		func(metric pdata.Metric) {
			metric.SetName("system.disk.io_time")
			metric.SetDescription("Time disk spent activated. On Windows, this is calculated as the inverse of disk idle time.")
			metric.SetUnit("s")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.disk.merged",
		func(metric pdata.Metric) {
			metric.SetName("system.disk.merged")
			metric.SetDescription("The number of disk reads merged into single physical disk access operations.")
			metric.SetUnit("{operations}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.disk.operation_time",
		func(metric pdata.Metric) {
			metric.SetName("system.disk.operation_time")
			metric.SetDescription("Time spent in disk operations.")
			metric.SetUnit("s")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.disk.operations",
		func(metric pdata.Metric) {
			metric.SetName("system.disk.operations")
			metric.SetDescription("Disk operations count.")
			metric.SetUnit("{operations}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.disk.pending_operations",
		func(metric pdata.Metric) {
			metric.SetName("system.disk.pending_operations")
			metric.SetDescription("The queue size of pending I/O operations.")
			metric.SetUnit("{operations}")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(false)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
	&metricImpl{
		"system.disk.weighted_io_time",
		func(metric pdata.Metric) {
			metric.SetName("system.disk.weighted_io_time")
			metric.SetDescription("Time disk spent activated multiplied by the queue length.")
			metric.SetUnit("s")
			metric.SetDataType(pdata.MetricDataTypeSum)
			metric.Sum().SetIsMonotonic(true)
			metric.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
		},
	},
}

// M contains a set of methods for each metric that help with
// manipulating those metrics. M is an alias for Metrics
var M = Metrics

// Labels contains the possible metric labels that can be used.
var Labels = struct {
	// Device (Name of the disk.)
	Device string
	// Direction (Direction of flow of bytes/opertations (read or write).)
	Direction string
}{
	"device",
	"direction",
}

// L contains the possible metric labels that can be used. L is an alias for
// Labels.
var L = Labels

// LabelDirection are the possible values that the label "direction" can have.
var LabelDirection = struct {
	Read  string
	Write string
}{
	"read",
	"write",
}
