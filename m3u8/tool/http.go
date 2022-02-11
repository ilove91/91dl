package tool

import (
	"fmt"
	"io"

	"github.com/ilove91/91dl/utils"
)

func Get(url string) (io.ReadCloser, error) {
	c := utils.GetNewHttpClient(60)
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", resp.StatusCode)
	}
	return resp.Body, nil
}
