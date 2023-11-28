package olivetv

func init() {
	registerSite("douyin", &douyin{})
}

type douyin struct {
	base
}

func (this *douyin) Name() string {
	return "抖音"
}
