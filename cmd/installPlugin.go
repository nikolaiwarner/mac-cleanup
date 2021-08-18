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
	"os"

	"github.com/go-git/go-git/v5"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// installPluginCmd represents the installPlugin command
var installPluginCmd = &cobra.Command{
	Use:   "install-plugin {plugin}",
	Short: "Install a new plugin",
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := os.UserHomeDir()
		createDir(home + "/.mac-cleanup/plugins/" + args[0])
		_, err := git.PlainClone(home+"/.mac-cleanup/plugins/"+args[0], false, &git.CloneOptions{
			URL: "https://github.com/" + args[0] + ".git",
		})

		if err != nil {
			handleError(err, args[0])
		} else {
			log.Info("Plugin installed")
		}
	},
}

func init() {
	rootCmd.AddCommand(installPluginCmd)
	log.SetFormatter(&log.TextFormatter{})
}

func handleError(err error, plugin string) {
	if err == git.ErrRepositoryAlreadyExists {
		log.WithFields(log.Fields{
			"plugin": plugin,
		}).Error("Plugin already exists")
	}

	// if err != nil {
	// 	log.Error(err)
	// }
}

func createDir(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			log.WithFields(log.Fields{
				"dir": dir,
			}).Error(err)
		}
	}
}
