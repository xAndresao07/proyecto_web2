package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"

	"proyecto/internal/handlers"
	"proyecto/internal/middleware"
	"proyecto/internal/service"
	"proyecto/internal/storage"
)

func TestRutasTecnicos_ProtegidasDevuelven401(t *testing.T) {
	// 1. Configuramos el entorno "fake" (dependencias en memoria)
	repoFake := storage.NuevaMemoria()
	svcTecnicos := service.NuevoTecnicoService(repoFake)

	// Como pasamos nil al repo de usuarios, el servicio de Auth fallara si intentara loguear,
	// pero para el test del 401 no llegara a tocar la BD porque no hay token.
	authSvc := service.NuevoAuthService(nil)

	// CORRECCIÓN: Usamos handlers.Deps{} como lo exige tu código
	servidor := handlers.NewServer(handlers.Deps{
		Tecnicos: svcTecnicos,
		Auth:     authSvc,
	})

	// 2. Configuramos un enrutador de prueba con el middleware de Auth
	r := chi.NewRouter()
	r.Use(middleware.Auth(authSvc))

	// CORRECCIÓN: El handler se llama GetAllTecnicos
	r.Get("/api/v1/tecnicos", servidor.GetAllTecnicos)

	// 3. Simulamos una peticion HTTP GET SIN enviar token Bearer
	req := httptest.NewRequest(http.MethodGet, "/api/v1/tecnicos", nil)
	w := httptest.NewRecorder()

	// 4. Disparamos la peticion
	r.ServeHTTP(w, req)

	// 5. Validamos que el servidor nos haya rechazado con 401
	if w.Code != http.StatusUnauthorized {
		t.Errorf("se esperaba status %d (Unauthorized), se obtuvo %d", http.StatusUnauthorized, w.Code)
	}
}
