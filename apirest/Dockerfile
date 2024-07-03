# Utiliza la imagen oficial de Go como base
FROM golang:1.18-alpine

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el archivo main.go
COPY main.go .

# Comando para ejecutar la aplicación
CMD ["go", "run", "main.go"]
