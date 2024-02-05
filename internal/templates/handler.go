package templates

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"mymodule/internal/config"
	"mymodule/internal/handlers"
	"net/http"
)

var _ handlers.Ihandler = &Handler{}

type Handler struct{}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/", h.home)
	mux.HandleFunc("/about", h.about)
	mux.HandleFunc("/contact", h.contact)
}

var cfg = config.GetConfig()

func connectToDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable client_encoding=UTF8",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)

	db, err := sql.Open(cfg.DbDriver, psqlInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}

	return db, err
}

type TestData struct {
	Id int
}

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	db, err := connectToDB()
	if err != nil {
		log.Println(err)
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT id FROM post")
	if err != nil {
		log.Println(err)
		return
	}

	defer rows.Close()

	var testData []TestData

	for rows.Next() {
		var t TestData

		if err = rows.Scan(&t.Id); err != nil {
			log.Println(err)
			return
		}
		testData = append(testData, t)
	}

	fmt.Fprintf(w, "%v", testData)
}
func (h *Handler) about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About")
}
func (h *Handler) contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact")
}
