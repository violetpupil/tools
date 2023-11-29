package olivetv

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

var (
	ErrNotSupported = errors.New("streamer not supported")
	ErrSiteInvalid  = errors.New("site invalid")
)

// TV 直播间
type TV struct {
	SiteID   string
	SiteName string
	RoomID   string

	roomName  string // 直播标题
	streamURL string // 拉流地址
	roomOn    bool   // 是否在直播

	// 获取信息所需参数
	proxy string
}

func New(siteID, roomID string) (*TV, error) {
	site, ok := Sniff(siteID)
	if !ok {
		return nil, ErrNotSupported
	}

	t := &TV{
		SiteID:   siteID,
		SiteName: site.Name(),
		RoomID:   roomID,
	}
	return t, nil
}

func NewWithURL(roomURL string) (*TV, error) {
	u := RoomURL(roomURL)
	t, err := u.Stream()
	if err != nil {
		err = fmt.Errorf("%s (err msg = %s)", ErrNotSupported, err)
		return nil, err
	}

	return t, nil
}

// RoomURL 直播间地址
type RoomURL string

// SiteID 从直播间地址获取网站域，作为网站id
// 解析失败返回空
func (this RoomURL) SiteID() string {
	u, err := url.Parse(string(this))
	if err != nil {
		return ""
	}
	eTLDPO, err := publicsuffix.EffectiveTLDPlusOne(u.Hostname())
	if err != nil {
		return ""
	}
	siteID := strings.Split(eTLDPO, ".")[0]
	return siteID
}

// Stream 初始化直播间对象
func (this RoomURL) Stream() (*TV, error) {
	siteID := this.SiteID()
	site, ok := Sniff(siteID)
	if !ok {
		return nil, ErrSiteInvalid
	}
	siteName := site.Name()
	roomID := site.RoomID(this)
	return &TV{SiteID: siteID, SiteName: siteName, RoomID: roomID}, nil
}
