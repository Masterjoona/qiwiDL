package main

import (
	"flag"
	"regexp"
	"strings"
	"sync"

	"github.com/AlecAivazis/survey/v2"
	"github.com/vbauerster/mpb/v8"
)

type file struct {
	name string
	url  string
	//size int64
}

func main() {
	var host string
	var concurrentDownloads int
	var qiwiFolder string
	var directory string

	flag.StringVar(&qiwiFolder, "folder", "", "Qiwi folder url")
	flag.IntVar(&concurrentDownloads, "concurrent", 3, "Number of concurrent downloads (how many files to download at once, this is not the amount of simultaneous connections to one file)")
	flag.StringVar(&directory, "directory", "", "Directory to download to")
	flag.StringVar(&host, "host", "https://qiwi.lol/", "When you go download a file and hover your mouse over the download button, you should see a link like https://texturepackguy.com/abcdefg1234567890. The host is the https://texturepackguy.com/ part.")
	flag.Parse()

	if qiwiFolder == "" || !strings.Contains(qiwiFolder, "qiwi.gg") {
		panic("Please provide a valid Qiwi folder url")
	}

	body := FetchFolder(qiwiFolder)

	fileUrls := regexp.MustCompile(`<a target="_blank" href="https:\/\/qiwi\.gg\/file\/(.*?)">`)
	fileMatches := fileUrls.FindAllStringSubmatch(body, -1)

	filenames := regexp.MustCompile(`<p>(.*?.{1,5})<\/p>`)
	filenameMatches := filenames.FindAllStringSubmatch(body, -1)

	extensions := regexp.MustCompile(`(\..{1,5})<\/p>`)
	extMatches := extensions.FindAllStringSubmatch(body, -1)

	/*
		fileSizes := regexp.MustCompile(`<div class="DownloadButton_FileSize__.*?">(.*?)<\/div>`)
		fileSizeMatches := fileSizes.FindAllStringSubmatch(body, -1)
	*/

	downloadURLs := make([]string, len(fileMatches))
	var options []string = make([]string, len(fileMatches))

	for i, match := range fileMatches {
		downloadURLs[i] = host + match[1] + extMatches[i][1]
		options[i] = filenameMatches[i][1]
	}

	var excludeFiles []string
	prompt := &survey.MultiSelect{
		Message:  "Select files to EXCLUDE:",
		Options:  options,
		PageSize: 10,
	}
	survey.AskOne(prompt, &excludeFiles)

	var actualDownloadURLs []file
	for i, url := range downloadURLs {
		if !Contains(excludeFiles, options[i]) {
			actualDownloadURLs = append(actualDownloadURLs, file{options[i], url /*ParseSize(fileSizeMatches[i][1])*/})
		}
	}
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, concurrentDownloads)

	p := mpb.New(mpb.WithWaitGroup(&wg))

	/*totalP := mpb.New()
	totalBar := totalP.AddBar(CalculateTotalSize(actualDownloadURLs),
	mpb.PrependDecorators(
		decor.Name("Total: ", decor.WC{W: 8, C: decor.DidentRight}),
	),
	mpb.AppendDecorators(
		decor.CountersKibiByte("% .2f / % .2f | ", decor.WCSyncWidth, decor.WC{W: 15}),
		decor.Percentage(decor.WC{W: 2}),
	))
	*/

	for index, file := range actualDownloadURLs {
		wg.Add(1)
		go DownloadFile(file.url, &wg, semaphore, file.name, p, index, len(actualDownloadURLs), directory /*totalBar*/)
	}

	wg.Wait()

}
