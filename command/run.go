package command

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/go-olive/olive/foundation/olivetv"
	"github.com/spf13/cobra"
	"golang.org/x/net/publicsuffix"
)

// run run 子命令
type run struct {
	roomURL string
}

func newRun() *cobra.Command {
	c := new(run)

	cc := &cobra.Command{
		Use:   "run",
		Short: "Start the olive engine.",
		Long:  `Start the olive engine.`,
		Run:   c.run,
	}

	cc.Flags().StringVarP(&c.roomURL, "url", "u", "", "room url")

	return cc
}

// run 执行函数
func (c *run) run(*cobra.Command, []string) {
	// 解析网站 id
	u, err := url.Parse(c.roomURL)
	if err != nil {
		return
	}
	eTLDPO, err := publicsuffix.EffectiveTLDPlusOne(u.Hostname())
	if err != nil {
		return
	}
	siteID := strings.Split(eTLDPO, ".")[0]

	// 选择网站
	site, ok := olivetv.Sniff(siteID)
	if !ok {
		return
	}

	// 创建直播对象
	tv, err := site.Permit(olivetv.RoomURL(c.roomURL))
	if err != nil {
		return
	}

	// TODO
	err = site.Snap(tv)
	if err != nil {
		return
	}
	fmt.Println(tv.StreamURL())
}
