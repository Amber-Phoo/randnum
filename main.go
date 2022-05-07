package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

type CLIOptions struct {
	Port      string
	StaticDir string
	Mode      string
}

func (opt *CLIOptions) options() {
	p := os.Getenv("PORT")
	if "" == p {
		p = "8080"
	}
	port := flag.String("port", fmt.Sprintf(":%s", p), "Port to listen on")
	staticDir := flag.String("staticDir", "./static", "Directory of static resources")
	mode := flag.String("mode", "debug", "Gin mode: debug, test, release")

	flag.Parse()

	opt.Port = *port
	opt.StaticDir = *staticDir
	opt.Mode = *mode
}

func index(ctx *gin.Context) {

	nums := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "10",
		"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
		"21", "22", "23", "24", "25", "26", "27", "28", "29", "30",
	}

	count, err := strconv.Atoi(ctx.Query("count"))
	if nil != err {
		count = 1
	}

	if count > len(nums) {
		count = len(nums)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	ctx.HTML(http.StatusOK, "result", gin.H{
		"count": count,
		"nums":  nums[0:count],
	})
}

func notFound(ctx *gin.Context) {
	ctx.HTML(http.StatusNotFound, "404", nil)
}

func main() {

	fmt.Println("Random number generator")

	options := CLIOptions{}
	options.options()

	gin.SetMode(options.Mode)

	router := gin.Default()

	router.HTMLRender = ginview.Default()

	router.GET("/", index)
	router.GET("/index.html", index)

	router.Use(static.Serve("/", static.LocalFile(options.StaticDir, false)))

	router.Use(notFound)

	log.Printf("Starting application on port %s in %s mode", options.Port, options.Mode)

	if err := router.Run(options.Port); nil != err {
		log.Fatalf("Error: %v\n", err)
	}

}
