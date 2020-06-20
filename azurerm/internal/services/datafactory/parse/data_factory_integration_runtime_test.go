package parse

import "testing"

func TestParseDataFactoryIntegrationRuntimeID(t *testing.T) {
	testData := []struct {
		Name     string
		Input    string
		Expected *DataFactoryIntegrationRuntimeId
	}{
		{
			Name:     "Empty",
			Input:    "",
			Expected: nil,
		},
		{
			Name:     "No Data Factory segment",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/",
			Expected: nil,
		},
		{
			Name:     "No Integration Runtime name",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.DataFactory/factories/factory1/integrationruntimes/",
			Expected: nil,
		},
		{
			Name:     "Case incorrect in path element",
			Input:    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/microsoft.dataFactory/factories/factory1/Integrationruntimes/integrationRuntimeName",
			Expected: nil,
		},
		{
			Name:  "Valid",
			Input: "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myGroup1/providers/Microsoft.DataFactory/factories/factory1/integrationruntimes/integrationRuntimeName",
			Expected: &DataFactoryIntegrationRuntimeId{
				ResourceGroup: "myGroup1",
				Name:          "integrationRuntimeName",
				DataFactory:   "factory1",
			},
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Name)

		actual, err := DataFactoryIntegrationRuntimeID(v.Input)
		if err != nil {
			if v.Expected == nil {
				continue
			}

			t.Fatalf("Expected a value but got an error: %s", err)
		}

		if actual.Name != v.Expected.Name {
			t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
		}

		if actual.ResourceGroup != v.Expected.ResourceGroup {
			t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
		}
	}
}
