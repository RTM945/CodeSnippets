#include "SimpleSocket.hpp"

RTM::SimpleSocket::SimpleSocket() : sock(-1){}

RTM::SimpleSocket::~SimpleSocket()
{
    if (sock != -1)
    {
        close(sock);
    }
}

void RTM::SimpleSocket::create()
{
    sock = ::socket(AF_INET, SOCK_STREAM, 0);
    if (sock < 0) 
    {
        perror("failed to create sock");
        exit(EXIT_FAILURE);
    }
    int on = 1;
    if (setsockopt(sock, SOL_SOCKET, SO_REUSEADDR, (const char*) &on, sizeof(on)) == -1)
    {
        perror("failed to setsockopt");
        exit(EXIT_FAILURE);
    }
    
}

void RTM::SimpleSocket::bind(int port)
{
    addr.sin_family = AF_INET;
    addr.sin_addr.s_addr = INADDR_ANY;
    addr.sin_port = htons(port);
    int res = ::bind(sock, (struct sockaddr*)&addr, sizeof(addr));
    if (res < 0)
    {
        perror("failed to bind port");
        exit(EXIT_FAILURE);
    }
}

void RTM::SimpleSocket::listen(int backlog)
{
    int res = ::listen(sock, backlog);
    if (res < 0)
    {
        perror("failed to listen");
        exit(EXIT_FAILURE);
    }
}

void RTM::SimpleSocket::connect(std::string host, int port)
{
    addr.sin_family = AF_INET;
    addr.sin_addr.s_addr = INADDR_ANY;
    addr.sin_port = htons(port);
    int res = inet_pton(AF_INET, host.c_str(), &addr.sin_addr);
    if (errno == EAFNOSUPPORT)
    {
        perror("failed to connect");
        exit(EAFNOSUPPORT);
    }
    res = ::connect(sock, (sockaddr*)&addr, sizeof(addr));
    if (res < 0)
    {
        perror("failed to connect");
        exit(EXIT_FAILURE);
    }
}

bool RTM::SimpleSocket::accept(SimpleSocket& connector)
{
    std::cout<<"my sock=" << sock << std::endl;
    int addr_length = sizeof(addr);
    connector.sock = ::accept(sock, (sockaddr*)&addr, (socklen_t*)&addr_length);
    std::cout<<"new connector sock=" << connector.sock << std::endl;
    if (connector.sock < 0) 
    {
        return false;
    }
    return true;
}

bool RTM::SimpleSocket::send(const std::string msg)
{
    int res = ::send(sock, msg.c_str(), msg.size(), 0);
    if (res < 0)
    {
        return false;
    }
    return true;
}

int RTM::SimpleSocket::recv(std::string& msg)
{
    char buf[256];
    int res = ::recv(sock, buf, 255, 0);
    if (res < 0)
    {
        std::cout << "recv fail from sock"<< sock << "error = " << errno << "\n";
        
    } 
    else
    {
        msg = buf;
    }
    return res;
}