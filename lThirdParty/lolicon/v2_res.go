package lolicon

type V2Res struct {
	Error string `json:"error"`
	Data  []struct {
		Pid        int      `json:"pid"`
		P          int      `json:"p"`
		Uid        int      `json:"uid"`
		Title      string   `json:"title"`
		Author     string   `json:"author"`
		R18        bool     `json:"r18"`
		Width      int      `json:"width"`
		Height     int      `json:"height"`
		Tags       []string `json:"tags"`
		Ext        string   `json:"ext"`
		AiType     int      `json:"aiType"`
		UploadDate int64    `json:"uploadDate"`
		Urls       struct {
			Original string `json:"original"`
			Regular  string `json:"regular"`
			Small    string `json:"small"`
			Thumb    string `json:"thumb"`
			Mini     string `json:"mini"`
		} `json:"urls"`
	} `json:"data"`
}
