#include "tokens.h"
#include <fcntl.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <unistd.h>

char prev_command[MAX_INPUT_LENGTH] = "";

ssize_t getline(char **lineptr, size_t *n, FILE *stream); // getline workaround for implicit decleration but still uses built in...

void cd_command(char *directory);
void help_command();
Token *extract_tokens(Token *original_tokens, int start, int end);
void execute_preexisting_command(Token *originalTokens, int start, int end);
void execute_recursive(Token *tokens, int start, int end);
int find_char(Token *tokens, int start, int end, char *target);
int count_tokens(Token *tokens);
void handle_seq(Token *tokens, int start, int end, int index);
void handle_pipe(Token *tokens, int start, int end, int index);
void handle_input(Token *tokens, int start, int end, int index);
void handle_output(Token *tokens, int start, int end, int index);
void handle_prev();

// This drives the interactive shell
int main(int argc, char **argv) {
    char input[MAX_INPUT_LENGTH];
    Token *tokens;

    printf("Welcome to mini-shell.\n");

    while (1) {
        printf("shell $ ");

        // Check if we reached the end of file (Ctrl-D)
        if (fgets(input, MAX_INPUT_LENGTH, stdin) == NULL) {
            if (feof(stdin)) {
                printf("\nBye bye.\n");
                break;
            }
        }

        // Check if we exited
        if (strcmp(input, "exit\n") == 0) {
            printf("Bye bye.\n");
            break;
        }

        tokens = tokenize(input);                               // tokenize the input command
        execute_recursive(tokens, 0, count_tokens(tokens) - 1); // execute this command
        strcpy(prev_command, input);                            // Store the command in prev_command (for prev)
        free(tokens);
    }

    return 0;
}

// Handles the cd command
void cd_command(char *directory) {
    if (chdir(directory) != 0) {
        perror("cd");
    }
}

// Handles the help command
void help_command() {
    printf("\'cd\': Change current directory.\n");
    printf("\'source\': Execute commands from a file.\n");
    printf("\'prev\': Repeat the previous command.\n");
    printf("\'help\': Display help for built-in commands.\n");
}

void source_command(char *filename) {
    FILE *file = fopen(filename, "r");
    if (file == NULL) {
        perror("fopen");
        exit(EXIT_FAILURE);
    }

    char *line = NULL;
    size_t len = 0;

    while (getline(&line, &len, file) != -1) {
        // Remove the trailing newline character from the line
        size_t line_length = strlen(line);
        if (line_length > 0 && line[line_length - 1] == '\n') {
            line[line_length - 1] = '\0';
        }

        // Tokenize the line (you'll need to implement this)
        Token *tokens = tokenize(line);

        // Determine the start and end of the tokens for the execute_recursive function
        int start = 0;
        int end = count_tokens(tokens);

        // Execute the command using the execute_recursive function
        execute_recursive(tokens, start, end);
    }

    fclose(file);
    if (line) {
        free(line);
    }
}

void handle_prev() {
    // printf("Previous command: %s", prev_command);
    Token *tokens = tokenize(prev_command);
    execute_recursive(tokens, 0, count_tokens(tokens) - 1);
    free(tokens);
}

Token *extract_tokens(Token *original_tokens, int start, int end) {
    int num_tokens = end - start + 1; // Calculate the number of tokens to extract

    Token *extracted_tokens = malloc((num_tokens + 1) * sizeof(Token)); // Allocate memory for extracted tokens
    if (extracted_tokens == NULL) {
        perror("Memory allocation error");
        exit(EXIT_FAILURE);
    }

    for (int i = start; i <= end; i++) {
        extracted_tokens[i - start] = original_tokens[i]; // Copy tokens from original list to extracted list
    }

    extracted_tokens[num_tokens].value[0] = '\0'; // Add a sentinel token with an empty value to mark the end

    return extracted_tokens;
}

// Executes a pre-existing command using execvp
// Makes a child process so that we don't kill the shell
void execute_preexisting_command(Token *originalTokens, int start, int end) {

    Token *tokens = extract_tokens(originalTokens, start, end);

    pid_t pid = fork();

    if (pid == 0) {
        // This is the child process
        int i = 0;
        char *args[MAX_INPUT_LENGTH]; // Assuming a maximum length for arguments

        while (tokens[i].type == 'W' && tokens[i].value[0] != '\0') {
            args[i] = tokens[i].value;
            i++;
        }

        args[i] = NULL; // Make sure the last element is NULL for execvp

        if (execvp(args[0], args) == -1) {
            fprintf(stderr, "%s: command not found\n", args[0]);
            _exit(1); // Make sure to exit child on execvp failure
        }
    } else if (pid > 0) {
        // This is the parent process
        int status;
        wait(&status); // Wait for the child process to finish
        if (WIFEXITED(status)) {
        } else {
            printf("Child process did not exit normally\n");
        }
    } else {
        perror("fork");
    }
    free(tokens);
}

// finds the target char and returns the index
int find_char(Token *tokens, int start, int end, char *target) {
    for (int i = start; i < end; i++) {
        if (strcmp(tokens[i].value, target) == 0) {
            return i;
        }
    }
    return -1;
}

// counts all the elements in the array of tokens
int count_tokens(Token *tokens) {
    int count = 0;
    while (tokens[count].value[0] != '\0') {
        count++;
    }
    return count;
}

