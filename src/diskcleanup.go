package main

import (
	"flag"
	"fmt"
	"os/exec"
	"path/filepath"

	"golang.org/x/sys/windows/registry"
)

const (
	//The registry path containing the directories we want to search through
	REG_BASE = `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\VolumeCaches`
	//The base of the key name which we will set in the above directory
	TARGET_FLAG = "StateFlags"
	//The value of the key which indicates to clean the directory
	TARGET_VALUE = 2
	//The path to cleanmgr on the machine
	CLEAN_MGR = `C:\Windows\System32\cleanmgr.exe`
)

var (
	//The folders which we want to clean. These are the subfolders of REG_BASE
	TARGETS = []string{
		"Temporary Setup Files",
		"Update Cleanup",
	}
)

func main() {
	//Parse the -profile flag from command line arguments
	profPtr := flag.Int("profile", 123, "The cleanup profile number")
	flag.Parse()
	//Create the key value we need, which is in the form of StateFlagsNNNN, where NNNN is
	//on the range [0, 9999] and left padded by zeroes to four digits if the value is <1000
	//Example:
	//	profile=123 -> StateFlags0123
	targetString := fmt.Sprintf("%v%04d", TARGET_FLAG, *profPtr)

	for _, val := range TARGETS {
		targetReg := filepath.Join(REG_BASE, val)
		fmt.Printf("Opening registry for writing: [%v]\n", targetReg)
		k, err := registry.OpenKey(registry.LOCAL_MACHINE, targetReg, registry.WRITE)
		//This SHOULD never error out because these keys are default keys, but if it does,
		//we just move on and scratch our heads.
		if err != nil {
			fmt.Printf("Error opening registry [HKLM:%v]:\n%v\n", REG_BASE, err.Error())
			continue
		}

		fmt.Printf("Writing DWORD value [%v] = [%v]\n", targetString, TARGET_VALUE)
		err = k.SetDWordValue(targetString, TARGET_VALUE)
		if err != nil {
			fmt.Printf("Failed to write registry value [%v] = [%v]\n", targetString, TARGET_VALUE)
		}
	}

	//Run the disk cleanup utility with the /sagerun:N flag
	arguments := fmt.Sprintf("/sagerun:%v", *profPtr)
	fmt.Printf("Running command [%v %v]", CLEAN_MGR, arguments)
	var cmd = exec.Command(CLEAN_MGR, arguments)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Disk cleanup failed:\n%v", err.Error())
	}
}
