package database

type Blacklist struct {
	Model
	Sha1           string `gorm:"unique;not null;index"`
	AverageHash    string `gorm:"unique;not null;index"`
	DifferenceHash string `gorm:"unique;not null;index"`
}

func (Blacklist) TableName() string {
	return "Blacklist"
}

// Saves the data to the database
func (b *Blacklist) Save() {
	DB.Save(&b)
}
