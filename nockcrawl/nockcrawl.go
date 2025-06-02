package nockcrawl

import (
	"nock/httputils"
)

const Reset = "\033[0m"

type NockCrawl struct {
	opt    *OptionsCrawl
	client *httputils.HTTPClient
	Href   string
}
