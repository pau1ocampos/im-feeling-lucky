# im-feeling-lucky
I'm feeling lucky or just `lucky` is a CLI written in Go that generates Euromillions keys. If you win using this generator, please share a bit of the prize with me. :smile:

This is just a small project that I worked on to practice my Go skills.

The key generation is based purely in computer randomization. It's possible match the generated key with the past draw results (very unlikely event) - in case of a repeated key a new one is generated.

## Third party data

`lucky` CLI can work in a standalone manner by passing `disable-check` check flag to `draw` command.   
All the other commands are leveraged by using data from [euro-millions](https://www.euro-millions.com/) website.   
Since there's no API available the data is scraped from source page HTML.   
The only data scraped are the number and the stars from each past draw.   
The data scraped is only used to check if the generated key is _repeated_, meaning that if the generated key was already drawn in the past.   
It's possible to store the parsed data into a `json` file, intention is this that that data to be used to feed the `draw` command and nothing else.

## How it works

### Generate numbers and starts

```shell
lucky draw --disable-check
```

### Generate numbers and stars using the match check

This command will scrape data from past draws and compare the generated key agains them. If generated key is a past draw, the CLI will generate a new one

```shell
lucky draw
```

### Generate numbners and stars using the match check providing the draws from a file

Same as [previous section](#generate-numbers-and-stars-using-the-match-check) but instead of scrapping directly from the web, it will use a `json` file.

```shell
lucky draw --file-path /path/to/json/file
```

Spec of json file:

```json
[
    {
        "numbers": [<int>],
        "stars": [<int>]
    }
]
```

### Scrape data from the web

```shell
lucky scrape --year 2004 --file-path /path/to/store/data --silent
```

* `year`: collects data from the provided year until the current year. There are no data previous to 2004, so if no year is passed or is passed a year prior to 2004, by default 2004 will be set for both scenarios
* `file-path`: the path to a file to store the data collected. Will be stored in `json` format. If no `file-path` is passed, the scrape will performed, however just printed to the screen in a non-standard format (just a print from the Go structure used to work with the data). If `silent` is flag is passed without `file-path` the scrape will be done but nothing shown, so it's kinda useless.
* `silent`: won't print the non-standard data (a print from the Go stucture used to work with data)

### Install

Go to the repository latest [release](https://github.com/pau1ocampos/im-feeling-lucky/releases/latest) and download the version which fits your computer architecture.

On alternative you can clone this repository and just run `make build` - you need to have makefile installed on your system.