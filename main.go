package main

import (
	"fmt"
	"io"
	"os"

	zebrash "github.com/lroentgenoil/zebrashMod"
	"github.com/lroentgenoil/zebrashMod/drawers"
	"github.com/lroentgenoil/zpl-builder/functions"
	"github.com/lroentgenoil/zpl-builder/handle"
)

func main() {
	inputBytes, err := io.ReadAll(os.Stdin)
	exitWithError(err)

	params, err := functions.ParseInputParams(inputBytes)
	exitWithError(err)

	// Procesar ZPL
	parser := zebrash.NewParser()
	res, err := parser.Parse(params.ZPL)
	if err != nil || len(res) == 0 {
		exitWithError(fmt.Errorf("error parsing ZPL: %v", err))
	}

	// Configurar opciones del drawer
	opts := drawers.DrawerOptions{
		LabelWidthMm:    params.LabelWidth,
		LabelHeightMm:   params.LabelHeight,
		Dpmm:            params.Dpmm,
		LabelBackground: params.LabelBackground,
		Resize:          params.Resize,
	}

	switch params.Formato {
	case "pdf":
		err = handle.PDFOutput(res, params, opts)
	default:
		err = handle.ImageOutput(res, params, opts)
	}

	exitWithError(err)
}

func exitWithError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
