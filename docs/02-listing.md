# Listing Time Entries

**Hours** can fetch your latest time entries, which can be easily processed by external scripts suitable for many use cases.

## Reporters

Reporters are plugins accepting a list of time entries and printing them to the terminal in suitable format. Availabe reporters are described below.

### Plain text

**Examples:**

Formats and styles the output with ANSI styles using the [_termenv_][termenv] library, unless the flag `--ansi=false` is specified.

```sh
hours list
hours list --ansi=false
hours list --output text
```

### Markdown

**Examples:**

Pipe to e.g. [_Pandoc_][pandoc] for additional processing, or to [_Glow_][glow] for fancier formatting.

```sh
hours list --output markdown
hours list --output markdown > report.md
hours list --output markdown | glow -
hours list --output markdown | pandoc
```

### JSON

**Examples:**

Pipe to e.g. [_jq_][jq] for additional processing.

```sh
hours list --output json
hours list --output json > report.json
hours list --output json | jq '.'
```

### CSV

**Examples:**

Save the output to a file and import to Excel, or wherever you need it.

```sh
hours list --output csv
hours list --output csv > report.csv
```

[termenv]: https://github.com/muesli/termenv
[pandoc]: https://pandoc.org/
[glow]: https://github.com/charmbracelet/glow
[jq]: https://stedolan.github.io/jq/
