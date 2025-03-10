// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package testclient

import (
	"context"
	"time"

	"github.com/Yamashou/gqlgenc/clientv2"
	"go.infratographer.com/x/gidx"
)

type TestClient interface {
	GetResourceProvider(ctx context.Context, id gidx.PrefixedID, interceptors ...clientv2.RequestInterceptor) (*GetResourceProvider, error)
	ListResourceProviders(ctx context.Context, id gidx.PrefixedID, orderBy *ResourceProviderOrder, interceptors ...clientv2.RequestInterceptor) (*ListResourceProviders, error)
	ResourceProviderCreate(ctx context.Context, input CreateResourceProviderInput, interceptors ...clientv2.RequestInterceptor) (*ResourceProviderCreate, error)
	ResourceProviderUpdate(ctx context.Context, id gidx.PrefixedID, input UpdateResourceProviderInput, interceptors ...clientv2.RequestInterceptor) (*ResourceProviderUpdate, error)
	ResourceProviderDelete(ctx context.Context, id gidx.PrefixedID, interceptors ...clientv2.RequestInterceptor) (*ResourceProviderDelete, error)
}

type Client struct {
	Client *clientv2.Client
}

func NewClient(cli clientv2.HttpClient, baseURL string, options *clientv2.Options, interceptors ...clientv2.RequestInterceptor) TestClient {
	return &Client{Client: clientv2.NewClient(cli, baseURL, options, interceptors...)}
}

type GetResourceProvider_ResourceProvider_Owner struct {
	ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
}

func (t *GetResourceProvider_ResourceProvider_Owner) GetID() *gidx.PrefixedID {
	if t == nil {
		t = &GetResourceProvider_ResourceProvider_Owner{}
	}
	return &t.ID
}

type GetResourceProvider_ResourceProvider struct {
	CreatedAt   time.Time                                  "json:\"createdAt\" graphql:\"createdAt\""
	Description *string                                    "json:\"description,omitempty\" graphql:\"description\""
	ID          gidx.PrefixedID                            "json:\"id\" graphql:\"id\""
	Name        string                                     "json:\"name\" graphql:\"name\""
	Owner       GetResourceProvider_ResourceProvider_Owner "json:\"owner\" graphql:\"owner\""
	UpdatedAt   time.Time                                  "json:\"updatedAt\" graphql:\"updatedAt\""
}

func (t *GetResourceProvider_ResourceProvider) GetCreatedAt() *time.Time {
	if t == nil {
		t = &GetResourceProvider_ResourceProvider{}
	}
	return &t.CreatedAt
}
func (t *GetResourceProvider_ResourceProvider) GetDescription() *string {
	if t == nil {
		t = &GetResourceProvider_ResourceProvider{}
	}
	return t.Description
}
func (t *GetResourceProvider_ResourceProvider) GetID() *gidx.PrefixedID {
	if t == nil {
		t = &GetResourceProvider_ResourceProvider{}
	}
	return &t.ID
}
func (t *GetResourceProvider_ResourceProvider) GetName() string {
	if t == nil {
		t = &GetResourceProvider_ResourceProvider{}
	}
	return t.Name
}
func (t *GetResourceProvider_ResourceProvider) GetOwner() *GetResourceProvider_ResourceProvider_Owner {
	if t == nil {
		t = &GetResourceProvider_ResourceProvider{}
	}
	return &t.Owner
}
func (t *GetResourceProvider_ResourceProvider) GetUpdatedAt() *time.Time {
	if t == nil {
		t = &GetResourceProvider_ResourceProvider{}
	}
	return &t.UpdatedAt
}

type ListResourceProviders_Entities_ResourceOwner_ResourceProvider_Edges_Node struct {
	ID   gidx.PrefixedID "json:\"id\" graphql:\"id\""
	Name string          "json:\"name\" graphql:\"name\""
}

func (t *ListResourceProviders_Entities_ResourceOwner_ResourceProvider_Edges_Node) GetID() *gidx.PrefixedID {
	if t == nil {
		t = &ListResourceProviders_Entities_ResourceOwner_ResourceProvider_Edges_Node{}
	}
	return &t.ID
}
func (t *ListResourceProviders_Entities_ResourceOwner_ResourceProvider_Edges_Node) GetName() string {
	if t == nil {
		t = &ListResourceProviders_Entities_ResourceOwner_ResourceProvider_Edges_Node{}
	}
	return t.Name
}

