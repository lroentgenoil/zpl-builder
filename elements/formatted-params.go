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
	Mosaico     bool
	MarginX     float64
	MarginY     float64
	Chunk       int
	TipoPapel   string
	Orientacion string
	Output      string
	UrlOutput   string
	Comprimir   bool
}
