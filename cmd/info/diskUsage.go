/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package info

import (
	"fmt"

	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
)

var (
	KB = uint64(1024)
	InputPath string

)

func checkPath(path string)  (string, error) {

	usage := du.NewDiskUsage(path)
	if usage == nil {
		return "", fmt.Errorf("the path specified not found: %v", path)
	}

	return fmt.Sprintf("Used: %d", usage.Used()/(KB*KB)), nil
	
}

// diskUsageCmd represents the diskUsage command
var diskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "prints disk usage of current directory",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := checkPath(InputPath); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {

	diskUsageCmd.Flags().StringVarP(&InputPath, "path", "p", ".", "path to check Disk usaged")

	if err := diskUsageCmd.MarkFlagRequired("path"); err != nil {
		fmt.Println(err)
	}
	//rootCmd.AddCommand(diskUsageCmd)
	InfoCmd.AddCommand(diskUsageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diskUsageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diskUsageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
