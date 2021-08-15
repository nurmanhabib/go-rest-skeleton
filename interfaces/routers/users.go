package routers

import (
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"github.com/nurmanhabib/go-rest-skeleton/pkg/response"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Email            string `json:"email" bson:"email"`
}

func handleUsers(e *gin.Engine) {
	// Setup the mgm default config
	err := mgm.SetDefaultConfig(nil, "skeleton", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	e.GET("users", func(c *gin.Context) {
		//panic(errors.New("user not found"))
		//time.Sleep(5 * time.Second)

		coll := mgm.Coll(&User{})
		find, err := coll.Find(nil, bson.D{})
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		var users []User

		err = find.All(nil, &users)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		response.JSON(c.Writer, response.SuccessResponse{Data: users})
	})

	e.POST("/users", func(c *gin.Context) {
		u := &User{
			Name:  "Habib",
			Email: "habib@mail.com",
		}

		err := mgm.Coll(u).Create(u)
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		//jsonData, _ := ioutil.ReadAll(c.Request.Body)
		err = c.Request.ParseForm()
		if err != nil {
			panic(err)
		}

		log.Printf("%s", c.Request.Form)

		//return

		//rollbar.Warning(f)
		//rollbar.Wait()

		var timer *time.Timer
		timer.Reset(10 * time.Second)

		response.JSON(c.Writer, response.SuccessResponse{
			Code: http.StatusCreated,
			Data: u,
		})
	})
}
