# aocplz

Plz! Give me some [Advent of Code](https://adventofcode.com/), and make it quick!

This program is intended to be a lickity split way to get the day's Advent of Code challenge and be on your way to a fast solution.

## How To Use

```
➜  ~ aocplz fetch --day=1
Fetching input and creating test file for AOC 2022 day 1
Created new directory: [...]/advent-of-code-2022/day-1      # creates a directory for the new day
Created new input.txt file with data                        # pulls the input data from the puzzle and puts into a file
Created new test files: [a.rb Gemfile]                      # creates template solution files for the day
Opening browser to relevant puzzle...                       # opens a browser window to the puzzle
 ✓ Fetch Complete, Happy Coding!
```

If you're impatient (like me) you can also run
```
➜  ~ aocplz fetch --day=25 --wait
𝘹 the puzzle for day 25 won't be ready until 526h31m47s from now
526h30m48s left to wait...
526h30m43s left to wait...
526h30m38s left to wait...
```

... which will wait until the puzzle is available, and then run above items once it is. This is for those of you waiting anxiously each night for the next day's puzzle.

You can also provide a `--year` flag if you're working on a previous year's. Otherwise it will default to this year's. If `--day` is not provided, it will default to today as well, but that can have unintended consequences since it uses your current local time.

Right now the only subcommand available is `fetch`. I've stubbed out a `solve` subcommand but am not sure exactly yet what that would look like and if it really would save any time.

## How It Works

This program is fairly opinionated about how you do Advent of Code. It assumes:
* That you have a parent `/advent-of-code` directory with children `/day-1`, `/day-2` etc directories
* That you want to use the same programming language for each day
* That you want to start with a fresh script each day
* The way it ships, it assumes you'll be coding in Ruby, but that's also configurable

### Environment Variables

This program requires a handful of environment variables in order to run. I put these in my `~/.zshrc` file so that I can run the program from any terminal in any directory.

```
AOCPLZ_SESSION_TOKEN=[can be stolen from your cookies in the browser]   # required
AOCPLZ_ROOT_DIR=[path to your parent aoc dir]/advent-of-code-2022       # required
AOCPLZ_TEMPLATES_DIR=[path to aocplz repo or your own]/templates
AOCPLZ_TEMPLATE_FILES=a.rb.tmpl,Gemfile.tmpl
```

The program will look in the `AOCPLZ_TEMPLATES_DIR` directory for comma-separated files directed by `AOCPLZ_TEMPLATE_FILES`, it will then copy these files into the day's folder of your `AOCPLZ_ROOT_DIR` directory.

The `AOCPLZ_SESSION_TOKEN` is used to fetch your personal input data for the puzzle.

If no templates directory or files are provided, the program will use this repo's latest files from Github.

### Customization

As shipped, the `a.rb` file is setup to use some helper methods from the `activesupport` gem and also comes ready to read in the input data from the day's puzzle.

But what if you want to code in Golang? Python? Typescript?

Easy. All you need to do is create a `templates` directory somewhere, and put all the files that you would want copied for each day into that directory. Make sure that they end with `.tmpl` (which will be stripped when they are copied into the day's puzzle folder).

Set your `templates` directory's full path to the `AOCPLZ_TEMPLATES_DIR` environment variable, and comma-separate the files in that folder in the `AOCPLZ_TEMPLATE_FILES` variable.

You can then put whatever you want into these files and they will get copied into each day!

(Note: a `TODO` here would be to programatically iterate over all files in the template folder removing the need for the `AOCPLZ_TEMPLATE_FILES` variable)

## Installation

You need Golang. Download it [here](https://go.dev/dl/) then make sure the `go version` runs succesfully.

Then, run:
```
go install github.com/joshmenden/aocplz@v0.4.2
```

And you are ready to code!
