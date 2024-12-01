# Advent of Code 2024
This year's advent of code is being solved in Go as a series of webservices.
Each webservice exposes `POST /1` and `POST /2` which accepts an `application/text`
with the puzzle input as data.

The services can be easily run with `docker-compose up` from the root directory.
The URL for a given solution is `http://localhost:${3000+day}/${part}` For example
if you have your puzzle input in a file called `input` and want to get the solution for
day 1 part 2, you would run `curl -X POST --data-binary @input http://localhost:3001/2`.

## Progress
* ⭐ means solved
* 🥸 means solved, but not a very satisfying solution
* 💩 means solved, but...

| Day | Solution | Day | Solution |
|-----|----------|-----|----------|
| 01  | ⭐        | 14  |          |
| 02  |          | 15  |          |
| 03  |          | 16  |          |
| 04  |          | 17  |          |
| 05  |          | 18  |          |
| 06  |          | 19  |          |
| 07  |          | 20  |          |
| 08  |          | 21  |          |
| 09  |          | 22  |          |
| 10  |          | 23  |          |
| 11  |          | 24  |          |
| 12  |          | 25  |          |
| 13  |          |     |          |
