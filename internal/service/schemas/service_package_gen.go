// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package schemas

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDiscoverer,
			TypeName: "aws_schemas_discoverer",
		},
		{
			Factory:  ResourceRegistry,
			TypeName: "aws_schemas_registry",
		},
		{
			Factory:  ResourceRegistryPolicy,
			TypeName: "aws_schemas_registry_policy",
		},
		{
			Factory:  ResourceSchema,
			TypeName: "aws_schemas_schema",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Schemas
}

var ServicePackage = &servicePackage{}
