package domain

import (
	errors "CentralBankTask/inter/Middleware/Error"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

// Valute structure for valute info
type Valute struct {
	Name  string
	Value float32
}

// ValCurs structure of central bank info in XML format
type ValCurs struct {
	Date   string   `xml:"Date,attr"`
	Valute []Valute `xml:"Valute"`
}

type ValuteValue struct {
	Date   time.Time `json:"Date"`
	Valute Valute    `json:"Valute"`
}

type DateInterval struct {
	BegD time.Time
	EndD time.Time
}

type DownloadInterval struct {
	DateSlice []time.Time
}

func (v *ValCurs) ReformatFile(url, filename string) error {
	err := DownloadFile(filename, url)
	if err != nil {
		return err
	}

	cmd := exec.Command("iconv", "-f", "cp1251", "-t", "utf8", filename, "-o", filename)
	_, err = cmd.CombinedOutput()

	cmd2 := exec.Command("sed", "-i", "s/ encoding=\"windows-1251\"//", filename)

	_, err = cmd2.CombinedOutput()

	resp, err := os.Open(filename)

	if err != nil {
		return err
	} else {

		defer func() {
			resp.Close()
			os.Remove(filename)
		}()
		body, err := ioutil.ReadAll(resp)
		if err != nil {
			return err
		}

		newBody := strings.Replace(string(body), ",", ".", -1)
		err = xml.Unmarshal([]byte(newBody), v)

		if err != nil {
			return err
		}

		return nil
	}
}

// DownloadFile func to download file by filepath and string
func DownloadFile(filename string, url string) error {
	out, err := os.Create("./" + filename)
	if err != nil {
		return &errors.Errors{
			Alias: errors.ErrCreate,
		}
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return &errors.Errors{
			Alias: errors.ErrGet,
		}
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return &errors.Errors{
			Alias: errors.ErrCopy,
		}
	}

	return nil
}
