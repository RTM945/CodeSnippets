#ifndef Connector_hpp
#define Connector_hpp

#include "SimpleSocket.hpp"

namespace RTM
{
class Connector : private SimpleSocket
{
public:
    Connector(std::string host, int port);
    virtual ~Connector(){};
    void send(const std::string&);
    void recv(std::string&);
};

};

#endif