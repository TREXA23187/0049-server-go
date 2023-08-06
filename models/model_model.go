package models

type ModelModel struct {
	MODEL
	Name            string `gorm:"size:36" json:"name"`
	Type            string `gorm:"size:36" json:"type"` // regression or classification
	IsDefault       bool   `gorm:"size:16" json:"is_default"`
	ModelFileName   string `gorm:"size:128" json:"model_file_name"`
	ModelFilePath   string `gorm:"size:128" json:"model_file_path"`
	IsGithub        bool   `gorm:"size:16" json:"is_github"`
	GithubLink      string `gorm:"size:128" json:"github_link"`
	HyperParameters string `gorm:"size:512" json:"hyper_parameters"`
	CreateUser      uint   `json:"create_user"`
}
