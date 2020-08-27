# Cgo header include reproduction

TL;DR: cgo seems to look in the current working directory for includes in preference to directories specified with -I, regardless of whether the include is specified with <> or "". I expect that the current directory is searched for includes only when they're included with "".

## Reproducing

There are four test cases:
    * `cangled` - C only - import a header file with <>
    * `cquoted` - C only - import a header file with ""
    * `goangled` - cgo - import a header file with <>
    * `goquoted` - cgo - import a header file with ""

Each test case directory has a file `cstuff/cdata.h`, which has a constant `char *the_text = "this is from the embedded directory";`. Each test case prints that constant out. There's also a `system/cstuff/cdata.h` which contains a _different_ version of that constant: `char *the_text = "this is from the system directory";`.

## `cangled`

```
$ cd cangled; gcc -I../system main.c; ./a.out
this is from the system directory
```

As expected - gcc does not search the current directory for includes specified with <>

## `cquoted`

```
$ cd cquoted; gcc -I../system main.c; ./a.out
this is from the embedded directory
```

As expected - gcc searches the current directory in preference to any -I flags for includes specified with ""

## `goquoted`

```
$ cd goquoted; go run main.go;
this is from the embedded directory
```

As expected - cgo resolves the include from the current directory because of the ""

## `goangled`

```
$ cd goangled; go run main.go;
this is from the embedded directory
```

This is _not_ what I expect; I expect that since the include is done with <>, that the current directory is not used to resolve it, and instead it should find the version in `system/cstuff/cdata.h` because of the -I flag.
