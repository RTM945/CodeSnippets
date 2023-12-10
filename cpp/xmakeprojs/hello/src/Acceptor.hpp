#ifndef Acceptor_hpp
#define Acceptor_hpp

#include "SimpleSocket.hpp"

namespace RTM
{
class Acceptor : private SimpleSocket
{
public:
    Acceptor(int port, int backlog);
    virtual ~Acceptor();
    void accept(SimpleSocket&);
    void send(const std::string&);
    void recv(std::string&);
};

};

#endif