#!/bin/bash

INFILE="swagger_en.json"
WORKINGFILE="swagger.temp.json"
OUTFILE="swagger.new.json"
SEDBACKUP=".bk"
GITUSERID="dozyio"
GITREPOID="loopring-dex-go-client/client/go"

export WORKINGFILE

cp ${INFILE} ${WORKINGFILE}

# fix 0 ints
sed -i ${SEDBACKUP} 's/"0[lL]"/0/g' ${WORKINGFILE}

# fix defaults
sed -i ${SEDBACKUP} 's/"type": "integer", "default": "None"/"type": "integer", "default": 0/g' ${WORKINGFILE}
sed -i ${SEDBACKUP} 's/"type": "integer", "default": "None "/"type": "integer", "default": 0/g' ${WORKINGFILE}
sed -i ${SEDBACKUP} 's/"type": "boolean", "default": "None"/"type": "boolean", "default": false/g' ${WORKINGFILE}
sed -i ${SEDBACKUP} 's/"type": "string", "default": "None"/"type": "string", "default": ""/g' ${WORKINGFILE}

# fix 0 response
sed -i ${SEDBACKUP} 's/"responses": {"0"/"responses": {"200"/g' ${WORKINGFILE}

# fix operationId
php operationId.php

# clean up sed backup
rm ${WORKINGFILE}${SEDBACKUP}

jq . ${WORKINGFILE} > ${OUTFILE}

rm ${WORKINGFILE}

# build Go client
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i /local/${OUTFILE} -g go -o /local/client/go --package-name loopring --git-user-id ${GITUSERID} --git-repo-id ${GITREPOID}
gofmt -w ./client/go
