# Unique Tic Tac Toe

## Context
_Given this assignment:_

"Write a program that plays every possible tic-tac-toe game, and then prints the number of possible valid, completed game states.

Valid meaning the game can be reached by following the rules of tic-tac-toe, alternating player by player.

Completed meaning the game is won by either player or tied.

Note:

1. Assume X always goes first

2. If two games end in the same state, they should only be counted once. Below is an example of one valid game state after O moves:

```
X O X      X O X

X _ _  ->  X O _

_ O _      _ O _


X O X      X O X

X O _  ->  X O _

_ _ _      _ O _
```
"

## Running this program
Assuming that you can Go installed locally, you can run this program easily by running the below
```go
oliver@Mac WorkWhileAssignment % go run main.go 
num of unique completed games: 958
```

You can also run this program via Docker by building the image and subsequently running the container.
```shell
oliver@Mac WorkWhileAssignment % docker build -t unique-tic-tac-toe .
...
oliver@Mac WorkWhileAssignment % docker run --rm unique-tic-tac-toe
num of unique completed games: 958
```

## Thoughts
This was interesting to work through and is seemingly straightforward once you give it a bit of thought. 
However, I'm left bothered by the fact that I have no idea if the solution is correct. :( I can't think of a mathematical proof or ultimate unit test to validate the answer given the problem description.
Worse still if you wanted to extend this logic to allow for different board dimensions, required number of connections, possible moves, etc. If I think of a clever proof, I'll update this repo to include it.

Could this implementation be optimized in some ways? Sure. I could get rid of available state and derive it on the fly based on the board state, I could more cleverly encode the board state, maybe flattening the board state into a 1D array could be beneficial, I could just directly check the 8 possible win states instead of iterating for each valid winning path, etc.
However, these were trade-offs that I felt comfortable making for my own sake to better visualize the problem and be able to better debug edge cases. I also wanted the logic to be extensible for different game rules. Maybe overkill, maybe inefficient, who knows! 

I hope this implementation suffices for a take home assignment. :)
In either case, I appreciated the brain tease implementing this!