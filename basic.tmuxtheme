#
# Basic theme
#

# Themepack format options
set -goq @themepack-status-left-area-left-format "#S"
set -goq @themepack-status-left-area-middle-format "#I"
set -goq @themepack-status-left-area-right-format "#P"
set -goq @themepack-status-right-area-left-format "#H"
set -goq @themepack-status-right-area-middle-format "%H:%M:%S"
set -goq @themepack-status-right-area-right-format "%d-%b-%y"
set -goq @themepack-window-status-current-format "#I:#W#F"
set -goq @themepack-window-status-format "#I:#W#F"

# Customizable prefixes and suffixes for @themepack-* format options
set -goq @themepack-status-left-area-left-prefix ""
set -goq @themepack-status-left-area-left-suffix ""
set -goq @themepack-status-left-area-middle-prefix ""
set -goq @themepack-status-left-area-middle-suffix ""
set -goq @themepack-status-left-area-right-prefix ""
set -goq @themepack-status-left-area-right-suffix ""
set -goq @themepack-status-right-area-left-prefix ""
set -goq @themepack-status-right-area-left-suffix ""
set -goq @themepack-status-right-area-middle-prefix ""
set -goq @themepack-status-right-area-middle-suffix ""
set -goq @themepack-status-right-area-right-prefix ""
set -goq @themepack-status-right-area-right-suffix ""
set -goq @themepack-window-status-current-prefix ""
set -goq @themepack-window-status-current-suffix ""
set -goq @themepack-window-status-prefix ""
set -goq @themepack-window-status-suffix ""

# Apply prefixes and suffixes to @themepack-* format options
set -gqF @themepack-status-left-area-left-format "#{@themepack-status-left-area-left-prefix}#{@themepack-status-left-area-left-format}#{@themepack-status-left-area-left-suffix}"
set -gqF @themepack-status-left-area-middle-format "#{@themepack-status-left-area-middle-prefix}#{@themepack-status-left-area-middle-format}#{@themepack-status-left-area-middle-suffix}"
set -gqF @themepack-status-left-area-right-format "#{@themepack-status-left-area-right-prefix}#{@themepack-status-left-area-right-format}#{@themepack-status-left-area-right-suffix}"
set -gqF @themepack-status-right-area-left-format "#{@themepack-status-right-area-left-prefix}#{@themepack-status-right-area-left-format}#{@themepack-status-right-area-left-suffix}"
set -gqF @themepack-status-right-area-middle-format "#{@themepack-status-right-area-middle-prefix}#{@themepack-status-right-area-middle-format}#{@themepack-status-right-area-middle-suffix}"
set -gqF @themepack-status-right-area-right-format "#{@themepack-status-right-area-right-prefix}#{@themepack-status-right-area-right-format}#{@themepack-status-right-area-right-suffix}"
set -gqF @themepack-window-status-current-format "#{@themepack-window-status-current-prefix}#{@themepack-window-status-current-format}#{@themepack-window-status-current-suffix}"
set -gqF @themepack-window-status-format "#{@themepack-window-status-prefix}#{@themepack-window-status-format}#{@themepack-window-status-suffix}"

