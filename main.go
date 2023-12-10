package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"

	"fyne.io/systray"
)

func requireCommand(name string) {
	_, err := exec.LookPath(name)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Required command not found:", name)
		os.Exit(1)
	}
}

func main() {
	requireCommand("asahi-bless")
	requireCommand("pkexec")

	systray.Run(onReady, func() {})
}

func rebootToMacOS(onlyOnce bool) {
	args := []string{"pkexec", "asahi-bless", "--set-boot-macos"}
	if onlyOnce {
		args = append(args, "-n")
	}
	cmd := exec.Command("pkexec", args...)
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to reboot to macOS:", err)
		os.Exit(1)
	}
}

//go:embed macos-hdd.png
var appIcon []byte

func onReady() {
	systray.SetTemplateIcon(appIcon, appIcon)
	systray.SetTitle("Asahi Reboot Switcher")
	systray.SetTooltip("Asahi Reboot Switcher")

	mReboot := systray.AddMenuItem("Reboot to macOS...", "Reboot to macOS now")
	systray.AddSeparator()
	mOnlyOnce := systray.AddMenuItemCheckbox("Only once", "Reboot to macOS once", false)
	mSetDefault := systray.AddMenuItemCheckbox("Set as default", "Boot to macOS by default", true)
	systray.AddSeparator()
	mQuitOrig := systray.AddMenuItem("Quit", "Quit application")

	for {
		select {
		case <-mReboot.ClickedCh:
			fmt.Println("clicked Reboot to macOS")
			rebootToMacOS(mOnlyOnce.Checked())
		case <-mOnlyOnce.ClickedCh:
			if !mOnlyOnce.Checked() {
				mSetDefault.Uncheck()
			}
			mOnlyOnce.Check()
		case <-mSetDefault.ClickedCh:
			if !mSetDefault.Checked() {
				mOnlyOnce.Uncheck()
			}
			mSetDefault.Check()
		case <-mQuitOrig.ClickedCh:
			fmt.Println("Quit")
			systray.Quit()
		}
	}
}
