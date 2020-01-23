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
					fmt.Printf("\nstart processing %q\n", arg)
					err := parser.ParseProject(arg)
					check(err)
					fmt.Printf("\nprocessing %q took %s\n", arg, time.Since(start).Round(time.Millisecond))
				}
			}
			if len(os.Args) >= 3 {
				fmt.Printf("\ntotal processing time was %s", time.Since(start).Round(time.Millisecond))
			}
			res, _ := json.Marshal(parser)
			err := ioutil.WriteFile(outputFile, res, 0644)
			fmt.Printf("\nwriting output to %q", outputFile)
			check(err)
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
