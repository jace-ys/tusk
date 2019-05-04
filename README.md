[![Build Status][ci-build]][ci-status]

[ci-build]:	https://img.shields.io/travis/com/jace-ys/taskar/master.svg?style=for-the-badge&logo=travis
[ci-status]: https://travis-ci.com/jace-ys/taskar

# Taskar

A command-line interface for managing your to-do list.

## Installation

Ensure Go (version >1.11) is installed and that your $GOPATH/bin is added to $PATH before installing Taskar.

Clone the repository, create the database path, and install the Taskar binary using the following commands:

```
git clone https://github.com/jace-ys/taskar
make install
```

## Usage

List of commands:

```
Usage:
  taskar [command]

Available Commands:
  add         Add a new task to your to-do list
  amend       Amend the name of a task on your to-do list
  comment     Add a comment message to a task on your to-do list
  done        Mark a task on your to-do list as complete
  find        Display a list of filtered tasks from your to-do list
  help        Help about any command
  list        Display all tasks on your to-do list
  list count  Count the number of tasks on your to-do list
  rm          Delete all tasks on your to-do list
  set-due     Set a due date for a task on your to-do list
  tag         Tag a category to a task on your to-do list
  unwatch     Unwatch a task on your to-do list
  watch       Watch a task on your to-do list
  watch count Count the number of watched tasks on your to-do list
  watch list  Display all watched tasks on your to-do list

Flags:
  -h, --help      help for taskar
      --version   version for taskar

Use "taskar [command] --help" for more information about a command.
```

## Custom Zsh Prompt

A sample [.zshrc](https://github.com/jace-ys/taskar/blob/master/zsh/powerlevel9k/.zshrc) file for the [powerlevel9k](https://github.com/bhilburn/powerlevel9k) zsh theme has been included.

Once Taskar is installed, add the following lines to your .zshrc file to add a custom prompt to your terminal:

```
POWERLEVEL9K_RIGHT_PROMPT_ELEMENTS=(tasks)

# tasks
POWERLEVEL9K_CUSTOM_TASKS="prompt_taskar"
POWERLEVEL9K_CUSTOM_TASKS_FOREGROUND="045"
POWERLEVEL9K_CUSTOM_TASKS_BACKGROUND="none"

# custom prompt that displays the number of watched tasks / total number of tasks
# eg. 3/5 task(s)
prompt_taskar() {
  if ! [ -x "$(command -v taskar)" ]; then
    exit 1
  else
    watch=$(taskar watch count)
    tasks=$(taskar list count)
    if [[ $watch =~ ^[0-9]+$ ]] && [[ $tasks =~ ^[0-9]+$ ]] ; then
      echo -e "$watch/$tasks task(s) \uf5c0"
    fi
  fi
}
```
