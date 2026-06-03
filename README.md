# pkgrun

Interactive CLI to run npm scripts using the package manager detected from your project's lockfile.

## What it does

1. Detects which Node package managers match lockfiles in the current directory (priority: Bun → pnpm → Yarn → npm)
2. Picks the first detected manager that is installed on your `PATH`
3. Reads `package.json` scripts from the current directory
4. Shows an interactive terminal menu to pick a script
5. Runs `<manager> run <script>`

## Requirements

- Go 1.26+
- A Node project with `package.json` in the current working directory
- At least one supported lockfile: `bun.lock`, `pnpm-lock.json`, `yarn.lock`, or `package-lock.json`
- The matching package manager binary on your `PATH`

## Install

```bash
go install github.com/proudynyu/pkgrun@latest
```

Or build from source:

```bash
git clone https://github.com/proudynyu/pkgrun.git
cd pkgrun
go build -o pkgrun .
```

## Usage

From a Node project root:

```bash
pkgrun
```

Example output:

```
-> Package managers identified: [yarn]
-> Using package manager: yarn
Select a script (↑↓ to move, Enter to confirm):
 |> build
  dev
  lint
```

After you confirm with Enter, pkgrun runs `yarn run <chosen-script>`.

## Lockfile priority

When multiple lockfiles exist, managers are listed in this order:

| Lockfile            | Manager |
|---------------------|---------|
| `bun.lock`          | bun     |
| `pnpm-lock.json`    | pnpm    |
| `yarn.lock`         | yarn    |
| `package-lock.json` | npm     |

The first manager in that list that is both detected **and** installed is used.

## Project layout

```
pkgrun/
├── main.go           # CLI entrypoint
├── src/
│   ├── app/          # Orchestration and exit codes
│   ├── cmd/          # Lockfile detection
│   ├── cwd/          # Working directory helpers
│   ├── file/         # package.json parsing
│   └── ui/           # Interactive script selector
└── go.mod
```

## Development

Build:

```bash
go build -o pkgrun .
```
Run tests:

```bash
go test ./...
```

With coverage:

```bash
go test ./... -cover
```

## Platform support

- **Script execution**: cross-platform (Go `os/exec`)
- **Interactive menu**: currently uses Linux `stty` on `/dev/tty`; macOS/Windows support is planned

Run from the directory that contains `package.json`. Parent-directory lookup is not supported yet.

## Roadmap

- [ ] Cross-platform terminal raw mode
- [ ] Walk up to find `package.json`
- [ ] CI with `go test` and `go vet`
- [ ] Monorepo / workspace support
