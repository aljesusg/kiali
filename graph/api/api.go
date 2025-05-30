package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kiali/kiali/business"
	"github.com/kiali/kiali/config"
	"github.com/kiali/kiali/graph"
	config_common "github.com/kiali/kiali/graph/config/common"
	"github.com/kiali/kiali/graph/telemetry/istio"
	"github.com/kiali/kiali/log"
	"github.com/kiali/kiali/observability"
	"github.com/kiali/kiali/prometheus"
	"github.com/kiali/kiali/prometheus/internalmetrics"
)

// GraphNamespaces generates a namespaces graph using the provided options
func GraphNamespaces(ctx context.Context, business *business.Layer, o graph.Options) (code int, graphConfig interface{}) {
	var end observability.EndFunc
	ctx, end = observability.StartSpan(
		ctx,
		"GraphNamespaces",
		observability.Attribute("package", "api"),
	)
	defer end()
	// time how long it takes to generate this graph
	promtimer := internalmetrics.GetGraphGenerationTimePrometheusTimer(o.GetGraphKind(), o.TelemetryOptions.GraphType, o.InjectServiceNodes)
	defer promtimer.ObserveDuration()

	switch o.TelemetryVendor {
	case graph.VendorIstio:
		prom, err := prometheus.NewClient()
		graph.CheckError(err)
		code, graphConfig = graphNamespacesIstio(ctx, business, prom, o)
	default:
		graph.Error(fmt.Sprintf("TelemetryVendor [%s] not supported", o.TelemetryVendor))
	}

	// update metrics
	if _, ok := graphConfig.(config_common.Config); ok {
		numNodes := len(graphConfig.(config_common.Config).Elements.Nodes)
		internalmetrics.SetGraphNodes(o.GetGraphKind(), o.TelemetryOptions.GraphType, o.InjectServiceNodes, numNodes)
	}

	return code, graphConfig
}

// graphNamespacesIstio provides a test hook that accepts mock clients
func graphNamespacesIstio(ctx context.Context, business *business.Layer, prom *prometheus.Client, o graph.Options) (code int, graphConfig interface{}) {

	// Create a 'global' object to store the business. Global only to the request.
	globalInfo := graph.NewGlobalInfo(business, prom, config.Get())

	trafficMap := istio.BuildNamespacesTrafficMap(ctx, o.TelemetryOptions, globalInfo)
	code, graphConfig = generateGraph(ctx, trafficMap, o)

	return code, graphConfig
}

// GraphNode generates a node graph using the provided options
func GraphNode(ctx context.Context, business *business.Layer, o graph.Options) (code int, graphConfig interface{}) {
	if len(o.Namespaces) != 1 {
		graph.Error("Node graph does not support the 'namespaces' query parameter or the 'all' namespace")
	}

	// time how long it takes to generate this graph
	promtimer := internalmetrics.GetGraphGenerationTimePrometheusTimer(o.GetGraphKind(), o.TelemetryOptions.GraphType, o.InjectServiceNodes)
	defer promtimer.ObserveDuration()

	switch o.TelemetryVendor {
	case graph.VendorIstio:
		prom, err := prometheus.NewClient()
		graph.CheckError(err)
		code, graphConfig = graphNodeIstio(ctx, business, prom, o)
	default:
		graph.Error(fmt.Sprintf("TelemetryVendor [%s] not supported", o.TelemetryVendor))
	}

	// update metrics
	if _, ok := graphConfig.(config_common.Config); ok {
		numNodes := len(graphConfig.(config_common.Config).Elements.Nodes)
		internalmetrics.SetGraphNodes(o.GetGraphKind(), o.TelemetryOptions.GraphType, o.InjectServiceNodes, numNodes)
	}

	return code, graphConfig
}

// graphNodeIstio provides a test hook that accepts mock clients
func graphNodeIstio(ctx context.Context, business *business.Layer, prom *prometheus.Client, o graph.Options) (code int, graphConfig interface{}) {

	// Create a 'global' object to store the business. Global only to the request.
	globalInfo := graph.NewGlobalInfo(business, prom, config.Get())
	globalInfo.Business = business
	globalInfo.PromClient = prom

	trafficMap, _ := istio.BuildNodeTrafficMap(ctx, o.TelemetryOptions, globalInfo)
	code, graphConfig = generateGraph(ctx, trafficMap, o)

	return code, graphConfig
}

func generateGraph(ctx context.Context, trafficMap graph.TrafficMap, o graph.Options) (int, interface{}) {
	zl := log.FromContext(ctx)

	zl.Trace().Msgf("Generating config for [%s] graph...", o.ConfigVendor)

	promtimer := internalmetrics.GetGraphMarshalTimePrometheusTimer(o.GetGraphKind(), o.TelemetryOptions.GraphType, o.InjectServiceNodes)
	defer promtimer.ObserveDuration()

	var vendorConfig interface{}
	switch o.ConfigVendor {
	case graph.VendorCommon:
		vendorConfig = config_common.NewConfig(trafficMap, o.ConfigOptions)
	default:
		vendorConfig = config_common.NewConfig(trafficMap, o.ConfigOptions)
		zl.Debug().Msgf("ConfigVendor [%s] not supported, defaulting to [Common]", o.ConfigVendor)
	}

	zl.Trace().Msgf("Done generating config for [%s] graph", o.ConfigVendor)
	return http.StatusOK, vendorConfig
}
