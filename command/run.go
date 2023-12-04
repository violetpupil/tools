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
	u, err := url.Parse(c.roomURL)
	if err != nil {
		return
	}
	eTLDPO, err := publicsuffix.EffectiveTLDPlusOne(u.Hostname())
	if err != nil {
		return
	}
	siteID := strings.Split(eTLDPO, ".")[0]

	site, ok := olivetv.Sniff(siteID)
	if !ok {
		return
	}

	fmt.Println(site)
}
