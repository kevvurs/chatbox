build-primer:
	$(MAKE) -C primer build

build-chatbox:
	$(MAKE) -C chatbox build

build: build-primer build-chatbox
	mkdir -p dist/
	mv primer/main.wasm dist/main.wasm
	mv chatbox/chat.wasm dist/chat.wasm
	cp primer/wasm_exec.js dist/wasm_exec.js
	cp index.html dist/index.html
