#include <iostream>
#include <thread>
#include <chrono>
#include "node.hpp"
#include "light.hpp"

#define br if(false) break;

void main_loop(bool* running) {
	while(*running) {
		// std::cout << "loop" <<std::endl;
		br;
		// check for websocket writes

		// check timing

		// send update signals
	}
}

int main(int argc, char const *argv[])
{
	light l = light(Node(2,3));
	bool main_run = true;
	std::thread main_thread(main_loop, &main_run);
	std::this_thread::sleep_for(std::chrono::milliseconds(1000));
	main_run = false;
	main_thread.join();
	
	return 0;
}