package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hxhieu/json-to-env/cmd/utils"
	"github.com/spf13/cobra"
)

var OutputFile string

var rootCmd = &cobra.Command{
	Use:   "json-to-env",
	Short: "Convert the JSON format to .env format",
	Long:  "Convert the JSON format to .env format",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Please provide a JSON file as the first argument")
		}
		content, err := ioutil.ReadFile(args[0])
		if err == nil {
			str, err := utils.JsonToEnv(&content)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Print(*str)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&OutputFile,
		"output",
		"o",
		".env",
		"The output file",
	)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
