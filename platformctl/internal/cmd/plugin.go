package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Execute plugins",
}

func init() {
	addPluginCommand(rootCmd)
}

func addPluginCommand(parent *cobra.Command) {
	pluginDir, err := homedir.Expand("~/.platformctl/plugin/")
	if err != nil {
		return
	}

	type plugin struct {
		path string
		name string
	}

	plugins := make([]plugin, 0)
	_ = filepath.Walk(pluginDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		const prefix = "platformctl-plugin-"
		if info.IsDir() && strings.HasPrefix(info.Name(), prefix) {
			plugins = append(plugins, plugin{
				path: path,
				name: strings.TrimPrefix(info.Name(), prefix),
			})
		}

		return nil
	})

	for _, plugin := range plugins {
		var concretePluginCmd = &cobra.Command{
			Use:                plugin.name,
			Short:              fmt.Sprintf("See %s --help for details", plugin.name),
			Args:               cobra.ArbitraryArgs,
			DisableFlagParsing: true,
			RunE: func(cmd *cobra.Command, args []string) error {
				var err error

				pe := path.Join(plugin.path, "main")

				e, err := homedir.Expand(pe)
				if err != nil {
					return err
				}

				bypassedArgs := os.Args[3:]
				c := exec.CommandContext(cmd.Context(), e, bypassedArgs...)

				c.Stdout = os.Stdout
				c.Stderr = os.Stderr

				return c.Run()
			},
		}

		pluginCmd.AddCommand(concretePluginCmd)
	}

	parent.AddCommand(pluginCmd)
}
