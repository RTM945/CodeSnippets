#include "SimpleSocket.hpp"

RTM::Acceptor::Acceptor(int port, int backlog) : SimpleSocket()
{
    SimpleSocket::create();
    SimpleSocket::bind(port);
    SimpleSocket::listen(backlog);
}

bool RTM::Acceptor::accept(Connector& connector)
{
    SimpleSocket& simpleSocket = connector;
    return SimpleSocket::accept(simpleSocket);
}