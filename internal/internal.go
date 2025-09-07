// Copyright (c) Kyle Huggins
// SPDX-License-Identifier: BSD-3-Clause

package internal

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/siderolabs/talos/pkg/machinery/client/config"
)

// Select prompts the user to choose a Talos context from the provided map. It returns the context
// name as string.
func Select(contexts map[string]*config.Context, preselect string) (string, error) {
	var contextNames []string

	for name := range contexts {
		contextNames = append(contextNames, name)
	}

	idx, err := fuzzyfinder.Find(contextNames, func(i int) string {
		return contextNames[i]
	}, fuzzyfinder.WithPreselected(func(i int) bool {
		if preselect == "" {
			return false
		}

		return contextNames[i] == preselect
	}))

	if err != nil {
		return "", err
	}

	return contextNames[idx], nil
}

// ValidateContext indicates if the provided context exists in the provided map of contexts.
func ValidateContext(contexts map[string]*config.Context, context string) bool {
	if _, ok := contexts[context]; !ok {
		return false
	}

	return true
}

func previousContextFilepath() string {
	return filepath.Join(xdg.DataHome, "talosctx")
}

// GetPreviousContext returns the previous selected context, if present.
func GetPreviousContext() (string, error) {
	filePath := previousContextFilepath()

	if _, err := os.Stat(filePath); err != nil {
		return "", err
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// SetPreviousContext writes the selected context to a data file.
func SetPreviousContext(context string) error {
	return os.WriteFile(previousContextFilepath(), []byte(context), fs.ModePerm)
}
