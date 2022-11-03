/**
 * @Author: Hardews
 * @Date: 2022/10/30 16:20
**/

package tool

import (
	"net/http"
)

func Get(url string, query [][2]string) *http.Request {
	var req *http.Request

	req, _ = http.NewRequest(http.MethodGet, url, nil)

	q := req.URL.Query()
	for _, res := range query {
		q.Add(res[0], res[1])
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.99 Safari/537.36")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Connection", "keep-alive")

	return req
}
