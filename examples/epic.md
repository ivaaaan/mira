# Create a tool that can convert Markdown text to Jira issues

We want to create a CLI tool that can convert Markdown outline to Jira issues.
Each heading in Markdown text will represent an issue in Jira. Nesting of issues can be achieved by using levels of Markdown heading.
The future versions we will also add support for labels.

- Hello world!
- Test
- Test

## Create a skeleton of the project

Initialize a skeleton of the project. We will use Go for this project.
Research and add all necessary tooling you think is needed. `cobra` and `viper` are good choices for CLI programs.

Here we go with some links: [Duck Duck Go](https://duckduckgo.com).

Does it support lists?

- I
- have
- no
- idea

### Create a Git repo

### Create Confulence page with docs

## Create parser

Parser should parse issues from markdown file and create a tree of `Task` structs, which can be later processed and pushed to Jira.

### Parse tree of tickets

