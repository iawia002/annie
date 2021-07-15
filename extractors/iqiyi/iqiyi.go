package iqiyi

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/iawia002/annie/extractors/types"
	"github.com/iawia002/annie/parser"
	"github.com/iawia002/annie/request"
	"github.com/iawia002/annie/utils"
)

type iqiyi struct {
	Code string `json:"code"`
	Data struct {
		VP struct {
			Du  string `json:"du"`
			Tkl []struct {
				Vs []struct {
					Bid   int    `json:"bid"`
					Scrsz string `json:"scrsz"`
					Vsize int64  `json:"vsize"`
					Fs    []struct {
						L string `json:"l"`
						B int64  `json:"b"`
					} `json:"fs"`
				} `json:"vs"`
			} `json:"tkl"`
		} `json:"vp"`
	} `json:"data"`
	Msg string `json:"msg"`
}

type iqiyiURL struct {
	L string `json:"l"`
}

const iqiyiReferer = "https://www.iq.com"

func getMacID() string {
	var macID string
	chars := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "n", "m", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}
	size := len(chars)
	for i := 0; i < 32; i++ {
		macID += chars[rand.Intn(size)]
	}
	return macID
}

func getVF(params string) string {
	var suffix string
	for j := 0; j < 8; j++ {
		for k := 0; k < 4; k++ {
			var v8 int
			v4 := 13 * (66*k + 27*j) % 35
			if v4 >= 10 {
				v8 = v4 + 88
			} else {
				v8 = v4 + 49
			}
			suffix += string(rune(v8)) // string(97) -> "a"
		}
	}
	params += suffix

	return utils.Md5(params)
}

func getVPS(tvid, vid string) (*iqiyi, error) {
	t := time.Now().Unix() * 1000
	host := "http://cache.video.qiyi.com"
	params := fmt.Sprintf(
		"/vps?tvid=%s&vid=%s&v=0&qypid=%s_12&src=01012001010000000000&t=%d&k_tag=1&k_uid=%s&rs=1",
		tvid, vid, tvid, t, getMacID(),
	)
	vf := getVF(params)
	apiURL := fmt.Sprintf("%s%s&vf=%s", host, params, vf)
	info, err := request.Get(apiURL, iqiyiReferer, nil)
	if err != nil {
		return nil, err
	}
	data := new(iqiyi)
	if err := json.Unmarshal([]byte(info), data); err != nil {
		return nil, err
	}
	return data, nil
}

type extractor struct{}

// New returns a youtube extractor.
func New() types.Extractor {
	return &extractor{}
}

// Extract is the main function to extract the data.
func (e *extractor) Extract(url string, _ types.Options) ([]*types.Data, error) {
	html, err := request.Get(url, iqiyiReferer, map[string]string{
		"Accept-Language":   "zh-TW",
	})
	if err != nil {
		return nil, err
	}
	tvid := utils.MatchOneOf(
		url,
		`#curid=(.+)_`,
		`tvid=([^&]+)`,
	)
	if tvid == nil {
		tvid = utils.MatchOneOf(
			html,
			`data-player-tvid="([^"]+)"`,
			`param\['tvid'\]\s*=\s*"(.+?)"`,
			`"tvid":"(\d+)"`,
			`"tvId":(\d+)`,
		)
	}
	if tvid == nil || len(tvid) < 2 {
		return nil, types.ErrURLParseFailed
	}

	vid := utils.MatchOneOf(
		url,
		`#curid=.+_(.*)$`,
		`vid=([^&]+)`,
	)
	if vid == nil {
		vid = utils.MatchOneOf(
			html,
			`data-player-videoid="([^"]+)"`,
			`param\['vid'\]\s*=\s*"(.+?)"`,
			`"vid":"(\w+)"`,
		)
	}
	if vid == nil || len(vid) < 2 {
		return nil, types.ErrURLParseFailed
	}

	doc, err := parser.GetDoc(html)
	if err != nil {
		return nil, err
	}
	title := strings.TrimSpace(doc.Find("span#pageMetaTitle").First().Text())
	sub := utils.MatchOneOf(
			html,
			`"subTitle":"([^"]+)","isoDuration":`,
		)
	if sub != nil || len(sub) > 1 {
	        title += " "
	        title += sub[1]
       }
	if title == "" {
		title = doc.Find("title").Text()
	}
	videoDatas, err := getVPS(tvid[1], vid[1])
	if err != nil {
		return nil, err
	}
	if videoDatas.Code != "A00000" {
		return nil, fmt.Errorf("can't play this video: %s", videoDatas.Msg)
	}

	streams := make(map[string]*types.Stream)
	urlPrefix := videoDatas.Data.VP.Du
	for _, video := range videoDatas.Data.VP.Tkl[0].Vs {
		urls := make([]*types.Part, len(video.Fs))
		for index, v := range video.Fs {
			realURLData, err := request.Get(urlPrefix+v.L, iqiyiReferer, nil)
			if err != nil {
				return nil, err
			}
			var realURL iqiyiURL
			if err = json.Unmarshal([]byte(realURLData), &realURL); err != nil {
				return nil, err
			}
			_, ext, err := utils.GetNameAndExt(realURL.L)
			if err != nil {
				return nil, err
			}
			urls[index] = &types.Part{
				URL:  realURL.L,
				Size: v.B,
				Ext:  ext,
			}
		}
		streams[strconv.Itoa(video.Bid)] = &types.Stream{
			Parts:   urls,
			Size:    video.Vsize,
			Quality: video.Scrsz,
		}
	}

	return []*types.Data{
		{
			Site:    "爱奇艺 iqiyi.com",
			Title:   title,
			Type:    types.DataTypeVideo,
			Streams: streams,
			URL:     url,
		},
	}, nil
}
