# Setup
* Go Installation 
    * https://golang.org/dl/
    * Code Editors:
        * VSCode https://code.visualstudio.com/
        * GoLand https://www.jetbrains.com/go/


## Lesson 1
* What is Go
* Brief History
* Computer setup
* Writing hello world
* compiling/running
* Vars, loops, etc

## Lesson 2
* methods
* imports
* intro to goroutines
* put everything together
* Structs

## Lesson 3 
* Show finished product
* JS code
* Starting Code
* intro to http library

## Lesson 4
* create ghost
* move a ghost
* create pacman
* move pacman

## Lesson 5
* pacman/ghost collision
* scoring

## Lesson 6
* Separate Packages

# Pacman API Definition

## Start the Game

#### Request

```GET /api/startGame```

#### Response

```200 OK```

## Track Pacman movements

#### Request

```GET /api/trackPacman?x=<x coordinate>&y=<y coordinate>```

#### Response

```200 OK```

## Add ghost to board

#### Request

```GET /api/addGhost?ghost=<id of ghost>```

#### Response

```200 OK```

## Randomly move ghosts

#### Request

```GET /api/moveGhost?ghost=<id of ghost>```

#### Response

```200 OK```

```
Possible Response Text : 
UP
DOWN
LEFT
RIGHT    
```

## Track ghosts

#### Request

```GET /api/trackGhost?ghost=<id of ghost>&x=<x coordinate>&y=<y coordinate>```

## Determine if pacman was caught

#### Request

```GET /api/isGameOver```

#### Response

```200 OK```

```
Possible Response Text : 
true
false    
```
