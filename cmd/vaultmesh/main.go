package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vaultmesh",
	Short: "VaultMesh is a decentralized encrypted P2P storage system",
}

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Manage the VaultMesh daemon",
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the VaultMesh daemon",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting VaultMesh daemon...")
	},
}

var uploadCmd = &cobra.Command{
	Use:   "upload [file]",
	Short: "Upload a file to the network",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Uploading file: %s\n", args[0])
	},
}

func init() {
	daemonCmd.AddCommand(startCmd)
	rootCmd.AddCommand(daemonCmd)
	rootCmd.AddCommand(uploadCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
