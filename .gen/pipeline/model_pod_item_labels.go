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

type PodItemLabels struct {
	App string `json:"app,omitempty"`
	Chart string `json:"chart,omitempty"`
	Release string `json:"release,omitempty"`
}
