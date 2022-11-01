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

3. **trace.go**
        
        // Cobra CLI - command with our own flags
        package cmd

        import (
                "fmt"
                "github.com/spf13/cobra"
        )

        // traceCmd represents the trace command
        var traceCmd = &cobra.Command{
                Use:   "trace",
                Short: "trace the ip",
                Long:  `trace the ip subcommand`,
                Run: func(cmd *cobra.Command, args []string) {
                        fmt.Println("trace subcommand called", args)
                        f, _ := cmd.Flags().GetString("ip")
                        fmt.Println("We have flags now, IP :", f)
                },
        }

        func init() {
                rootCmd.AddCommand(traceCmd)
                rootCmd.PersistentFlags().String("ip", "", "Enter IP address")
        }

        /*
        Output
        ~/Documents/Golang programs/cobra$ go run main.go trace 1.1 --ip str
        trace subcommand called [1.1]
        We have flags now, IP : str
        */


