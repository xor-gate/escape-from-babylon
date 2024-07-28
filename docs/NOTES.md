# Some notes to Escape from Babylon

## Well known paths (Windows)

* Python official install path for current user `%APPDATA\Local\Programs\Python\PythonXX`
* NPM global current user path: `%APPDATA%\Roaming\npm\node_modules\npm\bin`
* Go bin folder: `C:\Users\YourUsername\go\bin\go.exe`
* Rust: `C:\Users\YourUsername\.cargo\bin\rustc.exe`
* Haskel: `C:\Users\YourUsername\AppData\Roaming\local\bin\ghc.exe`
* FireFox: `C:\Users\<username>\AppData\Local\Mozilla Firefox\firefox.exe`
* Chrome: `C:\Users\<username>\AppData\Local\Google\Chrome\Application\chrome.exe`

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

## Windows

* Copy to well known current user binary path to semi related filenames
* Run via start menu item for current user, or via `schtasks`
  * <https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/schtasks-create>
  * <https://github.com/emersion/go-autostart>
