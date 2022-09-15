package star

import "github.com/xyy277/gallery/global/response"

type Authentication struct{}

func (authentication *Authentication) RequestToken(token string) (bool, response.Response) {
	var r bool
	var body response.Response

	return r, body
}

func (authentication *Authentication) RequestEnforce(path string, method string) bool {
	var r bool

	return r
}
