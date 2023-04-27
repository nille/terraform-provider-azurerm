package common

import (
	"fmt"
	"os"
	"strings"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/terraform-provider-azurerm/internal/features"
	"github.com/hashicorp/terraform-provider-azurerm/version"
)

type Authorizers struct {
	BatchManagement auth.Authorizer
	KeyVault        auth.Authorizer
	ResourceManager auth.Authorizer
	Storage         auth.Authorizer
	Synapse         auth.Authorizer

	// Some data-plane APIs require a token scoped for a specific endpoint
	AuthorizerFunc ApiAuthorizerFunc
}

type ApiAuthorizerFunc func(api environments.Api) (auth.Authorizer, error)

type ClientOptions struct {
	Authorizers *Authorizers
	Environment environments.Environment
	Features    features.UserFeatures

	SubscriptionId   string
	TenantId         string
	PartnerId        string
	TerraformVersion string

	CustomCorrelationRequestID  string
	DisableCorrelationRequestID bool

	DisableTerraformPartnerID bool
	SkipProviderReg           bool
	StorageUseAzureAD         bool

	// Keep these around for convenience with Autorest based clients, remove when we are no longer using autorest
	AzureEnvironment        azure.Environment
	ResourceManagerEndpoint string

	// Legacy authorizers for go-autorest
	AttestationAuthorizer     autorest.Authorizer
	BatchManagementAuthorizer autorest.Authorizer
	KeyVaultAuthorizer        autorest.Authorizer
	ResourceManagerAuthorizer autorest.Authorizer
	StorageAuthorizer         autorest.Authorizer
	SynapseAuthorizer         autorest.Authorizer
}

// Configure set up a resourcemanager.Client using an auth.Authorizer from hashicorp/go-azure-sdk
func (o ClientOptions) Configure(c *resourcemanager.Client, authorizer auth.Authorizer) {
	c.Authorizer = authorizer
	c.UserAgent = userAgent(c.UserAgent, o.TerraformVersion, o.PartnerId, o.DisableTerraformPartnerID)

	requestMiddlewares := make([]client.RequestMiddleware, 0)
	if !o.DisableCorrelationRequestID {
		id := o.CustomCorrelationRequestID
		if id == "" {
			id = correlationRequestID()
		}
		requestMiddlewares = append(requestMiddlewares, correlationRequestIDMiddleware(id))
	}
	requestMiddlewares = append(requestMiddlewares, requestLoggerMiddleware("AzureRM"))
	c.RequestMiddlewares = &requestMiddlewares

	c.ResponseMiddlewares = &[]client.ResponseMiddleware{
		responseLoggerMiddleware("AzureRM"),
	}
}

// ConfigureClient sets up an autorest.Client using an autorest.Authorizer
func (o ClientOptions) ConfigureClient(c *autorest.Client, authorizer autorest.Authorizer) {
	c.UserAgent = userAgent(c.UserAgent, o.TerraformVersion, o.PartnerId, o.DisableTerraformPartnerID)

	c.Authorizer = authorizer
	c.Sender = sender.BuildSender("AzureRM")
	c.SkipResourceProviderRegistration = o.SkipProviderReg
	if !o.DisableCorrelationRequestID {
		id := o.CustomCorrelationRequestID
		if id == "" {
			id = correlationRequestID()
		}
		c.RequestInspector = withCorrelationRequestID(id)
	}
}

func userAgent(userAgent, tfVersion, partnerID string, disableTerraformPartnerID bool) string {
	// FORK: this gives us the ability to add a Pulumi Specific user agent
	providerUserAgent := fmt.Sprintf("pulumi-azure/%s", version.ProviderVersion)
	userAgent = strings.TrimSpace(fmt.Sprintf("%s %s", userAgent, providerUserAgent))

	// append the CloudShell version to the user agent if it exists
	if azureAgent := os.Getenv("AZURE_HTTP_USER_AGENT"); azureAgent != "" {
		userAgent = fmt.Sprintf("%s %s", userAgent, azureAgent)
	}

	// only one pid can be interpreted currently
	// hence, send partner ID if present, otherwise send Pulumi GUID
	// unless users have opted out
	if partnerID == "" && !disableTerraformPartnerID {
		// FORK: Microsoft’s Pulumi Partner ID is this specific GUID
		partnerID = "a90539d8-a7a6-5826-95c4-1fbef22d4b22"
	}

	if partnerID != "" {
		// Tolerate partnerID UUIDs without the "pid-" prefix
		userAgent = fmt.Sprintf("%s pid-%s", userAgent, strings.TrimPrefix(partnerID, "pid-"))
	}

	return userAgent
}
