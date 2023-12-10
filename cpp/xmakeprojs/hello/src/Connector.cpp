#include "Connector.hpp"

RTM::Connector::Connector(std::string host, int port) {
    SimpleSocket::create();
    SimpleSocket::connect(host, port);
}

void RTM::Connector::send(const std::string& msg) 
{
    SimpleSocket::send(msg);
}

void RTM::Connector::recv(std::string& msg) 
{
    SimpleSocket::recv(msg);
}