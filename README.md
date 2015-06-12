# sendgrid-reporter

A comman-line tool to get sendgrid statistics using sendgrid web api.

## Installation

```
go get github.com/hiroakis/sendgrid-reporter
```

## Usage

* prepare

Set ENV of sendgrid user and key

```
export SENDGRID_API_USER=XXXXX@kke.com
export SENDGRID_API_KEY=XXXXXXXXXXXX
```

### Basic

* show general statistics

```
sendgrid-reporter stat
```

* show bounced mail

```
sendgrid-reporter bounce
```

* show blocked mail

```
sendgrid-reporter block
```

* show invalid mail

```
sendgrid-reporter invalid
```

* show spam report

```
sendgrid-reporter spam
```

### Using options

* -s: "The start of the date range for which to retrieve invalid emails. Date must be in YYYY-MM-DD format. Default: Today's date"
* -e: "The end of the date range for which to retrieve invalid emails. Date must be in YYYY-MM-DD format. Default: Today's date"
* -d: "Number of days in the past for which to retrieve invalid emails (includes today)."
* -l: "Optional field to limit the number of results returned."
* -o: "Optional beginning point in the list to retrieve from."
* -t: "Hard or Soft. Choose the type of bounce to search for."
* -email: "Optional email to search."

## License

MIT
