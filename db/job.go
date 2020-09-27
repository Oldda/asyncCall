package db

type Jobs struct{
	ID uint `gorm:"column:id;primaryKey"`
	Delay string `gorm:"column:delay" json:"delay"`
	RequestUrl string `gorm:"column:request_url" json:"request_url"`
	RequestParams string `gorm:"column:request_params" json:"request_params"`
	RequestTime uint `gorm:"column:request_time" json:"request_time"`
	CreateTime uint `gorm:"column:create_time" json:"create_time"`
	Status uint `gorm:"column:status" json:"status"`
}

func (this *Jobs) TableName() string {
  return "crm_jobs"
}