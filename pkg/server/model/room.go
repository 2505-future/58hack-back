package model

type Room struct {
	ID        string `json:"id"`
	HostID    string `json:"host_id"`
	Name      string `json:"name"`
	Capacity  int    `json:"capacity"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
