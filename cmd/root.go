// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package cmd

import (
	"errors"
	"fmt"
	"os"

	goversion "github.com/caarlos0/go-version"
	"github.com/hugginsio/talosctx/internal"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/siderolabs/talos/pkg/machinery/client/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "talosctx <context>",
	Args:  cobra.MaximumNArgs(1),
	Short: "Faster context switching for talosctl",
	Run: func(cmd *cobra.Command, args []string) {
		previousContext, _ := internal.GetPreviousContext()

		talosconfig, err := config.Open("")
		if err != nil {
			fmt.Println("Failed to open talosconfig:", err)
			os.Exit(5)
		}

		var chosenContext string
		if len(args) == 0 {
			chosenContext, err = internal.Select(talosconfig.Contexts, previousContext)
			if errors.Is(err, fuzzyfinder.ErrAbort) {
				fmt.Println("Cancelled, context unchanged.")
				os.Exit(0)
			}

			if err != nil {
				fmt.Println("Failure while selecting context:", err)
				os.Exit(5)
			}
		}

		if len(args) == 1 {
			if args[0] == "-" {
				if previousContext == "" {
					fmt.Println("No previous context found")
					os.Exit(4)
				}

				if !internal.ValidateContext(talosconfig.Contexts, previousContext) {
					fmt.Println("Invalid context:", previousContext)
					os.Exit(4)
				}

				chosenContext = previousContext
			} else {
				chosenContext = args[0]
				if !internal.ValidateContext(talosconfig.Contexts, chosenContext) {
					fmt.Println("Invalid context:", previousContext)
					os.Exit(4)
				}
			}

		}

		_ = internal.SetPreviousContext(talosconfig.Context)
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
