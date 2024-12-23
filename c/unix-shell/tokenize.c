#include "tokens.h"
#include <stdio.h>
#include <stdlib.h>

// A demo program JUST to show that the tokenizer function works (used for autograder)
int main(int argc, char **argv) {
    char input[MAX_INPUT_LENGTH];
    fgets(input, MAX_INPUT_LENGTH, stdin);

    Token *tokens = tokenize(input);

    int i = 0;
    while (tokens[i].type != '\0') {
        printf("%s\n", tokens[i].value);
        i++;
    }

    free(tokens);
    return 0;
}