type ListResourceProviders_Entities_ResourceOwner_ResourceProvider_Edges struct {
	Node *ListResourceProviders_Entities_ResourceOwner_ResourceProvider_Edges_Node "json:\"node,omitempty\" graphql:\"node\""
}

func (t *ListResourceProviders_Entities_ResourceOwner_ResourceProvider_Edges) GetNode() *ListResourceProviders_Entities_ResourceOwner_ResourceProvider_Edges_Node {
	if t == nil {
		t = &ListResourceProviders_Entities_ResourceOwner_ResourceProvider_Edges{}
	}
	return t.Node
}

type ListResourceProviders_Entities_ResourceOwner_ResourceProvider struct {
	Edges []*ListResourceProviders_Entities_ResourceOwner_ResourceProvider_Edges "json:\"edges,omitempty\" graphql:\"edges\""
}

func (t *ListResourceProviders_Entities_ResourceOwner_ResourceProvider) GetEdges() []*ListResourceProviders_Entities_ResourceOwner_ResourceProvider_Edges {
	if t == nil {
		t = &ListResourceProviders_Entities_ResourceOwner_ResourceProvider{}
	}
	return t.Edges
}

type ListResourceProviders_Entities_ResourceOwner struct {
	ResourceProvider ListResourceProviders_Entities_ResourceOwner_ResourceProvider "json:\"resourceProvider\" graphql:\"resourceProvider\""
}

func (t *ListResourceProviders_Entities_ResourceOwner) GetResourceProvider() *ListResourceProviders_Entities_ResourceOwner_ResourceProvider {
	if t == nil {
		t = &ListResourceProviders_Entities_ResourceOwner{}
	}
	return &t.ResourceProvider
}

type ResourceProviderCreate_ResourceProviderCreate_ResourceProvider_Owner struct {
	ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
}

func (t *ResourceProviderCreate_ResourceProviderCreate_ResourceProvider_Owner) GetID() *gidx.PrefixedID {
	if t == nil {
		t = &ResourceProviderCreate_ResourceProviderCreate_ResourceProvider_Owner{}
	}
	return &t.ID
}

type ResourceProviderCreate_ResourceProviderCreate_ResourceProvider struct {
	CreatedAt   time.Time                                                            "json:\"createdAt\" graphql:\"createdAt\""
	Description *string                                                              "json:\"description,omitempty\" graphql:\"description\""
	ID          gidx.PrefixedID                                                      "json:\"id\" graphql:\"id\""
	Name        string                                                               "json:\"name\" graphql:\"name\""
	Owner       ResourceProviderCreate_ResourceProviderCreate_ResourceProvider_Owner "json:\"owner\" graphql:\"owner\""
	UpdatedAt   time.Time                                                            "json:\"updatedAt\" graphql:\"updatedAt\""
}

func (t *ResourceProviderCreate_ResourceProviderCreate_ResourceProvider) GetCreatedAt() *time.Time {
	if t == nil {
		t = &ResourceProviderCreate_ResourceProviderCreate_ResourceProvider{}
	}
	return &t.CreatedAt
}
func (t *ResourceProviderCreate_ResourceProviderCreate_ResourceProvider) GetDescription() *string {
	if t == nil {
		t = &ResourceProviderCreate_ResourceProviderCreate_ResourceProvider{}
	}
	return t.Description
}
func (t *ResourceProviderCreate_ResourceProviderCreate_ResourceProvider) GetID() *gidx.PrefixedID {
	if t == nil {
		t = &ResourceProviderCreate_ResourceProviderCreate_ResourceProvider{}
	}
	return &t.ID
}
func (t *ResourceProviderCreate_ResourceProviderCreate_ResourceProvider) GetName() string {
	if t == nil {
		t = &ResourceProviderCreate_ResourceProviderCreate_ResourceProvider{}
	}
	return t.Name
}
func (t *ResourceProviderCreate_ResourceProviderCreate_ResourceProvider) GetOwner() *ResourceProviderCreate_ResourceProviderCreate_ResourceProvider_Owner {
	if t == nil {
		t = &ResourceProviderCreate_ResourceProviderCreate_ResourceProvider{}
	}
	return &t.Owner
}
func (t *ResourceProviderCreate_ResourceProviderCreate_ResourceProvider) GetUpdatedAt() *time.Time {
	if t == nil {
		t = &ResourceProviderCreate_ResourceProviderCreate_ResourceProvider{}
	}
	return &t.UpdatedAt
}

type ResourceProviderCreate_ResourceProviderCreate struct {
	ResourceProvider ResourceProviderCreate_ResourceProviderCreate_ResourceProvider "json:\"resourceProvider\" graphql:\"resourceProvider\""
}

