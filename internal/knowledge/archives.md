# Linux Archives and Compression

This document covers common Linux commands for creating, viewing, extracting, and managing archives and compressed files.

---

## 1. Archives vs Compression

An **archive** combines multiple files and directories into a single file.

A **compressed file** reduces the amount of storage space needed.

These are different concepts.

For example:

- `tar` → creates archives
- `gzip` → compresses data
- `tar.gz` → a tar archive compressed with gzip
- `zip` → combines and compresses files

A `.tar` file is not necessarily compressed.

A `.tar.gz` file is both archived and compressed.

---

# 2. tar

`tar` is one of the most common Linux tools for creating and extracting archives.

Basic syntax:

```bash
tar [options] archive files
````

---

## 3. Create a tar Archive

Create an archive containing files:

```bash
tar -cf archive.tar file1.txt file2.txt
```

Create an archive containing a directory:

```bash
tar -cf archive.tar my_directory/
```

### Important options

```text
-c    create an archive
-f    specify the archive filename
-v    verbose output
-x    extract an archive
-t    list archive contents
```

Example:

```bash
tar -cvf backup.tar my_directory/
```

This creates `backup.tar` containing `my_directory`.

---

# 4. List Contents of a tar Archive

To see what is inside an archive:

```bash
tar -tf archive.tar
```

For more detailed output:

```bash
tar -tvf archive.tar
```

This does not extract the archive.

---

# 5. Extract a tar Archive

Extract an archive into the current directory:

```bash
tar -xf archive.tar
```

Extract to a specific directory:

```bash
tar -xf archive.tar -C /path/to/directory/
```

The destination directory should normally exist before extraction.

---

# 6. gzip Compression

`gzip` compresses individual files.

Compress a file:

```bash
gzip file.txt
```

This normally produces:

```text
file.txt.gz
```

The original uncompressed file is normally replaced by the compressed version.

---

## Decompress gzip

Use:

```bash
gunzip file.txt.gz
```

Or:

```bash
gzip -d file.txt.gz
```

Both decompress the gzip file.

---

# 7. tar.gz Archives

A `.tar.gz` archive is a tar archive compressed using gzip.

It is also commonly written as:

```text
.tgz
```

These are commonly used for distributing Linux source code, applications, backups, and collections of files.

---

## Create a tar.gz Archive

```bash
tar -czf archive.tar.gz my_directory/
```

Verbose version:

```bash
tar -czvf archive.tar.gz my_directory/
```

Important options:

```text
-c    create
-z    use gzip compression
-v    verbose
-f    specify filename
```

---

## Extract tar.gz

```bash
tar -xzf archive.tar.gz
```

Verbose:

```bash
tar -xzvf archive.tar.gz
```

Extract to a specific directory:

```bash
tar -xzf archive.tar.gz -C /path/to/directory/
```

---

# 8. tar.bz2

A `.tar.bz2` archive uses bzip2 compression.

Create:

```bash
tar -cjf archive.tar.bz2 my_directory/
```

Extract:

```bash
tar -xjf archive.tar.bz2
```

Important option:

```text
-j    use bzip2 compression
```

---

# 9. tar.xz

A `.tar.xz` archive uses xz compression.

Create:

```bash
tar -cJf archive.tar.xz my_directory/
```

Extract:

```bash
tar -xJf archive.tar.xz
```

Important option:

```text
-J    use xz compression
```

---

# 10. zip

`zip` is another common archive and compression utility.

Create a zip archive:

```bash
zip archive.zip file1.txt file2.txt
```

To include a directory recursively:

```bash
zip -r archive.zip my_directory/
```

Important option:

```text
-r    recursively include directories
```

---

# 11. unzip

Extract a zip archive:

```bash
unzip archive.zip
```

Extract to a specific directory:

```bash
unzip archive.zip -d /path/to/directory/
```

List contents without extracting:

```bash
unzip -l archive.zip
```

---

# 12. Common Archive Formats

| Extension  | Tool        | Compression     |
| ---------- | ----------- | --------------- |
| `.tar`     | tar         | None            |
| `.tar.gz`  | tar + gzip  | gzip            |
| `.tgz`     | tar + gzip  | gzip            |
| `.tar.bz2` | tar + bzip2 | bzip2           |
| `.tar.xz`  | tar + xz    | xz              |
| `.zip`     | zip/unzip   | zip compression |
| `.gz`      | gzip        | gzip            |

---

# 13. Common tar Flag Combinations

### Create archive

```bash
tar -cf archive.tar directory/
```

### Create compressed gzip archive

```bash
tar -czf archive.tar.gz directory/
```

### Create compressed gzip archive with verbose output

```bash
tar -czvf archive.tar.gz directory/
```

### Extract archive

```bash
tar -xf archive.tar
```

### Extract gzip archive

```bash
tar -xzf archive.tar.gz
```

### List archive

```bash
tar -tf archive.tar
```

### List compressed archive

```bash
tar -tzf archive.tar.gz
```

---

# 14. Archive Specific Files

You can specify multiple files:

```bash
tar -cf documents.tar file1.txt file2.txt file3.txt
```

You can also use patterns:

```bash
tar -cf logs.tar *.log
```

This archives files ending in `.log` from the current directory.

---

# 15. Exclude Files

`tar` can exclude files or directories.

Example:

```bash
tar -czf backup.tar.gz project/ --exclude=project/node_modules
```

Exclude multiple patterns:

```bash
tar -czf backup.tar.gz project/ \
    --exclude=project/node_modules \
    --exclude=project/.git
