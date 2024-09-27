package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "music_service_client",
	Short: "A brief description of your application",
	Long: ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var _config string

func getHostPort() (string, string){
	viper.SetConfigFile(_config)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	host := viper.GetString("server_host")
	port := viper.GetString("server_port")
	return host, port
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(&_config, "config", "config.yaml", "")
}


