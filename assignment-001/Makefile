
build:
	gcc src/main.c src/calculate.c src/operations/add.c src/operations/sub.c src/operations/mult.c src/operations/div.c -o calculator

build-static:
	gcc src/main.c src/calculate.c src/libfunctions.a -o calculator

build-shared:
	gcc src/main.c src/calculate.c src/libfunctions.so -o calculator

run:
	./calculator

clean:
	rm calculator
