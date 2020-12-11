package douyin

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/iawia002/annie/extractors/types"
	"github.com/iawia002/annie/request"
	"github.com/iawia002/annie/utils"
)

// .video.play_addr.url_list
type data struct {
	ItemList []struct {
		Desc  string `json:"desc"`
		Video struct {
			PlayAddr struct {
				URLList []string `json:"url_list"`
			} `json:"play_addr"`
		} `json:"video"`
	} `json:"item_list"`
}

type extractor struct{}

// New returns a youtube extractor.
func New() types.Extractor {
	return &extractor{}
}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string, option types.Options) ([]*types.Data, error) {
	var err error
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.80 Mobile Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	reader := resp.Body
	doc, _ := htmlquery.Parse(reader)
	realURLs := htmlquery.FindOne(doc, "//a/@href")
	if realURLs == nil {
		return nil, types.ErrURLParseFailed
	}
	realURL := htmlquery.InnerText(realURLs)

	if err != nil {
		return nil, err
	}

	videoIDs := utils.MatchOneOf(realURL, `/video/(\d+)`)
	if len(videoIDs) == 0 {
		return nil, errors.New("unable to get video ID")
	}
	videoID := videoIDs[1]
	apiDataString, err := request.Get(
		fmt.Sprintf("https://www.iesdouyin.com/web/api/v2/aweme/iteminfo/?item_ids=%s", videoID),
		realURL, nil,
	)
	if err != nil {
		return nil, err
	}

	var apiData data
	if err = json.Unmarshal([]byte(apiDataString), &apiData); err != nil {
		return nil, err
	}
	awemeURL := apiData.ItemList[0].Video.PlayAddr.URLList[0]
	awemeURL = strings.Replace(awemeURL, "/playwm/", "/play/", 1)
	videoReq, err := http.NewRequest("GET", awemeURL, nil)
	if err != nil {
		return nil, err
	}
	videoReq.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.80 Mobile Safari/537.36")

	videoResp, err := client.Do(videoReq)
	if err != nil {
		return nil, err
	}
	videoReader := videoResp.Body
	videoDoc, _ := htmlquery.Parse(videoReader)
	videoURLs := htmlquery.FindOne(videoDoc, "//a/@href")
	if videoURLs == nil {
		return nil, types.ErrURLParseFailed
	}
	videoURL := htmlquery.InnerText(videoURLs)
	size, err := request.Size(videoURL, url)
	urlData := &types.Part{
		URL:  videoURL,
		Size: size,
		Ext:  "mp4",
	}
	streams := map[string]*types.Stream{
		"default": {
			Parts: []*types.Part{urlData},
			Size:  size,
		},
	}

	return []*types.Data{
		{
			Site:    "抖音 douyin.com",
			Title:   apiData.ItemList[0].Desc,
			Type:    types.DataTypeVideo,
			Streams: streams,
			URL:     realURL,
		},
	}, nil
}
