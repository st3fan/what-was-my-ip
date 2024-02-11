# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/

FROM golang:1.22 AS build
WORKDIR /usr/local/src
COPY . .
RUN CGO_ENABLED=0 go build -v

FROM alpine:3.19.1
COPY --from=build /usr/local/src/what-was-my-ip /usr/local/bin/what-was-my-ip
ENTRYPOINT ["/usr/local/bin/what-was-my-ip"]
