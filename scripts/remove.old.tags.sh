#!/usr/bin/env bash

for tag in $(git tag --sort=-v:refname | head -n 10); do
  git push --delete origin "$tag"
  git tag -d "$tag"
done
