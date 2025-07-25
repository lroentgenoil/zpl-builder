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

	initParams(&formatData)

	err := json.Unmarshal(inputBytes, &raw)
	if err != nil {
		return elements.FormattedParams{}, err
	}

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
	}

	formatData.ZPL = []byte(raw.ZPL)

	return formatData, nil
}

func initParams(params *elements.FormattedParams) {
	params.LabelWidth = 0
	params.LabelHeight = 0
	params.Dpmm = 0
	params.Formato = "jpg"
	params.Filas = 0
	params.Columnas = 0
	params.MarginX = 0
	params.MarginY = 0
	params.Mosaico = false
}
