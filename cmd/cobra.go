package cmd

import (
	"errors"
	"os"

	"snail/cmd/api"
	"snail/cmd/scripts"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "entry",
	SilenceUsage: false,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run:               func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(api.Server)       // 接口服务
	rootCmd.AddCommand(scripts.Register) // 脚本入口
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
