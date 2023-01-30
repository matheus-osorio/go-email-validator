package database

func autoMigrate() {
	db.AutoMigrate(&Email{})
}
