package onebot11

// https://github.com/botuniverse/onebot-11/blob/master/message/segment.md

// MessageSegment 表示一个消息段
type MessageSegment struct {
	// Type 消息段类型
	Type string `json:"type"`
	// Data 消息段数据
	Data interface{} `json:"data,omitempty"`
}

// MessageArray 表示消息段数组
type MessageArray []MessageSegment

// Text 创建文本消息段
func Text(text string) MessageSegment {
	return MessageSegment{
		Type: TextType,
		Data: TextData{
			Text: text,
		},
	}
}

// Face 创建 QQ 表情消息段
func Face(id string) MessageSegment {
	return MessageSegment{
		Type: FaceType,
		Data: FaceData{
			ID: id,
		},
	}
}

// Image 创建图片消息段
func Image(file string, imageType ...string) MessageSegment {
	data := ImageData{
		File: file,
	}
	if len(imageType) > 0 {
		data.Type = imageType[0]
	}
	return MessageSegment{
		Type: ImageType,
		Data: data,
	}
}

// At 创建@消息段
func At(qq string) MessageSegment {
	return MessageSegment{
		Type: AtType,
		Data: AtData{
			QQ: qq,
		},
	}
}

// Reply 创建回复消息段
func Reply(id string) MessageSegment {
	return MessageSegment{
		Type: ReplyType,
		Data: ReplyData{
			ID: id,
		},
	}
}

// Record 创建语音消息段
func Record(file string) MessageSegment {
	return MessageSegment{
		Type: RecordType,
		Data: RecordData{
			File: file,
		},
	}
}

// Video 创建视频消息段
func Video(file string) MessageSegment {
	return MessageSegment{
		Type: VideoType,
		Data: VideoData{
			File: file,
		},
	}
}

// Share 创建分享消息段
func Share(url string, title string, content string, image string) MessageSegment {
	return MessageSegment{
		Type: ShareType,
		Data: ShareData{
			URL:     url,
			Title:   title,
			Content: content,
			Image:   image,
		},
	}
}

// Poke 创建戳一戳(双击头像)消息段
func Poke(type_ string, id string) MessageSegment {
	return MessageSegment{
		Type: PokeType,
		Data: PokeData{
			Type: type_,
			ID:   id,
		},
	}
}
