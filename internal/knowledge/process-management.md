# Process Management in Linux

## Overview

A process is a running instance of a program.

Linux manages processes by assigning each process a unique Process ID (PID), allocating CPU and memory resources, and controlling how processes interact with the system.

Process management is useful for:

* Finding running programs
* Checking resource usage
* Starting and stopping programs
* Running programs in the background
* Finding processes consuming too many resources
* Terminating unresponsive programs
* Understanding parent and child processes

---

## 1. `ps`

The `ps` command displays information about currently running processes.

### Basic usage

```bash
ps
```

Shows processes associated with the current terminal.

### Show all processes

```bash
ps aux
```

Commonly used to display detailed information about processes running on the system.

Important columns include:

* `USER` - user who owns the process
* `PID` - process ID
* `%CPU` - CPU usage
* `%MEM` - memory usage
* `STAT` - process state
* `COMMAND` - command that started the process

### Another common form

```bash
ps -ef
```

Shows processes in a full-format listing.

### Find a specific process

```bash
ps aux | grep nginx
```

This searches the output for processes containing `nginx`.

---

## 2. `top`

The `top` command provides a real-time view of running processes and system resource usage.

```bash
top
```

It displays information such as:

* CPU usage
* Memory usage
* Load average
* Running processes
* Process IDs
* Process owners

Useful interactive keys inside `top` include:

```text
q
```

Quit `top`.

```text
k
```

Send a signal to a process.

```text
M
```

Sort processes by memory usage.

```text
P
```

Sort processes by CPU usage.

---

## 3. `htop`

`htop` is an interactive process viewer with a more user-friendly interface than `top`.

```bash
htop
```

It may not be installed by default.

On systems using `apt`:

```bash
sudo apt install htop
```

On Arch-based systems:

```bash
sudo pacman -S htop
```

Useful for:

* Monitoring CPU usage
* Monitoring memory usage
* Finding processes
* Killing processes interactively
* Viewing process hierarchy

---

## 4. Process IDs (PID)

Every running process has a Process ID (PID).

For example:

```text
PID
1234
```

The PID can be used to identify and control a process.

To find the PID of a process:

```bash
pgrep nginx
```

You can also use:

```bash
pidof nginx
```

Example:

```bash
pidof firefox
```

---

## 5. `pgrep`

`pgrep` searches for processes based on their name or other attributes.

```bash
pgrep nginx
```

Returns the PID of matching processes.

### Show process names with PIDs

```bash
pgrep -a nginx
```

### Search for a user-owned process

```bash
pgrep -u username
```

---

## 6. `pidof`

`pidof` finds the process IDs of a running program.

```bash
pidof firefox
```

Example output:

```text
1234 5678
```

Multiple PIDs can be returned if multiple instances are running.

---

## 7. Foreground and Background Processes

Normally, when you run a command in a terminal, it runs in the foreground.

Example:

```bash
sleep 60
```

The terminal waits until the command finishes.

A command can be started in the background by adding `&`:

```bash
sleep 60 &
```

The shell returns control to the terminal immediately.

---

## 8. `jobs`

The `jobs` command displays jobs running in the current shell.

```bash
jobs
```

Example:

```text
[1]+  Running    sleep 60 &
```

The number inside `[ ]` is the job ID.

---

## 9. `bg`

The `bg` command resumes a stopped job in the background.

For example:

```bash
sleep 100
```

Press:

```text
Ctrl+Z
```

This suspends the process.

Then:

```bash
bg
```

resumes it in the background.

---

## 10. `fg`

The `fg` command brings a background or stopped job to the foreground.

```bash
fg
```

For a specific job:

```bash
fg %1
```

Here `%1` refers to job number 1.

---

## 11. `Ctrl+C`

`Ctrl+C` sends an interrupt signal to the foreground process.

For example:

```bash
ping example.com
```

Pressing:

```text
Ctrl+C
```

normally stops the command.

The signal generated is generally `SIGINT`.

---

## 12. `Ctrl+Z`

`Ctrl+Z` suspends the foreground process.

Example:

```bash
sleep 100
```

Press:

```text
Ctrl+Z
```

The process becomes stopped rather than terminated.

It can then be resumed with:

```bash
fg
```

or:

```bash
bg
```

---

## 13. `kill`

The `kill` command sends a signal to a process.

Despite its name, `kill` does not always terminate a process. It can send different signals.

Basic usage:

```bash
kill PID
```

Example:

```bash
kill 1234
```

By default, `kill` normally sends `SIGTERM`.

---

## 14. `SIGTERM`

`SIGTERM` asks a process to terminate gracefully.

```bash
kill -TERM 1234
```

or:

```bash
kill -15 1234
```

Programs can handle `SIGTERM` and perform cleanup before exiting.

It should generally be preferred over immediately using `SIGKILL`.

---

## 15. `SIGKILL`

`SIGKILL` immediately terminates a process.

```bash
kill -KILL 1234
```

or:

```bash
kill -9 1234
```

A process cannot catch or ignore `SIGKILL`.

Use it when a process refuses to terminate normally.

Example:

```bash
kill -9 1234
```

### Important

Do not use `kill -9` as the first solution.

Prefer:

```bash
kill PID
```

and only use `SIGKILL` when necessary.

---

## 16. `pkill`

`pkill` sends a signal to processes based on their name or other attributes.

Example:

```bash
pkill firefox
```

This can terminate matching processes.

To send `SIGTERM` explicitly:

```bash
pkill -TERM firefox
```

To force termination:

```bash
pkill -KILL firefox
```

