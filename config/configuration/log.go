package configuration

import "time"

type Log struct {
	Path           string        `json:"path"`
	LinkPath       string        `json:"link_path"`
	Level          string        `json:"level"`
	FilenameFormat string        `json:"filename_format"`
	MaxAge         time.Duration `json:"max_age"`
	RotationTime   time.Duration `json:"rotation_time"`
}

func (l *Log) Process() {

}
