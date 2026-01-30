// Copyright (c) 2025 Arc Engineering
// SPDX-License-Identifier: MIT

package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	arcer "github.com/yourorg/arc-sdk/errors"
	"github.com/yourorg/arc-sdk/output"
)

// NewRootCmd creates the root command for arc-blog.
func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "arc-blog",
		Short: "Blog and article operations",
		Long:  "Fetch and manage blog posts or external articles (feature under active development).",
		Example: strings.TrimSpace(`
Example:
  # Fetch and save a single article (Phase 2 placeholder)
  arc-blog fetch --url https://example.com/post

Example:
  # Pull a feed into docs/research-external/blog/
  arc-blog fetch --playlist feed.xml --out-dir docs/research-external/blog

Example:
  # Pipe the fetched article into an analyzer workflow
  arc-blog fetch --url https://example.com/post --analyze --output json
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(fetchCmd())
	return cmd
}

func fetchCmd() *cobra.Command {
	var (
		articleURL string
		playlist   string
		outDir     string
		analyze    bool
		opts       output.OutputOptions
	)

	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "Fetch a blog/article (stub)",
		Example: strings.TrimSpace(`
Example:
  # Basic fetch placeholder
  arc-blog fetch --url https://example.com/post

Example:
  # Write the article into a specific directory
  arc-blog fetch --url https://example.com/post --out-dir docs/research-external/blog

Example:
  # Request analysis output (future hook) and emit JSON
  arc-blog fetch --url https://example.com/post --analyze --output json
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			if articleURL == "" && playlist == "" {
				return &arcer.CLIError{Msg: "provide --url or --playlist", Hint: "blog fetch requires at least one source"}
			}
			if err := opts.Resolve(); err != nil {
				return err
			}

			result := map[string]any{
				"url":       articleURL,
				"playlist":  playlist,
				"out_dir":   outDir,
				"analyze":   analyze,
				"status":    "stub",
				"next_step": "Phase 2 ingestion pipeline",
			}

			switch {
			case opts.Is(output.OutputJSON):
				enc := json.NewEncoder(cmd.OutOrStdout())
				enc.SetIndent("", "  ")
				return enc.Encode(result)
			case opts.Is(output.OutputYAML):
				enc := yaml.NewEncoder(cmd.OutOrStdout())
				defer enc.Close()
				return enc.Encode(result)
			case opts.Is(output.OutputQuiet):
				return nil
			default:
				fmt.Fprintf(cmd.OutOrStdout(), "blog fetch (stub) -> out_dir=%s analyze=%t\n", outDir, analyze)
				if articleURL != "" {
					fmt.Fprintf(cmd.OutOrStdout(), "  URL: %s\n", articleURL)
				}
				if playlist != "" {
					fmt.Fprintf(cmd.OutOrStdout(), "  Playlist: %s\n", playlist)
				}
				return nil
			}
		},
	}

	opts.AddOutputFlags(cmd, output.OutputTable)
	cmd.Flags().StringVar(&articleURL, "url", "", "Article URL to fetch")
	cmd.Flags().StringVar(&playlist, "playlist", "", "Playlist/feed file to ingest (Phase 2 placeholder)")
	cmd.Flags().StringVar(&outDir, "out-dir", "docs/research-external/blog", "Destination directory for fetched content")
	cmd.Flags().BoolVar(&analyze, "analyze", false, "Send fetched content into analyzer workflows (placeholder)")

	return cmd
}
