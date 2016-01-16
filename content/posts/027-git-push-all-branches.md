+++
date = "2013-09-10"
title = "git push all branches"
slug = "git-push-all-branches"
+++

If you use git branches a lot, you'll often push each branch after each commit.

This takes time...  
<!--more-->
Instead of pushing every single branch you can do `git push --all origin`.  
This will push all commits of all branches to origin. That really simple!

### Pushing all branches to default remote

Maybe you're even renaming your git remotes like I do.  
`git remote rename origin github` and you will have a remote called *github*.  
Now you would have to push all commits of all branches with `git push --all github`.

To simplify that aswell you can run `git push --all github -u` once and now all you'll have to do is `git push`. This will now by default push all branches to the default remote *github*.

Checkout [http://stackoverflow.com/questions/1914579/set-up-git-to-pull-and-push-all-branches](http://stackoverflow.com/questions/1914579/set-up-git-to-pull-and-push-all-branches) for even more information.

Pretty fancy, right?
