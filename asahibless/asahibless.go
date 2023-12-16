package asahibless

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/nohajc/asahi-reboot-switcher/util"
)

var asahiBlessCmd = util.RequireCommand("asahi-bless")

func SetBootMacOS(onlyOnce bool) error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
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
		return fmt.Errorf("failed to set boot disk to macOS: %w", err)
	}
	return nil
}

func SetBoot(volIdx int) error {
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}
	args := []string{exePath, asahiBlessCmd, "--set-boot", fmt.Sprintf("%d", volIdx), "-y"}
	fmt.Printf("executing: pkexec %s\n", strings.Join(args, " "))
	cmd := exec.Command("pkexec", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to set boot disk: %w", err)
	}
	return nil
}

type Volume struct {
	Active    bool
	Idx       int
	PartNames []string
}

func (v *Volume) ShortName() string {
	partNames := []string{}
	for _, p := range v.PartNames {
		if !strings.Contains(p, "Data") {
			partNames = append(partNames, p)
		}
	}
	return strings.Join(partNames, ", ")
}

func ListVolumes() ([]Volume, error) {
	exePath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("failed to get executable path: %w", err)
	}
	cmd := exec.Command("pkexec", exePath, asahiBlessCmd, "--list-volumes")
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list volumes: %w", err)
	}
	lines := strings.Split(string(out), "\n")
	volumes := make([]Volume, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		vol := Volume{}
		if line[0] == '*' {
			vol.Active = true
			line = line[1:]
		}
		n, err := fmt.Sscanf(line, "%d) ", &vol.Idx)

		if err != nil {
			return nil, fmt.Errorf("failed to parse volume line %q: %w", line, err)
		}

		line = line[n+2:]
		vol.PartNames = strings.Split(line, ", ")
		volumes = append(volumes, vol)
	}
	return volumes, nil
}
