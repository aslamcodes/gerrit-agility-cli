# Git CLI Tool

A simple command-line interface (CLI) tool for managing Git commits and pushes with ticket-based messages. This tool helps streamline the process of adding changes, committing them with a ticket number, and optionally pushing them to a remote branch.

## Features

- Add all changes to the staging area.
- Commit changes with a ticket number and a custom message.
- Optionally push changes to the `develop` branch.

## Prerequisites

- Go 1.24 or later
- Git installed on your machine

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/aslamcodes/gerrit-agility-cli.git
   cd gerrit-agility-cli
    ```

2. Build the Project

    ```bash
    go build -o gcommit
    ```

3. Install 
    ```bash
    go install ./...
    ```

## Usage

Run the CLI tool with the following command:

```bash
gcommit [-p]
