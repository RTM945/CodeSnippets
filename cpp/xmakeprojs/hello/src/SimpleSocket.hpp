#ifndef SimpleSocket_hpp
#define SimpleSocket_hpp

#include <iostream>
#include <string>
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
    bool send(const std::string);
    int recv(std::string&);
};
}

#endif
// wtf?
// https://tldp.org/LDP/LG/issue74/misc/tougher/