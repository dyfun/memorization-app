package seeders

import "fmt"

func Seed() {
	fmt.Println("Seeder started")

	// Run user seeder
	UserSeeder()
}
