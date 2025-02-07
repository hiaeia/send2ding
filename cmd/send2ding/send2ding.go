package send2ding

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/hiaeia/send2ding"
)

var (
	cfgFile  string
	dingTalk *send2ding.DingTalk
	Version  = "v1.0.2"
)

var cmd = &cobra.Command{
	Use:   "send2ding",
	Short: "send dingtalk message",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		token, err := cmd.Parent().PersistentFlags().GetString("token")
		if err != nil {
			fmt.Println(err.Error())
		}
		if token == "" {
			token = viper.GetString("token")
		}

		if token == "" {
			panic("Ding talk token not set")
		}

		secret, err := cmd.Parent().PersistentFlags().GetString("secret")
		if err != nil {
			fmt.Println(err.Error())
		}
		if secret == "" {
			secret = viper.GetString("secret")
		}

		dingTalk = send2ding.New(token, secret)
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/send2ding.toml)")
	cmd.PersistentFlags().String("token", "", "dingtalk robot token (require)")
	cmd.PersistentFlags().String("secret", "", "dingtalk robot secret")

	cmd.AddCommand(text, markdown, versionCmd, initCommand, link, feedcard)
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println("Error:", err.Error())
			os.Exit(1)
		}

		// check default config file
		_, err = os.Stat(fmt.Sprintf(`%s/send2ding.toml`, home))
		if err != nil {
			return
		}

		// Search config in home directory with name "send2ding" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("send2ding")
		viper.SetConfigType("toml")
	}

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Using config file:", viper.ConfigFileUsed())
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Dingbot",
	Long:  `All software has versions. This is Dingbot's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Dingbot Command", Version)
	},
}
