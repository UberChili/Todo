# Introduction
This is a personal project.

The basic idea is to tackle the classic to-do application but with a CLI functionality instead of the typical method of implementing it as a web application.

# Integration with bullet journaling and Obsidian
The greater idea behind this is that I've been using a classic physical Bullet Journal in my day to day, in order to write down some stuff, like thoughts, activities, appointments, and other things. This is as flexible as it can be considering it is pen and paper.
However, I also use Obsidian on my devices to take notes. Obsidian has a "Dailies" functionality which allows the user to create a new note for that specific day.

The idea is to be able to have a CLI interface to quickly "migrate" activities, todo elements, or appointments to my PC, so that I can quickly take a look at them from the terminal. At the same time, I want the tasks to be written into the Daily Note in my Obsidian Vault. This way, I can have access to my tasks on paper and on all my devices. It sounds simple enough but it should make for a small but decent improvement of productivity.

# Prerequisites
- Go 1.16 or later

# Instructions
1. Clone the repository:

``` sh
git clone git@github.com:UberChili/Todoapp.git
```

2. Set an environment variable in your profile, that points to your Dailies directory
For example, mine is:

``` sh
export DAILIES_DIRECTORY="$HOME/Notes/Dailies"
```

3. Build and run:

``` sh
go build -o todo
./todo
# or:
go run .
```

# Usage
List your open tasks like:

``` sh
todo

1       - [ ] Buy eggs
2       - [ ] Buy milk
3       - [ ] Start working on todo CLI app
```

To remove a task, specify its id:

``` sh
todo -r 2

todo

1       - [ ] Buy eggs
2       - [ ] Start working on todo CLI app
```

## TODO Add new task
