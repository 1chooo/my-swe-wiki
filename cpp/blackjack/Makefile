CC          = gcc
LD          = gcc
CFLAG       = -Wall
PROG_NAME   = main

SRC_DIR     = ./src
INC_DIR     = ./inc
OBJ_DIR     = ./obj
BIN_DIR     = ./bin
SRC_LIST = $(wildcard $(SRC_DIR)/*.c)
OBJ_LIST = $(wildcard $(OBJ_DIR)/*.o)

.PHONY: all clean $(PROG_NAME) compile

all: $(PROG_NAME)


compile:
	$(CC) $(CFLAG) $(SRC_LIST) -I $(INC_DIR) -c
# gcc -Wall src/main.c src/sta ... -I ./
copy: compile
	@mv $(notdir $(SRC_LIST:.c=.o)) $(OBJ_DIR)/

$(PROG_NAME): copy
	$(LD) $(OBJ_LIST) -o $(BIN_DIR)/$@

clean:
	rm -f $(BIN_DIR)/$(PROG_NAME) $(OBJ_DIR)/*.o