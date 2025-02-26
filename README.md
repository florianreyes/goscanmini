# GoScan-Mini
---

This is a small project to learn and practice Go, it scans ports and saves them to the clipboard.

## Installation

```go
// Clone the repository
git clone https://github.com/florianreyes/goscanmini

// cd into the repo
cd goscanmini

// cd into scripts
cd scripts

// Run the installator
./install.sh
```
- When prompted to move the file to /usr/local/bin you might need to input sudo password
- If `n` selected, the executable will be moved to the root of the `goscanmini` directory

## Usage

```go
goscanm -host [IP] -p [P1,P2,P3...PN] -max-threads [THREAD-NUM]
```
### Example

```go
goscanm -host 10.10.10.8 -p 80,443 -max-threads 50
```

----

- If ports not indicated it defaults to top 100 ports
- Flag `--all` to scan all 65535 ports.