package metricsproxy

import (
	component "github.com/openshift/hypershift/support/controlplane-component"
	"github.com/openshift/hypershift/support/serviceutil"

	corev1 "k8s.io/api/core/v1"
)

func adaptService(cpContext component.WorkloadContext, svc *corev1.Service) error {
	serviceutil.ApplyTrafficDistribution(svc, cpContext.HCP.Spec.Networking.TrafficDistribution)
	return nil
}
