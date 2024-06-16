package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Email representa la estructura de un correo electrónico
type Email struct {
	Subject string `json:"subject"`
	From    string `json:"from"`
	To      string `json:"to"`
	Body    string `json:"body"`
}

// Handlers struct tendrá métodos para manejar varias rutas
type Handlers struct {
	// Agrega cualquier dependencia que tus manejadores puedan necesitar, por ejemplo, una conexión a la base de datos
}

// NewHandlers retorna una nueva struct de Handlers
func NewHandlers() *Handlers {
	return &Handlers{}
}

// Home es el manejador para la ruta por defecto
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("¡Bienvenido a la página principal!"))
}

// Search es el manejador para la ruta /search
func (h *Handlers) Search(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("La funcionalidad de búsqueda va aquí."))
}

// GetEmails es el manejador para la ruta /emails
func (h *Handlers) GetEmails(w http.ResponseWriter, r *http.Request) {
	// Leer el archivo JSON
	fileData, err := ioutil.ReadFile("/Users/macbook/Documents/proyecto_v2_mail/Backend/cmd/api/handlers/allen-p.json")
	if err != nil {
		log.Fatalf("Error al leer el archivo: %v", err)
	}

	// Deserializar el JSON en una estructura Go
	var emailsData []map[string]string
	err = json.Unmarshal(fileData, &emailsData)
	if err != nil {
		log.Fatalf("Error al deserializar el JSON: %v", err)
	}

	// Procesar y extraer los campos relevantes
	var emails []Email
	for _, email := range emailsData {
		for _, content := range email {
			emails = append(emails, Email{
				Subject: extractField(content, "Subject: "),
				From:    extractField(content, "From: "),
				To:      extractField(content, "To: "),
				Body:    extractBody(content),
			})
		}
	}

	// Enviar la respuesta en formato JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emails)
}

func extractField(content, field string) string {
	fieldIndex := strings.Index(content, field)
	if fieldIndex == -1 {
		return ""
	}

	startIndex := fieldIndex + len(field)
	endIndex := strings.Index(content[startIndex:], "\n")
	if endIndex == -1 {
		return strings.TrimSpace(content[startIndex:])
	}

	return strings.TrimSpace(content[startIndex : startIndex+endIndex])
}

func extractBody(content string) string {
	bodyIndex := strings.Index(content, "\n\n")
	if bodyIndex == -1 {
		return ""
	}

	return strings.TrimSpace(content[bodyIndex+2:])
}
