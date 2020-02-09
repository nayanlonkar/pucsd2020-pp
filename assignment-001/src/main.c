#include "../include/functions.h"
#include <stdio.h>
#include <string.h>
#define MAX 500
#define PROMPT ">>> "

int main() {
	
	/* char expression[100]; */
	/* printf("Enter a expression : "); */
	/* scanf("%s", expression); */
	/* printf("\n"); */
	/* printf("%s", expression); */

	/* int a = 10, b = 5; */

	/* printf("addition is: %f\n", add(a, b)); */
	/* printf("subtraction is: %f\n", sub(a, b)); */
	/* printf("multiplication is: %f\n", mult(a, b)); */
	/* printf("division is: %f\n", div(a, b)); */
	/* printf("percentage is: %f\n", percentage(a, b)); */
	
	char expression[MAX];
	printf("Enter expression else press q to exit\n");

	while(1) {
		printf(PROMPT);

		/* fgets(expression, MAX, stdin); */
		scanf("%[^\n]%*c",expression);
		
		if (strcmp(expression,"q") == 0)
			break;
	}

	return 0;
}
