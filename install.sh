#!/usr/bin/env bash

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  curl -L "https://github.com/simonfalke-01/go-threaded-dl/releases/download/v1.0.0/gdl-linux" -o /usr/local/bin/gdl
elif [[ "$OSTYPE" == "darwin"* ]]; then
  curl -L "https://github.com/simonfalke-01/go-threaded-dl/releases/download/v1.0.0/gdl-darwin" -o /usr/local/bin/gdl
fi

if [[ -f "/usr/local/bin/gtd" ]]; then
  read -p "A legacy version of go-threaded-dl is found at /usr/local/bin/gtd, do you want to remove it? [y/N] " -n 1 -r
  echo
  if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Installation aborted."
    exit 1
  fi
fi

chmod +x /usr/local/bin/gdl

echo "Installed gtd to /usr/local/bin/gdl."