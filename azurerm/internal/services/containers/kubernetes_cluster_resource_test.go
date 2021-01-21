package containers_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/acceptance"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/containers/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

type KubernetesClusterResource struct {
}

var (
	olderKubernetesVersion   = "1.18.14"
	currentKubernetesVersion = "1.19.6"
)

func TestAccKubernetes_all(t *testing.T) {
	// NOTE: this test is no longer used, but this assignment kicks around temporarily
	// to allow us to migrate off this without causing conflicts in open PR's
	_ = map[string]map[string]func(t *testing.T){
		"auth":               kubernetesAuthTests,
		"clusterAddOn":       kubernetesAddOnTests,
		"datasource":         kubernetesDataSourceTests,
		"network":            kubernetesNetworkAuthTests,
		"nodePool":           kubernetesNodePoolTests,
		"nodePoolDataSource": kubernetesNodePoolDataSourceTests,
		"other":              kubernetesOtherTests,
		"scaling":            kubernetesScalingTests,
		"upgrade":            kubernetesUpgradeTests,
	}

	t.Skip("Skipping since this is being run Individually")
}

func (t KubernetesClusterResource) Exists(ctx context.Context, clients *clients.Client, state *terraform.InstanceState) (*bool, error) {
	id, err := parse.ClusterID(state.ID)
	if err != nil {
		return nil, err
	}

	resp, err := clients.Containers.KubernetesClustersClient.Get(ctx, id.ResourceGroup, id.ManagedClusterName)
	if err != nil {
		return nil, fmt.Errorf("reading Kubernetes Cluster (%s): %+v", id.String(), err)
	}

	return utils.Bool(resp.ID != nil), nil
}

func kubernetesClusterUpdateNodePoolCount(resourceName string, nodeCount int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := acceptance.AzureProvider.Meta().(*clients.Client).Containers.AgentPoolsClient
		ctx := acceptance.AzureProvider.Meta().(*clients.Client).StopContext

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", resourceName)
		}

		nodePoolName := rs.Primary.Attributes["default_node_pool.0.name"]
		clusterName := rs.Primary.Attributes["name"]
		resourceGroup := rs.Primary.Attributes["resource_group_name"]

		nodePool, err := client.Get(ctx, resourceGroup, clusterName, nodePoolName)
		if err != nil {
			return fmt.Errorf("Bad: Get on agentPoolsClient: %+v", err)
		}

		if nodePool.StatusCode == http.StatusNotFound {
			return fmt.Errorf("Bad: Node Pool %q (Kubernetes Cluster %q / Resource Group: %q) does not exist", nodePoolName, clusterName, resourceGroup)
		}

		if nodePool.ManagedClusterAgentPoolProfileProperties == nil {
			return fmt.Errorf("Bad: Node Pool %q (Kubernetes Cluster %q / Resource Group: %q): `properties` was nil", nodePoolName, clusterName, resourceGroup)
		}

		nodePool.ManagedClusterAgentPoolProfileProperties.Count = utils.Int32(int32(nodeCount))

		future, err := client.CreateOrUpdate(ctx, resourceGroup, clusterName, nodePoolName, nodePool)
		if err != nil {
			return fmt.Errorf("Bad: updating node pool %q: %+v", nodePoolName, err)
		}

		if err := future.WaitForCompletionRef(ctx, client.Client); err != nil {
			return fmt.Errorf("Bad: waiting for update of node pool %q: %+v", nodePoolName, err)
		}

		return nil
	}
}
