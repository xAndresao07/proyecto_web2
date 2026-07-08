// Package config carga la configuracion de la aplicacion desde variables de
// entorno (con soporte para un archivo .env opcional) y expone valores por
// defecto razonables para desarrollo.
package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Puerto       string        // puerto HTTP, ej ":8080"
	DBDriver     string        // motor de base de datos: "sqlite" (default) o "postgres"
	DBDsn        string        // DSN de PostgreSQL (solo se usa si DBDriver="postgres")
	RutaDB       string        // archivo SQLite, ej "proyecto.db" (solo si DBDriver="sqlite")
	Backend      string        // backend de productos/categorias: "gorm" (default) o "sqlc"
	JWTSecreto   []byte        // clave para firmar/verificar JWT
	JWTDuracion  time.Duration // validez del token
	ReadTimeout  time.Duration // timeout de lectura del servidor HTTP
	WriteTimeout time.Duration // timeout de escritura del servidor HTTP
}

func Cargar() Config {
	_ = godotenv.Load()
	log.Printf("gfgfdgfddgf: motor=%s, backend=%s", os.Getenv("DB_DRIVER"), os.Getenv("STORAGE"))
	return Config{
		Puerto:       conTexto("PUERTO", ":8080"),
		DBDriver:     conTexto("DB_DRIVER", "sqlite"),
		DBDsn:        conTexto("DB_DSN", ""),
		RutaDB:       conTexto("RUTA_DB", "proyecto.db"),
		Backend:      conTexto("STORAGE", "sqlc"),
		JWTSecreto:   []byte(conTexto("JWT_SECRETO", "palabra-secreta-para-jwt")),
		JWTDuracion:  conDuracion("JWT_DURACION", 24*time.Hour),
		ReadTimeout:  conDuracion("HTTP_READ_TIMEOUT", 10*time.Second),
		WriteTimeout: conDuracion("HTTP_WRITE_TIMEOUT", 10*time.Second),
	}
}

func conTexto(clave, porDefecto string) string {
	if v := os.Getenv(clave); v != "" {
		return v
	}
	return porDefecto
}

func conDuracion(clave string, porDefecto time.Duration) time.Duration {
	v := os.Getenv(clave)
	if v == "" {
		return porDefecto
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return porDefecto
	}
	return d
}
