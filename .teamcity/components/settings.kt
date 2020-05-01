// specifies the default hour (UTC) at which tests should be triggered, if enabled
var defaultStartHour = 0

// specifies the default level of parallelism per-service-package
var defaultParallelism = 10

var locations = mapOf(
        "public" to LocationConfiguration("westeurope", "eastus2", "francecentral", false),
        "germany" to LocationConfiguration("germanynortheast", "germanycentral", "", false)
)

// specifies the list of Azure Environments where tests should be run nightly
var runNightly = mapOf(
        "public" to true
)

// specifies a list of services which should be run with a custom test configuration
var serviceTestConfigurationOverrides = mapOf(
        // The API Management tests take ~45m each
        "apimanagement" to testConfiguration(10, defaultStartHour),

        // Compute is a large package
        "compute" to testConfiguration(10, defaultStartHour),

        // The AKS API has a low rate limit
        "containers" to testConfiguration(5, defaultStartHour),

        // Data Lake has a low quota
        "datalake" to testConfiguration(2, defaultStartHour),

        // Network is a large package
        "network" to testConfiguration(10, defaultStartHour),

        // SignalR only allows provisioning one "Free" instance at a time,
        // which is used in multiple tests
        "signalr" to testConfiguration(1, defaultStartHour)
)
