package main

import (
	"github.com/moosashah/go-crud/initializers"
	"github.com/moosashah/go-crud/tournament"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	initializers.DB.AutoMigrate(&tournament.Model{})
	//initializers.DB.Migrator().DropTable(&tournament.Model{})
}
