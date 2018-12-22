package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ko/cms-db/adminService/controllers"
)

func main() {

	// // Dummy endpoint to redirect elsewhere if someone hits /
	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.Redirect(w, r, "https://www.ovo.id/", http.StatusMovedPermanently)
	// })

	// http.ListenAndServe(":8080", mux)

	r := gin.Default()
	controllers.SetUpRoutes(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
