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

[V] Extract the model
  [V] Convert global var based logic into struct and its methods
  [V] Convert pieces into a map instead of a set of vars

[V] Testing framework
  [V] Write a test helper function that takes a string snapshot of
      `Tetris`'s current state
  [V] Write simplest test over the model
  [V] Make the test function print expected and got horizontally line
      by line
  [V] Amend testing framework with progression tests

[V] Implement steps/progressions, even without checking for the field boundaries

[.] Glue everything back
  [V] Make graphics dark green
  [V] Make `main()` add a _random_ piece in the beginning
  [V] Make it centered
  - Make the piece fall (not drop) by timer
  - Make the game react to `left` and `right` buttons

- Test over the graphics
  - Make the graphics output a string of a size of the playing field
  - Simply compare the string full of glyphs to the expected one,
    like this:
    ```
      <! . . . . . . . . . .!>
      <! . . . . . . . . . .!>
      <!====================!>
        \/\/\/\/\/\/\/\/\/\/
    ```

- Core game functions and tests for them
  [V] All types of pieces
  - Movement logic:
    - Sideways movement
    - Simple: down upon timer events
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
  - Refactor
    - `glyph` should be `[2]rune`, not a `string`
    - Extract testing framework from `core_test.go`, it feel too dirty
      when dumped in the same file
    - Try using beautiful Unicode characters like `𐘀`
      (string(1000000)) as key runes for pieces
  - Do a proper peer-review to adhere to modern industry standards

- Consider converting `core.go` and `graphics.go` into packages of
  their own. But ensure to write and list all the pros of _why_ am I
  doing it?

- Promotion and finish
  - Cool README
  - Add yourself here https://github.com/gdamore/tcell#examples
  - Add it to resume, maybe?
