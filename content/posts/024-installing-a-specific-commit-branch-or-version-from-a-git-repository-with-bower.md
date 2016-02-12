+++
date = "2013-08-27"
title = "Installing a specific commit, branch or version from a git repository with bower"
slug = "installing-a-specific-commit-branch-or-version-from-a-git-repository-with-bower"
+++

**Installing packages via bower is awesome.**

It's even more awesome that you're able to install git repositories with bower.  
This will become handy if you have private repositories.

<!--more-->
---

Installing git repositories with bower:

    bower install --save https://github.com/jquery/jquery.git

This will clone the jquery git repository from github to your local machine and then checkout the master branch to your bower folder.

It could only become a problem in the future, since you would get all the changes from the master branch.

Specifing a commit, branch or version of your git repository is thankfully not that hard.

To install a specific **commit**:

    bower install --save https://github.com/jquery/jquery.git#42cd19fb8f29db6fb06f49e5b94829ea7036c3c5

To install a specific **branch**:

    bower install --save https://github.com/jquery/jquery.git#develop

To install a specific **version**:

    bower install --save https://github.com/jquery/jquery.git#1.10.2

Hopefully you're now more safe to updates in the future.
