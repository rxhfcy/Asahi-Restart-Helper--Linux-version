# TODO LIST:

## Priority system for TODO items:
- **"1:" - Must have.** Critical, must not be omitted.
- **"2:" - Should have.** Important but not strictly necessary, can be delayed.
- **"3:" - Could have.** Desirable but not crucial, included if resources allow.
- **"4:" - Won't have.** Lowest priority, currently not planned.

## TODO items:
(please keep this TODO list up-to-date)

- [ ] (todo: add link to macOS version's TODO list)

# (NOTE: if you tweak something, please keep macOS/Linux TODOs synced!)

- [ ] 1: Decide if there should be a pref to also always change the default Startup Disk when restarting?
    - ...in the menu itself "[✔️] Also Change Default"?
    - ...or in the Preferences window?
        - would have to change the menu items to explain "Restart in macOS (and make default)..."
    - Will some users prefer that
        - reason why: normal restart is never surprising and loads the same OS next?
        - why not: just always use this app to restart into same OS too
    - "New prefs + score creep = bad"!

- [ ] 1: Figure out how to add submenu items in Go

- [ ] 1: Don't try to use clever macOS detection. A disk is a disk.
    - This disk is "Linux"
    - ...so the other one must be "macOS" (!)

- [ ] 1: Read and understand current code
    - How is asahi-bless called?
    - etc

- [ ] 1: Is the project name ok?

- [ ] 1: Mock UI to change current default Startup Disk (submenu)
    - remove checkboxes
    - omit current default Startup Disk from list
    - reason why in-app: Linux needs this to be in-app because no official GUI for this
        (the macOS version of the app also offers this, for parity)

- [ ] 2: Actually implement changing current default Startup Disk (submenu)

- [ ] 1: Add "Restart in Linux..."
    - for users who want to easily restart in Linux but keep macOS the default Startup Disk

- [ ] 2: UI: Actually show "(default)" for actual current default Startup Disk in menu

- [ ] 2: Implement Preferences dialog/window
    - main menu item

- [ ] 2: Implement setting "automatically load at startup" straight in preferences
    "It boils down to writing a config file under home directory which is then also accessible from DE configuration
    (at least on KDE, Gnome)"

- [ ] 2: Consider including / bundle / vendor asahi-bless binary directly inside app
    - don't depend on external program

- [ ] 1: Move "Quit" to preferences

- [ ] 1: Function to get list of every available Startup Disk
    - like this or is there a better method?
        1) call asahi-bless -l
        2) parse text: get numbers...
        3) parse text: get disk names...
        4) parse text: "*" means "default"...
        5) don't try to detect "Macintosh" drives, just assume that if there is only 1 other disk, 
            it's probably the "macOS" one
            - (very safe bet, Asahi Linux is essentially broken if there is no macOS disk)

- [ ] 1: UI: List all disks in main menu, allow booting to any of them
    - if just macOS and Linux, show "macOS" and "Linux"
    - if more than 2 disks, list every disk, including this one, with full disk names
        (e.g. "Restart in 'Fedora Linux'...")
    - show the text "(default)" after current default Startup Disk

- [ ] 2: UI: Actually show "(default)" for actual current default Startup Disk in menu

- [ ] 1: Make it possible to add app to Login Items using the app itself
    - in preferences
    - can I use some else's work?
    - if super simple, DIY?

- [ ] 2: Automatically add to / remove from Login Items if pref value disagrees with reality

- [ ] 2: Make searching for "restart", "reboot", "macOS" etc work
    - at least in KDE "start menu"?
    - tweak aliases, what else?
    - where (what files specifically?)

- [ ] 2: Proper Help dialog
    - help icon in upper right corner of Preferences or something?
    - explain what the app does and why
      - "this app uses Asahi `asahi-bless` tool with --next to temporarily make e.g. macOS load
      after restarting next time (i.e. without having to change the default Startup Disk value)"
    - bonus: teach the possibility of using `sudo asahi-bless` in Terminal.app and then restarting normally
    (asahi-bless CLI is currently somewhat suboptimal though, will hopefully change) -> change OS for next time only
    - bonus: teach newbies the essential skill of holding down the power button to get to boot picker (show animation?)
    - bonus: mention the macOS version of this tool
        and Apple's official bless / official System Options -> Startup Options
    - link to Asahi web page (explain that we are not affiliated, at least not yet)?

