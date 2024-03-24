#include <iostream>
#include "olc_net.h"

enum class CustomMsgTypes : uint32_t
{
    ServerAccept,
    ServerDeny,
    ServerPing,
    MessageAll,
    ServerMessage
};

class CustomServer : public olc::net::server_interface<CustomMsgTypes>
{
public:
    CustomServer(uint16_t nPort) : olc::net::server_interface<CustomMsgTypes>(nPort)
    {

    }

protected:
    virtual bool OnClientConnect(std::shared_ptr<olc::net::connection<CustomMsgTypes>> client)
    {
        return true;
    }


    virtual void OnClientDisConnect(std::shared_ptr<olc::net::connection<CustomMsgTypes>> client)
    {
        
    }

    virtual void OnMessage(std::shared_ptr<olc::net::connection<CustomMsgTypes>> client, olc::net::message<CustomMsgTypes>& msg)
    {
        
    }
};

int main() 
{
    CustomServer server(60000);
    server.Start();

    while (1) {
        server.Update();
    }
    return 0;
}