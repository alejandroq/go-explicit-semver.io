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
gosemver history <artifact-id> # artifact-id is an item in config versioning
gosemver increment <major|minor> <artifact-id>
gosemver increment <major|minor> <artifact-id> --suffix rc # for custom suffix
gosemver increment <major|minor> <artifact-id> --rm-suffix # to remove suffix with addition
gosemver checkout <artifact-id> <version> # checks out git tagged with the following version (read-only *no incrementing from here* - this is not Git)
```

### Template Variables

Templates will have access to the entire environment of Template variables found by using `gosemver versioning list`.

If artifact source is "src":

| ID     | Name  | Template Variable | Version | Git Tag |
| ------ | ----- | ----------------- | ------- | ------- |
| abc123 | `src` | {{ .SrcVersion }} | 0.0.1   | true    |


If artifact source is "README.md":

| ID     | Name        | Template Variable      | Version | Git Tag |
| ------ | ----------- | ---------------------- | ------- | ------- |
| xyz456 | `README.md` | {{ .ReadmeMdVersion }} | 0.0.1   | false   |


If artifact source is "1234.md":

| ID          | Name      | Template Variable | Version | Git Tag |
| ----------- | --------- | ----------------- | ------- | ------- |
| whythisname | `1234.md` | {{ .MdVersion }}  | 0.0.1   | false   |

*The above tables are also ideas for the output of the `versioning list` command*