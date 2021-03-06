// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gateway_test

import (
	"io/ioutil"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/uber/zanzibar/runtime"
	"github.com/uber/zanzibar/test/lib"
	"github.com/uber/zanzibar/test/lib/bench_gateway"
	"github.com/uber/zanzibar/test/lib/test_gateway"
	"github.com/uber/zanzibar/test/lib/util"

	exampleGateway "github.com/uber/zanzibar/examples/example-gateway/build/services/example-gateway"
)

func TestHealthCall(t *testing.T) {
	gateway, err := testGateway.CreateGateway(t, nil, &testGateway.Options{
		TestBinary:  util.DefaultMainFile("example-gateway"),
		ConfigFiles: util.DefaultConfigFiles("example-gateway"),
	})
	if !assert.NoError(t, err, "must be able to create gateway") {
		return
	}
	defer gateway.Close()

	assert.NotNil(t, gateway, "gateway exists")

	res, err := gateway.MakeRequest("GET", "/health", nil, nil)
	if !assert.NoError(t, err, "got http error") {
		return
	}

	assert.Equal(t, res.Status, "200 OK", "got http 200")
}

func BenchmarkHealthCall(b *testing.B) {
	gateway, err := benchGateway.CreateGateway(
		map[string]interface{}{
			"clients.baz.serviceName": "baz",
		},
		&testGateway.Options{
			TestBinary:            util.DefaultMainFile("example-gateway"),
			ConfigFiles:           util.DefaultConfigFiles("example-gateway"),
			KnownHTTPBackends:     []string{"bar", "contacts", "google-now"},
			KnownTChannelBackends: []string{"baz"},
		},
		exampleGateway.CreateGateway,
	)

	if err != nil {
		b.Error("got bootstrap err: " + err.Error())
		return
	}

	b.ResetTimer()

	// b.SetParallelism(100)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			res, err := gateway.MakeRequest("GET", "/health", nil, nil)
			if err != nil {
				b.Error("got http error: " + err.Error())
				break
			}
			if res.Status != "200 OK" {
				b.Error("got bad status error: " + res.Status)
				break
			}

			_, err = ioutil.ReadAll(res.Body)
			if err != nil {
				b.Error("could not write response: " + res.Status)
				break
			}
		}
	})

	b.StopTimer()
	gateway.Close()
	b.StartTimer()
}

func TestHealthMetrics(t *testing.T) {
	gateway, err := testGateway.CreateGateway(t, nil, &testGateway.Options{
		CountMetrics: true,
		TestBinary:   util.DefaultMainFile("example-gateway"),
		ConfigFiles:  util.DefaultConfigFiles("example-gateway"),
	})
	if !assert.NoError(t, err, "must be able to create gateway") {
		return
	}
	defer gateway.Close()

	cgateway := gateway.(*testGateway.ChildProcessGateway)

	// Expect three metrics
	cgateway.MetricsWaitGroup.Add(3)

	res, err := gateway.MakeRequest("GET", "/health", nil, nil)
	if !assert.NoError(t, err, "got http error") {
		return
	}

	assert.Equal(t, res.Status, "200 OK", "got http 200")

	cgateway.MetricsWaitGroup.Wait()
	metrics := cgateway.M3Service.GetMetrics()
	sort.Sort(lib.SortMetricsByName(metrics))

	assert.Equal(t, 3, len(metrics), "expected 3 metrics")

	latencyMetric := metrics[0]

	assert.Equal(t,
		"test-gateway.production.all-workers.inbound.calls.latency",
		latencyMetric.GetName(),
		"expected correct name",
	)

	value := *latencyMetric.MetricValue.Timer.I64Value

	assert.True(t, value > 1000,
		"expected timer to be >1000 nano seconds",
	)
	assert.True(t, value < 1000*1000*1000,
		"expected timer to be <1 second",
	)

	tags := latencyMetric.GetTags()
	assert.Equal(t, 4, len(tags), "expected 4 tags")
	expectedTags := map[string]string{
		"endpoint": "health",
		"handler":  "health",
		"service":  "test-gateway",
		"env":      "test",
	}
	for tag := range tags {
		assert.Equal(t,
			expectedTags[tag.GetTagName()],
			tag.GetTagValue(),
			"expected tag value to be correct",
		)
	}

	recvdMetric := metrics[1]

	assert.Equal(t,
		"test-gateway.production.all-workers.inbound.calls.recvd",
		recvdMetric.GetName(),
		"expected correct name",
	)

	value = *recvdMetric.MetricValue.Count.I64Value
	assert.Equal(t,
		int64(1),
		value,
		"expected counter to be 1",
	)

	tags = recvdMetric.GetTags()
	assert.Equal(t, 4, len(tags), "expected 4 tags")
	expectedTags = map[string]string{
		"endpoint": "health",
		"handler":  "health",
		"service":  "test-gateway",
		"env":      "test",
	}
	for tag := range tags {
		assert.Equal(t,
			expectedTags[tag.GetTagName()],
			tag.GetTagValue(),
			"expected tag value to be correct",
		)
	}

	statusCodeMetric := metrics[2]

	assert.Equal(t,
		"test-gateway.production.all-workers.inbound.calls.status.200",
		statusCodeMetric.GetName(),
		"expected correct name",
	)

	value = *statusCodeMetric.MetricValue.Count.I64Value
	assert.Equal(t,
		int64(1),
		value,
		"expected counter to be 1",
	)

	tags = statusCodeMetric.GetTags()
	assert.Equal(t, 4, len(tags), "expected 4 tags")
	expectedTags = map[string]string{
		"endpoint": "health",
		"handler":  "health",
		"service":  "test-gateway",
		"env":      "test",
	}
	for tag := range tags {
		assert.Equal(t,
			expectedTags[tag.GetTagName()],
			tag.GetTagValue(),
			"expected tag value to be correct",
		)
	}
}

