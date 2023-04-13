# go-threaded-dl
Multi-threaded content downloader written in Go

## Usage
```
gtd <url> <threads> <output>
```

### Example
```
gtd https://example.com/file.zip 4 ~/file.zip
```

## Installation
go-threaded-dl can be installed on your system by running one of the commands below in your terminal.
An installation script will be run. You must have git, and either `curl` or `wget` installed.

### Commands
| Method    | Command                                                                                            |
| :-------- |:---------------------------------------------------------------------------------------------------|
| **curl**  | `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/simonfalke-01/go-threaded-dl/main/install.sh)"` |
| **wget**  | `/bin/bash -c "$(wget -O- https://raw.githubusercontent.com/simonfalke-01/go-threaded-dl/main/install.sh)"`   |
| **fetch** | `/bin/bash -c "$(fetch -o - https://raw.githubusercontent.com/simonfalke-01/go-threaded-dl/main/install.sh)"` |

### Inspection
If you would like to inspect the script before running it, you can download the script by running the following command:
```bash
wget https://raw.githubusercontent.com/simonfalke-01/go-threaded-dl/main/install.sh
```
After you are done, you can run the script with
```bash
chmod +x ./install.sh && ./install.sh
```

## License
This project is licensed under there are literally no licenses to this (made by me). You can do whatever you want with this, but I would appreciate if you would give me credit.