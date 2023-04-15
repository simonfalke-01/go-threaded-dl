/*
Copyright © 2023 simonfalke ziqibrandonli@gmail.com
*/

package cmd

import (
	"fmt"
	"github.com/simonfalke-01/go-threaded-dl/v2/downloader"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var (
	url      string
	threads  int
	savePath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gdl [url]",
	Short: "Multi-threaded content downloader written in Go",
	Long: `gdl is a multi-threaded content downloader written in Go. It is designed to be as fast as possible, while also being as easy to use.
It is a lightweight alternative to other downloaders, such as aria2c, wget, and curl. Think of this as a multi-threaded version of wget.`,

	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		downloadWithArgs(args)
	},
}

func getCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}

func isValid(fp string) bool {
	// Check if fp is a directory
	if stat, err := os.Stat(fp); err == nil {
		if stat.IsDir() {
			return false
		}
	}

	// Check if file already exists
	if _, err := os.Stat(fp); err == nil {
		return true
	}

	// Attempt to create it
	var d []byte
	if err := os.WriteFile(fp, d, 0644); err == nil {
		err := os.Remove(fp)
		if err != nil {
			return false
		} // And delete it
		return true
	}

	return false
}

func checkIfExists(fp string) bool {
	if _, err := os.Stat(fp); err == nil {
		return true
	}

	return false
}

func appendToFileName(fp string, toAppend string) string {
	withoutExt := fp[:len(fp)-len(filepath.Ext(fp))]
	ext := filepath.Ext(fp)

	return withoutExt + toAppend + ext
}

func appendTillNotExist(fp string) string {
	if !checkIfExists(fp) {
		return fp
	}

	for i := 1; true; i++ {
		if !checkIfExists(appendToFileName(fp, fmt.Sprintf(" (%d)", i))) {
			return appendToFileName(fp, fmt.Sprintf(" (%d)", i))
		}
	}

	return fp
}

func downloadWithArgs(args []string) {
	url = args[0]

	if !isValid(savePath) {
		if !(savePath == "") {
			fmt.Println("Invalid save path. Using current directory instead.")
		}

		savePath = filepath.Join(getCurrentDirectory(), filepath.Base(url))
	}

	savePath = appendTillNotExist(savePath)

	if stat, err := os.Stat(filepath.Dir(savePath)); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(savePath), 0666); err != nil {
			panic(err)
		}
	} else if !stat.IsDir() {
		panic("Save path's parent is not a directory")
	}

	fmt.Println("[*] Downloading", url, "to", savePath, "with", threads, "threads")
	downloader.Download(downloader.Downloader{
		Url:      url,
		Threads:  threads,
		SavePath: savePath,
	})
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&threads, "threads", "t", 10, "Number of threads to use")

	rootCmd.Flags().StringVarP(&savePath, "save-path", "o", "", "Path to save file to")
}
