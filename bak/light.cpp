#include "light.hpp"
#include "node.hpp"

light::light(Node pos) : position(pos.X,pos.Y) {
	
}

light::~light() {}