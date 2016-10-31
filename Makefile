HEADERS=src/light.hpp src/node.hpp
SOURCES=src/main.cpp src/light.cpp src/node.cpp
CXXFLAGS+=-std=c++11 -pthread -Isrc
CXX=g++
OUT=bin/controller

all: $(OUT)

$(OUT): $(HEADERS) $(SOURCES)
	$(CXX) $(SOURCES) $(CXXFLAGS) -o $(OUT)

run: $(OUT)
	$(OUT)

clean: 
	rm $(OUT) 
