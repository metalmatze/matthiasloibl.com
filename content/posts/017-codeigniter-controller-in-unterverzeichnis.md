+++
date = "2012-07-13"
title = "CodeIgniter Controller in Unterverzeichnis"
slug = "codeigniter-controller-in-unterverzeichnis"
+++


[CodeIgniter](http://codeigniter.com), ein quelloffenes PHP-Web-Framework, mein lieblings PHP Framework. Neben den ganzen nützlichen Funktionen gibt es eine wirklich gute Sache, die ich noch nicht all zu lange weiß.

Es ist möglich Controller in Unterverzeichnisse zu sortieren.  
Das trägt wesentlich zu übersichtlichen Projekten auf Quellcode Ebene bei.

Angenommen wir möchten einen Login Bereich erstellen:

1.  /application/controllers/login.php wird erstellt
2.  /application/controllers/login.php wird mit sämtlichen gewünschten Funktionen voll geschrieben
3.  Der Controller wird schnell sehr sehr unübersichtlich, ich würde sagen ab 5 Funktionen aufwärts
4.  Kein Überblick, zu viel Code. CHAOS

Stattdessen kann man aber auch Sub-Controller verwenden:

1.  /application/controllers/login/ Verzeichnis erstellen
2.  /application/controllers/login/user.php erstellen
3.  /application/controllers/login/pics.php erstellen
4.  /application/controllers/login/some.php erstellen

Nun hat man für den Login Bereich anstatt einem Controller viele Einzelne.

[Link zum UserGuide](http://codeigniter.com/user_guide/general/controllers.html#subfolders)
