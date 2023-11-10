# Work 3

Like the task 2, but with rational fractions and,
most importantly, **GUI**.

# Usage

Put your matrix into the input field, like so:
```
1/2 1/3
1/4 1/5
```

Then press `Read` button to read the matrix.

After that, any of the other buttons may be used.

# How to run

1. Via Nix (flakes enabled)
```ShellSession
$ nix develop
$ go run .
```

2. Whatever else
First, install required C libraries
(see https://github.com/diamondburned/gotk4-examples#installing-gtk)

Then, again,
```ShellSession
$ go run .
```

If some garbage collector warnings arise,
just ignore them and pass whatever environment varialbe it will ask for.
