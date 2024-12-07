#!/usr/bin/env bash

ssh github_runner@$1 "sudo systemctl stop go-web-site"
scp -r ~/build github_runner@${{ secrets.SSH_HOST }}:~
chmod -R +x ~/build
ssh github_runner@$1 "sudo systemctl start go-web-site"