package razer

type appCreationRequest struct {
	Title           string                   `json:"title"`
	Description     string                   `json:"description"`
	Author          appCreationRequestAuthor `json:"author"`
	DeviceSupported []string                 `json:"device_supported"`
	Category        string                   `json:"category"`
}

type appCreationRequestAuthor struct {
	Name    string `json:"name"`
	Contact string `json:"contact"`
}

type appCreationResponse struct {
	SessionID int    `json:"sessionid"`
	URI       string `json:"uri"`
}

type effectCreationRequest struct {
	Effect string      `json:"effect"`
	Param  effectParam `json:"param"`
}

type effectParam struct {
	Color int `json:"color"`
}

type effectCreationResponse struct {
	ID     string `json:"id"`
	Result int    `json:"result"`
}

type effectApplyRequest struct {
	ID string `json:"id"`
}
