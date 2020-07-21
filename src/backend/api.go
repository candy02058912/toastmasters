package svr

type h1request struct {
	A int `json:"a" schema:"a"`
	B int `json:"b" schema:"b"`
}

type h1response struct {
	Answer    int   `json:"answer"`
	TimeStamp int64 `json:"time_stamp"`
}

type h2request struct {
	UUID string `json:"uuid" schema:"uuid"`
}

type h2response struct {
	StudyID string `json:"study_id"`
}
