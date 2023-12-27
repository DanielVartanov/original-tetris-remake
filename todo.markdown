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

[V] Glue everything back
  [V] Make graphics dark green
  [V] Make `main()` add a _random_ piece in the beginning
  [V] Make it centered
  [V] Make the piece fall (not drop) by timer
  [V] Make the game react to `left` and `right` buttons

[V] Test over the graphics
  [V] Make the graphics output a string of a size of the playing field
  [V] Simply compare the string full of glyphs to the expected one,
      like this:
      ```
        <! . . . . . . . . . .!>
        <! . . . . . . . . . .!>
        <!====================!>
          \/\/\/\/\/\/\/\/\/\/
      ```

[V] Core game functions and tests for them
  [V] All types of pieces
  [V] Movement logic:
    [V] Sideways movement
    [V] Simple: down upon timer events
    [V] Collision detection
      [V] When moving sideways
      [V] When falling (against the bottom)
    [V] Encapsulate all `== '■'`
    [V] Rotation
      [V] Ability to rotate upon `w` keypress
      [V] Boundaries collision detection
    [V] Piece drop movement upon `s` keypress
    [V] Filled cells

[.]- Make it a game
  [V] Introduce Tetris
     [V] It will have count ticks, 7 ticks per fall initiatlly
     [V] It will add pieces, `main.go` should not call `AddPiece` and
       should now know `Well` at all, it will have only `Tetris`
  [V] Bake-in piece when it cannot fall further
  [V] Add a new piece when the previous one is baked-in
  [V] Snap a line if filled
  [.] Test collision detection with filled cells
    - When moving sideways
    - When falling
    - When dropping
    - When rotating
  - Keep scores for the snapped lines
    - Extract score keeping to a separate class not to have to test `Tetris`
  - Introduce speed and increase it at score threshold

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
  - Next piece
    - Introduce `Graphics` that will render field, next piece, scores
      etc. Basically, something like `Field` that aggregates Field and
      other pieces of screen. It will take a `Tetris` (not just a
      `Well`) as a constructor argument.
  - Score calcs
  - Levels & speeding up
  - Make it react to arrow keys (requires stdin buffer)
    ```go
    if (first := reader.Read()) == '\x1b' {
        second, third := reader.Read(), reader.Read()
        return {first, second, third}
    } else {
        return first
    }
    ```
  - Hotkeys / help

- Correctness
  - Get notified on screen size change (zoom and/or re-size) and re-draw
  - Clean-up properly at SIGINT (C-c)
  - Ensure it runs well on Mac and Win
  - Check if usage of "reflect" package is okay, change otherwise
  - Refactor
    - Well
      - Convert all methods to references from values `(w Well)`
    - Game
      - Leave only system/terminal issues in `tetris.go`, everything
        related to game shall be in `game.go` or whatever: it will
        create `Tetris`, set up timer, initiate graphics etc (maybe
        only `Screen` is to be created at `main()` level as it is
        related to terminal)
    - Graphics
      - `glyph` should be `[2]rune`, not a `string`
      - field.Render()
		// nextPiece.Render()
		// other-ui-elements.Render()
		// ^^^ all this to be extracted to Graphics as well
        Graphics render everything to the same instance of `Screen`,
        and then that screen gets printed to terminal
        field, NextPiece and others know about tetris, while
		`viewport` and `Screen` don't
        And yes, it might be okay of `main()` knows about the screen
		too, creates a screen itself (since it's related to
		terminal!), passes it to Graphics, let graphics render things,
		and then takes the instance of `Screen` directly to print it
		to terminal.
        - Идея! Функция создания viewport'а находится у Screen'а, но
          вызывают её элементы интерфейса (потому что только они знают
          размеры!), далее хранят этот вьюпорт у себя, и рисуют уже
          только на нём, а screen не хранят вообще, им же нужен
          толкько viewport

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
