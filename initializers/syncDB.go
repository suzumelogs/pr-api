package initializers

import "pr-api/models"

func SyncDB() {
	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to migrate the database: " + err.Error())
	}
}