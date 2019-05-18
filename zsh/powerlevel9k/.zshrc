POWERLEVEL9K_RIGHT_PROMPT_ELEMENTS=(tasks)

# tasks
POWERLEVEL9K_CUSTOM_TASKS="prompt_tusk"
POWERLEVEL9K_CUSTOM_TASKS_FOREGROUND="045"
POWERLEVEL9K_CUSTOM_TASKS_BACKGROUND="none"

prompt_tusk() {
  if ! [ -x "$(command -v tusk)" ]; then
    exit 1
  else
    watch=$(tusk watch count)
    tasks=$(tusk list count)
    if [[ $watch =~ ^[0-9]+$ ]] && [[ $tasks =~ ^[0-9]+$ ]] ; then
      echo -e "$watch/$tasks task(s) \uf5c0"
    fi
  fi
}
