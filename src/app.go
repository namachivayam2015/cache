package main


import (
	"fmt"
	"github.com/gorilla/handlers"
	"test.mydomain.com/cache/cache"
	"test.mydomain.com/cache/util"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	defSize, er := strconv.Atoi(os.Getenv("DEFAULT_CACHE_SIZE"))
	if er != nil {
		fmt.Println("INFO : Default Size not specified")
		util.Size = 10
	} else {
		util.Size = defSize
	}

	timeOutThreshold, er := strconv.Atoi(os.Getenv("CACHE_READ_TIMEOUT_THRESHOLD"))
	if er != nil {
		fmt.Println("INFO : Default Read Timeout not specified")
	} else {
		util.ReadThresholdInSec = timeOutThreshold
	}

	go cacheGC()

	router := cache.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	var err = http.ListenAndServe(":8080", handlers.CORS(allowedOrigins, allowedMethods)(router))
	if err != nil {
		fmt.Print("err :: ", err)
	}

}

func cacheGC() {
	interval, err := strconv.Atoi(os.Getenv("CACHE_GC_INTERVAL"))
	if err != nil {
		fmt.Println("Fatal : Invalid GC Interval Specified!!")
		interval = 5
	}
	for {
		time.Sleep(time.Duration(interval)*time.Second)
		util.GarbageCollector()
	}
}

