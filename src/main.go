package main

type Downloader struct {
	url     string
	threads int

	savePath string
}

func main() {
	download("https://do-spaces-1.simonfalke.studio/file.txt", 10, "/Users/brandonli/file.txt")
}
