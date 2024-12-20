# aocplz

Plz! Give me some [Advent of Code](https://adventofcode.com/), and make it quick!

This program is intended to be a lickity split way to get the day's Advent of Code challenge and be on your way to a fast solution.

## How To Use

```
➜  ~ aocplz fetch --day=1
Fetching input and creating test file for AOC 2022 day 1
Created new directory: [...]/advent-of-code-2022/day-1      # creates a directory for the new day
[...]
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

- That you have a parent `/advent-of-code` directory with children for each year `/2023`, `/2024` etc directories
- That in the root of each year dir, you have some file ending in `.tmpl` which will be copied into each day's solution

### Environment Variables

This program requires two environment variables in order to run. I put these in my `~/.zshrc` file so that I can run the program from any terminal in any directory.

```
AOCPLZ_SESSION_TOKEN=[can be stolen from your cookies in the browser]   # required
AOCPLZ_ROOT_DIR=[path to your parent aoc dir]/advent-of-code-2022       # required
```

The `AOCPLZ_SESSION_TOKEN` is used to fetch your personal input data for the puzzle.

## Installation

You need Golang. Download it [here](https://go.dev/dl/) then make sure the `go version` runs succesfully.

Then, run:

```
go install github.com/joshmenden/aocplz
```

And you are ready to code!
