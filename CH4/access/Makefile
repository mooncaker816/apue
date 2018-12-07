ROOT = ../..
include $(ROOT)/OPTIONS

PROG = $(notdir $(CURDIR))

all: build_c build_go

build_c: $(PROG)
vpath %.c c
%:	%.c $(LIBAPUE)
	$(CC) $(CFLAGS) c/$@.c -o c/$@ $(LDFLAGS) $(LDLIBS)

build_go:
	go build -ldflags="-s -w" -o go/$(PROG) ./go

clean:
	# rm -f c/$(PROG) c/*.o go/$(PROG);
	rm -f `ls */* |grep -E -v ".go|.c"`;