- [ ] 2: Quit confirmation dialog
    - reason: user might click accidentally and not realize this removes the icon, explain that
        - especially if autostart on login not enabled: show a different message and a checkbox to add to Login Items?

- [ ] 2: Make sure that current password-free doesn't need to be installed/uninstalled(?)
    - Read the code

- [ ] 2: Detect if user cancelled an ongoing restart (i.e. used sudo bless to change --nextonly but didn't restart)
    - otherwise, next "normal" restart (can be weeks from now!) will be very surprising ("why did the wrong OS load")? 
    - before asahi-bless command, store "NVRAM before" variable?
    - detect cancellation with e.g. a timer and/or checking for user interaction after having already "restarted"
    - read previously stored "NVRAM before", compare to "NVRAM now" and undo

- [ ] 1: App icon OK? Use Asahi icon for now?

- [ ] 1: Tray icon OK? Use Asahi icon for now?
    - should a template? icon be used instead to look more native(?)

- [ ] 1: Name of project ok? (Compare with macOS version)
    - (See macOS version TODO for ideas)
    
- [ ] 2: GitHub blurb ok? (Compare with macOS version)
    - (See macOS version TODO for ideas)

- [ ] 2: Scope creep? What should the project (not) do? (Compare with macOS version)
    - Should name be just "Asahi Restart Helper" instead? Does the Asahi project want a more general "helper"?
    - note that macOS version of this application
        can offer much more functionality than the Linux version for technical reasons

- [ ] 3: Handle not having any other viable disks ("macOS")
    - why would anyone run the app with just the one OS? Removing the macOS disk is not supported by the Asahi project
    - explain that no other disk were detected
    - link to relevant help page?

- [ ] 3: Consider showing a user-friendly error if using x86_64
    - reason why not: why on earth would anyone try to run the application on an x86_64 machine?
    (Asahi Linux requires Apple Silicon)

- [ ] 3: Show something (Open the menu? Open preferences?) if user launches manually while app was already running
    - reason: maybe user is confused and doesn't notice the menubar icon?
    - reason why not: corner case, adds code complexity

- [ ] 3: Auto updater
    - use someone else's solution as reference?
    - if super simple, DIY?

- [ ] 3: Handle other errors (what? where?)

- [ ] 3: Wording/strings ok?
    - "Restart in macOS" is the wording Apple's x86_64 Boot Camp uses in the app it adds on Windows
    - I don't want to use "to" or "reboot" (be macOS-user-friendly, use familiar terminology)
    - what is the correct amount of detail in text and tooltips?
    - suggestions for wording/spelling (titles, tooltips)?

- [ ] 3: Write documentation
    - proper README.md
    - is anything else necessary?

- [ ] 3: Write tests
    - Unit tests
        - might need to sometimes use weird custom mock tests because the app writes to NVRAM, GitHub won't test that 
    - UI tests


----------------------------------------

## 4: WON'T HAVE (not worth it / out of scope / feature creep, revisit this list periodically):

- [ ] 4: CLI installer: not needed? Can Asahi installer or Fedora install the app automatically?

- [ ] 4: No need to ask for permissions from the system to restart without password?

- [ ] 4: Uninstaller not necessary?
    - will eventually be a normal package that can be uninstalled?
    - Any settings left behind will be harmless?- (don't leave privileged helper app and other cruft behind)?
    - remove from Login Items
    - clear all prefs

- [ ] 4: Warn first-time users about permission prompts (no permission prompts needed?)

- [ ] 4: If prefs not enabled, try to convince user to enable pref?
    - show button/checkbox in the main "Restart dialog"?
        - don't be deceptive or potentially destructive, users might thinks it's the "normal restart toggle"
          (users don't read text on screen)
    - reason why not: annoying + adds complexity, use an installer or write to prefs from an external app instead?
    - reason: most users would want both checkboxes to be enabled, i.e. password-free + show icon on macOS login?

- [ ] 4: Detect and indicate "real" Asahi Linux disks vs "real" macOS disks (e.g. detect step2)
    - reason why not: fickle and hacky, Asahi Linux installation implementation details may change in the future
    - reason why not: what would be the user benefit?

- [ ] 4: Different UX if user chooses to restart to the current default Startup Disk (i.e. sudo bless not necessary)
    - reason why not: why not just sudo bless anyway, realistically no harm will be caused by always writing to NVRAM
        (even though technically the flash chip NVRAM is using can wear down if written to infinitely many times) 
    - reason why not: not useful to annoy the user just to tell them information they already know?
        (completely reasonable to want to always Restart the same way, no matter what current default Startup Disk is)

- [ ] 4: (HOPEFULLY NOT NECESSARY) Some kind of first-use-wizard?
    - reason why not: UX should be good enough to not require wizard
    - reason why not: hopefully this can be achieved by settings prefs by external apps or preferences dialog
    - reason why not: even an installer is better than a wizard?

----------------------------------------

# Seek external feedback:

- [ ] Is there something fundamentally wrong with this approach, or is using sudo asahi-bless --next inevitable?
- [ ] Code review?
- [ ] Name: see macOS version of TODO
- [ ] Menubar icon? Can I use someone else's work?
- [ ] App icon? Can I use someone else's work?
- [ ] Is it allowed for an external application to explicitly ask for and "transfer" user's consent
    - [ ] To add *this* app to autostart?
    - [ ] OK to "transfer" consent via prefs or command line options for installer?
- [ ] UI text proofreading by a native speaker, preferably someone with the relevant technical knowledge
- [ ] UX testing by both beginners and experts? Suggestions?

----------------------------------------

## Manual tests (test before releasing a new version and after major OS updates?)

- [ ] Can some of these somehow be tested with Unit Tests / UI Tests?
- [ ] Does the app actually successfully do what it promises
    - Able to restart to any OS (all macOS and Linux disks)
    - Able to change the default Startup Disk
    - Other "Asahi Linux" tasks?
- [ ] Do the menus look wonky? What if the icon is moved to the right etc?
- [ ] Do the macOS/Linux versions of this app look and feel similar enough?
- [ ] Is cancelling before restarting handled gracefully?
- [ ] What does a brand new user (never used the app before) see? Is everything as simple and as automatic as possible?
- [ ] Does the app ask for OS level permissions? Can fewer permission be used somehow?
- [ ] What happens if the default Startup Disk is already Linux/macOS and user tries to use the app to restart to that?
    - Is the UX different? Should it be? (No? Just always have consistent UX?)
- [ ] What happens if there are no "macOS" disks available?
- [ ] What happens if the user cancels an active restarting process? Is NVRAM now different than before?
- [ ] Does everything look fine in both KDE Plasma and GNOME?
- [ ] Does making the app start at login work?
- [ ] Change all preferences in Options and restart the app, does reading/writing prefs work correctly?
- [ ] Does the app specifically ask user to make the app start at login if unset? Should it?
- [ ] Does the app automatically make reality agree with what value "auto start at login" pref has?
- [ ] Is behavior acceptable compared to normal Linux restart dialog / workflow? KDE vs GNOME?
- [ ] Has something changed in KDE / GNOME? Does something need to be changed in this app to compensate?
- [ ] Has a relevant Asahi Linux related thing changed? Does something need to be changed in this app to compensate?
- [ ] README and TODO: are macOS/Linux versions roughly similar, i.e. have they been kept in lockstep?
    - I sure hope so, must be annoying to fix it not
- [ ] The actual point of this project: Does using the Asahi Linux installer automatically add a "Restart in macOS"
    tray icon in a useful state (works without asking silly questions & autoloads)?

----------------------------------------

(move tasks here after completing them)
## DONE:
- [x] Original extremely nice app by nohajc (modified and renamed with explicit permission)
