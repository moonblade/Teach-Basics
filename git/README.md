# Git


## 1 - fork workflow - commit

1. Fork this repo to your own gitlab account
2. Set up an ssh key for your own account and check it out
3. Create a new commit on your version of the repo with a line saying task1 done on this readme
4. push the commit and create a merge request to the original repo.
5. Update your code, and instead of saying task 1 done, it should say task1 done with the link to the merge request 
6. wait for the code to merge
7. task 1 done https://gitlab.com/moonblade/jw-basics/-/merge_requests/1

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



