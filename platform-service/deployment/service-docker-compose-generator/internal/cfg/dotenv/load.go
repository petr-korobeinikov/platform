package dotenv

import "github.com/joho/godotenv"

func Load(files ...string) {
	for _, file := range files {
		if err := godotenv.Load(file); err != nil {
			continue
		}
	}
}
