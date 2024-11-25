package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	// req := ... // здесь нужно создать запрос к сервису

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// здесь нужно добавить необходимые проверки
}

// func TestMainHandlerWhenOk(t *testing.T) {
//     req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

//     responseRecorder := httptest.NewRecorder()
//     handler := http.HandlerFunc(mainHandle)
//     handler.ServeHTTP(responseRecorder, req)

//     if status := responseRecorder.Code; status != http.StatusOK {
//         t.Errorf("expected status code: %d, got %d", http.StatusOK, status)
//     }
// }

// func TestMainHandlerWhenMissingCount(t *testing.T) {
//     req := httptest.NewRequest("GET", "/cafe?city=moscow", nil)

//     responseRecorder := httptest.NewRecorder()
//     handler := http.HandlerFunc(mainHandle)
//     handler.ServeHTTP(responseRecorder, req)

//     if status := responseRecorder.Code; status != http.StatusBadRequest {
//         t.Errorf("expected status code: %d, got %d", http.StatusBadRequest, status)
//     }

//     expected := `count missing`
//     if responseRecorder.Body.String() != expected {
//         t.Errorf("expected body: %s, got %s", expected, responseRecorder.Body.String())
//     }
// }
