# Linux File Permissions and Ownership

Linux uses file permissions and ownership to control who can read, modify, or execute files and directories.

## Permission Model

Each file and directory has three basic permission categories:

* **user (u)** — the owner of the file
* **group (g)** — users belonging to the file's group
* **others (o)** — everyone else

The three basic permissions are:

* **read (r)**
* **write (w)**
* **execute (x)**

For a regular file:

* `r` allows reading the file.
* `w` allows modifying the file.
* `x` allows executing the file as a program.

For a directory:

* `r` allows listing directory entries.
* `w` allows creating, deleting, or renaming entries, subject to other permission constraints.
* `x` allows accessing entries within the directory and using the directory in pathnames.

---

## Viewing Permissions

Use `ls -l` to display file permissions and ownership.

```bash
ls -l file.txt
```

Example:

```text
-rw-r--r-- 1 almas users 1234 Aug 10 14:00 file.txt
```

The first part:

```text
-rw-r--r--
```

contains the file type and permissions.

The first character indicates the file type:

```text
-    regular file
d    directory
l    symbolic link
```

The remaining nine characters are divided into three groups:

```text
rw- r-- r--
 │   │   │
 │   │   └── others
 │   └────── group
 └────────── owner
```

For:

```text
-rw-r--r--
```

the owner has:

```text
rw-
```

The group has:

```text
r--
```

Others have:

```text
r--
```

---

# chmod

`chmod` changes the permissions of files and directories.

## Symbolic Mode

Basic syntax:

```bash
chmod [options] mode file
```

Give the owner execute permission:

```bash
chmod u+x script.sh
```

Remove execute permission from the owner:

```bash
chmod u-x script.sh
```

Give the group write permission:

```bash
chmod g+w file.txt
```

Remove write permission from others:

```bash
chmod o-w file.txt
```

Give everyone read permission:

```bash
chmod a+r file.txt
```

`a` means all users: owner, group, and others.

Multiple changes can be combined:

```bash
chmod u+rwx,g+rx,o-rwx script.sh
```

This gives the owner read/write/execute, gives the group read/execute, and removes all permissions from others.

---

## Numeric Mode

Permissions can also be represented using numbers.

The values are:

```text
r = 4
w = 2
x = 1
```

Add the values together for each permission category.

For example:

```text
rwx = 4 + 2 + 1 = 7
rw- = 4 + 2     = 6
r-x = 4 + 1     = 5
r-- = 4         = 4
-wx = 2 + 1     = 3
-w- = 2         = 2
--x = 1         = 1
--- = 0
```

Therefore:

```text
rwxr-xr-x
```

becomes:

```text
755
```

And:

```text
rw-r--r--
```

becomes:

```text
644
```

Example:

```bash
chmod 755 script.sh
```

This sets:

```text
owner   = rwx
group   = r-x
others  = r-x
```

Another common example:

```bash
chmod 644 file.txt
```

This sets:

```text
owner   = rw-
group   = r--
others  = r--
```

---

# Recursive Permissions

The `-R` option applies changes recursively to a directory and its contents.

Example:

```bash
chmod -R 755 directory/
```

This changes permissions for the directory and everything below it.

Be careful with recursive permission changes because they can unintentionally modify permissions of many files.

---

# chown

`chown` changes the owner of a file or directory.

Basic syntax:

```bash
chown [options] owner[:group] file
```

Change the owner:

```bash
sudo chown alice file.txt
```

Change the owner and group:

```bash
sudo chown alice:developers file.txt
```

Change the group while keeping the owner:

```bash
sudo chown :developers file.txt
```

Recursive ownership change:

```bash
sudo chown -R alice:developers project/
```

Use recursive ownership changes carefully.

---

# chgrp

`chgrp` changes the group ownership of a file or directory.

Example:

```bash
chgrp developers file.txt
```

Recursive example:

```bash
chgrp -R developers project/
```

---

# Ownership

Every file normally has:

* an owner
* a group

You can inspect ownership with:

```bash
ls -l
```

Example:

```text
-rw-r--r-- 1 almas developers 1234 Aug 10 14:00 file.txt
```

Here:

```text
owner = almas
group = developers
```

---

# umask

`umask` controls the default permission bits that are removed when new files and directories are created.

View the current umask:

```bash
umask
```

Example:

```text
0022
```

A common umask of `022` means that newly created objects do not normally receive write permission for group and others.

The exact resulting permissions also depend on the permissions requested by the program creating the file.

---

# Permission Troubleshooting

## Permission denied when executing a script

Check the permissions:

```bash
ls -l script.sh
```

If the execute permission is missing, add it:

```bash
chmod +x script.sh
```

Then run:

```bash
./script.sh
```

---

## Permission denied when accessing a directory

Check the directory permissions:

```bash
ls -ld directory/
```

Remember that directory `x` permission controls whether the directory can be accessed/traversed.

For example:

```bash
chmod u+x directory/
```

adds directory traversal permission for the owner.

---

## File is owned by another user

Check ownership:

```bash
ls -l file.txt
```

If you need to change ownership and have the required privileges:

```bash
sudo chown $USER file.txt
```

The appropriate ownership change depends on the situation; avoid changing ownership blindly.

---

## `sudo chmod` vs `sudo chown`

These commands solve different problems.

`chmod` changes:

```text
permissions
```

Example:

```bash
sudo chmod 644 file.txt
```

`chown` changes:

```text
ownership
```

Example:

```bash
sudo chown alice file.txt
```

Having the wrong owner and having insufficient permissions are related but different problems.

---

# Common Permission Patterns

## 644

```text
rw-r--r--
```

Commonly used for regular files.

```bash
chmod 644 file.txt
```

Owner can read/write.

Group and others can read.

---

## 755

```text
rwxr-xr-x
```

Commonly used for executable files and directories.

```bash
chmod 755 script.sh
```

Owner can read/write/execute.

Group and others can read/execute.

---

## 700

```text
rwx------
```

Only the owner has permissions.

```bash
chmod 700 private-directory/
```

---

## 600

```text
rw-------
```

Only the owner can read/write.

```bash
chmod 600 private.txt
```

---

# Quick Reference

| Task                        | Command                         |
| --------------------------- | ------------------------------- |
| View permissions            | `ls -l file`                    |
| View directory permissions  | `ls -ld directory`              |
| Add execute permission      | `chmod +x file`                 |
| Add owner execute           | `chmod u+x file`                |
| Remove write permission     | `chmod -w file`                 |
| Set permissions numerically | `chmod 755 file`                |
| Change owner                | `chown user file`               |
| Change owner and group      | `chown user:group file`         |
| Change group                | `chgrp group file`              |
| Change recursively          | `chmod -R ...` / `chown -R ...` |
| View umask                  | `umask`                         |

## Key Distinctions

```text
chmod  → permissions
chown  → owner
chgrp  → group
umask  → default permission mask for newly created objects
```
