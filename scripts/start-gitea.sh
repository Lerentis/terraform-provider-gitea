#!/bin/bash -e

WORK_DIR=`pwd`
GITEA_SDK_TEST_USERNAME=test01
GITEA_SDK_TEST_PASSWORD=test01

mkdir -p ${WORK_DIR}/test/conf/ ${WORK_DIR}/test/data/
wget --quiet "https://dl.gitea.io/gitea/main/gitea-main-linux-amd64" -O ${WORK_DIR}/test/gitea-main
chmod +x ${WORK_DIR}/test/gitea-main
echo "[security]" > ${WORK_DIR}/test/conf/app.ini
echo "INTERNAL_TOKEN = eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE1NTg4MzY4ODB9.LoKQyK5TN_0kMJFVHWUW0uDAyoGjDP6Mkup4ps2VJN4" >> ${WORK_DIR}/test/conf/app.ini
echo "INSTALL_LOCK   = true" >> ${WORK_DIR}/test/conf/app.ini
echo "SECRET_KEY     = 2crAW4UANgvLipDS6U5obRcFosjSJHQANll6MNfX7P0G3se3fKcCwwK3szPyGcbo" >> ${WORK_DIR}/test/conf/app.ini
echo "PASSWORD_COMPLEXITY = off" >> ${WORK_DIR}/test/conf/app.ini
echo "[database]" >> ${WORK_DIR}/test/conf/app.ini
echo "DB_TYPE = sqlite3" >> ${WORK_DIR}/test/conf/app.ini
echo "LOG_SQL = false" >> ${WORK_DIR}/test/conf/app.ini
echo "[repository]" >> ${WORK_DIR}/test/conf/app.ini
echo "ROOT = ${WORK_DIR}/test/data/" >> ${WORK_DIR}/test/conf/app.ini
echo "[server]" >> ${WORK_DIR}/test/conf/app.ini
echo "ROOT_URL = http://127.0.0.1:3000" >> ${WORK_DIR}/test/conf/app.ini
echo "DISABLE_ROUTER_LOG=false" >> ${WORK_DIR}/test/conf/app.ini
echo "[log]" >> ${WORK_DIR}/test/conf/app.ini
echo "COLORIZE=false" >> ${WORK_DIR}/test/conf/app.ini
${WORK_DIR}/test/gitea-main migrate -c ${WORK_DIR}/test/conf/app.ini
${WORK_DIR}/test/gitea-main admin create-user --username=${GITEA_SDK_TEST_USERNAME} --password=${GITEA_SDK_TEST_PASSWORD} --email=test01@gitea.io --admin=true --must-change-password=false --access-token -c ${WORK_DIR}/test/conf/app.ini | grep "Access token" > .token
${WORK_DIR}/test/gitea-main web -c ${WORK_DIR}/test/conf/app.ini &
sleep 3