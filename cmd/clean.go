/*
Copyright © 2021 Florian Wartner <florian@wartner.io>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"plugin"

	"github.com/spf13/cobra"
)

type Plugin interface {
	Clean()
}

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		loadPlugins()
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

func loadPlugins() {
	home, _ := os.UserHomeDir()
	err := filepath.Walk(home+"/.mac-cleanup/plugins",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			execPlugin(path)

			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

func execPlugin(path string) {

	plug, err := plugin.Open(path)

	symPlugin, err := plug.Lookup("Plugin")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var plugin Plugin
	plugin, ok := symPlugin.(Plugin)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	plugin.Clean()
}
