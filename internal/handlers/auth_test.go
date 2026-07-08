package handlers_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegistrar(t *testing.T) {
	h, _, _ := construirEntorno()

	t.Run("valido -> 201", func(t *testing.T) {
		// Tu API exige enviar el "rol"
		body := `{"email":"nuevo@uleam.edu.ec","password":"secreta123","rol":"solicitante"}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/auth/registrar", body, ""))
		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("email duplicado -> 409", func(t *testing.T) {
		body := `{"email":"dup@uleam.edu.ec","password":"secreta123","rol":"tecnico"}`
		ejecutar(h, jsonReq(http.MethodPost, "/api/v1/auth/registrar", body, ""))        // Primer registro
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/auth/registrar", body, "")) // Intento repetido
		assert.Equal(t, http.StatusConflict, rec.Code)
	})

	t.Run("campos vacios -> 400", func(t *testing.T) {
		// Petición sin rol ni contraseña (dispara tu ErrNombreVacio)
		body := `{"email":"incompleto@uleam.edu.ec"}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/auth/registrar", body, ""))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("JSON malformado -> 400", func(t *testing.T) {
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/auth/registrar", `{roto`, ""))
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})
}

func TestLogin(t *testing.T) {
	h, _, _ := construirEntorno()

	// Creamos un usuario inicial válido en el fake para poder probar el login
	credRegistro := `{"email":"ana@uleam.edu.ec","password":"secreta123","rol":"admin"}`
	ejecutar(h, jsonReq(http.MethodPost, "/api/v1/auth/registrar", credRegistro, ""))

	t.Run("correcto -> 200 + token", func(t *testing.T) {
		// Para loguearse solo necesitas email y password
		credLogin := `{"email":"ana@uleam.edu.ec","password":"secreta123"}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/auth/login", credLogin, ""))
		require.Equal(t, http.StatusOK, rec.Code)

		var resp struct {
			Token string `json:"token"`
		}
		require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
		assert.NotEmpty(t, resp.Token)
	})

	t.Run("contrasena incorrecta -> 401", func(t *testing.T) {
		malo := `{"email":"ana@uleam.edu.ec","password":"incorrecta"}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/auth/login", malo, ""))
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	})

	t.Run("usuario no existe -> 409", func(t *testing.T) {
		// Tu manejador de errores devuelve StatusConflict (409) para ErrEmailEnUso
		noExiste := `{"email":"fantasma@uleam.edu.ec","password":"123"}`
		rec := ejecutar(h, jsonReq(http.MethodPost, "/api/v1/auth/login", noExiste, ""))
		assert.Equal(t, http.StatusConflict, rec.Code)
	})
}
