# Go Explicit Semver

_v{{ .SrcVersion }} Alejandro Quesada_

**Very early work in progress. Feel free to contribute ideas!**

> Don't boil the ocean

## Goals

Application that can manage the semantic versioning of an application via a ledger explicitly called by a user via the CLI. The tool does the following things:

- [ ] Command line operations:
  - [ ] Viewable History
  - [x] Initilizable configuration
  - [ ] Tool being incited will quickly scope opted child directories/files for Patch release versioning based on diffs (ex. `v`)
  - [ ] Commands to increment Major, Minor or custom (rc, etc) versions for directories AND OR individual files (say wireframes, etc)
- [ ] Git tagging
  - [ ] Git tag for each versioned artifact

_For design ideas, view the [TODO.md](./TODO.md)_

This application is meant to reduce the overhead of this particular task in a simple and repeatable way. The value derived from the activity is evidently significant when paired with appropriate verifiers (ex. test suite(s)).

## What this is not

This is not a versioning control system like Git. Simply a tool to allow a user to explicitly set a semantic version and have this information propogate via Git tags and documentation. The tool is smart enough to determine **WHEN** source files have changed for the incrementing of patch releases.

## Future use cases in-mind

- Infrastructure as Code
  - Situation: When actively developing or maintaining an AWS CloudFormation document.
  - Oppurtunity: Track vanity mutations to your Infrastructure as Code. If team A is running version 1.0.0 and you have added a CloudFront dashboard for 1.1.0, quickly ascertain that they require a pending update. Or if there is an error, the process for a Cloud Engineer to replicate team A's infrastructure is simplified.
- Mobile development
  - Situation: Developing an iOS application.
  - Oppurtunity: Pushing changes to the app store with accurately maintained versions. Be able to more quickly alleviate user woes when their clients break on version "X" by having had explicitly maintained Git tags.
- Swagger development
  - Situation: You are defining your REST service with Swagger
  - Oppurtunity: Save key strokes and accurately maintain your Swagger version with this tool. Edit the Swagger document's template and let this tool update version information. Deploy the Templates `output` artifact.
- For versioning THIS tool
  - Oppurtunity: To test the idea in creating the idea. Very organic development.
- Versioning multiple documents for a project from: source code, design docs, etc
  - Oppurtunity: I like the monorepo, so meeting the usecase is important to me. I occasinoaly like to keep non-code documents in the repository as well for XYZ reason.
- Function as an I/O argument for `-ldflag` during a `go build`. Similar to Git hash invocations in said usecase.

### How will be handled by versioning where there are multiple authors

... tbd guid ... Use a simple list

## Templating system

Variables for templates are: ...

See [https://golang.org/pkg/html/template/](https://golang.org/pkg/html/template/).

## Config with `.semver/explicit-config.json`

Version {{ .SrcVersion }} example:

```json
{
  "versioning": ["src", "README.md", "TODO.md"],
  "templates": []
}
```

## Artifacts

A `.semver` directory that maintains various files pertaining to the state of the application. This directory is meant to be versioned (primary usecases of today call for Git) for consistency between developers.

## Resources

- (https://semver.org/)[https://semver.org/]
