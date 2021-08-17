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
	"os"
	"path/filepath"
	"plugin"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Plugin interface {
	Cleanup()
}

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		DeleteFiles("/Volumes/*/.Trashes/*")
		DeleteFiles("~/.Trash/*")
		DeleteFiles("/Library/Caches/*")
		DeleteFiles("/System/Library/Caches/*")
		DeleteFiles("~/Library/Caches/*")
		DeleteFiles("/private/var/folders/bh/*/*/*/*")
		DeleteFiles("/private/var/log/asl/*.asl")
		DeleteFiles("/Library/Logs/DiagnosticReports/*")
		DeleteFiles("/Library/Logs/CreativeCloud/*")
		DeleteFiles("/Library/Logs/adobegc.log")
		DeleteFiles("~/Library/Containers/com.apple.mail/Data/Library/Logs/Mail/*")
		DeleteFiles("~/Library/Logs/CoreSimulator/*")

		checkForPlugins()
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

func WalkMatch(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func checkForPlugins() {
	home, _ := os.UserHomeDir()
	root := home + "/.mac-cleanup/plugins/"
	files, _ := WalkMatch(root, "*.so")
	if len(files) != 0 {
		fmt.Println("Loading plugins..")
		loadPlugins()
		return
	}
}

func loadPlugins() {
	home, _ := os.UserHomeDir()
	root := home + "/.mac-cleanup/plugins/"
	files, err := WalkMatch(root, "*.so")
	if err != nil {
		log.Panic(err)
	}
	for _, file := range files {
		execPlugin(file)
	}
}

func execPlugin(path string) {
	plug, err := plugin.Open(path)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	symPlugin, err := plug.Lookup("Cleanup")
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	cleanup, ok := symPlugin.(Plugin)
	if !ok {
		log.Info("unexpected type from module symbol")
		os.Exit(1)
	}

	cleanup.Cleanup()
}

func DeleteFiles(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
