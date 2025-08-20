# ZPL to PDF/PNG Converter

Este proyecto convierte etiquetas en formato **ZPL (Zebra Programming Language)** a archivos **PDF** o **PNG**, utilizando una versión personalizada de la biblioteca [zebrash](https://github.com/ingridhq/zebrash). Es ideal para generar etiquetas de impresoras Zebra en entornos donde se necesita renderizar sin hardware.

## ✨ Características

- Parseo de comandos ZPL.
- Renderizado a PNG.
- Generación de documentos PDF con soporte de:
  - Tamaño de etiqueta personalizado.
  - Modo mosaico (múltiples etiquetas por página).
- Entrada por `stdin`, salida por `stdout`.

## 📦 Dependencias

- [Fork personalizado de zebrash](https://github.com/lroentgenoil/zebrashMod)
- [gofpdf](https://github.com/jung-kurt/gofpdf)

## 📋 Parámetros
todos los datos deben ir entre comillas

| Campo             | Tipo    | Default      | Descripción                                                                  |
| ----------------- | ------- | ------------ | ---------------------------------------------------------------------------- |
| `zpl`             | string  | "^XA^FO50,50^ADN,36,20^FDHello, ZPL!^FS^XZ" | Código ZPL.                                 |
| `formato`         | string  | `"png"`      | `"pdf"` o `"png"`.                                                           |
| `ancho`           | float64 | `"100"`      | Ancho de etiqueta en milímetros.                                             |
| `largo`           | float64 | `"50"`       | Alto de etiqueta en milímetros.                                              |
| `escala`          | int     | `"8dpmm"`    | Puntos por milímetro (6, 8, 12, 24).                                         |
| `mosaico`         | bool    | `"false"`    | `true` para agrupar etiquetas por página.                                    |
| `orientacion`     | string  | `"portrait"` | `"P"` (portrait) o `"L"` (landscape).                                        |
| `tipoPapel`       | string  | `"A4"`       | Tamaño de papel, por ejemplo `"A4"`.                                         |
| `filas`           | int     | `"1"`        | Cantidad de filas (si `mosaico` es `true`).                                  |
| `columnas`        | int     | `"1"`        | Cantidad de columnas (si `mosaico` es `true`).                               |
| `marginX`         | float64 | `"1"`        | Margen horizontal (mm)(si `mosaico` es `true`).                              |
| `marginY`         | float64 | `"1"`        | Margen vertical (mm)(si `mosaico` es `true`).                                |
| `chunk`           | int     | `"4000"`     | division de etiquetas por archivo (ayuda a manejar el consumo de RAM).       |
| `output`          | string  | `"binary"`   | `"binary"` o `"file"`.                                                       |
| `urlOutput`       | string  | `"./" `      | salida del archivo `"./"` (si `output` es `file`).                           |
| `comprimir`       | bool    | `"false"`    | ayuda a reducir el peso de los archivos echos en mosaicos (requiere de mas procesamiento por lo que puede reducir la velocidad de creación). |
| `resize`          | bool    | `"false"`    | redimenciona el tamaño de la etiqueta para ajustarla al valor máximo de los ejes X / Y encontrado de un elemento dentro de la etiqueta. |
| `labelBackground` | bool    | `"true"`     | Agrega o quita el Background de la etiqueta                                  |

Instalación de dependencias:
```bash
go get github.com/lroentgenoil/zebrashMod
go get github.com/jung-kurt/gofpdf
```

Puedes compilarlo con: (En la raíz del proyecto)
- en windows
```bash
go build -o ./build/zpl-builder.exe main.go
```
- en linux
```bash
set GOOS=linux 
set GOARCH=amd64 
go build -o ./build/zpl-builder main.go 
```


Ejecutarlo pasándole un JSON por stdin: (ejecutar donde se encuentre zpl-builder.exe)
- en windows
```bash
echo { "zpl": "^XA^FO50,50^ADN,36,20^FDHello, ZPL!^FS^XZ", "formato": "png", "ancho": "100", "largo": "50", "escala": "8", "mosaico": "false", "orientacion": "P", "tipoPapel": "A4", "filas": "1", "columnas": "1", "marginX": "1", "marginY": "1", "chunk": "4000", "output": "binary", "urlOutput": "./", "comprimir": "false", "resize": "true", "labelBackground": "true" } | zpl-builder.exe
```
- en linux
```bash
echo { "zpl": "^XA^FO50,50^ADN,36,20^FDHello, ZPL!^FS^XZ", "formato": "png", "ancho": "100", "largo": "50", "escala": "8", "mosaico": "false", "orientacion": "P", "tipoPapel": "A4", "filas": "1", "columnas": "1", "marginX": "1", "marginY": "1", "chunk": "4000", "output": "binary", "urlOutput": "./", "comprimir": "false", "resize": "true", "labelBackground": "true" } | ./zpl-builder
```
- ejemplo en PHP
```php
$path = base_path('/zebrashMod');

$command = escapeshellcmd($path);
$params = json_encode([
    "zpl" => "^XA^FO50,50^ADN,36,20^FDHello, ZPL!^FS^XZ", 
    "formato" => "png", 
    "ancho" => "100", 
    "largo" => "50", 
    "escala" => "8", 
    "mosaico" => "false", 
    "orientacion" => "P", 
    "tipoPapel" => "A4", 
    "filas" => "1", 
    "columnas" => "1", 
    "marginX" => "1", 
    "marginY" => "1", 
    "chunk" => "4000", 
    "output" => "binary", 
    "urlOutput" => "./",
    "comprimir" => "false",
    "resize" => "true",
    "labelBackground" => "true"
]);

$descriptorspec = [
    0 => ["pipe", "r"],
    1 => ["pipe", "w"],
    2 => ["pipe", "w"] 
];

$process = proc_open($command, $descriptorspec, $pipes);

if (is_resource($process)) {
    fwrite($pipes[0], $params);
    fclose($pipes[0]);

    $data = stream_get_contents($pipes[1]);
    fclose($pipes[1]);

    $error = stream_get_contents($pipes[2]);
    fclose($pipes[2]);

    proc_close($process);
}
```