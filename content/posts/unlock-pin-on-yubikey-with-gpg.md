+++
cover = ""
date = ""
draft = true
slug = ""
title = "Unlock PIN on YubiKey with gpg"

+++
It just happened, that I was doing something completely else and didn't pay attention when entering my YubiKey PIN incorrectly 3 times...

In the end I was stuck with a not working YubiKey gpg agent. That's when I started searching on how to easily unlock it with the admin key (called PUK) and couldn't find any easy guides.

If you need to unlock gpg on your YubiKey it is actually a gpg feature and not a YubiKey feature itself.

Therefore use `gpg --edit` to go into the settings of gpg. Once you're in there you could use `passwd` to change your PIN / password, if you weren't locked out.
Thus, enter the admin mode by entering `admin` and entering your admin password, also called PUK, and then then you can change the Pin with (some command). You need to enter a new PIN, which can be the PIN you used before. Maybe it's time to rotate it though?!

Exit both shells with Ctrl+D and you're ready to use your YubiKey with gpg just like before. 😊

Hey Kemal