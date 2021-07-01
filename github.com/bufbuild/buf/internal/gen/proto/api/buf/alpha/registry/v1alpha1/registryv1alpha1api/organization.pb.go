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

// Code generated by protoc-gen-go-api. DO NOT EDIT.

package registryv1alpha1api

import (
	context "context"
	v1alpha1 "github.com/bufbuild/buf/internal/gen/proto/go/buf/alpha/registry/v1alpha1"
)

// OrganizationService is the Organization service.
type OrganizationService interface {
	// GetOrganization gets a organization by ID.
	GetOrganization(ctx context.Context, id string) (organization *v1alpha1.Organization, err error)
	// GetOrganizationByName gets a organization by name.
	GetOrganizationByName(ctx context.Context, name string) (organization *v1alpha1.Organization, err error)
	// ListOrganizations lists all organizations.
	ListOrganizations(
		ctx context.Context,
		pageSize uint32,
		pageToken string,
		reverse bool,
	) (organizations []*v1alpha1.Organization, nextPageToken string, err error)
	// ListUserOrganizations lists all organizations a user is member of.
	ListUserOrganizations(
		ctx context.Context,
		userId string,
		pageSize uint32,
		pageToken string,
		reverse bool,
	) (organizations []*v1alpha1.Organization, nextPageToken string, err error)
	// CreateOrganization creates a new organization.
	CreateOrganization(ctx context.Context, name string) (organization *v1alpha1.Organization, err error)
	// UpdateOrganizationName updates a organization's name.
	UpdateOrganizationName(
		ctx context.Context,
		id string,
		newName string,
	) (organization *v1alpha1.Organization, err error)
	// UpdateOrganizationNameByName updates a organization's name by name.
	UpdateOrganizationNameByName(
		ctx context.Context,
		name string,
		newName string,
	) (organization *v1alpha1.Organization, err error)
	// DeleteOrganization deletes a organization.
	DeleteOrganization(ctx context.Context, id string) (err error)
	// DeleteOrganizationByName deletes a organization by name.
	DeleteOrganizationByName(ctx context.Context, name string) (err error)
	// AddOrganizationBaseRepositoryScope adds a base repository scope to an organization by ID.
	AddOrganizationBaseRepositoryScope(
		ctx context.Context,
		id string,
		repositoryScope v1alpha1.RepositoryScope,
	) (err error)
	// AddOrganizationBaseRepositoryScopeByName adds a base repository scope to an organization by name.
	AddOrganizationBaseRepositoryScopeByName(
		ctx context.Context,
		name string,
		repositoryScope v1alpha1.RepositoryScope,
	) (err error)
	// RemoveOrganizationBaseRepositoryScope removes a base repository scope from an organization by ID.
	RemoveOrganizationBaseRepositoryScope(
		ctx context.Context,
		id string,
		repositoryScope v1alpha1.RepositoryScope,
	) (err error)
	// RemoveOrganizationBaseRepositoryScopeByName removes a base repository scope from an organization by name.
	RemoveOrganizationBaseRepositoryScopeByName(
		ctx context.Context,
		name string,
		repositoryScope v1alpha1.RepositoryScope,
	) (err error)
}
