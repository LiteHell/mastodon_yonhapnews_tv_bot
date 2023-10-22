package bot

import (
	"io"
	"net/http"
	"os"
)

// https://stackoverflow.com/a/33845771
func downloadFile(out *os.File, url string) (err error) {
	// Create the file
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