Use process-name based commands carefully because they can affect multiple processes.

---

## 17. `killall`

`killall` can send a signal to processes based on their name.

```bash
killall firefox
```

This may terminate all matching `firefox` processes.

Force termination:

```bash
killall -9 firefox
```

Be careful when using commands that can affect multiple processes.

---

## 18. Process States

The `STAT` column in commands such as `ps` shows process state.

Common states include:

```text
R
```

Running or runnable.

```text
S
```

Sleeping.

```text
D
```

Uninterruptible sleep, commonly associated with waiting for I/O.

```text
T
```

Stopped or traced.

```text
Z
```

Zombie process.

---

## 19. Zombie Processes

A zombie process is a process that has finished execution but still has an entry in the process table because its parent has not yet collected its exit status.

You may see a zombie process with:

```bash
ps aux
```

or:

```bash
ps -ef
```

A zombie is already finished and cannot be killed normally because it is not actually running.

The parent process usually needs to properly reap it.

---

## 20. Parent and Child Processes

Processes can create other processes.

The process that creates another process is called the parent process.

The newly created process is called the child process.

You can view process relationships with:

```bash
ps -ef
```

The `PPID` column represents the Parent Process ID.

Example:

```text
UID   PID   PPID   CMD
user  2000  1000   bash
user  2100  2000   sleep 60
```

Here, process `2000` is the parent of process `2100`.

---

## 21. `pstree`

`pstree` displays processes in a tree structure.

```bash
pstree
```

This makes parent-child relationships easier to understand.

To include PIDs:

```bash
pstree -p
```

Example:

```text
systemd(1)
 ├─bash(1200)
 │  └─sleep(1300)
 └─sshd(900)
```

---

## 22. Finding Resource-Hungry Processes

Use:

```bash
top
```

or:

```bash
htop
```

To inspect CPU usage with `ps`:

```bash
ps aux --sort=-%cpu
```

To inspect memory usage:

```bash
ps aux --sort=-%mem
```

These commands are useful when investigating a system that feels slow.

---

## 23. `nice`

Linux processes have a scheduling priority represented by a nice value.

The `nice` command starts a process with a specified nice value.

Example:

```bash
nice -n 10 command
```

A higher nice value generally means the process receives lower CPU scheduling priority.

Nice values commonly range from:

```text
-20
```

to:

```text
19
```

Lower values generally mean higher priority.

---

## 24. `renice`

`renice` changes the nice value of an already-running process.

Example:

```bash
renice 10 -p 1234
```

This changes the nice value of process `1234`.

Changing a process to a higher priority may require elevated privileges.

---

## 25. `nohup`

`nohup` runs a command so that it can continue running after the terminal or shell session is closed.

Example:

```bash
nohup ./server &
```

Output is commonly redirected to:

```text
nohup.out
```

unless output is explicitly redirected elsewhere.

Example:

```bash
nohup ./server > server.log 2>&1 &
```

---

## 26. Running a Process in the Background

A command can be placed in the background using:

```bash
command &
```

Example:

```bash
./server &
```

The shell remains available for other commands.

To see background jobs:

```bash
jobs
```

---

## 27. Checking Whether a Process Exists

Use:

```bash
pgrep process_name
```

Example:

```bash
pgrep nginx
```

If a PID is returned, a matching process exists.

You can also use:

```bash
ps aux | grep nginx
```

However, `pgrep` is generally cleaner for process-name searches.

---

## 28. Common Troubleshooting Workflow

When a program appears to be running incorrectly:

### Step 1: Find the process

```bash
ps aux | grep program
```

or:

```bash
pgrep -a program
```

### Step 2: Check resource usage

```bash
top
```

or:

```bash
htop
```

### Step 3: Try graceful termination

```bash
kill PID
```

### Step 4: Check whether it stopped

```bash
pgrep program
```

### Step 5: Force termination only if necessary

```bash
kill -9 PID
```

---

## 29. Useful Command Summary

| Command   | Purpose                             |
| --------- | ----------------------------------- |
| `ps`      | Display processes                   |
| `ps aux`  | Detailed process listing            |
| `ps -ef`  | Full process listing                |
| `top`     | Real-time process monitoring        |
| `htop`    | Interactive process monitoring      |
| `pgrep`   | Find processes by name              |
| `pidof`   | Find PIDs of a program              |
| `pstree`  | Display process hierarchy           |
| `kill`    | Send a signal to a process          |
| `pkill`   | Signal processes by name            |
| `killall` | Signal processes by name            |
| `jobs`    | Show shell jobs                     |
| `bg`      | Resume a job in the background      |
| `fg`      | Bring a job to the foreground       |
| `nice`    | Start a process with a nice value   |
| `renice`  | Change a process's nice value       |
| `nohup`   | Keep a process running after logout |

---

## 30. Common Questions

### How do I see all running processes?

```bash
ps aux
```

or:

```bash
ps -ef
```

### How do I find a process by name?

```bash
pgrep process_name
```

### How do I kill a process?

```bash
kill PID
```

### How do I force kill a process?

```bash
kill -9 PID
```

### How do I monitor processes in real time?

```bash
top
```

or:

```bash
htop
```

### How do I run a command in the background?

```bash
command &
```

### How do I bring a background job to the foreground?

```bash
fg
```

### How do I resume a stopped job in the background?

```bash
bg
```

### How do I find processes using lots of CPU?

```bash
ps aux --sort=-%cpu
```

### How do I find processes using lots of memory?

```bash
ps aux --sort=-%mem
```

### How do I see the parent-child relationship between processes?

```bash
pstree -p
```