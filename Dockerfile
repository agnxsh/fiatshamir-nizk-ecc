FROM go:1.18

WORKDIR /fiatshamir-nizk-ecc

COPY . .

RUN apt update

RUN go build