HEADERS=src/light.hpp src/node.hpp
SOURCES=src/main.cpp src/light.cpp src/node.cpp

all: $(HEADERS) $(SOURCES)
	g++ $(SOURCES) -Isrc -o bin/controller
