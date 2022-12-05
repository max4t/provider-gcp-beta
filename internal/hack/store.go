package hack

import (
	"context"

	"github.com/upbound/upjet/pkg/config"
	tjcontroller "github.com/upbound/upjet/pkg/controller"
	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/terraform"
)

var _ tjcontroller.Store = (*Store)(nil)
var _ terraform.StoreCleaner = (*Store)(nil)

type Store struct {
	*terraform.WorkspaceStore
}

func (s *Store) Workspace(ctx context.Context, c resource.SecretClient, tr resource.Terraformed, ts terraform.Setup, cfg *config.Resource) (*terraform.Workspace, error) {
	return s.WorkspaceStore.Workspace(ctx, c, &aliasedTerraformed{tr}, ts, cfg)
}

type aliasedTerraformed struct {
	resource.Terraformed
}

func (t *aliasedTerraformed) GetParameters() (map[string]any, error) {
	m, err := t.Terraformed.GetParameters()
	if err != nil {
		return m, err
	}

	m["provider"] = "google-beta"

	return m, nil
}