func (t *ResourceProviderCreate_ResourceProviderCreate) GetResourceProvider() *ResourceProviderCreate_ResourceProviderCreate_ResourceProvider {
	if t == nil {
		t = &ResourceProviderCreate_ResourceProviderCreate{}
	}
	return &t.ResourceProvider
}

type ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider_Owner struct {
	ID gidx.PrefixedID "json:\"id\" graphql:\"id\""
}

func (t *ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider_Owner) GetID() *gidx.PrefixedID {
	if t == nil {
		t = &ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider_Owner{}
	}
	return &t.ID
}

type ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider struct {
	CreatedAt   time.Time                                                            "json:\"createdAt\" graphql:\"createdAt\""
	Description *string                                                              "json:\"description,omitempty\" graphql:\"description\""
	ID          gidx.PrefixedID                                                      "json:\"id\" graphql:\"id\""
	Name        string                                                               "json:\"name\" graphql:\"name\""
	Owner       ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider_Owner "json:\"owner\" graphql:\"owner\""
	UpdatedAt   time.Time                                                            "json:\"updatedAt\" graphql:\"updatedAt\""
}

func (t *ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider) GetCreatedAt() *time.Time {
	if t == nil {
		t = &ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider{}
	}
	return &t.CreatedAt
}
func (t *ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider) GetDescription() *string {
	if t == nil {
		t = &ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider{}
	}
	return t.Description
}
func (t *ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider) GetID() *gidx.PrefixedID {
	if t == nil {
		t = &ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider{}
	}
	return &t.ID
}
func (t *ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider) GetName() string {
	if t == nil {
		t = &ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider{}
	}
	return t.Name
}
func (t *ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider) GetOwner() *ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider_Owner {
	if t == nil {
		t = &ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider{}
	}
	return &t.Owner
}
func (t *ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider) GetUpdatedAt() *time.Time {
	if t == nil {
		t = &ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider{}
	}
	return &t.UpdatedAt
}

type ResourceProviderUpdate_ResourceProviderUpdate struct {
	ResourceProvider ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider "json:\"resourceProvider\" graphql:\"resourceProvider\""
}

func (t *ResourceProviderUpdate_ResourceProviderUpdate) GetResourceProvider() *ResourceProviderUpdate_ResourceProviderUpdate_ResourceProvider {
	if t == nil {
		t = &ResourceProviderUpdate_ResourceProviderUpdate{}
	}
	return &t.ResourceProvider
}

type ResourceProviderDelete_ResourceProviderDelete struct {
	DeletedID gidx.PrefixedID "json:\"deletedID\" graphql:\"deletedID\""
}

func (t *ResourceProviderDelete_ResourceProviderDelete) GetDeletedID() *gidx.PrefixedID {
	if t == nil {
		t = &ResourceProviderDelete_ResourceProviderDelete{}
	}
	return &t.DeletedID
}

type GetResourceProvider struct {
	ResourceProvider GetResourceProvider_ResourceProvider "json:\"resourceProvider\" graphql:\"resourceProvider\""
}

func (t *GetResourceProvider) GetResourceProvider() *GetResourceProvider_ResourceProvider {
	if t == nil {
		t = &GetResourceProvider{}
	}
	return &t.ResourceProvider
}

type ListResourceProviders struct {
	Entities []*ListResourceProviders_Entities_ResourceOwner "json:\"_entities\" graphql:\"_entities\""
}

func (t *ListResourceProviders) GetEntities() []*ListResourceProviders_Entities_ResourceOwner {
	if t == nil {
		t = &ListResourceProviders{}
	}
	return t.Entities
}

type ResourceProviderCreate struct {
	ResourceProviderCreate ResourceProviderCreate_ResourceProviderCreate "json:\"resourceProviderCreate\" graphql:\"resourceProviderCreate\""
}

func (t *ResourceProviderCreate) GetResourceProviderCreate() *ResourceProviderCreate_ResourceProviderCreate {
	if t == nil {
		t = &ResourceProviderCreate{}
	}
	return &t.ResourceProviderCreate
}

type ResourceProviderUpdate struct {
	ResourceProviderUpdate ResourceProviderUpdate_ResourceProviderUpdate "json:\"resourceProviderUpdate\" graphql:\"resourceProviderUpdate\""
}

func (t *ResourceProviderUpdate) GetResourceProviderUpdate() *ResourceProviderUpdate_ResourceProviderUpdate {
	if t == nil {
		t = &ResourceProviderUpdate{}
	}
	return &t.ResourceProviderUpdate
}

