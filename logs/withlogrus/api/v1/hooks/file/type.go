package file

type FileConfiguration struct {
	Enable   bool                 `json:"enable"`
	Dir      string               `json:"dir"`
	BaseName string               `json:"base-name"`
	Route    bool                 `json:"route"`
	Rotate   *RotateConfiguration `json:"rotate"`
}

type RotateConfiguration struct {
	Enable     bool `json:"enable"`
	MaxSize    int  `json:"max_size"`
	MaxBackups int  `json:"max_backups"`
	MaxAge     int  `json:"max_age"`
	Compress   bool `json:"compress"`
}
