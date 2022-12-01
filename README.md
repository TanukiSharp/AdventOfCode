# Overview

This repository contains my implementations for the [Advent of Code (AoC)](https://adventofcode.com) problems.

## Note about input data

It is required to be logged in to [AoC] in order to achieve them, therefore everyone gets different data and so no one can spoil the answers. The drawback is that it's a pain in the butt for a console application to login to [AoC], and so it is so much simpler to grab data from the browser and drop it in a file for the console application to consume it.

## [2022](https://adventofcode.com/2022)

For the 2022 edition, I decided to implement the puzzles' solution in C#, but I might also try to implement them in Go.

Create a class named `DayXX` where `XX` must be a two digits number, representing the day of the month.<br/>
The class name for first day must be `Day01`.

Drop the input data for a given day in a file named `XX.txt`, so data for the first day must be named `01.txt`.<br/>
You can drop the file in the same folder as the binary, or in any directory up the directory tree structure.<br/>
This allows to drop the input files where the `.csproj` / `.cs` file is located and run more or less from anywhere without caring too much about where the executable looks for the files.

[AoC]: https://adventofcode.com
