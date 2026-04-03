package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func DownloadFile(url, destDir string) error {
	fileName := filepath.Base(url)
	filePath := filepath.Join(destDir, fileName)

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	fmt.Println("Downloading", url)
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		_ = os.Remove(filePath)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		_ = os.Remove(filePath)
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("Downloading %s took %s\n", fileName, time.Since(start))
	return nil
}

func SequentialDownloader(urls []string, destDir string) error {

	for _, url := range urls {
		err := DownloadFile(url, destDir)
		if err != nil {
			fmt.Println("Error downloading... ", url, err)
		}
	}
	return nil
}

type Result struct {
	URL      string
	Filename string
	Size     int64
	Duration time.Duration
	Error    error
}

func ConcurrentDownloader(urls []string, concurrency int, destDir string) error {
	results := make(chan Result)

	err := os.MkdirAll(destDir, 0755)
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	limitter := make(chan struct{}, concurrency)

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()
			limitter <- struct{}{}
			defer func() { <-limitter }()

			start := time.Now()
			filename := filepath.Base(url)
			filePath := filepath.Join(destDir, filename)

			out, err := os.Create(filePath)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}
			defer out.Close()

			resp, err := http.Get(url)
			if err != nil {
				results <- Result{URL: url, Error: err}
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				results <- Result{URL: url, Error: fmt.Errorf("bad status: %s", resp.Status)}
				return
			}
			size, err := io.Copy(out, resp.Body)
			if err != nil {
				results <- Result{URL: url, Error: err}
				return
			}
			timeSince := time.Since(start)
			results <- Result{URL: url, Filename: filename, Size: size, Duration: timeSince, Error: nil}

		}(url)

	}
	go func() {
		wg.Wait()
		close(results)
	}()

	var totalSize int64
	var errors []error
	start := time.Now()

	for result := range results {
		if result.Error != nil {
			fmt.Printf("Error downloading %s: %s\n", result.URL, result.Error.Error())
			errors = append(errors, result.Error)
		} else {
			totalSize += result.Size
			fmt.Printf("Downloaded %s (%d bytes) in %s\n", result.Filename, result.Size, result.Duration)
		}
	}

	startedSince := time.Since(start)
	fmt.Printf("All downloads completed in %s, Total: %d bytes\n", startedSince, totalSize)
	if len(errors) > 0 {
		return fmt.Errorf("errors downloading: %+v", errors)
	}

	return nil
}

func main() {

	urls := []string{"https://go.dev/images/go-logo-white.svg", "https://t4.ftcdn.net/jpg/06/45/69/17/360_F_645691769_kDGC0KZcRNw8Zmag6oLNnfJ9Gflp17oo.jpg"}

	err := ConcurrentDownloader(urls, 3, "./downloads")
	if err != nil {
		fmt.Println(err)
	}
}
