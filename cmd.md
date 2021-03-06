# cmd

credits to: <https://stackoverflow.com/a/28891643/6023007>

## Multiple Commands

- `&` separates commands on a line.
- `&&` executes this command only if previous command's errorlevel is 0.
- `||` (not used above) executes this command only if previous command's errorlevel is NOT 0

## Redirection

- `>` output to a file
- `>>` append output to a file
- `<` input from a file
- `|` output of one command into the input of another command

## Literals

- `^` escapes any of the above, including itself, if needed to be passed to a program
- `"` parameters with spaces must be enclosed in quotes

## Copy

- `+` used with copy to concatenate files. E.G. copy file1+file2 newfile
- `,` used with copy to indicate missing parameters. This updates the files modified date. E.G. `copy /b file1,,`

## Variables

- `%variablename%` a inbuilt or user set environmental variable
- `!variablename!` a user set environmental variable expanded at execution time, turned with SetLocal EnableDelayedExpansion command
- `%<number>` (`%1`) the nth command line parameter passed to a batch file. %0 is the batchfile's name.
- `%*` (`%*`) the entire command line.
- `%<a letter>` or `%%<a letter>` (`%A` or `%%A`) the variable in a for loop. Single `%` sign at command prompt and double `%` sign in a batch file.
