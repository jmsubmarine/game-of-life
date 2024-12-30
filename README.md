# game-of-life
This project is an implementation of Conway's Game of Life in Go. It simulates a cellular automaton on a toroidal grid, where each cell is either "alive" (O) or "dead" ( ). The state of the universe evolves based on predefined rules. 

**Rules of the Game:**
Any live cell with 2 or 3 live neighbors survives.
Any dead cell with exactly 3 live neighbors becomes alive.
All other live cells die in the next generation.
