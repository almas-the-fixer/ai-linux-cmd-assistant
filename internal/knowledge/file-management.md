# Linux File and Directory Management

Linux provides a collection of commands for creating, viewing, copying, moving, renaming, and deleting files and directories.

## ls

`ls` lists directory contents.

### Basic usage

```bash
ls
```

List a specific directory:

```bash
ls /var/log
```

### Useful options

`-l` — use a long listing format.

```bash
ls -l
```

`-a` — include hidden files.

```bash
ls -a
```

`-h` — display file sizes in a human-readable format. Usually combined with `-l`.

```bash
ls -lh
```

`-R` — recursively list subdirectories.

```bash
ls -R
```

`-t` — sort by modification time.

```bash
ls -lt
```

`-S` — sort by file size.

```bash
ls -lS
```

A commonly useful combination:

```bash
ls -lah
```

This displays hidden files in long format with human-readable sizes.

---

# pwd

`pwd` prints the current working directory.

```bash
pwd
```

Example output:

```text
/home/almas/projects
```

This is useful when you need to know where you currently are before performing file operations.

---

# cd

`cd` changes the current working directory.

Go to a directory:

```bash
cd /var/log
```

Go to a directory relative to the current directory:

```bash
cd projects
```

Go to the parent directory:

```bash
cd ..
```

Go to the user's home directory:

```bash
cd ~
```

Return to the previous working directory:

```bash
cd -
```

---

# mkdir

`mkdir` creates directories.

### Basic usage

```bash
mkdir project
```

Create multiple directories:

```bash
mkdir project1 project2 project3
```

### Creating parent directories

The `-p` option creates missing parent directories.

```bash
mkdir -p project/src/components
```

Without `-p`, the command fails if the parent directories do not already exist.

---

# rmdir

`rmdir` removes empty directories.

```bash
rmdir empty-directory
```

It normally fails if the directory contains files or other directories.

For removing a directory and its contents, `rm -r` is generally used instead.

---

# touch

`touch` updates file timestamps and can create an empty file if the specified file does not exist.

Create an empty file:

```bash
touch file.txt
```

Create multiple files:

```bash
touch file1.txt file2.txt file3.txt
```

`touch` does not specifically mean "create a file"; its primary purpose is updating file access/modification timestamps, with creation being a consequence when the file does not exist.

---

# cp

`cp` copies files or directories.

### Copy a file

```bash
cp source.txt destination.txt
```

Copy a file into a directory:

```bash
cp source.txt /tmp/
```

### Copy multiple files

```bash
cp file1.txt file2.txt /tmp/
```

### Copy directories

The `-r` option copies directories recursively.

```bash
cp -r source-directory destination-directory
```

### Preserve attributes

The `-p` option attempts to preserve file mode, ownership, and timestamps.

```bash
cp -p source.txt destination.txt
```

The `-a` option enables archive mode and is commonly used when copying directory trees while preserving metadata.

```bash
cp -a source-directory backup-directory
```

### Interactive copying

`-i` asks before overwriting an existing destination.

```bash
cp -i source.txt destination.txt
```

### Avoid overwriting

`-n` prevents overwriting an existing destination where supported by the implementation.

```bash
cp -n source.txt destination.txt
```

### Verbose output

`-v` displays what is being copied.

```bash
cp -v source.txt destination.txt
```

---

# mv

`mv` moves or renames files and directories.

### Rename a file

```bash
mv old-name.txt new-name.txt
```

### Move a file

```bash
mv file.txt /tmp/
```

### Move multiple files

```bash
mv file1.txt file2.txt /tmp/
```

### Rename a directory

```bash
mv old-directory new-directory
```

### Interactive mode

`-i` asks before overwriting an existing destination.

```bash
mv -i source.txt destination.txt
```

### Verbose mode

`-v` displays the operations being performed.

```bash
mv -v source.txt /tmp/
```

Unlike `cp`, `mv` does not create a second copy of the original file as a normal move operation; the directory entry is moved/renamed as appropriate.

---

# rm

`rm` removes files and directory entries.

### Remove a file

```bash
rm file.txt
```

### Remove multiple files

```bash
rm file1.txt file2.txt
```

### Interactive removal

`-i` asks for confirmation before removing files.

```bash
rm -i file.txt
```

### Remove directories recursively

`-r` or `-R` removes directories and their contents recursively.

```bash
rm -r directory/
```

### Force removal

`-f` ignores nonexistent files and does not prompt for confirmation.

```bash
rm -f file.txt
```

### Recursive force removal

```bash
rm -rf directory/
```

This can remove an entire directory tree without interactive confirmation.

**IMPORTANT: rm -rf / is DANGEROUS and running it with sudo can Cause Destructive data loss it can Wipe Out Full OS if ran with SUDO!** 
**Use `rm -rf` carefully.** Mistakes in the path can cause destructive data loss.

---

# File Paths

Linux commands commonly accept either absolute or relative paths.

## Absolute path

An absolute path starts from `/`.

```bash
/home/almas/projects/app/main.go
```

It identifies a location independently of the current working directory.

## Relative path

A relative path is interpreted from the current working directory.

```bash
projects/app/main.go
```

If the current directory is:

```text
/home/almas
```

then the relative path refers to:

```text
/home/almas/projects/app/main.go
```

---

# Special Path Components

`.` refers to the current directory.

```bash
ls .
```

`..` refers to the parent directory.

```bash
ls ..
```

`~` normally expands to the current user's home directory in a shell.

