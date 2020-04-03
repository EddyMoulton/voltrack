package stocks

type Stock struct {
	Code        string `gorm:"PRIMARY_KEY"`
	Description string
}
