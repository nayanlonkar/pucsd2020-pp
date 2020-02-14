
build:
	gcc src/main.c src/operations/add.c src/operations/sub.c src/operations/mult.c src/operations/div.c -o calculator

build-static:
	gcc src/main.c src/libfunctions.a -o calculator

run:
	./calculator


clean:
	rm calculator
