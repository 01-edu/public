#!/usr/bin/env bash

# Install VSCode

script_dir="$(cd -P "$(dirname "$BASH_SOURCE")" && pwd)"
cd $script_dir
. set.sh

wget https://github.com/VSCodium/vscodium/releases/download/1.41.1/codium_1.41.1-1576787344_amd64.deb
dpkg -i codium_1.40.2-1574798581_amd64.deb

ln -s /usr/bin/codium /usr/local/bin/code

# Set-up all users
for dir in $(ls -1d /home/* 2>/dev/null ||:)
do
	# Disable most of the telemetry and auto-updates
	mkdir -p $dir/.config/VSCodium/User
	cat <<-'EOF'> $dir/.config/VSCodium/User/settings.json
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
	usr=$(echo "$dir" | rev | cut -d/ -f1 | rev)
	chown -R $usr:$usr $dir ||:
done
