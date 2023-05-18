name_web=picow-rgbw-web
build_path=./build
bin_path=~/.local/bin
config_path=~/.config/
systemd_path=~/.config/systemd/user

.PHONY: build

build:
	@cd .. && (\
		git clone --branch=main https://github.com/knackwurstking/svelteui.git || ( \
			cd svelteui && git pull \
		) \
	)
	@cd frontend && npm install && npm run build
	@go mod tidy
	@echo "build ${name_web} -> ${build_path}/${name_web}"
	@go build -v -o ${build_path}/${name_web} ./cmd/${name_web}

install:
	@systemctl --user stop picow-rgbw-web || exit 0
	@mkdir -p ${bin_path} && cp ${build_path}/${name_web} ${bin_path}/${name_web}
	@if [ ! -e ${config_path}/${name_web}/config.json ]; then \
		mkdir -p ${config_path}/${name_web}; \
		cp ./cmd/picow-rgbw-web/config/example.config.json ${config_path}/${name_web}/config.json; \
	fi
	@if [ ! -e ${systemd_path}/${name_web}.service ]; \
		then mkdir -p ${systemd_path}; \
		cp ./cmd/picow-rgbw-web/config/${name_web}.service ${systemd_path}/${name_web}.service; \
	fi

service:
	systemctl --user daemon-reload
	systemctl --user restart picow-rgbw-web
	systemctl --user enable picow-rgbw-web
	journalctl --user --follow --output cat -u picow-rgbw-web
