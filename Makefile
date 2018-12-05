CHS = CH1 CH3
LANG = c go

all:
	for ch in $(CHS); do \
		for prog in $$ch/**; do \
			echo "building $$prog:"; \
			$(MAKE) -C $$prog; \
		done; \
	done

clean:
	for ch in $(CHS); do \
		for prog in $$ch/**; do \
			echo "cleaning $$prog:"; \
			$(MAKE) clean -C $$prog; \
		done; \
	done
	