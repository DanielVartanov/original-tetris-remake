- Outline the end goal, how would it look o satisfy me?
  https://www.firstversions.com/2015/11/tetris.html

[V] Print the playing field 10x20
```
  <! . . . . . . . . . .!>
  <! . . . . . . . . . .!>
  <!====================!>
    \/\/\/\/\/\/\/\/\/\/
```

[V] Make it run in a non-canonical term way, and actually draw/render the playing field in a loop

[V] Listen to SIGINT in the loop, exit on SIGINT and SIGTERM

[.] Listen to keyboard, exit on C-c, q

- Visible objects:
  - Occupied & free cells
  [V] Playing field boundaries
  - Figures
- Rendering all of the visible objects ^^^
  - Draw them just being static, simply place then somehow and make it
    render, that's it!

- Movement logic:
  - Simple: down upon timer events
  - Simple: left and right upon user's input
  - Rotation
    - Perhaps each figure shall have its own rules of rotation:
      compare a T to a square, for instance

- Final similarities to the original
  - Read screen size on init, draw from the middle
  - Next figure
  - Score calcs
  - Levels & speeding up
  - Hotkeys / help

- Correctness
  - Get notified on screen size change (zoom and/or re-size) and re-draw
  - Clean-up properly at SIGINT (C-c)
  - Ensure it runs well on Mac and Win
  - Do a proper peer-review to adhere to modern industry standards


- Promotion and finish
  - Cool README
  - Add yourself here https://github.com/gdamore/tcell#examples
  - Add it to resume, maybe?
