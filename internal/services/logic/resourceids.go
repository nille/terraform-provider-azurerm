package logic

//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=IntegrationAccount -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=IntegrationAccountCertificate -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/account1/certificates/certificate1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=IntegrationAccountSchema -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/integrationAccount1/schemas/schema1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=IntegrationAccountSession -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Logic/integrationAccounts/integrationAccount1/sessions/session1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=IntegrationServiceEnvironment -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/group1/providers/Microsoft.Logic/integrationServiceEnvironments/ise1
