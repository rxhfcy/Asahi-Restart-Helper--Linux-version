```diff
- TODO: We're evaluating 4 different UI options
- and need help deciding on the best compromise:
```

Linux version:

- [Option A (README)](./README_OPTION_A.md): Boot menu.
This option lets users select from all disks and restart **WITHOUT CHANGING** the default 'Startup Disk'.
However, it does provide a submenu for changing the default 'Startup Disk'.

- [Option B (README)](./README_OPTION_B.md): Essentially the same as [Option A](./README_OPTION_A.md),
but with an additional optional checkbox menu item
that allows users to **CHANGE** the default 'Startup Disk' upon restarting.

- **"Perhaps TOO Simple"** [Option X (README)](./README_OPTION_X.md): This option **ONLY** offers a
"Restart in macOS..." option.
This will **ALWAYS CHANGE** the default 'Startup Disk' to macOS (imitates Apple's Boot Camp behavior).
It **DOES NOT** provide a submenu to change the default 'Startup Disk' back to Linux.

- **"Simple"** [Option Y (README)](./README_OPTION_Y.md): Similar to [Option X](./README_OPTION_X.md)
(only "Restart in macOS", change default 'Startup Disk' to macOS),
but it also includes a submenu that allow users to **CHANGE** the default 'Startup Disk' back to Linux for convenience.

PS. There is also a [corresponding macOS version](https://github.com/rxhfcy/Asahi-Restart-Helper--macOS-version) of this application, which currently uses "Option Y".

Once we've determined the best compromise for the UI and scope (probably "Option Y"?), we will make the UI of both the macOS and Linux versions match as closely as possible.