```bash
cd ~
```

`/` is the root of the filesystem hierarchy.

---

# Hidden Files

In Linux, filenames beginning with `.` are normally treated as hidden by commands such as `ls` unless explicitly requested.

Example:

```text
.bashrc
.gitconfig
.config/
```

Show hidden files:

```bash
ls -a
```

Show hidden files with details:

```bash
ls -la
```

---

# Wildcards and Globbing

The shell expands wildcard patterns before executing many commands.

`*` matches zero or more characters.

```bash
ls *.txt
```

This can match:

```text
file.txt
notes.txt
report.txt
```

`?` matches a single character.

```bash
ls file?.txt
```

This could match:

```text
file1.txt
file2.txt
```

Character classes use square brackets.

```bash
ls file[123].txt
```

This can match:

```text
file1.txt
file2.txt
file3.txt
```

Important: these patterns are **shell globbing**, not regular expressions.

---

# Symbolic Links

A symbolic link is a filesystem object that refers to another path.

Create a symbolic link:

```bash
ln -s target link-name
```

Example:

```bash
ln -s /var/log/app.log app.log
```

Inspect a symbolic link:

```bash
ls -l app.log
```

The output may look like:

```text
app.log -> /var/log/app.log
```

Remove a symbolic link with `rm`:

```bash
rm app.log
```

Removing the symbolic link does not normally remove its target.

---

# Hard Links

A hard link is another directory entry referring to the same underlying file data.

Create one:

```bash
ln existing.txt hardlink.txt
```

Unlike a symbolic link, a hard link does not contain a path pointing to the target.

Hard links generally cannot be created for directories by ordinary users, and they normally cannot cross filesystem boundaries.

---

# File Information

## stat

`stat` displays detailed information about a file or filesystem object.

```bash
stat file.txt
```

It can provide information such as:

* file type
* permissions
* ownership
* size
* inode
* access time
* modification time
* status-change time

---

# File Types

The `file` command attempts to determine the type of a file based on its contents and other information.

```bash
file document.txt
```

Example output might identify the object as ASCII text, UTF-8 text, an executable, an image, or another recognized format.

This can be useful when a filename extension is misleading or missing.

---

# Disk Usage

## du

`du` estimates filesystem space used by files and directories.

Check the size of a directory:

```bash
du -sh directory/
```

`-s` gives a summary.

`-h` uses human-readable units.

Check the sizes of entries in the current directory:

```bash
du -sh *
```

---

# Disk Free Space

## df

`df` reports filesystem disk space usage.

```bash
df
```

Human-readable output:

```bash
df -h
```

This is different from `du`.

```text
du → space used by files/directories
df → space available/used on filesystems
```

---

# Common Operations

## Create a directory and enter it

```bash
mkdir project
cd project
```

Or:

```bash
mkdir -p project/src
cd project/src
```

## Create a file

```bash
touch file.txt
```

## Copy a file

```bash
cp file.txt backup.txt
```

## Rename a file

```bash
mv old.txt new.txt
```

## Move a file

```bash
mv file.txt /tmp/
```

## Remove a file

```bash
rm file.txt
```

## Copy a directory

```bash
cp -r source/ destination/
```

## Remove a directory and its contents

```bash
rm -r directory/
```

---

# Common Troubleshooting

## `cp: cannot stat`

Example:

```text
cp: cannot stat 'file.txt': No such file or directory
```

Common causes:

* The file does not exist.
* The path is incorrect.
* The filename is misspelled.
* The current working directory is not what you expected.

Check:

```bash
pwd
ls
```

Then verify the path.

---

## `mv: cannot stat`

Similar to `cp`, this usually means the source path cannot be found.

Check:

```bash
pwd
ls
```

and verify the source path.

---

## `rm: cannot remove ...: Is a directory`

`rm` without recursive mode does not remove directories.

For an empty directory:

```bash
rmdir directory/
```

For a directory containing files:

```bash
rm -r directory/
```

---

## `Permission denied`

Check the permissions and ownership:

```bash
ls -l file
```

For a directory:

```bash
ls -ld directory/
```

See the permissions documentation for `chmod`, `chown`, and ownership troubleshooting.

---

# Important Command Distinctions

```text
pwd     → show current directory
cd      → change directory
ls      → list directory contents
mkdir   → create directory
rmdir   → remove empty directory
touch   → update timestamps / create missing file
cp      → copy
mv      → move or rename
rm      → remove
ln      → create links
stat    → show detailed file information
file    → identify file type
du      → estimate file/directory disk usage
df      → report filesystem disk usage
```

## Quick Reference

| Goal                      | Command                      |
| ------------------------- | ---------------------------- |
| Show current directory    | `pwd`                        |
| List files                | `ls`                         |
| List hidden files         | `ls -la`                     |
| Change directory          | `cd directory`               |
| Go to parent              | `cd ..`                      |
| Create directory          | `mkdir directory`            |
| Create nested directories | `mkdir -p path/to/directory` |
| Create/update a file      | `touch file`                 |
| Copy file                 | `cp source destination`      |
| Copy directory            | `cp -r source destination`   |
| Move/rename               | `mv source destination`      |
| Remove file               | `rm file`                    |
| Remove directory          | `rm -r directory`            |
| Create symbolic link      | `ln -s target link`          |
| File information          | `stat file`                  |
| Identify file type        | `file file`                  |
| Directory disk usage      | `du -sh directory`           |
| Filesystem disk usage     | `df -h`                      |
