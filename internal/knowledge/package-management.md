# Linux Package Management

This document covers common Linux package management commands and concepts.

Package managers are used to install, remove, update, search, and inspect software packages on Linux systems.

Different Linux distributions use different package managers.

---

# 1. What Is a Package?

A package is a collection of files and metadata required to install software.

A package may contain:

- Executable programs
- Libraries
- Configuration files
- Documentation
- Metadata
- Dependency information

Package managers handle installing and removing these packages while resolving dependencies.

---

# 2. Common Linux Package Managers

Different Linux distributions commonly use different package management systems.

| Distribution | Package Manager | Package Format |
|---|---|---|
| Debian | apt | `.deb` |
| Ubuntu | apt | `.deb` |
| Linux Mint | apt | `.deb` |
| Fedora | dnf | `.rpm` |
| RHEL | dnf | `.rpm` |
| CentOS Stream | dnf | `.rpm` |
| Arch Linux | pacman | `.pkg.tar.*` |
| openSUSE | zypper | `.rpm` |

The commands in this document are grouped by package manager.

---

# 3. APT

`apt` is commonly used on Debian-based distributions such as Debian and Ubuntu.

Check whether apt is available:

```bash
command -v apt
````

---

# 4. Update Package Information

Before installing software, it is often useful to update the local package lists.

```bash
sudo apt update
```

This downloads updated package information from configured repositories.

It does not normally upgrade installed packages.

---

# 5. Upgrade Installed Packages

Upgrade installed packages:

```bash
sudo apt upgrade
```

To perform a more comprehensive upgrade:

```bash
sudo apt full-upgrade
```

`full-upgrade` may install or remove packages when necessary to complete the upgrade.

---

# 6. Install a Package

Install a package:

```bash
sudo apt install package-name
```

Example:

```bash
sudo apt install curl
```

Install multiple packages:

```bash
sudo apt install curl git wget
```

APT resolves required dependencies automatically.

---

# 7. Remove a Package

Remove an installed package:

```bash
sudo apt remove package-name
```

Example:

```bash
sudo apt remove nginx
```

Removing a package may leave configuration files behind.

To remove the package and its system configuration files:

```bash
sudo apt purge package-name
```

---

# 8. Remove Unused Dependencies

APT can remove packages that were installed as dependencies but are no longer required:

```bash
sudo apt autoremove
```

Review the packages that will be removed before confirming.

---

# 9. Search for Packages

Search package names and descriptions:

```bash
apt search keyword
```

Example:

```bash
apt search nginx
```

---

# 10. Show Package Information

Display information about a package:

```bash
apt show package-name
```

Example:

```bash
apt show curl
```

This can show:

* Package version
* Description
* Dependencies
* Repository
* Installed size
* Package status

---

# 11. Check Whether a Package Is Installed

Use:

```bash
dpkg -s package-name
```

Example:

```bash
dpkg -s curl
```

You can also use:

```bash
apt list --installed
```

To search installed packages:

```bash
apt list --installed | grep package-name
```

---

# 12. List Upgradable Packages

```bash
apt list --upgradable
```

This shows installed packages for which newer versions are available.

---

# 13. Find Which Package Provides a File

On Debian-based systems, `dpkg -S` can search installed package ownership:

```bash
dpkg -S /path/to/file
```

Example:

```bash
dpkg -S /usr/bin/curl
```

This is useful when you have a file and want to determine which installed package provided it.

---

# 14. Installing a Local `.deb` File

APT can install a local `.deb` file:

```bash
sudo apt install ./package.deb
```

The `./` is important when referring to a package in the current directory.

Another tool is `dpkg`:

```bash
sudo dpkg -i package.deb
```

If dependencies are missing after using `dpkg`, APT can often resolve them:

```bash
sudo apt install -f
```

---

# 15. DNF

`dnf` is commonly used by Fedora and modern Red Hat-based distributions.

Check whether dnf is available:

```bash
command -v dnf
```

---

# 16. Search Packages with DNF

```bash
dnf search keyword
```

Example:

```bash
dnf search nginx
```

---

# 17. Install Packages with DNF

```bash
sudo dnf install package-name
```

Example:

```bash
sudo dnf install git
```

Multiple packages:

```bash
sudo dnf install git curl wget
```

---

# 18. Remove Packages with DNF

```bash
sudo dnf remove package-name
```

Example:

```bash
sudo dnf remove nginx
```

---

# 19. Update Packages with DNF

Update installed packages:

```bash
sudo dnf upgrade
```

On many systems:

```bash
sudo dnf update
```

may also be used.

---

# 20. Show Package Information with DNF

```bash
dnf info package-name
```

Example:

```bash
dnf info nginx
```

---

# 21. List Installed Packages with DNF

```bash
dnf list installed
```

Search installed packages:

```bash
dnf list installed | grep nginx
```

---

# 22. Find Which Package Provides a File

DNF can search repositories for packages providing a particular file:

```bash
dnf provides /path/to/file
```

Example:

```bash
dnf provides /usr/bin/htop
```

This can help determine which package contains a command.

---

# 23. Pacman

`pacman` is the primary package manager for Arch Linux and Arch-based distributions.

Check whether pacman is available:

```bash
command -v pacman
```

---

# 24. Install Packages with Pacman

Install a package:

```bash
sudo pacman -S package-name
```

Example:

```bash
sudo pacman -S git
```

Install multiple packages:

```bash
sudo pacman -S git curl wget
```

---

# 25. Update Packages with Pacman

Update the package database and upgrade installed packages:

```bash
sudo pacman -Syu
```

The common meaning is:

```text
-S    synchronize packages
-y    refresh package databases
-u    upgrade installed packages
```

---

# 26. Remove Packages with Pacman

Remove a package:

```bash
sudo pacman -R package-name
```

Remove a package and dependencies that are no longer required:

```bash
sudo pacman -Rs package-name
```

Use removal commands carefully.

---

# 27. Search Packages with Pacman

Search repositories:

```bash
pacman -Ss keyword
```

Example:

```bash
pacman -Ss nginx
```

---

# 28. Search Installed Packages

Search locally installed packages:

```bash
pacman -Qs keyword
```

---

# 29. Show Package Information

For a package in repositories:

```bash
pacman -Si package-name
```

For an installed package:

```bash
pacman -Qi package-name
```

---

# 30. List Installed Packages

```bash
pacman -Q
```

To display explicitly installed packages:

```bash
pacman -Qe
```

---

# 31. Find Which Package Owns a File

For an installed file:

```bash
pacman -Qo /path/to/file
```

Example:

```bash
pacman -Qo /usr/bin/bash
```

---

# 32. Package Repositories

Package managers normally download packages from configured repositories.

Repositories provide:

* Package files
* Package metadata
* Version information
* Dependency information
* Security updates

Repository configuration differs between distributions.

Do not blindly add random third-party repositories.

Prefer official repositories when possible.

---

# 33. Dependencies

Software packages often depend on other packages.

For example:

```text
Application
    |
    +-- Library A
    |
    +-- Library B
            |
            +-- Library C
