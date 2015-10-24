TARGETS=go_only c_only go_to_c_adds go_to_c_sum go_concurrent go_to_c_concurrent go_to_go_http

.PHONY: clean $(TARGETS)

all: $(TARGETS)

go_only:
	make -C $@

c_only:
	make -C $@

go_to_c_adds:
	make -C $@

go_to_c_sum:
	make -C $@

go_concurrent:
	make -C $@

go_to_c_concurrent:
	make -C $@

go_to_go_http:
	make -C $@

clean:
	rm -f go_only/main c_only/main go_to_c_adds/main go_to_c_sum/main \
		go_only/libharmonic.* go_concurrent/main go_to_c_concurrent/main \
		go_to_go_http/main
