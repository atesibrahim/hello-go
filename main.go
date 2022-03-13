package main

import (
	"net/http"

	"github.com/atesibrahim/hello-go/controllers"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

	/*u := models.User{
		ID:        2,
		FirstName: "Ibrahim",
		LastName:  "Ates",
	}
	fmt.Println(u)*/

}
