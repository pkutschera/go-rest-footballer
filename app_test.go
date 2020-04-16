package main

import (
	"github.com/fergusstrange/embedded-postgres"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var app Application

const createTableClub = `
	CREATE TABLE IF NOT EXISTS public.club (
		id integer NOT NULL DEFAULT nextval('club_id_seq'::regclass),
		name "char"[] NOT NULL,
		founded "char"[] NOT NULL,
		CONSTRAINT club_pkey PRIMARY KEY (id)
	)
`

const createTableFootballer = `
	CREATE TABLE IF NOT EXISTS public.footballer (
		id integer NOT NULL DEFAULT nextval('footballer_id_seq'::regclass),
		firstname "char"[] NOT NULL,
		lastname "char"[] NOT NULL,
		"position" "char"[] NOT NULL,
		clubid integer NOT NULL DEFAULT nextval('footballer_club_seq'::regclass),
		CONSTRAINT footballer_pkey PRIMARY KEY (id),
		CONSTRAINT "clubid constraint" FOREIGN KEY (clubid)
			REFERENCES public.club (id) MATCH SIMPLE
			ON UPDATE NO ACTION
			ON DELETE NO ACTION
			NOT VALID
	)
`

const createClubIdSequence = `
	CREATE SEQUENCE IF NOT EXISTS public.club_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1
`

const createFootballerClubSequence = `
	CREATE SEQUENCE IF NOT EXISTS public.footballer_club_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1
`

const createFootballerIdSequence = `
	CREATE SEQUENCE IF NOT EXISTS public.footballer_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1
`

func TestMain(m *testing.M) {
	postgres := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().
		Username("postgres").
		Password("admin").
		Database("postgres").
		Version(embeddedpostgres.V12).
		RuntimePath("/tmp").
		Port(9876).
		StartTimeout(45 * time.Second))

	if err := postgres.Start(); err == nil {
		log.Fatal(err)
	}
	defer postgres.Stop()

	app.Initialize("postgres", "admin", "postgres", 9876)
	defer app.DB.Close()

	createSchema()
	m.Run()
}

func createSchema() {
	if _, err := app.DB.Exec(createClubIdSequence); err != nil {
		log.Fatal(err)
	}
	if _, err := app.DB.Exec(createTableClub); err != nil {
		log.Fatal(err)
	}
	if _, err := app.DB.Exec(createFootballerIdSequence); err != nil {
		log.Fatal(err)
	}
	if _, err := app.DB.Exec(createFootballerClubSequence); err != nil {
		log.Fatal(err)
	}
	if _, err := app.DB.Exec(createTableFootballer); err != nil {
		log.Fatal(err)
	}
}

func TestEmptyTable(t *testing.T) {
	req, _ := http.NewRequest("GET", "/products", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
