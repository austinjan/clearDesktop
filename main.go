package main

import (
	"flag"
	"fmt"
	"io"

	"os"
	"path/filepath"
	"strings"
)

const backupFolderName = "desktop_backup"

func getDesktopPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Desktop"), nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func backupShortcuts(desktopPath string) error {
	backupPathUser := filepath.Join(desktopPath, backupFolderName)

	if err := os.MkdirAll(backupPathUser, 0755); err != nil {
		return err
	}

	// Backup Public Desktop
	return backupFromPath(desktopPath, backupPathUser)
}

func restoreShortcuts(desktopPath string) error {
	backupPathUser := filepath.Join(desktopPath, backupFolderName)

	// Restore Public Desktop
	return restoreToPath(backupPathUser, desktopPath)
}

func backupFromPath(sourcePath, backupPath string) error {
	files, err := os.ReadDir(sourcePath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".lnk") || strings.HasSuffix(file.Name(), ".url") {
			src := filepath.Join(sourcePath, file.Name())
			dst := filepath.Join(backupPath, file.Name())

			if err := copyFile(src, dst); err != nil {
				return err
			}
			if err := os.Remove(src); err != nil {
				return err
			}
		}
	}
	return nil
}

func restoreToPath(backupPath, destinationPath string) error {
	files, err := os.ReadDir(backupPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".lnk") || strings.HasSuffix(file.Name(), ".url") {
			src := filepath.Join(backupPath, file.Name())
			dst := filepath.Join(destinationPath, file.Name())

			if err := copyFile(src, dst); err != nil {
				return err
			}
			if err := os.Remove(src); err != nil {
				return err
			}
		}
	}
	return nil
}

func deleteShortcuts(desktopPath string) error {
	files, err := os.ReadDir(desktopPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".lnk") {
			if err := os.Remove(filepath.Join(desktopPath, file.Name())); err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	// Add all desktop folder name
	userDesktopPath, err := getDesktopPath()
	if err != nil {
		fmt.Println("Error getting user desktop path:", err)
		return
	}
	publicDesktopPath := `C:\Users\Public\Desktop`
	oneNotePath := `C:\Users\austi\OneDrive\桌面`

	pathSlice := []string{
		userDesktopPath, publicDesktopPath, oneNotePath,
	}

	backup := flag.Bool("backup", false, "run backup")
	restore := flag.Bool("restore", false, "restore shortcut")
	clear := flag.Bool("clear", false, "clear shortcut")
	flag.Parse()

	// Backup shortcuts
	if *backup {
		fmt.Println("Backup....")
		for _, p := range pathSlice {
			if err := backupShortcuts(p); err != nil {
				fmt.Printf("Backup %s fail", p)
			}
		}

	}
	if *restore {
		fmt.Println("Restore...")
		for _, p := range pathSlice {
			if err := restoreShortcuts(p); err != nil {
				fmt.Printf("Restore %s fail", p)
			}
		}
	}

	if *clear {
		fmt.Println("Clear....")
		for _, p := range pathSlice {
			if err := deleteShortcuts(p); err != nil {
				fmt.Printf("Clear %s fail", p)
			}
		}

	}

	fmt.Println("Done")

	// Delete shortcuts
	// if err := deleteShortcuts(desktopPath); err != nil {
	// 	fmt.Println("Error deleting shortcuts:", err)
	// }

	// Restore shortcuts
	// if err := restoreShortcuts(desktopPath); err != nil {
	// 	fmt.Println("Error restoring shortcuts:", err)
	// }
}
