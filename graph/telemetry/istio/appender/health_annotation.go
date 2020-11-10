package appender

import (
	"github.com/kiali/kiali/config"
	"github.com/kiali/kiali/graph"
	"github.com/kiali/kiali/models"
)

const AnnotationAppenderName = "healthAnnotation"

// Name: annotation
type AnnotationAppender struct{}

// Name implements Appender
func (a AnnotationAppender) Name() string {
	return AnnotationAppenderName
}

// AppendGraph implements Appender
func (a AnnotationAppender) AppendGraph(trafficMap graph.TrafficMap, globalInfo *graph.AppenderGlobalInfo, namespaceInfo *graph.AppenderNamespaceInfo) {
	if len(trafficMap) == 0 {
		return
	}

	if getWorkloadList(namespaceInfo) == nil {
		workloadList, err := globalInfo.Business.Workload.GetWorkloadList(namespaceInfo.Namespace)
		graph.CheckError(err)
		namespaceInfo.Vendor[workloadListKey] = &workloadList
	}

	a.applyAnnotationPresence(trafficMap, namespaceInfo)
}

func (a *AnnotationAppender) applyAnnotationPresence(trafficMap graph.TrafficMap, namespaceInfo *graph.AppenderNamespaceInfo) {
	for _, n := range trafficMap {
		// Skip the check if this node is outside the requested namespace, we limit badging to the requested namespaces
		if n.Namespace != namespaceInfo.Namespace {
			continue
		}

		// We whitelist istio components because they may not report telemetry using injected sidecars.
		if config.IsIstioNamespace(n.Namespace) {
			continue
		}

		// dead nodes tell no tales (er, have no pods)
		if isDead, ok := n.Metadata[graph.IsDead]; ok && isDead.(bool) {
			continue
		}

		// get the workloads for the node and check to see if they have sidecars. Note that
		// if there are no workloads/pods we don't flag it as missing sidecars.  No pods means
		// no missing sidecars.  (In most cases this means it was flagged as dead, and handled above)
		switch n.NodeType {
		case graph.NodeTypeWorkload:
			if workload, found := getWorkload(n.Workload, namespaceInfo); found {
				annotations, err := models.GetHealthConfiguration(workload.Annotations)
				if err != nil {
					break
				}
				n.Metadata[graph.Annotation] = annotations
			}
		case graph.NodeTypeService:
			if srv, found := getService(n.Service, namespaceInfo); found {
				annotations, err := models.GetHealthConfiguration(srv.Annotations)
				if err != nil {
					break
				}
				n.Metadata[graph.Annotation] = annotations
			}
		default:
			continue
		}
	}
}
