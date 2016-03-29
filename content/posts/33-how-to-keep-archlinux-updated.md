+++
date = "2016-03-27T16:21:38+01:00"
title = "How to keep archlinux updated"
slug = "how-to-keep-archlinux-updated"
cover = "posts/33/archey.png"
+++

This post is for beginners that want to get started with pacman.  
It's only a summary.

You can find a complete guide for pacman in the 
[archlinux.org wiki](https://wiki.archlinux.org/index.php/Pacman#Usage)
[[de]](https://wiki.archlinux.de/title/Pacman#Anwendung).  

<!--more-->

## Install, uninstall and search for packages


#### Install or update packages:
```bash
sudo pacman -S [package1] [package2]
```

#### Uninstall packages:
```bash
sudo pacman -Rns [package1] [package2]
```
*One or more packages are uninstalled with all dependencies (if those are no longer used by other programs).*

#### Search for packages
```bash
pacman -Ss [package]
```
*Searches for installable packages. It's enough to provide pieces of the package's name.*

---

Search for already installed packages:
```bash
pacman -Qs [package]
```
*This is perfect to check if you've already installed a package.*

To get a complete list of all explicit installed packages use `pacman -Qet` and with all dependencies use `pacman -Qqe`.  
This could be useful if you e.g. want to install most of the packages on another system.

## Update your system

Upgrade your entire system.
```bash
sudo pacman -Syu
```
*Do this daily to weekly, to stay up to date.*

---

Delete the local database and download it again:
```bash
sudo pacman -Syy
```
*You shouldn't need to run this often, if you stay updated.  
If pacman can't find a package you want to install, run this.*

#### Run the following every week or so

Delete unsued and obsolete packages from `/var/cache/pacman/pkg` where updates are downloaded to:
```bash
sudo pacman -Sc
```
This will free some space on `/var`.

---

Optimize pacman's database to work faster:
```bash
sudo pacman-optimize
```
*This will enable pacman to faster resolve dependencies of packages.*

---

Delete all orphaned packages:
```bash
sudo pacman -Rsn `pacman -Qdtq`
```
*A package is orphaned if no other package requires it and it wasn't explicitly installed.*

**Be careful with this, it could delete packages that you want to keep. Double check!**

---

Please comment below if you find any errors, typos or wrong tricks.
