# Emulate Terminal

Emulate Terminal is a simple command-line application built with Go and the Cobra library. It demonstrates basic CLI functionality including greeting and version commands.

## Installation

To install Emulate Terminal, make sure you have Go installed on your system, then run:

```
go get github.com/JanDez/emulate_terminal
```

## Usage

After installation, you can run the application using:

```
emulate_terminal [command]
```

Available commands:

- `greet`: Prints a greeting message
- `version`: Displays the current version of the application
- `help`: Shows help information about available commands

### Examples

1. To get a greeting:
   ```
   emulate_terminal greet
   ```

2. To check the version:
   ```
   emulate_terminal version
   ```

3. To get help:
   ```
   emulate_terminal --help
   ```

## Development

This project uses the Cobra library for CLI functionality. The main components are:

- `main.go`: The entry point of the application
- `cmd/root.go`: Defines the root command and basic configuration
- `cmd/greet.go`: Implements the greet command
- `cmd/version.go`: Implements the version command (defined in root.go)

To add new commands, create a new file in the `cmd` directory and use the Cobra library to define the command structure.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.