#include "Connector.hpp"
#include <iostream>
#include <string>

int main(int argc, char** argv)
{
    RTM::Connector connector("localhost", 8888);
    std::string reply;
    connector.send("Hello");
    connector.recv(reply);
    std::cout << reply << std::endl;
    return 0;
}