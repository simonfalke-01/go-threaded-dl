# <img align="center" src="https://github.com/simonfalke-01/go-threaded-dl/blob/main/images/gdl_icon.webp?raw=true" width=300>
![go_badge](https://github.com/simonfalke-01/go-threaded-dl/actions/workflows/go.yml/badge.svg) ![go_macos_badge](https://github.com/simonfalke-01/go-threaded-dl/actions/workflows/go-macos.yml/badge.svg) ![release](https://badgen.net/github/release/simonfalke-01/go-threaded-dl) </br>
Multi-threaded content downloader written in Go. By downloading a file with multiple connections (multiple threads), download speed can be drastically improved. </br>
Downloading with 100 threads: </br>
<img align="center" src="https://github.com/simonfalke-01/go-threaded-dl/blob/main/images/100_threads.png?raw=true">

## Usage
Provide URL. Threads and save path are optional. Default is 10 threads and the current directory.
```
gdl <url>
gdl <url> -t <threads>
gdl <url> -o <save-path>
gdl <url> -t <threads> -o <save-path>
```

### Example
```
gdl https://do-spaces-1.simonfalke.studio/Hello\! -t 20 -o ~/Hello\!
```

## Installation
go-threaded-dl can be installed on your system by running one of the commands below in your terminal.
An installation script will be run. You must have git, and either `curl` or `wget` installed.</br>
Latest builds can be found under GitHub Actions [here](https://github.com/simonfalke-01/go-threaded-dl/actions).

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

### Building
A version of Go must be installed. Clone the repository and run:
```
go build -v -o gdl .
```

## License
This project is licensed under there are literally no licenses to this (made by me). You can do whatever you want with this, but I would appreciate if you would give me credit.
