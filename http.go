package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

func DownloadFile(url string, wg *sync.WaitGroup, semaphore chan struct{}, fileName string, progress *mpb.Progress, index int, amountOfUrls int, directory string /*, totalProgess *mpb.Bar*/) {
	defer wg.Done()
	defer func() { <-semaphore }()
	semaphore <- struct{}{}
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error downloading %s: %s\n", url, err)
		return
	}
	defer resp.Body.Close()

	if directory != "" {
		if _, err := os.Stat(directory); os.IsNotExist(err) {
			os.Mkdir(directory, 0755)
		}
		fileName = directory + "/" + fileName
	}
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file %s: %s\n", fileName, err)
		return
	}
	defer file.Close()
	bar := progress.AddBar(resp.ContentLength,
		mpb.PrependDecorators(
			decor.Name(fmt.Sprintf("[%d/%d] %s ", index+1, amountOfUrls, fileName), decor.WC{W: len(fmt.Sprintf("%d/%d", amountOfUrls, amountOfUrls)) + 1, C: decor.DidentRight}),
			decor.CountersKibiByte("% .2f / % .2f | ", decor.WCSyncWidth, decor.WC{W: 15}),
			decor.Percentage(decor.WC{W: 2}),
		),
		mpb.AppendDecorators(
			decor.EwmaSpeed(decor.SizeB1024(0), "% .2f", 30),
			decor.EwmaETA(decor.ET_STYLE_GO, 100, decor.WC{W: 4}),
		),
	)
	reader := bar.ProxyReader(resp.Body)
	//anotherReader := totalProgess.ProxyReader(reader)
	_, err = io.Copy(file, reader)
	if err != nil {
		fmt.Printf("Error copying data for %s: %s\n", url, err)
		return
	}
}

func FetchFolder(qiwiFolder string) string {
	resp, err := http.Get(qiwiFolder)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return string(body)
}
