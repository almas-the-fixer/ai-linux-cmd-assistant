# Linux File Searching and Text Searching

This document covers common Linux commands used to search for files, directories, and text.

## grep

`grep` searches input files for lines matching a specified pattern.

### Basic syntax

```bash
grep [OPTION]... PATTERN [FILE]...
```

### Basic examples

Search for the word `error` in a file:

```bash
grep "error" logfile.txt
```

Search multiple files:

```bash
grep "error" app.log system.log
```

Search recursively through a directory:

```bash
grep -r "error" /var/log/
```

### Useful options

`-i` — ignore case when matching.

```bash
grep -i "error" logfile.txt
```

`-n` — show line numbers for matching lines.

```bash
grep -n "error" logfile.txt
```

`-r` or `--recursive` — search recursively through directories.

```bash
grep -r "TODO" .
```

`-v` — invert the match and show lines that do not match.

```bash
grep -v "debug" logfile.txt
```

`-c` — print the number of matching lines instead of the matching lines.

```bash
grep -c "error" logfile.txt
```

`-l` — print only the names of files containing a match.

```bash
grep -l "error" *.log
```

`-w` — match a whole word.

```bash
grep -w "root" users.txt
```

`-E` — use extended regular expressions.

```bash
grep -E "error|warning" logfile.txt
```

`-F` — treat the pattern as a fixed string rather than a regular expression.

```bash
grep -F "192.168.1.10" logfile.txt
```

### Combining options

Search recursively, ignore case, and show line numbers:

```bash
grep -rin "error" /var/log/
```

### Searching command output

`grep` can read from standard input, making it useful in pipelines.

```bash
ps aux | grep nginx
```

This searches the output of `ps aux` for lines containing `nginx`.

### Important distinction

Shell wildcard expansion and `grep` patterns are not the same thing.

For example:

```bash
*.log
```

is a shell glob, while:

```bash
.*\.log$
```

is a regular expression pattern.

---

## find

`find` searches for files and directories by traversing a directory hierarchy.

### Basic syntax

```bash
find [starting-point...] [expression]
```

### Find files by name

```bash
find . -name "file.txt"
```

This searches from the current directory for files or directories named `file.txt`.

Case-insensitive name matching:

```bash
find . -iname "file.txt"
```

### Find files by type

Find regular files:

```bash
find . -type f
```

Find directories:

```bash
find . -type d
```

### Find files by extension

```bash
find . -type f -name "*.log"
```

This searches recursively for regular files ending in `.log`.

### Find files by size

Find files larger than 100 MB:

```bash
find . -type f -size +100M
```

Find files smaller than 10 MB:

```bash
find . -type f -size -10M
```

### Find recently modified files

Find files modified within the last day:

```bash
find . -type f -mtime -1
```

### Execute a command on results

`find` can perform actions on matching files.

For example:

```bash
find . -type f -name "*.tmp" -delete
```

This deletes matching files.

Use destructive actions carefully.

### Combining conditions

Find regular files whose names end in `.log`:

```bash
find . -type f -name "*.log"
```

Find files larger than 100 MB:

```bash
find /var -type f -size +100M
```

### find and grep together

`find` can locate files while `grep` searches their contents.

```bash
find . -type f -name "*.log" -exec grep -n "error" {} \;
```

This finds `.log` files and searches each one for `error`.

---

## which

`which` is commonly used to locate the executable that would be invoked for a command.

Example:

```bash
which python
```

Possible output:

```text
/usr/bin/python
```

It can be useful for checking which executable is being found through the user's `PATH`.

---

## whereis

`whereis` locates the binary, source, and manual page associated with a command when those locations are known.

Example:

```bash
whereis bash
```

Possible output may include paths for:

* the executable
* source files
* manual pages

---

## Common command selection

Use `grep` when:

```text
You want to search the CONTENT of files.
```

Example:

```bash
grep "error" logfile.txt
```

Use `find` when:

```text
You want to search for FILES or DIRECTORIES.
```

Example:

```bash
find . -name "*.log"
```

Use `which` when:

```text
You want to know which executable will be found for a command.
```

Example:

```bash
which python
```

Use `whereis` when:

```text
You want to locate the binary, source, or manual page associated with a command.
```

Example:

```bash
whereis bash
```

## Common troubleshooting

### `grep` returns nothing

Possible reasons include:

* The pattern does not occur in the input.
* The search is case-sensitive and the capitalization differs.
* The wrong file was specified.
* The pattern is being interpreted as a regular expression.

Try:

```bash
grep -i "pattern" file.txt
```

or use `-F` when searching for a literal string:

```bash
grep -F "literal.text" file.txt
```

### `find` returns nothing

Check:

* The starting directory is correct.
* The filename pattern is correct.
* The `-type` condition matches the object you are searching for.
* Case sensitivity is appropriate.

For example:

```bash
find . -iname "*.log"
```

can be useful when filename capitalization is unknown.

## Quick reference

| Goal                          | Command                       |
| ----------------------------- | ----------------------------- |
| Search text inside a file     | `grep "pattern" file`         |
| Search text recursively       | `grep -r "pattern" directory` |
| Ignore case                   | `grep -i "pattern" file`      |
| Show matching line numbers    | `grep -n "pattern" file`      |
| Find a file by name           | `find . -name "file"`         |
| Find regular files            | `find . -type f`              |
| Find directories              | `find . -type d`              |
| Find by extension             | `find . -name "*.log"`        |
| Find large files              | `find . -size +100M`          |
| Locate an executable          | `which command`               |
| Locate binary/source/man page | `whereis command`             |
