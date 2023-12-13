package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/getlantern/systray"
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

func main() {
	_ = requireCommand("pkexec")

	if len(os.Args) > 1 {
		callAsahiBlessAndReboot(os.Args[1:])
		return
	}

	systray.Run(onReady, func() {})
}

func callAsahiBlessAndReboot(args []string) {
	{
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to set boot volume:", err)
		}
	}

	time.Sleep(1 * time.Second)

	{
		cmd := exec.Command("reboot")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed to reboot to macOS:", err)
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
}

//go:embed macos-hdd.png
var appIcon []byte

func onReady() {
	systray.SetTemplateIcon(appIcon, appIcon)
	// systray.SetTitle("Asahi Reboot Switcher")
	systray.SetTooltip("Asahi Reboot Switcher")

	mReboot := systray.AddMenuItem("Reboot to macOS...", "Reboot to macOS now")
	mOnlyOnce := mReboot.AddSubMenuItem("Only once", "Reboot to macOS once")
	mSetDefault := mReboot.AddSubMenuItem("Set as default", "Boot to macOS by default")
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
