#pragma once

typedef int coord_t;

class Node {
public:
	Node(coord_t,coord_t);
	~Node();
// private:
	coord_t X,Y;
};