```

Package managers resolve these dependencies automatically when possible.

If a package cannot be installed because of dependency problems, inspect the error before attempting random fixes.

---

# 34. Package Not Found

If a package cannot be found, first check:

```bash
apt search package-name
```

or:

```bash
dnf search package-name
```

or:

```bash
pacman -Ss package-name
```

Also check whether the package manager's package databases are up to date.

For APT:

```bash
sudo apt update
```

For Arch:

```bash
sudo pacman -Sy
```

Avoid partial upgrades on Arch. Prefer:

```bash
sudo pacman -Syu
```

instead of upgrading only the package database.

---

# 35. Command Not Found After Installation

If a package appears to be installed but the command cannot be found:

Check whether the command exists:

```bash
command -v command-name
```

Check the PATH:

```bash
echo "$PATH"
```

Search for the executable:

```bash
find /usr -name command-name 2>/dev/null
```

The package may install an executable with a different name than the package itself.

---

# 36. Check Installed Version

APT:

```bash
apt list --installed package-name
```

Or:

```bash
dpkg -s package-name
```

DNF:

```bash
dnf info package-name
```

Pacman:

```bash
pacman -Qi package-name
```

For many commands, the program itself also supports:

```bash
command-name --version
```

---

# 37. Package Manager vs Manual Installation

Package managers are generally preferred because they can:

* Track installed files
* Resolve dependencies
* Upgrade software
* Remove software cleanly
* Apply distribution security updates

Manually downloading binaries or compiling software from source may be appropriate in some situations, but it makes package tracking and upgrades more manual.

---

# 38. Common Permission Errors

Package installation usually requires administrator privileges because system directories are being modified.

Example:

```bash
apt install curl
```

may fail with a permissions error.

The normal command on Debian-based systems is:

```bash
sudo apt install curl
```

Similarly:

```bash
sudo dnf install curl
```

and:

```bash
sudo pacman -S curl
```

Do not recommend `sudo` blindly if the command is only querying package information.

For example, these usually do not require root:

```bash
apt search curl
apt show curl
dnf search curl
pacman -Ss curl
```

---

# 39. Common Package Management Troubleshooting

## APT package database problems

If APT reports problems with package configuration, inspect the error carefully.

A commonly used command for unfinished package configuration is:

```bash
sudo dpkg --configure -a
```

Then:

```bash
sudo apt install -f
```

These commands should be used based on the actual error rather than blindly.

---

## Package Manager Is Locked

APT may report that another process is using the package database.

Check running processes:

```bash
ps aux | grep apt
```

or:

```bash
ps aux | grep dpkg
```

Another package-management process may currently be running.

Do not immediately delete lock files. Determine which process is holding the lock first.

---

# 40. Package Manager Quick Reference

## Debian / Ubuntu

```text
sudo apt update
sudo apt upgrade
sudo apt install package
sudo apt remove package
sudo apt purge package
sudo apt autoremove
apt search package
apt show package
apt list --installed
apt list --upgradable
```

## Fedora / RHEL

```text
sudo dnf install package
sudo dnf remove package
sudo dnf upgrade
dnf search package
dnf info package
dnf list installed
dnf provides /path/to/file
```

## Arch Linux

```text
sudo pacman -S package
sudo pacman -Syu
sudo pacman -R package
sudo pacman -Rs package
pacman -Ss package
pacman -Qs package
pacman -Qi package
pacman -Qo /path/to/file
```

---

# 41. Important Concepts for the Assistant

When answering package-management questions:

* First identify the user's Linux distribution when the command depends on it.
* Debian and Ubuntu commonly use `apt`.
* Fedora and modern RHEL-based systems commonly use `dnf`.
* Arch Linux uses `pacman`.
* Package managers handle dependencies.
* Prefer official repositories when possible.
* Do not recommend `sudo` for commands that do not require administrative privileges.
* Do not tell users to blindly delete package-manager lock files.
* Do not recommend random repository changes without understanding the distribution.
* On Arch Linux, avoid recommending partial upgrades.
* When troubleshooting a package installation failure, inspect the actual error message first.
* `command -v` can determine whether an installed command is available in the current `PATH`.
* Package names and command names are not necessarily identical.