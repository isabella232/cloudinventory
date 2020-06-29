package cmd

import (
        "encoding/json"
        "fmt"
        "github.com/adobe/cloudinventory/azurevnetcollector"
        "github.com/spf13/cobra"
        "io/ioutil"
        "strings"
)

// azurevnetCmd represents the azure command
var azurevnetCmd = &cobra.Command{
        Use:   "azurevnet",
        Short: "Dump Azure inventory. Currently supports Virtual networks",
        Run: func(cmd *cobra.Command, args []string) {
                path := cmd.Flag("path").Value.String()
                inputPath := cmd.Flag("inputPath").Value.String()
                var col azurevnetcollector.AzureCollector
                var err error
                if inputPath != "" {
                        data, err := ioutil.ReadFile(inputPath)
                        if err != nil {
                                fmt.Println("File reading error", err)
                                return
                        }
                        s := string(data)
                        subID := strings.Split(s, " ")
                        col, err = azurevnetcollector.NewAzureCollectorUserDefined(subID)
                        if err != nil {
                                fmt.Printf("Failed to create Azure vnet collector: %v\n", err)
                                return
                        }

                } else {
                        col, err = azurevnetcollector.NewAzureCollector()
                        if err != nil {
                                fmt.Printf("Failed to create Azure vnet collector: %v\n", err)
                                return
                        }
                }
                // Create a map per service
                result := make(map[string]interface{})
                err = collectVNets(col, result)
                if err != nil {
                        return
                }

                fmt.Printf("Dumping to %s\n", path)
                jsonBytes, err := json.MarshalIndent(result, "", "    ")
                if err != nil {
                        fmt.Printf("Error Marshalling JSON: %v\n", err)
                }
                err = ioutil.WriteFile(path, jsonBytes, 0644)
                if err != nil {
                        fmt.Printf("Error writing file: %v\n", err)
                }

        },
}

func collectVNets(col azurevnetcollector.AzureCollector, result map[string]interface{}) error {
        instances, err := col.CollectVirtualNetworks()
        if err != nil {
                fmt.Printf("Failed to gather virtual networks Data: %v\n", err)
                return err
        }
        fmt.Printf("Gathered virtual networks across %d subscriptions\n", len(instances))
        result["vnet"] = instances
        return nil
}

func init() {
        dumpCmd.AddCommand(azurevnetCmd)
        azurevnetCmd.PersistentFlags().StringP("inputPath", "i", "", "file path to take subscriptionIDs as input")
}