func TestRuntimeMetrics(t *testing.T) {
	gateway, err := testGateway.CreateGateway(t, nil, &testGateway.Options{
		CountMetrics:         true,
		EnableRuntimeMetrics: true,
		TestBinary:           util.DefaultMainFile("example-gateway"),
		ConfigFiles:          util.DefaultConfigFiles("example-gateway"),
	})
	if !assert.NoError(t, err, "must be able to create gateway") {
		return
	}
	defer gateway.Close()

	cgateway := gateway.(*testGateway.ChildProcessGateway)

	// Expect 30 runtime metrics
	cgateway.MetricsWaitGroup.Add(30)

	cgateway.MetricsWaitGroup.Wait()
	metrics := cgateway.M3Service.GetMetrics()
	sort.Sort(lib.SortMetricsByName(metrics))

	assert.Equal(t, 30, len(metrics), "expected 30 metrics")

	testData := []string{
		"test-gateway.production.per-worker.runtime.cpu.cgoCalls",
		"test-gateway.production.per-worker.runtime.cpu.count",
		"test-gateway.production.per-worker.runtime.cpu.goMaxProcs",
		"test-gateway.production.per-worker.runtime.cpu.goroutines",
		"test-gateway.production.per-worker.runtime.mem.alloc",
		"test-gateway.production.per-worker.runtime.mem.frees",
		"test-gateway.production.per-worker.runtime.mem.gc.count",
		"test-gateway.production.per-worker.runtime.mem.gc.cpuFraction",
		"test-gateway.production.per-worker.runtime.mem.gc.last",
		"test-gateway.production.per-worker.runtime.mem.gc.next",
		"test-gateway.production.per-worker.runtime.mem.gc.pause",
		"test-gateway.production.per-worker.runtime.mem.gc.pauseTotal",
		"test-gateway.production.per-worker.runtime.mem.gc.sys",
		"test-gateway.production.per-worker.runtime.mem.heap.alloc",
		"test-gateway.production.per-worker.runtime.mem.heap.idle",
		"test-gateway.production.per-worker.runtime.mem.heap.inuse",
		"test-gateway.production.per-worker.runtime.mem.heap.objects",
		"test-gateway.production.per-worker.runtime.mem.heap.released",
		"test-gateway.production.per-worker.runtime.mem.heap.sys",
		"test-gateway.production.per-worker.runtime.mem.lookups",
		"test-gateway.production.per-worker.runtime.mem.malloc",
		"test-gateway.production.per-worker.runtime.mem.otherSys",
		"test-gateway.production.per-worker.runtime.mem.stack.inuse",
		"test-gateway.production.per-worker.runtime.mem.stack.mcacheInuse",
		"test-gateway.production.per-worker.runtime.mem.stack.mcacheSys",
		"test-gateway.production.per-worker.runtime.mem.stack.mspanInuse",
		"test-gateway.production.per-worker.runtime.mem.stack.mspanSys",
		"test-gateway.production.per-worker.runtime.mem.stack.sys",
		"test-gateway.production.per-worker.runtime.mem.sys",
		"test-gateway.production.per-worker.runtime.mem.total",
	}

	for i, m := range metrics {
		assert.Equal(t, testData[i], m.Name, "expected correct name")

		tags := m.Tags
		assert.Equal(t, 3, len(tags), "expected 3 tags")
		expectedTags := map[string]string{
			"env":     "test",
			"service": "test-gateway",
			"host":    zanzibar.GetHostname(),
		}
		for tag := range tags {
			assert.Equal(t, expectedTags[tag.GetTagName()], tag.GetTagValue(), "expected tag value to be correct")
		}
	}
}
