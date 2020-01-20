#include <stdio.h>
#include "../include/functions.h"

int main() {
	
	/* char expression[100]; */
	/* printf("Enter a expression : "); */
	/* scanf("%s", expression); */
	/* printf("\n"); */
	/* printf("%s", expression); */

	int a = 10, b = 5;

	printf("addition is: %f\n", add(a, b));
	printf("subtraction is: %f\n", sub(a, b));
	printf("multiplication is: %f\n", mult(a, b));
	printf("division is: %f\n", div(a, b));
	printf("percentage is: %f\n", percentage(a, b));

	return 0;
}
