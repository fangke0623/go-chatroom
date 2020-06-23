package fileHandle

type FileEntity struct {
	FileName    string  `json:"fileName"`
	FileContent string  `json:"fileContent"`
	FileSize    float32 `json:"fileSize"`
}
type FileConf struct {
	Address string `json:"address"`
}
