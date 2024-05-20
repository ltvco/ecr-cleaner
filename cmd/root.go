package cmd

import (
	"fmt"
	"os"

	"github.com/ltvco/ecr-cleaner/clean"
	"github.com/spf13/cobra"
)

var (
	untaggedOnly bool
	dryRun       bool
	verbose      bool
	keepLast     int
	arn          string
	region       string
)

func init() {
	// rootCmd.PersistentFlags().StringVarP(&arn, "arn", "a", "arn:aws:iam::722014088219:role/devops-read-only", "The AWS profile to use (default: 'arn:aws:iam::722014088219:role/devops-read-only')")
	rootCmd.PersistentFlags().BoolVarP(&untaggedOnly, "untagged-only", "u", true, "Only delete untagged images (default: true)")
	rootCmd.PersistentFlags().IntVarP(&keepLast, "keep-latest-count", "k", 10, "Keep the latest X images (default: 10)")
	rootCmd.PersistentFlags().StringVarP(&arn, "arn", "a", "arn:aws:iam::722014088219:role/devops-read-only", "The AWS profile to use (default: 'arn:aws:iam::722014088219:role/devops-read-only')")
	rootCmd.PersistentFlags().StringVarP(&region, "region", "r", "us-east-1", "AWS region (default: us-east-1)")
	rootCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "d", true, "Dry run of the clean action (default: true)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Use to see details of images to be deleted")
}

var rootCmd = &cobra.Command{
	Use:   "ecr-cleaner",
	Short: "clean amazon elastic container registries",
	Run: func(cmd *cobra.Command, args []string) {
		clean.CleanRepos(untaggedOnly, keepLast, arn, region, dryRun, verbose)
	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
