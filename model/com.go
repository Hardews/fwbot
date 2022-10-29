/**
 * @Author: Hardews
 * @Date: 2022/10/30 0:52
**/

package model

// 事件相关结构体，详情见https://docs.go-cqhttp.org/event

type Com struct {
	Time     int64  `json:"time,omitempty"`
	SelfId   int64  `json:"self_id,omitempty"`
	PostType string `json:"post_type,omitempty"`
}

type Message struct {
	Com
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageId   int32  `json:"message_id"`
	UserId      int64  `json:"user_id"`
	Messages    string `json:"message"`
	RawMessage  string `json:"raw_message"`
	NoticeType  string `json:"notice_type"`
	Font        string `json:"font"`
	GroupId     int64  `json:"group_id"`
	Sender
}

type Sender struct {
	UserId   int64  `json:"user_id"`
	NickName string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
}
