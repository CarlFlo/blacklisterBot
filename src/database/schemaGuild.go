package database

type Guild struct {
	Model
	GuildID string
}

func (Guild) TableName() string {
	return "Blacklist"
}

// Saves the data to the database
func (g *Guild) Save() error {
	tx := DB.Save(&g)
	return tx.Error
}
