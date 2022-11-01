**Cobra**
It is an open source go module to build the powerful CLI such as kubectl (kubernetes CLI is made in cobra).

Example:

1. **main.go**

        package main

        import "cobra/cmd"

        func main() {
          cmd.Execute()
        }
       
 2. **root.go**
 
        package cmd

        import (
          "fmt"
          "os"

          "github.com/spf13/cobra"
        )

        var rootCmd = &cobra.Command{
          Use:   "myapp",
          Short: "Short desc",
          Long:  "Long desc",
          Run: func(cmd *cobra.Command, args []string) {

          },
        }

        func Execute() {
          if err := rootCmd.Execute(); err != nil {
            fmt.Println("Error occured!")
            os.Exit(1)
          }
        }
