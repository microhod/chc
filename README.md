# chc - Character Count

`chc` returns a sorted list of the most common characters in a file or directory. If given a directory, it walks all subdirectories recursively.

## Example

Running on this repository, reading only `go` files and counting only special characters, gives:

```console
$ chc -i='[^a-zA-Z0-9\s]' -f='.*\.go' .
" 76
) 70
( 70
. 62
, 54
} 40
{ 40
: 38
= 35
] 18
\ 18
[ 18
* 12
! 12
% 10
' 6
& 5
_ 3
; 3
/ 3
> 2
+ 2
< 1
```

## Usage

```
Usage of chc:
  -f string
        filter files to read (regex) (default ".*")
  -i string
        characters to include (regex) (default ".*")
  -v    enable verbose output
```
