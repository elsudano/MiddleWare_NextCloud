###############################################################################
# Variable que guarda los ficheros que hay que compilar
SRC_DIR = src

###############################################################################
# Extraemos los nombres de los ficheros para los binarios
BIN_DIR = bin

###############################################################################
# Variables que no cambian para cualquier compilación
# Extención de los ficheros fuente
EXT = go # pueden ser go

###############################################################################
# Nombre del programa Principal
NAME = main

###############################################################################
# Comando para eliminar los ficheros
RM = rm -Rf

###############################################################################
# Indica cual es el compilador para los fuentes
COMPILER = go

.PHONY: help
.SECONDARY:

all: help

makedir:
	mkdir ./$(BIN_DIR)

run:
	$(COMPILER) run $(SRC_DIR)/$(NAME).$(EXT)

build: clean makedir
	$(COMPILER) build -o $(BIN_DIR)/$(NAME) $(SRC_DIR)/$(NAME).$(EXT)

docu:
	doxygen ./doc/doxys/dox_config

clean:
	$(RM) $(BIN_DIR) doc/html/* doc/latex/*

touch:
	touch $(SRC_DIR)/*

help:
	@echo "Available targets:"
	@echo "- run       Run Application whitout compile"
	@echo "- build       Generate bin file on directory $(BIN_DIR)"
	@echo "- clean       Clean up the source directory $(SRC_DIR) and bin directory $(BIN_DIR)"
	@echo "- test        Run tests"
	@echo "- help        This info"
	@echo
	@echo "Available variables:"
	@echo "- SRC_DIR      default: $(SRC_DIR)"
	@echo
