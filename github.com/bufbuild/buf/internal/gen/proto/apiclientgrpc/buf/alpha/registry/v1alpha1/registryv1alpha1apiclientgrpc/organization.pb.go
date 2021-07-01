// Copyright 2020-2021 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-apiclientgrpc. DO NOT EDIT.

package registryv1alpha1apiclientgrpc

import (
	context "context"
	v1alpha1 "github.com/bufbuild/buf/internal/gen/proto/go/buf/alpha/registry/v1alpha1"
	zap "go.uber.org/zap"
)

type organizationService struct {
	logger          *zap.Logger
	client          v1alpha1.OrganizationServiceClient
	contextModifier func(context.Context) context.Context
}

// GetOrganization gets a organization by ID.
func (s *organizationService) GetOrganization(ctx context.Context, id string) (organization *v1alpha1.Organization, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.GetOrganization(
		ctx,
		&v1alpha1.GetOrganizationRequest{
			Id: id,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Organization, nil
}

// GetOrganizationByName gets a organization by name.
func (s *organizationService) GetOrganizationByName(ctx context.Context, name string) (organization *v1alpha1.Organization, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.GetOrganizationByName(
		ctx,
		&v1alpha1.GetOrganizationByNameRequest{
			Name: name,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Organization, nil
}

// ListOrganizations lists all organizations.
func (s *organizationService) ListOrganizations(
	ctx context.Context,
	pageSize uint32,
	pageToken string,
	reverse bool,
) (organizations []*v1alpha1.Organization, nextPageToken string, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.ListOrganizations(
		ctx,
		&v1alpha1.ListOrganizationsRequest{
			PageSize:  pageSize,
			PageToken: pageToken,
			Reverse:   reverse,
		},
	)
	if err != nil {
		return nil, "", err
	}
	return response.Organizations, response.NextPageToken, nil
}

// ListUserOrganizations lists all organizations a user is member of.
func (s *organizationService) ListUserOrganizations(
	ctx context.Context,
	userId string,
	pageSize uint32,
	pageToken string,
	reverse bool,
) (organizations []*v1alpha1.Organization, nextPageToken string, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.ListUserOrganizations(
		ctx,
		&v1alpha1.ListUserOrganizationsRequest{
			UserId:    userId,
			PageSize:  pageSize,
			PageToken: pageToken,
			Reverse:   reverse,
		},
	)
	if err != nil {
		return nil, "", err
	}
	return response.Organizations, response.NextPageToken, nil
}

// CreateOrganization creates a new organization.
func (s *organizationService) CreateOrganization(ctx context.Context, name string) (organization *v1alpha1.Organization, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.CreateOrganization(
		ctx,
		&v1alpha1.CreateOrganizationRequest{
			Name: name,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Organization, nil
}

// UpdateOrganizationName updates a organization's name.
func (s *organizationService) UpdateOrganizationName(
	ctx context.Context,
	id string,
	newName string,
) (organization *v1alpha1.Organization, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UpdateOrganizationName(
		ctx,
		&v1alpha1.UpdateOrganizationNameRequest{
			Id:      id,
			NewName: newName,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Organization, nil
}

// UpdateOrganizationNameByName updates a organization's name by name.
func (s *organizationService) UpdateOrganizationNameByName(
	ctx context.Context,
	name string,
	newName string,
) (organization *v1alpha1.Organization, _ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	response, err := s.client.UpdateOrganizationNameByName(
		ctx,
		&v1alpha1.UpdateOrganizationNameByNameRequest{
			Name:    name,
			NewName: newName,
		},
	)
	if err != nil {
		return nil, err
	}
	return response.Organization, nil
}

// DeleteOrganization deletes a organization.
func (s *organizationService) DeleteOrganization(ctx context.Context, id string) (_ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	_, err := s.client.DeleteOrganization(
		ctx,
		&v1alpha1.DeleteOrganizationRequest{
			Id: id,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// DeleteOrganizationByName deletes a organization by name.
func (s *organizationService) DeleteOrganizationByName(ctx context.Context, name string) (_ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	_, err := s.client.DeleteOrganizationByName(
		ctx,
		&v1alpha1.DeleteOrganizationByNameRequest{
			Name: name,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// AddOrganizationBaseRepositoryScope adds a base repository scope to an organization by ID.
func (s *organizationService) AddOrganizationBaseRepositoryScope(
	ctx context.Context,
	id string,
	repositoryScope v1alpha1.RepositoryScope,
) (_ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	_, err := s.client.AddOrganizationBaseRepositoryScope(
		ctx,
		&v1alpha1.AddOrganizationBaseRepositoryScopeRequest{
			Id:              id,
			RepositoryScope: repositoryScope,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// AddOrganizationBaseRepositoryScopeByName adds a base repository scope to an organization by name.
func (s *organizationService) AddOrganizationBaseRepositoryScopeByName(
	ctx context.Context,
	name string,
	repositoryScope v1alpha1.RepositoryScope,
) (_ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	_, err := s.client.AddOrganizationBaseRepositoryScopeByName(
		ctx,
		&v1alpha1.AddOrganizationBaseRepositoryScopeByNameRequest{
			Name:            name,
			RepositoryScope: repositoryScope,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// RemoveOrganizationBaseRepositoryScope removes a base repository scope from an organization by ID.
func (s *organizationService) RemoveOrganizationBaseRepositoryScope(
	ctx context.Context,
	id string,
	repositoryScope v1alpha1.RepositoryScope,
) (_ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	_, err := s.client.RemoveOrganizationBaseRepositoryScope(
		ctx,
		&v1alpha1.RemoveOrganizationBaseRepositoryScopeRequest{
			Id:              id,
			RepositoryScope: repositoryScope,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// RemoveOrganizationBaseRepositoryScopeByName removes a base repository scope from an organization by name.
func (s *organizationService) RemoveOrganizationBaseRepositoryScopeByName(
	ctx context.Context,
	name string,
	repositoryScope v1alpha1.RepositoryScope,
) (_ error) {
	if s.contextModifier != nil {
		ctx = s.contextModifier(ctx)
	}
	_, err := s.client.RemoveOrganizationBaseRepositoryScopeByName(
		ctx,
		&v1alpha1.RemoveOrganizationBaseRepositoryScopeByNameRequest{
			Name:            name,
			RepositoryScope: repositoryScope,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
