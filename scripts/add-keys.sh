#!/usr/bin/env bash

mkdir ~/.ssh/
touch ~/.ssh/id_rsa
chmod 600 ~/.ssh/id_rsa
echo $1 > ~/.ssh/id_rsa
ssh-keyscan -t rsa -H $2 >> ~/.ssh/known_hosts
echo "Host $2\n  AddKeysToAgent yes\n  IdentityFile ~/.ssh/id_rsa\n" > ~/.ssh/config