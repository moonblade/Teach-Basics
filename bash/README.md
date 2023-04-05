# Bash

## 1 - Better tools

1. The terminal in mac sucks. Install iterm
2. If interested, use quake mode so that you're able to summon a terminal in any screen with a shortcut
3. Install oh my zsh on it as well (has some aliases and autocompletes that are handy)

## 2 - Common commands

1. Find out what ls, echo, touch, mkdir, cat, chmod do

## 2.5 - Bash background

1. Relative path and absolute path
2. . and ..
3. manual pages 
4. Optional and required arguments
5. Subcommands with $()

## 3 - Pipe

1. Find what pipe | is used for in linux
2. Use commands grep, less, more, head, tail
3. Escape character
4. Wildcards

## 4 - Redirection

1. Redirect output (stdout) of a command to a file
2. Redirect error stream (stderr) of a command to a file, but keep stdout in console itself
3. Redirect both stdout and stderr to same file
4. redirect stdout to out.log and error to err.log

## 5 - Cli arguments

1. Make a simple bash script that takes one argument and prints "hello \<argument\>"
2. Make it so that if no <argument> is provided, it takes a default value.
3. Read up on difference between default value expansions, [Docs](https://www.gnu.org/software/bash/manual/html_node/Shell-Parameter-Expansion.html#Shell-Parameter-Expansion)
4. Change it so that if no <argument> is provided, it gives usage instructions, similar to what happens if you run `mv` without arguments.
5. for the above commands, do variations with optional parameters like "command --optional-argument". [Help](https://unix.stackexchange.com/questions/331522/how-do-i-parse-optional-arguments-in-a-bash-script-if-no-order-is-given)

## 6 - More commands

1. Find out how to use find, sed, awk (print) commands 
2. [find](https://www.redhat.com/sysadmin/linux-find-command#:~:text=The%20find%20command%20is%20one,and%20not%20just%20by%20filename.)
3. [sed](https://phoenixnap.com/kb/linux-sed#:~:text=The%20Linux%20sed%20command%20is%20most%20commonly%20used%20for%20substituting,%2F%20\)%20for%20separating%20text%20fields.)
4. [awk](https://www.javatpoint.com/linux-awk-command#:~:text=The%20awk%20command%20is%20used,language%20used%20for%20text%20scripting.)

## 7 - Source, aliases and functions

1. Add to path so that your binaries work everywhere
2. Add aliases
3. Create a function mcd, that creates a directory and then cd's into it

<details>
<summary> sample </sample>

```bash
function book() {
  filename=${1}
  if [[ -z "$filename" ]]
  then
    echo "Usage: book [filename]"
    return 0
  fi
  # if [[ "${filename: -4}" == "epub" ]]
  # then
  #   echo "File $filename is epub, converting to mobi"
  #   /Applications/calibre.app/Contents/MacOS/ebook-convert $filename ${filename%epub}mobi
  #   filename=${filename%epub}mobi
  # fi
  if [[ "${filename: -4}" == "mobi" || "${filename: -4}" == "epub" ]]
  # if [[ "${filename: -4}" == "mobi" ]]
  then
    echo "Sending $filename to kindle"
    echo "book: $filename" | mutt -s "book: $filename" -a $filename -- email@kindle.com
  else
    echo "Unknown format"
  fi
}
```
<details>


## 8 - env variables

1. make a program that prints the env variable $MY_DUMMY_VAR
2. run it and give the env variable in the same line
3. export env variable then run it
4. export env variable and also give it in the same line
5. Export it in one terminal and run it in another terminal
6. Export it in bashrc so that its permanent, then run it in a new terminal

## 9 - Vim basics (optional)

1. navigate around, words, end of line, end of page, end of document
2. Make some changes, in current space and in next line
3. Save and quit
4. Reload the page and remove unsaved changes
5. Undo and redo
6. Search and replace
7. macros - save actions and replay them

## 10 - Basic scripting
1. Run 'make bash-scripting-1' 
2. remove every png file in 'bash/resources/cleanup' 
3. remove every csv file in 'bash/resources/cleanup' that start with the name file and are numbered between 30 and 60
4. remove every jpg file in 'bash/resources/cleanup' that have a file with the same name but with csv extention
5. For every pdf file, with the character f in its name, move it to the folder fpdf
6. for every txt file, if there it has the characters 00 inside it, move it to folder txt00

## 11 - Makefiles
1. Sometimes you want to save commands along with your repos, and complicated commands or sets of commands are too much to remember, or have no idea about when you come back to the repo a year later

