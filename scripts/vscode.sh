#!/usr/bin/env bash

# Install VSCode

script_dir="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $script_dir
. set.sh

wget https://github.com/VSCodium/vscodium/releases/download/1.40.2/codium_1.40.2-1574798581_amd64.deb
dpkg -i codium_1.40.2-1574798581_amd64.deb

ln -s /usr/bin/codium /usr/local/bin/code

# Set-up all users
for DIR in $(ls -1d /home/* 2>/dev/null ||:)
do
	# Disable most of the telemetry and auto-updates
	mkdir -p $DIR/.config/VSCodium/User
	cat <<-'EOF'> $DIR/.config/VSCodium/User/settings.json
	{
	    "telemetry.enableCrashReporter": false,
	    "update.enableWindowsBackgroundUpdates": false,
	    "update.mode": "none",
	    "update.showReleaseNotes": false,
	    "extensions.autoCheckUpdates": false,
	    "extensions.autoUpdate": false,
	    "workbench.enableExperiments": false,
	    "workbench.settings.enableNaturalLanguageSearch": false,
	    "npm.fetchOnlinePackageInfo": false
	}
	EOF

	# Fix rights
	USR=$(echo "$DIR" | rev | cut -d/ -f1 | rev)
	chown -R $USR:$USR $DIR ||:
done
