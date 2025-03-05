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
				// :D не успел сделать запросы, но Calc разделяет выражение и раскидывает в Taske (при чем максимально рационально)
				// но так как это не закончил, оставил старую Calc2, чтоб хоть как то считалось, короче строго не судите пжжж
			}
		} ()
	}
}
