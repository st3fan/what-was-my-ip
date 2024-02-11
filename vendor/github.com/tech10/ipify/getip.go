package ipify

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/jpillora/backoff"
)

func getIp(url string) (string, error) {
	if url == "" {
		return "", errors.New("ip address retrieval URL is blank")
	}
	b := &backoff.Backoff{
		Jitter: true,
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", wrapError(err)
	}
	req.Header.Set("User-Agent", USER_AGENT)
	for tries := 0; tries < MAX_TRIES; tries++ {
		var resp *http.Response
		resp, err = client.Do(req)
		if err != nil {
			d := b.Duration()
			time.Sleep(d)
			continue
		}
		defer resp.Body.Close()
		var ip []byte
		ip, err = io.ReadAll(resp.Body)
		if err != nil {
			return "", wrapError(err)
		}
		if resp.StatusCode != http.StatusOK {
			description := fmt.Sprintf("Received invalid status code %d from ipify: %s. The service might be experiencing issues.", resp.StatusCode, resp.Status)
			return "", errors.New(description)
		}
		return string(ip), nil
	}
	return "", wrapError(err)
}

func wrapError(inner error) error {
	description := fmt.Sprintf("The request failed with error %s. This is most likely due to a networking error of some sort.", inner.Error())
	return errors.New(description)
}
