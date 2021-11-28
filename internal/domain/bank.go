package domain

import (
	errors "CentralBankTask/internal/Middleware/Error"
	"encoding/xml"
	"fmt"
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
	var path string
	path = "./files/"

	cmd := exec.Command("iconv", "-f", "cp1251", "-t", "utf8", path+filename, "-o", path+filename)
	_, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}

	cmd2 := exec.Command("sed", "-i", "s/ encoding=\"windows-1251\"//", path+filename)

	_, err = cmd2.CombinedOutput()
	if err != nil {
		return err
	}

	resp, err := os.Open(path + filename)

	if err != nil {
		fmt.Println("Невозможно найти файл или открыть")
		return err
	} else {

		defer func() {
			resp.Close()
			os.Remove(path + filename)
		}()
		body, err := ioutil.ReadAll(resp)
		if err != nil {
			return err
		}

		newBody := strings.Replace(string(body), ",", ".", -1)
		err = xml.Unmarshal([]byte(newBody), v)

		if err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	}
}

// DownloadFile func to download file by filepath and string
func DownloadFile(filepath string, url string) error {
	out, err := os.Create("files/" + filepath)
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
