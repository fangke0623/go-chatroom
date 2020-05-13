package exception

type Error struct {
	ErrorCode int `json:"errorCode"`

	ErrorMsg string `json:"errorMsg"`
}
