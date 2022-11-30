/*
Copyright 2022 Max T.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this RegionTargetTCPProxy
func (mg *RegionTargetTCPProxy) GetTerraformResourceType() string {
	return "google_compute_region_target_tcp_proxy"
}

// GetConnectionDetailsMapping for this RegionTargetTCPProxy
func (tr *RegionTargetTCPProxy) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this RegionTargetTCPProxy
func (tr *RegionTargetTCPProxy) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this RegionTargetTCPProxy
func (tr *RegionTargetTCPProxy) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this RegionTargetTCPProxy
func (tr *RegionTargetTCPProxy) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this RegionTargetTCPProxy
func (tr *RegionTargetTCPProxy) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this RegionTargetTCPProxy
func (tr *RegionTargetTCPProxy) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this RegionTargetTCPProxy using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *RegionTargetTCPProxy) LateInitialize(attrs []byte) (bool, error) {
	params := &RegionTargetTCPProxyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *RegionTargetTCPProxy) GetTerraformSchemaVersion() int {
	return 0
}