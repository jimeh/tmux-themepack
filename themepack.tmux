#!/usr/bin/env bash

CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

theme_option="@themepack"
default_theme='basic'

get_tmux_option() {
  local option="$1"
  local default_value="$2"
  local option_value="$(tmux show-option -gqv "$option")"

  if [ -n "$option_value" ]; then
    echo "$option_value"
  else
    echo "$default_value"
  fi
}

main() {
  local theme="$(get_tmux_option "$theme_option" "$default_theme")"
  if [ -f "$CURRENT_DIR/${theme}.tmuxtheme" ]; then
    tmux source-file "$CURRENT_DIR/${theme}.tmuxtheme"
  else
    tmux source-file "$CURRENT_DIR/powerline/${theme}.tmuxtheme"
  fi
}

main
