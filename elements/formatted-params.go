package elements

// FormattedParams contiene los campos convertidos y listos para usar
type FormattedParams struct {
	LabelWidth      float64
	LabelHeight     float64
	Dpmm            int
	ZPL             []byte
	Formato         string
	Filas           int
	Columnas        int
	Mosaico         bool
	MarginX         float64
	MarginY         float64
	Chunk           int
	TipoPapel       string
	Orientacion     string
	Output          string
	UrlOutput       string
	Comprimir       bool
	Resize          bool
	LabelBackground bool
}

func (p FormattedParams) WithDefaults() FormattedParams {
	res := p

	if res.LabelWidth < 1 {
		res.LabelWidth = 100
	}
	if res.LabelHeight < 1 {
		res.LabelHeight = 50
	}
	validDpmm := map[int]bool{6: true, 8: true, 12: true, 24: true}
	if !validDpmm[res.Dpmm] {
		res.Dpmm = 8
	}
	validFormato := map[string]bool{"png": true, "pdf": true}
	if !validFormato[res.Formato] {
		res.Formato = "png"
	}
	if res.Mosaico {
		res.Mosaico = true
	}
	if res.Filas < 1 {
		res.Filas = 1
	}
	if res.Columnas < 1 {
		res.Columnas = 1
	}
	if res.MarginX < 1 {
		res.MarginX = 1
	}
	if res.MarginY < 1 {
		res.MarginY = 1
	}
	if res.Chunk < 1 {
		res.Chunk = 4000
	}
	validTipoPapel := map[string]bool{"A3": true, "A4": true, "A5": true, "Letter": true, "Legal": true, "Tabloid": true}
	if !validTipoPapel[res.TipoPapel] {
		res.TipoPapel = "A4"
	}
	validOrientacion := map[string]bool{"portrait": true, "landscape": true}
	if !validOrientacion[res.Orientacion] {
		res.Orientacion = "portrait"
	}
	validOutput := map[string]bool{"file": true, "binary": true}
	if !validOutput[res.Output] {
		res.Output = "binary"
	}
	if res.UrlOutput == "" {
		res.UrlOutput = "./"
	}
	if res.Comprimir {
		res.Comprimir = true
	}
	if res.Resize {
		res.Resize = true
	}
	if res.LabelBackground {
		res.LabelBackground = true
	}
	if len(res.ZPL) == 0 {
		res.ZPL = []byte("^XA^FO50,50^ADN,36,20^FDHello, ZPL!^FS^XZ")
	}

	return res
}
