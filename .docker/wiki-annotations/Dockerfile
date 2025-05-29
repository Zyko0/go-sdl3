FROM golang:1.24

ENV DEBIAN_FRONTEND=noninteractive

RUN apt update -y

# Apparently wget is not installed by default
RUN apt-get install -y wget

# Install add-apt-repository
RUN apt-get install -y software-properties-common
# Clone libsdl-org/sdlwiki repository
RUN apt-get install -y git
RUN git clone https://github.com/libsdl-org/sdlwiki.git --depth 1

COPY *.go ./
RUN go mod init sdlwiki-annotations

# Run the annotations extraction script
RUN go run .