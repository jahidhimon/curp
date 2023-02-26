Output ./demo/curp.gif

Require echo

Set PlaybackSpeed 0.6
Set FontSize 32
Set Width 1200
Set Height 600
Set LoopOffset 50%

Type "make build" Sleep 500ms  Enter
Sleep 2s
Type "./bin/curp" Sleep 500ms Enter
Sleep 2s
Type "./bin/curp --from eur --to bdt" Sleep 500ms Enter
Sleep 2s
Type "./bin/curp --from myr --to bdt --amount 34" Sleep 500ms Enter

Sleep 4s
