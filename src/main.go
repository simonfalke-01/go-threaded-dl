package main

import (
	"os"
	"strconv"
)

type Downloader struct {
	url     string
	threads int

	savePath string
}

func printHelp() {
	println("Usage: gtd [url] [threads] [save path]")
	println("Example: gtd https://example.com/file.zip 10 /home/user/Downloads/file.zip")
}

func IsValid(fp string) bool {
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

func checkValidDownloader(d Downloader) {
	if !IsValid(d.savePath) {
		println("Error: save path is not valid")
		printHelp()
		os.Exit(1)
	}
}

func main() {
	// basic cli stuff
	// parse args
	// create downloader
	// call download

	// parse args
	args := os.Args[1:]
	if len(args) == 1 && (args[0] == "help" || args[0] == "-h" || args[0] == "--help") {
		printHelp()
		os.Exit(0)
	} else if len(args) < 3 {
		println("Error: not enough arguments")
		printHelp()
		os.Exit(1)
	} else if len(args) > 3 {
		println("Error: too many arguments")
		printHelp()
		os.Exit(1)
	} else {
		url := args[0]

		threads, err := strconv.Atoi(args[1])
		if err != nil {
			println("Error: invalid thread count")
			printHelp()
			os.Exit(1)
		}

		savePath := args[2]

		downloader := Downloader{
			url:     url,
			threads: threads,

			savePath: savePath,
		}

		checkValidDownloader(downloader)

		println("Downloading", downloader.url, "to", downloader.savePath, "with", downloader.threads, "threads")

		download(downloader)
	}
}