void handle_seq(Token *tokens, int start, int end, int index) {
    // Create a child process to execute the first part of the sequence
    pid_t left_pid = fork();

    if (left_pid == -1) {
        perror("fork");
        exit(EXIT_FAILURE);
    } else if (left_pid == 0) {
        // This is the child process (left part)
        execute_recursive(tokens, start, index - 1);
        exit(0); // Child process exits
    } else {
        // This is the parent process
        int status;
        // Wait for the first child to finish
        waitpid(left_pid, &status, 0);
    }

    // Create a child process to execute the second part of the sequence
    pid_t right_pid = fork();

    if (right_pid == -1) {
        perror("fork");
        exit(EXIT_FAILURE);
    } else if (right_pid == 0) {
        // This is the child process (right part)
        execute_recursive(tokens, index + 1, end);
        exit(0); // Child process exits
    } else {
        // This is the parent process
        int status;
        // Wait for the second child to finish
        waitpid(right_pid, &status, 0);
    }
}

void handle_pipe(Token *tokens, int start, int end, int index) {
    int pipefd[2];
    if (pipe(pipefd) == -1) {
        perror("pipe");
        exit(EXIT_FAILURE);
    }

    pid_t left_pid = fork();

    if (left_pid == -1) {
        perror("fork");
        exit(EXIT_FAILURE);
    } else if (left_pid == 0) {
        // This is the child process (left part)
        close(pipefd[0]);                            // Close read end of the pipe
        dup2(pipefd[1], STDOUT_FILENO);              // Redirect standard output to the write end of the pipe
        close(pipefd[1]);                            // Close write end of the pipe
        execute_recursive(tokens, start, index - 1); // Execute left side
        exit(0);                                     // Child process exits
    } else {
        // This is the parent process
        close(pipefd[1]); // Close write end of the pipe
        int status;
        waitpid(left_pid, &status, 0);
    }

    pid_t right_pid = fork(); // create right process
    if (right_pid == -1) {
        perror("fork");
        exit(EXIT_FAILURE);
    } else if (right_pid == 0) {
        // This is the child process (right part)
        close(pipefd[1]);                          // Close write end of the pipe
        dup2(pipefd[0], STDIN_FILENO);             // Redirect standard input to the read end of the pipe
        close(pipefd[0]);                          // Close read end of the pipe
        execute_recursive(tokens, index + 1, end); // Execute right side
        exit(0);                                   // Child process exits
    } else {
        // This is the parent process
        close(pipefd[0]); // Close read end of the pipe
        int status;
        waitpid(right_pid, &status, 0);
    }
}

void handle_input(Token *tokens, int start, int end, int index) {
    // Create a child process for input redirection
    pid_t input_redirect_pid = fork();
    if (input_redirect_pid == -1) {
        perror("fork");
        exit(EXIT_FAILURE);
    } else if (input_redirect_pid == 0) {
        // This is the child process (input redirection)
        int input_fd = open(tokens[index + 1].value, O_RDONLY);
        if (input_fd == -1) {
            perror("open");
            exit(EXIT_FAILURE);
        }

        chmod(tokens[index + 1].value, 0666); // Read and write permissions for owner, group, and others

        dup2(input_fd, STDIN_FILENO);
        close(input_fd);
        execute_recursive(tokens, start, index - 1);
        exit(0); // Child process exits
    } else {
        // This is the parent process
        int status;
        waitpid(input_redirect_pid, &status, 0);
    }
}

void handle_output(Token *tokens, int start, int end, int index) {
    // Create a child process for output redirection
    pid_t output_redirect_pid = fork();
    if (output_redirect_pid == -1) {
        perror("fork");
        exit(EXIT_FAILURE);
    } else if (output_redirect_pid == 0) {
        // This is the child process (output redirection)
        // int output_fd = open(tokens[output_redirect + 1].value, O_WRONLY | O_CREAT | O_TRUNC, S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH);
        int output_fd = open(tokens[index + 1].value, O_WRONLY | O_CREAT | O_TRUNC);
        if (output_fd == -1) {
            perror("open");
            exit(EXIT_FAILURE);
        }

        // Set the file permissions to be readable and writable by all users
        chmod(tokens[index + 1].value, 0666); // Read and write permissions for owner, group, and others

        dup2(output_fd, STDOUT_FILENO);
        close(output_fd);
        execute_recursive(tokens, start, index - 1);
        exit(0); // Child process exits
    } else {
        // This is the parent process
        int status;
        waitpid(output_redirect_pid, &status, 0);
    }
}

void execute_recursive(Token *tokens, int start, int end) {

    // char* input_file = NULL;
    // char* output_file = NULL;
    int seq_result = find_char(tokens, start, end, ";");
    int pipe_result = find_char(tokens, start, end, "|");
    int input_redirect = find_char(tokens, start, end, "<");
    int output_redirect = find_char(tokens, start, end, ">");

    // If a semicolon was found
    if (seq_result != -1) {
        handle_seq(tokens, start, end, seq_result);
    } else if (pipe_result != -1) {
        handle_pipe(tokens, start, end, pipe_result);
    } else if (input_redirect != -1) {
        handle_input(tokens, start, end, input_redirect);
    } else if (output_redirect != -1) {
        handle_output(tokens, start, end, output_redirect);
    } else {
        int i;
        for (i = start; i <= end; i++) {

            if (strcmp(tokens[i].value, "cd") == 0) {
                cd_command(tokens[i + 1].value);
                return;
            } else if (strcmp(tokens[i].value, "source") == 0) {
                source_command(tokens[i + 1].value);
                return;
            } else if (strcmp(tokens[i].value, "prev") == 0) {
                handle_prev();
                return;
            } else if (strcmp(tokens[i].value, "help") == 0) {
                help_command();
                return;
            }
            execute_preexisting_command(tokens, start, end);
            return;
        }
    }
}
