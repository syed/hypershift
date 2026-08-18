package serviceutil

import corev1 "k8s.io/api/core/v1"

// ApplyTrafficDistribution sets the TrafficDistribution field on a Service
// spec when a non-empty value is provided. Use "PreferSameZone" to make
// kube-proxy prefer same-AZ endpoints, reducing cross-AZ data transfer.
func ApplyTrafficDistribution(svc *corev1.Service, trafficDistribution string) {
	if trafficDistribution != "" {
		svc.Spec.TrafficDistribution = &trafficDistribution
	}
}
