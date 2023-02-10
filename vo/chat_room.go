package vo

type ChatRoom struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Introduce   string `json:"introduce"`
	HeadImage   string `json:"head_image"`
	AddUserRole int64  `json:"add_user_role"`
	Del         bool   `json:"del"`
	GmtCreate   int64  `json:"gmt_create"`
	GmtModified int64  `json:"gmt_modified"`
}
