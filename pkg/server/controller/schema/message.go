package schema

type JoinRequest struct {
	Type    string      `json:"type"`
	Message JoinMessage `json:"message"`
}

type JoinMessage struct {
	ID      string `json:"id"`
	IconUrl string `json:"icon_url"`
	Power   int    `json:"power"`
	Weight  int    `json:"weight"`
	Volume  int    `json:"volume"`
	Cd      int    `json:"cd"`
}

type LeaveRequest struct {
	Type    string       `json:"type"`
	Message LeaveMessage `json:"message"`
}

type LeaveMessage struct {
	ID string `json:"id"`
}

type StartRequest struct {
	Type string `json:"type"`
}

type ActionRequest struct {
	Type    string        `json:"type"`
	Message ActionMessage `json:"message"`
}

type ActionMessage struct {
	ID        string    `json:"id"`
	Angle     []float32 `json:"angle"`
	PullPower int       `json:"pull_power"`
}
