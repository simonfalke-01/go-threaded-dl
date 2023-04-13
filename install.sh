#!/usr/bin/env bash

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  curl -L "https://github.com/simonfalke-01/go-threaded-dl/releases/download/v0.0.1/gtd-linux" -o /usr/local/bin/gtd
elif [[ "$OSTYPE" == "darwin"* ]]; then
  curl -L "https://github.com/simonfalke-01/go-threaded-dl/releases/download/v0.0.1/gtd-darwin" -o /usr/local/bin/gtd
fi

chmod +x /usr/local/bin/gtd

echo "Installed gtd to /usr/local/bin/gtd."