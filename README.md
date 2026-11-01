# Jevi

### A CLI Trivia Quiz Game

Jevi is a command-line quiz game that fetches questions from the Open Trivia Database (OpenTDB) API.

## Features

- Fetches questions from OpenTDB API
- Choose number of questions, category and difficulty
- Suppots multiple types of questions
- Shuffles answers randomly
- Tracks your score

## Installation

You can directly download the binaries.

[Download for Linux](https://github.com/shantanuparte/Quiz-game/releases/download/v1.1.1/Quiz-game)

## Build for yourself

Clone the repo ¯\\(ツ)/¯

```
git clone https://github.com/shantanuparte/Quiz-game.git
```

Build it:

```
go build -o jevi .
```

## Usage

You can also use `./jevi --help`.

**1. Amount:** Set the number of qusetions (1-50).

```
./jevi -amount 10
```

**2. Category:** Choose a category using its ID (9-32).

```
./jevi -category 18
```

To see all available categories:

```
./jevi -categories
```

**3. Difficulty:** Choose between easy, medium and hard.

```
./jevi -difficulty easy
```

**4. Question Type:** Choose between multiple choice and boolean.

```
./jevi -type multiple
```

You can also combine flags:

```
./jevi -amount 15 -category 18 -difficulty medium -type multiple
```

## License

I'm using MIT so feel free to use.




