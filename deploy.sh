#!/bin/bash
hugo -b="/"
rsync -avzu --delete --progress -h -r public/* metalmatze.de:web/matthiasloibl.com
