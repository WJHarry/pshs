# pshs
Easy to search and apply history commands in Powershell (Windows).

The "history" command in PowerShell is session-level. pshs provides global-level history browsing and supports fuzzy search. Find your history and just "Enter" to execute it again.

This plugin is inspired by [hstr](https://github.com/dvorka/hstr).

![example](https://github.com/WJHarry/pshs/blob/b28d75314f7cbc184ca6e573e9d6f1b489c201cf/example.gif)


## Install
### From msi
1. Download and unzip the [pshs.zip](https://github.com/WJHarry/pshs/releases/latest)
2. double click the "pshs.msi", it will add pshs to your environment variables automatically.

### From source
```bash
git clone https://github.com/WJHarry/pshs.git pshs
cd pshs
make install
```

## Usage
Ctrl-R: switch Normal mode and Regexp mode
Down/Ctrl-N: Next
Up/Ctrl-P: Previous
Ctrl-C: Quit
Enter: Select and execute
