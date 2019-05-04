POWERLEVEL9K_RIGHT_PROMPT_ELEMENTS=(tasks)

# tasks
POWERLEVEL9K_CUSTOM_TASKS="prompt_taskar"
POWERLEVEL9K_CUSTOM_TASKS_FOREGROUND="045"
POWERLEVEL9K_CUSTOM_TASKS_BACKGROUND="none"

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
