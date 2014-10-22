jot: drop dead simple CLI note taking
--------------------------------------

`jot` is a tiny note taking utility written in Go. The goal here was to make
the act of taking quick plaintext notes, or jots, absolutely frictionless.
Likewise, it aims to provide a central set of tools for searching through and
keeping track of existing jots.

`jot` was built on the principal that quick notes generally don't need to be 
named (much like sticky notes), that they should all be plaintext and in a central
location, and they should be easily viewable, editable and searchable via the 
command line.

All jots are saved in `$HOME/.jot`, and are opened in your favorite editor (which
right now is `vim`; this will become configurable in the future)

### Basic usage

* `jot`: opens a new jot with a unique SHA filename in your jot directory
* `jot <jotname>`: opens a new (or existing) jot named `jotname`
* `jot ls`: lists all jots, along with their title (e.g. the first line of the jot)
* `jot grep <pattern>`: searches through all jots for the given plaintext pattern

### To do:
* implement `jot lock` and `jot unlock`, which encrypts/decrypts all jots
* support other people's favorite text editors
* bash/zsh autocompletion for jot names
* `jot rm`
