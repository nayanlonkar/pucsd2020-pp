package model

type Files struct {
	Id       int64  `json:"id,omitempty" key:"primary" autoincr:"1" column:"id"`
	FileName string `json:"file_name" column:"file_name"`
	FileType string `json:"file_type" column:"file_type"`
	ParentId int64  `json:"parent_id" column:"parent_id"`
}

func (file *Files) Table() string {
	return "files"
}

func (file *Files) String() string {
	return Stringify(file)
}
