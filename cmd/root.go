// Copyright © 2018 ROBERT ALLEN <roball24@icloud.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "zpt [files | directories]",
	Short: "Password Encrypt In Zip Files",
	Long: `zpt is a convenience wrapper on UNIX 'zip' command with encryption. 
	run 'unzip' to retreive contents of you're password encrypted zip.`,
	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Zipping and Encrypting with zpt...\n\n")

		if _, err := os.Stat(outName); err == nil {
			fmt.Printf("  Output '%s' already exists. Exiting...\n", outName)
			return
		}

		// check all files / dirs exist
		for _, fname := range args {
			_, err := os.Stat(fname)
			if err != nil {
				fmt.Printf("\t%s does not exist\n", fname)
				return
			}
		}

		// pass files to zip
		zipArgs := append([]string{"-er", outName}, args...)
		c := exec.Command("zip", zipArgs...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		if err := c.Run(); err != nil {
			log.Fatal(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var outName string

func init() {
	rootCmd.Flags().StringVarP(&outName, "output", "o", "", "Output directory.")
	rootCmd.MarkFlagRequired("output")
}
