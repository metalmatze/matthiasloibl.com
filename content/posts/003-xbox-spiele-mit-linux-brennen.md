+++
date = "2010-10-02"
title = "Xbox Spiele mit Linux brennen"
slug = "xbox-spiele-mit-linux-brennen"
+++

Gestern habe ich mir Call of Duty Modern Warfare 2 für die Xbox gekauft.

Da ich allerdings weiß wie ich mit DVDs umgehe, wollte ich das Spiel sichern.
Schnell ein Abbild der DVD in Form einer Iso erstellt.

Weitaus schwieriger gestaltet sich jedoch das Brennen des Spieles wieder auf eine DoubleLayer DVD. Doch ich wurde fündig.

Es ist zwar kein grafisches Programm aber immerhin ein Befehl für die Bash. Also das Kommando kopiert und aus probiert:

<pre><code>growisofs -use-the-force-luke=dao -use-the-force-luke=break:1913760 -dvd-compat -Z /dev/scd0=CODMW2.iso -speed=2</code></pre>  
Anpassen muss man lediglich die Option von -Z.
/dev/scd0 war bei mir das DVD-Laufwerk und `dnl-codmw2.iso` ist das Abbild.

Dann kann man das Spiel wieder auf eine DVD brennen und es spielen.