```

This is useful when creating backups while avoiding large or unnecessary directories.

---

# 16. Extract a Specific File

You do not always need to extract everything.

First list the archive:

```bash
tar -tf archive.tar
```

Then extract a specific file:

```bash
tar -xf archive.tar path/to/file.txt
```

For a gzip archive:

```bash
tar -xzf archive.tar.gz path/to/file.txt
```

---

# 17. Check Archive Before Extracting

Before extracting an archive, it is often useful to inspect its contents:

```bash
tar -tf archive.tar.gz
```

This helps determine:

* What files will be extracted
* Whether the archive contains a top-level directory
* Whether unexpected paths are present

---

# 18. Common Errors

## "tar: command not found"

The `tar` program may not be installed.

Check:

```bash
which tar
```

or:

```bash
command -v tar
```

---

## "gzip: stdin: not in gzip format"

The file may not actually be gzip-compressed.

Check the file type:

```bash
file archive.tar.gz
```

Do not assume the file's extension accurately describes its contents.

---

## "Cannot open: No such file or directory"

Check whether the archive exists:

```bash
ls -l archive.tar.gz
```

Check your current directory:

```bash
pwd
```

---

## Permission denied while extracting

The current user may not have permission to write to the destination directory.

Check:

```bash
ls -ld /path/to/directory
```

You may need to choose a directory you have permission to write to.

Avoid using `sudo` automatically. First determine why permission is denied.

---

# 19. Useful Commands for Identifying Archives

Use `file`:

```bash
file archive.tar.gz
```

Example output might indicate:

```text
gzip compressed data
```

or:

```text
POSIX tar archive
```

This is useful when the filename extension is misleading or unknown.

---

# 20. Practical Examples

### Backup a directory

```bash
tar -czf backup.tar.gz important_files/
```

### Extract a backup

```bash
tar -xzf backup.tar.gz
```

### See what is inside a backup

```bash
tar -tzf backup.tar.gz
```

### Create a zip archive

```bash
zip -r project.zip project/
```

### Extract a zip archive

```bash
unzip project.zip
```

### Extract somewhere else

```bash
unzip project.zip -d extracted/
```

---

# 21. Quick Reference

```text
tar -cf archive.tar files/          Create tar archive
tar -xf archive.tar                 Extract tar archive
tar -tf archive.tar                 List tar contents

tar -czf archive.tar.gz files/     Create gzip archive
tar -xzf archive.tar.gz             Extract gzip archive
tar -tzf archive.tar.gz             List gzip archive

tar -cjf archive.tar.bz2 files/    Create bzip2 archive
tar -xjf archive.tar.bz2             Extract bzip2 archive

tar -cJf archive.tar.xz files/     Create xz archive
tar -xJf archive.tar.xz             Extract xz archive

gzip file.txt                       Compress file
gunzip file.txt.gz                  Decompress gzip

zip archive.zip files               Create zip archive
zip -r archive.zip directory/       Zip directory recursively
unzip archive.zip                   Extract zip
unzip -l archive.zip                List zip contents
```

---

# 22. Key Concepts for the Assistant

When answering archive-related questions:

* Explain the difference between archiving and compression when relevant.
* `tar` primarily handles archives.
* `gzip`, `bzip2`, and `xz` provide compression.
* `.tar.gz` means tar + gzip.
* `.tar.bz2` means tar + bzip2.
* `.tar.xz` means tar + xz.
* `zip` combines archiving and compression.
* Use `tar -t` to inspect an archive without extracting it.
* Use `file` when the actual file format is uncertain.
* Do not recommend `sudo` automatically for permission errors.
* Warn users before destructive extraction or overwriting files when appropriate.