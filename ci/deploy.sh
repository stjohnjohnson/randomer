#!/bin/sh -e

# Swapping between PR and Prod
[ "$SD_PULL_REQUEST" != "" ] && export HEROKU_APP=$HEROKU_APP_PR

mkdir ~/.ssh
echo Adding Heroku.com to Known Hosts
ssh-keyscan -H heroku.com >> ~/.ssh/known_hosts

echo Validating fingerprint
ssh-keygen -l -f ~/.ssh/known_hosts | grep $HEROKU_FINGERPRINT

echo Adding RSA key
echo $HEROKU_SSH > ~/.ssh/id_rsa
chmod 600 ~/.ssh/*

echo Deploying to Heroku $HEROKU_APP
git push -f "git@heroku.com:$HEROKU_APP.git" HEAD:master

export HEROKU_URL=https://$HEROKU_APP.herokuapp.com/
echo Deployed to $HEROKU_URL
