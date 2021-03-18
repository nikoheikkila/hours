# Installation

**NOTE:** Hours requires a Toggl Track account and an [API key][api] which you can set up for free. Furthermore, you should have enough privileges on your account since Hours fetches time entries by looking up from Toggl workspaces.

## Option 1: Download Binaries

Go to the [**releases**][releases] page and download the preferred archive. Each release is compiled for **macOS**, **Linux**, and **Windows** platforms (32-bit & 64-bit).

## Option 2: Use Go

Make sure you have at least **Go 1.15+** installed (Tip: run `brew install go`) and `$GOPATH/bin` is included in your `$PATH`.

```sh
go get github.com/nikoheikkila/hours
```

## Configuring

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

[api]: https://track.toggl.com/profile
[releases]: https://github.com/nikoheikkila/hours/releases/latest
