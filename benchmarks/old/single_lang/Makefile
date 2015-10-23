all: go_only c_only

go_only:
	go build go_only.go

c_only: c_only.c
	gcc -O3 -o $@ $<

clean:
	rm -f go_only c_only

