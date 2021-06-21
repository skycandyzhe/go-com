package file

import (
	"os"
	"path"
	"powershellDeal/common/logger"

	"github.com/gabriel-vasile/mimetype"
)

type FileObj struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Size      string `json:"size"`
	Crc32     string `json:"crc32"`
	Md5       string `json:"md5"`
	Sha256    string `json:"sha256"`
	Sha512    string `json:"sha512"`
	Ssdeep    string `json:"ssdeep"`
	Type      string `json:"type"`
	Yara      string `json:"yara"`
	Urls      string `json:"urls"`
	Ext       string `json:"ext"`
	MimeExt   string `json:"mimeext"`
	temporary bool
}

func NewFileObj(filepath string, temporary bool) *FileObj {

	var obj *FileObj = &FileObj{
		Name: path.Base(filepath),
		Path: filepath,
		// Size:      "",
		// Crc32:     "",
		// Md5:       "",
		// Sha256:    "",
		// Sha512:    "",
		// Ssdeep:    "",
		// Type:      "",
		// Yara:      "",
		// Urls:      "",
		// Ext:       "",
		temporary: temporary,
	}
	obj.getext()
	obj.getType()
	return obj
}
func (u *FileObj) getext() {
	u.Ext = path.Ext(u.Name)
}
func (u *FileObj) getType() {
	mime, err := mimetype.DetectFile(u.Path)
	if err != nil {
		logger.Logger.Warn("Get file Mimetype failure path:", u.Path, err)

	}
	u.Type = mime.String()
	u.MimeExt = mime.Extension()
	// if u.MimeExt != "" && u.Ext != "" {
	// 	if u.MimeExt != u.Ext {
	// 		u. = append(u., "ExtNotSame:5")
	// 	}
	// }
}
func (u *FileObj) Close() {
	if u.temporary {
		err := os.Remove(u.Path)
		if err != nil {
			logger.Logger.Info("file remove error:", err)
		}
	}
}
