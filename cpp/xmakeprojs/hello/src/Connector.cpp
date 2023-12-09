#include "SimpleSocket.hpp"

RTM::Connector::Connector() : SimpleSocket() {}

void RTM::Connector::connect(std::string host, int port)
{   
    SimpleSocket::create();
    SimpleSocket::connect(host, port);
}
