/*
 * Pipeline API
 *
 * Pipeline v0.3.0 swagger
 *
 * API version: master
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipeline

type NodeItemMetadataLabels struct {
	BetaKubernetesIoarch string `json:"beta.kubernetes.io/arch,omitempty"`
	BetaKubernetesIofluentdDsReady string `json:"beta.kubernetes.io/fluentd-ds-ready,omitempty"`
	BetaKubernetesIoinstanceType string `json:"beta.kubernetes.io/instance-type,omitempty"`
	BetaKubernetesIoos string `json:"beta.kubernetes.io/os,omitempty"`
	CloudGoogleComgkeNodepool string `json:"cloud.google.com/gke-nodepool,omitempty"`
	FailureDomainBetaKubernetesIoregion string `json:"failure-domain.beta.kubernetes.io/region,omitempty"`
	FailureDomainBetaKubernetesIozone string `json:"failure-domain.beta.kubernetes.io/zone,omitempty"`
	KubernetesIohostname string `json:"kubernetes.io/hostname,omitempty"`
}
