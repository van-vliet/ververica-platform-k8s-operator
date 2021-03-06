/*
 * Application Manager API
 *
 * Application Manager APIs to control Apache Flink jobs
 *
 * API version: 2.1.0
 * Contact: platform@ververica.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package appmanagerapi

type JobStatus struct {
	Failure *Failure          `json:"failure,omitempty"`
	Started *JobStatusStarted `json:"started,omitempty"`
	State   string            `json:"state,omitempty"`
}
