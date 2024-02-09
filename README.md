
---

# Website Health Checker

The Website Health Checker is a small command-line tool written in Go that allows users to check the health status of a website. It performs a simple health check similar to the `ping` command but specifically tailored for websites.

## Installation

To use the Website Health Checker, you can either compile the source code yourself or download the precompiled executable for your platform.

### Compiling from Source

1. Make sure you have Go installed on your system. If not, you can download and install it from [the official Go website]([(https://go.dev/)]).
2. Clone this repository to your local machine.
   ```
   git clone https://github.com/nishantnz/go-sitestatus
   ```
3. Navigate to the project directory.
   ```
   cd go-sitestatus
   ```
5. Compile the source code to generate the executable.
   ```
   go build -o WUD.exe
   ```

### Using Precompiled Executable

1. Download the precompiled executable(WUD.exe) for your platform from the [Releases](https://github.com/nishantnz/go-sitestatus/releases/tag/v.1) page.

## Usage

Once you have the executable installed, you can use the following command to check the health status of a website.
#### Note: only run the exe file from the directory where you have installed it.

```
WUD.exe --domain example.com
```

Replace `example.com` with the domain name you want to check. By default, the tool will use port 80 for the health check. If you want to specify a different port, you can use the `--port` flag:

```
WUD.exe --domain example.com --port 443
```

This command will check the health of the website `example.com` on port `443`.

### Command-Line Options

- `--domain` or `-d`: Specifies the domain name to check (required).
- `--port` or `-p`: Specifies the port number to use for the health check (optional, default is port 80).

---
