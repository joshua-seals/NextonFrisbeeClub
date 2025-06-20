FROM golang as build 

COPY . /opt/nextonfrisbeeclub/

WORKDIR /opt/nextonfrisbeeclub/
RUN go build -o /cmd/web/ nextonfrisbeeclub 



