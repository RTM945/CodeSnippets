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
    SimpleSocket::accept(simpleSocket);
}

void RTM::Acceptor::send(const std::string& msg) 
{
    SimpleSocket::send(msg);
}

std::string RTM::Acceptor::recv() 
{
    return SimpleSocket::recv();
}