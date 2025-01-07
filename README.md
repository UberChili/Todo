# Introduction
This is a personal project.

The basic idea is to implement the classic to-do application but with a CLI functionality instead of the typical method of implementing it as a web application. This is because I plan to use this myself for personal use in my daily life, and because I really like using the terminal.

# Integration with bullet journaling and Obsidian
I've been using a classic physical Bullet Journal in my day to day, in order to write down some stuff, like thoughts, activities, appointments, and other things. This is as flexible as it can be considering it is pen and paper. There is a great deal of flexibility about being able to use pen and paper for taking every day notes.
However, I also spend a lot of time in the computer and like terminal applications. I also use Obsidian on my devices to take and sync notes. This is very useful considering Obsidian's mobile application.

Obsidian has a "Dailies" functionality which allows the user to create a new note for that specific day.

The idea is to be able to have a CLI interface to "migrate" activities, todo elements, or appointments to be able to view them on my computers, so that I can quickly take a look at them from the terminal, like this, I want the tasks to be written into the Daily Note in my Obsidian Vault. This way, I can have access to my tasks on paper and on all my devices, eve my phone! It sounds simple enough but it should make for a small but decent improvement of productivity.

# Prerequisites
- Go 1.16 or later

# Instructions
1. Clone the repository:

``` sh
git clone git@github.com:UberChili/Todoapp.git
```

2. Set an environment variable in your profile that points to your dailies directory.

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
1       Configure NeoVim
2       Buy eggs
3       Buy milk
```
This will list your open tasks in the format shown above. All your open tasks will use markdown format in its corresponding file.

For example, a file opened in a text editor would look like this:
``` markdown
## Work
- [x] Start working on todo CLI app
- [x] This is a completed task
- [ ] Configure NeoVim

## Store
- [ ] Buy eggs
- [ ] Buy milk
```
To remove a task, use "-r" and specify its id:
``` sh
todo -r 2

todo

1       Configure NeoVim
2       Buy milk
```
Add a new task using the -add flag:

``` sh
$ todo -add "Write some code"
$ todo

1       Write some code
```
This will check if the file with the correct nomenclature already exists. If it doesn't, it will create it and add the new and only task. If it already exists, it will append the new task to the end of the file.

At first, I was considering the idea about establishing a special format or area for the tasks, something that, with some existing tasks and after creating a new task, in markdown could look like:
``` markdown
# Tasks for the day
- [ ] Buy eggs
- [ ] Buy milk
- [ ] New example would be inserted right here, and not in the end of the file.

This is some example test that was contained in the note before adding the third task, above.
Oh it was a good day!
```
I decided to not go for this approach and to instead just add new tasks at the bottom of each note because due to the usage of Markdown files as notes, the user could enter new tasks anywhere in the file. Maybe the user wants to add a task especifically next to some text block in order to remember some relationship. It would not be nice to alter that every time we add a new task. 

## TODO
- Prettier formatting. Maybe some colors or maybe using a library to make this a real TUI.
