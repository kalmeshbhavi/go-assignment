package http

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"bytes"

	"encoding/json"

	"github.com/gorilla/mux"

	"github.com/kalmeshbhavi/go-assignment/engine"
	"github.com/kalmeshbhavi/go-assignment/providers/database"
)

var (
	router mux.Router
)

func TestMain(m *testing.M) {
	//configs := config.GetConfigs()
	connString := database.GetConnectionString()
	provider := database.NewProvider(connString)

	EnsureTableExists(provider)
	// todo: init database (ex: create table, clear previous data, etc.)

	e := engine.NewEngine(provider)

	router = *mux.NewRouter()

	httpAdapter := NewHTTPAdapter(e)
	adapters := getServerAdapters()

	router.Handle("/knight", httpAdapter.ServerApply(httpAdapter.getAll(), adapters...)).Methods("GET")
	router.Handle("/knight", httpAdapter.ServerApply(httpAdapter.create(), adapters...)).Methods("POST")
	router.Handle("/knight/{id}", httpAdapter.ServerApply(httpAdapter.get(), adapters...)).Methods("GET")

	httptest.NewServer(&router)

	code := m.Run()
	clearTable(provider)
	provider.Close()
	os.Exit(code)
}

func TestPostKnightBipolelm(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/knight", bytes.NewBuffer([]byte(`{"name":"Bipolelm","strength":10,"weapon_power":20}`)))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Fatal("Server error: Returned ", recorder.Code, " instead of ", http.StatusCreated)
	}
}

func TestPostKnightElrynd(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/knight", bytes.NewBuffer([]byte(`{"name":"Elrynd","strength":10,"weapon_power":50}`)))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Fatal("Server error: Returned ", recorder.Code, " instead of ", http.StatusCreated)
	}
}

func TestPostKnightBadData(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/knight", bytes.NewBuffer([]byte(`{"name":"FAILED"}`)))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Fatal("Server error: Returned ", recorder.Code, " instead of ", http.StatusBadRequest)
	}

	response := map[string]interface{}{}

	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if _, has := response["code"]; !has {
		t.Fatal("Response error: Expected code field")
	}

	if _, has := response["message"]; !has {
		t.Fatal("Response error: Expected message field")
	}
}

func TestPostKnightBadType(t *testing.T) {
	req, err := http.NewRequest(http.MethodPost, "/knight", bytes.NewBuffer([]byte(`name:"Bipolelm"`)))
	req.Header.Add("Content-Type", "text/plain")
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Fatal("Server error: Returned ", recorder.Code, " instead of ", http.StatusBadRequest)
	}
}

func TestGetKnights(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/knight", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatal("Server error: Returned ", recorder.Code, " instead of ", http.StatusOK)
	}

	var response []map[string]interface{}

	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if len(response) != 2 {
		t.Fatal("Response error: Expected 2 knights")
	}

	knight := response[0]

	if _, has := knight["id"]; !has {
		t.Fatal("Response error: Expected id field in knight object")
	}
	if _, has := knight["name"]; !has {
		t.Fatal("Response error: Expected name field in knight object")
	}
	if _, has := knight["strength"]; !has {
		t.Fatal("Response error: Expected strength field in knight object")
	}
	if _, has := knight["weapon_power"]; !has {
		t.Fatal("Response error: Expected weapon_power field in knight object")
	}

	if response[0]["id"].(string) == response[1]["id"].(string) {
		t.Fatal("Response error: Expected not same id for each knights")
	}
}

func TestGetKnightNotFound(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/knight/123456789", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Fatal("Server error: Returned ", recorder.Code, " instead of ", http.StatusNotFound)
	}

	response := map[string]interface{}{}

	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}

	if _, has := response["code"]; !has {
		t.Fatal("Response error: Expected code field")
	}

	if _, has := response["message"]; !has {
		t.Fatal("Response error: Expected message field")
	}

	if response["message"].(string) != "Knight #123456789 not found." {
		log.Println(response["message"])

		t.Fatal("Response error: Expected error message 'Knight #123456789 not found.'")
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS knights
(
    id SERIAL,
    name TEXT NOT NULL,
    strength int NOT NULL default 1,
    weapon_power int NOT NULL default 0,
    CONSTRAINT knights_pkey PRIMARY KEY (id)
)`

func EnsureTableExists(provider *database.Provider) {
	if _, err := provider.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable(provider *database.Provider) {
	provider.DB.Exec("DELETE FROM knights")
	provider.DB.Exec("ALTER SEQUENCE knights_id_seq RESTART WITH 1")
}
