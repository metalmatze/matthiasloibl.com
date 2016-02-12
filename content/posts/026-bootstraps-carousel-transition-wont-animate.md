+++
date = "2013-09-05"
title = "Bootstrap's carousel transition won't animate"
slug = "bootstraps-carousel-transition-wont-animate"
+++

Twitter Bootstrap's carousel seems like a easy way to create a simple [carousel animation](http://getbootstrap.com/javascript/#carousel) on your website.

<!--more-->
---

![Bootstrap 3 Carousel](/posts/26/carousel.png)

Because I'm using [bower](http://bower.io) and [grunt](http://gruntjs.com) I decided to manually import bootstrap's javascript.

I started by concatenating `bootstrap/js/carousel.js` to my remaining javascript.  
Simple as that it worked after adding the css and creating the HTML.

Sadly to say the **sliding transition didn't work**.

I spend almost an hour trying to figure out what was wrong in my HTML.  
There wasn't anything wrong.  
The CSS was imported from bootstrap so it shouldn't have been the problem.

### Solution

While taking a look at all the available javascript files from bootstrap I saw that  
**there is a file called `bootstrap/js/transition.js`.**

After concatenating this `bootstrap/js/transition.js` aswell as `bootstrap/js/carousel.js` to my javascript files, the carousel worked with the expected sliding animation.
