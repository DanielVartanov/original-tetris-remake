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

[V] Listen to keyboard, exit on C-c and q

[V] Visible objects:
  [V] Playing field boundaries
  [V] One sample figure (not moving or anything, just getting displayed
      at given coordinates)

- Extract the model (to `tetris-core`, maybe? We'll figure out later)
  - Read about objects and "classes" in Go

- Testing framework
  - Write simplest test over the model
  ```
     // var actions = []func(){
     // 	func() { /*game.progress()*/ },
     // 	func() { /*game.progress()*/ },
     // 	func() { /*piece.turn(-1)*/ },
     // 	func() { /*game.progress()*/ },
     // }

     // var expectedStages = [][]string{
     // 	{"  xx  ",   "      ",   "      "},
     // 	{"  xx  ",   "  xx  ",   "      "},
     // 	{"      ",   "  xx  ",   "  xx  "},
     // 	{"      ",   "      ",   "  xx  "},
     // 	{"      ",   "      ",   "      "},
     // }

     // var expectedStages = [][][]string{
     // 	{{"  xx  "},   {"      "},   {"      "}},
     // 	{{"  xx  "},   {"  xx  "},   {"      "}},
     // 	{{"      "},   {"  xx  "},   {"  xx  "}},
     // 	{{"      "},   {"      "},   {"  xx  "}},
     // 	{{"      "},   {"      "},   {"      "}},
     // }
  ```

  - Test over the graphics
    - Make the graphics output a string of a size of the playing field
      - A var/class of `screen` that simply contains a 2D array of
        runes which then gets printed
    - Simply compare the string full of glyphs to the expected one,
      like this:
      ```
        <! . . . . . . . . . .!>
        <! . . . . . . . . . .!>
        <!====================!>
          \/\/\/\/\/\/\/\/\/\/
      ```

- Core game functions
  - All types of pieces
  - Movement logic:
    - Simple: down upon timer events
    - Simple: left and right upon user's input
    - Rotation
      - Perhaps each figure shall have its own rules of rotation:
        compare a T to a square, for instance

- Final similarities to the original
  - Read screen size on init, draw from the middle
    - Make the graphics output a string of a size of the terminal
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
