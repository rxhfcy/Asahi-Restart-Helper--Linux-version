asahi-reboot-switcher: main.go asahibless/* util/*
	go build

install: asahi-reboot-switcher
	install -m 755 asahi-reboot-switcher /usr/bin/asahi-reboot-switcher
	install -m 644 asahi-reboot-switcher.policy /usr/share/polkit-1/actions/asahi-reboot-switcher.policy
	install -m 644 asahi-reboot-switcher.desktop /usr/share/applications/asahi-reboot-switcher.desktop
	install -m 644 asahi-reboot-switcher.png /usr/share/icons/asahi-reboot-switcher.png
