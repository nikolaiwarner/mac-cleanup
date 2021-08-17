
# MacCleanup

A cleanup script for macOS that cleans up your system.

## Features

By default it performs the following tasks:

* Empty the Trash on All Mounted Volumes and the Main HDD
* Clear System Log Files
* Cleanup Any Old Versions of Gems
* Cleanup iOS Applications
* Remove iOS Device Backups
* Cleanup Xcode Derived Data and Archives
* Reset iOS simulators
* Cleanup Homebrew Cache
* Remove Wget logs and hosts
* Clear Bash/ZSH history
* Purge Inactive Memory

It also can be extended using the built-in plugin system. (WIP)
  
## Installation

Install mac-cleanup using homebrew:

```bash
  brew tap mac-cleanup/tap
  brew install mac-cleanup/tap/mac-cleanup
```
    
## Usage

```bash
🗑️  Cleanup script for macOS (by github.com/fwartner)

Usage:
  mac-cleanup [command]

Available Commands:
  clean
  completion  generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help     help for mac-cleanup
  -t, --toggle   Help message for toggle

Use "mac-cleanup [command] --help" for more information about a command.
```

  
## Contributing

Contributions are always welcome!

Please adhere to this project's `code of conduct`.

  
## Support

For support, please join our discussions on this repository or create an issue.

  
## Feedback

If you have any feedback, please reach out to us in the discussions.

  
## Roadmap

- Finish Plugin-System


  
## License

This project is licensed under the [MIT](https://choosealicense.com/licenses/mit/) license.

  
## Authors

- [@fwartner](https://www.github.com/fwartner)

  
