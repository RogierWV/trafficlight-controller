HEADERS=src/light.hpp src/node.hpp
SOURCES=src/main.cpp src/light.cpp src/node.cpp
CXXFLAGS+=-std=c++11 -pthread -Isrc

all: $(HEADERS) $(SOURCES)
	g++ $(SOURCES) $(CXXFLAGS) -o bin/controller

run: all
	bin/controller

clean: 
	rm bin/*
