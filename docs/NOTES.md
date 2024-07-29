# Some notes to Escape from Babylon

## Well known paths (Windows)

* Python official install path for current user `%APPDATA\Local\Programs\Python\PythonXX`
* NPM global current user path: `%APPDATA%\Roaming\npm\node_modules\npm\bin`
* Go bin folder: `C:\Users\YourUsername\go\bin\go.exe`
* Rust: `C:\Users\YourUsername\.cargo\bin\rustc.exe`
* Haskel: `C:\Users\YourUsername\AppData\Roaming\local\bin\ghc.exe`
* FireFox: `C:\Users\<username>\AppData\Local\Mozilla Firefox\firefox.exe`
* Chrome: `C:\Users\<username>\AppData\Local\Google\Chrome\Application`
	* `chrome.exe`: The main executable for launching Google Chrome.
	* `chrome_proxy.exe`: A process used for managing proxy settings in Chrome.
	* `chrome_launcher.exe`: Typically used to start the Chrome browser with specific configurations.
	* `chrome.dll`: While not an .exe, chrome.dll is a crucial dynamic link library file used by Chrome. (For context, it is located in the same directory or subdirectories, but it’s not an executable file.)
	* `chrome_remote_desktop_host.exe`: If Chrome Remote Desktop is installed, this executable handles remote desktop connections.
	* `chrome_update.exe`: An executable for updating Chrome.

* Edge extensions: `C:\Users\<YourUsername>\AppData\Local\Microsoft\Edge\User Data\Default\Extensions`
* Opera: `C:\Users\<YourUsername>\AppData\Roaming\Opera Software\Opera Stable\Extensions`
* Firefox profile extensions: `C:\Users\<YourUsername>\AppData\Roaming\Mozilla\Firefox\Profiles\<ProfileName>\extensions`
* Chrome extensions and components: `C:\Users\<YourUsername>\AppData\Local\Google\Chrome\User Data\Default\Extensions`

Check if running under wine by testing if executables are present:

* `.wine/drive_c/windows/syswow64/wine*.exe`
* `.wine/drive_c/windows/system32/wine*.exe`

## Ultimate Packer for Executables (UPX)

* <https://www.ired.team/offensive-security/defense-evasion/t1045-software-packing-upx>
* <https://medium.com/@ankyrockstar26/unpacking-a-upx-malware-dca2cdd1a8de>
* <https://www.mosse-security.com/2020/09/29/upx-malware-evasion-technique.html?ref=nishtahir.com>
* <https://www.esecurityplanet.com/threats/upx-compression-detection-evasion/>

## Persistence and hiding

* Search for existing well known binary paths
* Copy argv[0] to well known binary path
* Register startup by system
  * schtasks (cmd) for system or local user
  * go-autostart: shortcut in start-menu
* Write state file of persistence to somewhere...

## Debugging release build

* The "VMK" environment variable is the VerboseModeKey which enables logging to stdout/stderr even in release build

## OS and emulator/environment detector

* Linux
  * Native
  * Msys (Windows)
  * CYGWIN (Windows)
  * [Microsoft WSL & WSLv2](https://github.com/microsoft/WSL/issues/4071)
  * [FreeBSD linuxemu](https://docs.freebsd.org/en/books/handbook/linuxemu/)
* Windows
  * WINE
  * ReactOS
  * Native
* Darwin (macOS)

## Windows

* Copy to well known current user binary path to semi related filenames
* Run via start menu item for current user, or via `schtasks`
  * <https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/schtasks-create>
  * <https://github.com/emersion/go-autostart>

## Detection

* <https://www.logpoint.com/en/blog/deep-dive-on-malicious-dlls/>
