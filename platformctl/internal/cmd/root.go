package cmd

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"

	"github.com/pkorobeinikov/platform/platform-lib/service/platform"
	"github.com/pkorobeinikov/platform/platform-lib/service/spec"
	"platformctl/internal/cfg"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "platformctl",
	Short: "The only tool for Cloud Native Microservices development",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if spec.ExistInCurrentDirectory() {
			_ = platform.CreateDirectory()

			if "minikube" == cfg.PlatformFlavorContainerRuntimeVM {
				runtimeEnvCmdName := fmt.Sprintf("%s-env", cfg.PlatformFlavorContainerRuntime)

				runtimeEnvCmdArgs := []string{
					"minikube",
					"--profile",
					cfg.PlatformMinikubeProfile,
					runtimeEnvCmdName,
				}

				mde := exec.CommandContext(context.TODO(), runtimeEnvCmdArgs[0], runtimeEnvCmdArgs[1:]...)
				env, err := mde.Output()
				if err != nil {
					return err
				}

				parsed, err := godotenv.Parse(bytes.NewBuffer(env))
				if err != nil {
					return err
				}

				for k, v := range parsed {
					if err := os.Setenv(k, v); err != nil {
						return os.Setenv(k, v)
					}
				}
			}
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.platformctl/platformctl.yaml)")
	rootCmd.PersistentFlags().StringVarP(&cfg.ServiceEnv, "service-env", "", "local", "service environment name")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(path.Join(home, ".platformctl"))
		viper.SetConfigName("platformctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	viper.SetDefault("platform.flavor.container-runtime", "docker")
	cfg.PlatformFlavorContainerRuntime = viper.GetString("platform.flavor.container-runtime")

	viper.SetDefault("platform.flavor.container-runtime-ctl", "docker")
	cfg.PlatformFlavorContainerRuntimeCtl = viper.GetString("platform.flavor.container-runtime-ctl")

	viper.SetDefault("platform.flavor.container-runtime-vm", "docker-desktop")
	cfg.PlatformFlavorContainerRuntimeVM = viper.GetString("platform.flavor.container-runtime-vm")

	viper.SetDefault("", "platform.minikube.profile")
	cfg.PlatformMinikubeProfile = viper.GetString("platform.minikube.profile")
}

const (
	cfgKeyGoEnvVars = "platform.go_env_vars"
)
