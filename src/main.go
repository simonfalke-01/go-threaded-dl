package main

type Downloader struct {
	url     string
	threads int

	savePath string
}

func main() {
	download("https://sgp1.digitaloceanspaces.com/do-spaces-1/paper.zip", 10, "paper.zip")
}
