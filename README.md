# gaudit - Github Audit Tool for Organizations

This is a command line tool to help you organize and analyze your github
repositories in your organization. You can create some basic rules to look
for conditions in your repositories and output the results as a CSV.


### Commands

```
gaudit

  gaudit update - Update working list of github repos
  gaudit list - List of working github repos
  gaudit diff <old> <new> - Difference since last update
  gaudit details [filter] - Detail list of github repos
  gaudit analyze - Analyze rules against working repos
  gaudit results [--verbose] - Show results of analysis
  gaudit append - Creates or updates the append file
  gaudit stats - Summarizes statistics data on the audit
  gaudit csv - Outputs to csv format
```


### Environment Variables

`GAUDIT_GITHUB_TOKEN` - The Github Token for your account.

`GAUDIT_ORGANIZATION` - Specify the Github Organization (does not work with personal accounts)

`GAUDIT_STORAGE` - The file to store state in (defaults to audit.store)

`GAUDIT_RULES` - Rules file for analysis (defaults to rules.yml)

`GAUDIT_APPEND` - Append file for annotating (defaults to append.yml)

`GAUDIT_DEBUG` - Set debug mode (defaults to false)


### Rules File

You can create a rules file to analyze your Github repositories.

Example:
```
-
  name: CircleCI Config
  action: exists
  resource: .circleci/config.yml
-
  name: TravisCI Config
  action: not_exists
  resource: .travis.yml
-
  name: README File
  action: exists
  resource: README.md
-
  name: License
  action: exists
  resource: LICENSE
-
  name: Copyright Notice
  action: exists
  resource: NOTICE.md
  type: public
-
  name: Copyright Date
  action: contains
  resource: NOTICE.md
  match: 2019
  type: public
```


### Releases

NEXT
- Added rules results to details
- Added rules results to stats
- Updated diff to exclude size/updated date
- Updated the stats date output

v0.1.1
- Added a Stats command

v0.1.0
- Initial release
