---
title: 'Minimal macOS setup'
description: 'My checklist for setting up a new Mac'
date: '2026-10-24'
cover:
  src: '../../assets/setup.png'
  alt: 'My minimal macOS setup'
---

I've set up three Macs from scratch in the past couple of months. By the third, I realized I was repeating the same steps, so I wrote them down for next time.

This is my checklist for a minimal macOS setup. I'll keep it updated as macOS and the tools change.

## Accounts and Homebrew

1. Sign in to iCloud.
2. Install Chrome and sign in.
3. Install the [Bitwarden extension](https://chromewebstore.google.com/detail/bitwarden-password-manage/nngceckbapebfimnlniiiahkandclblb) and sign in.
4. Install [Homebrew](https://brew.sh/).
5. Open my [Brewfile](https://github.com/hagelstam/dotfiles/blob/main/Brewfile) and download the raw file as `~/Downloads/Brewfile`. Then run:

```sh
brew bundle --file=~/Downloads/Brewfile
```

This installs all the apps and utilities I need (incl. Ghostty, Raycast, Zed, Starship). Remove any entries you don't need before running it.

## Font

Download [JetBrains Mono Nerd Font](https://www.nerdfonts.com/font-downloads), unzip it and install it with Font Book.

## System Settings

These are my preferences grouped by pane. Some options depend on the macOS version.

<details class="settings-checklist">
<summary>Desktop & Dock</summary>

- Minimized window animation: Scale Effect.
- Window title bar double-click action: Fill.
- Turn on minimizing windows into the application icon and automatic Dock hiding.
- Turn off opening animations and suggested or recent apps.
- **Desktop & Stage Manager**: turn off _Show Items_ both on the desktop and in Stage Manager. Set _Click wallpaper to show desktop_ to Only in Stage Manager.
- **Widgets**: turn off _Show Widgets_ both on the desktop and in Stage Manager.
- **Mission Control**: turn on _Group windows by application_.

</details>

<details class="settings-checklist">
<summary>Displays</summary>

Turn off _Automatically adjust brightness_ and _True Tone_.

</details>

<details class="settings-checklist">
<summary>Battery</summary>

- Set the energy mode on the power adapter to High Power.
- Under **Options**, turn on _Prevent automatic sleeping on power adapter when the display is off_.

</details>

<details class="settings-checklist">
<summary>Lock Screen</summary>

- _Turn display off on battery when inactive_: For 10 minutes.
- _Turn display off on power adapter when inactive_: For 30 minutes.
- Turn off _Show user name and photo_.

</details>

<details class="settings-checklist">
<summary>Keyboard</summary>

Turn off _Adjust keyboard brightness in low light_. Under **Keyboard Shortcuts**:

- **Spotlight**: disable both shortcuts to free up <kbd aria-label="Command Space">⌘ Space</kbd> for Raycast.
- **Mission Control**: disable _Quick Note_ and _Game Overlay_.
- **Modifier Keys**: remap <kbd>Caps Lock</kbd> to <kbd aria-label="Command">⌘</kbd>.
- **Keyboard**: set _Move focus to next window_ to <kbd aria-label="Command Section sign">⌘ §</kbd>.

</details>

<details class="settings-checklist">
<summary>Menu Bar</summary>

- Keep Wi-Fi, battery, focus and sound in the menu bar. Hide all the other items.
- Under **Battery**, turn on _Show percentage_.
- Under **Clock**, turn off _Show the day of the week_.
- Turn off _Show suggestions in Control Gallery_.

</details>

## Finder

Open the settings with <kbd aria-label="Command Comma">⌘ ,</kbd>

- **General**: uncheck everything under _Show these items on the desktop_. Set _New Finder windows show_ to your home folder and turn off _Suggest file names_.
- **Tags**: delete all unused tags.
- **Sidebar**: uncheck all the items you don't need.
- **Advanced**: turn on _Show all filename extensions_. Turn off the warnings before changing an extension or removing a file from iCloud Drive. Set _When performing a search_ to `Search the Current Folder`.

In your home folder, select **View > as List** and enable **View > Show Path Bar** and **View > Show Status Bar**. Reorder the sidebar as needed.

## Terminal and shell

Create the config files:

```sh
mkdir -p ~/.config/ghostty ~/.config/mise
touch ~/.config/ghostty/config ~/.config/mise/config.toml
touch ~/.zshrc ~/.gitconfig ~/.gitattributes
```

Copy the configs from my dotfiles into the corresponding files:

| Config | Destination |
| --- | --- |
| [Ghostty](https://github.com/hagelstam/dotfiles/blob/main/ghostty/config) | `~/.config/ghostty/config` |
| [mise](https://github.com/hagelstam/dotfiles/blob/main/mise/config.toml) | `~/.config/mise/config.toml` |
| [Zsh](https://github.com/hagelstam/dotfiles/blob/main/zsh/.zshrc) | `~/.zshrc` |
| [Git](https://github.com/hagelstam/dotfiles/blob/main/.gitconfig) | `~/.gitconfig` |
| [Git attributes](https://github.com/hagelstam/dotfiles/blob/main/.gitattributes) | `~/.gitattributes` |

Remove my `AWS_PROFILE` from the mise config and set your name and email in the Git config. Install the mise tools before opening a new shell, since the Zsh config uses them:

```sh
mise install
```

## Editor

I've been using [Zed](https://zed.dev/) lately, with the following extensions and configs.

### Extensions

- [Catppuccin Blur](https://zed.dev/extensions/catppuccin-blur)
- [Symbols](https://zed.dev/extensions/symbols)
- [Proto](https://zed.dev/extensions/proto)
- [TOML](https://zed.dev/extensions/toml)

### Settings

Copy my [settings](https://github.com/hagelstam/dotfiles/blob/main/zed/settings.json) to `~/.config/zed/settings.json`.

### Key bindings

Copy my [keymap](https://github.com/hagelstam/dotfiles/blob/main/zed/keymap.json) to `~/.config/zed/keymap.json`.

### Command line

Run `cli: install cli binary` from the command palette to enable `zed .` in the terminal.

## Raycast

Set <kbd aria-label="Command Space">⌘ Space</kbd> as the launch hotkey and configure the shortcuts below in settings.

### Shortcuts

<div class="shortcut-table">

| Action | Shortcut |
| --- | --- |
| Chrome | <kbd aria-label="Command Option 1">⌘ ⌥ 1</kbd> |
| Ghostty | <kbd aria-label="Command Option 2">⌘ ⌥ 2</kbd> |
| Zed | <kbd aria-label="Command Option 3">⌘ ⌥ 3</kbd> |
| Slack (once installed) | <kbd aria-label="Command Option 4">⌘ ⌥ 4</kbd> |
| Screenshot | <kbd aria-label="Command Option S">⌘ ⌥ S</kbd> |
| Search Emoji | <kbd aria-label="Command Option E">⌘ ⌥ E</kbd> |
| Lock Screen | <kbd aria-label="Command Option L">⌘ ⌥ L</kbd> |
| Pick Color | <kbd aria-label="Command Option C">⌘ ⌥ C</kbd> |

</div>

### Window management

Set window size cycling to `1/2, 2/3, and 1/3`. Then assign these shortcuts:

<div class="shortcut-table">

| Action | Shortcut |
| --- | --- |
| Left half | <kbd aria-label="Command Option Left arrow">⌘ ⌥ ←</kbd> |
| Right half | <kbd aria-label="Command Option Right arrow">⌘ ⌥ →</kbd> |
| Center third | <kbd aria-label="Command Option Up arrow">⌘ ⌥ ↑</kbd> |
| Maximize | <kbd aria-label="Command Option F">⌘ ⌥ F</kbd> |

</div>

### Extensions

- [Coffee](https://www.raycast.com/mooxl/coffee)
- [GIF Search](https://www.raycast.com/josephschmitt/gif-search)
- [Color Picker](https://www.raycast.com/thomas/color-picker)
- [Svgl](https://www.raycast.com/1weiho/svgl)
- [Kill Process](https://www.raycast.com/rolandleth/kill-process)
- [Port Manager](https://www.raycast.com/lucaschultz/port-manager)
- [Team Time](https://www.raycast.com/david_nakhapetian/team-time)
