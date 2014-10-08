jot, a drop dead simple CLI note taker
--------------------------------------

`jot` is a tiny note taking utility written in Go. The goal here was to make
the act of taking quick plaintext notes, or jots, absolutely frictionless.
Likewise, it aims to provide a central set of tools for searching through and
keeping track of existing jots.

### Basic usage
The basic usage is as simple as running `jot` with no arguments. This opens your
favorite text editor (right now, just vim) to a new jot, which will be saved in
`~/.jot` with a SHA based uuid filename.

The idea here is that the name of your jot isn't as important as its content,
since jot provides some nice (though simple) methods of perusing that content.

That said, if you'd like your jot to be named `something`, just run 
`jot something` instead.

### Viewing, searching, and listing jots

* To open an existing jot, run `jot <jotname>` where jotname is the SHA uuid of the
jot or its custom name
* To get list of all jots and their "title" (the first line of each jot), run `jot ls`
* To search through your jots for some pattern, run `jot grep <pattern>`

### To do:
* implement `jot lock` and `jot unlock`, which encrypts/decrypts all jots
