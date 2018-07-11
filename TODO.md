## Technical Artifact

### Poss. Dependencies

- https://github.com/olekukonko/tablewriter for nice ASCII tables

### Configuration/State Approach

When a user inits, a `.gosemver` directory will be produced. The state of the application is maintained in this directory. Each time the application is run, it will utilize go routines to check for Patch releases as well as updating all templates based on the output of the tracking file. A log file will sit amongst the tracking file as a way to log behaviours.

The `.gosemver` directory:

- `config.json`
- `ledger.json`
- `events.log`

### CLI Design

```sh
gosemver init <list of files directories for quick add to versioning> # i.e. `gosemver init src tests xyz.wireframe`
gosemver version # get application version
gosemver versioning list # get list of artifact-ids and current versions (in tabular form)
gosemver versioning add <file|directory> # append new artifact to versioning list
gosemver versioning add <file|directory> # append new artifact to versioning list
gosemver versioning remove <id> # remove an artifact by ID (will not be removed from history or ledger)
gosemver templates list # get list of templates
gosemver template add <file> # append a new template
gosemver template remove <id> # remove a template by ID
gosemver history
gosemver history <artifact-id> # artifact-id is an item in config versioning / consider pulling commit information per minor releases or an artifact?
gosemver increment <major|minor> <artifact-id>
gosemver increment <major|minor> <artifact-id> --suffix rc # for custom suffix
gosemver increment <major|minor> <artifact-id> --rm-suffix # to remove suffix with addition
gosemver checkout <artifact-id> <version> # checks out git tagged with the following version (read-only *no incrementing from here* - this is not Git)
```

### Template Variables

Templates will have access to the entire environment of Template variables found by using `gosemver versioning list`.

If artifact source is "src":

| ID     | Name  | Template Variable | Version |
| ------ | ----- | ----------------- | ------- |
| abc123 | `src` | {{ .SrcVersion }} | 0.0.1   |


If artifact source is "README.md":

| ID     | Name        | Template Variable      | Version |
| ------ | ----------- | ---------------------- | ------- |
| xyz456 | `README.md` | {{ .ReadmeMdVersion }} | 0.0.1   |


If artifact source is "1234.md":

| ID          | Name      | Template Variable | Version |
| ----------- | --------- | ----------------- | ------- |
| whythisname | `1234.md` | {{ .MdVersion }}  | 0.0.1   |

*The above tables are also ideas for the output of the `versioning list` command*

### Git Tagging

If artifact source is "src": `src-0.0.1`

If artifact source is "README.md": `README.md-0.0.1`

Git tags point to Git commit and therefore once Go-Explicit-Semver has set a tag, you are able to return to it. A common usecase may be if a bug is reported in a specific version, this tool abstracts a bit of the Git kung-fu.