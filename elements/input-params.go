package elements

// InputParams contiene los campos recibidos por JSON
type InputParams struct {
	Ancho           string `json:"ancho"`
	Largo           string `json:"largo"`
	Escala          string `json:"escala"`
	Formato         string `json:"formato"`
	ZPL             string `json:"zpl"`
	Filas           string `json:"filas"`
	Columnas        string `json:"columnas"`
	Mosaico         string `json:"mosaico"`
	MarginX         string `json:"marginX"`
	MarginY         string `json:"marginY"`
	Chunk           string `json:"chunk"`
	TipoPapel       string `json:"tipoPapel"`
	Orientacion     string `json:"orientacion"`
	Output          string `json:"output"`
	UrlOutput       string `json:"urlOutput"`
	Comprimir       string `json:"comprimir"`
	Resize          string `json:"resize"`
	LabelBackground string `json:"labelBackground"`
}

func (p InputParams) WithDefaults() InputParams {
	res := p

	if res.Ancho == "" {
		res.Ancho = "100"
	}
	if res.Largo == "" {
		res.Largo = "50"
	}
	if res.Escala == "" {
		res.Escala = "8dpmm"
	}
	if res.Formato == "" {
		res.Formato = "png"
	}
	if res.Mosaico == "" {
		res.Mosaico = "false"
	}
	if res.Filas == "" {
		res.Filas = "1"
	}
	if res.Columnas == "" {
		res.Columnas = "1"
	}
	if res.MarginX == "" {
		res.MarginX = "1"
	}
	if res.MarginY == "" {
		res.MarginY = "1"
	}
	if res.Chunk == "" {
		res.Chunk = "4000"
	}
	if res.TipoPapel == "" {
		res.TipoPapel = "A4"
	}
	if res.Orientacion == "" {
		res.Orientacion = "portrait"
	}
	if res.Output == "" {
		res.Output = "binary"
	}
	if res.UrlOutput == "" {
		res.UrlOutput = "./"
	}
	if res.Comprimir == "" {
		res.Comprimir = "false"
	}
	if res.Resize == "" {
		res.Resize = "false"
	}
	if res.LabelBackground == "" {
		res.LabelBackground = "true"
	}
	if res.ZPL == "" {
		res.ZPL = "^XA^FO50,50^ADN,36,20^FDHello, ZPL!^FS^XZ"
	}

	return res
}
