# E-pacseed

E-pacseed is a declarative system building and restoring tool for Arch Linux. It allows users to describe their target system state (packages from official repositories and AUR) in a simple YAML manifest, and then apply this state to any machine without manual intervention. It eliminates configuration drift and makes system reproducibility easy.

## Prerequisites

To build the project, you must have **Go** installed on your system. Also, since the tool interacts with system packages, `pacman`, `git`, and `base-devel` are required for the restore functionality to work properly.

## Build

To compile the project, run:
```
make build
```
This will create the executable binary in the `bin/` directory.

To clean up the compiled binaries, run:
```
make clean
```
## Usage

After a successful build, you can run the tool using:
```
./bin/pacseed [command]
```
### Commands

*   **`compile`**
    Scans your currently installed packages and generates a YAML manifest.
        ```
        ./bin/pacseed compile --path "path/to/save" --name "my-config"
        ```
    *(If `--path` is not provided, it saves to your home directory. If `--name` is not provided, it defaults to `repos`.)*

*   **`restore`**
    Reads the specified YAML manifest and installs all listed packages. It automatically checks for an AUR helper (yay or paru) and offers to install one if missing.
        ```
        ./bin/pacseed restore --path "path/to/config" --name "my-config"
        ```
    *(Both `--path` and `--name` are required for this command).*

*   **`check`**
    Validates the configuration file and provides information about the repositories listed in the manifest.
      ```
        ./bin/pacseed check --repos
        ./bin/pacseed check --aur
        ./bin/pacseed check --repo "package-name"
      ```
