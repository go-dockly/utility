package xclient

import (
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

// Do sends the request to the specified rest path and unmarshals the response into the
// desired results interface{} if not provided as null and under the condition that
// the received status response is as expected
func (cli *Client) Do(method, path string, params io.Reader, result interface{}) (actualStatusCode int, err error) {
	url := fmt.Sprintf("%s/%s", cli.baseURL, path)

	req, err := cli.assembleRequest(method, url, params)
	if err != nil {
		return 0, err
	}

	res, err := cli.http.Do(req)
	if err != nil {
		return 0, errors.Wrap(err, "request failed")
	}

	for i := 0; i < cli.config.MaxRetry; i++ {
		switch res.StatusCode {
		case http.StatusInternalServerError:
			err = cli.handleBackoff(i)
			if err != nil {
				return 0, err
			}

			continue
		case http.StatusTooManyRequests:
			err = cli.handleBackoff(i)
			if err != nil {
				return 0, err
			}

			continue
		default:
			break
		}

		// 	check in case we are not interested in the response body
		if result != nil {
			err = cli.readResponse(res.Body, result)
		}

	}

	res.Body.Close()
	return res.StatusCode, nil
}
