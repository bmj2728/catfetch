# CatFetch

[![Go Version](https://img.shields.io/badge/Go-1.25-00ADD8?style=flat&logo=go)](https://go.dev/)
[![Tests](https://github.com/bmj2728/catfetch/actions/workflows/test.yml/badge.svg)](https://github.com/bmj2728/catfetch/actions/workflows/test.yml)
[![Release](https://github.com/bmj2728/catfetch/actions/workflows/release.yml/badge.svg)](https://github.com/bmj2728/catfetch/actions/workflows/release.yml)
[![codecov](https://codecov.io/gh/bmj2728/catfetch/branch/main/graph/badge.svg)](https://codecov.io/gh/bmj2728/catfetch)

A desktop GUI application built with Go and Gio UI that fetches and displays random cat images from [cataas.com](https://cataas.com/).

## Gallery

<p align="center">
  <img src="screenshots/cat_fetch_prefetch.png" alt="CatFetch startup screen" width="45%">
  <img src="screenshots/cat_fetch_with_catpic.png" alt="CatFetch with cat image" width="45%">
</p>

## Prerequisites

- Go 1.25 or higher

## Getting Started

### Clone the Repository

```bash
git clone <repository-url>
cd catfetch
```

### Build the Application

```bash
go build -o build/catfetch ./cmd/catfetch
```

### Run the Application

From the project directory:

```bash
./build/catfetch
```

Or run directly without building:

```bash
go run ./cmd/catfetch/main.go
```

### Optional: Install System-wide

To run the application from anywhere without `./`:

**Linux/macOS (user-local install):**
```bash
cp build/catfetch ~/bin/catfetch
# Ensure ~/bin is in your PATH
```

**Linux (system-wide install):**
```bash
sudo cp build/catfetch /usr/local/bin/catfetch
```

Then you can run it from anywhere:
```bash
catfetch
```

## Usage

Launch the application and click the "Fetch Image" button to load a random cat picture. The image will automatically scale to fit the window while maintaining its aspect ratio.

## Roadmap

- **Cat History**: Browse previously fetched cat images
- **Text Overlays**: Add custom text overlays to cat images
- **Tag Search**: Search for cats by specific tags
- **Image Filters**: Add sliders and options to apply filters (sepia, blur, brightness, etc.) using cataas API parameters

## Development

### Pre-commit Hooks

This project uses pre-commit hooks for automated code quality checks. To set up:

1. **Install pre-commit:**
   ```bash
   # macOS
   brew install pre-commit

   # Linux
   pip install pre-commit
   # or
   pipx install pre-commit
   ```

2. **Install the git hooks:**
   ```bash
   pre-commit install
   ```

3. **Run hooks manually (optional):**
   ```bash
   pre-commit run --all-files
   ```

The hooks will automatically run on every commit and check for:
- Go code formatting
- Linting with golangci-lint
- Secret detection with gitleaks
- Common issues (trailing whitespace, merge conflicts, etc.)

### Running Tests

```bash
# Run all tests
go test ./...

# Run with coverage
go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

# View coverage report
go tool cover -html=coverage.out
```

## License

MIT License - see [LICENSE](LICENSE) file for details.

Copyright (c) 2026 NovelGit LLC