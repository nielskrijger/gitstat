package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/nielskrijger/gitstat/internal"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"time"
)

var (
	outputFile string

	rootCmd = &cobra.Command{
		Use:   "gitstat",
		Short: "Generates a json file of your git repo histories",
		Long: "This is a CLI program that generates a JSON file of your git history." +
			"\nAfter generating the JSON you can generate metrics and graphs using https://gitstat.com",
		Run: func(cmd *cobra.Command, args []string) {
			start := time.Now()
			parser := internal.NewParser()
			for i, arg := range os.Args {
				if i > 0 {
					err := parser.ParseProject(arg)
					check(err)
				}
			}
			res, _ := json.Marshal(parser)
			fmt.Printf("\nwriting output to %q", outputFile)
			err := ioutil.WriteFile(outputFile, res, 0644)
			check(err)
			
			fmt.Printf("\ntotal processing time was %s", time.Since(start).Round(time.Millisecond))
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&outputFile, "out", "o", "gitstat_result.json", "name of output file")
}

func check(err error) {
	if err == nil {
		return
	}

	fmt.Print(err)
}
