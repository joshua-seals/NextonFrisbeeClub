FROM golang:tip-alpine3.21 

COPY . /nextonfrisbeeclub/

WORKDIR /nextonfrisbeeclub/
RUN go build -o nfc ./cmd/web/  

EXPOSE 3000 
CMD [ "./nfc" ]



