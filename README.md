# ZPL to PDF/PNG Converter

Este proyecto convierte etiquetas en formato **ZPL (Zebra Programming Language)** a archivos **PDF** o **PNG**, utilizando una versión personalizada de la biblioteca `zebrash`. Es ideal para generar etiquetas de impresoras Zebra en entornos donde se necesita renderizar sin hardware.

## ✨ Características

- Parseo de comandos ZPL.
- Renderizado a PNG.
- Generación de documentos PDF con soporte de:
  - Tamaño de etiqueta personalizado.
  - Modo mosaico (múltiples etiquetas por página).
- Entrada por `stdin`, salida por `stdout`.

## 📦 Dependencias

- [Fork personalizado de zebrash](https://github.com/lroentgenoil/zebrash)
- [gofpdf](https://github.com/jung-kurt/gofpdf)

## 📋 Parámetros
todos los datos deben ir entre comillas

| Campo         | Tipo    | Descripción                                    |
| ------------- | ------- | ---------------------------------------------- |
| `zpl`         | string  | Código ZPL.                                    |
| `formato`     | string  | `"pdf"` o `"png"`.                             |
| `ancho`       | float64 | Ancho de etiqueta en milímetros.               |
| `largo`       | float64 | Alto de etiqueta en milímetros.                |
| `escala`      | int     | Puntos por milímetro (6, 8, 12, 24).           |
| `mosaico`     | bool    | `true` para agrupar etiquetas por página.      |
| `orientacion` | string  | `"P"` (portrait) o `"L"` (landscape).          |
| `tipoPapel`   | string  | Tamaño de papel, por ejemplo `"A4"`.           |
| `filas`       | int     | Cantidad de filas (si `mosaico` es `true`).    |
| `columnas`    | int     | Cantidad de columnas (si `mosaico` es `true`). |
| `marginX`     | float64 | Margen horizontal (mm)(si `mosaico` es `true`).|
| `marginY`     | float64 | Margen vertical (mm)(si `mosaico` es `true`).  |

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
GOOS=linux 
GOARCH=amd64 
go build -o ./build/zpl-builder main.go 
```


Ejecutarlo pasándole un JSON por stdin: (ejecutar donde se encuentre zpl-builder.exe)
- en windows
```bash
echo '{ "zpl": "^XA^FO50,50^ADN,36,20^FDHello, ZPL!^FS^XZ", "formato": "pdf", "ancho": "100", "largo": "50", "escala": "8", "mosaico": "true", "orientacion": "P", "tipoPapel": "A4", "filas": "4", "columnas": "2", "marginX": "5", "marginY": "5"}' | zpl-builder.exe
```
- en linux
```bash
echo '{ "zpl": "^XA^FO50,50^ADN,36,20^FDHello, ZPL!^FS^XZ", "formato": "pdf", "ancho": "100", "largo": "50", "escala": "8", "mosaico": "true", "orientacion": "P", "tipoPapel": "A4", "filas": "4", "columnas": "2", "marginX": "5", "marginY": "5"}' | ./zpl-builder
```
- ejemplo en PHP
```php
$path = base_path('/zebrashMod');

$command = escapeshellcmd($path);
$params = json_encode([
    "zpl" => "^XA^FO50,50^ADN,36,20^FDHello, ZPL!^FS^XZ", 
    "formato" => "pdf", 
    "ancho" => "100", 
    "largo" => "50", 
    "escala" => "8", 
    "mosaico" => "true", 
    "orientacion" => "P", 
    "tipoPapel" => "A4", 
    "filas" => "4", 
    "columnas" => "2", 
    "marginX" => "5", 
    "marginY" => "5"
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