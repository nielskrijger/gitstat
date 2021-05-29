package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/nielskrijger/gitstat/internal"
	"github.com/spf13/cobra"
	"io/ioutil"
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

			for _, arg := range args {
				start2 := time.Now()
				err := parser.ParseProject(arg)
				if err != nil {
					fmt.Printf("%s\n", Red(err))
				} else {
					fmt.Printf(Green("done (%v)\n"), time.Since(start2).Round(time.Millisecond))
				}
			}

			res, _ := json.Marshal(parser)
			fmt.Printf("\nwriting output to %q", outputFile)
			err := ioutil.WriteFile(outputFile, res, 0644)
			if err != nil {
				fmt.Printf("%s\n", Red(err))
			}

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

const (
	Reset   = "\x1b[0m"
	FgRed   = "\x1b[31m"
	FgGreen = "\x1b[32m"
)

func Red(obj interface{}) string {
	return fmt.Sprintf("%s%v%s", FgRed, obj, Reset)
}

func Green(obj interface{}) string {
	return fmt.Sprintf("%s%v%s", FgGreen, obj, Reset)
}
