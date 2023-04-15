#!/usr/bin/env bash

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  curl -L "https://github.com/simonfalke-01/go-threaded-dl/releases/latest/download/gdl-linux" -o /usr/local/bin/gdl
elif [[ "$OSTYPE" == "darwin"* ]]; then
  curl -L "https://github.com/simonfalke-01/go-threaded-dl/releases/latest/download/gdl-darwin" -o /usr/local/bin/gdl
fi

chmod +x /usr/local/bin/gdl

if [ -e /usr/local/bin/gtd ]; then
    read -p "A legacy version of go-threaded-dl is found at /usr/local/bin/gtd, do you want to remove it? [Y/n] " -r answer
    if [[ "$answer" =~ ^[Yy]$ ]]; then
        rm /usr/local/bin/gtd
    elif [[ "$answer" =~ ^[Nn]$ ]]; then
        echo "/usr/local/bin/gtd will not be removed."
    elif [ -z "$answer" ]; then
        rm /usr/local/bin/gtd
    else
        echo "Aborted."
        exit 1
    fi
fi

echo "Installed gdl to /usr/local/bin/gdl."