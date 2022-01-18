package cmd

import (
	"io/fs"
	"log"
	"os"

	"github.com/hxhieu/json-to-env/cmd/utils"
	"github.com/spf13/cobra"
)

var outputFile string
var fieldSeparator string

var rootCmd = &cobra.Command{
	Use:   "json-to-env",
	Short: "Convert the JSON format to .env format",
	Long:  "Convert the JSON format to .env format",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Please provide a JSON file as the first argument")
		}
		content, err := os.ReadFile(args[0])
		if err == nil {
			options := utils.JsonToEnvOption{
				FieldSeparator: fieldSeparator,
			}
			str, err := utils.JsonToEnv(&content, &options)
			if err != nil {
				log.Fatal(err)
			}
			os.WriteFile(outputFile, []byte(*str), fs.ModePerm)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&outputFile,
		"output",
		"o",
		".env",
		"The output file",
	)

	rootCmd.PersistentFlags().StringVarP(
		&fieldSeparator,
		"separator",
		"s",
		"__",
		"The nested fields separator",
	)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
