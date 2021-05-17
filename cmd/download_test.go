package cmd

import (
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/pkg/errors"
)

var downloadUrl = "https://github.com/unfor19/ops/releases/download/0.0.11rc/ops_0.0.11rc_linux_amd64"
var fileName string = getFileName(downloadUrl)

func DownloadTestsCleanup() {
	os.Remove(fileName)
}

func TestDownload(t *testing.T) {
	DownloadTestsCleanup()
	if err := exec.Command("ops", "download", "-u", downloadUrl).Run(); err != nil {
		log.Print(errors.Wrap(err, "downloading a file").Error())
		os.Exit(1)
	}
}

func TestDownloadExists(t *testing.T) {
	DownloadTestsCleanup()
	if err := exec.Command("ops", "download", "-u", downloadUrl).Run(); err != nil {
		log.Print(errors.Wrap(err, "downloading first attempt").Error())
		os.Exit(1)
	}

	if err := exec.Command("ops", "download", "-u", downloadUrl).Run(); err != nil {
		log.Println("error downloading second attempt")
		os.Exit(0)
	}
}

func TestDownloadRemoveExists(t *testing.T) {
	DownloadTestsCleanup()
	if err := exec.Command("ops", "download", "-u", downloadUrl).Run(); err != nil {
		log.Print(errors.Wrap(err, "first attempt").Error())
		os.Exit(1)
	}

	if err := exec.Command("ops", "download", "-u", downloadUrl, "--remove-existing").Run(); err != nil {
		log.Print(errors.Wrap(err, "using remove exists flag").Error())
		os.Exit(1)
	}
}
