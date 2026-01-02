FROM golang:1.25

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y

# Install add-apt-repository
# Update package list and install necessary dependencies
RUN apt-get update && apt-get install -y \
    gnupg2 \
    lsb-release \
    && echo "deb http://deb.debian.org/debian/ bullseye main" > /etc/apt/sources.list.d/bullseye.list \
    && apt-get update && apt-get install -y \
    wget \
    software-properties-common \
    apt-transport-https \
    ca-certificates \
    curl \
    && rm -rf /var/lib/apt/lists/*
# RUN apt-get update -y && apt-get install -y apt-transport-https ca-certificates && apt-get install -y software-properties-common
# Clone libsdl-org/sdlwiki repository
RUN apt-get install -y git
RUN git clone https://github.com/libsdl-org/sdlwiki.git --depth 1

COPY *.go ./
RUN go mod init sdlwiki-annotations

# Run the annotations extraction script
RUN go run .