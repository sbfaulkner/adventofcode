# Advent of Code Solutions

http://adventofcode.com/

## Day 1: Not Quite Lisp

A bit awkward, but it works.

## Day 2: I Was Told There Would Be No Math

Better than Day 1, but nothing special.

## Day 3: Perfectly Spherical Houses in a Vacuum

`Route#deliver` could use some refactoring, but generally a decent solution.

## Day 4: The Ideal Stocking Stuffer

Quite happy with how succinct the approach used here is, although performance might be a bit lacking.

## Day 5: Doesn't He Have Intern-Elves For This?

The two independent versions muddy this solution somewhat, but I think this is reasonable.

## Day 6: Probably a Fire Hazard

Decent solution, class naming could be better.

## Day 7: Some Assembly Required

A slightly awkward solution, but a proper lexer/parser seems like overkill.
Also, calculating the signal effectively destroys the circuit (as it caches values).

## Day 8: Matchsticks

The lexer should be decoupled from the decode/encode action... but it works.
Manually implementing the FSM was kind of fun. I should probably apply that technique more often.

## Day 9: All in a Single Night

Not the optimal solution (with larger datasets), and definitely considers twices as many routes as necessary.
However, kind of a nice solution taking advantage of Rubyisms.

## Day 10: Elves Look, Elves Say

Neat problem. I wonder what other approaches would yield for solutions.

## Day 11: Corporate Policy

I think I'd prefer to separate the password validation from the incrementing/finding for better separation of concerns.

## Day 12: JSAbacusFramework.io

Decent solution. There may be some way to combine the Array and Hash value extraction.
However, I suspect the requirement to ignore/skip properties prevents this generalization.

## Day 13: Knights of the Dinner Table

Used same approach as day 9... however, we know that we consider (at least) twice as many routes as necessary on day 9
(ie. the distance is the same forwards an backwards). Here, we know it's even less optimal, since the permutations can
now be both a) flipped (like day 9) and b) rotated, since the table is circular.

## Day 14: Reindeer Olympics

Used a mathematical approach to calculate position based on elapsed time. It would be interested to attempt this again
by tracking state and advancing reindeer based on velocity during simulation.

## Day 15: Science for Hungry People

This solution could do with some refactoring... and, I'm not sure I'm happy with the complexity of Recipe#combine.
