package models

type LogsTable struct {
	ID   uint   `gorm:"primaryKey" json:"-"`
	Name string `gorm:"type:text;index" json:"name"`

	PrevTime     string `gorm:"type:timestamptz" json:"time_before"`
	StartTime    string `gorm:"type:timestamptz" json:"time_current"`
	NextSchedule string `gorm:"type:timestamptz" json:"time_next"`

	Action       string `gorm:"type:text" json:"action"`
	Status       string `gorm:"type:text;index" json:"status"`
	ErrorMessage string `gorm:"type:text" json:"error,omitempty"`
}

type LogsBack struct {
	Name string `json:"name"`
	CMDB string `json:"cmdb"`
}

type Action struct {
	Name       string   `json:"name"`
	Activities []string `json:"activities"`
}
