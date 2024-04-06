#include "olc_net.h"

enum class CustomMsgTypes : uint32_t
{
    ServerAccept,
    ServerDeny,
    ServerPing,
    MessageAll,
    ServerMessage
};

class CustomClient : public olc::net::client_interface<CustomMsgTypes>
{

};

int main()
{
    CustomClient c;
    c.Connect("127.0.0.1", 60000);

    

    bool bQuit = false;
    while (!bQuit)
    {
        if (c.isConnected())
        {
            if (!c.Incoming().empty()) 
            {

            }
        } 
        else 
        {
            std::cout << "Server Down.\n";
            bQuit = true;
        }
    }

    return 0;
}