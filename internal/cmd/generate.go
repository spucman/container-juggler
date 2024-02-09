package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sgeisbacher/container-juggler/internal/generation"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var generateCmd = &cobra.Command{
	Use:   "generate [scenario]",
	Short: "generates docker-compose.yml for specified scenario",
	Long: `TODO: longer description

[scenario] defaults to all`,
	Args: cobra.RangeArgs(0, 1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		scenarios := viper.GetStringMapStringSlice("scenarios")
		var res []string
		for key := range scenarios {
			if len(toComplete) != 0 {
				if strings.HasPrefix(key, toComplete) {
					res = append(res, key)
				}
			} else {
				res = append(res, "blub")
			}
		}
		return res, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running generate command ...")
		generator := generation.CreateGenerator()
		scenario := ""
		if len(args) > 0 {
			scenario = args[0]
		}
		composeFile, err := os.Create("docker-compose.yml")
		if err != nil {
			log.Fatal("could not create docker-compose.yml")
		}
		if err := generator.Generate(scenario, composeFile); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
