# Type Flow

## State

- Die
  - value int
  - GetValue() int
- Slot
  - die Die
  - FillWithDie(Die)
  - GetDie() Die
- Column
  - slots []Slot
  - AddDie(Die)
  - RemoveDie(Die)
  - EmptySpaces() int
  - GetColumn() []Slot
- PlayerBoard
  - columns []Column
  - IsFull() bool
  - AddDieToCol(Die, Column)
  - GetBoard() []Column
- GameBoard
  - boards []PlayerBoard
  - ActivePlayer int
  - NextTurn()
  - AddDieToCol(Die, Column)

## Actions

- Roll
  - Roll() Die
- ColumnPlacement
  - columnNumber int
  - GetColumnPlacement(Agent)

## Agents

- User
  - AskForPlacement() int
- AI
  - AskForPlacement() int

# Program Structure

```
create game board
while not IsFull()
    switch active player
    AskForPlacement()
    create ColumnPlacement
    AddDieToCol
```