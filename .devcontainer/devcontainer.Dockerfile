# The Go setup is based on the [Medium article by Quentin McGaw](https://medium.com/@quentin.mcgaw/ultimate-go-dev-container-for-visual-studio-code-448f5e031911) from Sep 13, 2019.
# The node setup is based on the [Medium article by Csaba Apagyi](https://medium.com/geekculture/how-to-install-a-specific-node-js-version-in-an-alpine-docker-image-3edc1c2c64be) from Dec 7, 2021.

ARG GO_VERSION
ARG ALPINE_VERSION
ARG NODE_VERSION

FROM node:${NODE_VERSION}-alpine${ALPINE_VERSION} as node
FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION}
ARG NPM_VERSION
ARG GOPLS_VERSION
ARG GO_AIR_VERSION
ARG GO_TEMPL_VERSION
ARG GO_MIGRATE_DB_TAG
ARG GO_MIGRATE_VERSION
ARG USERNAME=vscode
ARG USER_UID=1000
ARG USER_GID=1000

# Node setup

COPY --from=node /usr/lib /usr/lib
COPY --from=node /usr/local/share /usr/local/share
COPY --from=node /usr/local/lib /usr/local/lib
COPY --from=node /usr/local/include /usr/local/include
COPY --from=node /usr/local/bin /usr/local/bin
COPY --from=node /opt /opt

# User setup

RUN adduser $USERNAME -s /bin/sh -D -u $USER_UID $USER_GID && \
    mkdir -p /etc/sudoers.d && \
    echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME && \
    chmod 0440 /etc/sudoers.d/$USERNAME

# install packages
RUN apk add -q --update --progress --no-cache git sudo openssh-client zsh nano postgresql16-client

# updating npm
RUN npm install -g npm@$NPM_VERSION

# installing gopls
RUN go install golang.org/x/tools/gopls@$GOPLS_VERSION

# installing cosmtrek/air for hot reloading
RUN go install github.com/cosmtrek/air@$GO_AIR_VERSION

# installing a-h/templ for templating
RUN go install github.com/a-h/templ/cmd/templ@$GO_TEMPL_VERSION

# intalling golang-migrate/migrate for database migrations
RUN go install -tags '${GO_MIGRATE_DB_TAG}' github.com/golang-migrate/migrate/v4/cmd/migrate@$GO_MIGRATE_VERSION

# Setup shell
USER $USERNAME
RUN sh -c "$(wget -O- https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)" "" --unattended &> /dev/null

ENV ENV="/home/$USERNAME/.ashrc"
ENV ZSH=/home/$USERNAME/.oh-my-zsh
ENV EDITOR=nano
ENV LANG=en_US.UTF-8

RUN echo 'ZSH_THEME="robbyrussell"' >> "/home/$USERNAME/.zshrc" \
    && echo 'ENABLE_CORRECTION="false"' >> "/home/$USERNAME/.zshrc" \
    && echo 'plugins=(git copyfile extract colorize dotenv encode64 golang)' >> "/home/$USERNAME/.zshrc" \
    && echo 'source $ZSH/oh-my-zsh.sh' >> "/home/$USERNAME/.zshrc"
RUN echo "exec `which zsh`" > "/home/$USERNAME/.ashrc"
USER root

# GOPATH pkg directory error fix
RUN chmod -R a+w /go/pkg

EXPOSE 3000