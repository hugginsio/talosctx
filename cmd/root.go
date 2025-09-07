// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"fmt"
	"os"

	goversion "github.com/caarlos0/go-version"
	"github.com/hugginsio/talosctx/internal/fzf"
	"github.com/siderolabs/talos/pkg/machinery/client/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "talosctx",
	Short: "Faster context switching for talosctl",
	Run: func(cmd *cobra.Command, args []string) {
		talosconfig, err := config.Open("")
		if err != nil {
			fmt.Println("Failed to open talosconfig:", err)
			os.Exit(5)
		}

		chosenContext, err := fzf.Select(talosconfig.Contexts)
		if err != nil {
			fmt.Println("Failure while selecting context:", err)
			os.Exit(5)
		}

		talosconfig.Context = chosenContext
		if err := talosconfig.Save(""); err != nil {
			fmt.Println("Failed to save talosconfig:", err)
			os.Exit(5)
		}

		fmt.Printf("Switched to context \"%s\".", chosenContext)
	},
}

func Execute() {
	rootCmd.Version = goversion.GetVersionInfo().GitVersion

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
