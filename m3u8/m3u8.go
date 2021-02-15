package m3u8

import (
	"github.com/ilove91/91dl/m3u8/downloader"
)

// Download m3u8
func Download(url, title, destDir string, concurrency int) error {
	downloader, err := downloader.NewTask(destDir, url, title)
	if err != nil {
		return err
	}
	if err := downloader.Start(concurrency); err != nil {
		return err
	}
	return nil
}
