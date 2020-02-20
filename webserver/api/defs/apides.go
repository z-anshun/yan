package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
} //用户的数据模型

type UserInfo struct {
	Id int `json:"id"`
}

type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
} //注销

type SignedIn struct {
	Success    bool   `json:"success"`
	SesssionId string `json:"sesssion_id"`
} //登录

type VideosInfo struct {
	Vides []*VideoInfo `json:"vides"`
}
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtime string
} //视频的模型

type Comments struct {
	Comments []*Comment `json:"comments"`
}
type Comment struct {
	Id      string
	VideoId string
	Author  string
	Content string
	Next    *Comments
} //评论模型

type User struct {
	Id        int
	LoginName string
	Pwd       string
}

type SimpleSession struct {
	Username string //登录名
	TTL      int64
	Sess     string
} //session模型

type NewVideo struct {
	AuthorId int    `json:"author_id"`
	Name     string `json:"name"`
}

type NewComment struct {
	AuthorId int    `json:"author_id"`
	Content  string `json:"content"`
	Vvid     int    `json:"vvid"`
}

type VideoAchieve struct {
	Vid      string `json:"vid"`
	Collect  int    `json:"collect"`
	Favour   int    `json:"favour"`
	Transmit int    `json:"transmit"`
}

//视频点赞这些的模型

type UserSelf struct {
	Name     string `json:"name"`
	Vid      string `json:"vid"`
	Collect  int    `json:"collect"`
	Favour   int    `json:"favour"`
	Transmit int    `json:"transmit"`
}

//用户是否点赞模型

type UserAchieve struct {
	Name     string `json:"name"`
	Favour   int    `json:"favour"`
	Fan      int    `json:"fan"`
	Transmit int    `json:"transmit"`
}

type Atten struct {
	AuthorName string `json:"author_name"`
}
type Attens struct {
	AttensAuthor []*Atten `json:"attens_author"`
}
