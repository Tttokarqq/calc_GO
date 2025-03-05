package demon

import(
	"fmt"
	"time"
	"github.com/MaksaNeNegr/calc_go/vars"
	// "os"
)


func Demon_func(){
	// vars.Load()
	for{
		fmt.Println(os.Getenv("GITHUB_USERNAME"))
		time.Sleep(2 * time.Second)
	}
}