type ResourceProviderDelete struct {
	ResourceProviderDelete ResourceProviderDelete_ResourceProviderDelete "json:\"resourceProviderDelete\" graphql:\"resourceProviderDelete\""
}

func (t *ResourceProviderDelete) GetResourceProviderDelete() *ResourceProviderDelete_ResourceProviderDelete {
	if t == nil {
		t = &ResourceProviderDelete{}
	}
	return &t.ResourceProviderDelete
}

const GetResourceProviderDocument = `query GetResourceProvider ($id: ID!) {
	resourceProvider(id: $id) {
		id
		name
		description
		owner {
			id
		}
		createdAt
		updatedAt
	}
}
`

func (c *Client) GetResourceProvider(ctx context.Context, id gidx.PrefixedID, interceptors ...clientv2.RequestInterceptor) (*GetResourceProvider, error) {
	vars := map[string]any{
		"id": id,
	}

	var res GetResourceProvider
	if err := c.Client.Post(ctx, "GetResourceProvider", GetResourceProviderDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

const ListResourceProvidersDocument = `query ListResourceProviders ($id: ID!, $orderBy: ResourceProviderOrder) {
	_entities(representations: [{__typename:"ResourceOwner",id:$id}]) {
		... on ResourceOwner {
			resourceProvider(orderBy: $orderBy) {
				edges {
					node {
						id
						name
					}
				}
			}
		}
	}
}
`

func (c *Client) ListResourceProviders(ctx context.Context, id gidx.PrefixedID, orderBy *ResourceProviderOrder, interceptors ...clientv2.RequestInterceptor) (*ListResourceProviders, error) {
	vars := map[string]any{
		"id":      id,
		"orderBy": orderBy,
	}

	var res ListResourceProviders
	if err := c.Client.Post(ctx, "ListResourceProviders", ListResourceProvidersDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

const ResourceProviderCreateDocument = `mutation ResourceProviderCreate ($input: CreateResourceProviderInput!) {
	resourceProviderCreate(input: $input) {
		resourceProvider {
			id
			name
			description
			owner {
				id
			}
			createdAt
			updatedAt
		}
	}
}
`

func (c *Client) ResourceProviderCreate(ctx context.Context, input CreateResourceProviderInput, interceptors ...clientv2.RequestInterceptor) (*ResourceProviderCreate, error) {
	vars := map[string]any{
		"input": input,
	}

	var res ResourceProviderCreate
	if err := c.Client.Post(ctx, "ResourceProviderCreate", ResourceProviderCreateDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

const ResourceProviderUpdateDocument = `mutation ResourceProviderUpdate ($id: ID!, $input: UpdateResourceProviderInput!) {
	resourceProviderUpdate(id: $id, input: $input) {
		resourceProvider {
			id
			name
			description
			owner {
				id
			}
			createdAt
			updatedAt
		}
	}
}
`

func (c *Client) ResourceProviderUpdate(ctx context.Context, id gidx.PrefixedID, input UpdateResourceProviderInput, interceptors ...clientv2.RequestInterceptor) (*ResourceProviderUpdate, error) {
	vars := map[string]any{
		"id":    id,
		"input": input,
	}

	var res ResourceProviderUpdate
	if err := c.Client.Post(ctx, "ResourceProviderUpdate", ResourceProviderUpdateDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

const ResourceProviderDeleteDocument = `mutation ResourceProviderDelete ($id: ID!) {
	resourceProviderDelete(id: $id) {
		deletedID
	}
}
`

func (c *Client) ResourceProviderDelete(ctx context.Context, id gidx.PrefixedID, interceptors ...clientv2.RequestInterceptor) (*ResourceProviderDelete, error) {
	vars := map[string]any{
		"id": id,
	}

	var res ResourceProviderDelete
	if err := c.Client.Post(ctx, "ResourceProviderDelete", ResourceProviderDeleteDocument, &res, vars, interceptors...); err != nil {
		if c.Client.ParseDataWhenErrors {
			return &res, err
		}

		return nil, err
	}

	return &res, nil
}

var DocumentOperationNames = map[string]string{
	GetResourceProviderDocument:    "GetResourceProvider",
	ListResourceProvidersDocument:  "ListResourceProviders",
	ResourceProviderCreateDocument: "ResourceProviderCreate",
	ResourceProviderUpdateDocument: "ResourceProviderUpdate",
	ResourceProviderDeleteDocument: "ResourceProviderDelete",
}
