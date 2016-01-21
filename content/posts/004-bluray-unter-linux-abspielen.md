+++
date = "2010-12-26"
title = "BluRay unter Linux abspielen"
slug = "bluray-unter-linux-abspielen"
+++

Heute möchte ich eine Möglichkeit zeigen, mit der man unter Linux aber auch unter Windows und Mac BluRays mehr oder weniger einfach angucken kann.

Dazu ist kein Zwischenspeichern oder langwieriges Konvertieren als Matroska Video (.mkv) nötig. Lediglich die BluRay ins BluRay-Laufwerk legen und mit ein paar Klicks gemütlich abspielen!

<img src="/posts/4/vlc-makemkv.png" class="picleft" />VLC ist der Video Player mit dem das Video später angeguckt wird, dazu wird der VLC Player mit Matroska Unterstützung benötigt, was in den meisten Fällen (Ubuntu, Windows und Mac) schon der Fall ist. Also einfach VLC Player installieren. Einen Download findet man unter <a href="http://www.videolan.org">http://www.videolan.org</a>.

Das zweite Programm ist MakeMKV. Unter <a href="http://www.makemkv.com/download">http://www.makemkv.com/download</a> kann man es herunterladen. Die Installation sollte unter Windows und Mac selbsterklärend und nicht all zu schwer sein!
Für die Linux Benutzer unter euch gibt es ein Download im <a href="http://www.makemkv.com/forum2/viewforum.php?f=3">Forum</a>

Speichert aus diesem Forum die beiden Dateien: <strong>makemkv_v[...]_bin.tar.gz</strong> und <strong>makemkv_v[...]_oss.tar.gz</strong>

Nun entpacken: <div class="code">tar -xf makemkv_v*_bin.tar.gz
tar -xf makemkv_v*_oss.tar.gz</div>
Nun in den ersten Ordner wechseln und kompilieren:<div class="code">cd makemkv_v*_bin
make -f makefile.linux
sudo make -f makefile.linux install</div>
Als nächstes in den zweiten Ordner und auch dies kompilieren und installieren:<div class="code">(cd ..)
cd makemkv_v*_oss
make -f makefile.linux
sudo make -f makefile.linux install</div>

<img src="/posts/4/makemkv.png" class="picleft" />Wenn all dies erledigt ist, sollte man nun MakeMKV öffnen können. Einfach makemkv in die Konsole oder sonst wo rein! :D
Es baut sich nun ein Fenster wie links auf. Zum Abspielen muss die BluRay nur geöffnet und einmal analysiert werden. Dies erreicht man über File > Open Disc oder ganz einfach in dem man auf die Abbildung mit der BluRay zur Festplatte klickt.
Die Kapitel der BluRay werden ausgelesen und angezeigt.

<img src="/posts/4/makemkv2.png" class="picright"/>Auf der linken Seite gibt es eine Übersicht mit allen Kapiteln der BluRay und kann diese dort zum Konvertieren/Streamen einzeln aktivieren bzw. deaktivieren. Ganz rechts gibt es einen Knopf unter "Make MKV" mit dessen Hilfe der ausgewählte Inhalt als .mkv auf die Festplatte kopiert wird. Allerdings 1:1 wodurch die Dateien ziemlich groß sind und später eventuell noch kleiner konvertiert werden sollten.

Aber eigentlich möchten wir den Film ja ganz normal angucken. Dazu müssen wir die BluRay als Stream zu Verfügung stellen. File > Stream oder den Knopf mit dem Computer und Satelliten in der oberen Toolbar klicken. Nun gibt es einen Netzwerkstream unter http://localhost:51000. Diesen Stream mit dem Browser aufrufen und das gewünschte Kapitel auswählen. Meistens ist der eigentliche Film title0.

Da man nun den Link zum Film hat wird dieser mit dem VLC Player aufgerufen.
Dazu im VLC Player "Medien > Netzwerkstream öffnen" oder einfach Strg+N und den Link einfügen: http://localhost:51000/stream/title0.m2ts

Viel Spaß beim BluRay gucken!!! :)

Zuletzt möchte ich noch auf einen Blog verweisen, der mir diese Möglichkeit gezeigt hat! Danke! <a href="http://linuxsolver.blogspot.com/2010/09/bluray-on-linux-no-rip-makemkv-and-vlc.html">http://linuxsolver.blogspot.com/2010/09/bluray-on-linux-no-rip-makemkv-and-vlc.html</a>
