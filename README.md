# Mira

With Mira, you can write your Jira epics in Markdown files and then create Jira issues from them.

## Concepts

This is an MVP, and concepts might change.

### One file per root issue

You should create one file for each issue you want to write. However, you can create child issues within that file.

Example:

```markdown

# Root issue

Root issue description

## Child issue 2

Child issue description

## Child issue 3
```

### Nesting

In theory, you can nest as many issues as you want. However, Jira's provider implementation supports only two levels at the moment.

Your root issue will be created with type `Epic`, and all nested issues with type `Task`. `Subtask` are not supported at the moment, but I will add them eventually.

### Parsers and providers

Right now only Jira and Markdown are supported. However, there is a room for extending with `Provider` and `Parser` interfaces.


## Usage

```
Usage:
  mira [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  push        Push tasks from a file to provider

Flags:
  -c, --config string   Path to a user config
  -h, --help            help for mira

Examples:
  mira push myepic.md
```

## Config

`~/.config/mira/config.toml`:

```toml
[jira]
url = "https://mycompany.atlassian.net/"
username = "<your email>"
api_token = "<your api token>"
project_key = "<your project key>"
```

## TODO

- [ ] Tests coverege
- [ ] Convert Markdown text to Jira format
- [ ] Support subtasks in Jira
- [ ] Login command
- [ ] Goreleaser
