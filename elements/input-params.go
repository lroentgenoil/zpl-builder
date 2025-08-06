package elements

// InputParams contiene los campos recibidos por JSON
type InputParams struct {
	Ancho       string `json:"ancho"`
	Largo       string `json:"largo"`
	Escala      string `json:"escala"`
	Formato     string `json:"formato"`
	ZPL         string `json:"zpl"`
	Filas       string `json:"filas"`
	Columnas    string `json:"columnas"`
	Mosaico     string `json:"mosaico"`
	MarginX     string `json:"marginX"`
	MarginY     string `json:"marginY"`
	Chunk       string `json:"chunk"`
	TipoPapel   string `json:"tipoPapel"`
	Orientacion string `json:"orientacion"`
	Output      string `json:"output"`
	UrlOutput   string `json:"urlOutput"`
	Comprimir   string `json:"comprimir"`
}
