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

	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
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

	for _, p := range plugins {
		func(p plugin) {
			var concretePluginCmd = &cobra.Command{
				Use:                p.name,
				Short:              fmt.Sprintf("See %s --help for details", p.name),
				Args:               cobra.ArbitraryArgs,
				DisableFlagParsing: true,
				RunE: func(cmd *cobra.Command, args []string) error {
					var err error

					s, err := spec.Read()
					if err != nil {
						return err
					}

					pe := path.Join(p.path, "main")

					e, err := homedir.Expand(pe)
					if err != nil {
						return err
					}

					bypassedArgs := os.Args[3:]
					c := exec.CommandContext(cmd.Context(), e, bypassedArgs...)

					envs := s.ShellEnvironmentFor("local")
					c.Env = append(os.Environ(), envs...)

					c.Stdout = os.Stdout
					c.Stderr = os.Stderr

					return c.Run()
				},
			}

			pluginCmd.AddCommand(concretePluginCmd)
		}(p)
	}

	parent.AddCommand(pluginCmd)
}
