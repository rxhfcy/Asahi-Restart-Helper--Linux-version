# asahi-reboot-switcher
GUI for asahi-bless - reboot to macOS system tray menu

## Install dependencies

```
dnf install gcc make gtk3-devel libappindicator-gtk3-devel asahi-bless
```

If asahi-bless version <= 0.2.1:
```
git clone https://github.com/WhatAmISupposedToPutHere/asahi-nvram.git
cd asahi-nvram/asahi-bless
cargo install --path .
```

Also make sure `~/.cargo/bin` is in `PATH`.

## Build & install asahi-reboot-switcher
```
make
sudo make install
```

## KDE

To automatically start after login, add `/usr/bin/asahi-reboot-switcher` to `System Settings -> Startup and Shutdown -> Autostart`.
(https://userbase.kde.org/System_Settings/Autostart)

## Gnome

To automatically start after login,  add `/usr/bin/asahi-reboot-switcher` to `Tweaks -> Startup Applications`
(https://help.gnome.org/users/gnome-help/stable/shell-apps-auto-start.html.en)

The following extension must also be installed and enabled:
https://extensions.gnome.org/extension/615/appindicator-support/
