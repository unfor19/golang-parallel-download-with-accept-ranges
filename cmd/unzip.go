/*
Copyright © 2021 Meir Gabay unfor19@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// unzipCmd represents the unzip command
var unzipCmd = &cobra.Command{
	Use:   "unzip",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		src, err := cmd.Flags().GetString("src")
		if err != nil {
			log.Fatal(err)
		}
		dst, err := cmd.Flags().GetString("dst")
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Unzipping", src, "to", dst)
		unzipError := unzipFunc(src, dst)
		if unzipError != nil {
			log.Fatal(unzipError)
		}
		log.Println("Successfully unzipped", src, "to", dst)
	},
}

func init() {
	rootCmd.AddCommand(unzipCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	unzipCmd.PersistentFlags().StringP("src", "s", "", "Source archive")
	unzipCmd.MarkPersistentFlagRequired("src")

	unzipCmd.PersistentFlags().StringP("dst", "d", "", "Target dir")
	unzipCmd.MarkPersistentFlagRequired("dst")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unzipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// Source: https://stackoverflow.com/a/24792688/5285732
func unzipFunc(src string, dest string) error {
	// src - zip file
	// dest -  auto creates target directory and extracts the files to it
	r, err := zip.OpenReader(src)
	if err != nil {
		log.Println("Failed to open source file", src)
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			log.Println("Failed to close file", src)
			panic(err)
		}
	}()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			log.Println("Failed to open output file", f.Name)
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				log.Println("Failed to close output file", f.Name)
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					log.Println("Failed to close file", f.Name())
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				log.Println("Failed to copy file", f.Name())
				log.Println(err)
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			log.Println("Failed to extract and write file", f.Name)
			return err
		}
	}

	return nil
}
