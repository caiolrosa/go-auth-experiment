package main

import (
	"errors"
	"guardian-api/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockAuthenticationService struct {
	mock.Mock
	ShouldError bool
	StatusCode  int
	User        models.User
}

type mockJWTService struct {
	mock.Mock
	ShouldError bool
	UID         string
}

type mockUserRepository struct {
	mock.Mock
	ShouldError bool
	User        models.User
}

func TestHealthCheck(t *testing.T) {
	app := &App{}

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
	app := &App{}

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

func TestHandleLogin(t *testing.T) {
	app := App{}
	server := echo.New()

	testCases := []struct {
		inputJSON          string
		inputUser          models.User
		expectedStatusCode int
		shouldError        bool
	}{
		{
			inputJSON:          `{"email": "test@test.com", "password": "123"}`,
			inputUser:          models.User{Email: "test@test.com", Password: "123"},
			expectedStatusCode: http.StatusOK,
			shouldError:        false,
		},
		{
			inputJSON:          `{"email": "test@test.com", "password": "123"}`,
			inputUser:          models.User{Email: "test@test.com", Password: "123456"},
			expectedStatusCode: http.StatusUnauthorized,
			shouldError:        true,
		},
	}

	for _, testCase := range testCases {
		app.AuthenticationService = &mockAuthenticationService{
			ShouldError: testCase.shouldError,
			User:        testCase.inputUser,
			StatusCode:  testCase.expectedStatusCode,
		}

		app.JWTService = &mockJWTService{
			ShouldError: testCase.shouldError,
			UID:         "123",
		}

		req := httptest.NewRequest(http.MethodPost, "/api/login", strings.NewReader(testCase.inputJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		context := server.NewContext(req, rec)

		if err := app.HandleLogin(context); err != nil {
			t.Error(err)
			return
		}

		res := rec.Result()
		assert.Equal(t, testCase.expectedStatusCode, res.StatusCode)
	}
}

func TestHandleRegister(t *testing.T) {
	app := App{}
	server := echo.New()

	testCases := []struct {
		inputJSON          string
		inputUser          models.User
		expectedStatusCode int
		shouldError        bool
	}{
		{
			inputJSON:          `{"email": "test@test.com", "password": "123"}`,
			inputUser:          models.User{Email: "test@test.com", Password: "123"},
			expectedStatusCode: http.StatusUnprocessableEntity,
			shouldError:        true,
		},
		{
			inputJSON:          `{"name": "test", "email": "test", "password": "123"}`,
			inputUser:          models.User{Email: "test@test.com", Password: "123456"},
			expectedStatusCode: http.StatusUnprocessableEntity,
			shouldError:        true,
		},
		{
			inputJSON:          `{"name": "test", "email": "test@test.com", "password": "123"}`,
			inputUser:          models.User{Email: "test@test.com", Password: "123456"},
			expectedStatusCode: http.StatusUnprocessableEntity,
			shouldError:        true,
		},
		{
			inputJSON:          `{"name": "test", "email": "test@test.com", "password": "12345678"}`,
			inputUser:          models.User{Email: "test@test.com", Password: "123456"},
			expectedStatusCode: http.StatusInternalServerError,
			shouldError:        true,
		},
		{
			inputJSON:          `{"name": "test", "email": "test@test.com", "password": "12345678"}`,
			inputUser:          models.User{Email: "test@test.com", Password: "123456"},
			expectedStatusCode: http.StatusOK,
			shouldError:        false,
		},
	}

	for _, testCase := range testCases {
		app.UserRepository = &mockUserRepository{
			ShouldError: testCase.shouldError,
			User:        testCase.inputUser,
		}

		app.JWTService = &mockJWTService{
			ShouldError: testCase.shouldError,
			UID:         "",
		}

		req := httptest.NewRequest(http.MethodPost, "/api/register", strings.NewReader(testCase.inputJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()
		context := server.NewContext(req, rec)

		if err := app.HandleRegister(context); err != nil {
			t.Error(err)
			return
		}

		res := rec.Result()
		assert.Equal(t, testCase.expectedStatusCode, res.StatusCode)
	}
}

func TestHandleAuth(t *testing.T) {
	app := App{}
	server := echo.New()

	testCases := []struct {
		UID                string
		inputCookie        http.Cookie
		inputUser          models.User
		expectedStatusCode int
		shouldError        bool
	}{
		{
			UID:                "123",
			inputCookie:        http.Cookie{Name: "should_error", Value: "123"},
			inputUser:          models.User{Email: "test@test.com", Password: "123456"},
			expectedStatusCode: http.StatusTemporaryRedirect,
			shouldError:        true,
		},
		{
			UID:                "",
			inputCookie:        http.Cookie{Name: cookieName, Value: "123"},
			inputUser:          models.User{Email: "test@test.com", Password: "123"},
			expectedStatusCode: http.StatusTemporaryRedirect,
			shouldError:        true,
		},
		{
			UID:                "asd",
			inputCookie:        http.Cookie{Name: cookieName, Value: "123"},
			inputUser:          models.User{Email: "test@test.com", Password: "123456"},
			expectedStatusCode: http.StatusTemporaryRedirect,
			shouldError:        true,
		},
		{
			UID:                "123",
			inputCookie:        http.Cookie{Name: cookieName, Value: "123"},
			inputUser:          models.User{ID: 123, Email: "test@test.com", Password: "123456"},
			expectedStatusCode: http.StatusTemporaryRedirect,
			shouldError:        true,
		},
		{
			UID:                "123",
			inputCookie:        http.Cookie{Name: cookieName, Value: "123"},
			inputUser:          models.User{ID: 123, Email: "test@test.com", Password: "123456"},
			expectedStatusCode: http.StatusOK,
			shouldError:        false,
		},
	}

	for _, testCase := range testCases {
		app.JWTService = &mockJWTService{
			ShouldError: testCase.shouldError,
			UID:         testCase.UID,
		}
		app.AuthenticationService = &mockAuthenticationService{
			ShouldError: testCase.shouldError,
			StatusCode:  testCase.expectedStatusCode,
			User:        testCase.inputUser,
		}

		req := httptest.NewRequest(http.MethodGet, "/api/auth", nil)
		rec := httptest.NewRecorder()
		context := server.NewContext(req, rec)
		req.AddCookie(&testCase.inputCookie)

		if err := app.HandleAuth(context); err != nil {
			t.Error(err)
			return
		}

		res := rec.Result()
		assert.Equal(t, testCase.expectedStatusCode, res.StatusCode)
	}
}

func (m *mockAuthenticationService) AuthenticateUser(email string, password string) (models.User, int, error) {
	if m.ShouldError {
		return models.User{}, m.StatusCode, errors.New("error")
	}

	return m.User, http.StatusOK, nil
}

func (m *mockAuthenticationService) AuthenticateUserByID(id int64) (models.User, error) {
	if m.ShouldError {
		return models.User{}, errors.New("error")
	}

	return m.User, nil
}

func (m *mockUserRepository) FindByID(ID int64) (models.User, error) {
	if m.ShouldError {
		return models.User{}, errors.New("error")
	}

	return m.User, nil
}

func (m *mockUserRepository) FindByEmail(email string) (models.User, error) {
	if m.ShouldError {
		return models.User{}, errors.New("error")
	}

	return m.User, nil
}

func (m *mockUserRepository) Save(user models.User) (models.User, error) {
	if m.ShouldError {
		return models.User{}, errors.New("error")
	}

	return m.User, nil
}

func (m *mockJWTService) ParseToken(token string) (string, error) {
	if m.ShouldError {
		return "", errors.New("error")
	}

	return m.UID, nil
}

func (m *mockJWTService) IssueToken(id string) (string, error) {
	// NOT IMPLEMENTED
	return "", nil
}
