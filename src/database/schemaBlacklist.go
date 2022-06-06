package database

type Blacklist struct {
	Model
	Sha1           string `gorm:"unique;not null;index"`
	AverageHash    uint64 `gorm:"not null;index"`
	DifferenceHash uint64 `gorm:"not null;index"`
	PerceptionHash uint64 `gorm:"not null;index"`
	URL            string `gorm:"unique;not null"`
}

func (Blacklist) TableName() string {
	return "Blacklist"
}

// Saves the data to the database
func (b *Blacklist) Save() {
	DB.Save(&b)
}

func (b *Blacklist) DeleteEntry() {
	DB.Delete(&b)
}
