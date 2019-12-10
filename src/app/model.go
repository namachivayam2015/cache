package app

type CacheData struct {
	Key string `json:"key"`
	Value string `json:"value"`
}

type CacheDataList []CacheData
