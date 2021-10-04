package helper

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2020-06-01/resources"

	"github.com/hashicorp/terraform-provider-azurerm/internal/services/mssql/parse"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

// FindDatabaseReplicationPartners looks for partner databases having one of the specified replication roles, by
// reading any replication links then attempting to discover and match the corresponding server/database resources for
// the other end of the link.
func FindDatabaseReplicationPartners(ctx context.Context, databasesClient *sql.DatabasesClient, replicationLinksClient *sql.ReplicationLinksClient, resourcesClient *resources.Client, id parse.DatabaseId, rolesToFind []sql.ReplicationRole) ([]sql.Database, error) {
	var partnerDatabases []sql.Database

	matchesRole := func(role sql.ReplicationRole) bool {
		for _, r := range rolesToFind {
			if r == role {
				return true
			}
		}
		return false
	}

	resp, err := replicationLinksClient.ListByDatabase(ctx, id.ResourceGroup, id.ServerName, id.Name)
	if err != nil {
		return nil, fmt.Errorf("reading Replication Links for %s: %+v", id, err)
	}
	if resp.Value == nil {
		return nil, fmt.Errorf("reading Replication Links for %s: response was nil", id)
	}

	for _, link := range *resp.Value {
		linkProps := link.ReplicationLinkProperties

		if linkProps == nil {
			log.Printf("[INFO] Replication Link Properties was nil for %s", id)
			continue
		}
		if linkProps.PartnerLocation == nil || linkProps.PartnerServer == nil || linkProps.PartnerDatabase == nil {
			log.Printf("[INFO] Replication Link Properties was invalid for %s", id)
			continue
		}

		log.Printf("[INFO] Replication Link found for %s", id)

		// Look for candidate partner SQL servers
		filter := fmt.Sprintf("(resourceType eq 'Microsoft.Sql/servers') and ((name eq '%s'))", *linkProps.PartnerServer)
		var resourceList []resources.GenericResourceExpanded
		for resourcesIterator, err := resourcesClient.ListComplete(ctx, filter, "", nil); resourcesIterator.NotDone(); err = resourcesIterator.NextWithContext(ctx) {
			if err != nil {
				return nil, fmt.Errorf("retrieving Partner SQL Servers with filter %q for %s: %+v", filter, id, err)
			}
			resourceList = append(resourceList, resourcesIterator.Value())
		}

		for _, server := range resourceList {
			if server.ID == nil {
				log.Printf("[INFO] Partner SQL Server ID was nil for %s", id)
				continue
			}

			partnerServerId, err := parse.ServerID(*server.ID)
			if err != nil {
				return nil, fmt.Errorf("parsing Partner SQL Server ID %q: %+v", *server.ID, err)
			}

			// Check if like-named server has a database named like the partner database, also with a replication link
			linksPossiblePartner, err := replicationLinksClient.ListByDatabase(ctx, partnerServerId.ResourceGroup, partnerServerId.Name, *linkProps.PartnerDatabase)
			if err != nil {
				if utils.ResponseWasNotFound(linksPossiblePartner.Response) {
					log.Printf("[INFO] no replication link found for Database %q (%s)", *linkProps.PartnerDatabase, partnerServerId)
					continue
				}
				return nil, fmt.Errorf("reading Replication Links for Database %s (%s): %+v", *linkProps.PartnerDatabase, partnerServerId, err)
			}
			if linksPossiblePartner.Value == nil {
				return nil, fmt.Errorf("reading Replication Links for Database %s (%s): response was nil", *linkProps.PartnerDatabase, partnerServerId)
			}

			for _, linkPossiblePartner := range *linksPossiblePartner.Value {
				if linkPossiblePartner.ReplicationLinkProperties == nil {
					log.Printf("[INFO] Replication Link Properties was nil for Database %s (%s)", *linkProps.PartnerDatabase, partnerServerId)
					continue
				}

				linkPropsPossiblePartner := *linkPossiblePartner.ReplicationLinkProperties

				// If the database has a replication link for the specified role and a matching partner location, we'll consider it a partner of this database
				if matchesRole(linkPropsPossiblePartner.Role) && *linkPossiblePartner.Location == *linkProps.PartnerLocation {
					partnerDatabaseId := parse.NewDatabaseID(partnerServerId.SubscriptionId, partnerServerId.ResourceGroup, partnerServerId.Name, *linkProps.PartnerDatabase)
					partnerDatabase, err := databasesClient.Get(ctx, partnerDatabaseId.ResourceGroup, partnerDatabaseId.ServerName, partnerDatabaseId.Name)
					if err != nil {
						return nil, fmt.Errorf("retrieving Partner %s: %+v", partnerDatabaseId, err)
					}
					if partnerDatabase.ID != nil {
						partnerDatabases = append(partnerDatabases, partnerDatabase)
					}
				}
			}
		}
	}

	return partnerDatabases, nil
}
