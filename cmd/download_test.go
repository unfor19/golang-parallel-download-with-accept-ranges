package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"testing"
)

var downloadUrl string = "https://github.com/unfor19/ops/releases/download/0.0.11rc/ops_0.0.11rc_linux_amd64"
var fileName string = getFileName(downloadUrl)

func DownloadTestsCleanup() {
	err := os.Remove(fileName)
	if err != nil {
		log.Println("Warning, did to remove existing file file")
	}
}

func TestDownload(t *testing.T) {
	DownloadTestsCleanup()
	cmd := exec.Command("ops", "download", "-u", downloadUrl)
	if runtime.GOOS == "windows" {
		outputFile, err := ioutil.TempFile("", ".*")
		if err != nil {
			log.Fatalf("Failed to create temp dir\n%s\n", err)
		}

		cmd.Stdout = os.NewFile(uintptr(syscall.Stdout), outputFile.Name())
		cmd.Stderr = os.NewFile(uintptr(syscall.Stderr), outputFile.Name())
	} else {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	log.Println("Downloading ...")
	err := cmd.Run()
	if err != nil {
		log.Fatalf("download file failed with:\n%s\n", err)
	}
}

func TestDownloadExists(t *testing.T) {
	DownloadTestsCleanup()
	cmd := exec.Command("ops", "download", "-u", downloadUrl)
	if runtime.GOOS != "windows" {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	err := cmd.Run()
	if err != nil {
		log.Fatalf("download first attempt failed with:\n%s\n", err)
	}

	err = cmd.Run()
	if err != nil {
		log.Println("Should not be able to download an existing file")
	} else {
		log.Fatalf("Should fail when attempting to download an existing file")
	}
}

func TestDownloadRemoveExists(t *testing.T) {
	DownloadTestsCleanup()
	cmd := exec.Command("ops", "download", "-u", downloadUrl)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("download first attempt failed with:\n%s\n", err)
	}

	cmd = exec.Command("ops", "download", "-u", downloadUrl, "--remove-existing")
	err = cmd.Run()
	if err != nil {
		log.Fatalf("using remove exists flag\n%s\n", err)
	}
}
