#include <stdio.h>
#include <stdlib.h>
#include <float.h>


struct Stack {
	int top; 		// stack top
	int capacity; 	// capacity of a stack
	float* array;  	// stack array
};

struct Stack* createStack(int capacity) {
	struct Stack* stack = (struct Stack*)malloc(sizeof(struct Stack));
	stack->array = (float*)malloc(sizeof(float)* capacity);
	stack->capacity = capacity;
	stack->top = -1;
	return stack;
}

int isEmpty(struct Stack* stack) {
	return (stack->top == -1);
}

int isFull(struct Stack* stack) {
	return (stack->top == stack->capacity - 1);
}

void push(struct Stack* stack, float data) {
	if (isFull(stack))
		return;
	else
		stack->array[++stack->top] = data;
}

float pop(struct Stack* stack) {
	if (isEmpty(stack))
		return FLT_MIN;
	return stack->array[stack->top--];
}

float peek(struct Stack* stack) {
	return (stack->array[stack->top]);
}


/* int main() { */
	/* struct Stack* stack = createStack(100); */
	/* return 0; */
/* } */








