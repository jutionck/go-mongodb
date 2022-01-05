package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mongodb/db"
	"go-mongodb/utils"
	"log"
	"os"
)

type Routes struct {
}

func (app Routes) StartGin() {
	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")

	mdb, err := db.InitResource()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := utils.InitContext()
	defer cancel()
	defer func() {
		if err = mdb.Db.Client().Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	r := gin.Default()
	publicRoute := r.Group("/api")
	NewStudentApi(publicRoute, mdb)

	apiBaseUrl := fmt.Sprintf("%s:%s", host, port)
	err = r.Run(apiBaseUrl)
	if err != nil {
		log.Fatal(err)
	}
}
