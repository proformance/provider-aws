/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta1 "github.com/upbound/provider-aws/apis/budgets/v1beta1"
	v1beta11 "github.com/upbound/provider-aws/apis/iam/v1beta1"
	common "github.com/upbound/provider-aws/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this BudgetResourceAssociation.
func (mg *BudgetResourceAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BudgetName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.BudgetNameRef,
		Selector:     mg.Spec.ForProvider.BudgetNameSelector,
		To: reference.To{
			List:    &v1beta1.BudgetList{},
			Managed: &v1beta1.Budget{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.BudgetName")
	}
	mg.Spec.ForProvider.BudgetName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BudgetNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceIDRef,
		Selector:     mg.Spec.ForProvider.ResourceIDSelector,
		To: reference.To{
			List:    &ProductList{},
			Managed: &Product{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceID")
	}
	mg.Spec.ForProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Constraint.
func (mg *Constraint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PortfolioID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PortfolioIDRef,
		Selector:     mg.Spec.ForProvider.PortfolioIDSelector,
		To: reference.To{
			List:    &PortfolioList{},
			Managed: &Portfolio{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortfolioID")
	}
	mg.Spec.ForProvider.PortfolioID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PortfolioIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProductID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ProductIDRef,
		Selector:     mg.Spec.ForProvider.ProductIDSelector,
		To: reference.To{
			List:    &ProductList{},
			Managed: &Product{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProductID")
	}
	mg.Spec.ForProvider.ProductID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProductIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PortfolioShare.
func (mg *PortfolioShare) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PortfolioID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.PortfolioIDRef,
		Selector:     mg.Spec.ForProvider.PortfolioIDSelector,
		To: reference.To{
			List:    &PortfolioList{},
			Managed: &Portfolio{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortfolioID")
	}
	mg.Spec.ForProvider.PortfolioID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PortfolioIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this PrincipalPortfolioAssociation.
func (mg *PrincipalPortfolioAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PortfolioID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PortfolioIDRef,
		Selector:     mg.Spec.ForProvider.PortfolioIDSelector,
		To: reference.To{
			List:    &PortfolioList{},
			Managed: &Portfolio{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortfolioID")
	}
	mg.Spec.ForProvider.PortfolioID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PortfolioIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PrincipalArn),
		Extract:      common.ARNExtractor(),
		Reference:    mg.Spec.ForProvider.PrincipalArnRef,
		Selector:     mg.Spec.ForProvider.PrincipalArnSelector,
		To: reference.To{
			List:    &v1beta11.UserList{},
			Managed: &v1beta11.User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PrincipalArn")
	}
	mg.Spec.ForProvider.PrincipalArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PrincipalArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ProductPortfolioAssociation.
func (mg *ProductPortfolioAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PortfolioID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PortfolioIDRef,
		Selector:     mg.Spec.ForProvider.PortfolioIDSelector,
		To: reference.To{
			List:    &PortfolioList{},
			Managed: &Portfolio{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PortfolioID")
	}
	mg.Spec.ForProvider.PortfolioID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PortfolioIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProductID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProductIDRef,
		Selector:     mg.Spec.ForProvider.ProductIDSelector,
		To: reference.To{
			List:    &ProductList{},
			Managed: &Product{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProductID")
	}
	mg.Spec.ForProvider.ProductID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProductIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ProvisionedProduct.
func (mg *ProvisionedProduct) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProductName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProductNameRef,
		Selector:     mg.Spec.ForProvider.ProductNameSelector,
		To: reference.To{
			List:    &ProductList{},
			Managed: &Product{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProductName")
	}
	mg.Spec.ForProvider.ProductName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProductNameRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProvisioningArtifactName),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ProvisioningArtifactNameRef,
		Selector:     mg.Spec.ForProvider.ProvisioningArtifactNameSelector,
		To: reference.To{
			List:    &ProvisioningArtifactList{},
			Managed: &ProvisioningArtifact{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProvisioningArtifactName")
	}
	mg.Spec.ForProvider.ProvisioningArtifactName = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProvisioningArtifactNameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ProvisioningArtifact.
func (mg *ProvisioningArtifact) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ProductID),
		Extract:      resource.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ProductIDRef,
		Selector:     mg.Spec.ForProvider.ProductIDSelector,
		To: reference.To{
			List:    &ProductList{},
			Managed: &Product{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ProductID")
	}
	mg.Spec.ForProvider.ProductID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ProductIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TagOptionResourceAssociation.
func (mg *TagOptionResourceAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.ResourceID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.ResourceIDRef,
		Selector:     mg.Spec.ForProvider.ResourceIDSelector,
		To: reference.To{
			List:    &ProductList{},
			Managed: &Product{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.ResourceID")
	}
	mg.Spec.ForProvider.ResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ResourceIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TagOptionID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TagOptionIDRef,
		Selector:     mg.Spec.ForProvider.TagOptionIDSelector,
		To: reference.To{
			List:    &TagOptionList{},
			Managed: &TagOption{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TagOptionID")
	}
	mg.Spec.ForProvider.TagOptionID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TagOptionIDRef = rsp.ResolvedReference

	return nil
}
