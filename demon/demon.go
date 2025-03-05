package demon

import(
	"fmt"
	"time"
	"github.com/MaksaNeNegr/calc_go/godotenv"
	"os"
)


func Demon_func(){
	for{
		fmt.Println(os.LookupEnv("GITHUB_USERNAME"))
		time.Sleep(2 * time.Second)
	}
}