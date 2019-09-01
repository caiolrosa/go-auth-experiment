package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type mockSQLClient struct {
	DB   *sql.DB
	Mock sqlmock.Sqlmock
}

func (m *mockSQLClient) GetConnection() (*sqlx.DB, error) {
	return sqlx.NewDb(m.DB, "sqlite3"), nil
}

func TestHealthCheck(t *testing.T) {
	dbMock, mock := sqlMock()
	app := &App{dbClient: &mockSQLClient{
		DB:   dbMock,
		Mock: mock,
	}}

	server := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/healthcheck", nil)
	rec := httptest.NewRecorder()
	context := server.NewContext(req, rec)

	if err := app.HandleHealthCheck(context); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, rec.Code, http.StatusOK)
}

func TestHandleLogout(t *testing.T) {
	dbMock, mock := sqlMock()
	app := &App{dbClient: &mockSQLClient{
		DB:   dbMock,
		Mock: mock,
	}}

	server := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/api/login", nil)
	rec := httptest.NewRecorder()
	context := server.NewContext(req, rec)

	if err := app.HandleLogout(context); err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, rec.Code, http.StatusOK)
}

func sqlMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}

	return db, mock
}
