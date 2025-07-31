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
	MarginX     string `json:"marginX"`
	MarginY     string `json:"marginY"`
	TipoPapel   string `json:"tipoPapel"`
	Mosaico     string `json:"mosaico"`
	Orientacion string `json:"orientacion"`
	Output      string `json:"output"`
	UrlOutput   string `json:"urlOutput"`
}
