# Tmux Themepack

A pack of various themes for Tmux.


## Installation

### Install manually

1.  Clone repo to local machine:

        git clone https://github.com/jimeh/tmux-themepack.git ~/.tmux-themepack

2.  Source desired theme in your `~/.tmux.conf`:

         source-file "${HOME}/.tmux-themepack/powerline/default/green.tmuxtheme"

    In some linux distributions you might have to remove the quotation marks
    from the `source-file` command to avoid a `no such file or directory` error:

         source-file ${HOME}/.tmux-themepack/powerline/default/green.tmuxtheme

### Install using [Tmux Plugin Manager](https://github.com/tmux-plugins/tpm)

1.  Add plugin to the list of TPM plugins in `.tmux.conf`:

        set -g @plugin 'jimeh/tmux-themepack'

2.  Press `prefix + I` to fetch the plugin and source it. The plugin should now
    be working.

Choose which theme is loaded by setting the `@themepack` option in your
`.tmux.conf`:

- `set -g @themepack 'basic'` (default)
- `set -g @themepack 'powerline/block/blue'`
- `set -g @themepack 'powerline/block/cyan'`
- `set -g @themepack 'powerline/default/green'`
- `set -g @themepack 'powerline/double/magenta'`
- `...`

## Themes

### Basic Themes

**Default (`default`):**

![basic](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/default-preview.png)

**Basic (`basic`):**

![basic](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/basic-preview.png)

### Powerline Themes

Inspired by the [Powerline][] VIM plugin, and requires the use of a powerline
compatible font in your terminal. You can find a number of such fonts in the
[powerline-fonts][] project.

If your preferred font isn't available there, please refer to Powerline's
documentation on [Font Patching][] to patch the font yourself.

[powerline]: https://github.com/Lokaltog/powerline
[powerline-fonts]: https://github.com/Lokaltog/powerline-fonts
[font patching]: https://powerline.readthedocs.org/en/latest/fontpatching.html#font-patching

**Powerline Blue (`powerline/default/blue`):**

![powerline-default-blue](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/default/blue-preview.png)

**Powerline Cyan (`powerline/default/cyan`):**

![powerline-default-cyan](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/default/cyan-preview.png)

**Powerline Gray (`powerline/default/gray`):**

![powerline-default-gray](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/default/gray-preview.png)

**Powerline Green (`powerline/default/green`):**

![powerline-default-green](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/default/green-preview.png)

**Powerline Magenta (`powerline/default/magenta`):**

![powerline-default-magenta](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/default/magenta-preview.png)

**Powerline Orange (`powerline/default/orange`):**

![powerline-default-orange](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/default/orange-preview.png)

**Powerline Purple (`powerline/default/purple`):**

![powerline-default-purple](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/default/purple-preview.png)

**Powerline Red (`powerline/default/red`):**

![powerline-default-red](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/default/red-preview.png)

**Powerline Yellow (`powerline/default/yellow`):**

![powerline-default-yellow](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/default/yellow-preview.png)

#### Block

Currently selected window is indicated by a colored block.

**Powerline Blue Block (`powerline/block/blue`):**

![powerline-block-blue](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/block/blue-preview.png)

**Powerline Cyan Block (`powerline/block/cyan`):**

![powerline-block-cyan](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/block/cyan-preview.png)

**Powerline Gray Block (`powerline/block/gray`):**

![powerline-block-gray](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/block/gray-preview.png)

**Powerline Green Block (`powerline/block/green`):**

![powerline-block-green](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/block/green-preview.png)

**Powerline Magenta Block (`powerline/block/magenta`):**

![powerline-block-magenta](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/block/magenta-preview.png)

**Powerline Orange Block (`powerline/block/orange`):**

![powerline-block-orange](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/block/orange-preview.png)

**Powerline Purple Block (`powerline/block/purple`):**

![powerline-block-purple](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/block/purple-preview.png)

**Powerline Red Block (`powerline/block/red`):**

![powerline-block-red](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/block/red-preview.png)

**Powerline Yellow Block (`powerline/block/yellow`):**

![powerline-block-yellow](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/block/yellow-preview.png)

#### Double

Both left and right far sides of the statusbar are colored, rather than just the
left side.

**Powerline Double Blue (`powerline/double/blue`):**

![powerline-double-blue](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/double/blue-preview.png)

**Powerline Double Cyan (`powerline/double/cyan`):**

![powerline-double-cyan](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/double/cyan-preview.png)

**Powerline Double Green (`powerline/double/green`):**

![powerline-double-green](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/double/green-preview.png)

**Powerline Double Magenta (`powerline/double/magenta`):**

![powerline-double-magenta](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/double/magenta-preview.png)

**Powerline Double Orange (`powerline/double/orange`):**

![powerline-double-orange](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/double/orange-preview.png)

**Powerline Double Purple (`powerline/double/purple`):**

![powerline-double-purple](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/double/purple-preview.png)

**Powerline Double Red (`powerline/double/red`):**

![powerline-double-red](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/double/red-preview.png)

**Powerline Double Yellow (`powerline/double/yellow`):**

![powerline-double-yellow](https://github.com/jimeh/tmux-themepack-previews/raw/1.0.0/powerline/double/yellow-preview.png)


## Tips

- Use different themes/colors on different machines by using some sort of
  wrapper around launching Tmux.


## Contributing

If you decide to contribute your own tmux themes, please try to base it on the
`default.tmuxtheme` theme. This ensures that switching between themes works as
it should and completely overwrites all settings from previous themes.

If it's not possible to base your theme on my default one, something is probably
missing from it. So please contribute a fix to the default theme too in that
case :)


## License

```
        DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
                    Version 2, December 2004

 Copyright (C) 2013 Jim Myhrberg

 Everyone is permitted to copy and distribute verbatim or modified
 copies of this license document, and changing it is allowed as long
 as the name is changed.

            DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
   TERMS AND CONDITIONS FOR COPYING, DISTRIBUTION AND MODIFICATION

  0. You just DO WHAT THE FUCK YOU WANT TO.
```
