package olivetv

import (
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// TV 直播间
type TV struct {
	SiteID string
	RoomID string
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
