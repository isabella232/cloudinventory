package azurecollector

import "testing"

// TestCollectSQLDBs test the function CollectSQLDBs
func TestCollectSQLDBs(t *testing.T) {
        if testing.Short() {
                t.Skip("Skipping test in short mode")
        }
        col, err := NewAzureCollector()
        if err != nil {
                t.Errorf("Failed to create collector: %v", err)
        }
        _, err = col.CollectSQLDBs()
        if err != nil {
                t.Errorf("Failed to collect SQL Databases: %v", err)
        }
}

// TestCollectVMS test the function CollectVMS
func TestCollectVMS(t *testing.T) {
        if testing.Short() {
                t.Skip("Skipping test in short mode")
        }
        col, err := NewAzureCollector()
        if err != nil {
                t.Errorf("Failed to create collector: %v", err)
        }
        _, err = col.CollectVMS()
        if err != nil {
                t.Errorf("Failed to collect Virtual Machines: %v", err)
        }
}

// TestCollectLoadBalancers test the function CollectLoadBalancers
func TestCollectLoadBalancers(t *testing.T) {
        if testing.Short() {
                t.Skip("Skipping test in short mode")
        }
        col, err := NewAzureCollector()
        if err != nil {
                t.Errorf("Failed to create collector: %v", err)
        }
        _, err = col.CollectLoadBalancers()
        if err != nil {
                t.Errorf("Failed to collect Load Balancers: %v", err)
        }
}