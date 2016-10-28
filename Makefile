HEADERS=src/light.hpp src/node.hpp
SOURCES=src/main.cpp src/light.cpp src/node.cpp
CXXFLAGS+=-std=c++11 -pthread -Isrc
CXX=g++

all: $(HEADERS) $(SOURCES)
	$(CXX) $(SOURCES) $(CXXFLAGS) -o bin/controller

run: all
	bin/controller

clean: 
	rm bin/controller
