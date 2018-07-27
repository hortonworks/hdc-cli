#!/usr/bin/env bash

: ${GIT_VERSION:=latest}
: ${GIT_TAG:=latest}
: ${BASE_URL:=https://127.0.0.1}

set-mock-image-tag() {
    echo "GitHub First Parent Tag is: ${GIT_VERSION}"
    echo "GitHub Tag is: ${GIT_TAG}"
    echo "CircleCI Branch is: ${CIRCLE_BRANCH}"
    echo "CircleCI Tag is: ${CIRCLE_TAG}"

    if [[ "${GIT_VERSION}" != *"-rc."* ]]; then
        export GIT_VERSION=latest
    fi
}

set-mock-image-tag

docker-compose -f docker-compose.yml -p cbreak up -d

sleep 30s