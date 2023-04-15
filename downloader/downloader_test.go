package downloader

import (
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	d := Downloader{
		Url:      "https://do-spaces-1.simonfalke.studio/Hello!",
		Threads:  10,
		SavePath: "Hello!",
	}
	Download(d)

	// check if file exists and delete it
	if _, err := os.Stat(d.SavePath); os.IsNotExist(err) {
		t.Errorf("File was not created")
	} else {
		err := os.Remove(d.SavePath)
		if err != nil {
			t.Errorf("Error deleting file: %v", err)
		}
	}
}
