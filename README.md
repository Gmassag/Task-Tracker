# Task Tracker CLI

A simple command-line interface (CLI) application for managing tasks. Built with Go as part of the [roadmap.sh Go development track](https://roadmap.sh/golang).

## Features

- ✅ Add new tasks
- 📋 List all tasks or filter by status
- ✏️ Update task descriptions
- 🗑️ Delete tasks
- ✅ Mark tasks as done
- 🔄 Mark tasks as in-progress
- 💾 Persistent storage using JSON file

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Gmassag/Task-Tracker
cd Task-Tracker
```

2. Build the application:
```bash
go build -o task-cli
```

## Usage

### Add a new task
```bash
task-cli add "Buy groceries"
```

### List tasks
```bash
# List all tasks
task-cli list

# Filter by status
task-cli list done
task-cli list todo
task-cli list in-progress
```

### Update a task
```bash
task-cli update 1 "Buy groceries and cook dinner"
```

### Delete a task
```bash
task-cli delete 1
```

### Mark task as done
```bash
task-cli mark-done 1
```

### Mark task as in-progress
```bash
task-cli mark-in-progress 1
```

## Task Properties

Each task contains the following information:
- **ID**: Unique identifier
- **Description**: Task description
- **Status**: `todo`, `in-progress`, or `done`
- **CreatedAt**: Timestamp when task was created
- **UpdatedAt**: Timestamp when task was last modified

## Data Storage

Tasks are stored in a `tasks.json` file in the same directory as the executable. The file is created automatically when you add your first task.

## Example Output

```bash
$ task-cli add "Learn Go programming"
Task added successfully (ID: 1)

$ task-cli add "Build a CLI app"
Task added successfully (ID: 2)

$ task-cli list
Tasks:
[1] [ ] Learn Go programming - todo
[2] [ ] Build a CLI app - todo

$ task-cli mark-in-progress 1
Task marked as in-progress: Learn Go programming

$ task-cli mark-done 2
Task marked as done: Build a CLI app

$ task-cli list
Tasks:
[1] [~] Learn Go programming - in-progress
[2] [✓] Build a CLI app - done
```

## Project URL

This project is part of the roadmap.sh Go development track:
https://roadmap.sh/golang

**Specific project:** https://roadmap.sh/projects/task-tracker

## Requirements

- Go 1.19 or higher
