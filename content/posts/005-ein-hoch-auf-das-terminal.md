+++
date = "2011-11-06"
title = "Ein Hoch auf das Terminal"
slug = "ein-hoch-auf-das-terminal"
+++

Möchte man eine große zip-Datei in viele einzelne kleine .rar-Dateien komprimieren gibt es einige Probleme. Einerseits ermöglicht es Ubuntu einem nur die zip zu entpacken und nicht erneut zu packen. Andererseits dauert es extrem lange 20GB mit den normalen Eigenschaften der GUI als .rar zu komprimieren.

Genau deshalb ein Hoch auf das Terminal.
Schnell mal im manual von rar gelesen und sofort einen kleinen Befehl zusammen geschraubt, um alles auf einen Schlag zu erledigen und das sogar noch sehr schnell.
`rar a -m0 -v921600 neues.rar alte.zip`

Jetzt werden 900MB Pakete erstellt die nicht komprimiert werden, wodurch alles sehr schnell abläuft, es eben aber nicht kleiner wird, was mir aber egal war.
