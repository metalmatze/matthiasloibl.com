+++
date = "2016-03-27T16:21:38+01:00"
title = "How to keep archlinux updated"
slug = "how-to-keep-archlinux-updated"
+++

You can find a complete overview of pacman in the 
[archlinux.org wiki](https://wiki.archlinux.org/index.php/Pacman#Usage)
[[de]](https://wiki.archlinux.de/title/Pacman#Anwendung).  
This post aims to summarize the important action to run with pacman to keep your system up to date.


## Install, uninstall and search for packages


#### Install or update packages:
```bash
sudo pacman -S [package1] [package2]
```

#### Uninstall packages:
```bash
sudo pacman -Rs [package1] [package2]
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

Delete all orphaned packages:
```bash
sudo pacman -Rsn `pacman -Qdtq`
```
*A package is orphaned if no other package requires it and it wasn't explicitly installed.*

---

Optimize pacman's database to work faster:
```bash
sudo pacman-optimize
```
*This will enable pacman to faster resolve dependencies of packages.*
