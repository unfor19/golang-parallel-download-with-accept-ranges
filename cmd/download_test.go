package cmd

import (
	"log"
	"os"
	"os/exec"
	"testing"
)

var downloadUrl string = "https://github.com/unfor19/ops/releases/download/0.0.11rc/ops_0.0.11rc_linux_amd64"
var fileName string = getFileName(downloadUrl)

func DownloadTestsCleanup() {
	os.Remove(fileName)
}

func TestDownload(t *testing.T) {
	DownloadTestsCleanup()
	cmd := exec.Command("ops", "download", "-u", downloadUrl)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("download file failed with:\n%s\n", err)
	}
}

func TestDownloadExists(t *testing.T) {
	DownloadTestsCleanup()
	cmd := exec.Command("ops", "download", "-u", downloadUrl)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
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
