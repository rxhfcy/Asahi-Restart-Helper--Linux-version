asahi-reboot-switcher: main.go
	go build -tags=legacy_appindicator

install: asahi-reboot-switcher
	install -m 755 asahi-reboot-switcher /usr/bin/asahi-reboot-switcher
	install -m 644 asahi-reboot-switcher.policy /usr/share/polkit-1/actions/asahi-reboot-switcher.policy
