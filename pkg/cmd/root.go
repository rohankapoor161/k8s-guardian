package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "guardian",
	Short: "Kubernetes deployment safety toolkit",
	Long: `Guardian provides safety gates, canary analysis, and incident prevention
for Kubernetes deployments. It helps platform teams enforce deployment standards
and prevent production incidents before they happen.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.guardian.yaml)")
	rootCmd.PersistentFlags().String("kubeconfig", "", "path to kubeconfig file")
	rootCmd.PersistentFlags().String("namespace", "default", "target namespace")
	rootCmd.PersistentFlags().Bool("verbose", false, "verbose output")
	
	viper.BindPFlag("kubeconfig", rootCmd.PersistentFlags().Lookup("kubeconfig"))
	viper.BindPFlag("namespace", rootCmd.PersistentFlags().Lookup("namespace"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".guardian")
		viper.SetConfigType("yaml")
		
		viper.AddConfigPath("$HOME/.guardian")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("GUARDIAN")

	if err := viper.ReadInConfig(); err == nil {
		if viper.GetBool("verbose") {
			fmt.Printf("Using config file: %s\n", viper.ConfigFileUsed())
		}
	}
}