# Desktop Shortcut Backup Tool

A simple Go application to backup and restore desktop shortcuts on Windows. The tool handles standard shortcuts (`.lnk` files), Internet shortcuts (`.url` files), and provides functionality to manage shortcuts from both the user's desktop and the public desktop.

## Features

- **Backup Shortcuts**: Quickly backup all desktop shortcuts to a specified directory.
- **Restore Shortcuts**: Easily restore all shortcuts from the backup directory to the desktop.


## Usage

1. Clone the repository or download the source code.
2. Navigate to the directory containing the Go program.
3. Uncomment the desired operation in the `main` function, i.e., Backup or Restore.
4. Run the program.

```bash
go run main.go -backup
go run main.go -restore
```

