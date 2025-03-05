package demon

import(
	"fmt"
	"time"
	"github.com/MaksaNeNegr/calc_go/godotenv"
	"os"
)


func Demon_func(){
	err := godotenv.Load()
	for{
		fmt.Println(os.Getenv("GITHUB_USERNAME"))
		time.Sleep(2 * time.Second)
	}
}