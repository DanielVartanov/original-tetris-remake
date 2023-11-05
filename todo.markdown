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

[.] Extract the model
  [V] Convert global var based logic into struct and its methods
  [.] Convert pieces into a map instead of a set of vars?

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
        - (!) Organise viewports as slices over the array of `screen`,
          those viewports, like a viewport over the playing field or a
          viewport over the "next figure box" are super convenient,
          inside those the coordinates always start with (0,0), it's
          super easy to draw within them, and they are independent of
          the terminal size, but the overall drawings will always be
          in the middle of the terminal window or wherever they need
          to be
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
  - Build a layer of `game.go` that surrounds the core mechanics
    - `Game` struct will control the randomness, the speed, it will
      contain the Timer, it will calculate the scores, be responsible
      for the "next figure" generation and so on.
      - `main()` will call `NewGame()` instead of `NewTetris()`,
        perhaps `Tetris` will be hidden behind `Game`, the `Game` will
        be facade for everything, even for `Graphics`.
      - Make sure to find how to leave a docstring for entire file (or
        maybe the main struct of the file?) and write that `Game` is
        what surrounds the core mechanics and makes the game
        competitive: speed, scores, victory and defeat conditions
        - Also add a docstring to `core.go`, simply say it defines
          core mechanics of Tetris. I guess, `piece.go` is
          self-explanatory and does not require a comment
  - Next figure
  - Score calcs
  - Levels & speeding up
  - Hotkeys / help

- Correctness
  - Get notified on screen size change (zoom and/or re-size) and re-draw
  - Clean-up properly at SIGINT (C-c)
  - Ensure it runs well on Mac and Win
  - Do a proper peer-review to adhere to modern industry standards

- Consider converting `core.go` and `graphics.go` into packages of
  their own. But ensure to write and list all the pros of _why_ am I
  doing it?

- Promotion and finish
  - Cool README
  - Add yourself here https://github.com/gdamore/tcell#examples
  - Add it to resume, maybe?
