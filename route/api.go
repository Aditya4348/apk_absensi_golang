package route

import (
	"apk_absesnsi_harian/controller"
	"database/sql"
	"net/http"
)

func InitRoutes(db *sql.DB) {
	http.HandleFunc("/user", controller.GetUserHandler(db))


	http.HandleFunc("/hobi", controller.GetHobiHandler(db))
	http.HandleFunc("/hobiByID", controller.GetHobiByIDHandler(db))
}
