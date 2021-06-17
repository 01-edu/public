package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func fatalln(a ...interface{}) {
	fmt.Println(a...)
	os.Exit(1)
}

func expect(target, err error, action string) {
	if err != nil && err != target && !errors.Is(err, target) {
		fatalln("Failed to", action, ":", err)
	}
}

func renderTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetColumnSeparator(" ")
	table.SetCenterSeparator(" ")
	table.SetRowSeparator(" ")
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetHeader(data[0])
	for _, row := range data[1:] {
		table.Append(row)
	}
	table.Render()
}

// run returns the output of the command
// In case of error, display the description, the command and the error
func run(description, name string, arg ...string) []byte {
	b, err := exec.Command(name, arg...).CombinedOutput()
	if err != nil {
		fmt.Println("Failed to", description)
		fmt.Printf("%s %#v\n", name, arg)
		if _, ok := err.(*exec.ExitError); ok {
			fatalln(string(b))
		}
		fatalln(err)
	}
	return b
}

type (
	size       int
	stringTrim string

	device struct {
		Path    string     // path to the device node
		RO      bool       // read-only device
		RM      bool       // removable device
		Hotplug bool       // removable or hotplug device (usb, pcmcia, ...)
		Model   stringTrim // device identifier
		Size    size       // size of the device
		Type    string     // device type
		Tran    string     // device transport type
		Vendor  stringTrim // device vendor
	}
)

func (s stringTrim) String() string {
	return strings.TrimSpace(string(s))
}

// format size with SI prefix
func (s size) String() string {
	f := float64(s)
	if f < 0 {
		f = 0
	}
	unit := " kMGTPEZY"
	for f >= 1000 {
		unit = unit[1:]
		f /= 1000
	}
	var prec int
	if unit[0] != ' ' && math.Round(f) < 10 {
		prec = 1
	}
	return fmt.Sprintf("%.*f %cB", prec, f, unit[0])
}

func main() {
	// Usage
	fmt.Println("While using this formatting tool:")
	fmt.Println("  - do not plug/unplug any drive")
	fmt.Println("  - make sure the drive is not in use")
	fmt.Print("Press ENTER to continue")
	bufio.NewScanner(os.Stdin).Scan()

	// Get a list of the block devices
	var data struct {
		Blockdevices []device
	}
	b := run("list the block devices",
		"lsblk",
		"--all",
		"--bytes",
		"--nodeps",
		"--json",
		"--output", "path,ro,rm,hotplug,model,size,type,tran,vendor",
	)
	expect(nil, json.Unmarshal(b, &data), "list the block devices")

	// Filter the block devices to get only the USB flash drives
	var devices []device
	for _, device := range data.Blockdevices {
		if !device.RO && device.RM && device.Hotplug && device.Type == "disk" && device.Tran == "usb" {
			devices = append(devices, device)
		}
	}

	// If no USB flash drive is found, exit the program with an explanation
	if len(devices) == 0 {
		fmt.Println("No available USB devices were found.")
		if len(data.Blockdevices) > 0 {
			fmt.Println("Here is a list of the block devices found:")
			table := [][]string{{"Path", "RO", "RM", "Hotplug", "Model", "Size", "Type", "Tran", "Vendor"}}
			for _, device := range data.Blockdevices {
				table = append(table, []string{
					device.Path,
					strconv.FormatBool(device.RO),
					strconv.FormatBool(device.RM),
					strconv.FormatBool(device.Hotplug),
					device.Model.String(),
					device.Size.String(),
					device.Type,
					device.Tran,
					device.Vendor.String(),
				})
			}
			renderTable(table)
		}
		return
	}

	// Display the found USB flash drives
	table := [][]string{{"Number", "Vendor", "Model", "Size"}}
	for i, device := range devices {
		table = append(table, []string{
			strconv.Itoa(i),
			device.Vendor.String(),
			device.Model.String(),
			device.Size.String(),
		})
	}
	renderTable(table)

	// Select the USB flash drive to format
	fmt.Println(`Enter the number of the disk you want to format (or "exit"):`)
	var choice string
	fmt.Scanln(&choice)
	if strings.ToLower(choice) == "exit" {
		return
	}
	var nb int
	nb, err := strconv.Atoi(choice)
	if err != nil || nb < 0 || nb >= len(devices) {
		fatalln("Wrong disk number:", choice)
	}
	device := devices[nb]

	// Get a list of the USB flash drive mount points
	b, err = os.ReadFile("/proc/mounts")
	expect(nil, err, "list the mount points")
	lines := strings.Split(string(b), "\n")
	var mountPoints []string
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		deviceName := fields[0]
		mountPoint := fields[1]
		if strings.Contains(deviceName, device.Path) {
			mountPoints = append(mountPoints, mountPoint)
		}
	}

	// Unmount the USB flash drive
	if len(mountPoints) > 0 {
		mountPointOccurrences := map[string]int{}
		for _, mountPoint := range mountPoints {
			mountPoint = mountPoint[1:] // remove leading '/'
			parts := strings.Split(mountPoint, "/")
			for i := 1; i <= len(parts); i++ {
				mountPointOccurrences[strings.Join(parts[:i], "/")]++
			}
		}
		for mountPoint, occurrences := range mountPointOccurrences {
			parts := strings.Split(mountPoint, "/")
			for i := len(parts) - 1; i > 0; i-- {
				parent := strings.Join(parts[:i], "/")
				if occurrences < mountPointOccurrences[parent] {
					delete(mountPointOccurrences, mountPoint)
				} else {
					delete(mountPointOccurrences, parent)
				}
			}
		}
		for mountPoint := range mountPointOccurrences {
			mountPoint = "/" + mountPoint // put back leading '/'
			fmt.Print("The selected device is mounted on ", mountPoint, ". Trying to unmount it... ")
			run("umount the selected device", "umount", "--recursive", mountPoint)
			fmt.Println("done")
		}
	}

	// Format USB flash drive
	fmt.Print("Formatting... ")
	run("erase disk data", "wipefs", "--all", device.Path)
	run("erase disk data", "sgdisk", "--zap-all", device.Path)
	run("create partition table", "sgdisk", "--largest-new", "0", device.Path)
	run("create partition table", "sgdisk", "--change-name", "1:01-home", device.Path)
	run("inform the OS of partition table changes", "partx", "--update", device.Path)
	run("format partition", "mkfs.f2fs", "-f", device.Path+"1")
	fmt.Println("done")
	fmt.Println(device.Vendor, device.Model, device.Size, "has been formatted, logout and login to use it")
}
