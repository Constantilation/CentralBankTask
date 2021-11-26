package domain

import (
	"CentralBankTask/internal/Utils"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Valute structure for valute info
type Valute struct {
	ID       string `xml:"ID,attr"`
	NumCode  int
	CharCode string
	Nominal  int
	Name     string
	Value    string
}

// ValCurs structure of central bank info in XML format
type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valute  []Valute `xml:"Valute"`
}

func (v *ValCurs) ReformatFile(url, filename string) error {
	err := Utils.DownloadFile(filename, url)
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

		newBody := strings.Replace(string(body), "<offers>", "", -1)
		newBody = strings.Replace(newBody, "</offers>", "", -1)
		err = xml.Unmarshal([]byte(newBody), v)

		if err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	}
}
