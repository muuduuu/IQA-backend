package config


import(
	"log"
	"os"

	"github.com/joho/godotenv"
)


func LoadEnv(
	if err:=godotenv.Load(); err!=nil{
			log.Printf("Env file wasnt found")
	}


)

func GetEnv(key string) string {
	val:= os.GetEnv(key)
	if val = "" {
		log.Fatalf("Envar not set")
	}

	return val
}