package stocks

// Stock is an object for storing a unique stock code along with a description
type Stock struct {
	Code        string `gorm:"PRIMARY_KEY"`
	Description string
}
