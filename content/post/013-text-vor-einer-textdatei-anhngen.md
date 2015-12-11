+++
date = "2012-04-25"
title = "Text vor einer Textdatei anhängen"
slug = "text-vor-einer-textdatei-anhngen"
+++

Gerade habe ich nach einer Möglichkeit gesucht den Textinhalt einer Datei vor eine andere Textdatei einzufügen.  
Jetzt habe ich folgenden Code für mich benutzt:

<pre><code>echo|cat insert.txt|cat - text.txt > /tmp/out && mv /tmp/out text.txt</code></pre>

*insert.txt* - Die Textdatei mit dem Inhalt der vor die *text.txt* Datei eingefügt werden soll.  
*text.txt* - Textdatei mit schon vorhandenem Inhalt

Praktisch erweist sich dies für mich, da ich gerne eine mit **yui minimizer** minimierte Stylesheets Datei vorher mit einem Kommentar versehen möchte.  
Dieser Kommentar ist vonnöten für die style.css die ein Wordpress Theme enthält.

Dies ist eine abgewandelte Version von [http://www.cyberciti.biz/faq/bash-prepend-text-lines-to-file](http://www.cyberciti.biz/faq/bash-prepend-text-lines-to-file)
