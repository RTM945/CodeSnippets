#ifndef SimpleSocket_hpp
#define SimpleSocket_hpp

#include <iostream>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <errno.h>

namespace RTM
{
class SimpleSocket
{
private:
    int sock = -1;
    sockaddr_in addr;
public:
    SimpleSocket();
    virtual ~SimpleSocket();
    void create();
    void bind(int port);
    void listen(int backlog);
    void connect(std::string host, int port);
    bool accept(SimpleSocket&);
    bool send(std::string);
    std::string recv();
};

class Connector : public SimpleSocket
{
public:
    Connector();
    void connect(std::string host, int port);
};

class Acceptor : public SimpleSocket
{
public:
    Acceptor(int port, int backlog);
    bool accept(Connector& connector);
};

};

#endif