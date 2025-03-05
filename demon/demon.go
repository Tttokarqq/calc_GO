package demon

import(
	// "fmt"
	"time"
	"os"
	// "syns"
	"strconv"
)


func Demon_func(){
	s, _ := strconv.Atoi(os.Getenv("COMPUTING_POWER"))
	for i := 0; i < s; i++{
		go func() {
			for{
				time.Sleep(10 * time.Second)

			}
		} ()
	}
}