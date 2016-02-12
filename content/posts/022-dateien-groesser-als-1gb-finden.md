+++
date = "2012-08-22"
title = "Dateien größer als 1GB finden"
slug = "dateien-groesser-als-1gb-finden"
+++

Soeben ist eine log-Datei im Home-Verzeichnis extrem groß geworden, so groß, dass am Ende auf der Partition 0 Byte frei waren.

Um den Übeltäter zu finden habe ich nach einer Möglichkeit gegoogelt große Dateien aufzuspüren.

<!--more-->
---

Folgender Befehl findet alle Dateien größer als 1GB auf der /home Partition

    find /home -size +1G

Der Übeltäter wurde gefunden und gelöscht! :)
