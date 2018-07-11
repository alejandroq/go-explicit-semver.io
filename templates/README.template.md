# Go Explicit Semver

_v{{ .SrcVersion }} Alejandro Quesada_

**Very early work in progress. Feel free to contribute ideas!**

## Goals

Application that can manage the semantic versioning of an application via a ledger explicitly called by a user via the CLI. The tool does the following things:

- [ ] Command line operations:
  - [ ] Viewable History
  - [x] Initilizable configuration
  - [ ] Tool being incited will quickly scope opted child directories/files for Patch release versioning based on diffs
  - [ ] Commands to increment Major, Minor or custom (rc, etc) versions for directories AND OR individual files (say wireframes, etc)
- [ ] Git tagging (need to decide how to do this when there are multiple versioned artifacts...)

_For design ideas, view the [TODO.md](./TODO.md)_

This application is meant to reduce the overhead of this particular task in a simple and repeatable way. The value derived from the activity is evidently significant when paired with appropriate verifiers (ex. test suite(s)).

## What this is not

This is not a versioning control system like Git. Simply a tool to allow a user to explicitly set a semantic version and have this information propogate via Git tags and documentation. The tool is smart enough to determine WHEN source files have changed for incrementing patch releases.

## Future use cases in-mind

- Editing and incrementing CloudFormation documents
- Developing an iOS app
- For versioning THIS tool
- Versioning multiple documents for a project from: source code, design docs, etc

### How will be handled by versioning where there are multiple authors

... tbd guid ...

## Templating system

Variables for templates are: ...

See [https://golang.org/pkg/html/template/](https://golang.org/pkg/html/template/).

## `.v/explicit-config.json`

Version {{ .SrcVersion }} example:

```json
{
  "versioning": [
    {
      "source": "src",
      "primary": true
    },
    {
      "source": "README.md",
      "primary": false
    },
    {
      "source": "TODO.md",
      "primary": false
    }
  ],
  "templates": []
}
```

## Artifacts

A `.v` directory that maintains various files pertaining to the state of the application. This directory is meant to be versioned (by Git or another client) for consistency between developers.
