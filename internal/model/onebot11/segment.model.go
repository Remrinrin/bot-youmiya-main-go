package onebot11

// 消息段类型常量
const (
	// Text 纯文本
	TextType = "text"
	// Face QQ 表情
	FaceType = "face"
	// Image 图片
	ImageType = "image"
	// Record 语音
	RecordType = "record"
	// Video 视频
	VideoType = "video"
	// At @某人
	AtType = "at"
	// RPS 猜拳魔法表情
	RPSType = "rps"
	// Dice 掷骰子魔法表情
	DiceType = "dice"
	// Shake 窗口抖动（戳一戳）
	ShakeType = "shake"
	// Poke 戳一戳
	PokeType = "poke"
	// Anonymous 匿名发消息
	AnonymousType = "anonymous"
	// Share 链接分享
	ShareType = "share"
	// Contact 推荐好友/群
	ContactType = "contact"
	// Location 位置
	LocationType = "location"
	// Music 音乐分享
	MusicType = "music"
	// Reply 回复
	ReplyType = "reply"
	// Forward 合并转发
	ForwardType = "forward"
	// Node 合并转发节点
	NodeType = "node"
	// XML XML 消息
	XMLType = "xml"
	// JSON JSON 消息
	JSONType = "json"
)

// TextData 文本消息段数据
type TextData struct {
	Text string `json:"text"`
}

// FaceData QQ 表情消息段数据
type FaceData struct {
	ID string `json:"id"`
}

// ImageData 图片消息段数据
type ImageData struct {
	File    string `json:"file"`
	Type    string `json:"type,omitempty"` // flash 表示闪照
	URL     string `json:"url,omitempty"`
	Cache   int    `json:"cache,omitempty"`   // 0/1
	Proxy   int    `json:"proxy,omitempty"`   // 0/1
	Timeout int    `json:"timeout,omitempty"` // 单位秒
}

// RecordData 语音消息段数据
type RecordData struct {
	File    string `json:"file"`
	Magic   int    `json:"magic,omitempty"` // 0/1
	URL     string `json:"url,omitempty"`
	Cache   int    `json:"cache,omitempty"`   // 0/1
	Proxy   int    `json:"proxy,omitempty"`   // 0/1
	Timeout int    `json:"timeout,omitempty"` // 单位秒
}

// VideoData 视频消息段数据
type VideoData struct {
	File    string `json:"file"`
	URL     string `json:"url,omitempty"`
	Cache   int    `json:"cache,omitempty"`   // 0/1
	Proxy   int    `json:"proxy,omitempty"`   // 0/1
	Timeout int    `json:"timeout,omitempty"` // 单位秒
}

// AtData @消息段数据
type AtData struct {
	QQ string `json:"qq"`
}

// RPSData 猜拳魔法表情消息段数据
type RPSData struct{}

// DiceData 掷骰子魔法表情消息段数据
type DiceData struct{}

// ShakeData 窗口抖动（戳一戳）消息段数据
type ShakeData struct{}

// PokeData 戳一戳消息段数据
type PokeData struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// AnonymousData 匿名发消息段数据
type AnonymousData struct {
	Ignore bool `json:"ignore,omitempty"` // 0/1
}

// ShareData 链接分享消息段数据
type ShareData struct {
	URL     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content,omitempty"`
	Image   string `json:"image,omitempty"`
}

// ContactData 推荐好友/群消息段数据
type ContactData struct {
	Type string `json:"type"` // qq/group
	ID   string `json:"id"`
}

// LocationData 位置消息段数据
type LocationData struct {
	Lat     string `json:"lat"`
	Lon     string `json:"lon"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

// MusicData 音乐分享消息段数据
type MusicData struct {
	Type    string `json:"type"`              // qq/163/xm/custom
	ID      string `json:"id,omitempty"`      // 歌曲 ID
	URL     string `json:"url,omitempty"`     // 点击后跳转目标 URL
	Audio   string `json:"audio,omitempty"`   // 音乐 URL
	Title   string `json:"title,omitempty"`   // 标题
	Content string `json:"content,omitempty"` // 内容描述
	Image   string `json:"image,omitempty"`   // 图片 URL
}

// ReplyData 回复消息段数据
type ReplyData struct {
	ID string `json:"id"`
}

// ForwardData 合并转发消息段数据
type ForwardData struct {
	ID string `json:"id"`
}

// NodeData 合并转发节点消息段数据
type NodeData struct {
	ID string `json:"id"`
}

// CustomNodeData 合并转发自定义节点消息段数据
type CustomNodeData struct {
	UserID   string      `json:"user_id"`
	Nickname string      `json:"nickname"`
	Content  interface{} `json:"content"` // 支持 message 数据类型
}

// XMLData XML 消息段数据
type XMLData struct {
	Data string `json:"data"`
}

// JSONData JSON 消息段数据
type JSONData struct {
	Data string `json:"data"`
}
