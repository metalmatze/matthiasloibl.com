+++
date = "2013-11-21"
title = "Deploying a website with git on uberspace"
slug = "deploying-a-website-with-git-on-uberspace"
+++

Are you using git yet? Are you deploying your website via `git push`? - You **should**!

The goal of this post is to give you a short overview of my way to get git automatically deploy a website via `git push` on [uberspace.de](https://uberspace.de/).

> I will use [*metalmatze.de*](http://metalmatze.de) as domain and *metalmatze* as examples (my site is not hosted on uberspace).
<!--more-->
---

### What we want to do
1. `git push uberspace master` - This will push new commits to uberspace and deploy your site.
2. Check [metalmatze.de](http://metalmatze.de) to see the changes we just pushed.

---

### Installation

Connect with ssh to uberspace.

    $ ssh metalmatze@pictor.uberspace.de

Create a new folder in your home directory (you should be there after connecting) called *metalmatze.de.git*

    $ mkdir metalmatze.de.git

Change into this directory with `cd metalmatze.de.git`.

Next we want to create a bare git repository
    
    $ git init --bare
    Initialized empty Git repository in /home/metalmatze/metalmatze.de.git/
    
**The following steps will take place in your *local* repository!**

Now we want to add uberspace as a remote repository to our local repository

    $ git remote add uberspace ssh://metalmatze@pictor.uberspace.de/home/metalmatze/metalmatze.de.git

**The following steps will take place on *uberspace* in the bare git repository!**

To get git to deploy our website into a specified folder we'll use a [post-receive hook](https://www.kernel.org/pub/software/scm/git/docs/githooks.html#post-receive).

Open `hooks/post-receive` with an editor like *vim* or *nano*.
    
    $ vim hooks/post-receive
    
Into this file we'll paste a simple bash script that will deploy the master branch into `/home/metalmatze/html/metalmatze.de`

    #!/bin/bash
    path="/home/metalmatze/html/metalmatze.de/"

    echo "========= GIT CHECKOUT ========="
    GIT_WORK_TREE=$path git checkout -f master

    echo "============= DONE ============="

Save and exit the editor.

Don't forget to change the file permissions to be executable.

    $ chmod u+x hooks/post-receive

At last we need to link [metalmatze.de](http://metalmatze.de) to our folder where we deployed metalmatze.de's into.

    $ cd /var/www/virtual/metalmatze

    $ ln -s html/metalmatze.de/public/ metalmatze.de

Now go on [metalmatze.de](http://metalmatze.de) with your browser and check out your automatically deployed website! 

Isn't that awesome?! 


### Benefits

We are now able to deploy our changes with only one git command, or even with a GUI.

---

Do you have any questions or comments? Please leave them below. Thanks! 
