package svr

type h1request struct {
	A int `json:"a" schema:"a"`
}

type h1response struct {
	Output    string `json:"output"`
	TimeStamp int64  `json:"time_stamp"`
}

type h2request struct {
	UUID string `json:"uuid" schema:"uuid"`
}

type h2response struct {
	StudyID string `json:"study_id"`
}
