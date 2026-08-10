# Linux Text Processing

Linux provides many command-line utilities for inspecting, searching, transforming,
and processing text. These commands are commonly used together through pipes (`|`)
to build powerful command-line workflows.

---

## 1. cat

`cat` reads files and writes their contents to standard output.

### Basic Usage

```bash
cat file.txt
````

### Display Multiple Files

```bash
cat file1.txt file2.txt
```

### Combine Files

```bash
cat file1.txt file2.txt > combined.txt
```

### Append Files

```bash
cat file1.txt file2.txt >> combined.txt
```

### Example

```bash
cat /etc/hosts
```

### Important

`cat` is mainly useful for displaying or combining files.

For large files, use tools such as `less` instead of printing the entire file
to the terminal.

---

## 2. less

`less` allows you to read a file interactively without printing the entire file
to the terminal at once.

```bash
less file.txt
```

Useful keys:

* `Space` - next page
* `b` - previous page
* `↑` / `↓` - move line by line
* `/pattern` - search
* `n` - next search result
* `q` - quit

### Example

```bash
less /var/log/syslog
```

`less` is especially useful for large files and log files.

---

## 3. head

`head` displays the beginning of a file.

```bash
head file.txt
```

By default, it displays the first 10 lines.

### Display a Specific Number of Lines

```bash
head -n 20 file.txt
```

You may also use:

```bash
head -20 file.txt
```

### Example

```bash
head -n 5 /var/log/syslog
```

---

## 4. tail

`tail` displays the end of a file.

```bash
tail file.txt
```

By default, it displays the last 10 lines.

### Display a Specific Number of Lines

```bash
tail -n 20 file.txt
```

### Follow a File

```bash
tail -f application.log
```

The `-f` option continuously displays new lines added to the file.

This is extremely useful for monitoring logs.

### Example

```bash
tail -f /var/log/syslog
```

Press `Ctrl+C` to stop following the file.

---

## 5. grep

`grep` searches text for a pattern and prints matching lines.

### Basic Usage

```bash
grep "error" application.log
```

This prints lines containing `error`.

### Case-Insensitive Search

```bash
grep -i "error" application.log
```

### Search Recursively

```bash
grep -r "TODO" .
```

This searches files recursively from the current directory.

### Show Line Numbers

```bash
grep -n "error" application.log
```

### Invert the Match

```bash
grep -v "success" application.log
```

This prints lines that do not contain `success`.

### Count Matches

```bash
grep -c "error" application.log
```

### Match Whole Words

```bash
grep -w "error" application.log
```

### Show Only Matching Filenames

```bash
grep -l "error" *.log
```

### Combine Options

```bash
grep -rin "error" /var/log/
```

This means:

* `-r` - recursive
* `-i` - case-insensitive
* `-n` - show line numbers

---

## 6. sort

`sort` sorts lines of text.

```bash
sort names.txt
```

### Reverse Order

```bash
sort -r names.txt
```

### Numeric Sorting

```bash
sort -n numbers.txt
```

Without `-n`, numbers may be sorted lexicographically rather than numerically.

### Sort by a Field

Suppose a file contains:

```text
alice 30
bob 20
charlie 40
```

Sort by the second field:

```bash
sort -k2 -n users.txt
```

### Sort and Remove Duplicates

```bash
sort -u names.txt
```

---

## 7. uniq

`uniq` removes or detects repeated adjacent lines.

```bash
uniq file.txt
```

Important: `uniq` only detects duplicates that are next to each other.

For example:

```text
apple
apple
banana
banana
apple
```

Running:

```bash
uniq file.txt
```

does not remove the final `apple`, because it is not adjacent to the first
`apple`.

To remove all duplicate lines:

```bash
sort file.txt | uniq
```

or:

```bash
sort -u file.txt
```

### Count Occurrences

```bash
uniq -c file.txt
```

A common pattern:

```bash
sort access.log | uniq -c
```

---

## 8. wc

`wc` counts lines, words, and bytes.

```bash
wc file.txt
```

### Count Lines

```bash
wc -l file.txt
```

### Count Words

```bash
wc -w file.txt
```

### Count Bytes

```bash
wc -c file.txt
```

### Example

```bash
wc -l application.log
```

This is useful when you need to know how many lines a file contains.

---

## 9. cut

`cut` extracts portions of each line.

It is especially useful with structured text separated by delimiters.

Suppose:

```text
alice:developer:1001
bob:admin:1002
charlie:user:1003
```

Extract the first field:

```bash
cut -d ':' -f 1 users.txt
```

Output:

```text
alice
bob
charlie
```

### Extract the Second Field

```bash
cut -d ':' -f 2 users.txt
```

### Extract Multiple Fields

```bash
cut -d ':' -f 1,3 users.txt
```

### Extract Character Ranges

```bash
cut -c 1-5 file.txt
```

Options:

* `-d` - delimiter
* `-f` - field
* `-c` - character positions

---

## 10. tr

`tr` translates or deletes characters.

### Convert Lowercase to Uppercase

```bash
echo "hello world" | tr 'a-z' 'A-Z'
```

Output:

```text
HELLO WORLD
```

### Convert Uppercase to Lowercase

```bash
echo "HELLO" | tr 'A-Z' 'a-z'
```

### Delete Characters

```bash
echo "hello123" | tr -d '0-9'
```

Output:

```text
hello
```

### Replace Characters

```bash
echo "hello world" | tr ' ' '_'
```

Output:

```text
hello_world
```

---

## 11. sed

`sed` is a stream editor used to transform text.

One of its most common uses is replacing text.

### Replace the First Occurrence on Each Line

```bash
sed 's/old/new/' file.txt
```

### Replace All Occurrences on Each Line

```bash
sed 's/old/new/g' file.txt
```

The `g` means global replacement.

### Delete Lines Containing a Pattern

```bash
sed '/error/d' application.log
```

### Print Specific Lines

```bash
sed -n '1,10p' file.txt
```

This prints lines 1 through 10.

### Edit a File Directly

```bash
sed -i 's/old/new/g' file.txt
```

Be careful with `-i` because it modifies the file.

---

## 12. awk

`awk` is a powerful text-processing language commonly used for processing
structured, column-based text.

Suppose:

```text
alice 100
bob 200
charlie 300
```

### Print the First Column

```bash
awk '{print $1}' users.txt
```

Output:

```text
alice
bob
charlie
```

### Print the Second Column

```bash
awk '{print $2}' users.txt
```

### Print Multiple Columns

```bash
awk '{print $1, $2}' users.txt
```

### Calculate a Value

```bash
awk '{sum += $2} END {print sum}' users.txt
```

### Use a Delimiter

For:

```text
alice:developer:1001
bob:admin:1002
```

Use:

```bash
awk -F ':' '{print $1}' users.txt
```

`-F` specifies the field separator.

---

## 13. diff

`diff` compares two files and shows their differences.

```bash
diff file1.txt file2.txt
```

This is useful for comparing configuration files, source code, or other text.

### Example

```bash
diff old.conf new.conf
```

A unified format is often easier to read:

```bash
diff -u old.conf new.conf
```

---

## 14. tee

`tee` reads standard input and writes it both to standard output and a file.

Example:

```bash
echo "hello" | tee output.txt
```

The text is displayed in the terminal and also written to `output.txt`.

### Append Instead of Overwrite

```bash
echo "hello" | tee -a output.txt
```

`tee` is useful when you want to see command output while also saving it.

---

# Pipes

A pipe (`|`) sends the standard output of one command to the standard input
of another command.

Example:

```bash
cat application.log | grep "error"
```

This sends the output of `cat` into `grep`.

In many cases, `cat` is unnecessary:

```bash
grep "error" application.log
```

Pipes become especially powerful when combining multiple tools.

### Count Errors

```bash
grep "error" application.log | wc -l
```

This:

1. Searches for lines containing `error`.
2. Sends those lines to `wc`.
3. Counts the resulting lines.

---

# Common Text-Processing Pipelines

## Count Errors in a Log

```bash
grep -i "error" application.log | wc -l
```

---

## Find the Most Common Values

```bash
sort values.txt | uniq -c | sort -nr
```

---

## Find the Top Five Results

```bash
sort values.txt | uniq -c | sort -nr | head -n 5
```

---

## Extract Usernames from `/etc/passwd`

```bash
cut -d ':' -f 1 /etc/passwd
```

---

## Find a Pattern and Display Selected Fields

```bash
grep "error" application.log | awk '{print $1, $5}'
```

---

## Monitor a Log for Errors

```bash
tail -f application.log | grep --line-buffered "error"
```

---

# Standard Input, Output, and Error

Linux programs normally have three standard streams.

## Standard Input

Standard input (`stdin`) is where a program receives input.

Example:

```bash
cat
```

The program can receive text from the terminal.

---

## Standard Output

Standard output (`stdout`) is where normal program output is written.

Example:

```bash
echo "hello"
```

The output normally appears in the terminal.

---

## Standard Error

Standard error (`stderr`) is used for error messages.

This allows normal output and error output to be handled separately.

---

# Output Redirection

## Redirect Standard Output

```bash
command > output.txt
```

This writes the output to `output.txt`.

If the file already exists, its contents are overwritten.

---

## Append Standard Output

```bash
command >> output.txt
```

This appends output to the file instead of overwriting it.

---

## Redirect Standard Error

```bash
command 2> errors.txt
```

This redirects standard error to `errors.txt`.

---

## Redirect Both Output and Errors

```bash
command > output.txt 2>&1
```

This redirects standard output to `output.txt` and then redirects standard
error to the same destination.

---

# Choosing the Right Tool

| Task                            | Command |   |
| ------------------------------- | ------- | - |
| Display a file                  | `cat`   |   |
| Read a large file interactively | `less`  |   |
| Show beginning of file          | `head`  |   |
| Show end of file                | `tail`  |   |
| Search for text                 | `grep`  |   |
| Sort lines                      | `sort`  |   |
| Remove adjacent duplicates      | `uniq`  |   |
| Count lines/words/bytes         | `wc`    |   |
| Extract columns/characters      | `cut`   |   |
| Translate/delete characters     | `tr`    |   |
| Edit/transform text streams     | `sed`   |   |
| Process structured text/columns | `awk`   |   |
| Compare files                   | `diff`  |   |
| Display and save output         | `tee`   |   |
| Connect commands                | `       | ` |

---

# Common Command Relationships

Many Linux commands become more useful when combined.

For example:

```bash
grep "error" application.log
```

Searches for errors.

```bash
grep "error" application.log | wc -l
```

Counts errors.

```bash
grep "error" application.log | cut -d ':' -f 1
```

Extracts a field from matching lines.

```bash
grep "error" application.log | sort | uniq -c
```

Groups identical matching lines and counts them.

```bash
tail -f application.log | grep --line-buffered "error"
```

Continuously watches a log for new error messages.

The main idea is that Linux text-processing utilities are designed to be
combined rather than used in isolation.

---

# Summary

The most important text-processing commands are:

* `cat` - display or combine files
* `less` - interactively read files
* `head` - display the beginning of a file
* `tail` - display the end of a file
* `grep` - search for patterns
* `sort` - sort lines
* `uniq` - detect or remove adjacent duplicates
* `wc` - count lines, words, and bytes
* `cut` - extract fields or characters
* `tr` - translate or delete characters
* `sed` - transform text streams
* `awk` - process structured text
* `diff` - compare files
* `tee` - display and save output
* `|` - connect commands together

Understanding these commands and how they work together provides a strong
foundation for Linux command-line text processing.