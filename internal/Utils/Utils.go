package Utils

import (
	"io"
	"net/http"
	"os"
)

// DownloadFile func to download file by filepath and string
func DownloadFile(filepath string, url string) error {

	// Create the file
	out, err := os.Create("files/" + filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
