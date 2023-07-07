package main

import (
	"fmt"
	"github.com/gozelle/cobra"
	"os"
	"runtime"
)

func main() {
	if err := commandRoot().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func commandRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "example",
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(2)
		},
	}
	rootCmd.AddCommand(commandServe())
	rootCmd.AddCommand(commandVersion())
	return rootCmd
}

type serveOptions struct {
	// Config file path
	config string
	
	// Flags
	webHTTPAddr  string
	webHTTPSAddr string
}

func commandServe() *cobra.Command {
	options := serveOptions{}
	
	cmd := &cobra.Command{
		Use:     "serve [flags] [config file]",
		Short:   "Launch Server",
		Example: "example serve config.yaml",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true
			options.config = args[0]
			
			fmt.Println("begin serve")
			
			return nil
		},
	}
	
	flags := cmd.Flags()
	
	flags.StringVar(&options.webHTTPAddr, "web-http-addr", "", "Web HTTP address")
	flags.StringVar(&options.webHTTPSAddr, "web-https-addr", "", "Web HTTPS address")
	
	return cmd
}

var version = "0.0.1"

func commandVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version and exit",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf(
				"Exapme Version: %s\nGo Version: %s\nGo OS/ARCH: %s %s\n",
				version,
				runtime.Version(),
				runtime.GOOS,
				runtime.GOARCH,
			)
		},
	}
}
