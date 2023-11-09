# Work 3

Like the task 2, but with rational fractions and,
most importantly, **GUI**.

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
