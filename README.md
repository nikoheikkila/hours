# ⏰ Hours

Working as a consultant requires you to mark your billed project hours periodically for invoicing. This is boring but luckily [**Toggl**][toggl] is there to help. Unfortunately, not everyone is using it directly.

**Hours** is a command-line tool built with **Go** for exporting the latest Toggl project entries from your workspace and printing a _nice_ report of them. Hand it over to your account manager or client.

> **NOTE:** Requires a Toggle account and an [API key][api] which you can set up for free. Furthermore, you should have enough privileges on your account since Hours fetches time entries by looking up from Toggl workspaces.

## Installation

Make sure you have at least **Go 1.15+** installed (Tip: run `brew install go`) and `$GOPATH/bin` is included in your `$PATH`.

```sh
go get github.com/nikoheikkila/hours
```

## Usage

Currently, using **Hours** requires the following environment variables:

- `TOGGL_API_TOKEN` – API token from your [Toggl settings][api]
- `TOGGL_WORKSPACE_ID` – Numeric workspace ID found in Toggl under `Organization -> Workspaces` (copy it from the address bar)

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

**Hours** uses reporters which are plugins accepting a list of time entries and printing them to the terminal in suitable format. As of now, available and upcoming reporters are:

- [x] Plain text (allows formatting the output with ANSI styles using the [_termenv_][termenv] library)
- [ ] Markdown (coming soon)
- [ ] CSV (coming soon)
- [ ] JSON (coming soon)

[toggl]: https://toggl.com/
[api]: https://track.toggl.com/profile
[termenv]: https://github.com/muesli/termenv
