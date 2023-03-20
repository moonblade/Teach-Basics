# Git


## 1 - fork workflow - commit

1. Fork this repo to your own gitlab account
2. Set up an ssh key for your own account and check it out
3. Create a new commit on your version of the repo with a line saying task1 done on this readme
4. push the commit and create a merge request to the original repo.
5. Update your code, and instead of saying task 1 done, it should say task1 done with the link to the merge request 
6. wait for the code to merge

## 2 - fork workflow - sync

1. now that you have already forked this repo, the original could change thats not present in your fork
2. Create another remote called upstream and add the link to the original repo
3. Pull in the changes from upstream and sync it with your fork
4. add a commit appending task two done here

## 3 - branch workflow 

1. get access to this repo directly with your username - by asking the admin
2. make a new branch with your changes (saying task 3 done)
3. push your new branch and create a merge request from that to main

## 4 - branch workflow - sync

1. reuse an old branch, but pull in new chagnes from main branch so that you can still create a merge request without any conflicts


## 5 - Conflict resolution

1. Pull the newest content in main 
2. Switch to a new branch called conflict
3. Change this line - add something here
4. Push to the branch conflict. See what happens

## 6 - Revert, abandon branch

1. While in main branch, Make a change in this line
2. commit it
3. Realize you committed it in main branch and that your main branch is now ruined
4. git checkout to an earlier commit, change the branch name of main to main_old, and checkout to new branch called main
5. checkout to branch to make your change in line one
6. commit it
7. Oh no, the change should have been in the last line
8. Revert the last change

## 7 - stash stuff you dont want

1. use stash to remove the clutter


## 8 - Dont add everything

1. 
```
for x in range {1..20}; do touch $x.ignore; done
```
2. add a line saying done
3. push it without pushing other files

## 9 - Git ignore files you need but not in git

1. Add the .ignore files to .gitignore, commit and push gitignore file

## 10 - Stash

1. Remove clutter using git stash
