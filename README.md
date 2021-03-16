# ⏰ Hours

Working as a consultant requires you to mark your billed project hours periodically for invoicing. This is boring but luckily [**Toggl**][toggl] is there to help. Unfortunately, not everyone is using it directly.

**Hours** is a command-line tool built with **Go** for exporting the latest Toggl project entries from your workspace and printing a _nice_ report of them. Hand it over to your account manager or client.

> **NOTE:** Requires a Toggl Track account and an [API key][api] which you can set up for free. Furthermore, you should have enough privileges on your account since Hours fetches time entries by looking up from Toggl workspaces.

## Installation

### Option 1: Download Binaries

Go to the [**releases**][releases] page and download the preferred archive. Each release is compiled for **macOS**, **Linux**, and **Windows** platforms (32-bit & 64-bit).

### Option 2: Use Go

Make sure you have at least **Go 1.15+** installed (Tip: run `brew install go`) and `$GOPATH/bin` is included in your `$PATH`.

```sh
go get github.com/nikoheikkila/hours
```

## Usage

First, configure the tool by creating a `~/.togglrc` file from the template below and replace dummy values with yours (note the YAML syntax).

```yml
# Find this from your Toggl user settings
api_token: your-toggl-api-token

# Find this under 'Organization -> Workspace' (copy it from the address bar)
workspace_id: your-toggl-workspace-id
```

Alternatively, you can use the environment variables `TOGGL_API_TOKEN` and `TOGGL_WORKSPACE_ID` if you don't want to use a configuration file.

After installing, run it from your shell like so:

```sh
hours <FLAGS>
```

```plain
Options:

--output <string>
    Output format for the reporter (default: 'text').

--since <string>
    ISO compatible date where to begin listing time entries (default: yesterday).

--until <string>
    ISO compatible date where to stop listing time entries (default: today).

--ansi=<bool>
    Set to false if you don't want text reporter to style its output (default: true).
```

## Reporters

**Hours** uses reporters which are plugins accepting a list of time entries and printing them to the terminal in suitable format. Availabe reporters are described below.

### Plain text

**Examples:**

Formats and styles the output with ANSI styles using the [_termenv_][termenv] library, unless the flag `--ansi=false` is specified.

```sh
hours
hours --ansi=false
hours --output text
```

### Markdown

**Examples:**

Pipe to e.g. [_Pandoc_][pandoc] for additional processing, or to [_Glow_][glow] for fancier formatting.

```sh
hours --output markdown
hours --output markdown > report.md
hours --output markdown | glow -
hours --output markdown | pandoc
```

### JSON

**Examples:**

Pipe to e.g. [_jq_][jq] for additional processing.

```sh
hours --output json
hours --output json > report.json
hours --output json | jq '.'
```

### CSV

**Examples:**

Save the output to a file and import to Excel, or wherever you need it.

```sh
hours --output csv
hours --output csv > report.csv
```

[toggl]: https://toggl.com/
[api]: https://track.toggl.com/profile
[termenv]: https://github.com/muesli/termenv
[pandoc]: https://pandoc.org/
[glow]: https://github.com/charmbracelet/glow
[jq]: https://stedolan.github.io/jq/
[releases]: https://github.com/nikoheikkila/hours/releases/latest
