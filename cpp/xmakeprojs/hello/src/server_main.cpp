#include "Acceptor.hpp"

int main(int argc, char** argv)
{
    RTM::Acceptor acceptor(8888, 10);
    while (true)
    {
        RTM::SimpleSocket connector;
        acceptor.accept(connector);
        std::string msg;
        connector.recv(msg);
        connector.send(msg);
    }
    return 0;
}