# Theme options
set -goq  @theme-clock-mode-colour red
set -goq  @theme-clock-mode-style 24
set -goq  @theme-display-panes-active-colour default
set -goq  @theme-display-panes-colour default
set -goq  @theme-message-bg default
set -goq  @theme-message-command-bg default
set -goq  @theme-message-command-fg default
set -goq  @theme-message-fg default
set -goq  @theme-mode-bg red
set -goq  @theme-mode-fg default
set -goq  @theme-pane-active-border-bg default
set -goq  @theme-pane-active-border-fg green
set -goq  @theme-pane-border-bg default
set -goq  @theme-pane-border-fg default
set -goq  @theme-status-bg black
set -goq  @theme-status-fg cyan
set -goq  @theme-status-interval 1
set -goq  @theme-status-justify centre
set -goqF @theme-status-left "#{@themepack-status-left-area-left-format} #[fg=white]» #[fg=yellow]#{@themepack-status-left-area-middle-format} #[fg=cyan]#{@themepack-status-left-area-right-format}"
set -goq  @theme-status-left-bg black
set -goq  @theme-status-left-fg green
set -goq  @theme-status-left-length 40
set -goqF @theme-status-right "#{@themepack-status-right-area-left-format} #[fg=white]« #[fg=yellow]#{@themepack-status-right-area-middle-format} #[fg=green]#{@themepack-status-right-area-right-format}"
set -goq  @theme-status-right-bg black
set -goq  @theme-status-right-fg cyan
set -goq  @theme-status-right-length 40
set -goq  @theme-window-status-activity-bg black
set -goq  @theme-window-status-activity-fg yellow
set -goq  @theme-window-status-current-bg red
set -goq  @theme-window-status-current-fg black
set -goqF @theme-window-status-current-format " #{@themepack-window-status-current-format} "
set -goqF @theme-window-status-format " #{@themepack-window-status-format} "
set -goq  @theme-window-status-separator ""

# Customizable prefixes and suffixes for @theme-* format options
set -goq @theme-status-left-prefix ""
set -goq @theme-status-left-suffix ""
set -goq @theme-status-right-prefix ""
set -goq @theme-status-right-suffix ""
set -goq @theme-window-status-current-prefix ""
set -goq @theme-window-status-current-suffix ""
set -goq @theme-window-status-prefix ""
set -goq @theme-window-status-suffix ""

# Apply prefixes and suffixes to @theme-* format options
set -gqF @theme-status-left "#{@theme-status-left-prefix}#{@theme-status-left}#{@theme-status-left-suffix}"
set -gqF @theme-status-right "#{@theme-status-right-prefix}#{@theme-status-right}#{@theme-status-right-suffix}"
set -gqF @theme-window-status-current-format "#{@theme-window-status-current-prefix}#{@theme-window-status-current-format}#{@theme-window-status-current-suffix}"
set -gqF @theme-window-status-format "#{@theme-window-status-prefix}#{@theme-window-status-format}#{@theme-window-status-suffix}"

# Apply @theme-* options to Tmux
set -gF  display-panes-active-colour "#{@theme-display-panes-active-colour}"
set -gF  display-panes-colour "#{@theme-display-panes-colour}"
set -gF  message-command-style "fg=#{@theme-message-command-fg},bg=#{@theme-message-command-bg}"
set -gF  message-style "fg=#{@theme-message-fg},bg=#{@theme-message-bg}"
set -gF  status-interval "#{@theme-status-interval}"
set -gF  status-justify "#{@theme-status-justify}"
set -gF  status-left "#{@theme-status-left}"
set -gF  status-left-length "#{@theme-status-left-length}"
set -gF  status-left-style "fg=#{@theme-status-left-fg},bg=#{@theme-status-left-bg}"
set -gF  status-right "#{@theme-status-right}"
set -gF  status-right-length "#{@theme-status-right-length}"
set -gF  status-right-style "fg=#{@theme-status-right-fg},bg=#{@theme-status-right-bg}"
set -gF  status-style "fg=#{@theme-status-fg},bg=#{@theme-status-bg}"
set -gwF clock-mode-colour "#{@theme-clock-mode-colour}"
set -gwF clock-mode-style "#{@theme-clock-mode-style}"
set -gwF mode-style "fg=#{@theme-mode-fg},bg=#{@theme-mode-bg}"
set -gwF pane-active-border-style "fg=#{@theme-pane-active-border-fg},bg=#{@theme-pane-active-border-bg}"
set -gwF pane-border-style "fg=#{@theme-pane-border-fg},bg=#{@theme-pane-border-bg}"
set -gwF window-status-activity-style "fg=#{@theme-window-status-activity-fg},bg=#{@theme-window-status-activity-bg}"
set -gwF window-status-current-format "#{@theme-window-status-current-format}"
set -gwF window-status-current-style "fg=#{@theme-window-status-current-fg},bg=#{@theme-window-status-current-bg}"
set -gwF window-status-format "#{@theme-window-status-format}"
set -gwF window-status-separator "#{@theme-window-status-separator}"
