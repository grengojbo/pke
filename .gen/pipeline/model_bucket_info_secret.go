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

type BucketInfoSecret struct {
	Id string `json:"id"`
	Name string `json:"name,omitempty"`
	// the secret identifier of the azure access information
	AccessId string `json:"accessId,omitempty"`
	// the secret name of the azure access information
	AccessName string `json:"accessName,omitempty"`
}
