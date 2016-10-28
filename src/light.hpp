#pragma once

#include "node.hpp"

class light //: kobject
{
public:
	light(Node);
	~light();
private:
	Node position;
};