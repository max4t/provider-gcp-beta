/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"github.com/upbound/upjet/pkg/terraform"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/max4t/provider-gcp-beta/apis/v1beta1"
)

const (
	keyProject = "project"

	keyCredentials = "credentials"
)

const (
	// error messages
	errNoProviderConfig   = "no providerConfigRef provided"
	errGetProviderConfig  = "cannot get referenced ProviderConfig"
	errTrackUsage         = "cannot track ProviderConfig usage"
	errExtractCredentials = "cannot extract credentials"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		// set provider configuration
		ps.Configuration = map[string]interface{}{
			keyProject: pc.Spec.ProjectID,
		}

		switch pc.Spec.Credentials.Source { //nolint:exhaustive
		case xpv1.CredentialsSourceInjectedIdentity:
			// We don't need to do anything here, as the TF Provider will take care of workloadIdentity etc.
		default:
			data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
			if err != nil {
				return ps, errors.Wrap(err, errExtractCredentials)
			}

			// set provider configuration keys for GCP credentials
			ps.Configuration[keyCredentials] = string(data)
		}
		return ps, nil
	}
}
