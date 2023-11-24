package kernel

// Show represents an individual show.
type Show struct {
	Platform     string `json:"platform"`      // 网站id
	RoomID       string `json:"room_id"`       // 房间id
	StreamerName string `json:"streamer_name"` // 网站名
	SplitRule    string `json:"split_rule"`    // 视频分段规则
}
