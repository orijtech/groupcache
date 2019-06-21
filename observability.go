/*
Copyright 2018 Google LLC.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package galaxycache

import (
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

const (
	unitDimensionless = "1"
	unitBytes         = "By"
	unitMillisecond   = "ms"
)

var (
	// Copied from https://github.com/census-instrumentation/opencensus-go/blob/ff7de98412e5c010eb978f11056f90c00561637f/plugin/ocgrpc/stats_common.go#L54
	defaultBytesDistribution = view.Distribution(0, 1024, 2048, 4096, 16384, 65536, 262144, 1048576, 4194304, 16777216, 67108864, 268435456, 1073741824, 4294967296)
	// Copied from https://github.com/census-instrumentation/opencensus-go/blob/ff7de98412e5c010eb978f11056f90c00561637f/plugin/ocgrpc/stats_common.go#L55
	defaultMillisecondsDistribution = view.Distribution(0, 0.01, 0.05, 0.1, 0.3, 0.6, 0.8, 1, 2, 3, 4, 5, 6, 8, 10, 13, 16, 20, 25, 30, 40, 50, 65, 80, 100, 130, 160, 200, 250, 300, 400, 500, 650, 800, 1000, 2000, 5000, 10000, 20000, 50000, 100000)
)

var (
	MGets            = stats.Int64("gets", "The number of Get requests", unitDimensionless)
	MCacheHits       = stats.Int64("cache_hits", "The number of times that either cache was good", unitDimensionless)
	MCacheMisses     = stats.Int64("cache_misses", "The number of times that either cache was not good", unitDimensionless)
	MStarAuthorityLoads       = stats.Int64("star_loads", "The number of remote loads or remote cache hits", unitDimensionless)
	MStarAuthorityErrors      = stats.Int64("star_errors", "The number of remote errors", unitDimensionless)
	MLoads           = stats.Int64("loads", "The number of gets/cacheHits", unitDimensionless)
	MLoadErrors      = stats.Int64("loads_errors", "The number of errors encountered during Get", unitDimensionless)
	MLoadsDeduped    = stats.Int64("loads_deduped", "The number of loads after singleflight", unitDimensionless)
	MLocalLoads      = stats.Int64("local_loads", "The number of good local loads", unitDimensionless)
	MLocalLoadErrors = stats.Int64("local_load_errors", "The number of bad local loads", unitDimensionless)
	MServerRequests  = stats.Int64("server_requests", "The number of Gets that came over the network from starAuthorities", unitDimensionless)
	MKeyLength       = stats.Int64("key_length", "The length of keys", unitBytes)
	MValueLength     = stats.Int64("value_length", "The length of values", unitBytes)

	MRoundtripLatencyMilliseconds = stats.Float64("roundtrip_latency", "Roundtrip latency in milliseconds", unitMillisecond)
)

var keyCommand, _ = tag.NewKey("command")

var AllViews = []*view.View{
	{Name: "galaxycache/gets", Description: "The number of Get requests", Measure: MGets, Aggregation: view.Count()},
	{Name: "galaxycache/cache_hits", Description: "The number of times that either cache was good", Measure: MCacheHits, Aggregation: view.Count()},
	{Name: "galaxycache/cache_misses", Description: "The number of times that either cache was not good", Measure: MCacheMisses, Aggregation: view.Count()},
	{Name: "galaxycache/star_loads", Description: "The number of remote loads or remote cache hits", Measure: MStarAuthorityLoads, Aggregation: view.Count()},
	{Name: "galaxycache/star_errors", Description: "The number of remote errors", Measure: MStarAuthorityErrors, Aggregation: view.Count()},
	{Name: "galaxycache/loads", Description: "The number of loads after singleflight", Measure: MLoads, Aggregation: view.Count()},
	{Name: "galaxycache/loads_deduped", Description: "The number of loads after singleflight", Measure: MLoadsDeduped, Aggregation: view.Count()},
	{Name: "galaxycache/local_loads", Description: "The number of good local loads", Measure: MLocalLoads, Aggregation: view.Count()},
	{Name: "galaxycache/local_load_errors", Description: "The number of bad local loads", Measure: MLocalLoadErrors, Aggregation: view.Count()},
	{Name: "galaxycache/server_requests", Description: "The number of Gets that came over the network from starAuthorities", Measure: MServerRequests, Aggregation: view.Count()},
	{Name: "galaxycache/key_length", Description: "The distribution of the key lengths", Measure: MKeyLength, Aggregation: defaultBytesDistribution},
	{Name: "galaxycache/value_length", Description: "The distribution of the value lengths", Measure: MValueLength, Aggregation: defaultBytesDistribution},
	{Name: "galaxycache/roundtrip_latency", Description: "The roundtrip latency", Measure: MRoundtripLatencyMilliseconds, Aggregation: defaultMillisecondsDistribution},
}

func sinceInMilliseconds(start time.Time) float64 {
	d := time.Since(start)
	return float64(d.Nanoseconds()) / 1e6
}
