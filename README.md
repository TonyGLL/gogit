# GoGit - A Git Replica in Go

GoGit is a simplified implementation of Git, the distributed version control system, built from scratch in Go. This project is intended for educational purposes, providing a hands-on approach to understanding the core concepts and inner workings of Git.

## Features

*   **Initialize a new repository:** Create a new `.gogit` directory to start tracking a project.
*   **Add files to the staging area:** Track new or modified files to be included in the next commit.
*   **Commit changes:** Save snapshots of the staging area to the project's history.
*   **View commit history:** Inspect the log of commits to see the project's evolution.
*   **Manage branches:** Create, list, and delete branches.
*   **Configure user information:** Set your name and email for commit attribution.
*   **Check repository status:** View the status of tracked and untracked files.

## Getting Started

### Prerequisites

*   Go 1.18 or higher

### Installation

1.  Clone the repository:
    ```sh
    git clone https://github.com/TonyGLL/gogit.git
    ```
2.  Navigate to the project directory:
    ```sh
    cd gogit
    ```
3.  Build the project:
    ```sh
    go build .
    ```

## Usage

GoGit provides the following commands:

*   `gogit init`: Initializes a new repository.
*   `gogit add <file>`: Adds a file to the staging area.
*   `gogit commit -m <message>`: Commits the staged changes.
*   `gogit log`: Displays the commit history.
*   `gogit branch`: Lists all branches.
*   `gogit branch <name>`: Creates a new branch.
*   `gogit branch -d <name>`: Deletes a branch.
*   `gogit config user.name <name>`: Sets the user's name.
*   `gogit config user.email <email>`: Sets the user's email.
*   `gogit status`: Shows the status of the repository.

## Contributing

Contributions are welcome! If you'd like to improve GoGit, please feel free to fork the repository, make your changes, and submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
