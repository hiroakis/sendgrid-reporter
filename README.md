# sendgrid-reporter

A comman-line tool to get sendgrid statistics using [sendgrid web api](https://sendgrid.com/docs/API_Reference/index.html).

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

### Pretty print

Use [jq](http://stedolan.github.io/jq/). The following is a example.

```
sendgrid-reporter -s 2015-06-12 stat | jq .
[
  {
    "delivered": 3777,
    "unsubscribes": 0,
    "repeat_bounces": 49,
    "invalid_domain": 0,
    "invalid_email": 47,
    "bounces": 27,
    "repeat_unsubscribes": 0,
    "unique_clicks": 0,
    "blocked": 6,
    "spam_drop": 0,
    "opens": 0,
    "repeat_spamreports": 0,
    "replies": 0,
    "date": "2015-06-12",
    "requests": 3904,
    "spamreports": 0,
    "clicks": 0,
    "umd_drop": 0,
    "unique_opens": 0
  },
  {
    "delivered": 157,
    "unsubscribes": 0,
    "repeat_bounces": 0,
    "invalid_domain": 0,
    "invalid_email": 1,
    "bounces": 3,
    "repeat_unsubscribes": 0,
    "unique_clicks": 0,
    "blocked": 0,
    "spam_drop": 0,
    "opens": 0,
    "repeat_spamreports": 0,
    "replies": 0,
    "date": "2015-06-13",
    "requests": 161,
    "spamreports": 0,
    "clicks": 0,
    "umd_drop": 0,
    "unique_opens": 0
  }
]
```

## License

MIT
