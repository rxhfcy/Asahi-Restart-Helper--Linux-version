package main

import (
	_ "embed"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"fyne.io/systray"
)

func requireCommand(name string) string {
	absPath, err := exec.LookPath(name)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Required command not found:", name)
		os.Exit(1)
	}
	return absPath
}

var asahiBlessCmd = requireCommand("asahi-bless")

func setupAutostart(homeDir string) {
	autostartDir := filepath.Join(homeDir, ".config", "autostart")
	autostartFile := filepath.Join(autostartDir, "asahi-reboot-switcher.desktop")

	// Check if the autostart file already exists
	if _, err := os.Stat(autostartFile); os.IsNotExist(err) {
		// Create the autostart directory if it doesn't exist
		if err := os.MkdirAll(autostartDir, 0755); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create autostart directory:", err)
			return
		}

		// Open the source file for reading
		srcFile, err := os.Open("/usr/share/applications/asahi-reboot-switcher.desktop")
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to open source file:", err)
			return
		}
		defer srcFile.Close()

		// Create the destination file for writing
		dstFile, err := os.Create(autostartFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to create destination file:", err)
			return
		}
		defer dstFile.Close()

		// Copy the contents from the source file to the destination file
		if _, err := io.Copy(dstFile, srcFile); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to copy file contents:", err)
			return
		}

		fmt.Println("Autostart file copied successfully.")
	}
}

func main() {
	currUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Running as: %+v\n", currUser)

	_ = requireCommand("pkexec")

	if len(os.Args) > 1 {
		callAsahiBless(os.Args[1:])
		return
	}

	if currUser.Uid == "0" {
		fmt.Fprintln(os.Stderr, "Should not run as root, exiting...")
		os.Exit(1)
	}

	setupAutostart(currUser.HomeDir)

	systray.Run(onReady, func() {})
}

func requestReboot() error {
	if os.Getenv("XDG_CURRENT_DESKTOP") == "KDE" {
		cmd := exec.Command("qdbus", "org.kde.ksmserver", "/KSMServer", "logout", "1", "1", "3")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
	if os.Getenv("XDG_CURRENT_DESKTOP") == "GNOME" {
		cmd := exec.Command("gnome-session-quit", "--reboot")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	// fallback
	cmd := exec.Command("pkexec", "reboot")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func callAsahiBless(args []string) {
	{
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to set boot volume:", err)
		}
	}
}

func rebootToMacOS(onlyOnce bool) {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to get executable path:", err)
	}
	args := []string{exePath, asahiBlessCmd, "--set-boot-macos", "-y"}
	if onlyOnce {
		args = append(args, "-n")
	}
	fmt.Printf("executing: pkexec %s\n", strings.Join(args, " "))
	cmd := exec.Command("pkexec", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to reboot to macOS:", err)
	}

	time.Sleep(1 * time.Second)

	{
		err := requestReboot()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to reboot to macOS:", err)
		}
	}
}

//go:embed asahi-reboot-switcher.png
var appIcon []byte

func onReady() {
	systray.SetTemplateIcon(appIcon, appIcon)
	// systray.SetTitle("Asahi Reboot Switcher")
	systray.SetTooltip("Restart in macOS (tray icon)")

	mReboot := systray.AddMenuItem("Restart in macOS...", "")
	mOnlyOnce := mReboot.AddSubMenuItem("Only once", "")
	mSetDefault := mReboot.AddSubMenuItem("Set as default", "")
	systray.AddSeparator()
	mQuitOrig := systray.AddMenuItem("Quit", "Quit application")

	for {
		select {
		case <-mReboot.ClickedCh:
			fmt.Println("clicked Reboot to macOS")
		case <-mOnlyOnce.ClickedCh:
			rebootToMacOS(true)
		case <-mSetDefault.ClickedCh:
			rebootToMacOS(false)
		case <-mQuitOrig.ClickedCh:
			fmt.Println("Quit")
			systray.Quit()
		}
	}
}
