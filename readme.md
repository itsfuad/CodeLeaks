# CodeLeaks

CodeLeaks is a Golang project designed to help developers identify and manage potential code leaks in their applications. This tool scans your codebase and provides insights into areas where sensitive information might be exposed.

## Features

- **Code Scanning**: Scans your codebase for potential leaks.
- **Reporting**: Generates detailed reports on identified issues.
- **Integration**: Easily integrates with CI/CD pipelines.

## Installation

To install CodeLeaks, use the following command:

```bash
go get github.com/itsfuad/CodeLeaks
```

## Usage

To use CodeLeaks, run the following command in your project directory:

```bash
Usage
  -d string
        Directory to scan (required)
  -e string
        Comma-separated list of files to exclude (e.g., file1.txt,file2.txt)
  -ex string
        Comma-separated list of extensions to exclude (e.g., .go,.py)
  -h    Show usage information
  -o string
        Comma-separated list of files to scan (e.g., file1.txt,file2.txt)
  -x string
        Comma-separated list of extensions to scan (e.g., .go,.py)
```

## Contact

For any questions or feedback, please open an issue on our [GitHub repository](https://github.com/itsfuad/CodeLeaks).
