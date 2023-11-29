package olivetv

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/imroc/req/v3"
	"github.com/violetpupil/gos/std/strings"
)

func init() {
	registerSite("tiktok", &tiktok{})
}

type tiktok struct {
	base
}

func (this *tiktok) Name() string {
	return "tiktok"
}

func (this *tiktok) RoomID(roomURL RoomURL) string {
	s := strings.TrimSuffix(string(roomURL), "/live")
	s = strings.SplitLast(s, "@")
	return ""
}

func (this *tiktok) RoomURL(roomID string) RoomURL {
	return RoomURL(fmt.Sprintf("https://www.tiktok.com/@%s/live", roomID))
}

type TiktokAutoGenerated struct {
	Data struct {
		Status int    `json:"status"`
		Title  string `json:"title"`
		Owner  struct {
			Nickname string `json:"nickname"`
		} `json:"owner"`
		StreamURL struct {
			RtmpPullURL string `json:"rtmp_pull_url"`
			FlvPullURL  struct {
				FullHd1 string `json:"FULL_HD1"`
				Hd1     string `json:"HD1"`
				Sd1     string `json:"SD1"`
				Sd2     string `json:"SD2"`
			} `json:"flv_pull_url"`
			HlsPullURL string `json:"hls_pull_url"`
		} `json:"stream_url"`
	} `json:"data"`
}

// set 抓取直播间信息
func (this *tiktok) set(tv *TV) error {
	// 抓取直播间页面
	liveURL := string(this.RoomURL(tv.RoomID))
	c := req.C()
	if tv.proxy != "" {
		c.SetProxyURL(tv.proxy)
	}
	resp, err := c.R().SetHeaders(map[string]string{
		"Referer":    liveURL,
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
	}).Get(liveURL)
	if err != nil {
		return err
	}
	if resp.GetStatusCode() != http.StatusOK {
		return errors.New("tiktok: network not available " + resp.GetStatus())
	}
	// 解析 tiktok room id
	res := regexp.MustCompile(`room_id=(.*?)"/>`).FindStringSubmatch(resp.String())
	if len(res) < 2 {
		return errors.New("tiktok: failed to find roomID")
	}
	roomID := res[1]

	// 抓取直播间信息
	var ag TiktokAutoGenerated
	_, err = req.C().R().
		SetHeaders(map[string]string{
			"Referer":    "https://www.tiktok.com/",
			"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36",
		}).
		SetQueryParam("room_id", roomID).
		SetSuccessResult(&ag).
		Get("https://webcast.tiktok.com/webcast/room/info/?aid=1988")
	if err != nil {
		return err
	}
	// 直播标题
	data := ag.Data
	if data.Title != "" {
		tv.roomName = data.Title
	} else {
		tv.roomName = tv.RoomID
	}
	// 拉流地址
	candi := []string{
		data.StreamURL.RtmpPullURL,
		data.StreamURL.FlvPullURL.FullHd1,
		data.StreamURL.FlvPullURL.Hd1,
		data.StreamURL.FlvPullURL.Sd1,
		data.StreamURL.FlvPullURL.Sd2,
	}
	for _, v := range candi {
		if v != "" {
			tv.roomOn = true
			tv.streamURL = v
			break
		}
	}
	return nil
}