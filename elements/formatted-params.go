package elements

// FormattedParams contiene los campos convertidos y listos para usar
type FormattedParams struct {
	LabelWidth  float64
	LabelHeight float64
	Dpmm        int
	ZPL         []byte
	Formato     string
	Filas       int
	Columnas    int
	MarginX     float64
	MarginY     float64
	TipoPapel   string
	Mosaico     bool
	Orientacion string
	Output      string
	UrlOutput   string
}
