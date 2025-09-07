// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package fzf

import (
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/siderolabs/talos/pkg/machinery/client/config"
)

// Select prompts the user to choose a Talos context from the provided map. It returns the context
// name as string.
func Select(contexts map[string]*config.Context) (string, error) {
	var contextNames []string

	for name, _ := range contexts {
		contextNames = append(contextNames, name)
	}

	idx, err := fuzzyfinder.Find(contextNames, func(i int) string {
		return contextNames[i]
	})

	if err != nil {
		return "", err
	}

	return contextNames[idx], nil
}
