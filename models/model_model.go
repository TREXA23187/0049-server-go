package models

type ModelModel struct {
	MODEL
	Name       string `gorm:"size:36" json:"name"`
	IsDefault  bool   `gorm:"size:16" json:"is_default"`
	FilePath   string `gorm:"size:128" json:"file_path"`
	IsGithub   bool   `gorm:"size:16" json:"is_github"`
	GithubLink string `gorm:"size:128" json:"github_link"`
}
