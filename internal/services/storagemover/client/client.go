package client

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/storagemover/2023-03-01/agents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagemover/2023-03-01/endpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/storagemover/2023-03-01/storagemovers"
	"github.com/hashicorp/terraform-provider-azurerm/internal/common"
)

type Client struct {
	StorageMoversClient *storagemovers.StorageMoversClient
	AgentsClient        *agents.AgentsClient
	EndpointsClient     *endpoints.EndpointsClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	storageMoversClient, err := storagemovers.NewStorageMoversClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building Storage Movers client: %+v", err)
	}
	o.Configure(storageMoversClient.Client, o.Authorizers.ResourceManager)

	agentsClient, err := agents.NewAgentsClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building agent client: %+v", err)
	}
	o.Configure(agentsClient.Client, o.Authorizers.ResourceManager)

	endpointsClient, err := endpoints.NewEndpointsClientWithBaseURI(o.Environment.ResourceManager)
	if err != nil {
		return nil, fmt.Errorf("building endpoints client: %+v", err)
	}
	o.Configure(endpointsClient.Client, o.Authorizers.ResourceManager)

	return &Client{
		StorageMoversClient: storageMoversClient,
		AgentsClient:        agentsClient,
		EndpointsClient:     endpointsClient,
	}, nil
}
