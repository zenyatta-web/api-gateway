# Utiliza la imagen oficial de Go como base
FROM golang:1.20-alpine as builder

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos de tu proyecto al contenedor
COPY . .

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia el binario compilado desde el contenedor de compilación
COPY --from=builder /app/main .

# Establece el puerto en el que la aplicación escuchará
EXPOSE 8082

# Comando para ejecutar la aplicación
CMD ["/go-api-rest"]
