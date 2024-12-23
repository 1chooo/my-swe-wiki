#ifndef TOKENS_H
#define TOKENS_H

#define MAX_INPUT_LENGTH 1024

// Define a struct to hold tokens
typedef struct {
    char type; // 'W' for word, 'S' for special character
    char value[MAX_INPUT_LENGTH];
} Token;

// Function to check if a character is a special token
int is_special(char ch);

// Function to tokenize the input into a list of tokens
Token *tokenize(const char *input);

#endif