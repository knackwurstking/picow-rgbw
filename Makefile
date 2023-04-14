name=picow-rgbw-web
bin_path=~/.local/bin
config_path=~/.config/
systemd_path=~/.config/systemd/user

build:
	@cd frontend && npm install && npm run build
	@go mod tidy
	@go build -o ./${name} ./cmd/${name}

install:
	@go build -o ${bin_path}/${name} ./cmd/${name}
	@mkdir -p ${bin_path} && cp ./${name} ${bin_path}/${name}
	@if [ ! -e ${config_path}/${name}/config.json ]; then \
		mkdir -p ${config_path}/${name}; \
		cp ./config/example.config.json ${config_path}/${name}/config.json; \
	fi
	@if [ ! -e ${systemd_path}/${name}.service ]; \
		then mkdir -p ${systemd_path}; \
		cp ./config/${name}.service ${systemd_path}/${name}.service; \
	fi

service_start:
	@systemctl --user daemon-reload
	@systemctl --user start picow-rgbw-web

service_stop:
	@systemctl --user daemon-reload
	@systemctl --user stop picow-rgbw-web

service_enable:
	@systemctl --user daemon-reload
	@systemctl --user enable picow-rgbw-web

service_disable:
	@systemctl --user daemon-reload
	@systemctl --user disable picow-rgbw-web

service_log:
	@journalctl --user --follow --output cat -u picow-rgbw-web
