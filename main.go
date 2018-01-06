package main

import (
	db "future/database"

	// "github.com/braintree/manners"
)

func main() {
	defer db.SqlDB.Close()
	router := initRouter()
	// router.RunTLS(":8080","cert.pem","key.pem")
	router.Run(":8080")
	// manners.ListenAndServe(":8888", router)
	//manners.ListenAndServeTLS(":8888","server.crt", "server.key", router)
}
