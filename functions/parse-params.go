package functions

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/lroentgenoil/zpl-builder/elements"
)

// ParseInputParams decodifica y convierte los datos de entrada
func ParseInputParams(inputBytes []byte) (elements.FormattedParams, error) {
	var raw elements.InputParams

	var formatData elements.FormattedParams

	err := json.Unmarshal(inputBytes, &raw)
	if err != nil {
		return elements.FormattedParams{}, err
	}

	raw = raw.WithDefaults()

	formatData.LabelWidth, err = strconv.ParseFloat(raw.Ancho, 64)
	if err != nil {
		return elements.FormattedParams{}, errors.New("error al convertir Ancho")
	}

	formatData.LabelHeight, err = strconv.ParseFloat(raw.Largo, 64)
	if err != nil {
		return elements.FormattedParams{}, errors.New("error al convertir Largo")
	}

	dpmmStr := strings.TrimSuffix(raw.Escala, "dpmm")
	formatData.Dpmm, err = strconv.Atoi(dpmmStr)
	if err != nil {
		return elements.FormattedParams{}, errors.New("error al convertir Escala")
	}

	formatData.Formato = strings.ToLower(raw.Formato)

	if formatData.Formato == "pdf" {
		formatData.Mosaico, err = strconv.ParseBool(raw.Mosaico)
		if err != nil {
			return elements.FormattedParams{}, errors.New("error al convertir Mosaico")
		}

		if formatData.Mosaico {
			formatData.Filas, err = strconv.Atoi(raw.Filas)
			if err != nil {
				return elements.FormattedParams{}, errors.New("error al convertir Filas")
			}

			formatData.Columnas, err = strconv.Atoi(raw.Columnas)
			if err != nil {
				return elements.FormattedParams{}, errors.New("error al convertir Columnas")
			}

			formatData.MarginX, err = strconv.ParseFloat(raw.MarginX, 64)
			if err != nil {
				return elements.FormattedParams{}, errors.New("error al convertir MarginX")
			}

			formatData.MarginY, err = strconv.ParseFloat(raw.MarginY, 64)
			if err != nil {
				return elements.FormattedParams{}, errors.New("error al convertir MarginY")
			}
		}

		formatData.Chunk, err = strconv.Atoi(raw.Chunk)
		if err != nil {
			return elements.FormattedParams{}, errors.New("error al convertir Chunk")
		}

		formatData.Comprimir, err = strconv.ParseBool(raw.Comprimir)
		if err != nil {
			return elements.FormattedParams{}, errors.New("error al convertir Comprimir")
		}
	}

	formatData.Output = strings.ToLower(raw.Output)
	if formatData.Output == "file" {
		formatData.UrlOutput = raw.UrlOutput
	}

	formatData.Resize, err = strconv.ParseBool(raw.Resize)
	if err != nil {
		return elements.FormattedParams{}, errors.New("error al convertir Resize")
	}

	formatData.LabelBackground, err = strconv.ParseBool(raw.LabelBackground)
	if err != nil {
		return elements.FormattedParams{}, errors.New("error al convertir LabelBackground")
	}

	formatData.ZPL = []byte(raw.ZPL)

	formatData = formatData.WithDefaults()

	return formatData, nil
}
