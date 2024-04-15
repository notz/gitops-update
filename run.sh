#!/bin/sh -l
set -e

FILE_NAME=$1
KEY=$2
VALUE=$3
GITHUB_DEPLOY_KEY=$4
GITHUB_ORG_AND_REPO=$5
COMMIT_MESSAGE=$6

mkdir -p ~/.ssh

cat <<EOF >~/.ssh/config
Hostname github.com
IdentityFile ~/.ssh/id_rsa
EOF

ssh-keyscan -t rsa github.com > ~/.ssh/known_hosts

git config --global user.email "41898282+github-actions[bot]@users.noreply.github.com"
git config --global user.name "github-actions[bot]"

# The key needs to be wrapped in double quotes
echo "$GITHUB_DEPLOY_KEY" > ~/.ssh/id_rsa
chmod 600 ~/.ssh/id_rsa

rm -rf $RUNNER_TEMP/infra-as-code-repo
git clone git@github.com:$GITHUB_ORG_AND_REPO.git $RUNNER_TEMP/infra-as-code-repo
wget https://raw.githubusercontent.com/notz/gitops-update/master/replace-key.py
python3 replace-key.py --file $RUNNER_TEMP/infra-as-code-repo/$FILE_NAME --key $KEY --value $VALUE
cd $RUNNER_TEMP/infra-as-code-repo
git add .
if ! git diff-index --quiet HEAD; then
  git commit -m "$COMMIT_MESSAGE"
  git push
fi
