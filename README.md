# hermes 📯

Command line interface for iMessage databases on MacOS. With hermes you can view your messages, export them and view your message statistics. 


## Installation
You can either download the source code from this repository and build it yourself via `go build -o hermes main.go` or you can 
install it via `go install github.com/f-ewald/hermes@latest`.

## Usage
To get started, execute the hermes binary via `./hermes` from the build folder or via `hermes`, if you used
`go install github.com/f-ewald/hermes`. All available options will then be listed as shown below:

```text
Hermes is a command-line interface for iMessage databases.
You can use it to display conversations and to view statistics.

Usage:
  hermes [command]

Available Commands:
  check        Check prerequisites for hermes
  completion   Generate the autocompletion script for the specified shell
  conversation Show conversations, find participants
  help         Help about any command
  statistics   Display message statistics

Flags:
      --config string   config file (default is $HOME/.hermes.yaml)
  -h, --help            help for hermes
  -o, --output string   The output format. Can be either json, yaml or text (default "text")

Use "hermes [command] --help" for more information about a command.
```

## Contributing
Contributions are welcome via pull request.

## License
Copyright 2022 Freddy Ewald

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.