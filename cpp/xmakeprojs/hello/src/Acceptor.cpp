#include "Acceptor.hpp"

RTM::Acceptor::Acceptor(int port, int backlog)
{
    SimpleSocket::create();
    SimpleSocket::bind(port);
    SimpleSocket::listen(backlog);
}

RTM::Acceptor::~Acceptor()
{
}

void RTM::Acceptor::accept(SimpleSocket& connector)
{
    SimpleSocket::accept(connector);
}

void RTM::Acceptor::send(const std::string& msg) 
{
    SimpleSocket::send(msg);
}

void RTM::Acceptor::recv(std::string& msg) 
{
    SimpleSocket::recv(msg);
}